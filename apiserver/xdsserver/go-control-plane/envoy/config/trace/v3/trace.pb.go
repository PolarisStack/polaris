// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/config/trace/v3/trace.proto

package envoy_config_trace_v3

import (
	fmt "fmt"
	v1 "github.com/census-instrumentation/opencensus-proto/gen-go/trace/v1"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	_ "github.com/polarismesh/polaris-server/apiserver/xdsserver/go-control-plane/envoy/annotations"
	v3 "github.com/polarismesh/polaris-server/apiserver/xdsserver/go-control-plane/envoy/config/core/v3"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	any "github.com/golang/protobuf/ptypes/any"
	_struct "github.com/golang/protobuf/ptypes/struct"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type LightstepConfig_PropagationMode int32

const (
	LightstepConfig_ENVOY         LightstepConfig_PropagationMode = 0
	LightstepConfig_LIGHTSTEP     LightstepConfig_PropagationMode = 1
	LightstepConfig_B3            LightstepConfig_PropagationMode = 2
	LightstepConfig_TRACE_CONTEXT LightstepConfig_PropagationMode = 3
)

var LightstepConfig_PropagationMode_name = map[int32]string{
	0: "ENVOY",
	1: "LIGHTSTEP",
	2: "B3",
	3: "TRACE_CONTEXT",
}

var LightstepConfig_PropagationMode_value = map[string]int32{
	"ENVOY":         0,
	"LIGHTSTEP":     1,
	"B3":            2,
	"TRACE_CONTEXT": 3,
}

func (x LightstepConfig_PropagationMode) String() string {
	return proto.EnumName(LightstepConfig_PropagationMode_name, int32(x))
}

func (LightstepConfig_PropagationMode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_a8ab8dce64ff7a05, []int{1, 0}
}

type ZipkinConfig_CollectorEndpointVersion int32

const (
	ZipkinConfig_hidden_envoy_deprecated_HTTP_JSON_V1 ZipkinConfig_CollectorEndpointVersion = 0 // Deprecated: Do not use.
	ZipkinConfig_HTTP_JSON                            ZipkinConfig_CollectorEndpointVersion = 1
	ZipkinConfig_HTTP_PROTO                           ZipkinConfig_CollectorEndpointVersion = 2
	ZipkinConfig_GRPC                                 ZipkinConfig_CollectorEndpointVersion = 3
)

var ZipkinConfig_CollectorEndpointVersion_name = map[int32]string{
	0: "hidden_envoy_deprecated_HTTP_JSON_V1",
	1: "HTTP_JSON",
	2: "HTTP_PROTO",
	3: "GRPC",
}

var ZipkinConfig_CollectorEndpointVersion_value = map[string]int32{
	"hidden_envoy_deprecated_HTTP_JSON_V1": 0,
	"HTTP_JSON":                            1,
	"HTTP_PROTO":                           2,
	"GRPC":                                 3,
}

func (x ZipkinConfig_CollectorEndpointVersion) String() string {
	return proto.EnumName(ZipkinConfig_CollectorEndpointVersion_name, int32(x))
}

func (ZipkinConfig_CollectorEndpointVersion) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_a8ab8dce64ff7a05, []int{2, 0}
}

type OpenCensusConfig_TraceContext int32

const (
	OpenCensusConfig_NONE                OpenCensusConfig_TraceContext = 0
	OpenCensusConfig_TRACE_CONTEXT       OpenCensusConfig_TraceContext = 1
	OpenCensusConfig_GRPC_TRACE_BIN      OpenCensusConfig_TraceContext = 2
	OpenCensusConfig_CLOUD_TRACE_CONTEXT OpenCensusConfig_TraceContext = 3
	OpenCensusConfig_B3                  OpenCensusConfig_TraceContext = 4
)

var OpenCensusConfig_TraceContext_name = map[int32]string{
	0: "NONE",
	1: "TRACE_CONTEXT",
	2: "GRPC_TRACE_BIN",
	3: "CLOUD_TRACE_CONTEXT",
	4: "B3",
}

var OpenCensusConfig_TraceContext_value = map[string]int32{
	"NONE":                0,
	"TRACE_CONTEXT":       1,
	"GRPC_TRACE_BIN":      2,
	"CLOUD_TRACE_CONTEXT": 3,
	"B3":                  4,
}

func (x OpenCensusConfig_TraceContext) String() string {
	return proto.EnumName(OpenCensusConfig_TraceContext_name, int32(x))
}

func (OpenCensusConfig_TraceContext) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_a8ab8dce64ff7a05, []int{5, 0}
}

type Tracing struct {
	Http                 *Tracing_Http `protobuf:"bytes,1,opt,name=http,proto3" json:"http,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *Tracing) Reset()         { *m = Tracing{} }
func (m *Tracing) String() string { return proto.CompactTextString(m) }
func (*Tracing) ProtoMessage()    {}
func (*Tracing) Descriptor() ([]byte, []int) {
	return fileDescriptor_a8ab8dce64ff7a05, []int{0}
}

func (m *Tracing) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Tracing.Unmarshal(m, b)
}
func (m *Tracing) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Tracing.Marshal(b, m, deterministic)
}
func (m *Tracing) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Tracing.Merge(m, src)
}
func (m *Tracing) XXX_Size() int {
	return xxx_messageInfo_Tracing.Size(m)
}
func (m *Tracing) XXX_DiscardUnknown() {
	xxx_messageInfo_Tracing.DiscardUnknown(m)
}

var xxx_messageInfo_Tracing proto.InternalMessageInfo

func (m *Tracing) GetHttp() *Tracing_Http {
	if m != nil {
		return m.Http
	}
	return nil
}

type Tracing_Http struct {
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Types that are valid to be assigned to ConfigType:
	//	*Tracing_Http_HiddenEnvoyDeprecatedConfig
	//	*Tracing_Http_TypedConfig
	ConfigType           isTracing_Http_ConfigType `protobuf_oneof:"config_type"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *Tracing_Http) Reset()         { *m = Tracing_Http{} }
func (m *Tracing_Http) String() string { return proto.CompactTextString(m) }
func (*Tracing_Http) ProtoMessage()    {}
func (*Tracing_Http) Descriptor() ([]byte, []int) {
	return fileDescriptor_a8ab8dce64ff7a05, []int{0, 0}
}

func (m *Tracing_Http) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Tracing_Http.Unmarshal(m, b)
}
func (m *Tracing_Http) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Tracing_Http.Marshal(b, m, deterministic)
}
func (m *Tracing_Http) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Tracing_Http.Merge(m, src)
}
func (m *Tracing_Http) XXX_Size() int {
	return xxx_messageInfo_Tracing_Http.Size(m)
}
func (m *Tracing_Http) XXX_DiscardUnknown() {
	xxx_messageInfo_Tracing_Http.DiscardUnknown(m)
}

var xxx_messageInfo_Tracing_Http proto.InternalMessageInfo

func (m *Tracing_Http) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type isTracing_Http_ConfigType interface {
	isTracing_Http_ConfigType()
}

type Tracing_Http_HiddenEnvoyDeprecatedConfig struct {
	HiddenEnvoyDeprecatedConfig *_struct.Struct `protobuf:"bytes,2,opt,name=hidden_envoy_deprecated_config,json=hiddenEnvoyDeprecatedConfig,proto3,oneof"`
}

type Tracing_Http_TypedConfig struct {
	TypedConfig *any.Any `protobuf:"bytes,3,opt,name=typed_config,json=typedConfig,proto3,oneof"`
}

func (*Tracing_Http_HiddenEnvoyDeprecatedConfig) isTracing_Http_ConfigType() {}

func (*Tracing_Http_TypedConfig) isTracing_Http_ConfigType() {}

func (m *Tracing_Http) GetConfigType() isTracing_Http_ConfigType {
	if m != nil {
		return m.ConfigType
	}
	return nil
}

// Deprecated: Do not use.
func (m *Tracing_Http) GetHiddenEnvoyDeprecatedConfig() *_struct.Struct {
	if x, ok := m.GetConfigType().(*Tracing_Http_HiddenEnvoyDeprecatedConfig); ok {
		return x.HiddenEnvoyDeprecatedConfig
	}
	return nil
}

func (m *Tracing_Http) GetTypedConfig() *any.Any {
	if x, ok := m.GetConfigType().(*Tracing_Http_TypedConfig); ok {
		return x.TypedConfig
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Tracing_Http) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Tracing_Http_HiddenEnvoyDeprecatedConfig)(nil),
		(*Tracing_Http_TypedConfig)(nil),
	}
}

type LightstepConfig struct {
	CollectorCluster     string                            `protobuf:"bytes,1,opt,name=collector_cluster,json=collectorCluster,proto3" json:"collector_cluster,omitempty"`
	AccessTokenFile      string                            `protobuf:"bytes,2,opt,name=access_token_file,json=accessTokenFile,proto3" json:"access_token_file,omitempty"`
	PropagationModes     []LightstepConfig_PropagationMode `protobuf:"varint,3,rep,packed,name=propagation_modes,json=propagationModes,proto3,enum=envoy.config.trace.v3.LightstepConfig_PropagationMode" json:"propagation_modes,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                          `json:"-"`
	XXX_unrecognized     []byte                            `json:"-"`
	XXX_sizecache        int32                             `json:"-"`
}

func (m *LightstepConfig) Reset()         { *m = LightstepConfig{} }
func (m *LightstepConfig) String() string { return proto.CompactTextString(m) }
func (*LightstepConfig) ProtoMessage()    {}
func (*LightstepConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_a8ab8dce64ff7a05, []int{1}
}

func (m *LightstepConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LightstepConfig.Unmarshal(m, b)
}
func (m *LightstepConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LightstepConfig.Marshal(b, m, deterministic)
}
func (m *LightstepConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LightstepConfig.Merge(m, src)
}
func (m *LightstepConfig) XXX_Size() int {
	return xxx_messageInfo_LightstepConfig.Size(m)
}
func (m *LightstepConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_LightstepConfig.DiscardUnknown(m)
}

var xxx_messageInfo_LightstepConfig proto.InternalMessageInfo

func (m *LightstepConfig) GetCollectorCluster() string {
	if m != nil {
		return m.CollectorCluster
	}
	return ""
}

func (m *LightstepConfig) GetAccessTokenFile() string {
	if m != nil {
		return m.AccessTokenFile
	}
	return ""
}

func (m *LightstepConfig) GetPropagationModes() []LightstepConfig_PropagationMode {
	if m != nil {
		return m.PropagationModes
	}
	return nil
}

type ZipkinConfig struct {
	CollectorCluster         string                                `protobuf:"bytes,1,opt,name=collector_cluster,json=collectorCluster,proto3" json:"collector_cluster,omitempty"`
	CollectorEndpoint        string                                `protobuf:"bytes,2,opt,name=collector_endpoint,json=collectorEndpoint,proto3" json:"collector_endpoint,omitempty"`
	TraceId_128Bit           bool                                  `protobuf:"varint,3,opt,name=trace_id_128bit,json=traceId128bit,proto3" json:"trace_id_128bit,omitempty"`
	SharedSpanContext        *wrappers.BoolValue                   `protobuf:"bytes,4,opt,name=shared_span_context,json=sharedSpanContext,proto3" json:"shared_span_context,omitempty"`
	CollectorEndpointVersion ZipkinConfig_CollectorEndpointVersion `protobuf:"varint,5,opt,name=collector_endpoint_version,json=collectorEndpointVersion,proto3,enum=envoy.config.trace.v3.ZipkinConfig_CollectorEndpointVersion" json:"collector_endpoint_version,omitempty"`
	XXX_NoUnkeyedLiteral     struct{}                              `json:"-"`
	XXX_unrecognized         []byte                                `json:"-"`
	XXX_sizecache            int32                                 `json:"-"`
}

func (m *ZipkinConfig) Reset()         { *m = ZipkinConfig{} }
func (m *ZipkinConfig) String() string { return proto.CompactTextString(m) }
func (*ZipkinConfig) ProtoMessage()    {}
func (*ZipkinConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_a8ab8dce64ff7a05, []int{2}
}

func (m *ZipkinConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ZipkinConfig.Unmarshal(m, b)
}
func (m *ZipkinConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ZipkinConfig.Marshal(b, m, deterministic)
}
func (m *ZipkinConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ZipkinConfig.Merge(m, src)
}
func (m *ZipkinConfig) XXX_Size() int {
	return xxx_messageInfo_ZipkinConfig.Size(m)
}
func (m *ZipkinConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_ZipkinConfig.DiscardUnknown(m)
}

var xxx_messageInfo_ZipkinConfig proto.InternalMessageInfo

func (m *ZipkinConfig) GetCollectorCluster() string {
	if m != nil {
		return m.CollectorCluster
	}
	return ""
}

func (m *ZipkinConfig) GetCollectorEndpoint() string {
	if m != nil {
		return m.CollectorEndpoint
	}
	return ""
}

func (m *ZipkinConfig) GetTraceId_128Bit() bool {
	if m != nil {
		return m.TraceId_128Bit
	}
	return false
}

func (m *ZipkinConfig) GetSharedSpanContext() *wrappers.BoolValue {
	if m != nil {
		return m.SharedSpanContext
	}
	return nil
}

func (m *ZipkinConfig) GetCollectorEndpointVersion() ZipkinConfig_CollectorEndpointVersion {
	if m != nil {
		return m.CollectorEndpointVersion
	}
	return ZipkinConfig_hidden_envoy_deprecated_HTTP_JSON_V1
}

type DynamicOtConfig struct {
	Library              string          `protobuf:"bytes,1,opt,name=library,proto3" json:"library,omitempty"`
	Config               *_struct.Struct `protobuf:"bytes,2,opt,name=config,proto3" json:"config,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *DynamicOtConfig) Reset()         { *m = DynamicOtConfig{} }
func (m *DynamicOtConfig) String() string { return proto.CompactTextString(m) }
func (*DynamicOtConfig) ProtoMessage()    {}
func (*DynamicOtConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_a8ab8dce64ff7a05, []int{3}
}

func (m *DynamicOtConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DynamicOtConfig.Unmarshal(m, b)
}
func (m *DynamicOtConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DynamicOtConfig.Marshal(b, m, deterministic)
}
func (m *DynamicOtConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DynamicOtConfig.Merge(m, src)
}
func (m *DynamicOtConfig) XXX_Size() int {
	return xxx_messageInfo_DynamicOtConfig.Size(m)
}
func (m *DynamicOtConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_DynamicOtConfig.DiscardUnknown(m)
}

var xxx_messageInfo_DynamicOtConfig proto.InternalMessageInfo

func (m *DynamicOtConfig) GetLibrary() string {
	if m != nil {
		return m.Library
	}
	return ""
}

func (m *DynamicOtConfig) GetConfig() *_struct.Struct {
	if m != nil {
		return m.Config
	}
	return nil
}

type DatadogConfig struct {
	CollectorCluster     string   `protobuf:"bytes,1,opt,name=collector_cluster,json=collectorCluster,proto3" json:"collector_cluster,omitempty"`
	ServiceName          string   `protobuf:"bytes,2,opt,name=service_name,json=serviceName,proto3" json:"service_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DatadogConfig) Reset()         { *m = DatadogConfig{} }
func (m *DatadogConfig) String() string { return proto.CompactTextString(m) }
func (*DatadogConfig) ProtoMessage()    {}
func (*DatadogConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_a8ab8dce64ff7a05, []int{4}
}

func (m *DatadogConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DatadogConfig.Unmarshal(m, b)
}
func (m *DatadogConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DatadogConfig.Marshal(b, m, deterministic)
}
func (m *DatadogConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DatadogConfig.Merge(m, src)
}
func (m *DatadogConfig) XXX_Size() int {
	return xxx_messageInfo_DatadogConfig.Size(m)
}
func (m *DatadogConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_DatadogConfig.DiscardUnknown(m)
}

var xxx_messageInfo_DatadogConfig proto.InternalMessageInfo

func (m *DatadogConfig) GetCollectorCluster() string {
	if m != nil {
		return m.CollectorCluster
	}
	return ""
}

func (m *DatadogConfig) GetServiceName() string {
	if m != nil {
		return m.ServiceName
	}
	return ""
}

type OpenCensusConfig struct {
	TraceConfig                *v1.TraceConfig                 `protobuf:"bytes,1,opt,name=trace_config,json=traceConfig,proto3" json:"trace_config,omitempty"`
	StdoutExporterEnabled      bool                            `protobuf:"varint,2,opt,name=stdout_exporter_enabled,json=stdoutExporterEnabled,proto3" json:"stdout_exporter_enabled,omitempty"`
	StackdriverExporterEnabled bool                            `protobuf:"varint,3,opt,name=stackdriver_exporter_enabled,json=stackdriverExporterEnabled,proto3" json:"stackdriver_exporter_enabled,omitempty"`
	StackdriverProjectId       string                          `protobuf:"bytes,4,opt,name=stackdriver_project_id,json=stackdriverProjectId,proto3" json:"stackdriver_project_id,omitempty"`
	StackdriverAddress         string                          `protobuf:"bytes,10,opt,name=stackdriver_address,json=stackdriverAddress,proto3" json:"stackdriver_address,omitempty"`
	StackdriverGrpcService     *v3.GrpcService                 `protobuf:"bytes,13,opt,name=stackdriver_grpc_service,json=stackdriverGrpcService,proto3" json:"stackdriver_grpc_service,omitempty"`
	ZipkinExporterEnabled      bool                            `protobuf:"varint,5,opt,name=zipkin_exporter_enabled,json=zipkinExporterEnabled,proto3" json:"zipkin_exporter_enabled,omitempty"`
	ZipkinUrl                  string                          `protobuf:"bytes,6,opt,name=zipkin_url,json=zipkinUrl,proto3" json:"zipkin_url,omitempty"`
	OcagentExporterEnabled     bool                            `protobuf:"varint,11,opt,name=ocagent_exporter_enabled,json=ocagentExporterEnabled,proto3" json:"ocagent_exporter_enabled,omitempty"`
	OcagentAddress             string                          `protobuf:"bytes,12,opt,name=ocagent_address,json=ocagentAddress,proto3" json:"ocagent_address,omitempty"`
	OcagentGrpcService         *v3.GrpcService                 `protobuf:"bytes,14,opt,name=ocagent_grpc_service,json=ocagentGrpcService,proto3" json:"ocagent_grpc_service,omitempty"`
	IncomingTraceContext       []OpenCensusConfig_TraceContext `protobuf:"varint,8,rep,packed,name=incoming_trace_context,json=incomingTraceContext,proto3,enum=envoy.config.trace.v3.OpenCensusConfig_TraceContext" json:"incoming_trace_context,omitempty"`
	OutgoingTraceContext       []OpenCensusConfig_TraceContext `protobuf:"varint,9,rep,packed,name=outgoing_trace_context,json=outgoingTraceContext,proto3,enum=envoy.config.trace.v3.OpenCensusConfig_TraceContext" json:"outgoing_trace_context,omitempty"`
	XXX_NoUnkeyedLiteral       struct{}                        `json:"-"`
	XXX_unrecognized           []byte                          `json:"-"`
	XXX_sizecache              int32                           `json:"-"`
}

func (m *OpenCensusConfig) Reset()         { *m = OpenCensusConfig{} }
func (m *OpenCensusConfig) String() string { return proto.CompactTextString(m) }
func (*OpenCensusConfig) ProtoMessage()    {}
func (*OpenCensusConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_a8ab8dce64ff7a05, []int{5}
}

func (m *OpenCensusConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OpenCensusConfig.Unmarshal(m, b)
}
func (m *OpenCensusConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OpenCensusConfig.Marshal(b, m, deterministic)
}
func (m *OpenCensusConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OpenCensusConfig.Merge(m, src)
}
func (m *OpenCensusConfig) XXX_Size() int {
	return xxx_messageInfo_OpenCensusConfig.Size(m)
}
func (m *OpenCensusConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_OpenCensusConfig.DiscardUnknown(m)
}

var xxx_messageInfo_OpenCensusConfig proto.InternalMessageInfo

func (m *OpenCensusConfig) GetTraceConfig() *v1.TraceConfig {
	if m != nil {
		return m.TraceConfig
	}
	return nil
}

func (m *OpenCensusConfig) GetStdoutExporterEnabled() bool {
	if m != nil {
		return m.StdoutExporterEnabled
	}
	return false
}

func (m *OpenCensusConfig) GetStackdriverExporterEnabled() bool {
	if m != nil {
		return m.StackdriverExporterEnabled
	}
	return false
}

func (m *OpenCensusConfig) GetStackdriverProjectId() string {
	if m != nil {
		return m.StackdriverProjectId
	}
	return ""
}

func (m *OpenCensusConfig) GetStackdriverAddress() string {
	if m != nil {
		return m.StackdriverAddress
	}
	return ""
}

func (m *OpenCensusConfig) GetStackdriverGrpcService() *v3.GrpcService {
	if m != nil {
		return m.StackdriverGrpcService
	}
	return nil
}

func (m *OpenCensusConfig) GetZipkinExporterEnabled() bool {
	if m != nil {
		return m.ZipkinExporterEnabled
	}
	return false
}

func (m *OpenCensusConfig) GetZipkinUrl() string {
	if m != nil {
		return m.ZipkinUrl
	}
	return ""
}

func (m *OpenCensusConfig) GetOcagentExporterEnabled() bool {
	if m != nil {
		return m.OcagentExporterEnabled
	}
	return false
}

func (m *OpenCensusConfig) GetOcagentAddress() string {
	if m != nil {
		return m.OcagentAddress
	}
	return ""
}

func (m *OpenCensusConfig) GetOcagentGrpcService() *v3.GrpcService {
	if m != nil {
		return m.OcagentGrpcService
	}
	return nil
}

func (m *OpenCensusConfig) GetIncomingTraceContext() []OpenCensusConfig_TraceContext {
	if m != nil {
		return m.IncomingTraceContext
	}
	return nil
}

func (m *OpenCensusConfig) GetOutgoingTraceContext() []OpenCensusConfig_TraceContext {
	if m != nil {
		return m.OutgoingTraceContext
	}
	return nil
}

type TraceServiceConfig struct {
	GrpcService          *v3.GrpcService `protobuf:"bytes,1,opt,name=grpc_service,json=grpcService,proto3" json:"grpc_service,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *TraceServiceConfig) Reset()         { *m = TraceServiceConfig{} }
func (m *TraceServiceConfig) String() string { return proto.CompactTextString(m) }
func (*TraceServiceConfig) ProtoMessage()    {}
func (*TraceServiceConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_a8ab8dce64ff7a05, []int{6}
}

func (m *TraceServiceConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TraceServiceConfig.Unmarshal(m, b)
}
func (m *TraceServiceConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TraceServiceConfig.Marshal(b, m, deterministic)
}
func (m *TraceServiceConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TraceServiceConfig.Merge(m, src)
}
func (m *TraceServiceConfig) XXX_Size() int {
	return xxx_messageInfo_TraceServiceConfig.Size(m)
}
func (m *TraceServiceConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_TraceServiceConfig.DiscardUnknown(m)
}

var xxx_messageInfo_TraceServiceConfig proto.InternalMessageInfo

func (m *TraceServiceConfig) GetGrpcService() *v3.GrpcService {
	if m != nil {
		return m.GrpcService
	}
	return nil
}

func init() {
	proto.RegisterEnum("envoy.config.trace.v3.LightstepConfig_PropagationMode", LightstepConfig_PropagationMode_name, LightstepConfig_PropagationMode_value)
	proto.RegisterEnum("envoy.config.trace.v3.ZipkinConfig_CollectorEndpointVersion", ZipkinConfig_CollectorEndpointVersion_name, ZipkinConfig_CollectorEndpointVersion_value)
	proto.RegisterEnum("envoy.config.trace.v3.OpenCensusConfig_TraceContext", OpenCensusConfig_TraceContext_name, OpenCensusConfig_TraceContext_value)
	proto.RegisterType((*Tracing)(nil), "envoy.config.trace.v3.Tracing")
	proto.RegisterType((*Tracing_Http)(nil), "envoy.config.trace.v3.Tracing.Http")
	proto.RegisterType((*LightstepConfig)(nil), "envoy.config.trace.v3.LightstepConfig")
	proto.RegisterType((*ZipkinConfig)(nil), "envoy.config.trace.v3.ZipkinConfig")
	proto.RegisterType((*DynamicOtConfig)(nil), "envoy.config.trace.v3.DynamicOtConfig")
	proto.RegisterType((*DatadogConfig)(nil), "envoy.config.trace.v3.DatadogConfig")
	proto.RegisterType((*OpenCensusConfig)(nil), "envoy.config.trace.v3.OpenCensusConfig")
	proto.RegisterType((*TraceServiceConfig)(nil), "envoy.config.trace.v3.TraceServiceConfig")
}

func init() { proto.RegisterFile("envoy/config/trace/v3/trace.proto", fileDescriptor_a8ab8dce64ff7a05) }

var fileDescriptor_a8ab8dce64ff7a05 = []byte{
	// 1302 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x56, 0x4d, 0x6f, 0xdb, 0xc6,
	0x16, 0x35, 0x69, 0xd9, 0x96, 0xaf, 0x24, 0x9b, 0x9e, 0x38, 0x31, 0x9f, 0xf2, 0xf1, 0x6c, 0x39,
	0x1f, 0x7e, 0x41, 0x9e, 0x04, 0xdb, 0x41, 0x5e, 0x9e, 0xdb, 0x45, 0x43, 0x59, 0x89, 0x1d, 0xa4,
	0x92, 0x40, 0x2b, 0x46, 0x3f, 0x16, 0xec, 0x98, 0x9c, 0xc8, 0x4c, 0x68, 0x0e, 0x31, 0x1c, 0x29,
	0x51, 0x56, 0x45, 0x57, 0x45, 0xbb, 0x29, 0x5a, 0x74, 0xd1, 0x02, 0xdd, 0xf7, 0x07, 0x74, 0xd5,
	0x7d, 0x81, 0x6e, 0xfb, 0x27, 0xda, 0x9f, 0x50, 0x20, 0xe8, 0xa2, 0xe0, 0x0c, 0x69, 0x51, 0x94,
	0xd5, 0x04, 0xc8, 0x8e, 0xbc, 0xe7, 0x9e, 0x33, 0x77, 0x0e, 0xef, 0xdc, 0x21, 0xac, 0x11, 0xbf,
	0x4f, 0x07, 0x35, 0x9b, 0xfa, 0x4f, 0xdc, 0x6e, 0x8d, 0x33, 0x6c, 0x93, 0x5a, 0x7f, 0x5b, 0x3e,
	0x54, 0x03, 0x46, 0x39, 0x45, 0xe7, 0x45, 0x4a, 0x55, 0xa6, 0x54, 0x25, 0xd2, 0xdf, 0x2e, 0xdf,
	0x18, 0x61, 0xda, 0x94, 0x09, 0x62, 0x97, 0x05, 0xb6, 0x15, 0x12, 0xd6, 0x77, 0x13, 0x7e, 0xf9,
	0x5f, 0x5d, 0x4a, 0xbb, 0x1e, 0xa9, 0x89, 0xb7, 0xa3, 0xde, 0x93, 0x1a, 0xf6, 0x07, 0x31, 0x74,
	0x29, 0x0b, 0x85, 0x9c, 0xf5, 0x6c, 0x1e, 0xa3, 0x57, 0xb2, 0xe8, 0x73, 0x86, 0x83, 0x80, 0xb0,
	0x30, 0xc6, 0x6f, 0xd1, 0x80, 0xf8, 0x36, 0xf1, 0xc3, 0x5e, 0x28, 0x73, 0x92, 0xfa, 0x37, 0xe5,
	0x83, 0x15, 0x97, 0x2c, 0xb3, 0xd7, 0x65, 0xbd, 0xd8, 0xf7, 0x29, 0xc7, 0xdc, 0xa5, 0x7e, 0x58,
	0x73, 0x48, 0xc0, 0x88, 0x2d, 0x5e, 0xe2, 0xa4, 0xcb, 0x3d, 0x27, 0xc0, 0x23, 0x39, 0x21, 0xc7,
	0xbc, 0x97, 0xac, 0xb8, 0x36, 0x06, 0xf7, 0x09, 0x0b, 0x5d, 0xea, 0xbb, 0x7e, 0xb2, 0xcc, 0x4a,
	0x1f, 0x7b, 0xae, 0x83, 0x39, 0xa9, 0x25, 0x0f, 0x12, 0xa8, 0xfc, 0xa1, 0xc2, 0x5c, 0x87, 0x61,
	0xdb, 0xf5, 0xbb, 0xe8, 0x7f, 0x90, 0x3b, 0xe6, 0x3c, 0xd0, 0x95, 0x55, 0x65, 0xa3, 0xb0, 0xb5,
	0x5e, 0x3d, 0xd3, 0xe1, 0x6a, 0x9c, 0x5d, 0xdd, 0xe3, 0x3c, 0x30, 0x05, 0xa1, 0xfc, 0x97, 0x02,
	0xb9, 0xe8, 0x15, 0x5d, 0x84, 0x9c, 0x8f, 0x4f, 0x88, 0x50, 0x98, 0x37, 0xe6, 0x5e, 0x19, 0x39,
	0xa6, 0xae, 0x2a, 0xa6, 0x08, 0xa2, 0x23, 0xb8, 0x72, 0xec, 0x3a, 0x0e, 0xf1, 0x2d, 0x21, 0x6c,
	0x25, 0xfb, 0x24, 0x4e, 0x6c, 0x89, 0xae, 0x8a, 0x85, 0x57, 0xaa, 0xd2, 0xe1, 0x6a, 0xe2, 0x70,
	0xf5, 0x40, 0xf8, 0x6f, 0xa8, 0xba, 0xb2, 0x37, 0x65, 0x5e, 0x94, 0x22, 0x8d, 0x48, 0x63, 0xf7,
	0x54, 0xa2, 0x2e, 0x14, 0xd0, 0xff, 0xa1, 0xc8, 0x07, 0xc1, 0x50, 0x71, 0x5a, 0x28, 0x2e, 0x8f,
	0x29, 0xde, 0xf3, 0x07, 0x7b, 0x53, 0x66, 0x41, 0xe4, 0x4a, 0xea, 0xce, 0x7f, 0xbe, 0xff, 0xe5,
	0xf3, 0x2b, 0x57, 0xa1, 0x72, 0xd6, 0xae, 0xb7, 0x46, 0x76, 0x6d, 0x94, 0xa0, 0x20, 0x71, 0x2b,
	0x12, 0xd8, 0xb9, 0x1a, 0x31, 0xff, 0x0d, 0x97, 0xff, 0x91, 0x59, 0xf9, 0x5d, 0x85, 0xc5, 0x47,
	0x6e, 0xf7, 0x98, 0x87, 0x9c, 0x04, 0x71, 0xb9, 0xb7, 0x61, 0xc9, 0xa6, 0x9e, 0x47, 0x6c, 0x4e,
	0x99, 0x65, 0x7b, 0xbd, 0x90, 0x13, 0x96, 0x35, 0x4f, 0x3b, 0xcd, 0xa8, 0xcb, 0x04, 0xb4, 0x0d,
	0x4b, 0xd8, 0xb6, 0x49, 0x18, 0x5a, 0x9c, 0x3e, 0x23, 0xbe, 0xf5, 0xc4, 0xf5, 0x88, 0xf0, 0x2e,
	0xc5, 0x5a, 0x94, 0x19, 0x9d, 0x28, 0xe1, 0xbe, 0xeb, 0x11, 0xc4, 0x60, 0x29, 0x60, 0x34, 0xc0,
	0x5d, 0xd1, 0x22, 0xd6, 0x09, 0x75, 0x48, 0xa8, 0x4f, 0xaf, 0x4e, 0x6f, 0x2c, 0x6c, 0xdd, 0x99,
	0xf0, 0xa5, 0x33, 0xd5, 0x56, 0xdb, 0x43, 0xfe, 0xfb, 0xd4, 0x21, 0x46, 0xe9, 0x95, 0x01, 0x5f,
	0x2b, 0x73, 0x95, 0x99, 0xcf, 0x14, 0x55, 0x53, 0x4c, 0x2d, 0x18, 0xc5, 0xc3, 0xca, 0x7d, 0x58,
	0xcc, 0x70, 0xd0, 0x3c, 0xcc, 0x34, 0x9a, 0x87, 0xad, 0x0f, 0xb5, 0x29, 0x54, 0x82, 0xf9, 0x47,
	0xfb, 0x0f, 0xf6, 0x3a, 0x07, 0x9d, 0x46, 0x5b, 0x53, 0xd0, 0x2c, 0xa8, 0xc6, 0xb6, 0xa6, 0xa2,
	0x25, 0x28, 0x75, 0xcc, 0x7b, 0xf5, 0x86, 0x55, 0x6f, 0x35, 0x3b, 0x8d, 0x0f, 0x3a, 0xda, 0xf4,
	0xce, 0xad, 0xc8, 0xe0, 0x1b, 0x70, 0xed, 0x6c, 0x83, 0x33, 0x65, 0x56, 0xbe, 0xca, 0x41, 0xf1,
	0x23, 0x37, 0x78, 0xe6, 0xfa, 0x6f, 0xe5, 0xf2, 0x1d, 0x40, 0x43, 0x16, 0xf1, 0x9d, 0x80, 0xba,
	0x3e, 0xcf, 0xda, 0x3c, 0x14, 0x6e, 0xc4, 0x19, 0xe8, 0x3a, 0x2c, 0xca, 0x73, 0xee, 0x3a, 0xd6,
	0xe6, 0xd6, 0xdd, 0x23, 0x97, 0x8b, 0x2e, 0xcc, 0x9b, 0x25, 0x11, 0xde, 0x77, 0x64, 0x10, 0x3d,
	0x84, 0x73, 0xe1, 0x31, 0x66, 0xc4, 0xb1, 0xc2, 0x00, 0xfb, 0x51, 0xc3, 0x72, 0xf2, 0x82, 0xeb,
	0x39, 0xd1, 0xb1, 0xe5, 0xb1, 0x8e, 0x35, 0x28, 0xf5, 0x0e, 0xb1, 0xd7, 0x23, 0xe6, 0x92, 0xa4,
	0x1d, 0x04, 0x38, 0xda, 0x60, 0x44, 0x42, 0x2f, 0xa1, 0x3c, 0x5e, 0xab, 0x15, 0x4f, 0x01, 0x7d,
	0x66, 0x55, 0xd9, 0x58, 0xd8, 0x7a, 0x77, 0xc2, 0x57, 0x4e, 0x5b, 0x55, 0xad, 0x67, 0xb7, 0x73,
	0x28, 0x35, 0x4c, 0xdd, 0x9e, 0x80, 0x54, 0x9e, 0x83, 0x3e, 0x89, 0x85, 0xb6, 0xe0, 0xea, 0xa4,
	0x23, 0xbf, 0xd7, 0xe9, 0xb4, 0xad, 0x87, 0x07, 0xad, 0xa6, 0x75, 0xb8, 0xa9, 0x4d, 0x95, 0xf3,
	0x3f, 0xfe, 0xf9, 0xd3, 0x97, 0xaa, 0x92, 0x57, 0xa2, 0xb6, 0x38, 0xc5, 0x34, 0x05, 0x2d, 0x00,
	0x88, 0xd7, 0xb6, 0xd9, 0xea, 0xb4, 0x34, 0x15, 0xe5, 0x21, 0xf7, 0xc0, 0x6c, 0xd7, 0xb5, 0xe9,
	0xd7, 0x1c, 0xd8, 0xf4, 0xb6, 0x2a, 0xdf, 0x28, 0xb0, 0xb8, 0x3b, 0xf0, 0xf1, 0x89, 0x6b, 0xb7,
	0x78, 0xdc, 0x15, 0x6b, 0x30, 0xe7, 0xb9, 0x47, 0x0c, 0xb3, 0x41, 0xb6, 0x17, 0x92, 0x38, 0xaa,
	0xc1, 0xec, 0x1b, 0x4d, 0x26, 0x33, 0x4e, 0x7b, 0x4d, 0xa3, 0x66, 0x2a, 0xa8, 0xfc, 0xa0, 0x40,
	0x69, 0x17, 0x73, 0xec, 0xd0, 0xee, 0x5b, 0x75, 0xea, 0x4d, 0x28, 0xc6, 0x77, 0x9b, 0x25, 0xa6,
	0x6f, 0xa6, 0x47, 0x0b, 0x31, 0xd8, 0xc4, 0x27, 0x64, 0xe7, 0x66, 0x54, 0xe1, 0x35, 0x58, 0x9f,
	0x50, 0x61, 0xba, 0x9a, 0xca, 0x77, 0x79, 0xd0, 0x5a, 0x01, 0xf1, 0xeb, 0xe2, 0x32, 0x8b, 0x4b,
	0xdc, 0x87, 0x62, 0xfa, 0x1a, 0x8b, 0x2f, 0x8b, 0xeb, 0xd5, 0xe1, 0xad, 0x27, 0xdd, 0x49, 0x44,
	0x37, 0xc5, 0x00, 0x24, 0x92, 0x6d, 0x16, 0xf8, 0xf0, 0x05, 0xdd, 0x81, 0x95, 0x90, 0x3b, 0xb4,
	0xc7, 0x2d, 0xf2, 0x22, 0xa0, 0x8c, 0x93, 0xa8, 0x77, 0xf1, 0x91, 0x47, 0x1c, 0xb1, 0x85, 0xbc,
	0x79, 0x5e, 0xc2, 0x8d, 0x18, 0x6d, 0x48, 0x10, 0xbd, 0x07, 0x97, 0x42, 0x8e, 0xed, 0x67, 0x0e,
	0x73, 0xfb, 0x11, 0x27, 0x4b, 0x96, 0xc7, 0xad, 0x9c, 0xca, 0xc9, 0x2a, 0xdc, 0x86, 0x0b, 0x69,
	0x85, 0x80, 0xd1, 0xa7, 0xc4, 0xe6, 0x96, 0xeb, 0x88, 0xe3, 0x37, 0x6f, 0x2e, 0xa7, 0xd0, 0xb6,
	0x04, 0xf7, 0x1d, 0x54, 0x83, 0x73, 0x69, 0x16, 0x76, 0x1c, 0x46, 0xc2, 0x50, 0x07, 0x41, 0x41,
	0x29, 0xe8, 0x9e, 0x44, 0xd0, 0xc7, 0xa0, 0xa7, 0x09, 0xe9, 0xbf, 0x10, 0xbd, 0x24, 0x7c, 0x5b,
	0x1b, 0x3d, 0x94, 0xd1, 0xff, 0x4a, 0x74, 0x26, 0x1f, 0xb0, 0xc0, 0x3e, 0x90, 0x89, 0x66, 0xba,
	0xd2, 0x54, 0x3c, 0x72, 0xef, 0xa5, 0xe8, 0xf1, 0x71, 0x03, 0x66, 0xa4, 0x7b, 0x12, 0xce, 0xee,
	0xfd, 0x32, 0x40, 0xcc, 0xeb, 0x31, 0x4f, 0x9f, 0x15, 0xc5, 0xcf, 0xcb, 0xc8, 0x63, 0xe6, 0xa1,
	0xbb, 0xa0, 0x53, 0x1b, 0x77, 0x89, 0x7f, 0xc6, 0x57, 0x29, 0x08, 0xdd, 0x0b, 0x31, 0x9e, 0x15,
	0xbe, 0x01, 0x8b, 0x09, 0x33, 0xb1, 0xa6, 0x28, 0xd4, 0x17, 0xe2, 0x70, 0x62, 0xcb, 0x01, 0x2c,
	0x27, 0x89, 0x23, 0x96, 0x2c, 0xbc, 0xa9, 0x25, 0x28, 0xa6, 0xa7, 0xed, 0x78, 0x0a, 0x17, 0x5c,
	0xdf, 0xa6, 0x27, 0xae, 0xdf, 0xb5, 0x4e, 0x1b, 0x54, 0x4c, 0xd4, 0xbc, 0xb8, 0xe4, 0x6e, 0x4f,
	0x18, 0x7f, 0xd9, 0x06, 0x3f, 0x6d, 0xd7, 0x88, 0x6b, 0x2e, 0x27, 0x9a, 0xe9, 0x68, 0xb4, 0x16,
	0xed, 0xf1, 0x2e, 0x1d, 0x5f, 0x6b, 0xfe, 0x6d, 0xd6, 0x4a, 0x34, 0xd3, 0xd1, 0xca, 0x27, 0x50,
	0x1c, 0x59, 0x3b, 0x0f, 0xb9, 0x66, 0xab, 0xd9, 0xd0, 0xa6, 0xc6, 0x2f, 0x4a, 0x05, 0x21, 0x58,
	0x88, 0x86, 0xa3, 0x25, 0xe3, 0xc6, 0x7e, 0x53, 0x53, 0xd1, 0x0a, 0x9c, 0xab, 0x3f, 0x6a, 0x3d,
	0xde, 0xb5, 0x32, 0xb7, 0x6a, 0x7c, 0xe1, 0xe6, 0x76, 0xfe, 0x1b, 0x8d, 0x84, 0x0d, 0xb8, 0x7e,
	0xf6, 0x48, 0xc8, 0xd6, 0xfc, 0x30, 0x97, 0x9f, 0xd3, 0xf2, 0x95, 0x6f, 0x15, 0x40, 0xa2, 0xae,
	0xd8, 0xff, 0xf8, 0x48, 0x37, 0xa1, 0x38, 0xf2, 0x49, 0x95, 0x37, 0xfc, 0xa4, 0x46, 0xfe, 0x95,
	0x31, 0xf3, 0x85, 0xf8, 0x8d, 0x28, 0x74, 0x87, 0xe1, 0x9d, 0x5a, 0x54, 0xdb, 0x4d, 0xd8, 0x98,
	0xfc, 0x6b, 0x35, 0x5a, 0x80, 0xf1, 0xce, 0xcf, 0x9f, 0xfe, 0xfa, 0xdb, 0xac, 0xaa, 0xa9, 0xb0,
	0xee, 0x52, 0xb9, 0x6c, 0xc0, 0xe8, 0x8b, 0xc1, 0xd9, 0x5f, 0xc4, 0x00, 0x21, 0xd1, 0x8e, 0x06,
	0x56, 0x5b, 0x39, 0x9a, 0x15, 0x93, 0x6b, 0xfb, 0xef, 0x00, 0x00, 0x00, 0xff, 0xff, 0xfb, 0x70,
	0x00, 0x91, 0x7f, 0x0c, 0x00, 0x00,
}
