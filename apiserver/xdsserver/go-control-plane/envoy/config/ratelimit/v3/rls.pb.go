// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/config/ratelimit/v3/rls.proto

package envoy_config_ratelimit_v3

import (
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	v3 "github.com/polarismesh/polaris-server/apiserver/xdsserver/go-control-plane/envoy/config/core/v3"
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

type RateLimitServiceConfig struct {
	GrpcService          *v3.GrpcService `protobuf:"bytes,2,opt,name=grpc_service,json=grpcService,proto3" json:"grpc_service,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *RateLimitServiceConfig) Reset()         { *m = RateLimitServiceConfig{} }
func (m *RateLimitServiceConfig) String() string { return proto.CompactTextString(m) }
func (*RateLimitServiceConfig) ProtoMessage()    {}
func (*RateLimitServiceConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_ca555cbfdc4e6f71, []int{0}
}

func (m *RateLimitServiceConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RateLimitServiceConfig.Unmarshal(m, b)
}
func (m *RateLimitServiceConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RateLimitServiceConfig.Marshal(b, m, deterministic)
}
func (m *RateLimitServiceConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RateLimitServiceConfig.Merge(m, src)
}
func (m *RateLimitServiceConfig) XXX_Size() int {
	return xxx_messageInfo_RateLimitServiceConfig.Size(m)
}
func (m *RateLimitServiceConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_RateLimitServiceConfig.DiscardUnknown(m)
}

var xxx_messageInfo_RateLimitServiceConfig proto.InternalMessageInfo

func (m *RateLimitServiceConfig) GetGrpcService() *v3.GrpcService {
	if m != nil {
		return m.GrpcService
	}
	return nil
}

func init() {
	proto.RegisterType((*RateLimitServiceConfig)(nil), "envoy.config.ratelimit.v3.RateLimitServiceConfig")
}

func init() {
	proto.RegisterFile("envoy/config/ratelimit/v3/rls.proto", fileDescriptor_ca555cbfdc4e6f71)
}

var fileDescriptor_ca555cbfdc4e6f71 = []byte{
	// 287 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x4e, 0xcd, 0x2b, 0xcb,
	0xaf, 0xd4, 0x4f, 0xce, 0xcf, 0x4b, 0xcb, 0x4c, 0xd7, 0x2f, 0x4a, 0x2c, 0x49, 0xcd, 0xc9, 0xcc,
	0xcd, 0x2c, 0xd1, 0x2f, 0x33, 0xd6, 0x2f, 0xca, 0x29, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17,
	0x92, 0x04, 0x2b, 0xd2, 0x83, 0x28, 0xd2, 0x83, 0x2b, 0xd2, 0x2b, 0x33, 0x96, 0x52, 0x47, 0xd1,
	0x9f, 0x9c, 0x5f, 0x94, 0x0a, 0xd2, 0x9a, 0x5e, 0x54, 0x90, 0x1c, 0x5f, 0x9c, 0x5a, 0x54, 0x96,
	0x99, 0x9c, 0x0a, 0x31, 0x43, 0x4a, 0xb6, 0x34, 0xa5, 0x20, 0x51, 0x3f, 0x31, 0x2f, 0x2f, 0xbf,
	0x24, 0xb1, 0x24, 0x33, 0x3f, 0xaf, 0x58, 0xbf, 0xb8, 0x24, 0xb1, 0xa4, 0x14, 0x6a, 0x85, 0x94,
	0x22, 0x86, 0x74, 0x59, 0x6a, 0x51, 0x71, 0x66, 0x7e, 0x5e, 0x66, 0x5e, 0x3a, 0x54, 0x89, 0x78,
	0x59, 0x62, 0x4e, 0x66, 0x4a, 0x62, 0x49, 0xaa, 0x3e, 0x8c, 0x01, 0x91, 0x50, 0x5a, 0xcb, 0xc8,
	0x25, 0x16, 0x94, 0x58, 0x92, 0xea, 0x03, 0x72, 0x54, 0x30, 0xc4, 0x56, 0x67, 0xb0, 0x83, 0x84,
	0xfc, 0xb8, 0x78, 0x90, 0xdd, 0x22, 0xc1, 0xa4, 0xc0, 0xa8, 0xc1, 0x6d, 0xa4, 0xa8, 0x87, 0xe2,
	0x21, 0x90, 0xab, 0xf5, 0xca, 0x8c, 0xf5, 0xdc, 0x8b, 0x0a, 0x92, 0xa1, 0xda, 0x9d, 0x38, 0x7e,
	0x39, 0xb1, 0x76, 0x31, 0x32, 0x09, 0x30, 0x06, 0x71, 0xa7, 0x23, 0x84, 0xad, 0xcc, 0x67, 0x1d,
	0xed, 0x90, 0x33, 0xe2, 0x32, 0xc0, 0x15, 0x20, 0x46, 0x7a, 0xd8, 0x1d, 0xe2, 0xc5, 0xc2, 0xc1,
	0x28, 0xc0, 0xe4, 0xc5, 0xc2, 0xc1, 0x2c, 0xc0, 0xe2, 0x64, 0xbb, 0xab, 0xe1, 0xc4, 0x45, 0x36,
	0x26, 0x01, 0x26, 0x2e, 0xf5, 0xcc, 0x7c, 0x88, 0x53, 0x0a, 0x8a, 0xf2, 0x2b, 0x2a, 0xf5, 0x70,
	0x06, 0xb3, 0x13, 0x47, 0x50, 0x4e, 0x71, 0x00, 0xc8, 0xb3, 0x01, 0x8c, 0x49, 0x6c, 0x60, 0x5f,
	0x1b, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x9c, 0x8c, 0x01, 0x9e, 0xbb, 0x01, 0x00, 0x00,
}
