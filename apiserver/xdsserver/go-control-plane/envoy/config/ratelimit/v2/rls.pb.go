// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/config/ratelimit/v2/rls.proto

package envoy_config_ratelimit_v2

import (
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	core "github.com/polarismesh/polaris-server/apiserver/xdsserver/go-control-plane/envoy/api/v2/core"
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
	GrpcService          *core.GrpcService `protobuf:"bytes,2,opt,name=grpc_service,json=grpcService,proto3" json:"grpc_service,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *RateLimitServiceConfig) Reset()         { *m = RateLimitServiceConfig{} }
func (m *RateLimitServiceConfig) String() string { return proto.CompactTextString(m) }
func (*RateLimitServiceConfig) ProtoMessage()    {}
func (*RateLimitServiceConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_3154ecf621be8917, []int{0}
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

func (m *RateLimitServiceConfig) GetGrpcService() *core.GrpcService {
	if m != nil {
		return m.GrpcService
	}
	return nil
}

func init() {
	proto.RegisterType((*RateLimitServiceConfig)(nil), "envoy.config.ratelimit.v2.RateLimitServiceConfig")
}

func init() {
	proto.RegisterFile("envoy/config/ratelimit/v2/rls.proto", fileDescriptor_3154ecf621be8917)
}

var fileDescriptor_3154ecf621be8917 = []byte{
	// 262 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x90, 0xcd, 0x4a, 0xc4, 0x30,
	0x14, 0x85, 0x49, 0xad, 0x43, 0xe9, 0xb8, 0x28, 0xb3, 0xf0, 0xa7, 0xa0, 0x88, 0x0a, 0xba, 0x4a,
	0xa0, 0xae, 0xdd, 0xd4, 0x85, 0x30, 0xba, 0x18, 0xea, 0x03, 0xc8, 0x35, 0x13, 0x4b, 0xa0, 0xe6,
	0xc6, 0xe4, 0x4e, 0x70, 0x76, 0xae, 0x7d, 0x24, 0x9f, 0xc0, 0xad, 0xaf, 0xe3, 0x4a, 0xda, 0xd4,
	0x9f, 0x8d, 0xbb, 0x90, 0x9c, 0xf3, 0x9d, 0x8f, 0xe4, 0xc7, 0xca, 0x04, 0x5c, 0x0b, 0x89, 0xe6,
	0x41, 0xb7, 0xc2, 0x01, 0xa9, 0x4e, 0x3f, 0x6a, 0x12, 0xa1, 0x12, 0xae, 0xf3, 0xdc, 0x3a, 0x24,
	0x9c, 0xed, 0x0d, 0x21, 0x1e, 0x43, 0xfc, 0x27, 0xc4, 0x43, 0x55, 0x9e, 0xc4, 0x3e, 0x58, 0xdd,
	0x57, 0x24, 0x3a, 0x25, 0x5a, 0x67, 0xe5, 0x9d, 0x57, 0x2e, 0x68, 0xa9, 0x22, 0xa0, 0xdc, 0x5f,
	0x2d, 0x2d, 0x08, 0x30, 0x06, 0x09, 0x48, 0xa3, 0xf1, 0xc2, 0x13, 0xd0, 0x6a, 0xe4, 0x97, 0x3b,
	0x01, 0x3a, 0xbd, 0x04, 0x52, 0xe2, 0xfb, 0x10, 0x1f, 0x8e, 0x9e, 0xf2, 0xed, 0x06, 0x48, 0xdd,
	0xf4, 0x6b, 0xb7, 0x91, 0x78, 0x39, 0x48, 0xcc, 0xae, 0xf3, 0xad, 0xbf, 0x3b, 0xbb, 0xc9, 0x21,
	0x3b, 0x9b, 0x56, 0x07, 0x3c, 0x9a, 0x82, 0xd5, 0x3c, 0x54, 0xbc, 0xd7, 0xe1, 0x57, 0xce, 0xca,
	0xb1, 0x5b, 0x67, 0x9f, 0xf5, 0xe6, 0x2b, 0x4b, 0x0a, 0xd6, 0x4c, 0xdb, 0xdf, 0xeb, 0x79, 0x9a,
	0xb1, 0x22, 0x99, 0xa7, 0xd9, 0x46, 0x91, 0xd6, 0x17, 0x6f, 0x2f, 0xef, 0x1f, 0x93, 0xa4, 0x60,
	0xf9, 0xa9, 0xc6, 0x88, 0xb3, 0x0e, 0x9f, 0xd7, 0xfc, 0xdf, 0x3f, 0xa8, 0xb3, 0xa6, 0xf3, 0x8b,
	0xde, 0x77, 0xc1, 0xee, 0x27, 0x83, 0xf8, 0xf9, 0x57, 0x00, 0x00, 0x00, 0xff, 0xff, 0x37, 0x76,
	0xe2, 0x11, 0x58, 0x01, 0x00, 0x00,
}
