// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/api/v2/discovery.proto

package envoy_api_v2

import (
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	core "github.com/polarismesh/polaris-server/apiserver/xdsserver/go-control-plane/envoy/api/v2/core"
	proto "github.com/golang/protobuf/proto"
	any "github.com/golang/protobuf/ptypes/any"
	status "google.golang.org/genproto/googleapis/rpc/status"
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

type DiscoveryRequest struct {
	VersionInfo          string         `protobuf:"bytes,1,opt,name=version_info,json=versionInfo,proto3" json:"version_info,omitempty"`
	Node                 *core.Node     `protobuf:"bytes,2,opt,name=node,proto3" json:"node,omitempty"`
	ResourceNames        []string       `protobuf:"bytes,3,rep,name=resource_names,json=resourceNames,proto3" json:"resource_names,omitempty"`
	TypeUrl              string         `protobuf:"bytes,4,opt,name=type_url,json=typeUrl,proto3" json:"type_url,omitempty"`
	ResponseNonce        string         `protobuf:"bytes,5,opt,name=response_nonce,json=responseNonce,proto3" json:"response_nonce,omitempty"`
	ErrorDetail          *status.Status `protobuf:"bytes,6,opt,name=error_detail,json=errorDetail,proto3" json:"error_detail,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *DiscoveryRequest) Reset()         { *m = DiscoveryRequest{} }
func (m *DiscoveryRequest) String() string { return proto.CompactTextString(m) }
func (*DiscoveryRequest) ProtoMessage()    {}
func (*DiscoveryRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2c7365e287e5c035, []int{0}
}

func (m *DiscoveryRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DiscoveryRequest.Unmarshal(m, b)
}
func (m *DiscoveryRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DiscoveryRequest.Marshal(b, m, deterministic)
}
func (m *DiscoveryRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DiscoveryRequest.Merge(m, src)
}
func (m *DiscoveryRequest) XXX_Size() int {
	return xxx_messageInfo_DiscoveryRequest.Size(m)
}
func (m *DiscoveryRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DiscoveryRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DiscoveryRequest proto.InternalMessageInfo

func (m *DiscoveryRequest) GetVersionInfo() string {
	if m != nil {
		return m.VersionInfo
	}
	return ""
}

func (m *DiscoveryRequest) GetNode() *core.Node {
	if m != nil {
		return m.Node
	}
	return nil
}

func (m *DiscoveryRequest) GetResourceNames() []string {
	if m != nil {
		return m.ResourceNames
	}
	return nil
}

func (m *DiscoveryRequest) GetTypeUrl() string {
	if m != nil {
		return m.TypeUrl
	}
	return ""
}

func (m *DiscoveryRequest) GetResponseNonce() string {
	if m != nil {
		return m.ResponseNonce
	}
	return ""
}

func (m *DiscoveryRequest) GetErrorDetail() *status.Status {
	if m != nil {
		return m.ErrorDetail
	}
	return nil
}

type DiscoveryResponse struct {
	VersionInfo          string             `protobuf:"bytes,1,opt,name=version_info,json=versionInfo,proto3" json:"version_info,omitempty"`
	Resources            []*any.Any         `protobuf:"bytes,2,rep,name=resources,proto3" json:"resources,omitempty"`
	Canary               bool               `protobuf:"varint,3,opt,name=canary,proto3" json:"canary,omitempty"`
	TypeUrl              string             `protobuf:"bytes,4,opt,name=type_url,json=typeUrl,proto3" json:"type_url,omitempty"`
	Nonce                string             `protobuf:"bytes,5,opt,name=nonce,proto3" json:"nonce,omitempty"`
	ControlPlane         *core.ControlPlane `protobuf:"bytes,6,opt,name=control_plane,json=controlPlane,proto3" json:"control_plane,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *DiscoveryResponse) Reset()         { *m = DiscoveryResponse{} }
func (m *DiscoveryResponse) String() string { return proto.CompactTextString(m) }
func (*DiscoveryResponse) ProtoMessage()    {}
func (*DiscoveryResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_2c7365e287e5c035, []int{1}
}

func (m *DiscoveryResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DiscoveryResponse.Unmarshal(m, b)
}
func (m *DiscoveryResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DiscoveryResponse.Marshal(b, m, deterministic)
}
func (m *DiscoveryResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DiscoveryResponse.Merge(m, src)
}
func (m *DiscoveryResponse) XXX_Size() int {
	return xxx_messageInfo_DiscoveryResponse.Size(m)
}
func (m *DiscoveryResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DiscoveryResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DiscoveryResponse proto.InternalMessageInfo

func (m *DiscoveryResponse) GetVersionInfo() string {
	if m != nil {
		return m.VersionInfo
	}
	return ""
}

func (m *DiscoveryResponse) GetResources() []*any.Any {
	if m != nil {
		return m.Resources
	}
	return nil
}

func (m *DiscoveryResponse) GetCanary() bool {
	if m != nil {
		return m.Canary
	}
	return false
}

func (m *DiscoveryResponse) GetTypeUrl() string {
	if m != nil {
		return m.TypeUrl
	}
	return ""
}

func (m *DiscoveryResponse) GetNonce() string {
	if m != nil {
		return m.Nonce
	}
	return ""
}

func (m *DiscoveryResponse) GetControlPlane() *core.ControlPlane {
	if m != nil {
		return m.ControlPlane
	}
	return nil
}

type DeltaDiscoveryRequest struct {
	Node                     *core.Node        `protobuf:"bytes,1,opt,name=node,proto3" json:"node,omitempty"`
	TypeUrl                  string            `protobuf:"bytes,2,opt,name=type_url,json=typeUrl,proto3" json:"type_url,omitempty"`
	ResourceNamesSubscribe   []string          `protobuf:"bytes,3,rep,name=resource_names_subscribe,json=resourceNamesSubscribe,proto3" json:"resource_names_subscribe,omitempty"`
	ResourceNamesUnsubscribe []string          `protobuf:"bytes,4,rep,name=resource_names_unsubscribe,json=resourceNamesUnsubscribe,proto3" json:"resource_names_unsubscribe,omitempty"`
	InitialResourceVersions  map[string]string `protobuf:"bytes,5,rep,name=initial_resource_versions,json=initialResourceVersions,proto3" json:"initial_resource_versions,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	ResponseNonce            string            `protobuf:"bytes,6,opt,name=response_nonce,json=responseNonce,proto3" json:"response_nonce,omitempty"`
	ErrorDetail              *status.Status    `protobuf:"bytes,7,opt,name=error_detail,json=errorDetail,proto3" json:"error_detail,omitempty"`
	XXX_NoUnkeyedLiteral     struct{}          `json:"-"`
	XXX_unrecognized         []byte            `json:"-"`
	XXX_sizecache            int32             `json:"-"`
}

func (m *DeltaDiscoveryRequest) Reset()         { *m = DeltaDiscoveryRequest{} }
func (m *DeltaDiscoveryRequest) String() string { return proto.CompactTextString(m) }
func (*DeltaDiscoveryRequest) ProtoMessage()    {}
func (*DeltaDiscoveryRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2c7365e287e5c035, []int{2}
}

func (m *DeltaDiscoveryRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeltaDiscoveryRequest.Unmarshal(m, b)
}
func (m *DeltaDiscoveryRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeltaDiscoveryRequest.Marshal(b, m, deterministic)
}
func (m *DeltaDiscoveryRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeltaDiscoveryRequest.Merge(m, src)
}
func (m *DeltaDiscoveryRequest) XXX_Size() int {
	return xxx_messageInfo_DeltaDiscoveryRequest.Size(m)
}
func (m *DeltaDiscoveryRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeltaDiscoveryRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeltaDiscoveryRequest proto.InternalMessageInfo

func (m *DeltaDiscoveryRequest) GetNode() *core.Node {
	if m != nil {
		return m.Node
	}
	return nil
}

func (m *DeltaDiscoveryRequest) GetTypeUrl() string {
	if m != nil {
		return m.TypeUrl
	}
	return ""
}

func (m *DeltaDiscoveryRequest) GetResourceNamesSubscribe() []string {
	if m != nil {
		return m.ResourceNamesSubscribe
	}
	return nil
}

func (m *DeltaDiscoveryRequest) GetResourceNamesUnsubscribe() []string {
	if m != nil {
		return m.ResourceNamesUnsubscribe
	}
	return nil
}

func (m *DeltaDiscoveryRequest) GetInitialResourceVersions() map[string]string {
	if m != nil {
		return m.InitialResourceVersions
	}
	return nil
}

func (m *DeltaDiscoveryRequest) GetResponseNonce() string {
	if m != nil {
		return m.ResponseNonce
	}
	return ""
}

func (m *DeltaDiscoveryRequest) GetErrorDetail() *status.Status {
	if m != nil {
		return m.ErrorDetail
	}
	return nil
}

type DeltaDiscoveryResponse struct {
	SystemVersionInfo    string      `protobuf:"bytes,1,opt,name=system_version_info,json=systemVersionInfo,proto3" json:"system_version_info,omitempty"`
	Resources            []*Resource `protobuf:"bytes,2,rep,name=resources,proto3" json:"resources,omitempty"`
	TypeUrl              string      `protobuf:"bytes,4,opt,name=type_url,json=typeUrl,proto3" json:"type_url,omitempty"`
	RemovedResources     []string    `protobuf:"bytes,6,rep,name=removed_resources,json=removedResources,proto3" json:"removed_resources,omitempty"`
	Nonce                string      `protobuf:"bytes,5,opt,name=nonce,proto3" json:"nonce,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *DeltaDiscoveryResponse) Reset()         { *m = DeltaDiscoveryResponse{} }
func (m *DeltaDiscoveryResponse) String() string { return proto.CompactTextString(m) }
func (*DeltaDiscoveryResponse) ProtoMessage()    {}
func (*DeltaDiscoveryResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_2c7365e287e5c035, []int{3}
}

func (m *DeltaDiscoveryResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeltaDiscoveryResponse.Unmarshal(m, b)
}
func (m *DeltaDiscoveryResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeltaDiscoveryResponse.Marshal(b, m, deterministic)
}
func (m *DeltaDiscoveryResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeltaDiscoveryResponse.Merge(m, src)
}
func (m *DeltaDiscoveryResponse) XXX_Size() int {
	return xxx_messageInfo_DeltaDiscoveryResponse.Size(m)
}
func (m *DeltaDiscoveryResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeltaDiscoveryResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeltaDiscoveryResponse proto.InternalMessageInfo

func (m *DeltaDiscoveryResponse) GetSystemVersionInfo() string {
	if m != nil {
		return m.SystemVersionInfo
	}
	return ""
}

func (m *DeltaDiscoveryResponse) GetResources() []*Resource {
	if m != nil {
		return m.Resources
	}
	return nil
}

func (m *DeltaDiscoveryResponse) GetTypeUrl() string {
	if m != nil {
		return m.TypeUrl
	}
	return ""
}

func (m *DeltaDiscoveryResponse) GetRemovedResources() []string {
	if m != nil {
		return m.RemovedResources
	}
	return nil
}

func (m *DeltaDiscoveryResponse) GetNonce() string {
	if m != nil {
		return m.Nonce
	}
	return ""
}

type Resource struct {
	Name                 string   `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Aliases              []string `protobuf:"bytes,4,rep,name=aliases,proto3" json:"aliases,omitempty"`
	Version              string   `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"`
	Resource             *any.Any `protobuf:"bytes,2,opt,name=resource,proto3" json:"resource,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Resource) Reset()         { *m = Resource{} }
func (m *Resource) String() string { return proto.CompactTextString(m) }
func (*Resource) ProtoMessage()    {}
func (*Resource) Descriptor() ([]byte, []int) {
	return fileDescriptor_2c7365e287e5c035, []int{4}
}

func (m *Resource) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Resource.Unmarshal(m, b)
}
func (m *Resource) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Resource.Marshal(b, m, deterministic)
}
func (m *Resource) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Resource.Merge(m, src)
}
func (m *Resource) XXX_Size() int {
	return xxx_messageInfo_Resource.Size(m)
}
func (m *Resource) XXX_DiscardUnknown() {
	xxx_messageInfo_Resource.DiscardUnknown(m)
}

var xxx_messageInfo_Resource proto.InternalMessageInfo

func (m *Resource) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Resource) GetAliases() []string {
	if m != nil {
		return m.Aliases
	}
	return nil
}

func (m *Resource) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *Resource) GetResource() *any.Any {
	if m != nil {
		return m.Resource
	}
	return nil
}

func init() {
	proto.RegisterType((*DiscoveryRequest)(nil), "envoy.api.v2.DiscoveryRequest")
	proto.RegisterType((*DiscoveryResponse)(nil), "envoy.api.v2.DiscoveryResponse")
	proto.RegisterType((*DeltaDiscoveryRequest)(nil), "envoy.api.v2.DeltaDiscoveryRequest")
	proto.RegisterMapType((map[string]string)(nil), "envoy.api.v2.DeltaDiscoveryRequest.InitialResourceVersionsEntry")
	proto.RegisterType((*DeltaDiscoveryResponse)(nil), "envoy.api.v2.DeltaDiscoveryResponse")
	proto.RegisterType((*Resource)(nil), "envoy.api.v2.Resource")
}

func init() { proto.RegisterFile("envoy/api/v2/discovery.proto", fileDescriptor_2c7365e287e5c035) }

var fileDescriptor_2c7365e287e5c035 = []byte{
	// 720 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x55, 0x4d, 0x6f, 0xd3, 0x4a,
	0x14, 0x95, 0xf3, 0xd5, 0x64, 0x92, 0x56, 0xe9, 0xbc, 0xbe, 0xd4, 0x8d, 0xfa, 0x1e, 0x21, 0x12,
	0x52, 0xa4, 0x4a, 0x36, 0x4a, 0x41, 0xaa, 0x10, 0x0b, 0x28, 0x61, 0x51, 0x16, 0x55, 0xe5, 0xaa,
	0x15, 0x3b, 0x6b, 0xe2, 0xdc, 0x56, 0x23, 0x9c, 0x19, 0x33, 0x63, 0x5b, 0x58, 0x62, 0x81, 0x10,
	0x7b, 0xb6, 0xfc, 0x16, 0x7e, 0x01, 0x5b, 0xf6, 0xfc, 0x0a, 0x76, 0x6c, 0x00, 0x79, 0x3c, 0x4e,
	0xec, 0x36, 0xaa, 0xb2, 0xf3, 0x9d, 0x7b, 0xe6, 0xce, 0x3d, 0xf7, 0x9c, 0x9b, 0xa0, 0x7d, 0x60,
	0x31, 0x4f, 0x6c, 0x12, 0x50, 0x3b, 0x1e, 0xdb, 0x33, 0x2a, 0x3d, 0x1e, 0x83, 0x48, 0xac, 0x40,
	0xf0, 0x90, 0xe3, 0x8e, 0xca, 0x5a, 0x24, 0xa0, 0x56, 0x3c, 0xee, 0x97, 0xb1, 0x1e, 0x17, 0x60,
	0x4f, 0x89, 0x84, 0x0c, 0xdb, 0xdf, 0xbb, 0xe6, 0xfc, 0xda, 0x07, 0x5b, 0x45, 0xd3, 0xe8, 0xca,
	0x26, 0x4c, 0x97, 0xe9, 0xef, 0xea, 0x94, 0x08, 0x3c, 0x5b, 0x86, 0x24, 0x8c, 0xa4, 0x4e, 0xfc,
	0x1f, 0xcd, 0x02, 0x62, 0x13, 0xc6, 0x78, 0x48, 0x42, 0xca, 0x99, 0xb4, 0xe7, 0xf4, 0x5a, 0x90,
	0x30, 0xaf, 0xf9, 0xdf, 0xad, 0x7c, 0xf1, 0xfa, 0xf0, 0x63, 0x05, 0x75, 0x27, 0x79, 0xcb, 0x0e,
	0xbc, 0x8d, 0x40, 0x86, 0xf8, 0x3e, 0xea, 0xc4, 0x20, 0x24, 0xe5, 0xcc, 0xa5, 0xec, 0x8a, 0x9b,
	0xc6, 0xc0, 0x18, 0xb5, 0x9c, 0xb6, 0x3e, 0x3b, 0x61, 0x57, 0x1c, 0x1f, 0xa0, 0x1a, 0xe3, 0x33,
	0x30, 0x2b, 0x03, 0x63, 0xd4, 0x1e, 0xef, 0x5a, 0x45, 0x96, 0x56, 0xca, 0xcb, 0x3a, 0xe5, 0x33,
	0x70, 0x14, 0x08, 0x3f, 0x40, 0x5b, 0x02, 0x24, 0x8f, 0x84, 0x07, 0x2e, 0x23, 0x73, 0x90, 0x66,
	0x75, 0x50, 0x1d, 0xb5, 0x9c, 0xcd, 0xfc, 0xf4, 0x34, 0x3d, 0xc4, 0x7b, 0xa8, 0x19, 0x26, 0x01,
	0xb8, 0x91, 0xf0, 0xcd, 0x9a, 0x7a, 0x72, 0x23, 0x8d, 0x2f, 0x84, 0xaf, 0x2b, 0x04, 0x9c, 0x49,
	0x70, 0x19, 0x67, 0x1e, 0x98, 0x75, 0x05, 0xd8, 0xcc, 0x4f, 0x4f, 0xd3, 0x43, 0xfc, 0x18, 0x75,
	0x40, 0x08, 0x2e, 0xdc, 0x19, 0x84, 0x84, 0xfa, 0x66, 0x43, 0x75, 0x87, 0xad, 0x6c, 0x78, 0x96,
	0x08, 0x3c, 0xeb, 0x5c, 0xb1, 0x77, 0xda, 0x0a, 0x37, 0x51, 0xb0, 0xe1, 0x2f, 0x03, 0x6d, 0x17,
	0x86, 0x90, 0x55, 0x5c, 0x67, 0x0a, 0x63, 0xd4, 0xca, 0x29, 0x48, 0xb3, 0x32, 0xa8, 0x8e, 0xda,
	0xe3, 0x9d, 0xfc, 0xb1, 0x5c, 0x44, 0xeb, 0x39, 0x4b, 0x9c, 0x25, 0x0c, 0xf7, 0x50, 0xc3, 0x23,
	0x8c, 0x88, 0xc4, 0xac, 0x0e, 0x8c, 0x51, 0xd3, 0xd1, 0xd1, 0x5d, 0xec, 0x77, 0x50, 0xbd, 0x48,
	0x3a, 0x0b, 0xf0, 0x04, 0x6d, 0x7a, 0x9c, 0x85, 0x82, 0xfb, 0x6e, 0xe0, 0x13, 0x06, 0x9a, 0xed,
	0xbd, 0x15, 0x5a, 0xbc, 0xc8, 0x70, 0x67, 0x29, 0xcc, 0xe9, 0x78, 0x85, 0x68, 0xf8, 0xa7, 0x8a,
	0xfe, 0x9d, 0x80, 0x1f, 0x92, 0x5b, 0x2e, 0xc8, 0x25, 0x36, 0xd6, 0x91, 0xb8, 0xd8, 0x7d, 0xa5,
	0xdc, 0xfd, 0x11, 0x32, 0xcb, 0xea, 0xbb, 0x32, 0x9a, 0x4a, 0x4f, 0xd0, 0x29, 0x68, 0x1f, 0xf4,
	0x4a, 0x3e, 0x38, 0xcf, 0xb3, 0xf8, 0x29, 0xea, 0xdf, 0xb8, 0x19, 0xb1, 0xe5, 0xdd, 0x9a, 0xba,
	0x6b, 0x96, 0xee, 0x5e, 0x2c, 0xf3, 0xf8, 0x3d, 0xda, 0xa3, 0x8c, 0x86, 0x94, 0xf8, 0xee, 0xa2,
	0x8a, 0x16, 0x4f, 0x9a, 0x75, 0x25, 0xd6, 0xb3, 0x32, 0xa9, 0x95, 0x73, 0xb0, 0x4e, 0xb2, 0x22,
	0x8e, 0xae, 0x71, 0xa9, 0x4b, 0xbc, 0x64, 0xa1, 0x48, 0x9c, 0x5d, 0xba, 0x3a, 0xbb, 0xc2, 0xb1,
	0x8d, 0x75, 0x1c, 0xbb, 0xb1, 0x96, 0x63, 0xfb, 0xaf, 0xd0, 0xfe, 0x5d, 0x6d, 0xe1, 0x2e, 0xaa,
	0xbe, 0x81, 0x44, 0x5b, 0x36, 0xfd, 0x4c, 0x3d, 0x14, 0x13, 0x3f, 0x02, 0xad, 0x4e, 0x16, 0x3c,
	0xa9, 0x1c, 0x19, 0xc3, 0x1f, 0x06, 0xea, 0xdd, 0x64, 0xae, 0x57, 0xc0, 0x42, 0xff, 0xc8, 0x44,
	0x86, 0x30, 0x77, 0x57, 0x6c, 0xc2, 0x76, 0x96, 0xba, 0x2c, 0xec, 0xc3, 0xa3, 0xdb, 0xfb, 0xd0,
	0x2b, 0x8f, 0x38, 0x6f, 0xb7, 0xb8, 0x11, 0x77, 0x38, 0xff, 0x00, 0x6d, 0x0b, 0x98, 0xf3, 0x18,
	0x66, 0xee, 0xb2, 0x70, 0x43, 0x09, 0xdf, 0xd5, 0x09, 0x67, 0x51, 0x67, 0xe5, 0x9a, 0x0c, 0x3f,
	0x19, 0xa8, 0x99, 0x63, 0x30, 0x46, 0xb5, 0xd4, 0x48, 0x6a, 0xf5, 0x5a, 0x8e, 0xfa, 0xc6, 0x26,
	0xda, 0x20, 0x3e, 0x25, 0x12, 0xa4, 0xb6, 0x54, 0x1e, 0xa6, 0x19, 0xcd, 0x5b, 0x53, 0xce, 0x43,
	0xfc, 0x10, 0x35, 0xf3, 0x7e, 0xf4, 0x4f, 0xe0, 0xea, 0xbd, 0x5f, 0xa0, 0x8e, 0x5f, 0xff, 0xfc,
	0xf2, 0xfb, 0x73, 0x7d, 0x1f, 0xf7, 0xb3, 0x71, 0x48, 0x10, 0x31, 0xf5, 0xc0, 0x5a, 0xfe, 0x5d,
	0xc4, 0x87, 0x5f, 0x3f, 0x7c, 0xfb, 0xde, 0xa8, 0x74, 0x0d, 0xd4, 0xa7, 0x3c, 0x9b, 0x5a, 0x20,
	0xf8, 0xbb, 0xa4, 0x34, 0xc0, 0xe3, 0xad, 0x85, 0x4a, 0x67, 0xe9, 0x43, 0x67, 0xc6, 0xb4, 0xa1,
	0x5e, 0x3c, 0xfc, 0x1b, 0x00, 0x00, 0xff, 0xff, 0x0f, 0x0b, 0x0b, 0xc6, 0x88, 0x06, 0x00, 0x00,
}
