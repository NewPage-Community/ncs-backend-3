// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: app/game/cvar/api/grpc/cvar.proto

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

type CVarInfo struct {
	Key                  string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value                string   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CVarInfo) Reset()         { *m = CVarInfo{} }
func (m *CVarInfo) String() string { return proto.CompactTextString(m) }
func (*CVarInfo) ProtoMessage()    {}
func (*CVarInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_d69c6fafa408c1f2, []int{0}
}
func (m *CVarInfo) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CVarInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CVarInfo.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CVarInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CVarInfo.Merge(m, src)
}
func (m *CVarInfo) XXX_Size() int {
	return m.Size()
}
func (m *CVarInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_CVarInfo.DiscardUnknown(m)
}

var xxx_messageInfo_CVarInfo proto.InternalMessageInfo

func (m *CVarInfo) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *CVarInfo) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type UpdateCVarsReq struct {
	GameId               int32    `protobuf:"varint,1,opt,name=game_id,json=gameId,proto3" json:"game_id,omitempty"`
	ModId                int32    `protobuf:"varint,2,opt,name=mod_id,json=modId,proto3" json:"mod_id,omitempty"`
	ServerId             int32    `protobuf:"varint,3,opt,name=server_id,json=serverId,proto3" json:"server_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateCVarsReq) Reset()         { *m = UpdateCVarsReq{} }
func (m *UpdateCVarsReq) String() string { return proto.CompactTextString(m) }
func (*UpdateCVarsReq) ProtoMessage()    {}
func (*UpdateCVarsReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_d69c6fafa408c1f2, []int{1}
}
func (m *UpdateCVarsReq) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *UpdateCVarsReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_UpdateCVarsReq.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *UpdateCVarsReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateCVarsReq.Merge(m, src)
}
func (m *UpdateCVarsReq) XXX_Size() int {
	return m.Size()
}
func (m *UpdateCVarsReq) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateCVarsReq.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateCVarsReq proto.InternalMessageInfo

func (m *UpdateCVarsReq) GetGameId() int32 {
	if m != nil {
		return m.GameId
	}
	return 0
}

func (m *UpdateCVarsReq) GetModId() int32 {
	if m != nil {
		return m.ModId
	}
	return 0
}

func (m *UpdateCVarsReq) GetServerId() int32 {
	if m != nil {
		return m.ServerId
	}
	return 0
}

type UpdateCVarsResp struct {
	Cvars                []*CVarInfo `protobuf:"bytes,1,rep,name=cvars,proto3" json:"cvars,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *UpdateCVarsResp) Reset()         { *m = UpdateCVarsResp{} }
func (m *UpdateCVarsResp) String() string { return proto.CompactTextString(m) }
func (*UpdateCVarsResp) ProtoMessage()    {}
func (*UpdateCVarsResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_d69c6fafa408c1f2, []int{2}
}
func (m *UpdateCVarsResp) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *UpdateCVarsResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_UpdateCVarsResp.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *UpdateCVarsResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateCVarsResp.Merge(m, src)
}
func (m *UpdateCVarsResp) XXX_Size() int {
	return m.Size()
}
func (m *UpdateCVarsResp) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateCVarsResp.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateCVarsResp proto.InternalMessageInfo

func (m *UpdateCVarsResp) GetCvars() []*CVarInfo {
	if m != nil {
		return m.Cvars
	}
	return nil
}

func init() {
	proto.RegisterType((*CVarInfo)(nil), "ncs.game.cvar.CVarInfo")
	proto.RegisterType((*UpdateCVarsReq)(nil), "ncs.game.cvar.UpdateCVarsReq")
	proto.RegisterType((*UpdateCVarsResp)(nil), "ncs.game.cvar.UpdateCVarsResp")
}

func init() { proto.RegisterFile("app/game/cvar/api/grpc/cvar.proto", fileDescriptor_d69c6fafa408c1f2) }

var fileDescriptor_d69c6fafa408c1f2 = []byte{
	// 309 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x4c, 0x2c, 0x28, 0xd0,
	0x4f, 0x4f, 0xcc, 0x4d, 0xd5, 0x4f, 0x2e, 0x4b, 0x2c, 0xd2, 0x4f, 0x2c, 0xc8, 0xd4, 0x4f, 0x2f,
	0x2a, 0x48, 0x06, 0xf3, 0xf4, 0x0a, 0x8a, 0xf2, 0x4b, 0xf2, 0x85, 0x78, 0xf3, 0x92, 0x8b, 0xf5,
	0x40, 0x4a, 0xf4, 0x40, 0x82, 0x52, 0x32, 0xe9, 0xf9, 0xf9, 0xe9, 0x39, 0xa9, 0x60, 0xa5, 0x89,
	0x79, 0x79, 0xf9, 0x25, 0x89, 0x25, 0x99, 0xf9, 0x79, 0xc5, 0x10, 0xc5, 0x4a, 0x46, 0x5c, 0x1c,
	0xce, 0x61, 0x89, 0x45, 0x9e, 0x79, 0x69, 0xf9, 0x42, 0x02, 0x5c, 0xcc, 0xd9, 0xa9, 0x95, 0x12,
	0x8c, 0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0x20, 0xa6, 0x90, 0x08, 0x17, 0x6b, 0x59, 0x62, 0x4e, 0x69,
	0xaa, 0x04, 0x13, 0x58, 0x0c, 0xc2, 0x51, 0x8a, 0xe5, 0xe2, 0x0b, 0x2d, 0x48, 0x49, 0x2c, 0x49,
	0x05, 0xe9, 0x2c, 0x0e, 0x4a, 0x2d, 0x14, 0x12, 0xe7, 0x62, 0x07, 0x59, 0x18, 0x9f, 0x99, 0x02,
	0xd6, 0xcd, 0x1a, 0xc4, 0x06, 0xe2, 0x7a, 0xa6, 0x08, 0x89, 0x72, 0xb1, 0xe5, 0xe6, 0xa7, 0x80,
	0xc4, 0x99, 0xc0, 0xe2, 0xac, 0xb9, 0xf9, 0x29, 0x9e, 0x29, 0x42, 0xd2, 0x5c, 0x9c, 0xc5, 0xa9,
	0x45, 0x65, 0xa9, 0x45, 0x20, 0x19, 0x66, 0xb0, 0x0c, 0x07, 0x44, 0xc0, 0x33, 0x45, 0xc9, 0x81,
	0x8b, 0x1f, 0xc5, 0xf8, 0xe2, 0x02, 0x21, 0x5d, 0x2e, 0x56, 0x90, 0x5f, 0x8a, 0x25, 0x18, 0x15,
	0x98, 0x35, 0xb8, 0x8d, 0xc4, 0xf5, 0x50, 0xbc, 0xa8, 0x07, 0xf3, 0x41, 0x10, 0x44, 0x95, 0x51,
	0x11, 0x17, 0x0b, 0x48, 0x48, 0x28, 0x8b, 0x8b, 0x1b, 0xc9, 0x24, 0x21, 0x59, 0x34, 0x6d, 0xa8,
	0x9e, 0x90, 0x92, 0xc3, 0x27, 0x5d, 0x5c, 0xa0, 0x24, 0xd3, 0x74, 0xf9, 0xc9, 0x64, 0x26, 0x31,
	0x25, 0x41, 0xa4, 0xf0, 0x2f, 0x05, 0xab, 0xb1, 0x62, 0xd4, 0x72, 0x12, 0x3b, 0xf1, 0x48, 0x8e,
	0xf1, 0xc2, 0x23, 0x39, 0xc6, 0x07, 0x8f, 0xe4, 0x18, 0x67, 0x3c, 0x96, 0x63, 0x88, 0x62, 0x01,
	0x45, 0x4c, 0x12, 0x1b, 0x38, 0x9c, 0x8d, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0x76, 0x38, 0x8d,
	0xa5, 0xb9, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CVarClient is the client API for CVar service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CVarClient interface {
	UpdateCVars(ctx context.Context, in *UpdateCVarsReq, opts ...grpc.CallOption) (*UpdateCVarsResp, error)
}

type cVarClient struct {
	cc *grpc.ClientConn
}

func NewCVarClient(cc *grpc.ClientConn) CVarClient {
	return &cVarClient{cc}
}

func (c *cVarClient) UpdateCVars(ctx context.Context, in *UpdateCVarsReq, opts ...grpc.CallOption) (*UpdateCVarsResp, error) {
	out := new(UpdateCVarsResp)
	err := c.cc.Invoke(ctx, "/ncs.game.cvar.CVar/UpdateCVars", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CVarServer is the server API for CVar service.
type CVarServer interface {
	UpdateCVars(context.Context, *UpdateCVarsReq) (*UpdateCVarsResp, error)
}

// UnimplementedCVarServer can be embedded to have forward compatible implementations.
type UnimplementedCVarServer struct {
}

func (*UnimplementedCVarServer) UpdateCVars(ctx context.Context, req *UpdateCVarsReq) (*UpdateCVarsResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateCVars not implemented")
}

func RegisterCVarServer(s *grpc.Server, srv CVarServer) {
	s.RegisterService(&_CVar_serviceDesc, srv)
}

func _CVar_UpdateCVars_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateCVarsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CVarServer).UpdateCVars(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ncs.game.cvar.CVar/UpdateCVars",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CVarServer).UpdateCVars(ctx, req.(*UpdateCVarsReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _CVar_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ncs.game.cvar.CVar",
	HandlerType: (*CVarServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UpdateCVars",
			Handler:    _CVar_UpdateCVars_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "app/game/cvar/api/grpc/cvar.proto",
}

func (m *CVarInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CVarInfo) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CVarInfo) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Value) > 0 {
		i -= len(m.Value)
		copy(dAtA[i:], m.Value)
		i = encodeVarintCvar(dAtA, i, uint64(len(m.Value)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Key) > 0 {
		i -= len(m.Key)
		copy(dAtA[i:], m.Key)
		i = encodeVarintCvar(dAtA, i, uint64(len(m.Key)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *UpdateCVarsReq) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *UpdateCVarsReq) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *UpdateCVarsReq) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.ServerId != 0 {
		i = encodeVarintCvar(dAtA, i, uint64(m.ServerId))
		i--
		dAtA[i] = 0x18
	}
	if m.ModId != 0 {
		i = encodeVarintCvar(dAtA, i, uint64(m.ModId))
		i--
		dAtA[i] = 0x10
	}
	if m.GameId != 0 {
		i = encodeVarintCvar(dAtA, i, uint64(m.GameId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *UpdateCVarsResp) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *UpdateCVarsResp) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *UpdateCVarsResp) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Cvars) > 0 {
		for iNdEx := len(m.Cvars) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Cvars[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintCvar(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintCvar(dAtA []byte, offset int, v uint64) int {
	offset -= sovCvar(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *CVarInfo) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Key)
	if l > 0 {
		n += 1 + l + sovCvar(uint64(l))
	}
	l = len(m.Value)
	if l > 0 {
		n += 1 + l + sovCvar(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *UpdateCVarsReq) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.GameId != 0 {
		n += 1 + sovCvar(uint64(m.GameId))
	}
	if m.ModId != 0 {
		n += 1 + sovCvar(uint64(m.ModId))
	}
	if m.ServerId != 0 {
		n += 1 + sovCvar(uint64(m.ServerId))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *UpdateCVarsResp) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Cvars) > 0 {
		for _, e := range m.Cvars {
			l = e.Size()
			n += 1 + l + sovCvar(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovCvar(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozCvar(x uint64) (n int) {
	return sovCvar(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *CVarInfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCvar
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
			return fmt.Errorf("proto: CVarInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CVarInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Key", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCvar
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthCvar
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCvar
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Key = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCvar
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthCvar
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCvar
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Value = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCvar(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthCvar
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthCvar
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
func (m *UpdateCVarsReq) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCvar
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
			return fmt.Errorf("proto: UpdateCVarsReq: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: UpdateCVarsReq: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field GameId", wireType)
			}
			m.GameId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCvar
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.GameId |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ModId", wireType)
			}
			m.ModId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCvar
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ModId |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ServerId", wireType)
			}
			m.ServerId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCvar
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ServerId |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipCvar(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthCvar
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthCvar
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
func (m *UpdateCVarsResp) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCvar
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
			return fmt.Errorf("proto: UpdateCVarsResp: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: UpdateCVarsResp: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Cvars", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCvar
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
				return ErrInvalidLengthCvar
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCvar
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Cvars = append(m.Cvars, &CVarInfo{})
			if err := m.Cvars[len(m.Cvars)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCvar(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthCvar
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthCvar
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
func skipCvar(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowCvar
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
					return 0, ErrIntOverflowCvar
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
					return 0, ErrIntOverflowCvar
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
				return 0, ErrInvalidLengthCvar
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupCvar
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthCvar
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthCvar        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowCvar          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupCvar = fmt.Errorf("proto: unexpected end of group")
)