// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.16.0
// source: envoy/extensions/filters/network/thrift_proxy/filters/ratelimit/v3/rate_limit.proto

package envoy_extensions_filters_network_thrift_proxy_filters_ratelimit_v3

import (
	_ "github.com/cncf/xds/go/udpa/annotations"
	v3 "github.com/hxx258456/ccgo/go-control-plane/envoy/config/ratelimit/v3"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	duration "github.com/golang/protobuf/ptypes/duration"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// [#next-free-field: 6]
type RateLimit struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The rate limit domain to use in the rate limit service request.
	Domain string `protobuf:"bytes,1,opt,name=domain,proto3" json:"domain,omitempty"`
	// Specifies the rate limit configuration stage. Each configured rate limit filter performs a
	// rate limit check using descriptors configured in the
	// :ref:`envoy_v3_api_msg_extensions.filters.network.thrift_proxy.v3.RouteAction` for the request.
	// Only those entries with a matching stage number are used for a given filter. If not set, the
	// default stage number is 0.
	//
	// .. note::
	//
	//  The filter supports a range of 0 - 10 inclusively for stage numbers.
	Stage uint32 `protobuf:"varint,2,opt,name=stage,proto3" json:"stage,omitempty"`
	// The timeout in milliseconds for the rate limit service RPC. If not
	// set, this defaults to 20ms.
	Timeout *duration.Duration `protobuf:"bytes,3,opt,name=timeout,proto3" json:"timeout,omitempty"`
	// The filter's behaviour in case the rate limiting service does
	// not respond back. When it is set to true, Envoy will not allow traffic in case of
	// communication failure between rate limiting service and the proxy.
	// Defaults to false.
	FailureModeDeny bool `protobuf:"varint,4,opt,name=failure_mode_deny,json=failureModeDeny,proto3" json:"failure_mode_deny,omitempty"`
	// Configuration for an external rate limit service provider. If not
	// specified, any calls to the rate limit service will immediately return
	// success.
	RateLimitService *v3.RateLimitServiceConfig `protobuf:"bytes,5,opt,name=rate_limit_service,json=rateLimitService,proto3" json:"rate_limit_service,omitempty"`
}

func (x *RateLimit) Reset() {
	*x = RateLimit{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_extensions_filters_network_thrift_proxy_filters_ratelimit_v3_rate_limit_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RateLimit) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RateLimit) ProtoMessage() {}

func (x *RateLimit) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_extensions_filters_network_thrift_proxy_filters_ratelimit_v3_rate_limit_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RateLimit.ProtoReflect.Descriptor instead.
func (*RateLimit) Descriptor() ([]byte, []int) {
	return file_envoy_extensions_filters_network_thrift_proxy_filters_ratelimit_v3_rate_limit_proto_rawDescGZIP(), []int{0}
}

func (x *RateLimit) GetDomain() string {
	if x != nil {
		return x.Domain
	}
	return ""
}

func (x *RateLimit) GetStage() uint32 {
	if x != nil {
		return x.Stage
	}
	return 0
}

func (x *RateLimit) GetTimeout() *duration.Duration {
	if x != nil {
		return x.Timeout
	}
	return nil
}

func (x *RateLimit) GetFailureModeDeny() bool {
	if x != nil {
		return x.FailureModeDeny
	}
	return false
}

func (x *RateLimit) GetRateLimitService() *v3.RateLimitServiceConfig {
	if x != nil {
		return x.RateLimitService
	}
	return nil
}

var File_envoy_extensions_filters_network_thrift_proxy_filters_ratelimit_v3_rate_limit_proto protoreflect.FileDescriptor

var file_envoy_extensions_filters_network_thrift_proxy_filters_ratelimit_v3_rate_limit_proto_rawDesc = []byte{
	0x0a, 0x53, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2f, 0x6e, 0x65, 0x74, 0x77, 0x6f,
	0x72, 0x6b, 0x2f, 0x74, 0x68, 0x72, 0x69, 0x66, 0x74, 0x5f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2f,
	0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2f, 0x72, 0x61, 0x74, 0x65, 0x6c, 0x69, 0x6d, 0x69,
	0x74, 0x2f, 0x76, 0x33, 0x2f, 0x72, 0x61, 0x74, 0x65, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x42, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74,
	0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2e,
	0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x74, 0x68, 0x72, 0x69, 0x66, 0x74, 0x5f, 0x70,
	0x72, 0x6f, 0x78, 0x79, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x72, 0x61, 0x74,
	0x65, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x2e, 0x76, 0x33, 0x1a, 0x23, 0x65, 0x6e, 0x76, 0x6f, 0x79,
	0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x72, 0x61, 0x74, 0x65, 0x6c, 0x69, 0x6d, 0x69,
	0x74, 0x2f, 0x76, 0x33, 0x2f, 0x72, 0x6c, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d,
	0x75, 0x64, 0x70, 0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x21, 0x75,
	0x64, 0x70, 0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f,
	0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64,
	0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd8, 0x02, 0x0a, 0x09, 0x52, 0x61,
	0x74, 0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x1f, 0x0a, 0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69,
	0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01,
	0x52, 0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x12, 0x1d, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x67,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x2a, 0x02, 0x18, 0x0a,
	0x52, 0x05, 0x73, 0x74, 0x61, 0x67, 0x65, 0x12, 0x33, 0x0a, 0x07, 0x74, 0x69, 0x6d, 0x65, 0x6f,
	0x75, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x07, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x12, 0x2a, 0x0a, 0x11,
	0x66, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x5f, 0x64, 0x65, 0x6e,
	0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0f, 0x66, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65,
	0x4d, 0x6f, 0x64, 0x65, 0x44, 0x65, 0x6e, 0x79, 0x12, 0x69, 0x0a, 0x12, 0x72, 0x61, 0x74, 0x65,
	0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x31, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x2e, 0x72, 0x61, 0x74, 0x65, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x2e, 0x76, 0x33,
	0x2e, 0x52, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x8a, 0x01, 0x02, 0x10,
	0x01, 0x52, 0x10, 0x72, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x3a, 0x3f, 0x9a, 0xc5, 0x88, 0x1e, 0x3a, 0x0a, 0x38, 0x65, 0x6e, 0x76, 0x6f,
	0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x2e,
	0x74, 0x68, 0x72, 0x69, 0x66, 0x74, 0x2e, 0x72, 0x61, 0x74, 0x65, 0x5f, 0x6c, 0x69, 0x6d, 0x69,
	0x74, 0x2e, 0x76, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x52, 0x61, 0x74, 0x65, 0x4c,
	0x69, 0x6d, 0x69, 0x74, 0x42, 0x6c, 0x0a, 0x50, 0x69, 0x6f, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79,
	0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65,
	0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x6e,
	0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x74, 0x68, 0x72, 0x69, 0x66, 0x74, 0x5f, 0x70, 0x72,
	0x6f, 0x78, 0x79, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x72, 0x61, 0x74, 0x65,
	0x6c, 0x69, 0x6d, 0x69, 0x74, 0x2e, 0x76, 0x33, 0x42, 0x0e, 0x52, 0x61, 0x74, 0x65, 0x4c, 0x69,
	0x6d, 0x69, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0xba, 0x80, 0xc8, 0xd1, 0x06, 0x02,
	0x10, 0x02, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_envoy_extensions_filters_network_thrift_proxy_filters_ratelimit_v3_rate_limit_proto_rawDescOnce sync.Once
	file_envoy_extensions_filters_network_thrift_proxy_filters_ratelimit_v3_rate_limit_proto_rawDescData = file_envoy_extensions_filters_network_thrift_proxy_filters_ratelimit_v3_rate_limit_proto_rawDesc
)

func file_envoy_extensions_filters_network_thrift_proxy_filters_ratelimit_v3_rate_limit_proto_rawDescGZIP() []byte {
	file_envoy_extensions_filters_network_thrift_proxy_filters_ratelimit_v3_rate_limit_proto_rawDescOnce.Do(func() {
		file_envoy_extensions_filters_network_thrift_proxy_filters_ratelimit_v3_rate_limit_proto_rawDescData = protoimpl.X.CompressGZIP(file_envoy_extensions_filters_network_thrift_proxy_filters_ratelimit_v3_rate_limit_proto_rawDescData)
	})
	return file_envoy_extensions_filters_network_thrift_proxy_filters_ratelimit_v3_rate_limit_proto_rawDescData
}

var file_envoy_extensions_filters_network_thrift_proxy_filters_ratelimit_v3_rate_limit_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_envoy_extensions_filters_network_thrift_proxy_filters_ratelimit_v3_rate_limit_proto_goTypes = []interface{}{
	(*RateLimit)(nil),                 // 0: envoy.extensions.filters.network.thrift_proxy.filters.ratelimit.v3.RateLimit
	(*duration.Duration)(nil),         // 1: google.protobuf.Duration
	(*v3.RateLimitServiceConfig)(nil), // 2: envoy.config.ratelimit.v3.RateLimitServiceConfig
}
var file_envoy_extensions_filters_network_thrift_proxy_filters_ratelimit_v3_rate_limit_proto_depIdxs = []int32{
	1, // 0: envoy.extensions.filters.network.thrift_proxy.filters.ratelimit.v3.RateLimit.timeout:type_name -> google.protobuf.Duration
	2, // 1: envoy.extensions.filters.network.thrift_proxy.filters.ratelimit.v3.RateLimit.rate_limit_service:type_name -> envoy.config.ratelimit.v3.RateLimitServiceConfig
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() {
	file_envoy_extensions_filters_network_thrift_proxy_filters_ratelimit_v3_rate_limit_proto_init()
}
func file_envoy_extensions_filters_network_thrift_proxy_filters_ratelimit_v3_rate_limit_proto_init() {
	if File_envoy_extensions_filters_network_thrift_proxy_filters_ratelimit_v3_rate_limit_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_envoy_extensions_filters_network_thrift_proxy_filters_ratelimit_v3_rate_limit_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RateLimit); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_envoy_extensions_filters_network_thrift_proxy_filters_ratelimit_v3_rate_limit_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_envoy_extensions_filters_network_thrift_proxy_filters_ratelimit_v3_rate_limit_proto_goTypes,
		DependencyIndexes: file_envoy_extensions_filters_network_thrift_proxy_filters_ratelimit_v3_rate_limit_proto_depIdxs,
		MessageInfos:      file_envoy_extensions_filters_network_thrift_proxy_filters_ratelimit_v3_rate_limit_proto_msgTypes,
	}.Build()
	File_envoy_extensions_filters_network_thrift_proxy_filters_ratelimit_v3_rate_limit_proto = out.File
	file_envoy_extensions_filters_network_thrift_proxy_filters_ratelimit_v3_rate_limit_proto_rawDesc = nil
	file_envoy_extensions_filters_network_thrift_proxy_filters_ratelimit_v3_rate_limit_proto_goTypes = nil
	file_envoy_extensions_filters_network_thrift_proxy_filters_ratelimit_v3_rate_limit_proto_depIdxs = nil
}
