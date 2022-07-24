// Code generated by protoc-gen-go. DO NOT EDIT.
// source: config_file.proto

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

type ConfigFileGroup struct {
	Id                   *wrappers.UInt64Value   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 *wrappers.StringValue   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Namespace            *wrappers.StringValue   `protobuf:"bytes,3,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Comment              *wrappers.StringValue   `protobuf:"bytes,4,opt,name=comment,proto3" json:"comment,omitempty"`
	CreateTime           *wrappers.StringValue   `protobuf:"bytes,5,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	CreateBy             *wrappers.StringValue   `protobuf:"bytes,6,opt,name=create_by,json=createBy,proto3" json:"create_by,omitempty"`
	ModifyTime           *wrappers.StringValue   `protobuf:"bytes,7,opt,name=modify_time,json=modifyTime,proto3" json:"modify_time,omitempty"`
	ModifyBy             *wrappers.StringValue   `protobuf:"bytes,8,opt,name=modify_by,json=modifyBy,proto3" json:"modify_by,omitempty"`
	FileCount            *wrappers.UInt64Value   `protobuf:"bytes,9,opt,name=fileCount,proto3" json:"fileCount,omitempty"`
	UserIds              []*wrappers.StringValue `protobuf:"bytes,10,rep,name=user_ids,proto3" json:"user_ids,omitempty"`
	GroupIds             []*wrappers.StringValue `protobuf:"bytes,11,rep,name=group_ids,proto3" json:"group_ids,omitempty"`
	RemoveUserIds        []*wrappers.StringValue `protobuf:"bytes,13,rep,name=remove_user_ids,proto3" json:"remove_user_ids,omitempty"`
	RemoveGroupIds       []*wrappers.StringValue `protobuf:"bytes,14,rep,name=remove_group_ids,proto3" json:"remove_group_ids,omitempty"`
	Editable             *wrappers.BoolValue     `protobuf:"bytes,15,opt,name=editable,proto3" json:"editable,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *ConfigFileGroup) Reset()         { *m = ConfigFileGroup{} }
func (m *ConfigFileGroup) String() string { return proto.CompactTextString(m) }
func (*ConfigFileGroup) ProtoMessage()    {}
func (*ConfigFileGroup) Descriptor() ([]byte, []int) {
	return fileDescriptor_config_file_092c8b107355d57c, []int{0}
}
func (m *ConfigFileGroup) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConfigFileGroup.Unmarshal(m, b)
}
func (m *ConfigFileGroup) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConfigFileGroup.Marshal(b, m, deterministic)
}
func (dst *ConfigFileGroup) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConfigFileGroup.Merge(dst, src)
}
func (m *ConfigFileGroup) XXX_Size() int {
	return xxx_messageInfo_ConfigFileGroup.Size(m)
}
func (m *ConfigFileGroup) XXX_DiscardUnknown() {
	xxx_messageInfo_ConfigFileGroup.DiscardUnknown(m)
}

var xxx_messageInfo_ConfigFileGroup proto.InternalMessageInfo

func (m *ConfigFileGroup) GetId() *wrappers.UInt64Value {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *ConfigFileGroup) GetName() *wrappers.StringValue {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *ConfigFileGroup) GetNamespace() *wrappers.StringValue {
	if m != nil {
		return m.Namespace
	}
	return nil
}

func (m *ConfigFileGroup) GetComment() *wrappers.StringValue {
	if m != nil {
		return m.Comment
	}
	return nil
}

func (m *ConfigFileGroup) GetCreateTime() *wrappers.StringValue {
	if m != nil {
		return m.CreateTime
	}
	return nil
}

func (m *ConfigFileGroup) GetCreateBy() *wrappers.StringValue {
	if m != nil {
		return m.CreateBy
	}
	return nil
}

func (m *ConfigFileGroup) GetModifyTime() *wrappers.StringValue {
	if m != nil {
		return m.ModifyTime
	}
	return nil
}

func (m *ConfigFileGroup) GetModifyBy() *wrappers.StringValue {
	if m != nil {
		return m.ModifyBy
	}
	return nil
}

func (m *ConfigFileGroup) GetFileCount() *wrappers.UInt64Value {
	if m != nil {
		return m.FileCount
	}
	return nil
}

func (m *ConfigFileGroup) GetUserIds() []*wrappers.StringValue {
	if m != nil {
		return m.UserIds
	}
	return nil
}

func (m *ConfigFileGroup) GetGroupIds() []*wrappers.StringValue {
	if m != nil {
		return m.GroupIds
	}
	return nil
}

func (m *ConfigFileGroup) GetRemoveUserIds() []*wrappers.StringValue {
	if m != nil {
		return m.RemoveUserIds
	}
	return nil
}

func (m *ConfigFileGroup) GetRemoveGroupIds() []*wrappers.StringValue {
	if m != nil {
		return m.RemoveGroupIds
	}
	return nil
}

func (m *ConfigFileGroup) GetEditable() *wrappers.BoolValue {
	if m != nil {
		return m.Editable
	}
	return nil
}

type ConfigFile struct {
	Id                   *wrappers.UInt64Value `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 *wrappers.StringValue `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Namespace            *wrappers.StringValue `protobuf:"bytes,3,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Group                *wrappers.StringValue `protobuf:"bytes,4,opt,name=group,proto3" json:"group,omitempty"`
	Content              *wrappers.StringValue `protobuf:"bytes,5,opt,name=content,proto3" json:"content,omitempty"`
	Format               *wrappers.StringValue `protobuf:"bytes,6,opt,name=format,proto3" json:"format,omitempty"`
	Comment              *wrappers.StringValue `protobuf:"bytes,7,opt,name=comment,proto3" json:"comment,omitempty"`
	Status               *wrappers.StringValue `protobuf:"bytes,8,opt,name=status,proto3" json:"status,omitempty"`
	Tags                 []*ConfigFileTag      `protobuf:"bytes,9,rep,name=tags,proto3" json:"tags,omitempty"`
	CreateTime           *wrappers.StringValue `protobuf:"bytes,10,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	CreateBy             *wrappers.StringValue `protobuf:"bytes,11,opt,name=create_by,json=createBy,proto3" json:"create_by,omitempty"`
	ModifyTime           *wrappers.StringValue `protobuf:"bytes,12,opt,name=modify_time,json=modifyTime,proto3" json:"modify_time,omitempty"`
	ModifyBy             *wrappers.StringValue `protobuf:"bytes,13,opt,name=modify_by,json=modifyBy,proto3" json:"modify_by,omitempty"`
	ReleaseTime          *wrappers.StringValue `protobuf:"bytes,14,opt,name=release_time,json=releaseTime,proto3" json:"release_time,omitempty"`
	ReleaseBy            *wrappers.StringValue `protobuf:"bytes,15,opt,name=release_by,json=releaseBy,proto3" json:"release_by,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *ConfigFile) Reset()         { *m = ConfigFile{} }
func (m *ConfigFile) String() string { return proto.CompactTextString(m) }
func (*ConfigFile) ProtoMessage()    {}
func (*ConfigFile) Descriptor() ([]byte, []int) {
	return fileDescriptor_config_file_092c8b107355d57c, []int{1}
}
func (m *ConfigFile) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConfigFile.Unmarshal(m, b)
}
func (m *ConfigFile) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConfigFile.Marshal(b, m, deterministic)
}
func (dst *ConfigFile) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConfigFile.Merge(dst, src)
}
func (m *ConfigFile) XXX_Size() int {
	return xxx_messageInfo_ConfigFile.Size(m)
}
func (m *ConfigFile) XXX_DiscardUnknown() {
	xxx_messageInfo_ConfigFile.DiscardUnknown(m)
}

var xxx_messageInfo_ConfigFile proto.InternalMessageInfo

func (m *ConfigFile) GetId() *wrappers.UInt64Value {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *ConfigFile) GetName() *wrappers.StringValue {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *ConfigFile) GetNamespace() *wrappers.StringValue {
	if m != nil {
		return m.Namespace
	}
	return nil
}

func (m *ConfigFile) GetGroup() *wrappers.StringValue {
	if m != nil {
		return m.Group
	}
	return nil
}

func (m *ConfigFile) GetContent() *wrappers.StringValue {
	if m != nil {
		return m.Content
	}
	return nil
}

func (m *ConfigFile) GetFormat() *wrappers.StringValue {
	if m != nil {
		return m.Format
	}
	return nil
}

func (m *ConfigFile) GetComment() *wrappers.StringValue {
	if m != nil {
		return m.Comment
	}
	return nil
}

func (m *ConfigFile) GetStatus() *wrappers.StringValue {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *ConfigFile) GetTags() []*ConfigFileTag {
	if m != nil {
		return m.Tags
	}
	return nil
}

func (m *ConfigFile) GetCreateTime() *wrappers.StringValue {
	if m != nil {
		return m.CreateTime
	}
	return nil
}

func (m *ConfigFile) GetCreateBy() *wrappers.StringValue {
	if m != nil {
		return m.CreateBy
	}
	return nil
}

func (m *ConfigFile) GetModifyTime() *wrappers.StringValue {
	if m != nil {
		return m.ModifyTime
	}
	return nil
}

func (m *ConfigFile) GetModifyBy() *wrappers.StringValue {
	if m != nil {
		return m.ModifyBy
	}
	return nil
}

func (m *ConfigFile) GetReleaseTime() *wrappers.StringValue {
	if m != nil {
		return m.ReleaseTime
	}
	return nil
}

func (m *ConfigFile) GetReleaseBy() *wrappers.StringValue {
	if m != nil {
		return m.ReleaseBy
	}
	return nil
}

type ConfigFileTag struct {
	Key                  *wrappers.StringValue `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value                *wrappers.StringValue `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *ConfigFileTag) Reset()         { *m = ConfigFileTag{} }
func (m *ConfigFileTag) String() string { return proto.CompactTextString(m) }
func (*ConfigFileTag) ProtoMessage()    {}
func (*ConfigFileTag) Descriptor() ([]byte, []int) {
	return fileDescriptor_config_file_092c8b107355d57c, []int{2}
}
func (m *ConfigFileTag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConfigFileTag.Unmarshal(m, b)
}
func (m *ConfigFileTag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConfigFileTag.Marshal(b, m, deterministic)
}
func (dst *ConfigFileTag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConfigFileTag.Merge(dst, src)
}
func (m *ConfigFileTag) XXX_Size() int {
	return xxx_messageInfo_ConfigFileTag.Size(m)
}
func (m *ConfigFileTag) XXX_DiscardUnknown() {
	xxx_messageInfo_ConfigFileTag.DiscardUnknown(m)
}

var xxx_messageInfo_ConfigFileTag proto.InternalMessageInfo

func (m *ConfigFileTag) GetKey() *wrappers.StringValue {
	if m != nil {
		return m.Key
	}
	return nil
}

func (m *ConfigFileTag) GetValue() *wrappers.StringValue {
	if m != nil {
		return m.Value
	}
	return nil
}

type ConfigFileRelease struct {
	Id                   *wrappers.UInt64Value `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 *wrappers.StringValue `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Namespace            *wrappers.StringValue `protobuf:"bytes,3,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Group                *wrappers.StringValue `protobuf:"bytes,4,opt,name=group,proto3" json:"group,omitempty"`
	FileName             *wrappers.StringValue `protobuf:"bytes,5,opt,name=file_name,json=fileName,proto3" json:"file_name,omitempty"`
	Content              *wrappers.StringValue `protobuf:"bytes,6,opt,name=content,proto3" json:"content,omitempty"`
	Comment              *wrappers.StringValue `protobuf:"bytes,7,opt,name=comment,proto3" json:"comment,omitempty"`
	Md5                  *wrappers.StringValue `protobuf:"bytes,8,opt,name=md5,proto3" json:"md5,omitempty"`
	Version              *wrappers.UInt64Value `protobuf:"bytes,9,opt,name=version,proto3" json:"version,omitempty"`
	CreateTime           *wrappers.StringValue `protobuf:"bytes,10,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	CreateBy             *wrappers.StringValue `protobuf:"bytes,11,opt,name=create_by,json=createBy,proto3" json:"create_by,omitempty"`
	ModifyTime           *wrappers.StringValue `protobuf:"bytes,12,opt,name=modify_time,json=modifyTime,proto3" json:"modify_time,omitempty"`
	ModifyBy             *wrappers.StringValue `protobuf:"bytes,13,opt,name=modify_by,json=modifyBy,proto3" json:"modify_by,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *ConfigFileRelease) Reset()         { *m = ConfigFileRelease{} }
func (m *ConfigFileRelease) String() string { return proto.CompactTextString(m) }
func (*ConfigFileRelease) ProtoMessage()    {}
func (*ConfigFileRelease) Descriptor() ([]byte, []int) {
	return fileDescriptor_config_file_092c8b107355d57c, []int{3}
}
func (m *ConfigFileRelease) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConfigFileRelease.Unmarshal(m, b)
}
func (m *ConfigFileRelease) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConfigFileRelease.Marshal(b, m, deterministic)
}
func (dst *ConfigFileRelease) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConfigFileRelease.Merge(dst, src)
}
func (m *ConfigFileRelease) XXX_Size() int {
	return xxx_messageInfo_ConfigFileRelease.Size(m)
}
func (m *ConfigFileRelease) XXX_DiscardUnknown() {
	xxx_messageInfo_ConfigFileRelease.DiscardUnknown(m)
}

var xxx_messageInfo_ConfigFileRelease proto.InternalMessageInfo

func (m *ConfigFileRelease) GetId() *wrappers.UInt64Value {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *ConfigFileRelease) GetName() *wrappers.StringValue {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *ConfigFileRelease) GetNamespace() *wrappers.StringValue {
	if m != nil {
		return m.Namespace
	}
	return nil
}

func (m *ConfigFileRelease) GetGroup() *wrappers.StringValue {
	if m != nil {
		return m.Group
	}
	return nil
}

func (m *ConfigFileRelease) GetFileName() *wrappers.StringValue {
	if m != nil {
		return m.FileName
	}
	return nil
}

func (m *ConfigFileRelease) GetContent() *wrappers.StringValue {
	if m != nil {
		return m.Content
	}
	return nil
}

func (m *ConfigFileRelease) GetComment() *wrappers.StringValue {
	if m != nil {
		return m.Comment
	}
	return nil
}

func (m *ConfigFileRelease) GetMd5() *wrappers.StringValue {
	if m != nil {
		return m.Md5
	}
	return nil
}

func (m *ConfigFileRelease) GetVersion() *wrappers.UInt64Value {
	if m != nil {
		return m.Version
	}
	return nil
}

func (m *ConfigFileRelease) GetCreateTime() *wrappers.StringValue {
	if m != nil {
		return m.CreateTime
	}
	return nil
}

func (m *ConfigFileRelease) GetCreateBy() *wrappers.StringValue {
	if m != nil {
		return m.CreateBy
	}
	return nil
}

func (m *ConfigFileRelease) GetModifyTime() *wrappers.StringValue {
	if m != nil {
		return m.ModifyTime
	}
	return nil
}

func (m *ConfigFileRelease) GetModifyBy() *wrappers.StringValue {
	if m != nil {
		return m.ModifyBy
	}
	return nil
}

type ConfigFileReleaseHistory struct {
	Id                   *wrappers.UInt64Value `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 *wrappers.StringValue `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Namespace            *wrappers.StringValue `protobuf:"bytes,3,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Group                *wrappers.StringValue `protobuf:"bytes,4,opt,name=group,proto3" json:"group,omitempty"`
	FileName             *wrappers.StringValue `protobuf:"bytes,5,opt,name=file_name,json=fileName,proto3" json:"file_name,omitempty"`
	Content              *wrappers.StringValue `protobuf:"bytes,6,opt,name=content,proto3" json:"content,omitempty"`
	Format               *wrappers.StringValue `protobuf:"bytes,7,opt,name=format,proto3" json:"format,omitempty"`
	Comment              *wrappers.StringValue `protobuf:"bytes,8,opt,name=comment,proto3" json:"comment,omitempty"`
	Md5                  *wrappers.StringValue `protobuf:"bytes,9,opt,name=md5,proto3" json:"md5,omitempty"`
	Type                 *wrappers.StringValue `protobuf:"bytes,10,opt,name=type,proto3" json:"type,omitempty"`
	Status               *wrappers.StringValue `protobuf:"bytes,11,opt,name=status,proto3" json:"status,omitempty"`
	Tags                 []*ConfigFileTag      `protobuf:"bytes,12,rep,name=tags,proto3" json:"tags,omitempty"`
	CreateTime           *wrappers.StringValue `protobuf:"bytes,13,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	CreateBy             *wrappers.StringValue `protobuf:"bytes,14,opt,name=create_by,json=createBy,proto3" json:"create_by,omitempty"`
	ModifyTime           *wrappers.StringValue `protobuf:"bytes,15,opt,name=modify_time,json=modifyTime,proto3" json:"modify_time,omitempty"`
	ModifyBy             *wrappers.StringValue `protobuf:"bytes,16,opt,name=modify_by,json=modifyBy,proto3" json:"modify_by,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *ConfigFileReleaseHistory) Reset()         { *m = ConfigFileReleaseHistory{} }
func (m *ConfigFileReleaseHistory) String() string { return proto.CompactTextString(m) }
func (*ConfigFileReleaseHistory) ProtoMessage()    {}
func (*ConfigFileReleaseHistory) Descriptor() ([]byte, []int) {
	return fileDescriptor_config_file_092c8b107355d57c, []int{4}
}
func (m *ConfigFileReleaseHistory) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConfigFileReleaseHistory.Unmarshal(m, b)
}
func (m *ConfigFileReleaseHistory) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConfigFileReleaseHistory.Marshal(b, m, deterministic)
}
func (dst *ConfigFileReleaseHistory) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConfigFileReleaseHistory.Merge(dst, src)
}
func (m *ConfigFileReleaseHistory) XXX_Size() int {
	return xxx_messageInfo_ConfigFileReleaseHistory.Size(m)
}
func (m *ConfigFileReleaseHistory) XXX_DiscardUnknown() {
	xxx_messageInfo_ConfigFileReleaseHistory.DiscardUnknown(m)
}

var xxx_messageInfo_ConfigFileReleaseHistory proto.InternalMessageInfo

func (m *ConfigFileReleaseHistory) GetId() *wrappers.UInt64Value {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *ConfigFileReleaseHistory) GetName() *wrappers.StringValue {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *ConfigFileReleaseHistory) GetNamespace() *wrappers.StringValue {
	if m != nil {
		return m.Namespace
	}
	return nil
}

func (m *ConfigFileReleaseHistory) GetGroup() *wrappers.StringValue {
	if m != nil {
		return m.Group
	}
	return nil
}

func (m *ConfigFileReleaseHistory) GetFileName() *wrappers.StringValue {
	if m != nil {
		return m.FileName
	}
	return nil
}

func (m *ConfigFileReleaseHistory) GetContent() *wrappers.StringValue {
	if m != nil {
		return m.Content
	}
	return nil
}

func (m *ConfigFileReleaseHistory) GetFormat() *wrappers.StringValue {
	if m != nil {
		return m.Format
	}
	return nil
}

func (m *ConfigFileReleaseHistory) GetComment() *wrappers.StringValue {
	if m != nil {
		return m.Comment
	}
	return nil
}

func (m *ConfigFileReleaseHistory) GetMd5() *wrappers.StringValue {
	if m != nil {
		return m.Md5
	}
	return nil
}

func (m *ConfigFileReleaseHistory) GetType() *wrappers.StringValue {
	if m != nil {
		return m.Type
	}
	return nil
}

func (m *ConfigFileReleaseHistory) GetStatus() *wrappers.StringValue {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *ConfigFileReleaseHistory) GetTags() []*ConfigFileTag {
	if m != nil {
		return m.Tags
	}
	return nil
}

func (m *ConfigFileReleaseHistory) GetCreateTime() *wrappers.StringValue {
	if m != nil {
		return m.CreateTime
	}
	return nil
}

func (m *ConfigFileReleaseHistory) GetCreateBy() *wrappers.StringValue {
	if m != nil {
		return m.CreateBy
	}
	return nil
}

func (m *ConfigFileReleaseHistory) GetModifyTime() *wrappers.StringValue {
	if m != nil {
		return m.ModifyTime
	}
	return nil
}

func (m *ConfigFileReleaseHistory) GetModifyBy() *wrappers.StringValue {
	if m != nil {
		return m.ModifyBy
	}
	return nil
}

type ClientConfigFileInfo struct {
	Namespace            *wrappers.StringValue `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Group                *wrappers.StringValue `protobuf:"bytes,2,opt,name=group,proto3" json:"group,omitempty"`
	FileName             *wrappers.StringValue `protobuf:"bytes,3,opt,name=file_name,json=fileName,proto3" json:"file_name,omitempty"`
	Content              *wrappers.StringValue `protobuf:"bytes,4,opt,name=content,proto3" json:"content,omitempty"`
	Version              *wrappers.UInt64Value `protobuf:"bytes,5,opt,name=version,proto3" json:"version,omitempty"`
	Md5                  *wrappers.StringValue `protobuf:"bytes,6,opt,name=md5,proto3" json:"md5,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *ClientConfigFileInfo) Reset()         { *m = ClientConfigFileInfo{} }
func (m *ClientConfigFileInfo) String() string { return proto.CompactTextString(m) }
func (*ClientConfigFileInfo) ProtoMessage()    {}
func (*ClientConfigFileInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_config_file_092c8b107355d57c, []int{5}
}
func (m *ClientConfigFileInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClientConfigFileInfo.Unmarshal(m, b)
}
func (m *ClientConfigFileInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClientConfigFileInfo.Marshal(b, m, deterministic)
}
func (dst *ClientConfigFileInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClientConfigFileInfo.Merge(dst, src)
}
func (m *ClientConfigFileInfo) XXX_Size() int {
	return xxx_messageInfo_ClientConfigFileInfo.Size(m)
}
func (m *ClientConfigFileInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ClientConfigFileInfo.DiscardUnknown(m)
}

var xxx_messageInfo_ClientConfigFileInfo proto.InternalMessageInfo

func (m *ClientConfigFileInfo) GetNamespace() *wrappers.StringValue {
	if m != nil {
		return m.Namespace
	}
	return nil
}

func (m *ClientConfigFileInfo) GetGroup() *wrappers.StringValue {
	if m != nil {
		return m.Group
	}
	return nil
}

func (m *ClientConfigFileInfo) GetFileName() *wrappers.StringValue {
	if m != nil {
		return m.FileName
	}
	return nil
}

func (m *ClientConfigFileInfo) GetContent() *wrappers.StringValue {
	if m != nil {
		return m.Content
	}
	return nil
}

func (m *ClientConfigFileInfo) GetVersion() *wrappers.UInt64Value {
	if m != nil {
		return m.Version
	}
	return nil
}

func (m *ClientConfigFileInfo) GetMd5() *wrappers.StringValue {
	if m != nil {
		return m.Md5
	}
	return nil
}

type ClientWatchConfigFileRequest struct {
	ClientIp             *wrappers.StringValue   `protobuf:"bytes,1,opt,name=client_ip,json=clientIp,proto3" json:"client_ip,omitempty"`
	ServiceName          *wrappers.StringValue   `protobuf:"bytes,2,opt,name=service_name,json=serviceName,proto3" json:"service_name,omitempty"`
	WatchFiles           []*ClientConfigFileInfo `protobuf:"bytes,3,rep,name=watch_files,json=watchFiles,proto3" json:"watch_files,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *ClientWatchConfigFileRequest) Reset()         { *m = ClientWatchConfigFileRequest{} }
func (m *ClientWatchConfigFileRequest) String() string { return proto.CompactTextString(m) }
func (*ClientWatchConfigFileRequest) ProtoMessage()    {}
func (*ClientWatchConfigFileRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_config_file_092c8b107355d57c, []int{6}
}
func (m *ClientWatchConfigFileRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClientWatchConfigFileRequest.Unmarshal(m, b)
}
func (m *ClientWatchConfigFileRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClientWatchConfigFileRequest.Marshal(b, m, deterministic)
}
func (dst *ClientWatchConfigFileRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClientWatchConfigFileRequest.Merge(dst, src)
}
func (m *ClientWatchConfigFileRequest) XXX_Size() int {
	return xxx_messageInfo_ClientWatchConfigFileRequest.Size(m)
}
func (m *ClientWatchConfigFileRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ClientWatchConfigFileRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ClientWatchConfigFileRequest proto.InternalMessageInfo

func (m *ClientWatchConfigFileRequest) GetClientIp() *wrappers.StringValue {
	if m != nil {
		return m.ClientIp
	}
	return nil
}

func (m *ClientWatchConfigFileRequest) GetServiceName() *wrappers.StringValue {
	if m != nil {
		return m.ServiceName
	}
	return nil
}

func (m *ClientWatchConfigFileRequest) GetWatchFiles() []*ClientConfigFileInfo {
	if m != nil {
		return m.WatchFiles
	}
	return nil
}

func init() {
	proto.RegisterType((*ConfigFileGroup)(nil), "v1.ConfigFileGroup")
	proto.RegisterType((*ConfigFile)(nil), "v1.ConfigFile")
	proto.RegisterType((*ConfigFileTag)(nil), "v1.ConfigFileTag")
	proto.RegisterType((*ConfigFileRelease)(nil), "v1.ConfigFileRelease")
	proto.RegisterType((*ConfigFileReleaseHistory)(nil), "v1.ConfigFileReleaseHistory")
	proto.RegisterType((*ClientConfigFileInfo)(nil), "v1.ClientConfigFileInfo")
	proto.RegisterType((*ClientWatchConfigFileRequest)(nil), "v1.ClientWatchConfigFileRequest")
}

func init() { proto.RegisterFile("config_file.proto", fileDescriptor_config_file_092c8b107355d57c) }

var fileDescriptor_config_file_092c8b107355d57c = []byte{
	// 760 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xec, 0x97, 0x4d, 0x6f, 0xd3, 0x4c,
	0x10, 0xc7, 0x95, 0x37, 0x27, 0x19, 0x37, 0x4d, 0xbb, 0x7a, 0x0e, 0xab, 0xaa, 0x7a, 0x54, 0x45,
	0x42, 0xe2, 0x80, 0xdc, 0x17, 0x4a, 0x45, 0x41, 0x08, 0x29, 0x95, 0x4a, 0x7b, 0xe1, 0x10, 0x0a,
	0x48, 0x5c, 0x22, 0x27, 0xd9, 0x84, 0x15, 0xb6, 0xd7, 0x78, 0x37, 0xae, 0xfc, 0x31, 0x39, 0xf4,
	0x03, 0x70, 0xe5, 0x43, 0x20, 0xb4, 0xbb, 0x36, 0x4e, 0x5b, 0xaa, 0xae, 0xed, 0x0b, 0x02, 0x4e,
	0x95, 0xed, 0xff, 0x7f, 0x66, 0x3d, 0x9e, 0xdf, 0x4c, 0x03, 0x9b, 0x53, 0x16, 0xcc, 0xe9, 0x62,
	0x3c, 0xa7, 0x1e, 0x71, 0xc2, 0x88, 0x09, 0x86, 0xea, 0xf1, 0xfe, 0xd6, 0xff, 0x0b, 0xc6, 0x16,
	0x1e, 0xd9, 0x55, 0x77, 0x26, 0xcb, 0xf9, 0xee, 0x65, 0xe4, 0x86, 0x21, 0x89, 0xb8, 0xd6, 0x0c,
	0xae, 0x2c, 0xe8, 0x9f, 0x28, 0xe7, 0x29, 0xf5, 0xc8, 0xab, 0x88, 0x2d, 0x43, 0xf4, 0x08, 0xea,
	0x74, 0x86, 0x6b, 0x3b, 0xb5, 0x87, 0xf6, 0xc1, 0xb6, 0xa3, 0x03, 0x38, 0x59, 0x00, 0xe7, 0xed,
	0x79, 0x20, 0x8e, 0x0e, 0xdf, 0xb9, 0xde, 0x92, 0x8c, 0xea, 0x74, 0x86, 0xf6, 0xa0, 0x19, 0xb8,
	0x3e, 0xc1, 0xf5, 0x3b, 0xf4, 0x6f, 0x44, 0x44, 0x83, 0x85, 0xd6, 0x2b, 0x25, 0x7a, 0x06, 0x5d,
	0xf9, 0x97, 0x87, 0xee, 0x94, 0xe0, 0x86, 0x81, 0x2d, 0x97, 0xa3, 0x23, 0x68, 0x4f, 0x99, 0xef,
	0x93, 0x40, 0xe0, 0xa6, 0x81, 0x33, 0x13, 0xa3, 0x17, 0x60, 0x4f, 0x23, 0xe2, 0x0a, 0x32, 0x16,
	0xd4, 0x27, 0xb8, 0x65, 0xe0, 0x05, 0x6d, 0xb8, 0xa0, 0x3e, 0x41, 0xc7, 0xd0, 0x4d, 0xed, 0x93,
	0x04, 0x5b, 0x06, 0xe6, 0x8e, 0x96, 0x0f, 0x13, 0x99, 0xd9, 0x67, 0x33, 0x3a, 0x4f, 0x74, 0xe6,
	0xb6, 0x49, 0x66, 0x6d, 0xc8, 0x32, 0xa7, 0xf6, 0x49, 0x82, 0x3b, 0x26, 0x99, 0xb5, 0x7c, 0x98,
	0xc8, 0x3a, 0xcb, 0x6e, 0x38, 0x61, 0xcb, 0x40, 0xe0, 0xae, 0xc1, 0xe7, 0xcc, 0xe5, 0xe8, 0x29,
	0x74, 0x96, 0x9c, 0x44, 0x63, 0x3a, 0xe3, 0x18, 0x76, 0x1a, 0xf7, 0x67, 0xcd, 0xd4, 0x32, 0xeb,
	0x42, 0xb6, 0x91, 0xb2, 0xda, 0x06, 0xd6, 0x5c, 0x8e, 0x4e, 0xa1, 0x1f, 0x11, 0x9f, 0xc5, 0x64,
	0xfc, 0x33, 0x79, 0xcf, 0x20, 0xc2, 0x4d, 0x13, 0x3a, 0x83, 0x8d, 0xf4, 0x56, 0x7e, 0x94, 0x75,
	0x83, 0x40, 0xb7, 0x5c, 0xe8, 0x08, 0x3a, 0x64, 0x46, 0x85, 0x3b, 0xf1, 0x08, 0xee, 0xab, 0x12,
	0x6e, 0xdd, 0x8a, 0x30, 0x64, 0xcc, 0x4b, 0xab, 0x90, 0x69, 0x07, 0x5f, 0x2c, 0x80, 0x9c, 0xab,
	0xdf, 0x1a, 0xa9, 0x03, 0x68, 0xa9, 0xf7, 0x35, 0x02, 0x4a, 0x4b, 0x35, 0x86, 0x81, 0x90, 0x18,
	0xb6, 0xcc, 0x30, 0x54, 0x62, 0x74, 0x08, 0xd6, 0x9c, 0x45, 0xbe, 0x2b, 0x8c, 0x20, 0x4a, 0xb5,
	0xab, 0xd0, 0xb7, 0x8b, 0x40, 0x7f, 0x08, 0x16, 0x17, 0xae, 0x58, 0x72, 0x23, 0x70, 0x52, 0x2d,
	0x7a, 0x00, 0x4d, 0xe1, 0x2e, 0x38, 0xee, 0xaa, 0x86, 0xd9, 0x74, 0xe2, 0x7d, 0x27, 0xff, 0x92,
	0x17, 0xee, 0x62, 0xa4, 0x1e, 0xdf, 0x9c, 0x28, 0x50, 0x65, 0xa2, 0xd8, 0x55, 0x26, 0xca, 0x5a,
	0x95, 0x89, 0xd2, 0x2b, 0x34, 0x51, 0x5e, 0xc2, 0x5a, 0x44, 0x3c, 0xe2, 0xf2, 0xf4, 0xa5, 0xd7,
	0x0d, 0xdc, 0x76, 0xea, 0x50, 0xb9, 0x9f, 0x03, 0x64, 0x01, 0x26, 0x49, 0x0a, 0xd4, 0x3d, 0x8d,
	0x9a, 0xea, 0x87, 0xc9, 0x80, 0x43, 0xef, 0xda, 0x87, 0x40, 0x0e, 0x34, 0x3e, 0x91, 0xe4, 0x4e,
	0xac, 0x56, 0xc3, 0x48, 0xa1, 0xec, 0xf4, 0x58, 0x5e, 0x19, 0x81, 0xa5, 0xa5, 0x83, 0x6f, 0x2d,
	0xd8, 0xcc, 0xb3, 0x8e, 0xf4, 0x61, 0xfe, 0x38, 0x9e, 0x8f, 0xf5, 0xaa, 0x18, 0xab, 0x63, 0x9a,
	0x10, 0xdd, 0x91, 0xf2, 0xd7, 0xf2, 0xa8, 0x2b, 0xa3, 0xc0, 0x2a, 0x32, 0x0a, 0xca, 0x42, 0xed,
	0x40, 0xc3, 0x9f, 0x3d, 0x31, 0x22, 0x5a, 0x0a, 0x65, 0x9e, 0x98, 0x44, 0x9c, 0xb2, 0xc0, 0x68,
	0x07, 0x66, 0xe2, 0xbf, 0x91, 0xef, 0xc1, 0x77, 0x0b, 0xf0, 0xad, 0x66, 0x3f, 0xa3, 0x5c, 0xb0,
	0x28, 0xf9, 0xd7, 0xf3, 0xd5, 0x7b, 0x3e, 0x5f, 0x7f, 0xed, 0x72, 0xeb, 0xaf, 0x53, 0x82, 0x94,
	0xae, 0x29, 0x29, 0x7b, 0xd0, 0x14, 0x49, 0x68, 0xd6, 0xea, 0x4a, 0xb9, 0xb2, 0x60, 0xed, 0x12,
	0x0b, 0x76, 0xad, 0xd0, 0x82, 0xed, 0x55, 0x01, 0x70, 0xbd, 0x0a, 0x80, 0xfd, 0x2a, 0x00, 0x6e,
	0x14, 0x02, 0xf0, 0x6b, 0x1d, 0xfe, 0x3b, 0xf1, 0x28, 0x09, 0x44, 0x5e, 0x91, 0xf3, 0x60, 0xce,
	0xae, 0xc3, 0x51, 0x2b, 0x09, 0x47, 0xbd, 0x24, 0x1c, 0x8d, 0xb2, 0x70, 0x34, 0x0b, 0x2e, 0x84,
	0x6c, 0x50, 0xb7, 0x8a, 0x0c, 0xea, 0xb4, 0xcd, 0x2d, 0xc3, 0x36, 0x1f, 0x5c, 0xd5, 0x60, 0x5b,
	0xd7, 0xf8, 0xbd, 0x2b, 0xa6, 0x1f, 0x57, 0xe7, 0xdd, 0xe7, 0x25, 0xe1, 0x42, 0x75, 0x8e, 0x7a,
	0x3e, 0xa6, 0xa1, 0x51, 0xad, 0x3b, 0x5a, 0x7e, 0x1e, 0xca, 0x7f, 0x90, 0x38, 0x89, 0x62, 0x3a,
	0x4d, 0x2b, 0x67, 0x52, 0x71, 0x3b, 0x75, 0xa8, 0xe2, 0x1d, 0x83, 0x7d, 0x29, 0x4f, 0xa5, 0x7e,
	0xc7, 0x73, 0xdc, 0x50, 0x88, 0x60, 0x85, 0xc8, 0x2f, 0xda, 0x62, 0x04, 0x4a, 0x2c, 0x2f, 0xf9,
	0xb0, 0xf9, 0xa1, 0x1e, 0xef, 0x4f, 0x2c, 0x95, 0xe3, 0xf1, 0x8f, 0x00, 0x00, 0x00, 0xff, 0xff,
	0x79, 0xcd, 0xba, 0x9a, 0x10, 0x10, 0x00, 0x00,
}
