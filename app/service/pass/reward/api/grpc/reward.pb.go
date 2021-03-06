// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: app/service/pass/reward/api/grpc/reward.proto

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

type Item struct {
	Level                int32    `protobuf:"varint,1,opt,name=level,proto3" json:"level,omitempty"`
	Id                   int32    `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	Amount               int32    `protobuf:"varint,3,opt,name=amount,proto3" json:"amount,omitempty"`
	Length               int64    `protobuf:"varint,4,opt,name=length,proto3" json:"length,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Item) Reset()         { *m = Item{} }
func (m *Item) String() string { return proto.CompactTextString(m) }
func (*Item) ProtoMessage()    {}
func (*Item) Descriptor() ([]byte, []int) {
	return fileDescriptor_e41ed0dd4bb3f6eb, []int{0}
}
func (m *Item) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Item) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Item.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Item) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Item.Merge(m, src)
}
func (m *Item) XXX_Size() int {
	return m.Size()
}
func (m *Item) XXX_DiscardUnknown() {
	xxx_messageInfo_Item.DiscardUnknown(m)
}

var xxx_messageInfo_Item proto.InternalMessageInfo

func (m *Item) GetLevel() int32 {
	if m != nil {
		return m.Level
	}
	return 0
}

func (m *Item) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Item) GetAmount() int32 {
	if m != nil {
		return m.Amount
	}
	return 0
}

func (m *Item) GetLength() int64 {
	if m != nil {
		return m.Length
	}
	return 0
}

type GetRewardsReq struct {
	Level                int32    `protobuf:"varint,1,opt,name=level,proto3" json:"level,omitempty"`
	Min                  int32    `protobuf:"varint,2,opt,name=min,proto3" json:"min,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetRewardsReq) Reset()         { *m = GetRewardsReq{} }
func (m *GetRewardsReq) String() string { return proto.CompactTextString(m) }
func (*GetRewardsReq) ProtoMessage()    {}
func (*GetRewardsReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_e41ed0dd4bb3f6eb, []int{1}
}
func (m *GetRewardsReq) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GetRewardsReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GetRewardsReq.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GetRewardsReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetRewardsReq.Merge(m, src)
}
func (m *GetRewardsReq) XXX_Size() int {
	return m.Size()
}
func (m *GetRewardsReq) XXX_DiscardUnknown() {
	xxx_messageInfo_GetRewardsReq.DiscardUnknown(m)
}

var xxx_messageInfo_GetRewardsReq proto.InternalMessageInfo

func (m *GetRewardsReq) GetLevel() int32 {
	if m != nil {
		return m.Level
	}
	return 0
}

func (m *GetRewardsReq) GetMin() int32 {
	if m != nil {
		return m.Min
	}
	return 0
}

type GetRewardsResp struct {
	Season               int32    `protobuf:"varint,1,opt,name=season,proto3" json:"season,omitempty"`
	FreeRewards          []*Item  `protobuf:"bytes,2,rep,name=free_rewards,json=freeRewards,proto3" json:"free_rewards,omitempty"`
	AdvRewards           []*Item  `protobuf:"bytes,3,rep,name=adv_rewards,json=advRewards,proto3" json:"adv_rewards,omitempty"`
	MaxLevel             int32    `protobuf:"varint,4,opt,name=max_level,json=maxLevel,proto3" json:"max_level,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetRewardsResp) Reset()         { *m = GetRewardsResp{} }
func (m *GetRewardsResp) String() string { return proto.CompactTextString(m) }
func (*GetRewardsResp) ProtoMessage()    {}
func (*GetRewardsResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_e41ed0dd4bb3f6eb, []int{2}
}
func (m *GetRewardsResp) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GetRewardsResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GetRewardsResp.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GetRewardsResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetRewardsResp.Merge(m, src)
}
func (m *GetRewardsResp) XXX_Size() int {
	return m.Size()
}
func (m *GetRewardsResp) XXX_DiscardUnknown() {
	xxx_messageInfo_GetRewardsResp.DiscardUnknown(m)
}

var xxx_messageInfo_GetRewardsResp proto.InternalMessageInfo

func (m *GetRewardsResp) GetSeason() int32 {
	if m != nil {
		return m.Season
	}
	return 0
}

func (m *GetRewardsResp) GetFreeRewards() []*Item {
	if m != nil {
		return m.FreeRewards
	}
	return nil
}

func (m *GetRewardsResp) GetAdvRewards() []*Item {
	if m != nil {
		return m.AdvRewards
	}
	return nil
}

func (m *GetRewardsResp) GetMaxLevel() int32 {
	if m != nil {
		return m.MaxLevel
	}
	return 0
}

func init() {
	proto.RegisterType((*Item)(nil), "ncs.pass.reward.Item")
	proto.RegisterType((*GetRewardsReq)(nil), "ncs.pass.reward.GetRewardsReq")
	proto.RegisterType((*GetRewardsResp)(nil), "ncs.pass.reward.GetRewardsResp")
}

func init() {
	proto.RegisterFile("app/service/pass/reward/api/grpc/reward.proto", fileDescriptor_e41ed0dd4bb3f6eb)
}

var fileDescriptor_e41ed0dd4bb3f6eb = []byte{
	// 377 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x52, 0x41, 0x4b, 0xc3, 0x30,
	0x18, 0x35, 0x6d, 0x37, 0x34, 0x9b, 0x53, 0x82, 0x8e, 0x32, 0x47, 0x2d, 0x3d, 0xed, 0x62, 0x0b,
	0x13, 0xd4, 0xb3, 0x17, 0x11, 0x3c, 0x48, 0x0f, 0x1e, 0x86, 0x30, 0xe2, 0x1a, 0x6b, 0xa1, 0x4d,
	0x62, 0x53, 0xeb, 0x60, 0xec, 0xe2, 0x5f, 0xf0, 0xe2, 0x1f, 0xf0, 0xee, 0xcd, 0xbf, 0xe0, 0x51,
	0xf0, 0x0f, 0xc8, 0xf4, 0x87, 0x48, 0xd2, 0xaa, 0x9b, 0xb2, 0x9d, 0xda, 0xf7, 0xf2, 0x5e, 0xbe,
	0xef, 0x85, 0x07, 0x77, 0x30, 0xe7, 0x9e, 0x20, 0x69, 0x1e, 0x0d, 0x88, 0xc7, 0xb1, 0x10, 0x5e,
	0x4a, 0x6e, 0x71, 0x1a, 0x78, 0x98, 0x47, 0x5e, 0x98, 0xf2, 0x41, 0x89, 0x5d, 0x9e, 0xb2, 0x8c,
	0xa1, 0x35, 0x3a, 0x10, 0xae, 0x94, 0xb9, 0x05, 0xdd, 0x6a, 0x87, 0x8c, 0x85, 0x31, 0x51, 0x72,
	0x4c, 0x29, 0xcb, 0x70, 0x16, 0x31, 0x2a, 0x0a, 0xb9, 0x73, 0x0e, 0x8d, 0xe3, 0x8c, 0x24, 0x68,
	0x03, 0x56, 0x62, 0x92, 0x93, 0xd8, 0x04, 0x36, 0xe8, 0x54, 0xfc, 0x02, 0xa0, 0x06, 0xd4, 0xa2,
	0xc0, 0xd4, 0x14, 0xa5, 0x45, 0x01, 0x6a, 0xc2, 0x2a, 0x4e, 0xd8, 0x0d, 0xcd, 0x4c, 0x5d, 0x71,
	0x25, 0x92, 0x7c, 0x4c, 0x68, 0x98, 0x5d, 0x99, 0x86, 0x0d, 0x3a, 0xba, 0x5f, 0x22, 0x67, 0x1f,
	0xae, 0x1e, 0x91, 0xcc, 0x57, 0x8b, 0x08, 0x9f, 0x5c, 0xcf, 0x19, 0xb3, 0x0e, 0xf5, 0x24, 0xa2,
	0xe5, 0x1c, 0xf9, 0xeb, 0x3c, 0x03, 0xd8, 0x98, 0x76, 0x0a, 0x2e, 0x67, 0x08, 0x82, 0x05, 0xa3,
	0xa5, 0xb7, 0x44, 0xe8, 0x00, 0xd6, 0x2f, 0x53, 0x42, 0xfa, 0x45, 0x5c, 0x61, 0x6a, 0xb6, 0xde,
	0xa9, 0x75, 0x37, 0xdd, 0x3f, 0xef, 0xe0, 0xca, 0x98, 0x7e, 0x4d, 0x4a, 0xcb, 0x5b, 0xd1, 0x1e,
	0xac, 0xe1, 0x20, 0xff, 0x31, 0xea, 0x8b, 0x8c, 0x10, 0x07, 0xf9, 0xb7, 0x6f, 0x0b, 0xae, 0x24,
	0x78, 0xd8, 0x2f, 0x82, 0x18, 0x6a, 0x99, 0xe5, 0x04, 0x0f, 0x4f, 0x24, 0xee, 0x3e, 0x01, 0x58,
	0x2d, 0x84, 0xe8, 0x11, 0x40, 0xf8, 0x1b, 0x02, 0x59, 0xff, 0x6e, 0x9e, 0x79, 0x9b, 0xd6, 0xf6,
	0xc2, 0x73, 0xc1, 0x9d, 0xb3, 0xbb, 0xb7, 0xcf, 0x7b, 0xed, 0x14, 0xd5, 0xa7, 0x6b, 0xd0, 0x6b,
	0xa3, 0xd6, 0x4c, 0x2d, 0xd4, 0x5a, 0xde, 0x48, 0x7d, 0xc6, 0x3d, 0x07, 0xd9, 0xf3, 0x4f, 0xbd,
	0x51, 0x12, 0xd1, 0xf1, 0x61, 0xf3, 0x65, 0x62, 0x81, 0xd7, 0x89, 0x05, 0xde, 0x27, 0x16, 0x78,
	0xf8, 0xb0, 0x96, 0x7a, 0x86, 0xec, 0xd5, 0x45, 0x55, 0x55, 0x64, 0xf7, 0x2b, 0x00, 0x00, 0xff,
	0xff, 0x12, 0xf2, 0x16, 0xef, 0x82, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// RewardClient is the client API for Reward service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RewardClient interface {
	GetRewards(ctx context.Context, in *GetRewardsReq, opts ...grpc.CallOption) (*GetRewardsResp, error)
}

type rewardClient struct {
	cc *grpc.ClientConn
}

func NewRewardClient(cc *grpc.ClientConn) RewardClient {
	return &rewardClient{cc}
}

func (c *rewardClient) GetRewards(ctx context.Context, in *GetRewardsReq, opts ...grpc.CallOption) (*GetRewardsResp, error) {
	out := new(GetRewardsResp)
	err := c.cc.Invoke(ctx, "/ncs.pass.reward.Reward/GetRewards", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RewardServer is the server API for Reward service.
type RewardServer interface {
	GetRewards(context.Context, *GetRewardsReq) (*GetRewardsResp, error)
}

// UnimplementedRewardServer can be embedded to have forward compatible implementations.
type UnimplementedRewardServer struct {
}

func (*UnimplementedRewardServer) GetRewards(ctx context.Context, req *GetRewardsReq) (*GetRewardsResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRewards not implemented")
}

func RegisterRewardServer(s *grpc.Server, srv RewardServer) {
	s.RegisterService(&_Reward_serviceDesc, srv)
}

func _Reward_GetRewards_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRewardsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RewardServer).GetRewards(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ncs.pass.reward.Reward/GetRewards",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RewardServer).GetRewards(ctx, req.(*GetRewardsReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _Reward_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ncs.pass.reward.Reward",
	HandlerType: (*RewardServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetRewards",
			Handler:    _Reward_GetRewards_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "app/service/pass/reward/api/grpc/reward.proto",
}

func (m *Item) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Item) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Item) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Length != 0 {
		i = encodeVarintReward(dAtA, i, uint64(m.Length))
		i--
		dAtA[i] = 0x20
	}
	if m.Amount != 0 {
		i = encodeVarintReward(dAtA, i, uint64(m.Amount))
		i--
		dAtA[i] = 0x18
	}
	if m.Id != 0 {
		i = encodeVarintReward(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x10
	}
	if m.Level != 0 {
		i = encodeVarintReward(dAtA, i, uint64(m.Level))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *GetRewardsReq) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetRewardsReq) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GetRewardsReq) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Min != 0 {
		i = encodeVarintReward(dAtA, i, uint64(m.Min))
		i--
		dAtA[i] = 0x10
	}
	if m.Level != 0 {
		i = encodeVarintReward(dAtA, i, uint64(m.Level))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *GetRewardsResp) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetRewardsResp) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GetRewardsResp) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.MaxLevel != 0 {
		i = encodeVarintReward(dAtA, i, uint64(m.MaxLevel))
		i--
		dAtA[i] = 0x20
	}
	if len(m.AdvRewards) > 0 {
		for iNdEx := len(m.AdvRewards) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.AdvRewards[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintReward(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.FreeRewards) > 0 {
		for iNdEx := len(m.FreeRewards) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.FreeRewards[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintReward(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if m.Season != 0 {
		i = encodeVarintReward(dAtA, i, uint64(m.Season))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintReward(dAtA []byte, offset int, v uint64) int {
	offset -= sovReward(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Item) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Level != 0 {
		n += 1 + sovReward(uint64(m.Level))
	}
	if m.Id != 0 {
		n += 1 + sovReward(uint64(m.Id))
	}
	if m.Amount != 0 {
		n += 1 + sovReward(uint64(m.Amount))
	}
	if m.Length != 0 {
		n += 1 + sovReward(uint64(m.Length))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *GetRewardsReq) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Level != 0 {
		n += 1 + sovReward(uint64(m.Level))
	}
	if m.Min != 0 {
		n += 1 + sovReward(uint64(m.Min))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *GetRewardsResp) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Season != 0 {
		n += 1 + sovReward(uint64(m.Season))
	}
	if len(m.FreeRewards) > 0 {
		for _, e := range m.FreeRewards {
			l = e.Size()
			n += 1 + l + sovReward(uint64(l))
		}
	}
	if len(m.AdvRewards) > 0 {
		for _, e := range m.AdvRewards {
			l = e.Size()
			n += 1 + l + sovReward(uint64(l))
		}
	}
	if m.MaxLevel != 0 {
		n += 1 + sovReward(uint64(m.MaxLevel))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovReward(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozReward(x uint64) (n int) {
	return sovReward(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Item) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowReward
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
			return fmt.Errorf("proto: Item: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Item: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Level", wireType)
			}
			m.Level = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowReward
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Level |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowReward
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			m.Amount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowReward
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Amount |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Length", wireType)
			}
			m.Length = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowReward
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Length |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipReward(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthReward
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthReward
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
func (m *GetRewardsReq) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowReward
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
			return fmt.Errorf("proto: GetRewardsReq: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetRewardsReq: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Level", wireType)
			}
			m.Level = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowReward
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Level |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Min", wireType)
			}
			m.Min = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowReward
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Min |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipReward(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthReward
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthReward
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
func (m *GetRewardsResp) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowReward
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
			return fmt.Errorf("proto: GetRewardsResp: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetRewardsResp: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Season", wireType)
			}
			m.Season = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowReward
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Season |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FreeRewards", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowReward
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
				return ErrInvalidLengthReward
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthReward
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.FreeRewards = append(m.FreeRewards, &Item{})
			if err := m.FreeRewards[len(m.FreeRewards)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AdvRewards", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowReward
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
				return ErrInvalidLengthReward
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthReward
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AdvRewards = append(m.AdvRewards, &Item{})
			if err := m.AdvRewards[len(m.AdvRewards)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxLevel", wireType)
			}
			m.MaxLevel = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowReward
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MaxLevel |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipReward(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthReward
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthReward
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
func skipReward(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowReward
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
					return 0, ErrIntOverflowReward
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
					return 0, ErrIntOverflowReward
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
				return 0, ErrInvalidLengthReward
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupReward
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthReward
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthReward        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowReward          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupReward = fmt.Errorf("proto: unexpected end of group")
)
