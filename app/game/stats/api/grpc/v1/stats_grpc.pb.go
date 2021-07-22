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

// StatsClient is the client API for Stats service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type StatsClient interface {
	Get(ctx context.Context, in *GetReq, opts ...grpc.CallOption) (*GetResp, error)
	GetAll(ctx context.Context, in *GetAllReq, opts ...grpc.CallOption) (*GetAllResp, error)
	Gets(ctx context.Context, in *GetsReq, opts ...grpc.CallOption) (*GetsResp, error)
	Set(ctx context.Context, in *SetReq, opts ...grpc.CallOption) (*SetResp, error)
	Sets(ctx context.Context, in *SetsReq, opts ...grpc.CallOption) (*SetResp, error)
	Incr(ctx context.Context, in *IncrReq, opts ...grpc.CallOption) (*IncrResp, error)
	Incrs(ctx context.Context, in *IncrsReq, opts ...grpc.CallOption) (*IncrResp, error)
}

type statsClient struct {
	cc grpc.ClientConnInterface
}

func NewStatsClient(cc grpc.ClientConnInterface) StatsClient {
	return &statsClient{cc}
}

func (c *statsClient) Get(ctx context.Context, in *GetReq, opts ...grpc.CallOption) (*GetResp, error) {
	out := new(GetResp)
	err := c.cc.Invoke(ctx, "/ncs.game.stats.v1.Stats/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *statsClient) GetAll(ctx context.Context, in *GetAllReq, opts ...grpc.CallOption) (*GetAllResp, error) {
	out := new(GetAllResp)
	err := c.cc.Invoke(ctx, "/ncs.game.stats.v1.Stats/GetAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *statsClient) Gets(ctx context.Context, in *GetsReq, opts ...grpc.CallOption) (*GetsResp, error) {
	out := new(GetsResp)
	err := c.cc.Invoke(ctx, "/ncs.game.stats.v1.Stats/Gets", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *statsClient) Set(ctx context.Context, in *SetReq, opts ...grpc.CallOption) (*SetResp, error) {
	out := new(SetResp)
	err := c.cc.Invoke(ctx, "/ncs.game.stats.v1.Stats/Set", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *statsClient) Sets(ctx context.Context, in *SetsReq, opts ...grpc.CallOption) (*SetResp, error) {
	out := new(SetResp)
	err := c.cc.Invoke(ctx, "/ncs.game.stats.v1.Stats/Sets", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *statsClient) Incr(ctx context.Context, in *IncrReq, opts ...grpc.CallOption) (*IncrResp, error) {
	out := new(IncrResp)
	err := c.cc.Invoke(ctx, "/ncs.game.stats.v1.Stats/Incr", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *statsClient) Incrs(ctx context.Context, in *IncrsReq, opts ...grpc.CallOption) (*IncrResp, error) {
	out := new(IncrResp)
	err := c.cc.Invoke(ctx, "/ncs.game.stats.v1.Stats/Incrs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StatsServer is the server API for Stats service.
// All implementations must embed UnimplementedStatsServer
// for forward compatibility
type StatsServer interface {
	Get(context.Context, *GetReq) (*GetResp, error)
	GetAll(context.Context, *GetAllReq) (*GetAllResp, error)
	Gets(context.Context, *GetsReq) (*GetsResp, error)
	Set(context.Context, *SetReq) (*SetResp, error)
	Sets(context.Context, *SetsReq) (*SetResp, error)
	Incr(context.Context, *IncrReq) (*IncrResp, error)
	Incrs(context.Context, *IncrsReq) (*IncrResp, error)
	mustEmbedUnimplementedStatsServer()
}

// UnimplementedStatsServer must be embedded to have forward compatible implementations.
type UnimplementedStatsServer struct {
}

func (UnimplementedStatsServer) Get(context.Context, *GetReq) (*GetResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedStatsServer) GetAll(context.Context, *GetAllReq) (*GetAllResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAll not implemented")
}
func (UnimplementedStatsServer) Gets(context.Context, *GetsReq) (*GetsResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Gets not implemented")
}
func (UnimplementedStatsServer) Set(context.Context, *SetReq) (*SetResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Set not implemented")
}
func (UnimplementedStatsServer) Sets(context.Context, *SetsReq) (*SetResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Sets not implemented")
}
func (UnimplementedStatsServer) Incr(context.Context, *IncrReq) (*IncrResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Incr not implemented")
}
func (UnimplementedStatsServer) Incrs(context.Context, *IncrsReq) (*IncrResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Incrs not implemented")
}
func (UnimplementedStatsServer) mustEmbedUnimplementedStatsServer() {}

// UnsafeStatsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to StatsServer will
// result in compilation errors.
type UnsafeStatsServer interface {
	mustEmbedUnimplementedStatsServer()
}

func RegisterStatsServer(s grpc.ServiceRegistrar, srv StatsServer) {
	s.RegisterService(&Stats_ServiceDesc, srv)
}

func _Stats_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StatsServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ncs.game.stats.v1.Stats/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StatsServer).Get(ctx, req.(*GetReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Stats_GetAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StatsServer).GetAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ncs.game.stats.v1.Stats/GetAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StatsServer).GetAll(ctx, req.(*GetAllReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Stats_Gets_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StatsServer).Gets(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ncs.game.stats.v1.Stats/Gets",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StatsServer).Gets(ctx, req.(*GetsReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Stats_Set_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StatsServer).Set(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ncs.game.stats.v1.Stats/Set",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StatsServer).Set(ctx, req.(*SetReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Stats_Sets_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StatsServer).Sets(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ncs.game.stats.v1.Stats/Sets",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StatsServer).Sets(ctx, req.(*SetsReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Stats_Incr_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IncrReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StatsServer).Incr(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ncs.game.stats.v1.Stats/Incr",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StatsServer).Incr(ctx, req.(*IncrReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Stats_Incrs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IncrsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StatsServer).Incrs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ncs.game.stats.v1.Stats/Incrs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StatsServer).Incrs(ctx, req.(*IncrsReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Stats_ServiceDesc is the grpc.ServiceDesc for Stats service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Stats_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ncs.game.stats.v1.Stats",
	HandlerType: (*StatsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _Stats_Get_Handler,
		},
		{
			MethodName: "GetAll",
			Handler:    _Stats_GetAll_Handler,
		},
		{
			MethodName: "Gets",
			Handler:    _Stats_Gets_Handler,
		},
		{
			MethodName: "Set",
			Handler:    _Stats_Set_Handler,
		},
		{
			MethodName: "Sets",
			Handler:    _Stats_Sets_Handler,
		},
		{
			MethodName: "Incr",
			Handler:    _Stats_Incr_Handler,
		},
		{
			MethodName: "Incrs",
			Handler:    _Stats_Incrs_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "app/game/stats/api/grpc/v1/stats.proto",
}