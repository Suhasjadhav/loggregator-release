// Code generated by protoc-gen-go.
// source: ingress.proto
// DO NOT EDIT!

package loggregator_v2

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

type IngressResponse struct {
}

func (m *IngressResponse) Reset()                    { *m = IngressResponse{} }
func (m *IngressResponse) String() string            { return proto.CompactTextString(m) }
func (*IngressResponse) ProtoMessage()               {}
func (*IngressResponse) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{0} }

func init() {
	proto.RegisterType((*IngressResponse)(nil), "loggregator.v2.IngressResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Ingress service

type IngressClient interface {
	Sender(ctx context.Context, opts ...grpc.CallOption) (Ingress_SenderClient, error)
}

type ingressClient struct {
	cc *grpc.ClientConn
}

func NewIngressClient(cc *grpc.ClientConn) IngressClient {
	return &ingressClient{cc}
}

func (c *ingressClient) Sender(ctx context.Context, opts ...grpc.CallOption) (Ingress_SenderClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Ingress_serviceDesc.Streams[0], c.cc, "/loggregator.v2.Ingress/Sender", opts...)
	if err != nil {
		return nil, err
	}
	x := &ingressSenderClient{stream}
	return x, nil
}

type Ingress_SenderClient interface {
	Send(*Envelope) error
	CloseAndRecv() (*IngressResponse, error)
	grpc.ClientStream
}

type ingressSenderClient struct {
	grpc.ClientStream
}

func (x *ingressSenderClient) Send(m *Envelope) error {
	return x.ClientStream.SendMsg(m)
}

func (x *ingressSenderClient) CloseAndRecv() (*IngressResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(IngressResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for Ingress service

type IngressServer interface {
	Sender(Ingress_SenderServer) error
}

func RegisterIngressServer(s *grpc.Server, srv IngressServer) {
	s.RegisterService(&_Ingress_serviceDesc, srv)
}

func _Ingress_Sender_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(IngressServer).Sender(&ingressSenderServer{stream})
}

type Ingress_SenderServer interface {
	SendAndClose(*IngressResponse) error
	Recv() (*Envelope, error)
	grpc.ServerStream
}

type ingressSenderServer struct {
	grpc.ServerStream
}

func (x *ingressSenderServer) SendAndClose(m *IngressResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *ingressSenderServer) Recv() (*Envelope, error) {
	m := new(Envelope)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _Ingress_serviceDesc = grpc.ServiceDesc{
	ServiceName: "loggregator.v2.Ingress",
	HandlerType: (*IngressServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Sender",
			Handler:       _Ingress_Sender_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "ingress.proto",
}

func init() { proto.RegisterFile("ingress.proto", fileDescriptor3) }

var fileDescriptor3 = []byte{
	// 123 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0xcd, 0xcc, 0x4b, 0x2f,
	0x4a, 0x2d, 0x2e, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0xcb, 0xc9, 0x4f, 0x4f, 0x2f,
	0x4a, 0x4d, 0x4f, 0x2c, 0xc9, 0x2f, 0xd2, 0x2b, 0x33, 0x92, 0xe2, 0x4b, 0xcd, 0x2b, 0x4b, 0xcd,
	0xc9, 0x2f, 0x48, 0x85, 0xc8, 0x2b, 0x09, 0x72, 0xf1, 0x7b, 0x42, 0x34, 0x04, 0xa5, 0x16, 0x17,
	0xe4, 0xe7, 0x15, 0xa7, 0x1a, 0x05, 0x71, 0xb1, 0x43, 0x85, 0x84, 0xdc, 0xb9, 0xd8, 0x82, 0x53,
	0xf3, 0x52, 0x52, 0x8b, 0x84, 0x24, 0xf4, 0x50, 0x0d, 0xd2, 0x73, 0x85, 0x9a, 0x23, 0x25, 0x8f,
	0x2e, 0x83, 0x66, 0x9e, 0x12, 0x83, 0x06, 0x63, 0x12, 0x1b, 0xd8, 0x36, 0x63, 0x40, 0x00, 0x00,
	0x00, 0xff, 0xff, 0xfb, 0x5d, 0x75, 0xe0, 0x9e, 0x00, 0x00, 0x00,
}
