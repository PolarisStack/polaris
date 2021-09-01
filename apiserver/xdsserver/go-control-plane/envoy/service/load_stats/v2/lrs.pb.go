// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/service/load_stats/v2/lrs.proto

package envoy_service_load_stats_v2

import (
	context "context"
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	core "github.com/polarismesh/polaris-server/apiserver/xdsserver/go-control-plane/envoy/api/v2/core"
	endpoint "github.com/polarismesh/polaris-server/apiserver/xdsserver/go-control-plane/envoy/api/v2/endpoint"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	duration "github.com/golang/protobuf/ptypes/duration"
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

type LoadStatsRequest struct {
	Node                 *core.Node               `protobuf:"bytes,1,opt,name=node,proto3" json:"node,omitempty"`
	ClusterStats         []*endpoint.ClusterStats `protobuf:"bytes,2,rep,name=cluster_stats,json=clusterStats,proto3" json:"cluster_stats,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *LoadStatsRequest) Reset()         { *m = LoadStatsRequest{} }
func (m *LoadStatsRequest) String() string { return proto.CompactTextString(m) }
func (*LoadStatsRequest) ProtoMessage()    {}
func (*LoadStatsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_cd7332d279836518, []int{0}
}

func (m *LoadStatsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoadStatsRequest.Unmarshal(m, b)
}
func (m *LoadStatsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoadStatsRequest.Marshal(b, m, deterministic)
}
func (m *LoadStatsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoadStatsRequest.Merge(m, src)
}
func (m *LoadStatsRequest) XXX_Size() int {
	return xxx_messageInfo_LoadStatsRequest.Size(m)
}
func (m *LoadStatsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_LoadStatsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_LoadStatsRequest proto.InternalMessageInfo

func (m *LoadStatsRequest) GetNode() *core.Node {
	if m != nil {
		return m.Node
	}
	return nil
}

func (m *LoadStatsRequest) GetClusterStats() []*endpoint.ClusterStats {
	if m != nil {
		return m.ClusterStats
	}
	return nil
}

type LoadStatsResponse struct {
	Clusters                  []string           `protobuf:"bytes,1,rep,name=clusters,proto3" json:"clusters,omitempty"`
	LoadReportingInterval     *duration.Duration `protobuf:"bytes,2,opt,name=load_reporting_interval,json=loadReportingInterval,proto3" json:"load_reporting_interval,omitempty"`
	ReportEndpointGranularity bool               `protobuf:"varint,3,opt,name=report_endpoint_granularity,json=reportEndpointGranularity,proto3" json:"report_endpoint_granularity,omitempty"`
	XXX_NoUnkeyedLiteral      struct{}           `json:"-"`
	XXX_unrecognized          []byte             `json:"-"`
	XXX_sizecache             int32              `json:"-"`
}

func (m *LoadStatsResponse) Reset()         { *m = LoadStatsResponse{} }
func (m *LoadStatsResponse) String() string { return proto.CompactTextString(m) }
func (*LoadStatsResponse) ProtoMessage()    {}
func (*LoadStatsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_cd7332d279836518, []int{1}
}

func (m *LoadStatsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoadStatsResponse.Unmarshal(m, b)
}
func (m *LoadStatsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoadStatsResponse.Marshal(b, m, deterministic)
}
func (m *LoadStatsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoadStatsResponse.Merge(m, src)
}
func (m *LoadStatsResponse) XXX_Size() int {
	return xxx_messageInfo_LoadStatsResponse.Size(m)
}
func (m *LoadStatsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_LoadStatsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_LoadStatsResponse proto.InternalMessageInfo

func (m *LoadStatsResponse) GetClusters() []string {
	if m != nil {
		return m.Clusters
	}
	return nil
}

func (m *LoadStatsResponse) GetLoadReportingInterval() *duration.Duration {
	if m != nil {
		return m.LoadReportingInterval
	}
	return nil
}

func (m *LoadStatsResponse) GetReportEndpointGranularity() bool {
	if m != nil {
		return m.ReportEndpointGranularity
	}
	return false
}

func init() {
	proto.RegisterType((*LoadStatsRequest)(nil), "envoy.service.load_stats.v2.LoadStatsRequest")
	proto.RegisterType((*LoadStatsResponse)(nil), "envoy.service.load_stats.v2.LoadStatsResponse")
}

func init() {
	proto.RegisterFile("envoy/service/load_stats/v2/lrs.proto", fileDescriptor_cd7332d279836518)
}

var fileDescriptor_cd7332d279836518 = []byte{
	// 448 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0xc1, 0x8e, 0xd3, 0x30,
	0x14, 0x45, 0x71, 0x0b, 0xa3, 0xe0, 0x01, 0x31, 0x44, 0xa0, 0x66, 0x3a, 0x80, 0xaa, 0x02, 0x22,
	0x08, 0x61, 0xa3, 0xb0, 0x67, 0x11, 0x40, 0x80, 0x54, 0xa1, 0x21, 0xfd, 0x80, 0xca, 0x6d, 0x1e,
	0x95, 0xa5, 0xe0, 0x17, 0x6c, 0xc7, 0xa2, 0x3b, 0x96, 0xb0, 0x61, 0xc1, 0xe7, 0xf0, 0x05, 0x2c,
	0xd8, 0xf0, 0x3b, 0xac, 0x50, 0x62, 0x67, 0x26, 0xc3, 0x62, 0xc4, 0xae, 0xee, 0x3b, 0xf7, 0xf9,
	0xde, 0x1b, 0xd3, 0xfb, 0xa0, 0x1c, 0xee, 0xb8, 0x01, 0xed, 0xe4, 0x06, 0x78, 0x85, 0xa2, 0x5c,
	0x19, 0x2b, 0xac, 0xe1, 0x2e, 0xe3, 0x95, 0x36, 0xac, 0xd6, 0x68, 0x31, 0x3e, 0xea, 0x30, 0x16,
	0x30, 0x76, 0x8a, 0x31, 0x97, 0x4d, 0x6f, 0xf9, 0x1d, 0xa2, 0x96, 0xad, 0x68, 0x83, 0x1a, 0xf8,
	0x5a, 0x18, 0xf0, 0xd2, 0xe9, 0x83, 0x33, 0x53, 0x50, 0x65, 0x8d, 0x52, 0x59, 0x7f, 0x93, 0x86,
	0x1a, 0xb5, 0x0d, 0xe0, 0x9d, 0x2d, 0xe2, 0xb6, 0x02, 0xde, 0x9d, 0xd6, 0xcd, 0x7b, 0x5e, 0x36,
	0x5a, 0x58, 0x89, 0x2a, 0xcc, 0x6f, 0x37, 0x65, 0x2d, 0xb8, 0x50, 0x0a, 0x6d, 0xf7, 0xb7, 0xe1,
	0xad, 0x83, 0x26, 0x58, 0x9c, 0x4e, 0x9c, 0xa8, 0x64, 0x29, 0x2c, 0xf0, 0xfe, 0x87, 0x1f, 0xcc,
	0xbf, 0x12, 0x7a, 0xb0, 0x40, 0x51, 0x2e, 0x5b, 0xbf, 0x05, 0x7c, 0x6c, 0xc0, 0xd8, 0xf8, 0x11,
	0xbd, 0xa8, 0xb0, 0x84, 0x84, 0xcc, 0x48, 0xba, 0x9f, 0x4d, 0x98, 0xcf, 0x27, 0x6a, 0xc9, 0x5c,
	0xc6, 0xda, 0x08, 0xec, 0x2d, 0x96, 0x50, 0x74, 0x50, 0xfc, 0x9a, 0x5e, 0xdd, 0x54, 0x8d, 0xb1,
	0xa0, 0x7d, 0xe8, 0x64, 0x34, 0x1b, 0xa7, 0xfb, 0xd9, 0xdd, 0xb3, 0xaa, 0x3e, 0x1a, 0x7b, 0xee,
	0x59, 0x7f, 0xdf, 0x95, 0xcd, 0xe0, 0x34, 0xff, 0x45, 0xe8, 0xf5, 0x81, 0x17, 0x53, 0xa3, 0x32,
	0x10, 0xdf, 0xa3, 0x51, 0xa0, 0x4c, 0x42, 0x66, 0xe3, 0xf4, 0x72, 0x1e, 0xfd, 0xc9, 0x2f, 0x7d,
	0x27, 0xa3, 0x88, 0x14, 0x27, 0x93, 0xf8, 0x1d, 0x9d, 0x0c, 0x4a, 0x93, 0x6a, 0xbb, 0x92, 0xca,
	0x82, 0x76, 0xa2, 0x4a, 0x46, 0x5d, 0x8a, 0x43, 0xe6, 0x1b, 0x64, 0x7d, 0x83, 0xec, 0x45, 0x68,
	0xb0, 0xb8, 0xd9, 0x2a, 0x8b, 0x5e, 0xf8, 0x26, 0xe8, 0xe2, 0x67, 0xf4, 0xc8, 0x6f, 0x5b, 0xf5,
	0xe6, 0x57, 0x5b, 0x2d, 0x54, 0x53, 0x09, 0x2d, 0xed, 0x2e, 0x19, 0xcf, 0x48, 0x1a, 0x15, 0x87,
	0x1e, 0x79, 0x19, 0x88, 0x57, 0xa7, 0x40, 0xf6, 0x8d, 0xd0, 0x1b, 0x8b, 0xe1, 0xe6, 0xa5, 0x7f,
	0x20, 0xb1, 0xa3, 0xd7, 0x96, 0x56, 0x83, 0xf8, 0x70, 0x12, 0x36, 0x7e, 0xcc, 0xce, 0x79, 0x43,
	0xec, 0xdf, 0x0f, 0x34, 0x65, 0xff, 0x8b, 0xfb, 0x0e, 0xe7, 0x17, 0x52, 0xf2, 0x84, 0xe4, 0xf9,
	0x8f, 0xcf, 0x3f, 0x7f, 0xef, 0x8d, 0x0e, 0x08, 0x7d, 0x28, 0xd1, 0x6f, 0xa8, 0x35, 0x7e, 0xda,
	0x9d, 0xb7, 0x2c, 0x8f, 0x16, 0xda, 0x1c, 0xb7, 0x95, 0x1d, 0x93, 0x2f, 0x84, 0xac, 0xf7, 0xba,
	0xfa, 0x9e, 0xfe, 0x0d, 0x00, 0x00, 0xff, 0xff, 0xe7, 0xc9, 0x50, 0x71, 0x1b, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// LoadReportingServiceClient is the client API for LoadReportingService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type LoadReportingServiceClient interface {
	StreamLoadStats(ctx context.Context, opts ...grpc.CallOption) (LoadReportingService_StreamLoadStatsClient, error)
}

type loadReportingServiceClient struct {
	cc *grpc.ClientConn
}

func NewLoadReportingServiceClient(cc *grpc.ClientConn) LoadReportingServiceClient {
	return &loadReportingServiceClient{cc}
}

func (c *loadReportingServiceClient) StreamLoadStats(ctx context.Context, opts ...grpc.CallOption) (LoadReportingService_StreamLoadStatsClient, error) {
	stream, err := c.cc.NewStream(ctx, &_LoadReportingService_serviceDesc.Streams[0], "/envoy.service.load_stats.v2.LoadReportingService/StreamLoadStats", opts...)
	if err != nil {
		return nil, err
	}
	x := &loadReportingServiceStreamLoadStatsClient{stream}
	return x, nil
}

type LoadReportingService_StreamLoadStatsClient interface {
	Send(*LoadStatsRequest) error
	Recv() (*LoadStatsResponse, error)
	grpc.ClientStream
}

type loadReportingServiceStreamLoadStatsClient struct {
	grpc.ClientStream
}

func (x *loadReportingServiceStreamLoadStatsClient) Send(m *LoadStatsRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *loadReportingServiceStreamLoadStatsClient) Recv() (*LoadStatsResponse, error) {
	m := new(LoadStatsResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// LoadReportingServiceServer is the server API for LoadReportingService service.
type LoadReportingServiceServer interface {
	StreamLoadStats(LoadReportingService_StreamLoadStatsServer) error
}

// UnimplementedLoadReportingServiceServer can be embedded to have forward compatible implementations.
type UnimplementedLoadReportingServiceServer struct {
}

func (*UnimplementedLoadReportingServiceServer) StreamLoadStats(srv LoadReportingService_StreamLoadStatsServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamLoadStats not implemented")
}

func RegisterLoadReportingServiceServer(s *grpc.Server, srv LoadReportingServiceServer) {
	s.RegisterService(&_LoadReportingService_serviceDesc, srv)
}

func _LoadReportingService_StreamLoadStats_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(LoadReportingServiceServer).StreamLoadStats(&loadReportingServiceStreamLoadStatsServer{stream})
}

type LoadReportingService_StreamLoadStatsServer interface {
	Send(*LoadStatsResponse) error
	Recv() (*LoadStatsRequest, error)
	grpc.ServerStream
}

type loadReportingServiceStreamLoadStatsServer struct {
	grpc.ServerStream
}

func (x *loadReportingServiceStreamLoadStatsServer) Send(m *LoadStatsResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *loadReportingServiceStreamLoadStatsServer) Recv() (*LoadStatsRequest, error) {
	m := new(LoadStatsRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _LoadReportingService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "envoy.service.load_stats.v2.LoadReportingService",
	HandlerType: (*LoadReportingServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamLoadStats",
			Handler:       _LoadReportingService_StreamLoadStats_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "envoy/service/load_stats/v2/lrs.proto",
}
