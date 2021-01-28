// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: app/bot/kaiheila/api/grpc/v1/kaiheila.proto

package v1

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

type SendMessageReq struct {
	Type                 int32    `protobuf:"varint,1,opt,name=type,proto3" json:"type,omitempty"`
	ChannelId            string   `protobuf:"bytes,2,opt,name=channel_id,json=channelId,proto3" json:"channel_id,omitempty"`
	Content              string   `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"`
	Quote                string   `protobuf:"bytes,4,opt,name=quote,proto3" json:"quote,omitempty"`
	Nonce                string   `protobuf:"bytes,5,opt,name=nonce,proto3" json:"nonce,omitempty"`
	TempTargetId         string   `protobuf:"bytes,6,opt,name=temp_target_id,json=tempTargetId,proto3" json:"temp_target_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SendMessageReq) Reset()         { *m = SendMessageReq{} }
func (m *SendMessageReq) String() string { return proto.CompactTextString(m) }
func (*SendMessageReq) ProtoMessage()    {}
func (*SendMessageReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_b0f486f7c8ed4f23, []int{0}
}
func (m *SendMessageReq) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SendMessageReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SendMessageReq.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SendMessageReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SendMessageReq.Merge(m, src)
}
func (m *SendMessageReq) XXX_Size() int {
	return m.Size()
}
func (m *SendMessageReq) XXX_DiscardUnknown() {
	xxx_messageInfo_SendMessageReq.DiscardUnknown(m)
}

var xxx_messageInfo_SendMessageReq proto.InternalMessageInfo

func (m *SendMessageReq) GetType() int32 {
	if m != nil {
		return m.Type
	}
	return 0
}

func (m *SendMessageReq) GetChannelId() string {
	if m != nil {
		return m.ChannelId
	}
	return ""
}

func (m *SendMessageReq) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

func (m *SendMessageReq) GetQuote() string {
	if m != nil {
		return m.Quote
	}
	return ""
}

func (m *SendMessageReq) GetNonce() string {
	if m != nil {
		return m.Nonce
	}
	return ""
}

func (m *SendMessageReq) GetTempTargetId() string {
	if m != nil {
		return m.TempTargetId
	}
	return ""
}

type SendMessageResp struct {
	MsgId                string   `protobuf:"bytes,1,opt,name=msg_id,json=msgId,proto3" json:"msg_id,omitempty"`
	MsgTimestamp         int64    `protobuf:"varint,2,opt,name=msg_timestamp,json=msgTimestamp,proto3" json:"msg_timestamp,omitempty"`
	Nonce                string   `protobuf:"bytes,3,opt,name=nonce,proto3" json:"nonce,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SendMessageResp) Reset()         { *m = SendMessageResp{} }
func (m *SendMessageResp) String() string { return proto.CompactTextString(m) }
func (*SendMessageResp) ProtoMessage()    {}
func (*SendMessageResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_b0f486f7c8ed4f23, []int{1}
}
func (m *SendMessageResp) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SendMessageResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SendMessageResp.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SendMessageResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SendMessageResp.Merge(m, src)
}
func (m *SendMessageResp) XXX_Size() int {
	return m.Size()
}
func (m *SendMessageResp) XXX_DiscardUnknown() {
	xxx_messageInfo_SendMessageResp.DiscardUnknown(m)
}

var xxx_messageInfo_SendMessageResp proto.InternalMessageInfo

func (m *SendMessageResp) GetMsgId() string {
	if m != nil {
		return m.MsgId
	}
	return ""
}

func (m *SendMessageResp) GetMsgTimestamp() int64 {
	if m != nil {
		return m.MsgTimestamp
	}
	return 0
}

func (m *SendMessageResp) GetNonce() string {
	if m != nil {
		return m.Nonce
	}
	return ""
}

func init() {
	proto.RegisterType((*SendMessageReq)(nil), "ncs.bot.kaiheila.v1.SendMessageReq")
	proto.RegisterType((*SendMessageResp)(nil), "ncs.bot.kaiheila.v1.SendMessageResp")
}

func init() {
	proto.RegisterFile("app/bot/kaiheila/api/grpc/v1/kaiheila.proto", fileDescriptor_b0f486f7c8ed4f23)
}

var fileDescriptor_b0f486f7c8ed4f23 = []byte{
	// 336 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x91, 0xdf, 0x4e, 0xc2, 0x30,
	0x14, 0xc6, 0xad, 0xfc, 0x51, 0x1a, 0xc4, 0xa4, 0x6a, 0xb2, 0x10, 0x5c, 0x08, 0x70, 0x41, 0x62,
	0xb2, 0x05, 0x7d, 0x03, 0xbd, 0x22, 0x86, 0x9b, 0xc9, 0x95, 0x5e, 0x90, 0xb2, 0x9d, 0x94, 0x45,
	0xd6, 0x16, 0x7a, 0x24, 0xf1, 0x4d, 0x7c, 0x0b, 0x5f, 0xc3, 0x4b, 0x1f, 0xc1, 0xcc, 0x17, 0x31,
	0x6b, 0x01, 0x25, 0x31, 0x7a, 0xb7, 0xef, 0x77, 0xf6, 0xb5, 0x5f, 0xbf, 0x43, 0x2f, 0xb8, 0xd6,
	0xe1, 0x54, 0x61, 0xf8, 0xc8, 0xd3, 0x19, 0xa4, 0x73, 0x1e, 0x72, 0x9d, 0x86, 0x62, 0xa9, 0xe3,
	0x70, 0x35, 0xd8, 0xc2, 0x40, 0x2f, 0x15, 0x2a, 0x76, 0x22, 0x63, 0x13, 0x4c, 0x15, 0x06, 0x5b,
	0xbe, 0x1a, 0x34, 0x5b, 0x42, 0x29, 0x31, 0x07, 0xeb, 0xe3, 0x52, 0x2a, 0xe4, 0x98, 0x2a, 0x69,
	0x9c, 0xa5, 0xf3, 0x4a, 0x68, 0xe3, 0x0e, 0x64, 0x32, 0x02, 0x63, 0xb8, 0x80, 0x08, 0x16, 0x8c,
	0xd1, 0x32, 0x3e, 0x6b, 0xf0, 0x48, 0x9b, 0xf4, 0x2b, 0x91, 0xfd, 0x66, 0xe7, 0x94, 0xc6, 0x33,
	0x2e, 0x25, 0xcc, 0x27, 0x69, 0xe2, 0xed, 0xb7, 0x49, 0xbf, 0x16, 0xd5, 0xd6, 0x64, 0x98, 0x30,
	0x8f, 0x1e, 0xc4, 0x4a, 0x22, 0x48, 0xf4, 0x4a, 0x76, 0xb6, 0x91, 0xec, 0x94, 0x56, 0x16, 0x4f,
	0x0a, 0xc1, 0x2b, 0x5b, 0xee, 0x44, 0x41, 0xa5, 0x92, 0x31, 0x78, 0x15, 0x47, 0xad, 0x60, 0x3d,
	0xda, 0x40, 0xc8, 0xf4, 0x04, 0xf9, 0x52, 0x00, 0x16, 0x17, 0x55, 0xed, 0xb8, 0x5e, 0xd0, 0xb1,
	0x85, 0xc3, 0xa4, 0x13, 0xd3, 0xe3, 0x9d, 0xc0, 0x46, 0xb3, 0x33, 0x5a, 0xcd, 0x8c, 0x28, 0x0c,
	0xc4, 0x9d, 0x97, 0x19, 0x31, 0x4c, 0x58, 0x97, 0x1e, 0x15, 0x18, 0xd3, 0x0c, 0x0c, 0xf2, 0x4c,
	0xdb, 0xdc, 0xa5, 0xa8, 0x9e, 0x19, 0x31, 0xde, 0xb0, 0xef, 0x28, 0xa5, 0x1f, 0x51, 0x2e, 0x05,
	0x3d, 0xbc, 0x5d, 0x77, 0xc8, 0x1e, 0x5c, 0x43, 0x37, 0xee, 0xb5, 0x23, 0x23, 0x58, 0x37, 0xf8,
	0xa5, 0xe8, 0x60, 0xb7, 0xc6, 0x66, 0xef, 0xff, 0x9f, 0x8c, 0xbe, 0x0e, 0xde, 0x72, 0x9f, 0xbc,
	0xe7, 0x3e, 0xf9, 0xc8, 0x7d, 0xf2, 0xf2, 0xe9, 0xef, 0xdd, 0xb7, 0xfe, 0xda, 0xf8, 0xb4, 0x6a,
	0xd7, 0x76, 0xf5, 0x15, 0x00, 0x00, 0xff, 0xff, 0xfd, 0xd2, 0xe9, 0xe4, 0x18, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// KaiheilaClient is the client API for Kaiheila service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type KaiheilaClient interface {
	SendChannelMsg(ctx context.Context, in *SendMessageReq, opts ...grpc.CallOption) (*SendMessageResp, error)
}

type kaiheilaClient struct {
	cc *grpc.ClientConn
}

func NewKaiheilaClient(cc *grpc.ClientConn) KaiheilaClient {
	return &kaiheilaClient{cc}
}

func (c *kaiheilaClient) SendChannelMsg(ctx context.Context, in *SendMessageReq, opts ...grpc.CallOption) (*SendMessageResp, error) {
	out := new(SendMessageResp)
	err := c.cc.Invoke(ctx, "/ncs.bot.kaiheila.v1.Kaiheila/SendChannelMsg", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// KaiheilaServer is the server API for Kaiheila service.
type KaiheilaServer interface {
	SendChannelMsg(context.Context, *SendMessageReq) (*SendMessageResp, error)
}

// UnimplementedKaiheilaServer can be embedded to have forward compatible implementations.
type UnimplementedKaiheilaServer struct {
}

func (*UnimplementedKaiheilaServer) SendChannelMsg(ctx context.Context, req *SendMessageReq) (*SendMessageResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendChannelMsg not implemented")
}

func RegisterKaiheilaServer(s *grpc.Server, srv KaiheilaServer) {
	s.RegisterService(&_Kaiheila_serviceDesc, srv)
}

func _Kaiheila_SendChannelMsg_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendMessageReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KaiheilaServer).SendChannelMsg(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ncs.bot.kaiheila.v1.Kaiheila/SendChannelMsg",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KaiheilaServer).SendChannelMsg(ctx, req.(*SendMessageReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _Kaiheila_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ncs.bot.kaiheila.v1.Kaiheila",
	HandlerType: (*KaiheilaServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendChannelMsg",
			Handler:    _Kaiheila_SendChannelMsg_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "app/bot/kaiheila/api/grpc/v1/kaiheila.proto",
}

func (m *SendMessageReq) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SendMessageReq) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SendMessageReq) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.TempTargetId) > 0 {
		i -= len(m.TempTargetId)
		copy(dAtA[i:], m.TempTargetId)
		i = encodeVarintKaiheila(dAtA, i, uint64(len(m.TempTargetId)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.Nonce) > 0 {
		i -= len(m.Nonce)
		copy(dAtA[i:], m.Nonce)
		i = encodeVarintKaiheila(dAtA, i, uint64(len(m.Nonce)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.Quote) > 0 {
		i -= len(m.Quote)
		copy(dAtA[i:], m.Quote)
		i = encodeVarintKaiheila(dAtA, i, uint64(len(m.Quote)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Content) > 0 {
		i -= len(m.Content)
		copy(dAtA[i:], m.Content)
		i = encodeVarintKaiheila(dAtA, i, uint64(len(m.Content)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.ChannelId) > 0 {
		i -= len(m.ChannelId)
		copy(dAtA[i:], m.ChannelId)
		i = encodeVarintKaiheila(dAtA, i, uint64(len(m.ChannelId)))
		i--
		dAtA[i] = 0x12
	}
	if m.Type != 0 {
		i = encodeVarintKaiheila(dAtA, i, uint64(m.Type))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *SendMessageResp) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SendMessageResp) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SendMessageResp) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Nonce) > 0 {
		i -= len(m.Nonce)
		copy(dAtA[i:], m.Nonce)
		i = encodeVarintKaiheila(dAtA, i, uint64(len(m.Nonce)))
		i--
		dAtA[i] = 0x1a
	}
	if m.MsgTimestamp != 0 {
		i = encodeVarintKaiheila(dAtA, i, uint64(m.MsgTimestamp))
		i--
		dAtA[i] = 0x10
	}
	if len(m.MsgId) > 0 {
		i -= len(m.MsgId)
		copy(dAtA[i:], m.MsgId)
		i = encodeVarintKaiheila(dAtA, i, uint64(len(m.MsgId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintKaiheila(dAtA []byte, offset int, v uint64) int {
	offset -= sovKaiheila(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *SendMessageReq) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Type != 0 {
		n += 1 + sovKaiheila(uint64(m.Type))
	}
	l = len(m.ChannelId)
	if l > 0 {
		n += 1 + l + sovKaiheila(uint64(l))
	}
	l = len(m.Content)
	if l > 0 {
		n += 1 + l + sovKaiheila(uint64(l))
	}
	l = len(m.Quote)
	if l > 0 {
		n += 1 + l + sovKaiheila(uint64(l))
	}
	l = len(m.Nonce)
	if l > 0 {
		n += 1 + l + sovKaiheila(uint64(l))
	}
	l = len(m.TempTargetId)
	if l > 0 {
		n += 1 + l + sovKaiheila(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *SendMessageResp) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.MsgId)
	if l > 0 {
		n += 1 + l + sovKaiheila(uint64(l))
	}
	if m.MsgTimestamp != 0 {
		n += 1 + sovKaiheila(uint64(m.MsgTimestamp))
	}
	l = len(m.Nonce)
	if l > 0 {
		n += 1 + l + sovKaiheila(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovKaiheila(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozKaiheila(x uint64) (n int) {
	return sovKaiheila(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *SendMessageReq) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowKaiheila
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
			return fmt.Errorf("proto: SendMessageReq: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SendMessageReq: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			m.Type = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowKaiheila
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Type |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChannelId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowKaiheila
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
				return ErrInvalidLengthKaiheila
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthKaiheila
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ChannelId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Content", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowKaiheila
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
				return ErrInvalidLengthKaiheila
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthKaiheila
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Content = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Quote", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowKaiheila
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
				return ErrInvalidLengthKaiheila
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthKaiheila
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Quote = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Nonce", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowKaiheila
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
				return ErrInvalidLengthKaiheila
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthKaiheila
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Nonce = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TempTargetId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowKaiheila
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
				return ErrInvalidLengthKaiheila
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthKaiheila
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TempTargetId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipKaiheila(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthKaiheila
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthKaiheila
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
func (m *SendMessageResp) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowKaiheila
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
			return fmt.Errorf("proto: SendMessageResp: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SendMessageResp: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MsgId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowKaiheila
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
				return ErrInvalidLengthKaiheila
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthKaiheila
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MsgId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MsgTimestamp", wireType)
			}
			m.MsgTimestamp = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowKaiheila
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MsgTimestamp |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Nonce", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowKaiheila
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
				return ErrInvalidLengthKaiheila
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthKaiheila
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Nonce = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipKaiheila(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthKaiheila
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthKaiheila
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
func skipKaiheila(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowKaiheila
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
					return 0, ErrIntOverflowKaiheila
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
					return 0, ErrIntOverflowKaiheila
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
				return 0, ErrInvalidLengthKaiheila
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupKaiheila
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthKaiheila
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthKaiheila        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowKaiheila          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupKaiheila = fmt.Errorf("proto: unexpected end of group")
)
