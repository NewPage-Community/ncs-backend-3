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

// SkinClient is the client API for Skin service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SkinClient interface {
	GetSkins(ctx context.Context, in *GetSkinsReq, opts ...grpc.CallOption) (*GetSkinsResp, error)
	GetInfo(ctx context.Context, in *GetInfoReq, opts ...grpc.CallOption) (*GetInfoResp, error)
	SetUsedSkin(ctx context.Context, in *SetUsedSkinReq, opts ...grpc.CallOption) (*SetUsedSkinResp, error)
}

type skinClient struct {
	cc grpc.ClientConnInterface
}

func NewSkinClient(cc grpc.ClientConnInterface) SkinClient {
	return &skinClient{cc}
}

func (c *skinClient) GetSkins(ctx context.Context, in *GetSkinsReq, opts ...grpc.CallOption) (*GetSkinsResp, error) {
	out := new(GetSkinsResp)
	err := c.cc.Invoke(ctx, "/ncs.game.skin.v1.Skin/GetSkins", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *skinClient) GetInfo(ctx context.Context, in *GetInfoReq, opts ...grpc.CallOption) (*GetInfoResp, error) {
	out := new(GetInfoResp)
	err := c.cc.Invoke(ctx, "/ncs.game.skin.v1.Skin/GetInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *skinClient) SetUsedSkin(ctx context.Context, in *SetUsedSkinReq, opts ...grpc.CallOption) (*SetUsedSkinResp, error) {
	out := new(SetUsedSkinResp)
	err := c.cc.Invoke(ctx, "/ncs.game.skin.v1.Skin/SetUsedSkin", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SkinServer is the server API for Skin service.
// All implementations must embed UnimplementedSkinServer
// for forward compatibility
type SkinServer interface {
	GetSkins(context.Context, *GetSkinsReq) (*GetSkinsResp, error)
	GetInfo(context.Context, *GetInfoReq) (*GetInfoResp, error)
	SetUsedSkin(context.Context, *SetUsedSkinReq) (*SetUsedSkinResp, error)
	mustEmbedUnimplementedSkinServer()
}

// UnimplementedSkinServer must be embedded to have forward compatible implementations.
type UnimplementedSkinServer struct {
}

func (UnimplementedSkinServer) GetSkins(context.Context, *GetSkinsReq) (*GetSkinsResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSkins not implemented")
}
func (UnimplementedSkinServer) GetInfo(context.Context, *GetInfoReq) (*GetInfoResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetInfo not implemented")
}
func (UnimplementedSkinServer) SetUsedSkin(context.Context, *SetUsedSkinReq) (*SetUsedSkinResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetUsedSkin not implemented")
}
func (UnimplementedSkinServer) mustEmbedUnimplementedSkinServer() {}

// UnsafeSkinServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SkinServer will
// result in compilation errors.
type UnsafeSkinServer interface {
	mustEmbedUnimplementedSkinServer()
}

func RegisterSkinServer(s grpc.ServiceRegistrar, srv SkinServer) {
	s.RegisterService(&Skin_ServiceDesc, srv)
}

func _Skin_GetSkins_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSkinsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SkinServer).GetSkins(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ncs.game.skin.v1.Skin/GetSkins",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SkinServer).GetSkins(ctx, req.(*GetSkinsReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Skin_GetInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetInfoReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SkinServer).GetInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ncs.game.skin.v1.Skin/GetInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SkinServer).GetInfo(ctx, req.(*GetInfoReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Skin_SetUsedSkin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetUsedSkinReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SkinServer).SetUsedSkin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ncs.game.skin.v1.Skin/SetUsedSkin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SkinServer).SetUsedSkin(ctx, req.(*SetUsedSkinReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Skin_ServiceDesc is the grpc.ServiceDesc for Skin service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Skin_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ncs.game.skin.v1.Skin",
	HandlerType: (*SkinServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetSkins",
			Handler:    _Skin_GetSkins_Handler,
		},
		{
			MethodName: "GetInfo",
			Handler:    _Skin_GetInfo_Handler,
		},
		{
			MethodName: "SetUsedSkin",
			Handler:    _Skin_SetUsedSkin_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "app/game/skin/api/grpc/v1/skin.proto",
}
