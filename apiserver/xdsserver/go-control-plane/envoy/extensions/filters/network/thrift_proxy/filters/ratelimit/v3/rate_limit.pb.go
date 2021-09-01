// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/extensions/filters/network/thrift_proxy/filters/ratelimit/v3/rate_limit.proto

package envoy_extensions_filters_network_thrift_proxy_filters_ratelimit_v3

import (
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	v3 "github.com/polarismesh/polaris-server/apiserver/xdsserver/go-control-plane/envoy/config/ratelimit/v3"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	duration "github.com/golang/protobuf/ptypes/duration"
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

type RateLimit struct {
	Domain               string                     `protobuf:"bytes,1,opt,name=domain,proto3" json:"domain,omitempty"`
	Stage                uint32                     `protobuf:"varint,2,opt,name=stage,proto3" json:"stage,omitempty"`
	Timeout              *duration.Duration         `protobuf:"bytes,3,opt,name=timeout,proto3" json:"timeout,omitempty"`
	FailureModeDeny      bool                       `protobuf:"varint,4,opt,name=failure_mode_deny,json=failureModeDeny,proto3" json:"failure_mode_deny,omitempty"`
	RateLimitService     *v3.RateLimitServiceConfig `protobuf:"bytes,5,opt,name=rate_limit_service,json=rateLimitService,proto3" json:"rate_limit_service,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *RateLimit) Reset()         { *m = RateLimit{} }
func (m *RateLimit) String() string { return proto.CompactTextString(m) }
func (*RateLimit) ProtoMessage()    {}
func (*RateLimit) Descriptor() ([]byte, []int) {
	return fileDescriptor_920e335ffb05829c, []int{0}
}

func (m *RateLimit) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RateLimit.Unmarshal(m, b)
}
func (m *RateLimit) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RateLimit.Marshal(b, m, deterministic)
}
func (m *RateLimit) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RateLimit.Merge(m, src)
}
func (m *RateLimit) XXX_Size() int {
	return xxx_messageInfo_RateLimit.Size(m)
}
func (m *RateLimit) XXX_DiscardUnknown() {
	xxx_messageInfo_RateLimit.DiscardUnknown(m)
}

var xxx_messageInfo_RateLimit proto.InternalMessageInfo

func (m *RateLimit) GetDomain() string {
	if m != nil {
		return m.Domain
	}
	return ""
}

func (m *RateLimit) GetStage() uint32 {
	if m != nil {
		return m.Stage
	}
	return 0
}

func (m *RateLimit) GetTimeout() *duration.Duration {
	if m != nil {
		return m.Timeout
	}
	return nil
}

func (m *RateLimit) GetFailureModeDeny() bool {
	if m != nil {
		return m.FailureModeDeny
	}
	return false
}

func (m *RateLimit) GetRateLimitService() *v3.RateLimitServiceConfig {
	if m != nil {
		return m.RateLimitService
	}
	return nil
}

func init() {
	proto.RegisterType((*RateLimit)(nil), "envoy.extensions.filters.network.thrift_proxy.filters.ratelimit.v3.RateLimit")
}

func init() {
	proto.RegisterFile("envoy/extensions/filters/network/thrift_proxy/filters/ratelimit/v3/rate_limit.proto", fileDescriptor_920e335ffb05829c)
}

var fileDescriptor_920e335ffb05829c = []byte{
	// 436 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x52, 0xb1, 0x6e, 0xd4, 0x40,
	0x10, 0xd5, 0x9a, 0xdc, 0x25, 0x59, 0x04, 0x1c, 0x6e, 0x30, 0x91, 0x12, 0x0c, 0x34, 0xa7, 0x14,
	0xbb, 0x4a, 0xae, 0x41, 0x69, 0x90, 0x4c, 0x4a, 0x90, 0x4e, 0xce, 0x07, 0x58, 0x9b, 0x78, 0xec,
	0xac, 0xf0, 0xed, 0x9c, 0xd6, 0x63, 0x93, 0xeb, 0x28, 0x11, 0x9f, 0xc0, 0xa7, 0xd0, 0x23, 0xd1,
	0xd2, 0xf1, 0x2d, 0xa9, 0xd0, 0xed, 0xda, 0xbe, 0xe4, 0x28, 0xd3, 0xed, 0xce, 0x9b, 0x79, 0x6f,
	0x66, 0xde, 0xf0, 0x0b, 0x30, 0x2d, 0xae, 0x24, 0xdc, 0x10, 0x98, 0x5a, 0xa3, 0xa9, 0x65, 0xa1,
	0x2b, 0x02, 0x5b, 0x4b, 0x03, 0xf4, 0x05, 0xed, 0x67, 0x49, 0xd7, 0x56, 0x17, 0x94, 0x2d, 0x2d,
	0xde, 0xac, 0x06, 0xd0, 0x2a, 0x82, 0x4a, 0x2f, 0x34, 0xc9, 0x76, 0xe6, 0x3e, 0x99, 0xfb, 0x89,
	0xa5, 0x45, 0xc2, 0x30, 0x71, 0xa4, 0x62, 0x43, 0x2a, 0xba, 0x3a, 0xd1, 0x91, 0x8a, 0xbb, 0xa4,
	0x03, 0x38, 0x90, 0x8a, 0x76, 0x76, 0xf0, 0xd6, 0x37, 0x76, 0x85, 0xa6, 0xd0, 0xe5, 0x96, 0x5e,
	0x55, 0x7b, 0xa1, 0x83, 0xa3, 0x12, 0xb1, 0xac, 0x40, 0xba, 0xdf, 0x65, 0x53, 0xc8, 0xbc, 0xb1,
	0x8a, 0x34, 0x9a, 0x0e, 0x3f, 0x6c, 0xf2, 0xa5, 0x92, 0xca, 0x18, 0x24, 0x17, 0xae, 0x65, 0x4d,
	0x8a, 0x9a, 0xbe, 0xfc, 0xf5, 0x7f, 0x70, 0x0b, 0x76, 0xdd, 0xb0, 0x36, 0x65, 0x97, 0xf2, 0xa2,
	0x55, 0x95, 0xce, 0x15, 0x81, 0xec, 0x1f, 0x1e, 0x78, 0xf3, 0x37, 0xe0, 0xfb, 0xa9, 0x22, 0xf8,
	0xb8, 0xee, 0x2a, 0x7c, 0xc5, 0xc7, 0x39, 0x2e, 0x94, 0x36, 0x11, 0x8b, 0xd9, 0x74, 0x3f, 0xd9,
	0xbd, 0x4d, 0x76, 0x6c, 0x10, 0xb3, 0xb4, 0x0b, 0x87, 0x87, 0x7c, 0x54, 0x93, 0x2a, 0x21, 0x0a,
	0x62, 0x36, 0x7d, 0xe2, 0xf0, 0xe3, 0x20, 0xe2, 0xa9, 0x8f, 0x86, 0x33, 0xbe, 0x4b, 0x7a, 0x01,
	0xd8, 0x50, 0xf4, 0x28, 0x66, 0xd3, 0xc7, 0xa7, 0x2f, 0x85, 0x1f, 0x4d, 0xf4, 0xa3, 0x89, 0xf3,
	0x6e, 0xb4, 0xb4, 0xcf, 0x0c, 0x8f, 0xf9, 0xf3, 0x42, 0xe9, 0xaa, 0xb1, 0x90, 0x2d, 0x30, 0x87,
	0x2c, 0x07, 0xb3, 0x8a, 0x76, 0x62, 0x36, 0xdd, 0x4b, 0x9f, 0x75, 0xc0, 0x27, 0xcc, 0xe1, 0x1c,
	0xcc, 0x2a, 0xd4, 0x3c, 0xdc, 0xd8, 0x94, 0xd5, 0x60, 0x5b, 0x7d, 0x05, 0xd1, 0xc8, 0x69, 0x9d,
	0x08, 0xef, 0x97, 0xdf, 0xf5, 0x3d, 0x1b, 0xc4, 0x30, 0xe2, 0x85, 0x2f, 0xf9, 0xe0, 0x72, 0x92,
	0xbd, 0xdb, 0x64, 0xf4, 0x9d, 0x05, 0x13, 0x96, 0x4e, 0xec, 0x56, 0xc6, 0xd9, 0xfb, 0x1f, 0xbf,
	0xbe, 0x1d, 0x9d, 0xf1, 0x77, 0xf7, 0x48, 0xbd, 0xc7, 0x9d, 0xef, 0xe2, 0xce, 0xc5, 0xb4, 0xa7,
	0xaa, 0x5a, 0x5e, 0xab, 0x93, 0x8d, 0x52, 0x52, 0xfd, 0xfc, 0xfa, 0xfb, 0xcf, 0x38, 0x98, 0x04,
	0x7c, 0xae, 0xd1, 0xf7, 0xe6, 0x0f, 0xe5, 0xe1, 0x67, 0x95, 0x3c, 0x1d, 0x64, 0xe6, 0xeb, 0xc5,
	0xce, 0xd9, 0xe5, 0xd8, 0x6d, 0x78, 0xf6, 0x2f, 0x00, 0x00, 0xff, 0xff, 0xf3, 0xb1, 0x3f, 0xe9,
	0x0a, 0x03, 0x00, 0x00,
}
