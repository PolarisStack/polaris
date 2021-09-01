// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/service/accesslog/v3/als.proto

package envoy_service_accesslog_v3

import (
	context "context"
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	v3 "github.com/polarismesh/polaris-server/apiserver/xdsserver/go-control-plane/envoy/config/core/v3"
	v31 "github.com/polarismesh/polaris-server/apiserver/xdsserver/go-control-plane/envoy/data/accesslog/v3"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type StreamAccessLogsResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StreamAccessLogsResponse) Reset()         { *m = StreamAccessLogsResponse{} }
func (m *StreamAccessLogsResponse) String() string { return proto.CompactTextString(m) }
func (*StreamAccessLogsResponse) ProtoMessage()    {}
func (*StreamAccessLogsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7df9b8b42f8221f5, []int{0}
}

func (m *StreamAccessLogsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StreamAccessLogsResponse.Unmarshal(m, b)
}
func (m *StreamAccessLogsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StreamAccessLogsResponse.Marshal(b, m, deterministic)
}
func (m *StreamAccessLogsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StreamAccessLogsResponse.Merge(m, src)
}
func (m *StreamAccessLogsResponse) XXX_Size() int {
	return xxx_messageInfo_StreamAccessLogsResponse.Size(m)
}
func (m *StreamAccessLogsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_StreamAccessLogsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_StreamAccessLogsResponse proto.InternalMessageInfo

type StreamAccessLogsMessage struct {
	Identifier *StreamAccessLogsMessage_Identifier `protobuf:"bytes,1,opt,name=identifier,proto3" json:"identifier,omitempty"`
	// Types that are valid to be assigned to LogEntries:
	//	*StreamAccessLogsMessage_HttpLogs
	//	*StreamAccessLogsMessage_TcpLogs
	LogEntries           isStreamAccessLogsMessage_LogEntries `protobuf_oneof:"log_entries"`
	XXX_NoUnkeyedLiteral struct{}                             `json:"-"`
	XXX_unrecognized     []byte                               `json:"-"`
	XXX_sizecache        int32                                `json:"-"`
}

func (m *StreamAccessLogsMessage) Reset()         { *m = StreamAccessLogsMessage{} }
func (m *StreamAccessLogsMessage) String() string { return proto.CompactTextString(m) }
func (*StreamAccessLogsMessage) ProtoMessage()    {}
func (*StreamAccessLogsMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_7df9b8b42f8221f5, []int{1}
}

func (m *StreamAccessLogsMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StreamAccessLogsMessage.Unmarshal(m, b)
}
func (m *StreamAccessLogsMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StreamAccessLogsMessage.Marshal(b, m, deterministic)
}
func (m *StreamAccessLogsMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StreamAccessLogsMessage.Merge(m, src)
}
func (m *StreamAccessLogsMessage) XXX_Size() int {
	return xxx_messageInfo_StreamAccessLogsMessage.Size(m)
}
func (m *StreamAccessLogsMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_StreamAccessLogsMessage.DiscardUnknown(m)
}

var xxx_messageInfo_StreamAccessLogsMessage proto.InternalMessageInfo

func (m *StreamAccessLogsMessage) GetIdentifier() *StreamAccessLogsMessage_Identifier {
	if m != nil {
		return m.Identifier
	}
	return nil
}

type isStreamAccessLogsMessage_LogEntries interface {
	isStreamAccessLogsMessage_LogEntries()
}

type StreamAccessLogsMessage_HttpLogs struct {
	HttpLogs *StreamAccessLogsMessage_HTTPAccessLogEntries `protobuf:"bytes,2,opt,name=http_logs,json=httpLogs,proto3,oneof"`
}

type StreamAccessLogsMessage_TcpLogs struct {
	TcpLogs *StreamAccessLogsMessage_TCPAccessLogEntries `protobuf:"bytes,3,opt,name=tcp_logs,json=tcpLogs,proto3,oneof"`
}

func (*StreamAccessLogsMessage_HttpLogs) isStreamAccessLogsMessage_LogEntries() {}

func (*StreamAccessLogsMessage_TcpLogs) isStreamAccessLogsMessage_LogEntries() {}

func (m *StreamAccessLogsMessage) GetLogEntries() isStreamAccessLogsMessage_LogEntries {
	if m != nil {
		return m.LogEntries
	}
	return nil
}

func (m *StreamAccessLogsMessage) GetHttpLogs() *StreamAccessLogsMessage_HTTPAccessLogEntries {
	if x, ok := m.GetLogEntries().(*StreamAccessLogsMessage_HttpLogs); ok {
		return x.HttpLogs
	}
	return nil
}

func (m *StreamAccessLogsMessage) GetTcpLogs() *StreamAccessLogsMessage_TCPAccessLogEntries {
	if x, ok := m.GetLogEntries().(*StreamAccessLogsMessage_TcpLogs); ok {
		return x.TcpLogs
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*StreamAccessLogsMessage) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*StreamAccessLogsMessage_HttpLogs)(nil),
		(*StreamAccessLogsMessage_TcpLogs)(nil),
	}
}

type StreamAccessLogsMessage_Identifier struct {
	Node                 *v3.Node `protobuf:"bytes,1,opt,name=node,proto3" json:"node,omitempty"`
	LogName              string   `protobuf:"bytes,2,opt,name=log_name,json=logName,proto3" json:"log_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StreamAccessLogsMessage_Identifier) Reset()         { *m = StreamAccessLogsMessage_Identifier{} }
func (m *StreamAccessLogsMessage_Identifier) String() string { return proto.CompactTextString(m) }
func (*StreamAccessLogsMessage_Identifier) ProtoMessage()    {}
func (*StreamAccessLogsMessage_Identifier) Descriptor() ([]byte, []int) {
	return fileDescriptor_7df9b8b42f8221f5, []int{1, 0}
}

func (m *StreamAccessLogsMessage_Identifier) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StreamAccessLogsMessage_Identifier.Unmarshal(m, b)
}
func (m *StreamAccessLogsMessage_Identifier) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StreamAccessLogsMessage_Identifier.Marshal(b, m, deterministic)
}
func (m *StreamAccessLogsMessage_Identifier) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StreamAccessLogsMessage_Identifier.Merge(m, src)
}
func (m *StreamAccessLogsMessage_Identifier) XXX_Size() int {
	return xxx_messageInfo_StreamAccessLogsMessage_Identifier.Size(m)
}
func (m *StreamAccessLogsMessage_Identifier) XXX_DiscardUnknown() {
	xxx_messageInfo_StreamAccessLogsMessage_Identifier.DiscardUnknown(m)
}

var xxx_messageInfo_StreamAccessLogsMessage_Identifier proto.InternalMessageInfo

func (m *StreamAccessLogsMessage_Identifier) GetNode() *v3.Node {
	if m != nil {
		return m.Node
	}
	return nil
}

func (m *StreamAccessLogsMessage_Identifier) GetLogName() string {
	if m != nil {
		return m.LogName
	}
	return ""
}

type StreamAccessLogsMessage_HTTPAccessLogEntries struct {
	LogEntry             []*v31.HTTPAccessLogEntry `protobuf:"bytes,1,rep,name=log_entry,json=logEntry,proto3" json:"log_entry,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *StreamAccessLogsMessage_HTTPAccessLogEntries) Reset() {
	*m = StreamAccessLogsMessage_HTTPAccessLogEntries{}
}
func (m *StreamAccessLogsMessage_HTTPAccessLogEntries) String() string {
	return proto.CompactTextString(m)
}
func (*StreamAccessLogsMessage_HTTPAccessLogEntries) ProtoMessage() {}
func (*StreamAccessLogsMessage_HTTPAccessLogEntries) Descriptor() ([]byte, []int) {
	return fileDescriptor_7df9b8b42f8221f5, []int{1, 1}
}

func (m *StreamAccessLogsMessage_HTTPAccessLogEntries) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StreamAccessLogsMessage_HTTPAccessLogEntries.Unmarshal(m, b)
}
func (m *StreamAccessLogsMessage_HTTPAccessLogEntries) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StreamAccessLogsMessage_HTTPAccessLogEntries.Marshal(b, m, deterministic)
}
func (m *StreamAccessLogsMessage_HTTPAccessLogEntries) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StreamAccessLogsMessage_HTTPAccessLogEntries.Merge(m, src)
}
func (m *StreamAccessLogsMessage_HTTPAccessLogEntries) XXX_Size() int {
	return xxx_messageInfo_StreamAccessLogsMessage_HTTPAccessLogEntries.Size(m)
}
func (m *StreamAccessLogsMessage_HTTPAccessLogEntries) XXX_DiscardUnknown() {
	xxx_messageInfo_StreamAccessLogsMessage_HTTPAccessLogEntries.DiscardUnknown(m)
}

var xxx_messageInfo_StreamAccessLogsMessage_HTTPAccessLogEntries proto.InternalMessageInfo

func (m *StreamAccessLogsMessage_HTTPAccessLogEntries) GetLogEntry() []*v31.HTTPAccessLogEntry {
	if m != nil {
		return m.LogEntry
	}
	return nil
}

type StreamAccessLogsMessage_TCPAccessLogEntries struct {
	LogEntry             []*v31.TCPAccessLogEntry `protobuf:"bytes,1,rep,name=log_entry,json=logEntry,proto3" json:"log_entry,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *StreamAccessLogsMessage_TCPAccessLogEntries) Reset() {
	*m = StreamAccessLogsMessage_TCPAccessLogEntries{}
}
func (m *StreamAccessLogsMessage_TCPAccessLogEntries) String() string {
	return proto.CompactTextString(m)
}
func (*StreamAccessLogsMessage_TCPAccessLogEntries) ProtoMessage() {}
func (*StreamAccessLogsMessage_TCPAccessLogEntries) Descriptor() ([]byte, []int) {
	return fileDescriptor_7df9b8b42f8221f5, []int{1, 2}
}

func (m *StreamAccessLogsMessage_TCPAccessLogEntries) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StreamAccessLogsMessage_TCPAccessLogEntries.Unmarshal(m, b)
}
func (m *StreamAccessLogsMessage_TCPAccessLogEntries) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StreamAccessLogsMessage_TCPAccessLogEntries.Marshal(b, m, deterministic)
}
func (m *StreamAccessLogsMessage_TCPAccessLogEntries) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StreamAccessLogsMessage_TCPAccessLogEntries.Merge(m, src)
}
func (m *StreamAccessLogsMessage_TCPAccessLogEntries) XXX_Size() int {
	return xxx_messageInfo_StreamAccessLogsMessage_TCPAccessLogEntries.Size(m)
}
func (m *StreamAccessLogsMessage_TCPAccessLogEntries) XXX_DiscardUnknown() {
	xxx_messageInfo_StreamAccessLogsMessage_TCPAccessLogEntries.DiscardUnknown(m)
}

var xxx_messageInfo_StreamAccessLogsMessage_TCPAccessLogEntries proto.InternalMessageInfo

func (m *StreamAccessLogsMessage_TCPAccessLogEntries) GetLogEntry() []*v31.TCPAccessLogEntry {
	if m != nil {
		return m.LogEntry
	}
	return nil
}

func init() {
	proto.RegisterType((*StreamAccessLogsResponse)(nil), "envoy.service.accesslog.v3.StreamAccessLogsResponse")
	proto.RegisterType((*StreamAccessLogsMessage)(nil), "envoy.service.accesslog.v3.StreamAccessLogsMessage")
	proto.RegisterType((*StreamAccessLogsMessage_Identifier)(nil), "envoy.service.accesslog.v3.StreamAccessLogsMessage.Identifier")
	proto.RegisterType((*StreamAccessLogsMessage_HTTPAccessLogEntries)(nil), "envoy.service.accesslog.v3.StreamAccessLogsMessage.HTTPAccessLogEntries")
	proto.RegisterType((*StreamAccessLogsMessage_TCPAccessLogEntries)(nil), "envoy.service.accesslog.v3.StreamAccessLogsMessage.TCPAccessLogEntries")
}

func init() {
	proto.RegisterFile("envoy/service/accesslog/v3/als.proto", fileDescriptor_7df9b8b42f8221f5)
}

var fileDescriptor_7df9b8b42f8221f5 = []byte{
	// 578 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x94, 0xcf, 0x8b, 0xd3, 0x4e,
	0x18, 0xc6, 0x77, 0xb2, 0xfb, 0xdd, 0x66, 0xdf, 0x5e, 0xca, 0x7c, 0x85, 0x2d, 0x01, 0xb5, 0x16,
	0xc1, 0xa2, 0x90, 0x40, 0xbb, 0x82, 0x16, 0x54, 0x1a, 0x7f, 0x6c, 0x17, 0xdc, 0x52, 0xb3, 0xc5,
	0xa3, 0xcb, 0x6c, 0x32, 0x1b, 0x07, 0xd2, 0x99, 0x92, 0x99, 0x0d, 0xf6, 0xa6, 0xb7, 0xc5, 0xa3,
	0x07, 0x0f, 0xfe, 0x15, 0xde, 0x84, 0xbd, 0x79, 0x10, 0xbc, 0xfa, 0xdf, 0xc8, 0x9e, 0x64, 0x92,
	0x34, 0x6a, 0x9b, 0x0a, 0xed, 0xad, 0x64, 0xde, 0xf7, 0xf3, 0x3c, 0xcf, 0xfb, 0x4e, 0x07, 0x6e,
	0x52, 0x9e, 0x88, 0xa9, 0x23, 0x69, 0x9c, 0x30, 0x9f, 0x3a, 0xc4, 0xf7, 0xa9, 0x94, 0x91, 0x08,
	0x9d, 0xa4, 0xe3, 0x90, 0x48, 0xda, 0x93, 0x58, 0x28, 0x81, 0xad, 0xb4, 0xca, 0xce, 0xab, 0xec,
	0xa2, 0xca, 0x4e, 0x3a, 0xd6, 0xf5, 0x8c, 0xe0, 0x0b, 0x7e, 0xca, 0x42, 0xc7, 0x17, 0x31, 0xd5,
	0xbd, 0x27, 0x44, 0xd2, 0xac, 0xd9, 0xba, 0x95, 0x15, 0x04, 0x44, 0x91, 0x39, 0x7e, 0x81, 0xc9,
	0x0a, 0xaf, 0x9e, 0x05, 0x13, 0xe2, 0x10, 0xce, 0x85, 0x22, 0x8a, 0x09, 0x2e, 0x1d, 0xa9, 0x88,
	0x3a, 0xcb, 0x4d, 0x58, 0x37, 0x16, 0x8e, 0x13, 0x1a, 0x4b, 0x26, 0x38, 0xe3, 0x33, 0xc2, 0x6e,
	0x42, 0x22, 0x16, 0x10, 0x45, 0x9d, 0xd9, 0x8f, 0xec, 0xa0, 0xf9, 0x12, 0xea, 0x47, 0x2a, 0xa6,
	0x64, 0xdc, 0x4b, 0x35, 0x9f, 0x8b, 0x50, 0x7a, 0x54, 0x4e, 0x04, 0x97, 0xb4, 0xdb, 0xfd, 0xf4,
	0xed, 0xfc, 0xda, 0x5d, 0xe8, 0x2c, 0xcd, 0xd8, 0xb6, 0x97, 0xf5, 0x36, 0xbf, 0x56, 0x60, 0x77,
	0xfe, 0xf0, 0x90, 0x4a, 0x49, 0x42, 0x8a, 0x5f, 0x01, 0xb0, 0x80, 0x72, 0xc5, 0x4e, 0x19, 0x8d,
	0xeb, 0xa8, 0x81, 0x5a, 0xd5, 0xf6, 0x43, 0x7b, 0xf9, 0x24, 0xed, 0x25, 0x20, 0xfb, 0xa0, 0xa0,
	0x78, 0x7f, 0x10, 0x71, 0x08, 0x3b, 0xaf, 0x95, 0x9a, 0x1c, 0x47, 0x22, 0x94, 0x75, 0x23, 0xc5,
	0xf7, 0xd7, 0xc1, 0xf7, 0x47, 0xa3, 0x61, 0xf1, 0xf5, 0x29, 0x57, 0x31, 0xa3, 0xb2, 0xbf, 0xe1,
	0x99, 0x1a, 0xae, 0xeb, 0x70, 0x00, 0xa6, 0xf2, 0x73, 0x9d, 0xcd, 0x54, 0x67, 0x7f, 0x1d, 0x9d,
	0xd1, 0xe3, 0x32, 0x99, 0x8a, 0xf2, 0x53, 0x15, 0xeb, 0x33, 0x02, 0xf8, 0x9d, 0x14, 0xdf, 0x83,
	0x2d, 0x2e, 0x02, 0x9a, 0xcf, 0xcd, 0xca, 0x05, 0xb3, 0x5b, 0x66, 0xeb, 0x5b, 0xa6, 0xa5, 0x06,
	0x22, 0xa0, 0xae, 0x79, 0xe9, 0xfe, 0xf7, 0x1e, 0x19, 0x35, 0xe4, 0xa5, 0x1d, 0xb8, 0x09, 0x66,
	0x24, 0xc2, 0x63, 0x4e, 0xc6, 0x34, 0x1d, 0xcb, 0x8e, 0x5b, 0xb9, 0x74, 0xb7, 0x62, 0xa3, 0x81,
	0xbc, 0x4a, 0x24, 0xc2, 0x01, 0x19, 0xd3, 0xee, 0x13, 0xbd, 0xf3, 0x47, 0xf0, 0x60, 0x85, 0x9d,
	0x2f, 0x6e, 0xc3, 0xba, 0x40, 0x70, 0xa5, 0x6c, 0x7a, 0xd8, 0x83, 0x1d, 0x6d, 0x81, 0x72, 0x15,
	0x4f, 0xeb, 0xa8, 0xb1, 0xd9, 0xaa, 0xb6, 0xef, 0xe4, 0x09, 0xf4, 0xdf, 0xe0, 0xef, 0x79, 0x2d,
	0x10, 0xa6, 0x69, 0xa4, 0x0f, 0xc8, 0x30, 0x91, 0xa7, 0xa3, 0xa4, 0xdf, 0xba, 0x03, 0x6d, 0xf9,
	0x00, 0xf6, 0xd7, 0xb0, 0x5c, 0xe6, 0xd1, 0xfa, 0x82, 0xe0, 0xff, 0x92, 0x95, 0xe0, 0x17, 0x8b,
	0xde, 0x6f, 0x2f, 0xf5, 0x3e, 0x0f, 0x28, 0xb7, 0x7e, 0xa8, 0xad, 0xf7, 0xe1, 0xd9, 0x1a, 0xd6,
	0x4b, 0x1c, 0x76, 0xef, 0x6b, 0xdc, 0x1e, 0xb4, 0x57, 0xc7, 0xb9, 0x18, 0xaa, 0xb3, 0x70, 0x3a,
	0xeb, 0xe6, 0x4f, 0x17, 0xb5, 0x3f, 0x22, 0xa8, 0x15, 0x95, 0x47, 0x19, 0x0d, 0xbf, 0x43, 0x50,
	0x9b, 0x87, 0xe0, 0xce, 0x1a, 0xd7, 0xde, 0xda, 0x5b, 0xa5, 0xa9, 0x78, 0x58, 0x36, 0x5a, 0xc8,
	0xed, 0x5d, 0xbc, 0xfd, 0xfe, 0x63, 0xdb, 0xa8, 0x19, 0xd0, 0x62, 0x22, 0xa3, 0x4c, 0x62, 0xf1,
	0x66, 0xfa, 0x0f, 0xa0, 0x6b, 0xf6, 0x22, 0x39, 0xd4, 0x4f, 0xde, 0x10, 0x9d, 0x23, 0x74, 0xb2,
	0x9d, 0x3e, 0x7f, 0x9d, 0x5f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x13, 0x97, 0x10, 0x5b, 0xe7, 0x05,
	0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AccessLogServiceClient is the client API for AccessLogService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AccessLogServiceClient interface {
	StreamAccessLogs(ctx context.Context, opts ...grpc.CallOption) (AccessLogService_StreamAccessLogsClient, error)
}

type accessLogServiceClient struct {
	cc *grpc.ClientConn
}

func NewAccessLogServiceClient(cc *grpc.ClientConn) AccessLogServiceClient {
	return &accessLogServiceClient{cc}
}

func (c *accessLogServiceClient) StreamAccessLogs(ctx context.Context, opts ...grpc.CallOption) (AccessLogService_StreamAccessLogsClient, error) {
	stream, err := c.cc.NewStream(ctx, &_AccessLogService_serviceDesc.Streams[0], "/envoy.service.accesslog.v3.AccessLogService/StreamAccessLogs", opts...)
	if err != nil {
		return nil, err
	}
	x := &accessLogServiceStreamAccessLogsClient{stream}
	return x, nil
}

type AccessLogService_StreamAccessLogsClient interface {
	Send(*StreamAccessLogsMessage) error
	CloseAndRecv() (*StreamAccessLogsResponse, error)
	grpc.ClientStream
}

type accessLogServiceStreamAccessLogsClient struct {
	grpc.ClientStream
}

func (x *accessLogServiceStreamAccessLogsClient) Send(m *StreamAccessLogsMessage) error {
	return x.ClientStream.SendMsg(m)
}

func (x *accessLogServiceStreamAccessLogsClient) CloseAndRecv() (*StreamAccessLogsResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(StreamAccessLogsResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// AccessLogServiceServer is the server API for AccessLogService service.
type AccessLogServiceServer interface {
	StreamAccessLogs(AccessLogService_StreamAccessLogsServer) error
}

// UnimplementedAccessLogServiceServer can be embedded to have forward compatible implementations.
type UnimplementedAccessLogServiceServer struct {
}

func (*UnimplementedAccessLogServiceServer) StreamAccessLogs(srv AccessLogService_StreamAccessLogsServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamAccessLogs not implemented")
}

func RegisterAccessLogServiceServer(s *grpc.Server, srv AccessLogServiceServer) {
	s.RegisterService(&_AccessLogService_serviceDesc, srv)
}

func _AccessLogService_StreamAccessLogs_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(AccessLogServiceServer).StreamAccessLogs(&accessLogServiceStreamAccessLogsServer{stream})
}

type AccessLogService_StreamAccessLogsServer interface {
	SendAndClose(*StreamAccessLogsResponse) error
	Recv() (*StreamAccessLogsMessage, error)
	grpc.ServerStream
}

type accessLogServiceStreamAccessLogsServer struct {
	grpc.ServerStream
}

func (x *accessLogServiceStreamAccessLogsServer) SendAndClose(m *StreamAccessLogsResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *accessLogServiceStreamAccessLogsServer) Recv() (*StreamAccessLogsMessage, error) {
	m := new(StreamAccessLogsMessage)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _AccessLogService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "envoy.service.accesslog.v3.AccessLogService",
	HandlerType: (*AccessLogServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamAccessLogs",
			Handler:       _AccessLogService_StreamAccessLogs_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "envoy/service/accesslog/v3/als.proto",
}
