//go:build !386
// +build !386

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

package xds_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/google/uuid"
	v3clusterpb "github.com/hxx258456/ccgo/go-control-plane/envoy/config/cluster/v3"
	v3endpointpb "github.com/hxx258456/ccgo/go-control-plane/envoy/config/endpoint/v3"
	v3listenerpb "github.com/hxx258456/ccgo/go-control-plane/envoy/config/listener/v3"
	v3routepb "github.com/hxx258456/ccgo/go-control-plane/envoy/config/route/v3"
	grpc "github.com/hxx258456/ccgo/grpc"
	"github.com/hxx258456/ccgo/grpc/credentials/insecure"
	"github.com/hxx258456/ccgo/grpc/internal/envconfig"
	xdsinternal "github.com/hxx258456/ccgo/grpc/internal/xds"
	testpb "github.com/hxx258456/ccgo/grpc/test/grpc_testing"
	"github.com/hxx258456/ccgo/grpc/xds"
	"github.com/hxx258456/ccgo/grpc/xds/internal/testutils"
	"github.com/hxx258456/ccgo/grpc/xds/internal/testutils/e2e"
	"github.com/hxx258456/ccgo/grpc/xds/internal/xdsclient/xdsresource"
)

// TestClientSideFederation tests that federation is supported.
//
// In this test, some xDS responses contain resource names in another authority
// (in the new resource name style):
// - LDS: old style, no authority (default authority)
// - RDS: new style, in a different authority
// - CDS: old style, no authority (default authority)
// - EDS: new style, in a different authority
func (s) TestClientSideFederation(t *testing.T) {
	oldXDSFederation := envconfig.XDSFederation
	envconfig.XDSFederation = true
	defer func() { envconfig.XDSFederation = oldXDSFederation }()

	// Start a management server as the default authority.
	serverDefaultAuth, err := e2e.StartManagementServer()
	if err != nil {
		t.Fatalf("Failed to spin up the xDS management server: %v", err)
	}
	t.Cleanup(serverDefaultAuth.Stop)

	// Start another management server as the other authority.
	const nonDefaultAuth = "non-default-auth"
	serverAnotherAuth, err := e2e.StartManagementServer()
	if err != nil {
		t.Fatalf("Failed to spin up the xDS management server: %v", err)
	}
	t.Cleanup(serverAnotherAuth.Stop)

	// Create a bootstrap file in a temporary directory.
	nodeID := uuid.New().String()
	bootstrapContents, err := xdsinternal.BootstrapContents(xdsinternal.BootstrapOptions{
		Version:                            xdsinternal.TransportV3,
		NodeID:                             nodeID,
		ServerURI:                          serverDefaultAuth.Address,
		ServerListenerResourceNameTemplate: e2e.ServerListenerResourceNameTemplate,
		// Specify the address of the non-default authority.
		Authorities: map[string]string{nonDefaultAuth: serverAnotherAuth.Address},
	})
	if err != nil {
		t.Fatalf("Failed to create bootstrap file: %v", err)
	}

	resolver, err := xds.NewXDSResolverWithConfigForTesting(bootstrapContents)
	if err != nil {
		t.Fatalf("Failed to create xDS resolver for testing: %v", err)
	}
	port, cleanup := clientSetup(t, &testService{})
	defer cleanup()

	const serviceName = "my-service-client-side-xds"
	// LDS is old style name.
	ldsName := serviceName
	// RDS is new style, with the non default authority.
	rdsName := testutils.BuildResourceName(xdsresource.RouteConfigResource, nonDefaultAuth, "route-"+serviceName, nil)
	// CDS is old style name.
	cdsName := "cluster-" + serviceName
	// EDS is new style, with the non default authority.
	edsName := testutils.BuildResourceName(xdsresource.EndpointsResource, nonDefaultAuth, "endpoints-"+serviceName, nil)

	// Split resources, put LDS/CDS in the default authority, and put RDS/EDS in
	// the other authority.
	resourcesDefault := e2e.UpdateOptions{
		NodeID: nodeID,
		// This has only LDS and CDS.
		Listeners:      []*v3listenerpb.Listener{e2e.DefaultClientListener(ldsName, rdsName)},
		Clusters:       []*v3clusterpb.Cluster{e2e.DefaultCluster(cdsName, edsName, e2e.SecurityLevelNone)},
		SkipValidation: true,
	}
	resourcesAnother := e2e.UpdateOptions{
		NodeID: nodeID,
		// This has only RDS and EDS.
		Routes:         []*v3routepb.RouteConfiguration{e2e.DefaultRouteConfig(rdsName, ldsName, cdsName)},
		Endpoints:      []*v3endpointpb.ClusterLoadAssignment{e2e.DefaultEndpoint(edsName, "localhost", []uint32{port})},
		SkipValidation: true,
	}

	ctx, cancel := context.WithTimeout(context.Background(), defaultTestTimeout)
	defer cancel()
	// This has only LDS and CDS.
	if err := serverDefaultAuth.Update(ctx, resourcesDefault); err != nil {
		t.Fatal(err)
	}
	// This has only RDS and EDS.
	if err := serverAnotherAuth.Update(ctx, resourcesAnother); err != nil {
		t.Fatal(err)
	}

	// Create a ClientConn and make a successful RPC.
	cc, err := grpc.Dial(fmt.Sprintf("xds:///%s", serviceName), grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithResolvers(resolver))
	if err != nil {
		t.Fatalf("failed to dial local test server: %v", err)
	}
	defer cc.Close()

	client := testpb.NewTestServiceClient(cc)
	if _, err := client.EmptyCall(ctx, &testpb.Empty{}, grpc.WaitForReady(true)); err != nil {
		t.Fatalf("rpc EmptyCall() failed: %v", err)
	}
}
