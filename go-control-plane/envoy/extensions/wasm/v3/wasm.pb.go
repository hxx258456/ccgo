// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.16.0
// source: envoy/extensions/wasm/v3/wasm.proto

package envoy_extensions_wasm_v3

import (
	_ "github.com/cncf/xds/go/udpa/annotations"
	v3 "github.com/hxx258456/ccgo/go-control-plane/envoy/config/core/v3"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	any "github.com/golang/protobuf/ptypes/any"
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

// Configuration for restricting Proxy-Wasm capabilities available to modules.
type CapabilityRestrictionConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The Proxy-Wasm capabilities which will be allowed. Capabilities are mapped by
	// name. The *SanitizationConfig* which each capability maps to is currently unimplemented and ignored,
	// and so should be left empty.
	//
	// The capability names are given in the
	// `Proxy-Wasm ABI <https://github.com/proxy-wasm/spec/tree/master/abi-versions/vNEXT>`_.
	// Additionally, the following WASI capabilities from
	// `this list <https://github.com/WebAssembly/WASI/blob/master/phases/snapshot/docs.md#modules>`_
	// are implemented and can be allowed:
	// *fd_write*, *fd_read*, *fd_seek*, *fd_close*, *fd_fdstat_get*, *environ_get*, *environ_sizes_get*,
	// *args_get*, *args_sizes_get*, *proc_exit*, *clock_time_get*, *random_get*.
	AllowedCapabilities map[string]*SanitizationConfig `protobuf:"bytes,1,rep,name=allowed_capabilities,json=allowedCapabilities,proto3" json:"allowed_capabilities,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *CapabilityRestrictionConfig) Reset() {
	*x = CapabilityRestrictionConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_extensions_wasm_v3_wasm_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CapabilityRestrictionConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CapabilityRestrictionConfig) ProtoMessage() {}

func (x *CapabilityRestrictionConfig) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_extensions_wasm_v3_wasm_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CapabilityRestrictionConfig.ProtoReflect.Descriptor instead.
func (*CapabilityRestrictionConfig) Descriptor() ([]byte, []int) {
	return file_envoy_extensions_wasm_v3_wasm_proto_rawDescGZIP(), []int{0}
}

func (x *CapabilityRestrictionConfig) GetAllowedCapabilities() map[string]*SanitizationConfig {
	if x != nil {
		return x.AllowedCapabilities
	}
	return nil
}

// Configuration for sanitization of inputs to an allowed capability.
//
// NOTE: This is currently unimplemented.
type SanitizationConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SanitizationConfig) Reset() {
	*x = SanitizationConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_extensions_wasm_v3_wasm_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SanitizationConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SanitizationConfig) ProtoMessage() {}

func (x *SanitizationConfig) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_extensions_wasm_v3_wasm_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SanitizationConfig.ProtoReflect.Descriptor instead.
func (*SanitizationConfig) Descriptor() ([]byte, []int) {
	return file_envoy_extensions_wasm_v3_wasm_proto_rawDescGZIP(), []int{1}
}

// Configuration for a Wasm VM.
// [#next-free-field: 8]
type VmConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// An ID which will be used along with a hash of the wasm code (or the name of the registered Null
	// VM plugin) to determine which VM will be used for the plugin. All plugins which use the same
	// *vm_id* and code will use the same VM. May be left blank. Sharing a VM between plugins can
	// reduce memory utilization and make sharing of data easier which may have security implications.
	// [#comment: TODO: add ref for details.]
	VmId string `protobuf:"bytes,1,opt,name=vm_id,json=vmId,proto3" json:"vm_id,omitempty"`
	// The Wasm runtime type.
	// Available Wasm runtime types are registered as extensions. The following runtimes are included
	// in Envoy code base:
	//
	// .. _extension_envoy.wasm.runtime.null:
	//
	// **envoy.wasm.runtime.null**: Null sandbox, the Wasm module must be compiled and linked into the
	// Envoy binary. The registered name is given in the *code* field as *inline_string*.
	//
	// .. _extension_envoy.wasm.runtime.v8:
	//
	// **envoy.wasm.runtime.v8**: `V8 <https://v8.dev/>`_-based WebAssembly runtime.
	//
	// .. _extension_envoy.wasm.runtime.wamr:
	//
	// **envoy.wasm.runtime.wamr**: `WAMR <https://github.com/bytecodealliance/wasm-micro-runtime/>`_-based WebAssembly runtime.
	// This runtime is not enabled in the official build.
	//
	// .. _extension_envoy.wasm.runtime.wavm:
	//
	// **envoy.wasm.runtime.wavm**: `WAVM <https://wavm.github.io/>`_-based WebAssembly runtime.
	// This runtime is not enabled in the official build.
	//
	// .. _extension_envoy.wasm.runtime.wasmtime:
	//
	// **envoy.wasm.runtime.wasmtime**: `Wasmtime <https://wasmtime.dev/>`_-based WebAssembly runtime.
	// This runtime is not enabled in the official build.
	//
	// [#extension-category: envoy.wasm.runtime]
	Runtime string `protobuf:"bytes,2,opt,name=runtime,proto3" json:"runtime,omitempty"`
	// The Wasm code that Envoy will execute.
	Code *v3.AsyncDataSource `protobuf:"bytes,3,opt,name=code,proto3" json:"code,omitempty"`
	// The Wasm configuration used in initialization of a new VM
	// (proxy_on_start). `google.protobuf.Struct` is serialized as JSON before
	// passing it to the plugin. `google.protobuf.BytesValue` and
	// `google.protobuf.StringValue` are passed directly without the wrapper.
	Configuration *any.Any `protobuf:"bytes,4,opt,name=configuration,proto3" json:"configuration,omitempty"`
	// Allow the wasm file to include pre-compiled code on VMs which support it.
	// Warning: this should only be enable for trusted sources as the precompiled code is not
	// verified.
	AllowPrecompiled bool `protobuf:"varint,5,opt,name=allow_precompiled,json=allowPrecompiled,proto3" json:"allow_precompiled,omitempty"`
	// If true and the code needs to be remotely fetched and it is not in the cache then NACK the configuration
	// update and do a background fetch to fill the cache, otherwise fetch the code asynchronously and enter
	// warming state.
	NackOnCodeCacheMiss bool `protobuf:"varint,6,opt,name=nack_on_code_cache_miss,json=nackOnCodeCacheMiss,proto3" json:"nack_on_code_cache_miss,omitempty"`
	// Specifies environment variables to be injected to this VM which will be available through
	// WASI's ``environ_get`` and ``environ_get_sizes`` system calls. Note that these functions are mostly implicitly
	// called in your language's standard library, so you do not need to call them directly and you can access to env
	// vars just like when you do on native platforms.
	// Warning: Envoy rejects the configuration if there's conflict of key space.
	EnvironmentVariables *EnvironmentVariables `protobuf:"bytes,7,opt,name=environment_variables,json=environmentVariables,proto3" json:"environment_variables,omitempty"`
}

func (x *VmConfig) Reset() {
	*x = VmConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_extensions_wasm_v3_wasm_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VmConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VmConfig) ProtoMessage() {}

func (x *VmConfig) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_extensions_wasm_v3_wasm_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VmConfig.ProtoReflect.Descriptor instead.
func (*VmConfig) Descriptor() ([]byte, []int) {
	return file_envoy_extensions_wasm_v3_wasm_proto_rawDescGZIP(), []int{2}
}

func (x *VmConfig) GetVmId() string {
	if x != nil {
		return x.VmId
	}
	return ""
}

func (x *VmConfig) GetRuntime() string {
	if x != nil {
		return x.Runtime
	}
	return ""
}

func (x *VmConfig) GetCode() *v3.AsyncDataSource {
	if x != nil {
		return x.Code
	}
	return nil
}

func (x *VmConfig) GetConfiguration() *any.Any {
	if x != nil {
		return x.Configuration
	}
	return nil
}

func (x *VmConfig) GetAllowPrecompiled() bool {
	if x != nil {
		return x.AllowPrecompiled
	}
	return false
}

func (x *VmConfig) GetNackOnCodeCacheMiss() bool {
	if x != nil {
		return x.NackOnCodeCacheMiss
	}
	return false
}

func (x *VmConfig) GetEnvironmentVariables() *EnvironmentVariables {
	if x != nil {
		return x.EnvironmentVariables
	}
	return nil
}

type EnvironmentVariables struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The keys of *Envoy's* environment variables exposed to this VM. In other words, if a key exists in Envoy's environment
	// variables, then that key-value pair will be injected. Note that if a key does not exist, it will be ignored.
	HostEnvKeys []string `protobuf:"bytes,1,rep,name=host_env_keys,json=hostEnvKeys,proto3" json:"host_env_keys,omitempty"`
	// Explicitly given key-value pairs to be injected to this VM in the form of "KEY=VALUE".
	KeyValues map[string]string `protobuf:"bytes,2,rep,name=key_values,json=keyValues,proto3" json:"key_values,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *EnvironmentVariables) Reset() {
	*x = EnvironmentVariables{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_extensions_wasm_v3_wasm_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EnvironmentVariables) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EnvironmentVariables) ProtoMessage() {}

func (x *EnvironmentVariables) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_extensions_wasm_v3_wasm_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EnvironmentVariables.ProtoReflect.Descriptor instead.
func (*EnvironmentVariables) Descriptor() ([]byte, []int) {
	return file_envoy_extensions_wasm_v3_wasm_proto_rawDescGZIP(), []int{3}
}

func (x *EnvironmentVariables) GetHostEnvKeys() []string {
	if x != nil {
		return x.HostEnvKeys
	}
	return nil
}

func (x *EnvironmentVariables) GetKeyValues() map[string]string {
	if x != nil {
		return x.KeyValues
	}
	return nil
}

// Base Configuration for Wasm Plugins e.g. filters and services.
// [#next-free-field: 7]
type PluginConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// A unique name for a filters/services in a VM for use in identifying the filter/service if
	// multiple filters/services are handled by the same *vm_id* and *root_id* and for
	// logging/debugging.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// A unique ID for a set of filters/services in a VM which will share a RootContext and Contexts
	// if applicable (e.g. an Wasm HttpFilter and an Wasm AccessLog). If left blank, all
	// filters/services with a blank root_id with the same *vm_id* will share Context(s).
	RootId string `protobuf:"bytes,2,opt,name=root_id,json=rootId,proto3" json:"root_id,omitempty"`
	// Configuration for finding or starting VM.
	//
	// Types that are assignable to Vm:
	//	*PluginConfig_VmConfig
	Vm isPluginConfig_Vm `protobuf_oneof:"vm"`
	// Filter/service configuration used to configure or reconfigure a plugin
	// (proxy_on_configuration).
	// `google.protobuf.Struct` is serialized as JSON before
	// passing it to the plugin. `google.protobuf.BytesValue` and
	// `google.protobuf.StringValue` are passed directly without the wrapper.
	Configuration *any.Any `protobuf:"bytes,4,opt,name=configuration,proto3" json:"configuration,omitempty"`
	// If there is a fatal error on the VM (e.g. exception, abort(), on_start or on_configure return false),
	// then all plugins associated with the VM will either fail closed (by default), e.g. by returning an HTTP 503 error,
	// or fail open (if 'fail_open' is set to true) by bypassing the filter. Note: when on_start or on_configure return false
	// during xDS updates the xDS configuration will be rejected and when on_start or on_configuration return false on initial
	// startup the proxy will not start.
	FailOpen bool `protobuf:"varint,5,opt,name=fail_open,json=failOpen,proto3" json:"fail_open,omitempty"`
	// Configuration for restricting Proxy-Wasm capabilities available to modules.
	CapabilityRestrictionConfig *CapabilityRestrictionConfig `protobuf:"bytes,6,opt,name=capability_restriction_config,json=capabilityRestrictionConfig,proto3" json:"capability_restriction_config,omitempty"`
}

func (x *PluginConfig) Reset() {
	*x = PluginConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_extensions_wasm_v3_wasm_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PluginConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PluginConfig) ProtoMessage() {}

func (x *PluginConfig) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_extensions_wasm_v3_wasm_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PluginConfig.ProtoReflect.Descriptor instead.
func (*PluginConfig) Descriptor() ([]byte, []int) {
	return file_envoy_extensions_wasm_v3_wasm_proto_rawDescGZIP(), []int{4}
}

func (x *PluginConfig) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *PluginConfig) GetRootId() string {
	if x != nil {
		return x.RootId
	}
	return ""
}

func (m *PluginConfig) GetVm() isPluginConfig_Vm {
	if m != nil {
		return m.Vm
	}
	return nil
}

func (x *PluginConfig) GetVmConfig() *VmConfig {
	if x, ok := x.GetVm().(*PluginConfig_VmConfig); ok {
		return x.VmConfig
	}
	return nil
}

func (x *PluginConfig) GetConfiguration() *any.Any {
	if x != nil {
		return x.Configuration
	}
	return nil
}

func (x *PluginConfig) GetFailOpen() bool {
	if x != nil {
		return x.FailOpen
	}
	return false
}

func (x *PluginConfig) GetCapabilityRestrictionConfig() *CapabilityRestrictionConfig {
	if x != nil {
		return x.CapabilityRestrictionConfig
	}
	return nil
}

type isPluginConfig_Vm interface {
	isPluginConfig_Vm()
}

type PluginConfig_VmConfig struct {
	VmConfig *VmConfig `protobuf:"bytes,3,opt,name=vm_config,json=vmConfig,proto3,oneof"` // TODO: add referential VM configurations.
}

func (*PluginConfig_VmConfig) isPluginConfig_Vm() {}

// WasmService is configured as a built-in *envoy.wasm_service* :ref:`WasmService
// <config_wasm_service>` This opaque configuration will be used to create a Wasm Service.
type WasmService struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// General plugin configuration.
	Config *PluginConfig `protobuf:"bytes,1,opt,name=config,proto3" json:"config,omitempty"`
	// If true, create a single VM rather than creating one VM per worker. Such a singleton can
	// not be used with filters.
	Singleton bool `protobuf:"varint,2,opt,name=singleton,proto3" json:"singleton,omitempty"`
}

func (x *WasmService) Reset() {
	*x = WasmService{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_extensions_wasm_v3_wasm_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WasmService) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WasmService) ProtoMessage() {}

func (x *WasmService) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_extensions_wasm_v3_wasm_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WasmService.ProtoReflect.Descriptor instead.
func (*WasmService) Descriptor() ([]byte, []int) {
	return file_envoy_extensions_wasm_v3_wasm_proto_rawDescGZIP(), []int{5}
}

func (x *WasmService) GetConfig() *PluginConfig {
	if x != nil {
		return x.Config
	}
	return nil
}

func (x *WasmService) GetSingleton() bool {
	if x != nil {
		return x.Singleton
	}
	return false
}

var File_envoy_extensions_wasm_v3_wasm_proto protoreflect.FileDescriptor

var file_envoy_extensions_wasm_v3_wasm_proto_rawDesc = []byte{
	0x0a, 0x23, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x77, 0x61, 0x73, 0x6d, 0x2f, 0x76, 0x33, 0x2f, 0x77, 0x61, 0x73, 0x6d, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x18, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74,
	0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x77, 0x61, 0x73, 0x6d, 0x2e, 0x76, 0x33, 0x1a,
	0x1f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x63, 0x6f,
	0x72, 0x65, 0x2f, 0x76, 0x33, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x75, 0x64, 0x70,
	0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69,
	0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x97, 0x02, 0x0a, 0x1b, 0x43, 0x61, 0x70, 0x61, 0x62, 0x69, 0x6c, 0x69,
	0x74, 0x79, 0x52, 0x65, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x12, 0x81, 0x01, 0x0a, 0x14, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x64, 0x5f,
	0x63, 0x61, 0x70, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x69, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x4e, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e,
	0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x77, 0x61, 0x73, 0x6d, 0x2e, 0x76, 0x33, 0x2e, 0x43, 0x61,
	0x70, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x52, 0x65, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x41, 0x6c, 0x6c, 0x6f, 0x77, 0x65,
	0x64, 0x43, 0x61, 0x70, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x69, 0x65, 0x73, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x52, 0x13, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x64, 0x43, 0x61, 0x70, 0x61, 0x62,
	0x69, 0x6c, 0x69, 0x74, 0x69, 0x65, 0x73, 0x1a, 0x74, 0x0a, 0x18, 0x41, 0x6c, 0x6c, 0x6f, 0x77,
	0x65, 0x64, 0x43, 0x61, 0x70, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x69, 0x65, 0x73, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x42, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74,
	0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x77, 0x61, 0x73, 0x6d, 0x2e, 0x76, 0x33, 0x2e,
	0x53, 0x61, 0x6e, 0x69, 0x74, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x14, 0x0a,
	0x12, 0x53, 0x61, 0x6e, 0x69, 0x74, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x22, 0x81, 0x03, 0x0a, 0x08, 0x56, 0x6d, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x12, 0x13, 0x0a, 0x05, 0x76, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x76, 0x6d, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x07, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52,
	0x07, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x39, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x33, 0x2e, 0x41, 0x73,
	0x79, 0x6e, 0x63, 0x44, 0x61, 0x74, 0x61, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x04, 0x63,
	0x6f, 0x64, 0x65, 0x12, 0x3a, 0x0a, 0x0d, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79,
	0x52, 0x0d, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x2b, 0x0a, 0x11, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x5f, 0x70, 0x72, 0x65, 0x63, 0x6f, 0x6d, 0x70,
	0x69, 0x6c, 0x65, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x10, 0x61, 0x6c, 0x6c, 0x6f,
	0x77, 0x50, 0x72, 0x65, 0x63, 0x6f, 0x6d, 0x70, 0x69, 0x6c, 0x65, 0x64, 0x12, 0x34, 0x0a, 0x17,
	0x6e, 0x61, 0x63, 0x6b, 0x5f, 0x6f, 0x6e, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x5f, 0x63, 0x61, 0x63,
	0x68, 0x65, 0x5f, 0x6d, 0x69, 0x73, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x13, 0x6e,
	0x61, 0x63, 0x6b, 0x4f, 0x6e, 0x43, 0x6f, 0x64, 0x65, 0x43, 0x61, 0x63, 0x68, 0x65, 0x4d, 0x69,
	0x73, 0x73, 0x12, 0x63, 0x0a, 0x15, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e,
	0x74, 0x5f, 0x76, 0x61, 0x72, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x2e, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x77, 0x61, 0x73, 0x6d, 0x2e, 0x76, 0x33, 0x2e, 0x45, 0x6e, 0x76,
	0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x56, 0x61, 0x72, 0x69, 0x61, 0x62, 0x6c, 0x65,
	0x73, 0x52, 0x14, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x56, 0x61,
	0x72, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x73, 0x22, 0xd6, 0x01, 0x0a, 0x14, 0x45, 0x6e, 0x76, 0x69,
	0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x56, 0x61, 0x72, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x73,
	0x12, 0x22, 0x0a, 0x0d, 0x68, 0x6f, 0x73, 0x74, 0x5f, 0x65, 0x6e, 0x76, 0x5f, 0x6b, 0x65, 0x79,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0b, 0x68, 0x6f, 0x73, 0x74, 0x45, 0x6e, 0x76,
	0x4b, 0x65, 0x79, 0x73, 0x12, 0x5c, 0x0a, 0x0a, 0x6b, 0x65, 0x79, 0x5f, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x3d, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79,
	0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x77, 0x61, 0x73, 0x6d,
	0x2e, 0x76, 0x33, 0x2e, 0x45, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x56,
	0x61, 0x72, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x73, 0x2e, 0x4b, 0x65, 0x79, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x09, 0x6b, 0x65, 0x79, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x73, 0x1a, 0x3c, 0x0a, 0x0e, 0x4b, 0x65, 0x79, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01,
	0x22, 0xd8, 0x02, 0x0a, 0x0c, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x72, 0x6f, 0x6f, 0x74, 0x5f, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x6f, 0x6f, 0x74, 0x49, 0x64, 0x12, 0x41,
	0x0a, 0x09, 0x76, 0x6d, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x22, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x77, 0x61, 0x73, 0x6d, 0x2e, 0x76, 0x33, 0x2e, 0x56, 0x6d, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x48, 0x00, 0x52, 0x08, 0x76, 0x6d, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x12, 0x3a, 0x0a, 0x0d, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x0d,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1b, 0x0a,
	0x09, 0x66, 0x61, 0x69, 0x6c, 0x5f, 0x6f, 0x70, 0x65, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x08, 0x66, 0x61, 0x69, 0x6c, 0x4f, 0x70, 0x65, 0x6e, 0x12, 0x79, 0x0a, 0x1d, 0x63, 0x61,
	0x70, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x5f, 0x72, 0x65, 0x73, 0x74, 0x72, 0x69, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x35, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x77, 0x61, 0x73, 0x6d, 0x2e, 0x76, 0x33, 0x2e, 0x43, 0x61, 0x70,
	0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x52, 0x65, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x1b, 0x63, 0x61, 0x70, 0x61, 0x62, 0x69,
	0x6c, 0x69, 0x74, 0x79, 0x52, 0x65, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x42, 0x04, 0x0a, 0x02, 0x76, 0x6d, 0x22, 0x6b, 0x0a, 0x0b, 0x57,
	0x61, 0x73, 0x6d, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3e, 0x0a, 0x06, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x65, 0x6e, 0x76,
	0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x77, 0x61,
	0x73, 0x6d, 0x2e, 0x76, 0x33, 0x2e, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x52, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x69,
	0x6e, 0x67, 0x6c, 0x65, 0x74, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x73,
	0x69, 0x6e, 0x67, 0x6c, 0x65, 0x74, 0x6f, 0x6e, 0x42, 0x3d, 0x0a, 0x26, 0x69, 0x6f, 0x2e, 0x65,
	0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e,
	0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x77, 0x61, 0x73, 0x6d, 0x2e,
	0x76, 0x33, 0x42, 0x09, 0x57, 0x61, 0x73, 0x6d, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0xba,
	0x80, 0xc8, 0xd1, 0x06, 0x02, 0x10, 0x02, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_envoy_extensions_wasm_v3_wasm_proto_rawDescOnce sync.Once
	file_envoy_extensions_wasm_v3_wasm_proto_rawDescData = file_envoy_extensions_wasm_v3_wasm_proto_rawDesc
)

func file_envoy_extensions_wasm_v3_wasm_proto_rawDescGZIP() []byte {
	file_envoy_extensions_wasm_v3_wasm_proto_rawDescOnce.Do(func() {
		file_envoy_extensions_wasm_v3_wasm_proto_rawDescData = protoimpl.X.CompressGZIP(file_envoy_extensions_wasm_v3_wasm_proto_rawDescData)
	})
	return file_envoy_extensions_wasm_v3_wasm_proto_rawDescData
}

var file_envoy_extensions_wasm_v3_wasm_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_envoy_extensions_wasm_v3_wasm_proto_goTypes = []interface{}{
	(*CapabilityRestrictionConfig)(nil), // 0: envoy.extensions.wasm.v3.CapabilityRestrictionConfig
	(*SanitizationConfig)(nil),          // 1: envoy.extensions.wasm.v3.SanitizationConfig
	(*VmConfig)(nil),                    // 2: envoy.extensions.wasm.v3.VmConfig
	(*EnvironmentVariables)(nil),        // 3: envoy.extensions.wasm.v3.EnvironmentVariables
	(*PluginConfig)(nil),                // 4: envoy.extensions.wasm.v3.PluginConfig
	(*WasmService)(nil),                 // 5: envoy.extensions.wasm.v3.WasmService
	nil,                                 // 6: envoy.extensions.wasm.v3.CapabilityRestrictionConfig.AllowedCapabilitiesEntry
	nil,                                 // 7: envoy.extensions.wasm.v3.EnvironmentVariables.KeyValuesEntry
	(*v3.AsyncDataSource)(nil),          // 8: envoy.config.core.v3.AsyncDataSource
	(*any.Any)(nil),                     // 9: google.protobuf.Any
}
var file_envoy_extensions_wasm_v3_wasm_proto_depIdxs = []int32{
	6,  // 0: envoy.extensions.wasm.v3.CapabilityRestrictionConfig.allowed_capabilities:type_name -> envoy.extensions.wasm.v3.CapabilityRestrictionConfig.AllowedCapabilitiesEntry
	8,  // 1: envoy.extensions.wasm.v3.VmConfig.code:type_name -> envoy.config.core.v3.AsyncDataSource
	9,  // 2: envoy.extensions.wasm.v3.VmConfig.configuration:type_name -> google.protobuf.Any
	3,  // 3: envoy.extensions.wasm.v3.VmConfig.environment_variables:type_name -> envoy.extensions.wasm.v3.EnvironmentVariables
	7,  // 4: envoy.extensions.wasm.v3.EnvironmentVariables.key_values:type_name -> envoy.extensions.wasm.v3.EnvironmentVariables.KeyValuesEntry
	2,  // 5: envoy.extensions.wasm.v3.PluginConfig.vm_config:type_name -> envoy.extensions.wasm.v3.VmConfig
	9,  // 6: envoy.extensions.wasm.v3.PluginConfig.configuration:type_name -> google.protobuf.Any
	0,  // 7: envoy.extensions.wasm.v3.PluginConfig.capability_restriction_config:type_name -> envoy.extensions.wasm.v3.CapabilityRestrictionConfig
	4,  // 8: envoy.extensions.wasm.v3.WasmService.config:type_name -> envoy.extensions.wasm.v3.PluginConfig
	1,  // 9: envoy.extensions.wasm.v3.CapabilityRestrictionConfig.AllowedCapabilitiesEntry.value:type_name -> envoy.extensions.wasm.v3.SanitizationConfig
	10, // [10:10] is the sub-list for method output_type
	10, // [10:10] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_envoy_extensions_wasm_v3_wasm_proto_init() }
func file_envoy_extensions_wasm_v3_wasm_proto_init() {
	if File_envoy_extensions_wasm_v3_wasm_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_envoy_extensions_wasm_v3_wasm_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CapabilityRestrictionConfig); i {
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
		file_envoy_extensions_wasm_v3_wasm_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SanitizationConfig); i {
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
		file_envoy_extensions_wasm_v3_wasm_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VmConfig); i {
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
		file_envoy_extensions_wasm_v3_wasm_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EnvironmentVariables); i {
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
		file_envoy_extensions_wasm_v3_wasm_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PluginConfig); i {
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
		file_envoy_extensions_wasm_v3_wasm_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WasmService); i {
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
	file_envoy_extensions_wasm_v3_wasm_proto_msgTypes[4].OneofWrappers = []interface{}{
		(*PluginConfig_VmConfig)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_envoy_extensions_wasm_v3_wasm_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_envoy_extensions_wasm_v3_wasm_proto_goTypes,
		DependencyIndexes: file_envoy_extensions_wasm_v3_wasm_proto_depIdxs,
		MessageInfos:      file_envoy_extensions_wasm_v3_wasm_proto_msgTypes,
	}.Build()
	File_envoy_extensions_wasm_v3_wasm_proto = out.File
	file_envoy_extensions_wasm_v3_wasm_proto_rawDesc = nil
	file_envoy_extensions_wasm_v3_wasm_proto_goTypes = nil
	file_envoy_extensions_wasm_v3_wasm_proto_depIdxs = nil
}
