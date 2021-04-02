// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: app/bot/qq/api/grpc/v1/qq.proto

package v1

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type SendGroupMessageReq struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	AutoEscape           bool     `protobuf:"varint,2,opt,name=auto_escape,json=autoEscape,proto3" json:"auto_escape,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SendGroupMessageReq) Reset()         { *m = SendGroupMessageReq{} }
func (m *SendGroupMessageReq) String() string { return proto.CompactTextString(m) }
func (*SendGroupMessageReq) ProtoMessage()    {}
func (*SendGroupMessageReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_653ff383dcae9fe4, []int{0}
}
func (m *SendGroupMessageReq) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SendGroupMessageReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SendGroupMessageReq.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SendGroupMessageReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SendGroupMessageReq.Merge(m, src)
}
func (m *SendGroupMessageReq) XXX_Size() int {
	return m.Size()
}
func (m *SendGroupMessageReq) XXX_DiscardUnknown() {
	xxx_messageInfo_SendGroupMessageReq.DiscardUnknown(m)
}

var xxx_messageInfo_SendGroupMessageReq proto.InternalMessageInfo

func (m *SendGroupMessageReq) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *SendGroupMessageReq) GetAutoEscape() bool {
	if m != nil {
		return m.AutoEscape
	}
	return false
}

type SendGroupMessageResp struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SendGroupMessageResp) Reset()         { *m = SendGroupMessageResp{} }
func (m *SendGroupMessageResp) String() string { return proto.CompactTextString(m) }
func (*SendGroupMessageResp) ProtoMessage()    {}
func (*SendGroupMessageResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_653ff383dcae9fe4, []int{1}
}
func (m *SendGroupMessageResp) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SendGroupMessageResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SendGroupMessageResp.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SendGroupMessageResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SendGroupMessageResp.Merge(m, src)
}
func (m *SendGroupMessageResp) XXX_Size() int {
	return m.Size()
}
func (m *SendGroupMessageResp) XXX_DiscardUnknown() {
	xxx_messageInfo_SendGroupMessageResp.DiscardUnknown(m)
}

var xxx_messageInfo_SendGroupMessageResp proto.InternalMessageInfo

func init() {
	proto.RegisterType((*SendGroupMessageReq)(nil), "ncs.bot.qq.v1.SendGroupMessageReq")
	proto.RegisterType((*SendGroupMessageResp)(nil), "ncs.bot.qq.v1.SendGroupMessageResp")
}

func init() { proto.RegisterFile("app/bot/qq/api/grpc/v1/qq.proto", fileDescriptor_653ff383dcae9fe4) }

var fileDescriptor_653ff383dcae9fe4 = []byte{
	// 203 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x4f, 0x2c, 0x28, 0xd0,
	0x4f, 0xca, 0x2f, 0xd1, 0x2f, 0x2c, 0xd4, 0x4f, 0x2c, 0xc8, 0xd4, 0x4f, 0x2f, 0x2a, 0x48, 0xd6,
	0x2f, 0x33, 0xd4, 0x2f, 0x2c, 0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0xcd, 0x4b, 0x2e,
	0xd6, 0x4b, 0xca, 0x2f, 0xd1, 0x2b, 0x2c, 0xd4, 0x2b, 0x33, 0x54, 0x0a, 0xe0, 0x12, 0x0e, 0x4e,
	0xcd, 0x4b, 0x71, 0x2f, 0xca, 0x2f, 0x2d, 0xf0, 0x4d, 0x2d, 0x2e, 0x4e, 0x4c, 0x4f, 0x0d, 0x4a,
	0x2d, 0x14, 0x92, 0xe0, 0x62, 0xcf, 0x85, 0xf0, 0x24, 0x18, 0x15, 0x18, 0x35, 0x38, 0x83, 0x60,
	0x5c, 0x21, 0x79, 0x2e, 0xee, 0xc4, 0xd2, 0x92, 0xfc, 0xf8, 0xd4, 0xe2, 0xe4, 0xc4, 0x82, 0x54,
	0x09, 0x26, 0x05, 0x46, 0x0d, 0x8e, 0x20, 0x2e, 0x90, 0x90, 0x2b, 0x58, 0x44, 0x49, 0x8c, 0x4b,
	0x04, 0xd3, 0xc4, 0xe2, 0x02, 0xa3, 0x44, 0x2e, 0xa6, 0xc0, 0x40, 0xa1, 0x68, 0x2e, 0x01, 0x74,
	0x59, 0x21, 0x25, 0x3d, 0x14, 0x37, 0xe9, 0x61, 0x71, 0x90, 0x94, 0x32, 0x41, 0x35, 0xc5, 0x05,
	0x4e, 0x1a, 0x27, 0x1e, 0xc9, 0x31, 0x5e, 0x78, 0x24, 0xc7, 0xf8, 0xe0, 0x91, 0x1c, 0xe3, 0x8c,
	0xc7, 0x72, 0x0c, 0x51, 0x62, 0xd8, 0x83, 0x23, 0x89, 0x0d, 0x1c, 0x18, 0xc6, 0x80, 0x00, 0x00,
	0x00, 0xff, 0xff, 0x57, 0xd2, 0x87, 0x19, 0x2f, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// QQClient is the client API for QQ service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QQClient interface {
	SendGroupMessage(ctx context.Context, in *SendGroupMessageReq, opts ...grpc.CallOption) (*SendGroupMessageResp, error)
}

type qQClient struct {
	cc *grpc.ClientConn
}

func NewQQClient(cc *grpc.ClientConn) QQClient {
	return &qQClient{cc}
}

func (c *qQClient) SendGroupMessage(ctx context.Context, in *SendGroupMessageReq, opts ...grpc.CallOption) (*SendGroupMessageResp, error) {
	out := new(SendGroupMessageResp)
	err := c.cc.Invoke(ctx, "/ncs.bot.qq.v1.QQ/SendGroupMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QQServer is the server API for QQ service.
type QQServer interface {
	SendGroupMessage(context.Context, *SendGroupMessageReq) (*SendGroupMessageResp, error)
}

// UnimplementedQQServer can be embedded to have forward compatible implementations.
type UnimplementedQQServer struct {
}

func (*UnimplementedQQServer) SendGroupMessage(ctx context.Context, req *SendGroupMessageReq) (*SendGroupMessageResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendGroupMessage not implemented")
}

func RegisterQQServer(s *grpc.Server, srv QQServer) {
	s.RegisterService(&_QQ_serviceDesc, srv)
}

func _QQ_SendGroupMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendGroupMessageReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QQServer).SendGroupMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ncs.bot.qq.v1.QQ/SendGroupMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QQServer).SendGroupMessage(ctx, req.(*SendGroupMessageReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _QQ_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ncs.bot.qq.v1.QQ",
	HandlerType: (*QQServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendGroupMessage",
			Handler:    _QQ_SendGroupMessage_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "app/bot/qq/api/grpc/v1/qq.proto",
}

func (m *SendGroupMessageReq) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SendGroupMessageReq) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SendGroupMessageReq) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.AutoEscape {
		i--
		if m.AutoEscape {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x10
	}
	if len(m.Message) > 0 {
		i -= len(m.Message)
		copy(dAtA[i:], m.Message)
		i = encodeVarintQq(dAtA, i, uint64(len(m.Message)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *SendGroupMessageResp) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SendGroupMessageResp) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SendGroupMessageResp) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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

func encodeVarintQq(dAtA []byte, offset int, v uint64) int {
	offset -= sovQq(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *SendGroupMessageReq) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Message)
	if l > 0 {
		n += 1 + l + sovQq(uint64(l))
	}
	if m.AutoEscape {
		n += 2
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *SendGroupMessageResp) Size() (n int) {
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

func sovQq(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQq(x uint64) (n int) {
	return sovQq(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *SendGroupMessageReq) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQq
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
			return fmt.Errorf("proto: SendGroupMessageReq: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SendGroupMessageReq: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Message", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQq
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
				return ErrInvalidLengthQq
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQq
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Message = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AutoEscape", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQq
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.AutoEscape = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipQq(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthQq
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthQq
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
func (m *SendGroupMessageResp) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQq
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
			return fmt.Errorf("proto: SendGroupMessageResp: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SendGroupMessageResp: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipQq(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthQq
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthQq
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
func skipQq(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowQq
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
					return 0, ErrIntOverflowQq
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
					return 0, ErrIntOverflowQq
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
				return 0, ErrInvalidLengthQq
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupQq
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthQq
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthQq        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowQq          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupQq = fmt.Errorf("proto: unexpected end of group")
)