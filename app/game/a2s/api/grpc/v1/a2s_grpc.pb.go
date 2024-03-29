// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package v1

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

// A2SClient is the client API for A2S service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type A2SClient interface {
	A2SQuery(ctx context.Context, in *A2SQueryReq, opts ...grpc.CallOption) (*A2SQueryResp, error)
}

type a2SClient struct {
	cc grpc.ClientConnInterface
}

func NewA2SClient(cc grpc.ClientConnInterface) A2SClient {
	return &a2SClient{cc}
}

func (c *a2SClient) A2SQuery(ctx context.Context, in *A2SQueryReq, opts ...grpc.CallOption) (*A2SQueryResp, error) {
	out := new(A2SQueryResp)
	err := c.cc.Invoke(ctx, "/ncs.game.a2s.v1.A2S/A2SQuery", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// A2SServer is the server API for A2S service.
// All implementations must embed UnimplementedA2SServer
// for forward compatibility
type A2SServer interface {
	A2SQuery(context.Context, *A2SQueryReq) (*A2SQueryResp, error)
	mustEmbedUnimplementedA2SServer()
}

// UnimplementedA2SServer must be embedded to have forward compatible implementations.
type UnimplementedA2SServer struct {
}

func (UnimplementedA2SServer) A2SQuery(context.Context, *A2SQueryReq) (*A2SQueryResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method A2SQuery not implemented")
}
func (UnimplementedA2SServer) mustEmbedUnimplementedA2SServer() {}

// UnsafeA2SServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to A2SServer will
// result in compilation errors.
type UnsafeA2SServer interface {
	mustEmbedUnimplementedA2SServer()
}

func RegisterA2SServer(s grpc.ServiceRegistrar, srv A2SServer) {
	s.RegisterService(&A2S_ServiceDesc, srv)
}

func _A2S_A2SQuery_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(A2SQueryReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(A2SServer).A2SQuery(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ncs.game.a2s.v1.A2S/A2SQuery",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(A2SServer).A2SQuery(ctx, req.(*A2SQueryReq))
	}
	return interceptor(ctx, in, info, handler)
}

// A2S_ServiceDesc is the grpc.ServiceDesc for A2S service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var A2S_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ncs.game.a2s.v1.A2S",
	HandlerType: (*A2SServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "A2SQuery",
			Handler:    _A2S_A2SQuery_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "app/game/a2s/api/grpc/v1/a2s.proto",
}
