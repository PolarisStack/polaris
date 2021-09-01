// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/config/filter/http/dynamic_forward_proxy/v2alpha/dynamic_forward_proxy.proto

package envoy_config_filter_http_dynamic_forward_proxy_v2alpha

import (
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	v2alpha "github.com/polarismesh/polaris-server/apiserver/xdsserver/go-control-plane/envoy/config/common/dynamic_forward_proxy/v2alpha"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
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

type FilterConfig struct {
	DnsCacheConfig       *v2alpha.DnsCacheConfig `protobuf:"bytes,1,opt,name=dns_cache_config,json=dnsCacheConfig,proto3" json:"dns_cache_config,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *FilterConfig) Reset()         { *m = FilterConfig{} }
func (m *FilterConfig) String() string { return proto.CompactTextString(m) }
func (*FilterConfig) ProtoMessage()    {}
func (*FilterConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_85a2356b260c47da, []int{0}
}

func (m *FilterConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FilterConfig.Unmarshal(m, b)
}
func (m *FilterConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FilterConfig.Marshal(b, m, deterministic)
}
func (m *FilterConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FilterConfig.Merge(m, src)
}
func (m *FilterConfig) XXX_Size() int {
	return xxx_messageInfo_FilterConfig.Size(m)
}
func (m *FilterConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_FilterConfig.DiscardUnknown(m)
}

var xxx_messageInfo_FilterConfig proto.InternalMessageInfo

func (m *FilterConfig) GetDnsCacheConfig() *v2alpha.DnsCacheConfig {
	if m != nil {
		return m.DnsCacheConfig
	}
	return nil
}

type PerRouteConfig struct {
	// Types that are valid to be assigned to HostRewriteSpecifier:
	//	*PerRouteConfig_HostRewrite
	//	*PerRouteConfig_AutoHostRewriteHeader
	HostRewriteSpecifier isPerRouteConfig_HostRewriteSpecifier `protobuf_oneof:"host_rewrite_specifier"`
	XXX_NoUnkeyedLiteral struct{}                              `json:"-"`
	XXX_unrecognized     []byte                                `json:"-"`
	XXX_sizecache        int32                                 `json:"-"`
}

func (m *PerRouteConfig) Reset()         { *m = PerRouteConfig{} }
func (m *PerRouteConfig) String() string { return proto.CompactTextString(m) }
func (*PerRouteConfig) ProtoMessage()    {}
func (*PerRouteConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_85a2356b260c47da, []int{1}
}

func (m *PerRouteConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PerRouteConfig.Unmarshal(m, b)
}
func (m *PerRouteConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PerRouteConfig.Marshal(b, m, deterministic)
}
func (m *PerRouteConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PerRouteConfig.Merge(m, src)
}
func (m *PerRouteConfig) XXX_Size() int {
	return xxx_messageInfo_PerRouteConfig.Size(m)
}
func (m *PerRouteConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_PerRouteConfig.DiscardUnknown(m)
}

var xxx_messageInfo_PerRouteConfig proto.InternalMessageInfo

type isPerRouteConfig_HostRewriteSpecifier interface {
	isPerRouteConfig_HostRewriteSpecifier()
}

type PerRouteConfig_HostRewrite struct {
	HostRewrite string `protobuf:"bytes,1,opt,name=host_rewrite,json=hostRewrite,proto3,oneof"`
}

type PerRouteConfig_AutoHostRewriteHeader struct {
	AutoHostRewriteHeader string `protobuf:"bytes,2,opt,name=auto_host_rewrite_header,json=autoHostRewriteHeader,proto3,oneof"`
}

func (*PerRouteConfig_HostRewrite) isPerRouteConfig_HostRewriteSpecifier() {}

func (*PerRouteConfig_AutoHostRewriteHeader) isPerRouteConfig_HostRewriteSpecifier() {}

func (m *PerRouteConfig) GetHostRewriteSpecifier() isPerRouteConfig_HostRewriteSpecifier {
	if m != nil {
		return m.HostRewriteSpecifier
	}
	return nil
}

func (m *PerRouteConfig) GetHostRewrite() string {
	if x, ok := m.GetHostRewriteSpecifier().(*PerRouteConfig_HostRewrite); ok {
		return x.HostRewrite
	}
	return ""
}

func (m *PerRouteConfig) GetAutoHostRewriteHeader() string {
	if x, ok := m.GetHostRewriteSpecifier().(*PerRouteConfig_AutoHostRewriteHeader); ok {
		return x.AutoHostRewriteHeader
	}
	return ""
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*PerRouteConfig) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*PerRouteConfig_HostRewrite)(nil),
		(*PerRouteConfig_AutoHostRewriteHeader)(nil),
	}
}

func init() {
	proto.RegisterType((*FilterConfig)(nil), "envoy.config.filter.http.dynamic_forward_proxy.v2alpha.FilterConfig")
	proto.RegisterType((*PerRouteConfig)(nil), "envoy.config.filter.http.dynamic_forward_proxy.v2alpha.PerRouteConfig")
}

func init() {
	proto.RegisterFile("envoy/config/filter/http/dynamic_forward_proxy/v2alpha/dynamic_forward_proxy.proto", fileDescriptor_85a2356b260c47da)
}

var fileDescriptor_85a2356b260c47da = []byte{
	// 410 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x92, 0xbd, 0xae, 0xd3, 0x30,
	0x14, 0xc7, 0xaf, 0x2b, 0xdd, 0x0a, 0xdc, 0xaa, 0xaa, 0x02, 0x2d, 0x51, 0xf9, 0x10, 0xea, 0xc4,
	0x14, 0x4b, 0xad, 0x84, 0x58, 0x9b, 0x56, 0x55, 0xc6, 0x28, 0x03, 0x6b, 0x64, 0x12, 0xa7, 0xb1,
	0x94, 0xd8, 0x91, 0x7d, 0xd2, 0x8f, 0x8d, 0x85, 0x85, 0x85, 0x95, 0x47, 0xe0, 0x19, 0xd8, 0x91,
	0x58, 0x79, 0x15, 0x46, 0x06, 0x84, 0x62, 0xa7, 0x40, 0x44, 0x55, 0xa4, 0xbb, 0x45, 0xe7, 0x77,
	0xfc, 0xf3, 0x39, 0xf1, 0x1f, 0x47, 0x4c, 0xec, 0xe5, 0x89, 0x24, 0x52, 0x64, 0x7c, 0x47, 0x32,
	0x5e, 0x00, 0x53, 0x24, 0x07, 0xa8, 0x48, 0x7a, 0x12, 0xb4, 0xe4, 0x49, 0x9c, 0x49, 0x75, 0xa0,
	0x2a, 0x8d, 0x2b, 0x25, 0x8f, 0x27, 0xb2, 0x5f, 0xd0, 0xa2, 0xca, 0xe9, 0x65, 0xea, 0x55, 0x4a,
	0x82, 0x74, 0x5e, 0x1a, 0xa7, 0x67, 0x9d, 0x9e, 0x75, 0x7a, 0x8d, 0xd3, 0xbb, 0x7c, 0xaa, 0x75,
	0xce, 0x56, 0x9d, 0x59, 0x12, 0x59, 0x96, 0x52, 0xfc, 0x6f, 0x0c, 0xa1, 0xe3, 0x84, 0x26, 0x39,
	0xb3, 0x57, 0xcf, 0x9e, 0xd5, 0x69, 0x45, 0x09, 0x15, 0x42, 0x02, 0x05, 0x2e, 0x85, 0x26, 0x25,
	0xdf, 0x29, 0x0a, 0x67, 0xfe, 0xf4, 0x1f, 0xae, 0x81, 0x42, 0xad, 0x5b, 0xfc, 0x68, 0x4f, 0x0b,
	0x9e, 0x52, 0x60, 0xe4, 0xfc, 0x61, 0xc1, 0xfc, 0x1d, 0xc2, 0xc3, 0xad, 0x59, 0x64, 0x6d, 0xa6,
	0x73, 0x6a, 0x3c, 0xfe, 0x7d, 0x77, 0x6c, 0x27, 0x76, 0xd1, 0x73, 0xf4, 0x62, 0xb0, 0x58, 0x79,
	0x9d, 0xf5, 0xed, 0x1a, 0xd7, 0x37, 0xf7, 0x36, 0x42, 0xaf, 0x1b, 0x93, 0x95, 0xfb, 0xf7, 0x7e,
	0xf8, 0xb7, 0xef, 0x51, 0x6f, 0x8c, 0xa2, 0x51, 0xda, 0x21, 0xf3, 0x2f, 0x08, 0x8f, 0x42, 0xa6,
	0x22, 0x59, 0x43, 0x5b, 0x72, 0x56, 0x78, 0x98, 0x4b, 0x0d, 0xb1, 0x62, 0x07, 0xc5, 0x81, 0x99,
	0x29, 0xee, 0xfb, 0x4f, 0xbe, 0x7f, 0xfc, 0xf9, 0xe1, 0x76, 0x8a, 0x1f, 0xfe, 0xcd, 0xe2, 0x82,
	0x03, 0x53, 0xb4, 0x08, 0x6e, 0xa2, 0x41, 0x53, 0x8f, 0x6c, 0xd9, 0x79, 0x8d, 0x5d, 0x5a, 0x83,
	0x8c, 0x3b, 0xbd, 0x39, 0xa3, 0x29, 0x53, 0x6e, 0xcf, 0xe8, 0x1e, 0x1b, 0xdd, 0x04, 0x3f, 0xb8,
	0xd0, 0x12, 0xdc, 0x44, 0x93, 0xe6, 0x78, 0xf0, 0xc7, 0x18, 0x18, 0xe0, 0xbb, 0x78, 0xda, 0xe9,
	0xd7, 0x15, 0x4b, 0x78, 0xc6, 0x99, 0xf2, 0x3f, 0x21, 0x23, 0x7c, 0x75, 0xce, 0x0a, 0x3b, 0x02,
	0x13, 0xba, 0x79, 0x90, 0x36, 0x2f, 0xfa, 0x6a, 0x60, 0x96, 0x9f, 0xdf, 0x7e, 0xfd, 0xd6, 0xef,
	0x8d, 0x11, 0xde, 0x70, 0x69, 0xff, 0xb7, 0x25, 0x77, 0x4b, 0x9e, 0xef, 0x6e, 0x2c, 0xde, 0x5a,
	0x1a, 0x36, 0x30, 0x6c, 0x1e, 0x3e, 0x44, 0x6f, 0xfa, 0x26, 0x01, 0xcb, 0x5f, 0x01, 0x00, 0x00,
	0xff, 0xff, 0x67, 0x0c, 0x51, 0x47, 0x2a, 0x03, 0x00, 0x00,
}
