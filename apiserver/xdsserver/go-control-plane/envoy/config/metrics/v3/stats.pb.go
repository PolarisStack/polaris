// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/config/metrics/v3/stats.proto

package envoy_config_metrics_v3

import (
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	v31 "github.com/polarismesh/polaris-server/apiserver/xdsserver/go-control-plane/envoy/config/core/v3"
	v3 "github.com/polarismesh/polaris-server/apiserver/xdsserver/go-control-plane/envoy/type/matcher/v3"
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

type StatsSink struct {
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Types that are valid to be assigned to ConfigType:
	//	*StatsSink_HiddenEnvoyDeprecatedConfig
	//	*StatsSink_TypedConfig
	ConfigType           isStatsSink_ConfigType `protobuf_oneof:"config_type"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *StatsSink) Reset()         { *m = StatsSink{} }
func (m *StatsSink) String() string { return proto.CompactTextString(m) }
func (*StatsSink) ProtoMessage()    {}
func (*StatsSink) Descriptor() ([]byte, []int) {
	return fileDescriptor_38edc339428e9971, []int{0}
}

func (m *StatsSink) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StatsSink.Unmarshal(m, b)
}
func (m *StatsSink) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StatsSink.Marshal(b, m, deterministic)
}
func (m *StatsSink) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StatsSink.Merge(m, src)
}
func (m *StatsSink) XXX_Size() int {
	return xxx_messageInfo_StatsSink.Size(m)
}
func (m *StatsSink) XXX_DiscardUnknown() {
	xxx_messageInfo_StatsSink.DiscardUnknown(m)
}

var xxx_messageInfo_StatsSink proto.InternalMessageInfo

func (m *StatsSink) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type isStatsSink_ConfigType interface {
	isStatsSink_ConfigType()
}

type StatsSink_HiddenEnvoyDeprecatedConfig struct {
	HiddenEnvoyDeprecatedConfig *_struct.Struct `protobuf:"bytes,2,opt,name=hidden_envoy_deprecated_config,json=hiddenEnvoyDeprecatedConfig,proto3,oneof"`
}

type StatsSink_TypedConfig struct {
	TypedConfig *any.Any `protobuf:"bytes,3,opt,name=typed_config,json=typedConfig,proto3,oneof"`
}

func (*StatsSink_HiddenEnvoyDeprecatedConfig) isStatsSink_ConfigType() {}

func (*StatsSink_TypedConfig) isStatsSink_ConfigType() {}

func (m *StatsSink) GetConfigType() isStatsSink_ConfigType {
	if m != nil {
		return m.ConfigType
	}
	return nil
}

// Deprecated: Do not use.
func (m *StatsSink) GetHiddenEnvoyDeprecatedConfig() *_struct.Struct {
	if x, ok := m.GetConfigType().(*StatsSink_HiddenEnvoyDeprecatedConfig); ok {
		return x.HiddenEnvoyDeprecatedConfig
	}
	return nil
}

func (m *StatsSink) GetTypedConfig() *any.Any {
	if x, ok := m.GetConfigType().(*StatsSink_TypedConfig); ok {
		return x.TypedConfig
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*StatsSink) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*StatsSink_HiddenEnvoyDeprecatedConfig)(nil),
		(*StatsSink_TypedConfig)(nil),
	}
}

type StatsConfig struct {
	StatsTags            []*TagSpecifier     `protobuf:"bytes,1,rep,name=stats_tags,json=statsTags,proto3" json:"stats_tags,omitempty"`
	UseAllDefaultTags    *wrappers.BoolValue `protobuf:"bytes,2,opt,name=use_all_default_tags,json=useAllDefaultTags,proto3" json:"use_all_default_tags,omitempty"`
	StatsMatcher         *StatsMatcher       `protobuf:"bytes,3,opt,name=stats_matcher,json=statsMatcher,proto3" json:"stats_matcher,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *StatsConfig) Reset()         { *m = StatsConfig{} }
func (m *StatsConfig) String() string { return proto.CompactTextString(m) }
func (*StatsConfig) ProtoMessage()    {}
func (*StatsConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_38edc339428e9971, []int{1}
}

func (m *StatsConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StatsConfig.Unmarshal(m, b)
}
func (m *StatsConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StatsConfig.Marshal(b, m, deterministic)
}
func (m *StatsConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StatsConfig.Merge(m, src)
}
func (m *StatsConfig) XXX_Size() int {
	return xxx_messageInfo_StatsConfig.Size(m)
}
func (m *StatsConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_StatsConfig.DiscardUnknown(m)
}

var xxx_messageInfo_StatsConfig proto.InternalMessageInfo

func (m *StatsConfig) GetStatsTags() []*TagSpecifier {
	if m != nil {
		return m.StatsTags
	}
	return nil
}

func (m *StatsConfig) GetUseAllDefaultTags() *wrappers.BoolValue {
	if m != nil {
		return m.UseAllDefaultTags
	}
	return nil
}

func (m *StatsConfig) GetStatsMatcher() *StatsMatcher {
	if m != nil {
		return m.StatsMatcher
	}
	return nil
}

type StatsMatcher struct {
	// Types that are valid to be assigned to StatsMatcher:
	//	*StatsMatcher_RejectAll
	//	*StatsMatcher_ExclusionList
	//	*StatsMatcher_InclusionList
	StatsMatcher         isStatsMatcher_StatsMatcher `protobuf_oneof:"stats_matcher"`
	XXX_NoUnkeyedLiteral struct{}                    `json:"-"`
	XXX_unrecognized     []byte                      `json:"-"`
	XXX_sizecache        int32                       `json:"-"`
}

func (m *StatsMatcher) Reset()         { *m = StatsMatcher{} }
func (m *StatsMatcher) String() string { return proto.CompactTextString(m) }
func (*StatsMatcher) ProtoMessage()    {}
func (*StatsMatcher) Descriptor() ([]byte, []int) {
	return fileDescriptor_38edc339428e9971, []int{2}
}

func (m *StatsMatcher) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StatsMatcher.Unmarshal(m, b)
}
func (m *StatsMatcher) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StatsMatcher.Marshal(b, m, deterministic)
}
func (m *StatsMatcher) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StatsMatcher.Merge(m, src)
}
func (m *StatsMatcher) XXX_Size() int {
	return xxx_messageInfo_StatsMatcher.Size(m)
}
func (m *StatsMatcher) XXX_DiscardUnknown() {
	xxx_messageInfo_StatsMatcher.DiscardUnknown(m)
}

var xxx_messageInfo_StatsMatcher proto.InternalMessageInfo

type isStatsMatcher_StatsMatcher interface {
	isStatsMatcher_StatsMatcher()
}

type StatsMatcher_RejectAll struct {
	RejectAll bool `protobuf:"varint,1,opt,name=reject_all,json=rejectAll,proto3,oneof"`
}

type StatsMatcher_ExclusionList struct {
	ExclusionList *v3.ListStringMatcher `protobuf:"bytes,2,opt,name=exclusion_list,json=exclusionList,proto3,oneof"`
}

type StatsMatcher_InclusionList struct {
	InclusionList *v3.ListStringMatcher `protobuf:"bytes,3,opt,name=inclusion_list,json=inclusionList,proto3,oneof"`
}

func (*StatsMatcher_RejectAll) isStatsMatcher_StatsMatcher() {}

func (*StatsMatcher_ExclusionList) isStatsMatcher_StatsMatcher() {}

func (*StatsMatcher_InclusionList) isStatsMatcher_StatsMatcher() {}

func (m *StatsMatcher) GetStatsMatcher() isStatsMatcher_StatsMatcher {
	if m != nil {
		return m.StatsMatcher
	}
	return nil
}

func (m *StatsMatcher) GetRejectAll() bool {
	if x, ok := m.GetStatsMatcher().(*StatsMatcher_RejectAll); ok {
		return x.RejectAll
	}
	return false
}

func (m *StatsMatcher) GetExclusionList() *v3.ListStringMatcher {
	if x, ok := m.GetStatsMatcher().(*StatsMatcher_ExclusionList); ok {
		return x.ExclusionList
	}
	return nil
}

func (m *StatsMatcher) GetInclusionList() *v3.ListStringMatcher {
	if x, ok := m.GetStatsMatcher().(*StatsMatcher_InclusionList); ok {
		return x.InclusionList
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*StatsMatcher) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*StatsMatcher_RejectAll)(nil),
		(*StatsMatcher_ExclusionList)(nil),
		(*StatsMatcher_InclusionList)(nil),
	}
}

type TagSpecifier struct {
	TagName string `protobuf:"bytes,1,opt,name=tag_name,json=tagName,proto3" json:"tag_name,omitempty"`
	// Types that are valid to be assigned to TagValue:
	//	*TagSpecifier_Regex
	//	*TagSpecifier_FixedValue
	TagValue             isTagSpecifier_TagValue `protobuf_oneof:"tag_value"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *TagSpecifier) Reset()         { *m = TagSpecifier{} }
func (m *TagSpecifier) String() string { return proto.CompactTextString(m) }
func (*TagSpecifier) ProtoMessage()    {}
func (*TagSpecifier) Descriptor() ([]byte, []int) {
	return fileDescriptor_38edc339428e9971, []int{3}
}

func (m *TagSpecifier) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TagSpecifier.Unmarshal(m, b)
}
func (m *TagSpecifier) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TagSpecifier.Marshal(b, m, deterministic)
}
func (m *TagSpecifier) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TagSpecifier.Merge(m, src)
}
func (m *TagSpecifier) XXX_Size() int {
	return xxx_messageInfo_TagSpecifier.Size(m)
}
func (m *TagSpecifier) XXX_DiscardUnknown() {
	xxx_messageInfo_TagSpecifier.DiscardUnknown(m)
}

var xxx_messageInfo_TagSpecifier proto.InternalMessageInfo

func (m *TagSpecifier) GetTagName() string {
	if m != nil {
		return m.TagName
	}
	return ""
}

type isTagSpecifier_TagValue interface {
	isTagSpecifier_TagValue()
}

type TagSpecifier_Regex struct {
	Regex string `protobuf:"bytes,2,opt,name=regex,proto3,oneof"`
}

type TagSpecifier_FixedValue struct {
	FixedValue string `protobuf:"bytes,3,opt,name=fixed_value,json=fixedValue,proto3,oneof"`
}

func (*TagSpecifier_Regex) isTagSpecifier_TagValue() {}

func (*TagSpecifier_FixedValue) isTagSpecifier_TagValue() {}

func (m *TagSpecifier) GetTagValue() isTagSpecifier_TagValue {
	if m != nil {
		return m.TagValue
	}
	return nil
}

func (m *TagSpecifier) GetRegex() string {
	if x, ok := m.GetTagValue().(*TagSpecifier_Regex); ok {
		return x.Regex
	}
	return ""
}

func (m *TagSpecifier) GetFixedValue() string {
	if x, ok := m.GetTagValue().(*TagSpecifier_FixedValue); ok {
		return x.FixedValue
	}
	return ""
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*TagSpecifier) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*TagSpecifier_Regex)(nil),
		(*TagSpecifier_FixedValue)(nil),
	}
}

type StatsdSink struct {
	// Types that are valid to be assigned to StatsdSpecifier:
	//	*StatsdSink_Address
	//	*StatsdSink_TcpClusterName
	StatsdSpecifier      isStatsdSink_StatsdSpecifier `protobuf_oneof:"statsd_specifier"`
	Prefix               string                       `protobuf:"bytes,3,opt,name=prefix,proto3" json:"prefix,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                     `json:"-"`
	XXX_unrecognized     []byte                       `json:"-"`
	XXX_sizecache        int32                        `json:"-"`
}

func (m *StatsdSink) Reset()         { *m = StatsdSink{} }
func (m *StatsdSink) String() string { return proto.CompactTextString(m) }
func (*StatsdSink) ProtoMessage()    {}
func (*StatsdSink) Descriptor() ([]byte, []int) {
	return fileDescriptor_38edc339428e9971, []int{4}
}

func (m *StatsdSink) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StatsdSink.Unmarshal(m, b)
}
func (m *StatsdSink) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StatsdSink.Marshal(b, m, deterministic)
}
func (m *StatsdSink) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StatsdSink.Merge(m, src)
}
func (m *StatsdSink) XXX_Size() int {
	return xxx_messageInfo_StatsdSink.Size(m)
}
func (m *StatsdSink) XXX_DiscardUnknown() {
	xxx_messageInfo_StatsdSink.DiscardUnknown(m)
}

var xxx_messageInfo_StatsdSink proto.InternalMessageInfo

type isStatsdSink_StatsdSpecifier interface {
	isStatsdSink_StatsdSpecifier()
}

type StatsdSink_Address struct {
	Address *v31.Address `protobuf:"bytes,1,opt,name=address,proto3,oneof"`
}

type StatsdSink_TcpClusterName struct {
	TcpClusterName string `protobuf:"bytes,2,opt,name=tcp_cluster_name,json=tcpClusterName,proto3,oneof"`
}

func (*StatsdSink_Address) isStatsdSink_StatsdSpecifier() {}

func (*StatsdSink_TcpClusterName) isStatsdSink_StatsdSpecifier() {}

func (m *StatsdSink) GetStatsdSpecifier() isStatsdSink_StatsdSpecifier {
	if m != nil {
		return m.StatsdSpecifier
	}
	return nil
}

func (m *StatsdSink) GetAddress() *v31.Address {
	if x, ok := m.GetStatsdSpecifier().(*StatsdSink_Address); ok {
		return x.Address
	}
	return nil
}

func (m *StatsdSink) GetTcpClusterName() string {
	if x, ok := m.GetStatsdSpecifier().(*StatsdSink_TcpClusterName); ok {
		return x.TcpClusterName
	}
	return ""
}

func (m *StatsdSink) GetPrefix() string {
	if m != nil {
		return m.Prefix
	}
	return ""
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*StatsdSink) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*StatsdSink_Address)(nil),
		(*StatsdSink_TcpClusterName)(nil),
	}
}

type DogStatsdSink struct {
	// Types that are valid to be assigned to DogStatsdSpecifier:
	//	*DogStatsdSink_Address
	DogStatsdSpecifier   isDogStatsdSink_DogStatsdSpecifier `protobuf_oneof:"dog_statsd_specifier"`
	Prefix               string                             `protobuf:"bytes,3,opt,name=prefix,proto3" json:"prefix,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                           `json:"-"`
	XXX_unrecognized     []byte                             `json:"-"`
	XXX_sizecache        int32                              `json:"-"`
}

func (m *DogStatsdSink) Reset()         { *m = DogStatsdSink{} }
func (m *DogStatsdSink) String() string { return proto.CompactTextString(m) }
func (*DogStatsdSink) ProtoMessage()    {}
func (*DogStatsdSink) Descriptor() ([]byte, []int) {
	return fileDescriptor_38edc339428e9971, []int{5}
}

func (m *DogStatsdSink) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DogStatsdSink.Unmarshal(m, b)
}
func (m *DogStatsdSink) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DogStatsdSink.Marshal(b, m, deterministic)
}
func (m *DogStatsdSink) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DogStatsdSink.Merge(m, src)
}
func (m *DogStatsdSink) XXX_Size() int {
	return xxx_messageInfo_DogStatsdSink.Size(m)
}
func (m *DogStatsdSink) XXX_DiscardUnknown() {
	xxx_messageInfo_DogStatsdSink.DiscardUnknown(m)
}

var xxx_messageInfo_DogStatsdSink proto.InternalMessageInfo

type isDogStatsdSink_DogStatsdSpecifier interface {
	isDogStatsdSink_DogStatsdSpecifier()
}

type DogStatsdSink_Address struct {
	Address *v31.Address `protobuf:"bytes,1,opt,name=address,proto3,oneof"`
}

func (*DogStatsdSink_Address) isDogStatsdSink_DogStatsdSpecifier() {}

func (m *DogStatsdSink) GetDogStatsdSpecifier() isDogStatsdSink_DogStatsdSpecifier {
	if m != nil {
		return m.DogStatsdSpecifier
	}
	return nil
}

func (m *DogStatsdSink) GetAddress() *v31.Address {
	if x, ok := m.GetDogStatsdSpecifier().(*DogStatsdSink_Address); ok {
		return x.Address
	}
	return nil
}

func (m *DogStatsdSink) GetPrefix() string {
	if m != nil {
		return m.Prefix
	}
	return ""
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*DogStatsdSink) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*DogStatsdSink_Address)(nil),
	}
}

type HystrixSink struct {
	NumBuckets           int64    `protobuf:"varint,1,opt,name=num_buckets,json=numBuckets,proto3" json:"num_buckets,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HystrixSink) Reset()         { *m = HystrixSink{} }
func (m *HystrixSink) String() string { return proto.CompactTextString(m) }
func (*HystrixSink) ProtoMessage()    {}
func (*HystrixSink) Descriptor() ([]byte, []int) {
	return fileDescriptor_38edc339428e9971, []int{6}
}

func (m *HystrixSink) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HystrixSink.Unmarshal(m, b)
}
func (m *HystrixSink) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HystrixSink.Marshal(b, m, deterministic)
}
func (m *HystrixSink) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HystrixSink.Merge(m, src)
}
func (m *HystrixSink) XXX_Size() int {
	return xxx_messageInfo_HystrixSink.Size(m)
}
func (m *HystrixSink) XXX_DiscardUnknown() {
	xxx_messageInfo_HystrixSink.DiscardUnknown(m)
}

var xxx_messageInfo_HystrixSink proto.InternalMessageInfo

func (m *HystrixSink) GetNumBuckets() int64 {
	if m != nil {
		return m.NumBuckets
	}
	return 0
}

func init() {
	proto.RegisterType((*StatsSink)(nil), "envoy.config.metrics.v3.StatsSink")
	proto.RegisterType((*StatsConfig)(nil), "envoy.config.metrics.v3.StatsConfig")
	proto.RegisterType((*StatsMatcher)(nil), "envoy.config.metrics.v3.StatsMatcher")
	proto.RegisterType((*TagSpecifier)(nil), "envoy.config.metrics.v3.TagSpecifier")
	proto.RegisterType((*StatsdSink)(nil), "envoy.config.metrics.v3.StatsdSink")
	proto.RegisterType((*DogStatsdSink)(nil), "envoy.config.metrics.v3.DogStatsdSink")
	proto.RegisterType((*HystrixSink)(nil), "envoy.config.metrics.v3.HystrixSink")
}

func init() {
	proto.RegisterFile("envoy/config/metrics/v3/stats.proto", fileDescriptor_38edc339428e9971)
}

var fileDescriptor_38edc339428e9971 = []byte{
	// 834 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x55, 0x4d, 0x6f, 0xe2, 0x56,
	0x14, 0xc5, 0x30, 0x93, 0x81, 0x6b, 0x18, 0x51, 0x0b, 0x15, 0x66, 0xa6, 0x93, 0x21, 0xcc, 0xd0,
	0xd2, 0x69, 0x65, 0x4b, 0x64, 0x15, 0xa4, 0x2e, 0x70, 0xa8, 0x84, 0xd2, 0x0f, 0xa5, 0x26, 0xea,
	0xa2, 0x1b, 0xeb, 0x61, 0x3f, 0x9c, 0xd7, 0x18, 0x3f, 0xeb, 0xf9, 0x99, 0xc2, 0x2e, 0xcb, 0xfe,
	0x84, 0xaa, 0xbf, 0xa0, 0xfb, 0xee, 0xba, 0xaf, 0xd4, 0x5d, 0xd5, 0x5f, 0xd3, 0xaa, 0xab, 0xea,
	0x7d, 0x40, 0x21, 0x09, 0x89, 0x2a, 0x75, 0x87, 0xdf, 0xb9, 0xf7, 0xf8, 0x9c, 0xe3, 0x7b, 0x1f,
	0xf0, 0x1a, 0x27, 0x0b, 0xba, 0x72, 0x02, 0x9a, 0xcc, 0x48, 0xe4, 0xcc, 0x31, 0x67, 0x24, 0xc8,
	0x9c, 0xc5, 0xb1, 0x93, 0x71, 0xc4, 0x33, 0x3b, 0x65, 0x94, 0x53, 0xab, 0x29, 0x8b, 0x6c, 0x55,
	0x64, 0xeb, 0x22, 0x7b, 0x71, 0xfc, 0xbc, 0xb3, 0xd3, 0x1d, 0x50, 0x86, 0x45, 0x2b, 0x0a, 0x43,
	0x86, 0x33, 0xdd, 0xbc, 0xae, 0xe1, 0xab, 0x14, 0x3b, 0x73, 0xc4, 0x83, 0x4b, 0xcc, 0x14, 0x3f,
	0x23, 0x49, 0xa4, 0x6b, 0x9e, 0x45, 0x94, 0x46, 0x31, 0x76, 0xe4, 0xd3, 0x34, 0x9f, 0x39, 0x28,
	0x59, 0x69, 0xe8, 0xbd, 0x9b, 0x50, 0xc6, 0x59, 0x1e, 0x70, 0x8d, 0x1e, 0xde, 0x44, 0xbf, 0x63,
	0x28, 0x4d, 0x31, 0x5b, 0xbf, 0xfc, 0x65, 0x1e, 0xa6, 0xc8, 0x41, 0x49, 0x42, 0x39, 0xe2, 0x84,
	0x26, 0x99, 0xf4, 0x95, 0xaf, 0xe1, 0xa3, 0x5b, 0xf0, 0x02, 0xb3, 0x8c, 0xd0, 0xe4, 0x5f, 0x69,
	0xcd, 0x05, 0x8a, 0x49, 0x88, 0x38, 0x76, 0xd6, 0x3f, 0x14, 0xd0, 0xf9, 0xd3, 0x80, 0xca, 0x44,
	0x84, 0x34, 0x21, 0xc9, 0x95, 0x65, 0xc1, 0xa3, 0x04, 0xcd, 0x71, 0xcb, 0x68, 0x1b, 0xbd, 0x8a,
	0x27, 0x7f, 0x5b, 0x53, 0x38, 0xbc, 0x24, 0x61, 0x88, 0x13, 0x5f, 0x46, 0xe0, 0x87, 0x38, 0x65,
	0x38, 0x40, 0x1c, 0x87, 0xbe, 0x4a, 0xac, 0x55, 0x6c, 0x1b, 0x3d, 0xb3, 0xdf, 0xb4, 0x95, 0x0b,
	0x7b, 0xed, 0xc2, 0x9e, 0x48, 0x8f, 0x6e, 0xb1, 0x65, 0x8c, 0x0b, 0xde, 0x0b, 0x45, 0xf2, 0xa9,
	0xe0, 0x18, 0x6d, 0x28, 0x4e, 0x25, 0x83, 0x75, 0x02, 0x55, 0x91, 0xec, 0x86, 0xb1, 0x24, 0x19,
	0x1b, 0xb7, 0x18, 0x87, 0xc9, 0x6a, 0x5c, 0xf0, 0x4c, 0x59, 0xab, 0x5a, 0x07, 0xbd, 0x1f, 0x7f,
	0xfd, 0xfe, 0xf0, 0x35, 0x1c, 0xdd, 0xfd, 0x71, 0xfb, 0xf6, 0xc6, 0x9c, 0x5b, 0x03, 0x53, 0xc1,
	0xbe, 0xe8, 0xef, 0xfc, 0x50, 0x04, 0x53, 0x82, 0x5a, 0xc3, 0x08, 0x40, 0x4e, 0x8b, 0xcf, 0x51,
	0x94, 0xb5, 0x8c, 0x76, 0xa9, 0x67, 0xf6, 0xbb, 0xf6, 0x9e, 0x99, 0xb1, 0x2f, 0x50, 0x34, 0x49,
	0x71, 0x40, 0x66, 0x04, 0x33, 0xaf, 0x22, 0x1b, 0x2f, 0x50, 0x94, 0x59, 0x9f, 0x41, 0x23, 0xcf,
	0xb0, 0x8f, 0xe2, 0xd8, 0x0f, 0xf1, 0x0c, 0xe5, 0x31, 0x57, 0x7c, 0x2a, 0xa3, 0xe7, 0xb7, 0x1c,
	0xb9, 0x94, 0xc6, 0x5f, 0xa3, 0x38, 0xc7, 0xde, 0x3b, 0x79, 0x86, 0x87, 0x71, 0x3c, 0x52, 0x5d,
	0x92, 0xec, 0x0c, 0x6a, 0x4a, 0x92, 0x9e, 0x38, 0x9d, 0xcb, 0x7e, 0x55, 0xd2, 0xcf, 0x17, 0xaa,
	0xd8, 0xab, 0x66, 0x5b, 0x4f, 0x83, 0xb7, 0x22, 0xa7, 0xae, 0xde, 0x94, 0x7d, 0x39, 0xa9, 0x28,
	0x44, 0x34, 0xd5, 0x6d, 0x2a, 0xeb, 0x15, 0x00, 0xc3, 0xdf, 0xe2, 0x80, 0x0b, 0x63, 0x72, 0x3a,
	0xca, 0xe3, 0x82, 0x57, 0x51, 0x67, 0xc3, 0x38, 0xb6, 0xbe, 0x82, 0xa7, 0x78, 0x19, 0xc4, 0xb9,
	0x98, 0x3a, 0x3f, 0x26, 0x19, 0xd7, 0x86, 0x7b, 0x5a, 0xaa, 0x48, 0xdc, 0xd6, 0x2e, 0x84, 0xd0,
	0xcf, 0x49, 0xc6, 0x27, 0x72, 0x77, 0xf4, 0x2b, 0xc6, 0x05, 0xaf, 0xb6, 0x61, 0x10, 0xa8, 0xa0,
	0x24, 0xc9, 0x0e, 0x65, 0xe9, 0xbf, 0x53, 0x6e, 0x18, 0x04, 0x3a, 0xf8, 0x48, 0x64, 0xf0, 0x3e,
	0xbc, 0xb9, 0x37, 0x03, 0xdd, 0xed, 0x36, 0x6e, 0x84, 0x6f, 0x95, 0xfe, 0x72, 0x8d, 0xce, 0x4f,
	0x06, 0x54, 0xb7, 0xbf, 0xbd, 0xf5, 0x0c, 0xca, 0x1c, 0x45, 0xfe, 0xd6, 0xda, 0x3c, 0xe1, 0x28,
	0xfa, 0x52, 0x6c, 0x4e, 0x1b, 0x1e, 0x33, 0x1c, 0xe1, 0xa5, 0xcc, 0xa2, 0xe2, 0x96, 0xff, 0x76,
	0x1f, 0xb3, 0x52, 0xef, 0x5a, 0x44, 0xa7, 0x00, 0xeb, 0x08, 0xcc, 0x19, 0x59, 0xe2, 0xd0, 0x5f,
	0x88, 0x11, 0x90, 0x06, 0x2b, 0xe3, 0x82, 0x07, 0xf2, 0x50, 0x8e, 0xc5, 0x83, 0x9a, 0xb7, 0xc5,
	0xb8, 0x26, 0x54, 0x84, 0x18, 0xc9, 0xd6, 0xf9, 0xdd, 0x00, 0x90, 0x8e, 0x42, 0xb9, 0xdb, 0x27,
	0xf0, 0x44, 0x5f, 0x69, 0x52, 0xa7, 0xd9, 0x7f, 0xb9, 0x3b, 0x46, 0xe2, 0xde, 0x13, 0x39, 0x0e,
	0x55, 0xd1, 0xb8, 0xe0, 0xad, 0xeb, 0xad, 0xb7, 0x50, 0xe7, 0x41, 0xea, 0x8b, 0x28, 0x39, 0x66,
	0xca, 0x6b, 0x51, 0x6b, 0x7d, 0xca, 0x83, 0xf4, 0x54, 0x01, 0xd2, 0xf4, 0xbb, 0x70, 0x90, 0x32,
	0x3c, 0x23, 0x4b, 0xe5, 0xc6, 0xd3, 0x4f, 0x83, 0x0f, 0x85, 0x8f, 0x37, 0xd0, 0xb9, 0x37, 0x7b,
	0xa9, 0xd4, 0x6d, 0x42, 0x5d, 0x26, 0x1f, 0xfa, 0xd9, 0x26, 0x66, 0x19, 0xfe, 0xcf, 0x06, 0xd4,
	0x46, 0x34, 0xfa, 0x7f, 0x4c, 0xed, 0x13, 0xfa, 0xb1, 0x10, 0xfa, 0x01, 0x74, 0xf7, 0x09, 0xdd,
	0x11, 0xe0, 0xbe, 0x80, 0x46, 0x48, 0x23, 0xff, 0x4e, 0xbd, 0x67, 0x8f, 0xca, 0xc5, 0x7a, 0xa9,
	0xf3, 0x0d, 0x98, 0xe3, 0x95, 0xf8, 0xa3, 0x58, 0x4a, 0xc9, 0xaf, 0xc0, 0x4c, 0xf2, 0xb9, 0x3f,
	0xcd, 0x83, 0x2b, 0xcc, 0x95, 0xec, 0x92, 0x07, 0x49, 0x3e, 0x77, 0xd5, 0xc9, 0x83, 0x9b, 0xba,
	0x45, 0xe6, 0x7e, 0xf2, 0xcb, 0xf5, 0x6f, 0x7f, 0x1c, 0x14, 0xeb, 0x45, 0xe8, 0x12, 0xaa, 0xac,
	0xa7, 0x8c, 0x2e, 0x57, 0xfb, 0x6e, 0x08, 0x57, 0x4d, 0xc4, 0xb9, 0xb8, 0x7e, 0xce, 0x8d, 0xe9,
	0x81, 0xbc, 0x87, 0x8e, 0xff, 0x09, 0x00, 0x00, 0xff, 0xff, 0x29, 0xc5, 0x9a, 0xd9, 0x40, 0x07,
	0x00, 0x00,
}
