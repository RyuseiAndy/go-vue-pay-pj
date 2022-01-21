// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package paymentservice

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

// PayManagerClient is the client API for PayManager service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PayManagerClient interface {
	Charge(ctx context.Context, in *PayRequest, opts ...grpc.CallOption) (*PayResponse, error)
}

type payManagerClient struct {
	cc grpc.ClientConnInterface
}

func NewPayManagerClient(cc grpc.ClientConnInterface) PayManagerClient {
	return &payManagerClient{cc}
}

func (c *payManagerClient) Charge(ctx context.Context, in *PayRequest, opts ...grpc.CallOption) (*PayResponse, error) {
	out := new(PayResponse)
	err := c.cc.Invoke(ctx, "/paymentservice.PayManager/Charge", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PayManagerServer is the server API for PayManager service.
// All implementations must embed UnimplementedPayManagerServer
// for forward compatibility
type PayManagerServer interface {
	Charge(context.Context, *PayRequest) (*PayResponse, error)
	mustEmbedUnimplementedPayManagerServer()
}

// UnimplementedPayManagerServer must be embedded to have forward compatible implementations.
type UnimplementedPayManagerServer struct {
}

func (UnimplementedPayManagerServer) Charge(context.Context, *PayRequest) (*PayResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Charge not implemented")
}
func (UnimplementedPayManagerServer) mustEmbedUnimplementedPayManagerServer() {}

// UnsafePayManagerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PayManagerServer will
// result in compilation errors.
type UnsafePayManagerServer interface {
	mustEmbedUnimplementedPayManagerServer()
}

func RegisterPayManagerServer(s grpc.ServiceRegistrar, srv PayManagerServer) {
	s.RegisterService(&PayManager_ServiceDesc, srv)
}

func _PayManager_Charge_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PayRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PayManagerServer).Charge(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/paymentservice.PayManager/Charge",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PayManagerServer).Charge(ctx, req.(*PayRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PayManager_ServiceDesc is the grpc.ServiceDesc for PayManager service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PayManager_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "paymentservice.PayManager",
	HandlerType: (*PayManagerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Charge",
			Handler:    _PayManager_Charge_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/pay.proto",
}
