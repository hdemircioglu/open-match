// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api/query.proto

package pb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type QueryTicketsRequest struct {
	// A Pool is consists of a set of Filters.
	Pool                 *Pool    `protobuf:"bytes,1,opt,name=pool,proto3" json:"pool,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *QueryTicketsRequest) Reset()         { *m = QueryTicketsRequest{} }
func (m *QueryTicketsRequest) String() string { return proto.CompactTextString(m) }
func (*QueryTicketsRequest) ProtoMessage()    {}
func (*QueryTicketsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5ec7651f31a90698, []int{0}
}

func (m *QueryTicketsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QueryTicketsRequest.Unmarshal(m, b)
}
func (m *QueryTicketsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QueryTicketsRequest.Marshal(b, m, deterministic)
}
func (m *QueryTicketsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryTicketsRequest.Merge(m, src)
}
func (m *QueryTicketsRequest) XXX_Size() int {
	return xxx_messageInfo_QueryTicketsRequest.Size(m)
}
func (m *QueryTicketsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryTicketsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryTicketsRequest proto.InternalMessageInfo

func (m *QueryTicketsRequest) GetPool() *Pool {
	if m != nil {
		return m.Pool
	}
	return nil
}

type QueryTicketsResponse struct {
	// Tickets that satisfy all the filtering criteria.
	Tickets              []*Ticket `protobuf:"bytes,1,rep,name=tickets,proto3" json:"tickets,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *QueryTicketsResponse) Reset()         { *m = QueryTicketsResponse{} }
func (m *QueryTicketsResponse) String() string { return proto.CompactTextString(m) }
func (*QueryTicketsResponse) ProtoMessage()    {}
func (*QueryTicketsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_5ec7651f31a90698, []int{1}
}

func (m *QueryTicketsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QueryTicketsResponse.Unmarshal(m, b)
}
func (m *QueryTicketsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QueryTicketsResponse.Marshal(b, m, deterministic)
}
func (m *QueryTicketsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryTicketsResponse.Merge(m, src)
}
func (m *QueryTicketsResponse) XXX_Size() int {
	return xxx_messageInfo_QueryTicketsResponse.Size(m)
}
func (m *QueryTicketsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryTicketsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryTicketsResponse proto.InternalMessageInfo

func (m *QueryTicketsResponse) GetTickets() []*Ticket {
	if m != nil {
		return m.Tickets
	}
	return nil
}

func init() {
	proto.RegisterType((*QueryTicketsRequest)(nil), "openmatch.QueryTicketsRequest")
	proto.RegisterType((*QueryTicketsResponse)(nil), "openmatch.QueryTicketsResponse")
}

func init() { proto.RegisterFile("api/query.proto", fileDescriptor_5ec7651f31a90698) }

var fileDescriptor_5ec7651f31a90698 = []byte{
	// 516 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x52, 0xcd, 0x6e, 0x13, 0x3d,
	0x14, 0xd5, 0x4c, 0xaa, 0x56, 0x75, 0x3f, 0xa9, 0x1f, 0xe6, 0x47, 0x55, 0x84, 0x8a, 0x49, 0x37,
	0x69, 0x20, 0xe3, 0x74, 0xc8, 0x2a, 0x08, 0xa9, 0xa5, 0xed, 0xa2, 0x52, 0xc2, 0x4f, 0x8a, 0x58,
	0xb0, 0x73, 0x9c, 0xcb, 0x8c, 0x69, 0xc6, 0xd7, 0xb5, 0x3d, 0x29, 0x95, 0x58, 0xb1, 0x66, 0x05,
	0x1b, 0xc4, 0x23, 0xf0, 0x12, 0x3c, 0x04, 0xaf, 0x80, 0x78, 0x0e, 0x34, 0x33, 0x09, 0x0d, 0xb4,
	0xac, 0x46, 0xbe, 0xe7, 0xdc, 0x73, 0xce, 0x1c, 0x9b, 0xac, 0x0b, 0xa3, 0xf8, 0x69, 0x0e, 0xf6,
	0x3c, 0x32, 0x16, 0x3d, 0xd2, 0x55, 0x34, 0xa0, 0x33, 0xe1, 0x65, 0x5a, 0xa7, 0x05, 0x96, 0x81,
	0x73, 0x22, 0x01, 0x57, 0xc1, 0xf5, 0xdb, 0x09, 0x62, 0x32, 0x01, 0x5e, 0x40, 0x42, 0x6b, 0xf4,
	0xc2, 0x2b, 0xd4, 0x73, 0xf4, 0x7e, 0xf9, 0x91, 0xed, 0x04, 0x74, 0xdb, 0x9d, 0x89, 0x24, 0x01,
	0xcb, 0xd1, 0x94, 0x8c, 0xcb, 0xec, 0x46, 0x8f, 0x5c, 0x7f, 0x5e, 0x38, 0xbf, 0x50, 0xf2, 0x04,
	0xbc, 0x1b, 0xc2, 0x69, 0x0e, 0xce, 0xd3, 0x2d, 0xb2, 0x64, 0x10, 0x27, 0x1b, 0x01, 0x0b, 0x9a,
	0x6b, 0xf1, 0x7a, 0xf4, 0x3b, 0x50, 0xf4, 0x0c, 0x71, 0x32, 0x2c, 0xc1, 0xc6, 0x3e, 0xb9, 0xf1,
	0xe7, 0xae, 0x33, 0xa8, 0x1d, 0xd0, 0x7b, 0x64, 0xc5, 0x57, 0xa3, 0x8d, 0x80, 0xd5, 0x9a, 0x6b,
	0xf1, 0xb5, 0x85, 0xfd, 0x8a, 0x3c, 0x9c, 0x33, 0xe2, 0x0f, 0x01, 0xf9, 0xaf, 0x54, 0x39, 0x06,
	0x3b, 0x55, 0x12, 0xe8, 0xbb, 0xd9, 0x79, 0xa6, 0x4a, 0x37, 0x17, 0x96, 0xaf, 0x88, 0x5a, 0xbf,
	0xf3, 0x4f, 0xbc, 0x8a, 0xd3, 0xd8, 0x7e, 0xff, 0xfd, 0xc7, 0xa7, 0x70, 0xab, 0xb1, 0xc9, 0xa7,
	0x3b, 0x55, 0xcd, 0xae, 0xb2, 0xe2, 0xb3, 0x0c, 0xbd, 0x72, 0xd8, 0x0b, 0x5a, 0x9d, 0xe0, 0xf1,
	0xe7, 0xda, 0xc7, 0xbd, 0x9f, 0x21, 0xfd, 0x16, 0x90, 0x9b, 0x83, 0x01, 0xeb, 0x63, 0xa2, 0x24,
	0x6b, 0x1e, 0x08, 0x2f, 0x58, 0x5f, 0x9c, 0x83, 0xdd, 0x6e, 0x1c, 0x11, 0xf2, 0xd4, 0x80, 0x66,
	0x83, 0xc2, 0x90, 0xde, 0x4a, 0xbd, 0x37, 0xae, 0xc7, 0x79, 0x91, 0xa1, 0x5d, 0x85, 0x18, 0xc3,
	0xb4, 0xbe, 0x75, 0x71, 0x6e, 0x8f, 0x95, 0x93, 0xb9, 0x73, 0xbb, 0xd5, 0xad, 0x25, 0x16, 0x73,
	0xe3, 0x22, 0x89, 0x59, 0xeb, 0x25, 0xa1, 0x7b, 0x46, 0xc8, 0x14, 0x58, 0x1c, 0x75, 0x58, 0x5f,
	0x49, 0x28, 0xda, 0xdb, 0x9d, 0x4b, 0x26, 0xca, 0xa7, 0xf9, 0xa8, 0x60, 0xf2, 0x6a, 0xf5, 0x35,
	0xda, 0x44, 0x64, 0xe0, 0x16, 0xcc, 0xf8, 0x68, 0x82, 0x23, 0x9e, 0x09, 0xe7, 0xc1, 0xf2, 0xfe,
	0xd1, 0xfe, 0xe1, 0x93, 0xe3, 0xc3, 0xb8, 0xb6, 0x13, 0x75, 0x5a, 0x61, 0x10, 0xc6, 0xff, 0x0b,
	0x63, 0x26, 0x4a, 0x96, 0x17, 0xce, 0xdf, 0x38, 0xd4, 0xbd, 0x4b, 0x93, 0xe1, 0x43, 0x52, 0xeb,
	0x76, 0xba, 0xb4, 0x4b, 0x5a, 0x43, 0xf0, 0xb9, 0xd5, 0x30, 0x66, 0x67, 0x29, 0x68, 0xe6, 0x53,
	0x60, 0x16, 0x1c, 0xe6, 0x56, 0x02, 0x1b, 0x23, 0x38, 0xa6, 0xd1, 0x33, 0x78, 0xab, 0x9c, 0x8f,
	0xe8, 0x32, 0x59, 0xfa, 0x12, 0x06, 0x2b, 0xf6, 0x11, 0xd9, 0xb8, 0x28, 0x83, 0x1d, 0xa0, 0xcc,
	0x33, 0xd0, 0xd5, 0x03, 0xa3, 0x77, 0xaf, 0xae, 0x86, 0x3b, 0xe5, 0x81, 0x8f, 0x51, 0x3a, 0xfe,
	0x8a, 0xfd, 0x05, 0x2d, 0xfc, 0x97, 0x39, 0x49, 0xb8, 0x19, 0x7d, 0x0d, 0x57, 0x0b, 0xfd, 0x52,
	0x7e, 0xb4, 0x5c, 0xbe, 0xd8, 0x07, 0xbf, 0x02, 0x00, 0x00, 0xff, 0xff, 0xaa, 0x04, 0x46, 0x55,
	0x2f, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// QueryServiceClient is the client API for QueryService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QueryServiceClient interface {
	// QueryTickets gets a list of Tickets that match all Filters of the input Pool.
	//   - If the Pool contains no Filters, QueryTickets will return all Tickets in the state storage.
	// QueryTickets pages the Tickets by `storage.pool.size` and stream back response.
	//   - storage.pool.size is default to 1000 if not set, and has a mininum of 10 and maximum of 10000
	QueryTickets(ctx context.Context, in *QueryTicketsRequest, opts ...grpc.CallOption) (QueryService_QueryTicketsClient, error)
}

type queryServiceClient struct {
	cc *grpc.ClientConn
}

func NewQueryServiceClient(cc *grpc.ClientConn) QueryServiceClient {
	return &queryServiceClient{cc}
}

func (c *queryServiceClient) QueryTickets(ctx context.Context, in *QueryTicketsRequest, opts ...grpc.CallOption) (QueryService_QueryTicketsClient, error) {
	stream, err := c.cc.NewStream(ctx, &_QueryService_serviceDesc.Streams[0], "/openmatch.QueryService/QueryTickets", opts...)
	if err != nil {
		return nil, err
	}
	x := &queryServiceQueryTicketsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type QueryService_QueryTicketsClient interface {
	Recv() (*QueryTicketsResponse, error)
	grpc.ClientStream
}

type queryServiceQueryTicketsClient struct {
	grpc.ClientStream
}

func (x *queryServiceQueryTicketsClient) Recv() (*QueryTicketsResponse, error) {
	m := new(QueryTicketsResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// QueryServiceServer is the server API for QueryService service.
type QueryServiceServer interface {
	// QueryTickets gets a list of Tickets that match all Filters of the input Pool.
	//   - If the Pool contains no Filters, QueryTickets will return all Tickets in the state storage.
	// QueryTickets pages the Tickets by `storage.pool.size` and stream back response.
	//   - storage.pool.size is default to 1000 if not set, and has a mininum of 10 and maximum of 10000
	QueryTickets(*QueryTicketsRequest, QueryService_QueryTicketsServer) error
}

// UnimplementedQueryServiceServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServiceServer struct {
}

func (*UnimplementedQueryServiceServer) QueryTickets(req *QueryTicketsRequest, srv QueryService_QueryTicketsServer) error {
	return status.Errorf(codes.Unimplemented, "method QueryTickets not implemented")
}

func RegisterQueryServiceServer(s *grpc.Server, srv QueryServiceServer) {
	s.RegisterService(&_QueryService_serviceDesc, srv)
}

func _QueryService_QueryTickets_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(QueryTicketsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(QueryServiceServer).QueryTickets(m, &queryServiceQueryTicketsServer{stream})
}

type QueryService_QueryTicketsServer interface {
	Send(*QueryTicketsResponse) error
	grpc.ServerStream
}

type queryServiceQueryTicketsServer struct {
	grpc.ServerStream
}

func (x *queryServiceQueryTicketsServer) Send(m *QueryTicketsResponse) error {
	return x.ServerStream.SendMsg(m)
}

var _QueryService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "openmatch.QueryService",
	HandlerType: (*QueryServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "QueryTickets",
			Handler:       _QueryService_QueryTickets_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "api/query.proto",
}
