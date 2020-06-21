// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: app/user/sign/api/grpc/sign.proto

package grpc

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
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

type Info struct {
	Uid                  int64    `protobuf:"varint,1,opt,name=uid,proto3" json:"uid,omitempty"`
	SignTime             int64    `protobuf:"varint,2,opt,name=sign_time,json=signTime,proto3" json:"sign_time,omitempty"`
	SignDays             int32    `protobuf:"varint,3,opt,name=sign_days,json=signDays,proto3" json:"sign_days,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Info) Reset()         { *m = Info{} }
func (m *Info) String() string { return proto.CompactTextString(m) }
func (*Info) ProtoMessage()    {}
func (*Info) Descriptor() ([]byte, []int) {
	return fileDescriptor_d3e46b5a18f2e404, []int{0}
}
func (m *Info) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Info) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Info.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Info) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Info.Merge(m, src)
}
func (m *Info) XXX_Size() int {
	return m.Size()
}
func (m *Info) XXX_DiscardUnknown() {
	xxx_messageInfo_Info.DiscardUnknown(m)
}

var xxx_messageInfo_Info proto.InternalMessageInfo

func (m *Info) GetUid() int64 {
	if m != nil {
		return m.Uid
	}
	return 0
}

func (m *Info) GetSignTime() int64 {
	if m != nil {
		return m.SignTime
	}
	return 0
}

func (m *Info) GetSignDays() int32 {
	if m != nil {
		return m.SignDays
	}
	return 0
}

type InfoReq struct {
	Uid                  int64    `protobuf:"varint,1,opt,name=uid,proto3" json:"uid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InfoReq) Reset()         { *m = InfoReq{} }
func (m *InfoReq) String() string { return proto.CompactTextString(m) }
func (*InfoReq) ProtoMessage()    {}
func (*InfoReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_d3e46b5a18f2e404, []int{1}
}
func (m *InfoReq) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *InfoReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_InfoReq.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *InfoReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InfoReq.Merge(m, src)
}
func (m *InfoReq) XXX_Size() int {
	return m.Size()
}
func (m *InfoReq) XXX_DiscardUnknown() {
	xxx_messageInfo_InfoReq.DiscardUnknown(m)
}

var xxx_messageInfo_InfoReq proto.InternalMessageInfo

func (m *InfoReq) GetUid() int64 {
	if m != nil {
		return m.Uid
	}
	return 0
}

type InfoResp struct {
	Info                 *Info    `protobuf:"bytes,1,opt,name=info,proto3" json:"info,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InfoResp) Reset()         { *m = InfoResp{} }
func (m *InfoResp) String() string { return proto.CompactTextString(m) }
func (*InfoResp) ProtoMessage()    {}
func (*InfoResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_d3e46b5a18f2e404, []int{2}
}
func (m *InfoResp) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *InfoResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_InfoResp.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *InfoResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InfoResp.Merge(m, src)
}
func (m *InfoResp) XXX_Size() int {
	return m.Size()
}
func (m *InfoResp) XXX_DiscardUnknown() {
	xxx_messageInfo_InfoResp.DiscardUnknown(m)
}

var xxx_messageInfo_InfoResp proto.InternalMessageInfo

func (m *InfoResp) GetInfo() *Info {
	if m != nil {
		return m.Info
	}
	return nil
}

type SignReq struct {
	Uid                  int64    `protobuf:"varint,1,opt,name=uid,proto3" json:"uid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SignReq) Reset()         { *m = SignReq{} }
func (m *SignReq) String() string { return proto.CompactTextString(m) }
func (*SignReq) ProtoMessage()    {}
func (*SignReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_d3e46b5a18f2e404, []int{3}
}
func (m *SignReq) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SignReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SignReq.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SignReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SignReq.Merge(m, src)
}
func (m *SignReq) XXX_Size() int {
	return m.Size()
}
func (m *SignReq) XXX_DiscardUnknown() {
	xxx_messageInfo_SignReq.DiscardUnknown(m)
}

var xxx_messageInfo_SignReq proto.InternalMessageInfo

func (m *SignReq) GetUid() int64 {
	if m != nil {
		return m.Uid
	}
	return 0
}

type SignResp struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SignResp) Reset()         { *m = SignResp{} }
func (m *SignResp) String() string { return proto.CompactTextString(m) }
func (*SignResp) ProtoMessage()    {}
func (*SignResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_d3e46b5a18f2e404, []int{4}
}
func (m *SignResp) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SignResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SignResp.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SignResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SignResp.Merge(m, src)
}
func (m *SignResp) XXX_Size() int {
	return m.Size()
}
func (m *SignResp) XXX_DiscardUnknown() {
	xxx_messageInfo_SignResp.DiscardUnknown(m)
}

var xxx_messageInfo_SignResp proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Info)(nil), "ncs.user.sign.Info")
	proto.RegisterType((*InfoReq)(nil), "ncs.user.sign.InfoReq")
	proto.RegisterType((*InfoResp)(nil), "ncs.user.sign.InfoResp")
	proto.RegisterType((*SignReq)(nil), "ncs.user.sign.SignReq")
	proto.RegisterType((*SignResp)(nil), "ncs.user.sign.SignResp")
}

func init() { proto.RegisterFile("app/user/sign/api/grpc/sign.proto", fileDescriptor_d3e46b5a18f2e404) }

var fileDescriptor_d3e46b5a18f2e404 = []byte{
	// 303 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x4c, 0x2c, 0x28, 0xd0,
	0x2f, 0x2d, 0x4e, 0x2d, 0xd2, 0x2f, 0xce, 0x4c, 0xcf, 0xd3, 0x4f, 0x2c, 0xc8, 0xd4, 0x4f, 0x2f,
	0x2a, 0x48, 0x06, 0xf3, 0xf4, 0x0a, 0x8a, 0xf2, 0x4b, 0xf2, 0x85, 0x78, 0xf3, 0x92, 0x8b, 0xf5,
	0x40, 0x4a, 0xf4, 0x40, 0x82, 0x52, 0x32, 0xe9, 0xf9, 0xf9, 0xe9, 0x39, 0xa9, 0x60, 0xa5, 0x89,
	0x79, 0x79, 0xf9, 0x25, 0x89, 0x25, 0x99, 0xf9, 0x79, 0xc5, 0x10, 0xc5, 0x4a, 0x41, 0x5c, 0x2c,
	0x9e, 0x79, 0x69, 0xf9, 0x42, 0x02, 0x5c, 0xcc, 0xa5, 0x99, 0x29, 0x12, 0x8c, 0x0a, 0x8c, 0x1a,
	0xcc, 0x41, 0x20, 0xa6, 0x90, 0x34, 0x17, 0x27, 0x48, 0x7f, 0x7c, 0x49, 0x66, 0x6e, 0xaa, 0x04,
	0x13, 0x58, 0x9c, 0x03, 0x24, 0x10, 0x92, 0x99, 0x9b, 0x0a, 0x97, 0x4c, 0x49, 0xac, 0x2c, 0x96,
	0x60, 0x56, 0x60, 0xd4, 0x60, 0x85, 0x48, 0xba, 0x24, 0x56, 0x16, 0x2b, 0x49, 0x73, 0xb1, 0x83,
	0xcc, 0x0c, 0x4a, 0x2d, 0xc4, 0x34, 0x56, 0xc9, 0x98, 0x8b, 0x03, 0x22, 0x59, 0x5c, 0x20, 0xa4,
	0xce, 0xc5, 0x92, 0x99, 0x97, 0x96, 0x0f, 0x96, 0xe6, 0x36, 0x12, 0xd6, 0x43, 0x71, 0xb8, 0x1e,
	0x58, 0x19, 0x58, 0x01, 0xc8, 0xc4, 0xe0, 0xcc, 0xf4, 0x3c, 0xec, 0x26, 0x72, 0x71, 0x71, 0x40,
	0x24, 0x8b, 0x0b, 0x8c, 0x56, 0x32, 0x72, 0xb1, 0x80, 0x38, 0x42, 0x81, 0x50, 0x7f, 0x89, 0x61,
	0x33, 0x34, 0xb5, 0x50, 0x4a, 0x1c, 0xab, 0x78, 0x71, 0x81, 0x92, 0x44, 0xd3, 0xe5, 0x27, 0x93,
	0x99, 0x84, 0x84, 0x04, 0x90, 0x42, 0xb9, 0xba, 0x34, 0x33, 0xa5, 0x56, 0xc8, 0x0f, 0x6a, 0x34,
	0xba, 0x91, 0x50, 0x97, 0x61, 0x18, 0x09, 0x73, 0x94, 0x92, 0x28, 0xd8, 0x48, 0x7e, 0x25, 0x2e,
	0x84, 0x91, 0x56, 0x8c, 0x5a, 0x4e, 0x62, 0x27, 0x1e, 0xc9, 0x31, 0x5e, 0x78, 0x24, 0xc7, 0xf8,
	0xe0, 0x91, 0x1c, 0xe3, 0x8c, 0xc7, 0x72, 0x0c, 0x51, 0x2c, 0xa0, 0xa8, 0x4c, 0x62, 0x03, 0xc7,
	0x8c, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0xe6, 0x13, 0x3b, 0xa6, 0xeb, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// SignClient is the client API for Sign service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SignClient interface {
	Info(ctx context.Context, in *InfoReq, opts ...grpc.CallOption) (*InfoResp, error)
	Sign(ctx context.Context, in *SignReq, opts ...grpc.CallOption) (*SignResp, error)
}

type signClient struct {
	cc *grpc.ClientConn
}

func NewSignClient(cc *grpc.ClientConn) SignClient {
	return &signClient{cc}
}

func (c *signClient) Info(ctx context.Context, in *InfoReq, opts ...grpc.CallOption) (*InfoResp, error) {
	out := new(InfoResp)
	err := c.cc.Invoke(ctx, "/ncs.user.sign.Sign/Info", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *signClient) Sign(ctx context.Context, in *SignReq, opts ...grpc.CallOption) (*SignResp, error) {
	out := new(SignResp)
	err := c.cc.Invoke(ctx, "/ncs.user.sign.Sign/Sign", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SignServer is the server API for Sign service.
type SignServer interface {
	Info(context.Context, *InfoReq) (*InfoResp, error)
	Sign(context.Context, *SignReq) (*SignResp, error)
}

// UnimplementedSignServer can be embedded to have forward compatible implementations.
type UnimplementedSignServer struct {
}

func (*UnimplementedSignServer) Info(ctx context.Context, req *InfoReq) (*InfoResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Info not implemented")
}
func (*UnimplementedSignServer) Sign(ctx context.Context, req *SignReq) (*SignResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Sign not implemented")
}

func RegisterSignServer(s *grpc.Server, srv SignServer) {
	s.RegisterService(&_Sign_serviceDesc, srv)
}

func _Sign_Info_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InfoReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SignServer).Info(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ncs.user.sign.Sign/Info",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SignServer).Info(ctx, req.(*InfoReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Sign_Sign_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SignServer).Sign(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ncs.user.sign.Sign/Sign",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SignServer).Sign(ctx, req.(*SignReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _Sign_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ncs.user.sign.Sign",
	HandlerType: (*SignServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Info",
			Handler:    _Sign_Info_Handler,
		},
		{
			MethodName: "Sign",
			Handler:    _Sign_Sign_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "app/user/sign/api/grpc/sign.proto",
}

func (m *Info) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Info) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Info) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.SignDays != 0 {
		i = encodeVarintSign(dAtA, i, uint64(m.SignDays))
		i--
		dAtA[i] = 0x18
	}
	if m.SignTime != 0 {
		i = encodeVarintSign(dAtA, i, uint64(m.SignTime))
		i--
		dAtA[i] = 0x10
	}
	if m.Uid != 0 {
		i = encodeVarintSign(dAtA, i, uint64(m.Uid))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *InfoReq) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *InfoReq) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *InfoReq) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Uid != 0 {
		i = encodeVarintSign(dAtA, i, uint64(m.Uid))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *InfoResp) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *InfoResp) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *InfoResp) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Info != nil {
		{
			size, err := m.Info.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintSign(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *SignReq) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SignReq) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SignReq) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Uid != 0 {
		i = encodeVarintSign(dAtA, i, uint64(m.Uid))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *SignResp) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SignResp) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SignResp) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	return len(dAtA) - i, nil
}

func encodeVarintSign(dAtA []byte, offset int, v uint64) int {
	offset -= sovSign(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Info) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Uid != 0 {
		n += 1 + sovSign(uint64(m.Uid))
	}
	if m.SignTime != 0 {
		n += 1 + sovSign(uint64(m.SignTime))
	}
	if m.SignDays != 0 {
		n += 1 + sovSign(uint64(m.SignDays))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *InfoReq) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Uid != 0 {
		n += 1 + sovSign(uint64(m.Uid))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *InfoResp) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Info != nil {
		l = m.Info.Size()
		n += 1 + l + sovSign(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *SignReq) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Uid != 0 {
		n += 1 + sovSign(uint64(m.Uid))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *SignResp) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovSign(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozSign(x uint64) (n int) {
	return sovSign(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Info) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSign
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Info: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Info: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Uid", wireType)
			}
			m.Uid = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSign
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Uid |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SignTime", wireType)
			}
			m.SignTime = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSign
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SignTime |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SignDays", wireType)
			}
			m.SignDays = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSign
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SignDays |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipSign(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthSign
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthSign
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *InfoReq) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSign
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: InfoReq: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: InfoReq: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Uid", wireType)
			}
			m.Uid = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSign
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Uid |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipSign(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthSign
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthSign
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *InfoResp) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSign
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: InfoResp: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: InfoResp: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Info", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSign
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthSign
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthSign
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Info == nil {
				m.Info = &Info{}
			}
			if err := m.Info.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSign(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthSign
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthSign
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *SignReq) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSign
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: SignReq: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SignReq: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Uid", wireType)
			}
			m.Uid = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSign
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Uid |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipSign(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthSign
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthSign
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *SignResp) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSign
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: SignResp: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SignResp: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipSign(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthSign
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthSign
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipSign(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowSign
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowSign
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowSign
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthSign
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupSign
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthSign
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthSign        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowSign          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupSign = fmt.Errorf("proto: unexpected end of group")
)