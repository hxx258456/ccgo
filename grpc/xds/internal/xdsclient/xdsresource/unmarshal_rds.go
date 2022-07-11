/*
 *
 * Copyright 2021 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package xdsresource

import (
	"fmt"
	"regexp"
	"strings"
	"time"

	"github.com/golang/protobuf/proto"
	v3routepb "github.com/hxx258456/ccgo/go-control-plane/envoy/config/route/v3"
	v3typepb "github.com/hxx258456/ccgo/go-control-plane/envoy/type/v3"
	"github.com/hxx258456/ccgo/grpc/codes"
	"github.com/hxx258456/ccgo/grpc/internal/envconfig"
	"github.com/hxx258456/ccgo/grpc/internal/grpclog"
	"github.com/hxx258456/ccgo/grpc/internal/pretty"
	"github.com/hxx258456/ccgo/grpc/xds/internal/clusterspecifier"
	"github.com/hxx258456/ccgo/grpc/xds/internal/xdsclient/xdsresource/version"
	"google.golang.org/protobuf/types/known/anypb"
)

// UnmarshalRouteConfig processes resources received in an RDS response,
// validates them, and transforms them into a native struct which contains only
// fields we are interested in. The provided hostname determines the route
// configuration resources of interest.
func UnmarshalRouteConfig(opts *UnmarshalOptions) (map[string]RouteConfigUpdateErrTuple, UpdateMetadata, error) {
	update := make(map[string]RouteConfigUpdateErrTuple)
	md, err := processAllResources(opts, update)
	return update, md, err
}

func unmarshalRouteConfigResource(r *anypb.Any, logger *grpclog.PrefixLogger) (string, RouteConfigUpdate, error) {
	if !IsRouteConfigResource(r.GetTypeUrl()) {
		return "", RouteConfigUpdate{}, fmt.Errorf("unexpected resource type: %q ", r.GetTypeUrl())
	}
	rc := &v3routepb.RouteConfiguration{}
	if err := proto.Unmarshal(r.GetValue(), rc); err != nil {
		return "", RouteConfigUpdate{}, fmt.Errorf("failed to unmarshal resource: %v", err)
	}
	logger.Infof("Resource with name: %v, type: %T, contains: %v.", rc.GetName(), rc, pretty.ToJSON(rc))

	// TODO: Pass version.TransportAPI instead of relying upon the type URL
	v2 := r.GetTypeUrl() == version.V2RouteConfigURL
	u, err := generateRDSUpdateFromRouteConfiguration(rc, logger, v2)
	if err != nil {
		return rc.GetName(), RouteConfigUpdate{}, err
	}
	u.Raw = r
	return rc.GetName(), u, nil
}

// generateRDSUpdateFromRouteConfiguration checks if the provided
// RouteConfiguration meets the expected criteria. If so, it returns a
// RouteConfigUpdate with nil error.
//
// A RouteConfiguration resource is considered valid when only if it contains a
// VirtualHost whose domain field matches the server name from the URI passed
// to the gRPC channel, and it contains a clusterName or a weighted cluster.
//
// The RouteConfiguration includes a list of virtualHosts, which may have zero
// or more elements. We are interested in the element whose domains field
// matches the server name specified in the "xds:" URI. The only field in the
// VirtualHost proto that the we are interested in is the list of routes. We
// only look at the last route in the list (the default route), whose match
// field must be empty and whose route field must be set.  Inside that route
// message, the cluster field will contain the clusterName or weighted clusters
// we are looking for.
func generateRDSUpdateFromRouteConfiguration(rc *v3routepb.RouteConfiguration, logger *grpclog.PrefixLogger, v2 bool) (RouteConfigUpdate, error) {
	vhs := make([]*VirtualHost, 0, len(rc.GetVirtualHosts()))
	csps := make(map[string]clusterspecifier.BalancerConfig)
	if envconfig.XDSRLS {
		var err error
		csps, err = processClusterSpecifierPlugins(rc.ClusterSpecifierPlugins)
		if err != nil {
			return RouteConfigUpdate{}, fmt.Errorf("received route is invalid %v", err)
		}
	}
	// cspNames represents all the cluster specifiers referenced by Route
	// Actions - any cluster specifiers not referenced by a Route Action can be
	// ignored and not emitted by the xdsclient.
	var cspNames = make(map[string]bool)
	for _, vh := range rc.GetVirtualHosts() {
		routes, cspNs, err := routesProtoToSlice(vh.Routes, csps, logger, v2)
		if err != nil {
			return RouteConfigUpdate{}, fmt.Errorf("received route is invalid: %v", err)
		}
		for n := range cspNs {
			cspNames[n] = true
		}
		rc, err := generateRetryConfig(vh.GetRetryPolicy())
		if err != nil {
			return RouteConfigUpdate{}, fmt.Errorf("received route is invalid: %v", err)
		}
		vhOut := &VirtualHost{
			Domains:     vh.GetDomains(),
			Routes:      routes,
			RetryConfig: rc,
		}
		if !v2 {
			cfgs, err := processHTTPFilterOverrides(vh.GetTypedPerFilterConfig())
			if err != nil {
				return RouteConfigUpdate{}, fmt.Errorf("virtual host %+v: %v", vh, err)
			}
			vhOut.HTTPFilterConfigOverride = cfgs
		}
		vhs = append(vhs, vhOut)
	}

	// "For any entry in the RouteConfiguration.cluster_specifier_plugins not
	// referenced by an enclosed ActionType's cluster_specifier_plugin, the xDS
	// client should not provide it to its consumers." - RLS in xDS Design
	for name := range csps {
		if !cspNames[name] {
			delete(csps, name)
		}
	}

	return RouteConfigUpdate{VirtualHosts: vhs, ClusterSpecifierPlugins: csps}, nil
}

func processClusterSpecifierPlugins(csps []*v3routepb.ClusterSpecifierPlugin) (map[string]clusterspecifier.BalancerConfig, error) {
	cspCfgs := make(map[string]clusterspecifier.BalancerConfig)
	// "The xDS client will inspect all elements of the
	// cluster_specifier_plugins field looking up a plugin based on the
	// extension.typed_config of each." - RLS in xDS design
	for _, csp := range csps {
		cs := clusterspecifier.Get(csp.GetExtension().GetTypedConfig().GetTypeUrl())
		if cs == nil {
			// "If no plugin is registered for it, the resource will be NACKed."
			// - RLS in xDS design
			return nil, fmt.Errorf("cluster specifier %q of type %q was not found", csp.GetExtension().GetName(), csp.GetExtension().GetTypedConfig().GetTypeUrl())
		}
		lbCfg, err := cs.ParseClusterSpecifierConfig(csp.GetExtension().GetTypedConfig())
		if err != nil {
			// "If a plugin is found, the value of the typed_config field will
			// be passed to it's conversion method, and if an error is
			// encountered, the resource will be NACKED." - RLS in xDS design
			return nil, fmt.Errorf("error: %q parsing config %q for cluster specifier %q of type %q", err, csp.GetExtension().GetTypedConfig(), csp.GetExtension().GetName(), csp.GetExtension().GetTypedConfig().GetTypeUrl())
		}
		// "If all cluster specifiers are valid, the xDS client will store the
		// configurations in a map keyed by the name of the extension instance." -
		// RLS in xDS Design
		cspCfgs[csp.GetExtension().GetName()] = lbCfg
	}
	return cspCfgs, nil
}

func generateRetryConfig(rp *v3routepb.RetryPolicy) (*RetryConfig, error) {
	if rp == nil {
		return nil, nil
	}

	cfg := &RetryConfig{RetryOn: make(map[codes.Code]bool)}
	for _, s := range strings.Split(rp.GetRetryOn(), ",") {
		switch strings.TrimSpace(strings.ToLower(s)) {
		case "cancelled":
			cfg.RetryOn[codes.Canceled] = true
		case "deadline-exceeded":
			cfg.RetryOn[codes.DeadlineExceeded] = true
		case "internal":
			cfg.RetryOn[codes.Internal] = true
		case "resource-exhausted":
			cfg.RetryOn[codes.ResourceExhausted] = true
		case "unavailable":
			cfg.RetryOn[codes.Unavailable] = true
		}
	}

	if rp.NumRetries == nil {
		cfg.NumRetries = 1
	} else {
		cfg.NumRetries = rp.GetNumRetries().Value
		if cfg.NumRetries < 1 {
			return nil, fmt.Errorf("retry_policy.num_retries = %v; must be >= 1", cfg.NumRetries)
		}
	}

	backoff := rp.GetRetryBackOff()
	if backoff == nil {
		cfg.RetryBackoff.BaseInterval = 25 * time.Millisecond
	} else {
		cfg.RetryBackoff.BaseInterval = backoff.GetBaseInterval().AsDuration()
		if cfg.RetryBackoff.BaseInterval <= 0 {
			return nil, fmt.Errorf("retry_policy.base_interval = %v; must be > 0", cfg.RetryBackoff.BaseInterval)
		}
	}
	if max := backoff.GetMaxInterval(); max == nil {
		cfg.RetryBackoff.MaxInterval = 10 * cfg.RetryBackoff.BaseInterval
	} else {
		cfg.RetryBackoff.MaxInterval = max.AsDuration()
		if cfg.RetryBackoff.MaxInterval <= 0 {
			return nil, fmt.Errorf("retry_policy.max_interval = %v; must be > 0", cfg.RetryBackoff.MaxInterval)
		}
	}

	if len(cfg.RetryOn) == 0 {
		return &RetryConfig{}, nil
	}
	return cfg, nil
}

func routesProtoToSlice(routes []*v3routepb.Route, csps map[string]clusterspecifier.BalancerConfig, logger *grpclog.PrefixLogger, v2 bool) ([]*Route, map[string]bool, error) {
	var routesRet []*Route
	var cspNames = make(map[string]bool)
	for _, r := range routes {
		match := r.GetMatch()
		if match == nil {
			return nil, nil, fmt.Errorf("route %+v doesn't have a match", r)
		}

		if len(match.GetQueryParameters()) != 0 {
			// Ignore route with query parameters.
			logger.Warningf("route %+v has query parameter matchers, the route will be ignored", r)
			continue
		}

		pathSp := match.GetPathSpecifier()
		if pathSp == nil {
			return nil, nil, fmt.Errorf("route %+v doesn't have a path specifier", r)
		}

		var route Route
		switch pt := pathSp.(type) {
		case *v3routepb.RouteMatch_Prefix:
			route.Prefix = &pt.Prefix
		case *v3routepb.RouteMatch_Path:
			route.Path = &pt.Path
		case *v3routepb.RouteMatch_SafeRegex:
			regex := pt.SafeRegex.GetRegex()
			re, err := regexp.Compile(regex)
			if err != nil {
				return nil, nil, fmt.Errorf("route %+v contains an invalid regex %q", r, regex)
			}
			route.Regex = re
		default:
			return nil, nil, fmt.Errorf("route %+v has an unrecognized path specifier: %+v", r, pt)
		}

		if caseSensitive := match.GetCaseSensitive(); caseSensitive != nil {
			route.CaseInsensitive = !caseSensitive.Value
		}

		for _, h := range match.GetHeaders() {
			var header HeaderMatcher
			switch ht := h.GetHeaderMatchSpecifier().(type) {
			case *v3routepb.HeaderMatcher_ExactMatch:
				header.ExactMatch = &ht.ExactMatch
			case *v3routepb.HeaderMatcher_SafeRegexMatch:
				regex := ht.SafeRegexMatch.GetRegex()
				re, err := regexp.Compile(regex)
				if err != nil {
					return nil, nil, fmt.Errorf("route %+v contains an invalid regex %q", r, regex)
				}
				header.RegexMatch = re
			case *v3routepb.HeaderMatcher_RangeMatch:
				header.RangeMatch = &Int64Range{
					Start: ht.RangeMatch.Start,
					End:   ht.RangeMatch.End,
				}
			case *v3routepb.HeaderMatcher_PresentMatch:
				header.PresentMatch = &ht.PresentMatch
			case *v3routepb.HeaderMatcher_PrefixMatch:
				header.PrefixMatch = &ht.PrefixMatch
			case *v3routepb.HeaderMatcher_SuffixMatch:
				header.SuffixMatch = &ht.SuffixMatch
			default:
				return nil, nil, fmt.Errorf("route %+v has an unrecognized header matcher: %+v", r, ht)
			}
			header.Name = h.GetName()
			invert := h.GetInvertMatch()
			header.InvertMatch = &invert
			route.Headers = append(route.Headers, &header)
		}

		if fr := match.GetRuntimeFraction(); fr != nil {
			d := fr.GetDefaultValue()
			n := d.GetNumerator()
			switch d.GetDenominator() {
			case v3typepb.FractionalPercent_HUNDRED:
				n *= 10000
			case v3typepb.FractionalPercent_TEN_THOUSAND:
				n *= 100
			case v3typepb.FractionalPercent_MILLION:
			}
			route.Fraction = &n
		}

		switch r.GetAction().(type) {
		case *v3routepb.Route_Route:
			route.WeightedClusters = make(map[string]WeightedCluster)
			action := r.GetRoute()

			// Hash Policies are only applicable for a Ring Hash LB.
			if envconfig.XDSRingHash {
				hp, err := hashPoliciesProtoToSlice(action.HashPolicy, logger)
				if err != nil {
					return nil, nil, err
				}
				route.HashPolicies = hp
			}

			switch a := action.GetClusterSpecifier().(type) {
			case *v3routepb.RouteAction_Cluster:
				route.WeightedClusters[a.Cluster] = WeightedCluster{Weight: 1}
			case *v3routepb.RouteAction_WeightedClusters:
				wcs := a.WeightedClusters
				var totalWeight uint32
				for _, c := range wcs.Clusters {
					w := c.GetWeight().GetValue()
					if w == 0 {
						continue
					}
					wc := WeightedCluster{Weight: w}
					if !v2 {
						cfgs, err := processHTTPFilterOverrides(c.GetTypedPerFilterConfig())
						if err != nil {
							return nil, nil, fmt.Errorf("route %+v, action %+v: %v", r, a, err)
						}
						wc.HTTPFilterConfigOverride = cfgs
					}
					route.WeightedClusters[c.GetName()] = wc
					totalWeight += w
				}
				// envoy xds doc
				// default TotalWeight https://www.envoyproxy.io/docs/envoy/latest/api-v3/config/route/v3/route_components.proto.html#envoy-v3-api-field-config-route-v3-weightedcluster-total-weight
				wantTotalWeight := uint32(100)
				if tw := wcs.GetTotalWeight(); tw != nil {
					wantTotalWeight = tw.GetValue()
				}
				if totalWeight != wantTotalWeight {
					return nil, nil, fmt.Errorf("route %+v, action %+v, weights of clusters do not add up to total total weight, got: %v, expected total weight from response: %v", r, a, totalWeight, wantTotalWeight)
				}
				if totalWeight == 0 {
					return nil, nil, fmt.Errorf("route %+v, action %+v, has no valid cluster in WeightedCluster action", r, a)
				}
			case *v3routepb.RouteAction_ClusterHeader:
				continue
			case *v3routepb.RouteAction_ClusterSpecifierPlugin:
				if !envconfig.XDSRLS {
					return nil, nil, fmt.Errorf("route %+v, has an unknown ClusterSpecifier: %+v", r, a)
				}
				if _, ok := csps[a.ClusterSpecifierPlugin]; !ok {
					// "When processing RouteActions, if any action includes a
					// cluster_specifier_plugin value that is not in
					// RouteConfiguration.cluster_specifier_plugins, the
					// resource will be NACKed." - RLS in xDS design
					return nil, nil, fmt.Errorf("route %+v, action %+v, specifies a cluster specifier plugin %+v that is not in Route Configuration", r, a, a.ClusterSpecifierPlugin)
				}
				cspNames[a.ClusterSpecifierPlugin] = true
				route.ClusterSpecifierPlugin = a.ClusterSpecifierPlugin
			default:
				return nil, nil, fmt.Errorf("route %+v, has an unknown ClusterSpecifier: %+v", r, a)
			}

			msd := action.GetMaxStreamDuration()
			// Prefer grpc_timeout_header_max, if set.
			dur := msd.GetGrpcTimeoutHeaderMax()
			if dur == nil {
				dur = msd.GetMaxStreamDuration()
			}
			if dur != nil {
				d := dur.AsDuration()
				route.MaxStreamDuration = &d
			}

			var err error
			route.RetryConfig, err = generateRetryConfig(action.GetRetryPolicy())
			if err != nil {
				return nil, nil, fmt.Errorf("route %+v, action %+v: %v", r, action, err)
			}

			route.ActionType = RouteActionRoute

		case *v3routepb.Route_NonForwardingAction:
			// Expected to be used on server side.
			route.ActionType = RouteActionNonForwardingAction
		default:
			route.ActionType = RouteActionUnsupported
		}

		if !v2 {
			cfgs, err := processHTTPFilterOverrides(r.GetTypedPerFilterConfig())
			if err != nil {
				return nil, nil, fmt.Errorf("route %+v: %v", r, err)
			}
			route.HTTPFilterConfigOverride = cfgs
		}
		routesRet = append(routesRet, &route)
	}
	return routesRet, cspNames, nil
}

func hashPoliciesProtoToSlice(policies []*v3routepb.RouteAction_HashPolicy, logger *grpclog.PrefixLogger) ([]*HashPolicy, error) {
	var hashPoliciesRet []*HashPolicy
	for _, p := range policies {
		policy := HashPolicy{Terminal: p.Terminal}
		switch p.GetPolicySpecifier().(type) {
		case *v3routepb.RouteAction_HashPolicy_Header_:
			policy.HashPolicyType = HashPolicyTypeHeader
			policy.HeaderName = p.GetHeader().GetHeaderName()
			if rr := p.GetHeader().GetRegexRewrite(); rr != nil {
				regex := rr.GetPattern().GetRegex()
				re, err := regexp.Compile(regex)
				if err != nil {
					return nil, fmt.Errorf("hash policy %+v contains an invalid regex %q", p, regex)
				}
				policy.Regex = re
				policy.RegexSubstitution = rr.GetSubstitution()
			}
		case *v3routepb.RouteAction_HashPolicy_FilterState_:
			if p.GetFilterState().GetKey() != "io.grpc.channel_id" {
				logger.Infof("hash policy %+v contains an invalid key for filter state policy %q", p, p.GetFilterState().GetKey())
				continue
			}
			policy.HashPolicyType = HashPolicyTypeChannelID
		default:
			logger.Infof("hash policy %T is an unsupported hash policy", p.GetPolicySpecifier())
			continue
		}

		hashPoliciesRet = append(hashPoliciesRet, &policy)
	}
	return hashPoliciesRet, nil
}
