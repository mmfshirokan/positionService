// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.25.1
// source: balance.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// BalanceClient is the client API for Balance service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BalanceClient interface {
	Balancer(ctx context.Context, opts ...grpc.CallOption) (Balance_BalancerClient, error)
}

type balanceClient struct {
	cc grpc.ClientConnInterface
}

func NewBalanceClient(cc grpc.ClientConnInterface) BalanceClient {
	return &balanceClient{cc}
}

func (c *balanceClient) Balancer(ctx context.Context, opts ...grpc.CallOption) (Balance_BalancerClient, error) {
	stream, err := c.cc.NewStream(ctx, &Balance_ServiceDesc.Streams[0], "/pb.Balance/Balancer", opts...)
	if err != nil {
		return nil, err
	}
	x := &balanceBalancerClient{stream}
	return x, nil
}

type Balance_BalancerClient interface {
	Send(*RequestBalancer) error
	Recv() (*ResponseBalancer, error)
	grpc.ClientStream
}

type balanceBalancerClient struct {
	grpc.ClientStream
}

func (x *balanceBalancerClient) Send(m *RequestBalancer) error {
	return x.ClientStream.SendMsg(m)
}

func (x *balanceBalancerClient) Recv() (*ResponseBalancer, error) {
	m := new(ResponseBalancer)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// BalanceServer is the server API for Balance service.
// All implementations must embed UnimplementedBalanceServer
// for forward compatibility
type BalanceServer interface {
	Balancer(Balance_BalancerServer) error
	mustEmbedUnimplementedBalanceServer()
}

// UnimplementedBalanceServer must be embedded to have forward compatible implementations.
type UnimplementedBalanceServer struct {
}

func (UnimplementedBalanceServer) Balancer(Balance_BalancerServer) error {
	return status.Errorf(codes.Unimplemented, "method Balancer not implemented")
}
func (UnimplementedBalanceServer) mustEmbedUnimplementedBalanceServer() {}

// UnsafeBalanceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BalanceServer will
// result in compilation errors.
type UnsafeBalanceServer interface {
	mustEmbedUnimplementedBalanceServer()
}

func RegisterBalanceServer(s grpc.ServiceRegistrar, srv BalanceServer) {
	s.RegisterService(&Balance_ServiceDesc, srv)
}

func _Balance_Balancer_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(BalanceServer).Balancer(&balanceBalancerServer{stream})
}

type Balance_BalancerServer interface {
	Send(*ResponseBalancer) error
	Recv() (*RequestBalancer, error)
	grpc.ServerStream
}

type balanceBalancerServer struct {
	grpc.ServerStream
}

func (x *balanceBalancerServer) Send(m *ResponseBalancer) error {
	return x.ServerStream.SendMsg(m)
}

func (x *balanceBalancerServer) Recv() (*RequestBalancer, error) {
	m := new(RequestBalancer)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Balance_ServiceDesc is the grpc.ServiceDesc for Balance service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Balance_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.Balance",
	HandlerType: (*BalanceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Balancer",
			Handler:       _Balance_Balancer_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "balance.proto",
}
