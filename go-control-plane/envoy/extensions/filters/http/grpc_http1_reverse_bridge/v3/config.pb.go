// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.16.0
// source: envoy/extensions/filters/http/grpc_http1_reverse_bridge/v3/config.proto

package envoy_extensions_filters_http_grpc_http1_reverse_bridge_v3

import (
	_ "github.com/cncf/xds/go/udpa/annotations"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
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

// gRPC reverse bridge filter configuration
type FilterConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The content-type to pass to the upstream when the gRPC bridge filter is applied.
	// The filter will also validate that the upstream responds with the same content type.
	ContentType string `protobuf:"bytes,1,opt,name=content_type,json=contentType,proto3" json:"content_type,omitempty"`
	// If true, Envoy will assume that the upstream doesn't understand gRPC frames and
	// strip the gRPC frame from the request, and add it back in to the response. This will
	// hide the gRPC semantics from the upstream, allowing it to receive and respond with a
	// simple binary encoded protobuf. In order to calculate the `Content-Length` header value, Envoy
	// will buffer the upstream response unless :ref:`response_size_header
	// <envoy_v3_api_field_extensions.filters.http.grpc_http1_reverse_bridge.v3.FilterConfig.response_size_header>`
	// is set, in which case Envoy will use the value of an upstream header to calculate the content
	// length.
	WithholdGrpcFrames bool `protobuf:"varint,2,opt,name=withhold_grpc_frames,json=withholdGrpcFrames,proto3" json:"withhold_grpc_frames,omitempty"`
	// When :ref:`withhold_grpc_frames
	// <envoy_v3_api_field_extensions.filters.http.grpc_http1_reverse_bridge.v3.FilterConfig.withhold_grpc_frames>`
	// is true, this option controls how Envoy calculates the `Content-Length`. When
	// *response_size_header* is empty, Envoy will buffer the upstream response to calculate its
	// size. When *response_size_header* is set to a non-empty string, Envoy will stream the response
	// to the downstream and it will use the value of the response header with this name to set the
	// `Content-Length` header and gRPC frame size. If the header with this name is repeated, only
	// the first value will be used.
	//
	// Envoy will treat the upstream response as an error if this option is specified and the header
	// is missing or if the value does not match the actual response body size.
	ResponseSizeHeader string `protobuf:"bytes,3,opt,name=response_size_header,json=responseSizeHeader,proto3" json:"response_size_header,omitempty"`
}

func (x *FilterConfig) Reset() {
	*x = FilterConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_extensions_filters_http_grpc_http1_reverse_bridge_v3_config_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FilterConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FilterConfig) ProtoMessage() {}

func (x *FilterConfig) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_extensions_filters_http_grpc_http1_reverse_bridge_v3_config_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FilterConfig.ProtoReflect.Descriptor instead.
func (*FilterConfig) Descriptor() ([]byte, []int) {
	return file_envoy_extensions_filters_http_grpc_http1_reverse_bridge_v3_config_proto_rawDescGZIP(), []int{0}
}

func (x *FilterConfig) GetContentType() string {
	if x != nil {
		return x.ContentType
	}
	return ""
}

func (x *FilterConfig) GetWithholdGrpcFrames() bool {
	if x != nil {
		return x.WithholdGrpcFrames
	}
	return false
}

func (x *FilterConfig) GetResponseSizeHeader() string {
	if x != nil {
		return x.ResponseSizeHeader
	}
	return ""
}

// gRPC reverse bridge filter configuration per virtualhost/route/weighted-cluster level.
type FilterConfigPerRoute struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// If true, disables gRPC reverse bridge filter for this particular vhost or route.
	// If disabled is specified in multiple per-filter-configs, the most specific one will be used.
	Disabled bool `protobuf:"varint,1,opt,name=disabled,proto3" json:"disabled,omitempty"`
}

func (x *FilterConfigPerRoute) Reset() {
	*x = FilterConfigPerRoute{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_extensions_filters_http_grpc_http1_reverse_bridge_v3_config_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FilterConfigPerRoute) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FilterConfigPerRoute) ProtoMessage() {}

func (x *FilterConfigPerRoute) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_extensions_filters_http_grpc_http1_reverse_bridge_v3_config_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FilterConfigPerRoute.ProtoReflect.Descriptor instead.
func (*FilterConfigPerRoute) Descriptor() ([]byte, []int) {
	return file_envoy_extensions_filters_http_grpc_http1_reverse_bridge_v3_config_proto_rawDescGZIP(), []int{1}
}

func (x *FilterConfigPerRoute) GetDisabled() bool {
	if x != nil {
		return x.Disabled
	}
	return false
}

var File_envoy_extensions_filters_http_grpc_http1_reverse_bridge_v3_config_proto protoreflect.FileDescriptor

var file_envoy_extensions_filters_http_grpc_http1_reverse_bridge_v3_config_proto_rawDesc = []byte{
	0x0a, 0x47, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2f, 0x68, 0x74, 0x74, 0x70, 0x2f,
	0x67, 0x72, 0x70, 0x63, 0x5f, 0x68, 0x74, 0x74, 0x70, 0x31, 0x5f, 0x72, 0x65, 0x76, 0x65, 0x72,
	0x73, 0x65, 0x5f, 0x62, 0x72, 0x69, 0x64, 0x67, 0x65, 0x2f, 0x76, 0x33, 0x2f, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x3a, 0x65, 0x6e, 0x76, 0x6f, 0x79,
	0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x66, 0x69, 0x6c, 0x74,
	0x65, 0x72, 0x73, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x68, 0x74,
	0x74, 0x70, 0x31, 0x5f, 0x72, 0x65, 0x76, 0x65, 0x72, 0x73, 0x65, 0x5f, 0x62, 0x72, 0x69, 0x64,
	0x67, 0x65, 0x2e, 0x76, 0x33, 0x1a, 0x1d, 0x75, 0x64, 0x70, 0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x21, 0x75, 0x64, 0x70, 0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x69, 0x6e,
	0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xef, 0x01, 0x0a, 0x0c, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x12, 0x2a, 0x0a, 0x0c, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01,
	0x52, 0x0b, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x30, 0x0a,
	0x14, 0x77, 0x69, 0x74, 0x68, 0x68, 0x6f, 0x6c, 0x64, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x66,
	0x72, 0x61, 0x6d, 0x65, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x12, 0x77, 0x69, 0x74,
	0x68, 0x68, 0x6f, 0x6c, 0x64, 0x47, 0x72, 0x70, 0x63, 0x46, 0x72, 0x61, 0x6d, 0x65, 0x73, 0x12,
	0x30, 0x0a, 0x14, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65,
	0x5f, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x72,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x48, 0x65, 0x61, 0x64, 0x65,
	0x72, 0x3a, 0x4f, 0x9a, 0xc5, 0x88, 0x1e, 0x4a, 0x0a, 0x48, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x2e, 0x68, 0x74,
	0x74, 0x70, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x68, 0x74, 0x74, 0x70, 0x31, 0x5f, 0x72, 0x65,
	0x76, 0x65, 0x72, 0x73, 0x65, 0x5f, 0x62, 0x72, 0x69, 0x64, 0x67, 0x65, 0x2e, 0x76, 0x32, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x22, 0x8b, 0x01, 0x0a, 0x14, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x50, 0x65, 0x72, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x64,
	0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x64,
	0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x3a, 0x57, 0x9a, 0xc5, 0x88, 0x1e, 0x52, 0x0a, 0x50,
	0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x66, 0x69, 0x6c,
	0x74, 0x65, 0x72, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x68, 0x74,
	0x74, 0x70, 0x31, 0x5f, 0x72, 0x65, 0x76, 0x65, 0x72, 0x73, 0x65, 0x5f, 0x62, 0x72, 0x69, 0x64,
	0x67, 0x65, 0x2e, 0x76, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x46, 0x69, 0x6c, 0x74,
	0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x50, 0x65, 0x72, 0x52, 0x6f, 0x75, 0x74, 0x65,
	0x42, 0x61, 0x0a, 0x48, 0x69, 0x6f, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78,
	0x79, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x2e,
	0x67, 0x72, 0x70, 0x63, 0x5f, 0x68, 0x74, 0x74, 0x70, 0x31, 0x5f, 0x72, 0x65, 0x76, 0x65, 0x72,
	0x73, 0x65, 0x5f, 0x62, 0x72, 0x69, 0x64, 0x67, 0x65, 0x2e, 0x76, 0x33, 0x42, 0x0b, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0xba, 0x80, 0xc8, 0xd1, 0x06,
	0x02, 0x10, 0x02, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_envoy_extensions_filters_http_grpc_http1_reverse_bridge_v3_config_proto_rawDescOnce sync.Once
	file_envoy_extensions_filters_http_grpc_http1_reverse_bridge_v3_config_proto_rawDescData = file_envoy_extensions_filters_http_grpc_http1_reverse_bridge_v3_config_proto_rawDesc
)

func file_envoy_extensions_filters_http_grpc_http1_reverse_bridge_v3_config_proto_rawDescGZIP() []byte {
	file_envoy_extensions_filters_http_grpc_http1_reverse_bridge_v3_config_proto_rawDescOnce.Do(func() {
		file_envoy_extensions_filters_http_grpc_http1_reverse_bridge_v3_config_proto_rawDescData = protoimpl.X.CompressGZIP(file_envoy_extensions_filters_http_grpc_http1_reverse_bridge_v3_config_proto_rawDescData)
	})
	return file_envoy_extensions_filters_http_grpc_http1_reverse_bridge_v3_config_proto_rawDescData
}

var file_envoy_extensions_filters_http_grpc_http1_reverse_bridge_v3_config_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_envoy_extensions_filters_http_grpc_http1_reverse_bridge_v3_config_proto_goTypes = []interface{}{
	(*FilterConfig)(nil),         // 0: envoy.extensions.filters.http.grpc_http1_reverse_bridge.v3.FilterConfig
	(*FilterConfigPerRoute)(nil), // 1: envoy.extensions.filters.http.grpc_http1_reverse_bridge.v3.FilterConfigPerRoute
}
var file_envoy_extensions_filters_http_grpc_http1_reverse_bridge_v3_config_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_envoy_extensions_filters_http_grpc_http1_reverse_bridge_v3_config_proto_init() }
func file_envoy_extensions_filters_http_grpc_http1_reverse_bridge_v3_config_proto_init() {
	if File_envoy_extensions_filters_http_grpc_http1_reverse_bridge_v3_config_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_envoy_extensions_filters_http_grpc_http1_reverse_bridge_v3_config_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FilterConfig); i {
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
		file_envoy_extensions_filters_http_grpc_http1_reverse_bridge_v3_config_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FilterConfigPerRoute); i {
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
			RawDescriptor: file_envoy_extensions_filters_http_grpc_http1_reverse_bridge_v3_config_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_envoy_extensions_filters_http_grpc_http1_reverse_bridge_v3_config_proto_goTypes,
		DependencyIndexes: file_envoy_extensions_filters_http_grpc_http1_reverse_bridge_v3_config_proto_depIdxs,
		MessageInfos:      file_envoy_extensions_filters_http_grpc_http1_reverse_bridge_v3_config_proto_msgTypes,
	}.Build()
	File_envoy_extensions_filters_http_grpc_http1_reverse_bridge_v3_config_proto = out.File
	file_envoy_extensions_filters_http_grpc_http1_reverse_bridge_v3_config_proto_rawDesc = nil
	file_envoy_extensions_filters_http_grpc_http1_reverse_bridge_v3_config_proto_goTypes = nil
	file_envoy_extensions_filters_http_grpc_http1_reverse_bridge_v3_config_proto_depIdxs = nil
}