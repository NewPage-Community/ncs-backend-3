// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: app/game/store/api/grpc/store.proto

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

type HotSaleListReq struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HotSaleListReq) Reset()         { *m = HotSaleListReq{} }
func (m *HotSaleListReq) String() string { return proto.CompactTextString(m) }
func (*HotSaleListReq) ProtoMessage()    {}
func (*HotSaleListReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_e7984dd24fb15707, []int{0}
}
func (m *HotSaleListReq) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *HotSaleListReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_HotSaleListReq.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *HotSaleListReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HotSaleListReq.Merge(m, src)
}
func (m *HotSaleListReq) XXX_Size() int {
	return m.Size()
}
func (m *HotSaleListReq) XXX_DiscardUnknown() {
	xxx_messageInfo_HotSaleListReq.DiscardUnknown(m)
}

var xxx_messageInfo_HotSaleListReq proto.InternalMessageInfo

type HotSaleListResp struct {
	ItemsId              []int32  `protobuf:"varint,1,rep,packed,name=items_id,json=itemsId,proto3" json:"items_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HotSaleListResp) Reset()         { *m = HotSaleListResp{} }
func (m *HotSaleListResp) String() string { return proto.CompactTextString(m) }
func (*HotSaleListResp) ProtoMessage()    {}
func (*HotSaleListResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_e7984dd24fb15707, []int{1}
}
func (m *HotSaleListResp) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *HotSaleListResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_HotSaleListResp.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *HotSaleListResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HotSaleListResp.Merge(m, src)
}
func (m *HotSaleListResp) XXX_Size() int {
	return m.Size()
}
func (m *HotSaleListResp) XXX_DiscardUnknown() {
	xxx_messageInfo_HotSaleListResp.DiscardUnknown(m)
}

var xxx_messageInfo_HotSaleListResp proto.InternalMessageInfo

func (m *HotSaleListResp) GetItemsId() []int32 {
	if m != nil {
		return m.ItemsId
	}
	return nil
}

type BuyItemReq struct {
	Uid                  int64    `protobuf:"varint,1,opt,name=uid,proto3" json:"uid,omitempty"`
	ItemId               int32    `protobuf:"varint,2,opt,name=item_id,json=itemId,proto3" json:"item_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BuyItemReq) Reset()         { *m = BuyItemReq{} }
func (m *BuyItemReq) String() string { return proto.CompactTextString(m) }
func (*BuyItemReq) ProtoMessage()    {}
func (*BuyItemReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_e7984dd24fb15707, []int{2}
}
func (m *BuyItemReq) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *BuyItemReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_BuyItemReq.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *BuyItemReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BuyItemReq.Merge(m, src)
}
func (m *BuyItemReq) XXX_Size() int {
	return m.Size()
}
func (m *BuyItemReq) XXX_DiscardUnknown() {
	xxx_messageInfo_BuyItemReq.DiscardUnknown(m)
}

var xxx_messageInfo_BuyItemReq proto.InternalMessageInfo

func (m *BuyItemReq) GetUid() int64 {
	if m != nil {
		return m.Uid
	}
	return 0
}

func (m *BuyItemReq) GetItemId() int32 {
	if m != nil {
		return m.ItemId
	}
	return 0
}

type BuyItemResp struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BuyItemResp) Reset()         { *m = BuyItemResp{} }
func (m *BuyItemResp) String() string { return proto.CompactTextString(m) }
func (*BuyItemResp) ProtoMessage()    {}
func (*BuyItemResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_e7984dd24fb15707, []int{3}
}
func (m *BuyItemResp) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *BuyItemResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_BuyItemResp.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *BuyItemResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BuyItemResp.Merge(m, src)
}
func (m *BuyItemResp) XXX_Size() int {
	return m.Size()
}
func (m *BuyItemResp) XXX_DiscardUnknown() {
	xxx_messageInfo_BuyItemResp.DiscardUnknown(m)
}

var xxx_messageInfo_BuyItemResp proto.InternalMessageInfo

func init() {
	proto.RegisterType((*HotSaleListReq)(nil), "ncs.game.store.HotSaleListReq")
	proto.RegisterType((*HotSaleListResp)(nil), "ncs.game.store.HotSaleListResp")
	proto.RegisterType((*BuyItemReq)(nil), "ncs.game.store.BuyItemReq")
	proto.RegisterType((*BuyItemResp)(nil), "ncs.game.store.BuyItemResp")
}

func init() {
	proto.RegisterFile("app/game/store/api/grpc/store.proto", fileDescriptor_e7984dd24fb15707)
}

var fileDescriptor_e7984dd24fb15707 = []byte{
	// 304 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x4e, 0x2c, 0x28, 0xd0,
	0x4f, 0x4f, 0xcc, 0x4d, 0xd5, 0x2f, 0x2e, 0xc9, 0x2f, 0x4a, 0xd5, 0x4f, 0x2c, 0xc8, 0xd4, 0x4f,
	0x2f, 0x2a, 0x48, 0x86, 0x70, 0xf5, 0x0a, 0x8a, 0xf2, 0x4b, 0xf2, 0x85, 0xf8, 0xf2, 0x92, 0x8b,
	0xf5, 0x40, 0x8a, 0xf4, 0xc0, 0xa2, 0x52, 0x32, 0xe9, 0xf9, 0xf9, 0xe9, 0x39, 0x10, 0xc5, 0x89,
	0x79, 0x79, 0xf9, 0x25, 0x89, 0x25, 0x99, 0xf9, 0x79, 0xc5, 0x10, 0xd5, 0x4a, 0x02, 0x5c, 0x7c,
	0x1e, 0xf9, 0x25, 0xc1, 0x89, 0x39, 0xa9, 0x3e, 0x99, 0xc5, 0x25, 0x41, 0xa9, 0x85, 0x4a, 0x3a,
	0x5c, 0xfc, 0x28, 0x22, 0xc5, 0x05, 0x42, 0x92, 0x5c, 0x1c, 0x99, 0x25, 0xa9, 0xb9, 0xc5, 0xf1,
	0x99, 0x29, 0x12, 0x8c, 0x0a, 0xcc, 0x1a, 0xac, 0x41, 0xec, 0x60, 0xbe, 0x67, 0x8a, 0x92, 0x39,
	0x17, 0x97, 0x53, 0x69, 0xa5, 0x67, 0x49, 0x6a, 0x6e, 0x50, 0x6a, 0xa1, 0x90, 0x00, 0x17, 0x73,
	0x29, 0x58, 0x0d, 0xa3, 0x06, 0x73, 0x10, 0x88, 0x29, 0x24, 0xce, 0x05, 0x56, 0x0a, 0xd2, 0xc9,
	0xa4, 0xc0, 0xa8, 0xc1, 0x1a, 0xc4, 0x06, 0xe2, 0x7a, 0xa6, 0x28, 0xf1, 0x72, 0x71, 0xc3, 0x35,
	0x16, 0x17, 0x18, 0x5d, 0x61, 0xe4, 0x62, 0x0d, 0x06, 0xb9, 0x57, 0x28, 0x8e, 0x8b, 0x1d, 0x2a,
	0x21, 0x24, 0xa5, 0x87, 0xea, 0x17, 0x3d, 0x84, 0x55, 0x52, 0xd2, 0x38, 0xe5, 0x8a, 0x0b, 0x94,
	0xa4, 0x9a, 0x2e, 0x3f, 0x99, 0xcc, 0x24, 0xa2, 0xc4, 0x8f, 0x1c, 0x5a, 0x49, 0xa5, 0x95, 0x56,
	0x8c, 0x5a, 0x42, 0xd9, 0x5c, 0xdc, 0x48, 0xfe, 0x13, 0x92, 0x43, 0x37, 0x07, 0x35, 0x38, 0xa4,
	0xe4, 0xf1, 0xca, 0x17, 0x17, 0x28, 0x49, 0x83, 0xed, 0x12, 0x15, 0x12, 0x46, 0xb6, 0x2b, 0x23,
	0xbf, 0xa4, 0x38, 0x31, 0x27, 0xd5, 0x49, 0xec, 0xc4, 0x23, 0x39, 0xc6, 0x0b, 0x8f, 0xe4, 0x18,
	0x1f, 0x3c, 0x92, 0x63, 0x9c, 0xf1, 0x58, 0x8e, 0x21, 0x8a, 0x05, 0x14, 0x61, 0x49, 0x6c, 0xe0,
	0xd0, 0x37, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0xf6, 0x08, 0x45, 0x64, 0xd2, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// StoreClient is the client API for Store service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type StoreClient interface {
	BuyItem(ctx context.Context, in *BuyItemReq, opts ...grpc.CallOption) (*BuyItemResp, error)
	HotSaleList(ctx context.Context, in *HotSaleListReq, opts ...grpc.CallOption) (*HotSaleListResp, error)
}

type storeClient struct {
	cc *grpc.ClientConn
}

func NewStoreClient(cc *grpc.ClientConn) StoreClient {
	return &storeClient{cc}
}

func (c *storeClient) BuyItem(ctx context.Context, in *BuyItemReq, opts ...grpc.CallOption) (*BuyItemResp, error) {
	out := new(BuyItemResp)
	err := c.cc.Invoke(ctx, "/ncs.game.store.Store/BuyItem", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storeClient) HotSaleList(ctx context.Context, in *HotSaleListReq, opts ...grpc.CallOption) (*HotSaleListResp, error) {
	out := new(HotSaleListResp)
	err := c.cc.Invoke(ctx, "/ncs.game.store.Store/HotSaleList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StoreServer is the server API for Store service.
type StoreServer interface {
	BuyItem(context.Context, *BuyItemReq) (*BuyItemResp, error)
	HotSaleList(context.Context, *HotSaleListReq) (*HotSaleListResp, error)
}

// UnimplementedStoreServer can be embedded to have forward compatible implementations.
type UnimplementedStoreServer struct {
}

func (*UnimplementedStoreServer) BuyItem(ctx context.Context, req *BuyItemReq) (*BuyItemResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BuyItem not implemented")
}
func (*UnimplementedStoreServer) HotSaleList(ctx context.Context, req *HotSaleListReq) (*HotSaleListResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HotSaleList not implemented")
}

func RegisterStoreServer(s *grpc.Server, srv StoreServer) {
	s.RegisterService(&_Store_serviceDesc, srv)
}

func _Store_BuyItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BuyItemReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StoreServer).BuyItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ncs.game.store.Store/BuyItem",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StoreServer).BuyItem(ctx, req.(*BuyItemReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Store_HotSaleList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HotSaleListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StoreServer).HotSaleList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ncs.game.store.Store/HotSaleList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StoreServer).HotSaleList(ctx, req.(*HotSaleListReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _Store_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ncs.game.store.Store",
	HandlerType: (*StoreServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "BuyItem",
			Handler:    _Store_BuyItem_Handler,
		},
		{
			MethodName: "HotSaleList",
			Handler:    _Store_HotSaleList_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "app/game/store/api/grpc/store.proto",
}

func (m *HotSaleListReq) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *HotSaleListReq) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *HotSaleListReq) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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

func (m *HotSaleListResp) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *HotSaleListResp) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *HotSaleListResp) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.ItemsId) > 0 {
		dAtA2 := make([]byte, len(m.ItemsId)*10)
		var j1 int
		for _, num1 := range m.ItemsId {
			num := uint64(num1)
			for num >= 1<<7 {
				dAtA2[j1] = uint8(uint64(num)&0x7f | 0x80)
				num >>= 7
				j1++
			}
			dAtA2[j1] = uint8(num)
			j1++
		}
		i -= j1
		copy(dAtA[i:], dAtA2[:j1])
		i = encodeVarintStore(dAtA, i, uint64(j1))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *BuyItemReq) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BuyItemReq) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *BuyItemReq) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.ItemId != 0 {
		i = encodeVarintStore(dAtA, i, uint64(m.ItemId))
		i--
		dAtA[i] = 0x10
	}
	if m.Uid != 0 {
		i = encodeVarintStore(dAtA, i, uint64(m.Uid))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *BuyItemResp) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BuyItemResp) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *BuyItemResp) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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

func encodeVarintStore(dAtA []byte, offset int, v uint64) int {
	offset -= sovStore(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *HotSaleListReq) Size() (n int) {
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

func (m *HotSaleListResp) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.ItemsId) > 0 {
		l = 0
		for _, e := range m.ItemsId {
			l += sovStore(uint64(e))
		}
		n += 1 + sovStore(uint64(l)) + l
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *BuyItemReq) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Uid != 0 {
		n += 1 + sovStore(uint64(m.Uid))
	}
	if m.ItemId != 0 {
		n += 1 + sovStore(uint64(m.ItemId))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *BuyItemResp) Size() (n int) {
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

func sovStore(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozStore(x uint64) (n int) {
	return sovStore(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *HotSaleListReq) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowStore
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
			return fmt.Errorf("proto: HotSaleListReq: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: HotSaleListReq: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipStore(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthStore
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthStore
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
func (m *HotSaleListResp) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowStore
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
			return fmt.Errorf("proto: HotSaleListResp: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: HotSaleListResp: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType == 0 {
				var v int32
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowStore
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= int32(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				m.ItemsId = append(m.ItemsId, v)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowStore
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					packedLen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if packedLen < 0 {
					return ErrInvalidLengthStore
				}
				postIndex := iNdEx + packedLen
				if postIndex < 0 {
					return ErrInvalidLengthStore
				}
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				var elementCount int
				var count int
				for _, integer := range dAtA[iNdEx:postIndex] {
					if integer < 128 {
						count++
					}
				}
				elementCount = count
				if elementCount != 0 && len(m.ItemsId) == 0 {
					m.ItemsId = make([]int32, 0, elementCount)
				}
				for iNdEx < postIndex {
					var v int32
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowStore
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= int32(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					m.ItemsId = append(m.ItemsId, v)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field ItemsId", wireType)
			}
		default:
			iNdEx = preIndex
			skippy, err := skipStore(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthStore
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthStore
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
func (m *BuyItemReq) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowStore
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
			return fmt.Errorf("proto: BuyItemReq: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BuyItemReq: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Uid", wireType)
			}
			m.Uid = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStore
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
				return fmt.Errorf("proto: wrong wireType = %d for field ItemId", wireType)
			}
			m.ItemId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStore
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ItemId |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipStore(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthStore
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthStore
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
func (m *BuyItemResp) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowStore
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
			return fmt.Errorf("proto: BuyItemResp: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BuyItemResp: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipStore(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthStore
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthStore
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
func skipStore(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowStore
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
					return 0, ErrIntOverflowStore
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
					return 0, ErrIntOverflowStore
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
				return 0, ErrInvalidLengthStore
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupStore
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthStore
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthStore        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowStore          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupStore = fmt.Errorf("proto: unexpected end of group")
)