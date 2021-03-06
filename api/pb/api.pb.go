// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api/pb/api.proto

/*
Package pb is a generated protocol buffer package.

It is generated from these files:
	api/pb/api.proto

It has these top-level messages:
	GetBlockReq
	GetBlockRsp
	SubscribeBlockReq
	SubscribeBlockRsp
*/
package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type GetBlockReq struct {
	Number int64 `protobuf:"varint,1,opt,name=number" json:"number,omitempty"`
}

func (m *GetBlockReq) Reset()                    { *m = GetBlockReq{} }
func (m *GetBlockReq) String() string            { return proto.CompactTextString(m) }
func (*GetBlockReq) ProtoMessage()               {}
func (*GetBlockReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *GetBlockReq) GetNumber() int64 {
	if m != nil {
		return m.Number
	}
	return 0
}

type GetBlockRsp struct {
	Block string `protobuf:"bytes,1,opt,name=block" json:"block,omitempty"`
}

func (m *GetBlockRsp) Reset()                    { *m = GetBlockRsp{} }
func (m *GetBlockRsp) String() string            { return proto.CompactTextString(m) }
func (*GetBlockRsp) ProtoMessage()               {}
func (*GetBlockRsp) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *GetBlockRsp) GetBlock() string {
	if m != nil {
		return m.Block
	}
	return ""
}

type SubscribeBlockReq struct {
	Number int64 `protobuf:"varint,1,opt,name=number" json:"number,omitempty"`
}

func (m *SubscribeBlockReq) Reset()                    { *m = SubscribeBlockReq{} }
func (m *SubscribeBlockReq) String() string            { return proto.CompactTextString(m) }
func (*SubscribeBlockReq) ProtoMessage()               {}
func (*SubscribeBlockReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *SubscribeBlockReq) GetNumber() int64 {
	if m != nil {
		return m.Number
	}
	return 0
}

type SubscribeBlockRsp struct {
	Block string `protobuf:"bytes,1,opt,name=block" json:"block,omitempty"`
}

func (m *SubscribeBlockRsp) Reset()                    { *m = SubscribeBlockRsp{} }
func (m *SubscribeBlockRsp) String() string            { return proto.CompactTextString(m) }
func (*SubscribeBlockRsp) ProtoMessage()               {}
func (*SubscribeBlockRsp) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *SubscribeBlockRsp) GetBlock() string {
	if m != nil {
		return m.Block
	}
	return ""
}

func init() {
	proto.RegisterType((*GetBlockReq)(nil), "pb.GetBlockReq")
	proto.RegisterType((*GetBlockRsp)(nil), "pb.GetBlockRsp")
	proto.RegisterType((*SubscribeBlockReq)(nil), "pb.SubscribeBlockReq")
	proto.RegisterType((*SubscribeBlockRsp)(nil), "pb.SubscribeBlockRsp")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for BlockService service

type BlockServiceClient interface {
	GetBlock(ctx context.Context, in *GetBlockReq, opts ...grpc.CallOption) (*GetBlockRsp, error)
	SubscribeBlock(ctx context.Context, in *SubscribeBlockReq, opts ...grpc.CallOption) (BlockService_SubscribeBlockClient, error)
}

type blockServiceClient struct {
	cc *grpc.ClientConn
}

func NewBlockServiceClient(cc *grpc.ClientConn) BlockServiceClient {
	return &blockServiceClient{cc}
}

func (c *blockServiceClient) GetBlock(ctx context.Context, in *GetBlockReq, opts ...grpc.CallOption) (*GetBlockRsp, error) {
	out := new(GetBlockRsp)
	err := grpc.Invoke(ctx, "/pb.BlockService/GetBlock", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blockServiceClient) SubscribeBlock(ctx context.Context, in *SubscribeBlockReq, opts ...grpc.CallOption) (BlockService_SubscribeBlockClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_BlockService_serviceDesc.Streams[0], c.cc, "/pb.BlockService/SubscribeBlock", opts...)
	if err != nil {
		return nil, err
	}
	x := &blockServiceSubscribeBlockClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type BlockService_SubscribeBlockClient interface {
	Recv() (*SubscribeBlockRsp, error)
	grpc.ClientStream
}

type blockServiceSubscribeBlockClient struct {
	grpc.ClientStream
}

func (x *blockServiceSubscribeBlockClient) Recv() (*SubscribeBlockRsp, error) {
	m := new(SubscribeBlockRsp)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for BlockService service

type BlockServiceServer interface {
	GetBlock(context.Context, *GetBlockReq) (*GetBlockRsp, error)
	SubscribeBlock(*SubscribeBlockReq, BlockService_SubscribeBlockServer) error
}

func RegisterBlockServiceServer(s *grpc.Server, srv BlockServiceServer) {
	s.RegisterService(&_BlockService_serviceDesc, srv)
}

func _BlockService_GetBlock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBlockReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlockServiceServer).GetBlock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.BlockService/GetBlock",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlockServiceServer).GetBlock(ctx, req.(*GetBlockReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _BlockService_SubscribeBlock_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(SubscribeBlockReq)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(BlockServiceServer).SubscribeBlock(m, &blockServiceSubscribeBlockServer{stream})
}

type BlockService_SubscribeBlockServer interface {
	Send(*SubscribeBlockRsp) error
	grpc.ServerStream
}

type blockServiceSubscribeBlockServer struct {
	grpc.ServerStream
}

func (x *blockServiceSubscribeBlockServer) Send(m *SubscribeBlockRsp) error {
	return x.ServerStream.SendMsg(m)
}

var _BlockService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.BlockService",
	HandlerType: (*BlockServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetBlock",
			Handler:    _BlockService_GetBlock_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SubscribeBlock",
			Handler:       _BlockService_SubscribeBlock_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "api/pb/api.proto",
}

func init() { proto.RegisterFile("api/pb/api.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 180 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x48, 0x2c, 0xc8, 0xd4,
	0x2f, 0x48, 0xd2, 0x4f, 0x2c, 0xc8, 0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2a, 0x48,
	0x52, 0x52, 0xe5, 0xe2, 0x76, 0x4f, 0x2d, 0x71, 0xca, 0xc9, 0x4f, 0xce, 0x0e, 0x4a, 0x2d, 0x14,
	0x12, 0xe3, 0x62, 0xcb, 0x2b, 0xcd, 0x4d, 0x4a, 0x2d, 0x92, 0x60, 0x54, 0x60, 0xd4, 0x60, 0x0e,
	0x82, 0xf2, 0x94, 0x94, 0x91, 0x94, 0x15, 0x17, 0x08, 0x89, 0x70, 0xb1, 0x26, 0x81, 0xd8, 0x60,
	0x55, 0x9c, 0x41, 0x10, 0x8e, 0x92, 0x36, 0x97, 0x60, 0x70, 0x69, 0x52, 0x71, 0x72, 0x51, 0x66,
	0x52, 0x2a, 0x41, 0x13, 0x35, 0x31, 0x14, 0xe3, 0x32, 0xd7, 0xa8, 0x89, 0x91, 0x8b, 0x07, 0xac,
	0x24, 0x38, 0xb5, 0xa8, 0x2c, 0x33, 0x39, 0x55, 0x48, 0x8f, 0x8b, 0x03, 0xe6, 0x1a, 0x21, 0x7e,
	0xbd, 0x82, 0x24, 0x3d, 0x24, 0x2f, 0x48, 0xa1, 0x0a, 0x14, 0x17, 0x28, 0x31, 0x08, 0x39, 0x71,
	0xf1, 0xa1, 0xda, 0x25, 0x24, 0x0a, 0x52, 0x84, 0xe1, 0x58, 0x29, 0x6c, 0xc2, 0x20, 0x13, 0x0c,
	0x18, 0x93, 0xd8, 0xc0, 0x61, 0x66, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0x0f, 0xa0, 0x5a, 0x22,
	0x47, 0x01, 0x00, 0x00,
}
