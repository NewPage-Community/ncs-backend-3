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

// ItemsClient is the client API for Items service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ItemsClient interface {
	GetItems(ctx context.Context, in *GetItemsReq, opts ...grpc.CallOption) (*GetItemsResp, error)
}

type itemsClient struct {
	cc grpc.ClientConnInterface
}

func NewItemsClient(cc grpc.ClientConnInterface) ItemsClient {
	return &itemsClient{cc}
}

func (c *itemsClient) GetItems(ctx context.Context, in *GetItemsReq, opts ...grpc.CallOption) (*GetItemsResp, error) {
	out := new(GetItemsResp)
	err := c.cc.Invoke(ctx, "/ncs.service.backpack.items.v1.Items/GetItems", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ItemsServer is the server API for Items service.
// All implementations must embed UnimplementedItemsServer
// for forward compatibility
type ItemsServer interface {
	GetItems(context.Context, *GetItemsReq) (*GetItemsResp, error)
	mustEmbedUnimplementedItemsServer()
}

// UnimplementedItemsServer must be embedded to have forward compatible implementations.
type UnimplementedItemsServer struct {
}

func (UnimplementedItemsServer) GetItems(context.Context, *GetItemsReq) (*GetItemsResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetItems not implemented")
}
func (UnimplementedItemsServer) mustEmbedUnimplementedItemsServer() {}

// UnsafeItemsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ItemsServer will
// result in compilation errors.
type UnsafeItemsServer interface {
	mustEmbedUnimplementedItemsServer()
}

func RegisterItemsServer(s grpc.ServiceRegistrar, srv ItemsServer) {
	s.RegisterService(&Items_ServiceDesc, srv)
}

func _Items_GetItems_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetItemsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ItemsServer).GetItems(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ncs.service.backpack.items.v1.Items/GetItems",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ItemsServer).GetItems(ctx, req.(*GetItemsReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Items_ServiceDesc is the grpc.ServiceDesc for Items service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Items_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ncs.service.backpack.items.v1.Items",
	HandlerType: (*ItemsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetItems",
			Handler:    _Items_GetItems_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "app/service/backpack/items/api/grpc/v1/items.proto",
}
