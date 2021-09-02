// Code generated by protoc-gen-go. DO NOT EDIT.
// source: model.proto

package v1

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import wrappers "github.com/golang/protobuf/ptypes/wrappers"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type MatchString_MatchStringType int32

const (
	MatchString_EXACT MatchString_MatchStringType = 0
	MatchString_REGEX MatchString_MatchStringType = 1
)

var MatchString_MatchStringType_name = map[int32]string{
	0: "EXACT",
	1: "REGEX",
}
var MatchString_MatchStringType_value = map[string]int32{
	"EXACT": 0,
	"REGEX": 1,
}

func (x MatchString_MatchStringType) String() string {
	return proto.EnumName(MatchString_MatchStringType_name, int32(x))
}
func (MatchString_MatchStringType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_model_534f8598bf3352f4, []int{1, 0}
}

type MatchString_ValueType int32

const (
	MatchString_TEXT      MatchString_ValueType = 0
	MatchString_PARAMETER MatchString_ValueType = 1
	MatchString_VARIABLE  MatchString_ValueType = 2
)

var MatchString_ValueType_name = map[int32]string{
	0: "TEXT",
	1: "PARAMETER",
	2: "VARIABLE",
}
var MatchString_ValueType_value = map[string]int32{
	"TEXT":      0,
	"PARAMETER": 1,
	"VARIABLE":  2,
}

func (x MatchString_ValueType) String() string {
	return proto.EnumName(MatchString_ValueType_name, int32(x))
}
func (MatchString_ValueType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_model_534f8598bf3352f4, []int{1, 1}
}

type Location struct {
	Region               *wrappers.StringValue `protobuf:"bytes,1,opt,name=region,proto3" json:"region,omitempty"`
	Zone                 *wrappers.StringValue `protobuf:"bytes,2,opt,name=zone,proto3" json:"zone,omitempty"`
	Campus               *wrappers.StringValue `protobuf:"bytes,3,opt,name=campus,proto3" json:"campus,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *Location) Reset()         { *m = Location{} }
func (m *Location) String() string { return proto.CompactTextString(m) }
func (*Location) ProtoMessage()    {}
func (*Location) Descriptor() ([]byte, []int) {
	return fileDescriptor_model_534f8598bf3352f4, []int{0}
}
func (m *Location) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Location.Unmarshal(m, b)
}
func (m *Location) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Location.Marshal(b, m, deterministic)
}
func (dst *Location) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Location.Merge(dst, src)
}
func (m *Location) XXX_Size() int {
	return xxx_messageInfo_Location.Size(m)
}
func (m *Location) XXX_DiscardUnknown() {
	xxx_messageInfo_Location.DiscardUnknown(m)
}

var xxx_messageInfo_Location proto.InternalMessageInfo

func (m *Location) GetRegion() *wrappers.StringValue {
	if m != nil {
		return m.Region
	}
	return nil
}

func (m *Location) GetZone() *wrappers.StringValue {
	if m != nil {
		return m.Zone
	}
	return nil
}

func (m *Location) GetCampus() *wrappers.StringValue {
	if m != nil {
		return m.Campus
	}
	return nil
}

type MatchString struct {
	Type                 MatchString_MatchStringType `protobuf:"varint,1,opt,name=type,proto3,enum=v1.MatchString_MatchStringType" json:"type,omitempty"`
	Value                *wrappers.StringValue       `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	ValueType            MatchString_ValueType       `protobuf:"varint,3,opt,name=value_type,json=valueType,proto3,enum=v1.MatchString_ValueType" json:"value_type,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                    `json:"-"`
	XXX_unrecognized     []byte                      `json:"-"`
	XXX_sizecache        int32                       `json:"-"`
}

func (m *MatchString) Reset()         { *m = MatchString{} }
func (m *MatchString) String() string { return proto.CompactTextString(m) }
func (*MatchString) ProtoMessage()    {}
func (*MatchString) Descriptor() ([]byte, []int) {
	return fileDescriptor_model_534f8598bf3352f4, []int{1}
}
func (m *MatchString) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MatchString.Unmarshal(m, b)
}
func (m *MatchString) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MatchString.Marshal(b, m, deterministic)
}
func (dst *MatchString) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MatchString.Merge(dst, src)
}
func (m *MatchString) XXX_Size() int {
	return xxx_messageInfo_MatchString.Size(m)
}
func (m *MatchString) XXX_DiscardUnknown() {
	xxx_messageInfo_MatchString.DiscardUnknown(m)
}

var xxx_messageInfo_MatchString proto.InternalMessageInfo

func (m *MatchString) GetType() MatchString_MatchStringType {
	if m != nil {
		return m.Type
	}
	return MatchString_EXACT
}

func (m *MatchString) GetValue() *wrappers.StringValue {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *MatchString) GetValueType() MatchString_ValueType {
	if m != nil {
		return m.ValueType
	}
	return MatchString_TEXT
}

func init() {
	proto.RegisterType((*Location)(nil), "v1.Location")
	proto.RegisterType((*MatchString)(nil), "v1.MatchString")
	proto.RegisterEnum("v1.MatchString_MatchStringType", MatchString_MatchStringType_name, MatchString_MatchStringType_value)
	proto.RegisterEnum("v1.MatchString_ValueType", MatchString_ValueType_name, MatchString_ValueType_value)
}

func init() { proto.RegisterFile("model.proto", fileDescriptor_model_534f8598bf3352f4) }

var fileDescriptor_model_534f8598bf3352f4 = []byte{
	// 291 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x8e, 0xc1, 0x4e, 0xb3, 0x40,
	0x14, 0x85, 0x7f, 0x28, 0x6d, 0xe0, 0xf2, 0xab, 0x64, 0x56, 0x68, 0x8c, 0x1a, 0x36, 0xba, 0x9a,
	0x5a, 0xea, 0xc2, 0x2d, 0x9a, 0x89, 0x31, 0x69, 0x13, 0x33, 0x92, 0x86, 0x9d, 0xa1, 0x38, 0x22,
	0x09, 0x65, 0x26, 0x14, 0x30, 0xf5, 0x1d, 0x7c, 0x0f, 0x1f, 0xd3, 0x70, 0xb1, 0xa6, 0xe9, 0x8a,
	0xdd, 0xc9, 0xcc, 0x77, 0xee, 0x77, 0xc0, 0x5e, 0xc9, 0x57, 0x91, 0x53, 0x55, 0xca, 0x4a, 0x12,
	0xbd, 0x99, 0x9c, 0x9c, 0xa5, 0x52, 0xa6, 0xb9, 0x18, 0xe3, 0xcb, 0xb2, 0x7e, 0x1b, 0x7f, 0x94,
	0xb1, 0x52, 0xa2, 0x5c, 0x77, 0x8c, 0xf7, 0xad, 0x81, 0x39, 0x93, 0x49, 0x5c, 0x65, 0xb2, 0x20,
	0x37, 0x30, 0x2a, 0x45, 0x9a, 0xc9, 0xc2, 0xd5, 0x2e, 0xb4, 0x2b, 0xdb, 0x3f, 0xa5, 0x5d, 0x9b,
	0x6e, 0xdb, 0xf4, 0xb9, 0x2a, 0xb3, 0x22, 0x5d, 0xc4, 0x79, 0x2d, 0xf8, 0x2f, 0x4b, 0xae, 0xc1,
	0xf8, 0x94, 0x85, 0x70, 0xf5, 0x1e, 0x1d, 0x24, 0x5b, 0x4f, 0x12, 0xaf, 0x54, 0xbd, 0x76, 0x07,
	0x7d, 0x3c, 0x1d, 0xeb, 0x7d, 0xe9, 0x60, 0xcf, 0xe3, 0x2a, 0x79, 0xef, 0x3e, 0xc9, 0x14, 0x8c,
	0x6a, 0xa3, 0x04, 0x6e, 0x3d, 0xf4, 0xcf, 0x69, 0x33, 0xa1, 0x3b, 0xdf, 0xbb, 0x39, 0xdc, 0x28,
	0xc1, 0x11, 0x26, 0x3e, 0x0c, 0x9b, 0xf6, 0x6a, 0xaf, 0xb5, 0x1d, 0x4a, 0x6e, 0x01, 0x30, 0xbc,
	0xa0, 0x6e, 0x80, 0xba, 0xe3, 0x7d, 0x1d, 0x36, 0x50, 0x64, 0x35, 0xdb, 0xe8, 0x5d, 0xc2, 0xd1,
	0xde, 0x0c, 0x62, 0xc1, 0x90, 0x45, 0xc1, 0x7d, 0xe8, 0xfc, 0x6b, 0x23, 0x67, 0x0f, 0x2c, 0x72,
	0x34, 0xcf, 0x07, 0xeb, 0xef, 0x00, 0x31, 0xc1, 0x08, 0x59, 0xd4, 0x12, 0x07, 0x60, 0x3d, 0x05,
	0x3c, 0x98, 0xb3, 0x90, 0x71, 0x47, 0x23, 0xff, 0xc1, 0x5c, 0x04, 0xfc, 0x31, 0xb8, 0x9b, 0x31,
	0x47, 0x5f, 0x8e, 0x70, 0xf3, 0xf4, 0x27, 0x00, 0x00, 0xff, 0xff, 0x28, 0x8b, 0x9e, 0xbf, 0xf4,
	0x01, 0x00, 0x00,
}
