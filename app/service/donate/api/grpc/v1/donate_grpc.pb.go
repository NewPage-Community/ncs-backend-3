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

// DonateClient is the client API for Donate service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DonateClient interface {
	CreateDonate(ctx context.Context, in *CreateDonateReq, opts ...grpc.CallOption) (*CreateDonateResp, error)
	QueryDonate(ctx context.Context, in *QueryDonateReq, opts ...grpc.CallOption) (*QueryDonateResp, error)
	GetDonateList(ctx context.Context, in *GetDonateListReq, opts ...grpc.CallOption) (*GetDonateListResp, error)
	AddDonate(ctx context.Context, in *AddDonateReq, opts ...grpc.CallOption) (*AddDonateResp, error)
}

type donateClient struct {
	cc grpc.ClientConnInterface
}

func NewDonateClient(cc grpc.ClientConnInterface) DonateClient {
	return &donateClient{cc}
}

func (c *donateClient) CreateDonate(ctx context.Context, in *CreateDonateReq, opts ...grpc.CallOption) (*CreateDonateResp, error) {
	out := new(CreateDonateResp)
	err := c.cc.Invoke(ctx, "/ncs.service.donate.v1.Donate/CreateDonate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *donateClient) QueryDonate(ctx context.Context, in *QueryDonateReq, opts ...grpc.CallOption) (*QueryDonateResp, error) {
	out := new(QueryDonateResp)
	err := c.cc.Invoke(ctx, "/ncs.service.donate.v1.Donate/QueryDonate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *donateClient) GetDonateList(ctx context.Context, in *GetDonateListReq, opts ...grpc.CallOption) (*GetDonateListResp, error) {
	out := new(GetDonateListResp)
	err := c.cc.Invoke(ctx, "/ncs.service.donate.v1.Donate/GetDonateList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *donateClient) AddDonate(ctx context.Context, in *AddDonateReq, opts ...grpc.CallOption) (*AddDonateResp, error) {
	out := new(AddDonateResp)
	err := c.cc.Invoke(ctx, "/ncs.service.donate.v1.Donate/AddDonate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DonateServer is the server API for Donate service.
// All implementations must embed UnimplementedDonateServer
// for forward compatibility
type DonateServer interface {
	CreateDonate(context.Context, *CreateDonateReq) (*CreateDonateResp, error)
	QueryDonate(context.Context, *QueryDonateReq) (*QueryDonateResp, error)
	GetDonateList(context.Context, *GetDonateListReq) (*GetDonateListResp, error)
	AddDonate(context.Context, *AddDonateReq) (*AddDonateResp, error)
	mustEmbedUnimplementedDonateServer()
}

// UnimplementedDonateServer must be embedded to have forward compatible implementations.
type UnimplementedDonateServer struct {
}

func (UnimplementedDonateServer) CreateDonate(context.Context, *CreateDonateReq) (*CreateDonateResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateDonate not implemented")
}
func (UnimplementedDonateServer) QueryDonate(context.Context, *QueryDonateReq) (*QueryDonateResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryDonate not implemented")
}
func (UnimplementedDonateServer) GetDonateList(context.Context, *GetDonateListReq) (*GetDonateListResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDonateList not implemented")
}
func (UnimplementedDonateServer) AddDonate(context.Context, *AddDonateReq) (*AddDonateResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddDonate not implemented")
}
func (UnimplementedDonateServer) mustEmbedUnimplementedDonateServer() {}

// UnsafeDonateServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DonateServer will
// result in compilation errors.
type UnsafeDonateServer interface {
	mustEmbedUnimplementedDonateServer()
}

func RegisterDonateServer(s grpc.ServiceRegistrar, srv DonateServer) {
	s.RegisterService(&Donate_ServiceDesc, srv)
}

func _Donate_CreateDonate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateDonateReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DonateServer).CreateDonate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ncs.service.donate.v1.Donate/CreateDonate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DonateServer).CreateDonate(ctx, req.(*CreateDonateReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Donate_QueryDonate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryDonateReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DonateServer).QueryDonate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ncs.service.donate.v1.Donate/QueryDonate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DonateServer).QueryDonate(ctx, req.(*QueryDonateReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Donate_GetDonateList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDonateListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DonateServer).GetDonateList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ncs.service.donate.v1.Donate/GetDonateList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DonateServer).GetDonateList(ctx, req.(*GetDonateListReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Donate_AddDonate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddDonateReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DonateServer).AddDonate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ncs.service.donate.v1.Donate/AddDonate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DonateServer).AddDonate(ctx, req.(*AddDonateReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Donate_ServiceDesc is the grpc.ServiceDesc for Donate service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Donate_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ncs.service.donate.v1.Donate",
	HandlerType: (*DonateServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateDonate",
			Handler:    _Donate_CreateDonate_Handler,
		},
		{
			MethodName: "QueryDonate",
			Handler:    _Donate_QueryDonate_Handler,
		},
		{
			MethodName: "GetDonateList",
			Handler:    _Donate_GetDonateList_Handler,
		},
		{
			MethodName: "AddDonate",
			Handler:    _Donate_AddDonate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "app/service/donate/api/grpc/v1/donate.proto",
}
