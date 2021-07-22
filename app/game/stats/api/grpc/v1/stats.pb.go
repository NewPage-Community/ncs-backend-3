// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.6.1
// source: app/game/stats/api/grpc/v1/stats.proto

package v1

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type GetReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Range     string `protobuf:"bytes,1,opt,name=range,proto3" json:"range,omitempty"`
	StatsName string `protobuf:"bytes,2,opt,name=stats_name,json=statsName,proto3" json:"stats_name,omitempty"`
	Version   string `protobuf:"bytes,3,opt,name=version,proto3" json:"version,omitempty"`
	Uid       int64  `protobuf:"varint,4,opt,name=uid,proto3" json:"uid,omitempty"`
}

func (x *GetReq) Reset() {
	*x = GetReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_game_stats_api_grpc_v1_stats_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetReq) ProtoMessage() {}

func (x *GetReq) ProtoReflect() protoreflect.Message {
	mi := &file_app_game_stats_api_grpc_v1_stats_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetReq.ProtoReflect.Descriptor instead.
func (*GetReq) Descriptor() ([]byte, []int) {
	return file_app_game_stats_api_grpc_v1_stats_proto_rawDescGZIP(), []int{0}
}

func (x *GetReq) GetRange() string {
	if x != nil {
		return x.Range
	}
	return ""
}

func (x *GetReq) GetStatsName() string {
	if x != nil {
		return x.StatsName
	}
	return ""
}

func (x *GetReq) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *GetReq) GetUid() int64 {
	if x != nil {
		return x.Uid
	}
	return 0
}

type GetResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid   int64   `protobuf:"varint,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Score float32 `protobuf:"fixed32,2,opt,name=score,proto3" json:"score,omitempty"`
	Rank  int64   `protobuf:"varint,3,opt,name=rank,proto3" json:"rank,omitempty"`
}

func (x *GetResp) Reset() {
	*x = GetResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_game_stats_api_grpc_v1_stats_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetResp) ProtoMessage() {}

func (x *GetResp) ProtoReflect() protoreflect.Message {
	mi := &file_app_game_stats_api_grpc_v1_stats_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetResp.ProtoReflect.Descriptor instead.
func (*GetResp) Descriptor() ([]byte, []int) {
	return file_app_game_stats_api_grpc_v1_stats_proto_rawDescGZIP(), []int{1}
}

func (x *GetResp) GetUid() int64 {
	if x != nil {
		return x.Uid
	}
	return 0
}

func (x *GetResp) GetScore() float32 {
	if x != nil {
		return x.Score
	}
	return 0
}

func (x *GetResp) GetRank() int64 {
	if x != nil {
		return x.Rank
	}
	return 0
}

type GetAllReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Range     string `protobuf:"bytes,1,opt,name=range,proto3" json:"range,omitempty"`
	StatsName string `protobuf:"bytes,2,opt,name=stats_name,json=statsName,proto3" json:"stats_name,omitempty"`
	Version   string `protobuf:"bytes,3,opt,name=version,proto3" json:"version,omitempty"`
}

func (x *GetAllReq) Reset() {
	*x = GetAllReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_game_stats_api_grpc_v1_stats_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllReq) ProtoMessage() {}

func (x *GetAllReq) ProtoReflect() protoreflect.Message {
	mi := &file_app_game_stats_api_grpc_v1_stats_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllReq.ProtoReflect.Descriptor instead.
func (*GetAllReq) Descriptor() ([]byte, []int) {
	return file_app_game_stats_api_grpc_v1_stats_proto_rawDescGZIP(), []int{2}
}

func (x *GetAllReq) GetRange() string {
	if x != nil {
		return x.Range
	}
	return ""
}

func (x *GetAllReq) GetStatsName() string {
	if x != nil {
		return x.StatsName
	}
	return ""
}

func (x *GetAllReq) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

type GetAllResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Stats []*GetResp `protobuf:"bytes,1,rep,name=stats,proto3" json:"stats,omitempty"`
}

func (x *GetAllResp) Reset() {
	*x = GetAllResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_game_stats_api_grpc_v1_stats_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllResp) ProtoMessage() {}

func (x *GetAllResp) ProtoReflect() protoreflect.Message {
	mi := &file_app_game_stats_api_grpc_v1_stats_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllResp.ProtoReflect.Descriptor instead.
func (*GetAllResp) Descriptor() ([]byte, []int) {
	return file_app_game_stats_api_grpc_v1_stats_proto_rawDescGZIP(), []int{3}
}

func (x *GetAllResp) GetStats() []*GetResp {
	if x != nil {
		return x.Stats
	}
	return nil
}

type GetsReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Stats []*GetReq `protobuf:"bytes,1,rep,name=stats,proto3" json:"stats,omitempty"`
}

func (x *GetsReq) Reset() {
	*x = GetsReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_game_stats_api_grpc_v1_stats_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetsReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetsReq) ProtoMessage() {}

func (x *GetsReq) ProtoReflect() protoreflect.Message {
	mi := &file_app_game_stats_api_grpc_v1_stats_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetsReq.ProtoReflect.Descriptor instead.
func (*GetsReq) Descriptor() ([]byte, []int) {
	return file_app_game_stats_api_grpc_v1_stats_proto_rawDescGZIP(), []int{4}
}

func (x *GetsReq) GetStats() []*GetReq {
	if x != nil {
		return x.Stats
	}
	return nil
}

type StatsInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Range     string  `protobuf:"bytes,1,opt,name=range,proto3" json:"range,omitempty"`
	StatsName string  `protobuf:"bytes,2,opt,name=stats_name,json=statsName,proto3" json:"stats_name,omitempty"`
	Version   string  `protobuf:"bytes,3,opt,name=version,proto3" json:"version,omitempty"`
	Uid       int64   `protobuf:"varint,4,opt,name=uid,proto3" json:"uid,omitempty"`
	Score     float32 `protobuf:"fixed32,5,opt,name=score,proto3" json:"score,omitempty"`
	Rank      int64   `protobuf:"varint,6,opt,name=rank,proto3" json:"rank,omitempty"`
}

func (x *StatsInfo) Reset() {
	*x = StatsInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_game_stats_api_grpc_v1_stats_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StatsInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StatsInfo) ProtoMessage() {}

func (x *StatsInfo) ProtoReflect() protoreflect.Message {
	mi := &file_app_game_stats_api_grpc_v1_stats_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StatsInfo.ProtoReflect.Descriptor instead.
func (*StatsInfo) Descriptor() ([]byte, []int) {
	return file_app_game_stats_api_grpc_v1_stats_proto_rawDescGZIP(), []int{5}
}

func (x *StatsInfo) GetRange() string {
	if x != nil {
		return x.Range
	}
	return ""
}

func (x *StatsInfo) GetStatsName() string {
	if x != nil {
		return x.StatsName
	}
	return ""
}

func (x *StatsInfo) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *StatsInfo) GetUid() int64 {
	if x != nil {
		return x.Uid
	}
	return 0
}

func (x *StatsInfo) GetScore() float32 {
	if x != nil {
		return x.Score
	}
	return 0
}

func (x *StatsInfo) GetRank() int64 {
	if x != nil {
		return x.Rank
	}
	return 0
}

type GetsResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Stats []*StatsInfo `protobuf:"bytes,1,rep,name=stats,proto3" json:"stats,omitempty"`
}

func (x *GetsResp) Reset() {
	*x = GetsResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_game_stats_api_grpc_v1_stats_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetsResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetsResp) ProtoMessage() {}

func (x *GetsResp) ProtoReflect() protoreflect.Message {
	mi := &file_app_game_stats_api_grpc_v1_stats_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetsResp.ProtoReflect.Descriptor instead.
func (*GetsResp) Descriptor() ([]byte, []int) {
	return file_app_game_stats_api_grpc_v1_stats_proto_rawDescGZIP(), []int{6}
}

func (x *GetsResp) GetStats() []*StatsInfo {
	if x != nil {
		return x.Stats
	}
	return nil
}

type SetReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Range     string  `protobuf:"bytes,1,opt,name=range,proto3" json:"range,omitempty"`
	StatsName string  `protobuf:"bytes,2,opt,name=stats_name,json=statsName,proto3" json:"stats_name,omitempty"`
	Version   string  `protobuf:"bytes,3,opt,name=version,proto3" json:"version,omitempty"`
	Uid       int64   `protobuf:"varint,4,opt,name=uid,proto3" json:"uid,omitempty"`
	Score     float32 `protobuf:"fixed32,5,opt,name=score,proto3" json:"score,omitempty"`
}

func (x *SetReq) Reset() {
	*x = SetReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_game_stats_api_grpc_v1_stats_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetReq) ProtoMessage() {}

func (x *SetReq) ProtoReflect() protoreflect.Message {
	mi := &file_app_game_stats_api_grpc_v1_stats_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetReq.ProtoReflect.Descriptor instead.
func (*SetReq) Descriptor() ([]byte, []int) {
	return file_app_game_stats_api_grpc_v1_stats_proto_rawDescGZIP(), []int{7}
}

func (x *SetReq) GetRange() string {
	if x != nil {
		return x.Range
	}
	return ""
}

func (x *SetReq) GetStatsName() string {
	if x != nil {
		return x.StatsName
	}
	return ""
}

func (x *SetReq) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *SetReq) GetUid() int64 {
	if x != nil {
		return x.Uid
	}
	return 0
}

func (x *SetReq) GetScore() float32 {
	if x != nil {
		return x.Score
	}
	return 0
}

type SetsReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Stats []*SetReq `protobuf:"bytes,1,rep,name=stats,proto3" json:"stats,omitempty"`
}

func (x *SetsReq) Reset() {
	*x = SetsReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_game_stats_api_grpc_v1_stats_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetsReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetsReq) ProtoMessage() {}

func (x *SetsReq) ProtoReflect() protoreflect.Message {
	mi := &file_app_game_stats_api_grpc_v1_stats_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetsReq.ProtoReflect.Descriptor instead.
func (*SetsReq) Descriptor() ([]byte, []int) {
	return file_app_game_stats_api_grpc_v1_stats_proto_rawDescGZIP(), []int{8}
}

func (x *SetsReq) GetStats() []*SetReq {
	if x != nil {
		return x.Stats
	}
	return nil
}

type SetResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SetResp) Reset() {
	*x = SetResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_game_stats_api_grpc_v1_stats_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetResp) ProtoMessage() {}

func (x *SetResp) ProtoReflect() protoreflect.Message {
	mi := &file_app_game_stats_api_grpc_v1_stats_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetResp.ProtoReflect.Descriptor instead.
func (*SetResp) Descriptor() ([]byte, []int) {
	return file_app_game_stats_api_grpc_v1_stats_proto_rawDescGZIP(), []int{9}
}

type IncrReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Range     string  `protobuf:"bytes,1,opt,name=range,proto3" json:"range,omitempty"`
	StatsName string  `protobuf:"bytes,2,opt,name=stats_name,json=statsName,proto3" json:"stats_name,omitempty"`
	Version   string  `protobuf:"bytes,3,opt,name=version,proto3" json:"version,omitempty"`
	Uid       int64   `protobuf:"varint,4,opt,name=uid,proto3" json:"uid,omitempty"`
	Increment float32 `protobuf:"fixed32,5,opt,name=increment,proto3" json:"increment,omitempty"`
}

func (x *IncrReq) Reset() {
	*x = IncrReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_game_stats_api_grpc_v1_stats_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IncrReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IncrReq) ProtoMessage() {}

func (x *IncrReq) ProtoReflect() protoreflect.Message {
	mi := &file_app_game_stats_api_grpc_v1_stats_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IncrReq.ProtoReflect.Descriptor instead.
func (*IncrReq) Descriptor() ([]byte, []int) {
	return file_app_game_stats_api_grpc_v1_stats_proto_rawDescGZIP(), []int{10}
}

func (x *IncrReq) GetRange() string {
	if x != nil {
		return x.Range
	}
	return ""
}

func (x *IncrReq) GetStatsName() string {
	if x != nil {
		return x.StatsName
	}
	return ""
}

func (x *IncrReq) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *IncrReq) GetUid() int64 {
	if x != nil {
		return x.Uid
	}
	return 0
}

func (x *IncrReq) GetIncrement() float32 {
	if x != nil {
		return x.Increment
	}
	return 0
}

type IncrsReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Stats []*IncrReq `protobuf:"bytes,1,rep,name=stats,proto3" json:"stats,omitempty"`
}

func (x *IncrsReq) Reset() {
	*x = IncrsReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_game_stats_api_grpc_v1_stats_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IncrsReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IncrsReq) ProtoMessage() {}

func (x *IncrsReq) ProtoReflect() protoreflect.Message {
	mi := &file_app_game_stats_api_grpc_v1_stats_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IncrsReq.ProtoReflect.Descriptor instead.
func (*IncrsReq) Descriptor() ([]byte, []int) {
	return file_app_game_stats_api_grpc_v1_stats_proto_rawDescGZIP(), []int{11}
}

func (x *IncrsReq) GetStats() []*IncrReq {
	if x != nil {
		return x.Stats
	}
	return nil
}

type IncrResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *IncrResp) Reset() {
	*x = IncrResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_game_stats_api_grpc_v1_stats_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IncrResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IncrResp) ProtoMessage() {}

func (x *IncrResp) ProtoReflect() protoreflect.Message {
	mi := &file_app_game_stats_api_grpc_v1_stats_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IncrResp.ProtoReflect.Descriptor instead.
func (*IncrResp) Descriptor() ([]byte, []int) {
	return file_app_game_stats_api_grpc_v1_stats_proto_rawDescGZIP(), []int{12}
}

var File_app_game_stats_api_grpc_v1_stats_proto protoreflect.FileDescriptor

var file_app_game_stats_api_grpc_v1_stats_proto_rawDesc = []byte{
	0x0a, 0x26, 0x61, 0x70, 0x70, 0x2f, 0x67, 0x61, 0x6d, 0x65, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x73,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x74, 0x61,
	0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x11, 0x6e, 0x63, 0x73, 0x2e, 0x67, 0x61,
	0x6d, 0x65, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x69, 0x0a, 0x06, 0x47, 0x65, 0x74,
	0x52, 0x65, 0x71, 0x12, 0x14, 0x0a, 0x05, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x61,
	0x74, 0x73, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73,
	0x74, 0x61, 0x74, 0x73, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x03, 0x75, 0x69, 0x64, 0x22, 0x45, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x12,
	0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x75, 0x69,
	0x64, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02,
	0x52, 0x05, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x61, 0x6e, 0x6b, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x72, 0x61, 0x6e, 0x6b, 0x22, 0x5a, 0x0a, 0x09, 0x47,
	0x65, 0x74, 0x41, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x12, 0x14, 0x0a, 0x05, 0x72, 0x61, 0x6e, 0x67,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x1d,
	0x0a, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x73, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x73, 0x74, 0x61, 0x74, 0x73, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0x3e, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x41, 0x6c,
	0x6c, 0x52, 0x65, 0x73, 0x70, 0x12, 0x30, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x6e, 0x63, 0x73, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x2e,
	0x73, 0x74, 0x61, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x73, 0x22, 0x3a, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x73, 0x52,
	0x65, 0x71, 0x12, 0x2f, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x19, 0x2e, 0x6e, 0x63, 0x73, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x2e, 0x73, 0x74, 0x61,
	0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x52, 0x05, 0x73, 0x74,
	0x61, 0x74, 0x73, 0x22, 0x96, 0x01, 0x0a, 0x09, 0x53, 0x74, 0x61, 0x74, 0x73, 0x49, 0x6e, 0x66,
	0x6f, 0x12, 0x14, 0x0a, 0x05, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x73,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x74, 0x61,
	0x74, 0x73, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x75,
	0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x02, 0x52, 0x05, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x61, 0x6e, 0x6b,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x72, 0x61, 0x6e, 0x6b, 0x22, 0x3e, 0x0a, 0x08,
	0x47, 0x65, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x12, 0x32, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x6e, 0x63, 0x73, 0x2e, 0x67, 0x61,
	0x6d, 0x65, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x61, 0x74,
	0x73, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x73, 0x22, 0x7f, 0x0a, 0x06,
	0x53, 0x65, 0x74, 0x52, 0x65, 0x71, 0x12, 0x14, 0x0a, 0x05, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x1d, 0x0a, 0x0a,
	0x73, 0x74, 0x61, 0x74, 0x73, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x73, 0x74, 0x61, 0x74, 0x73, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x76,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x03, 0x75, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x63, 0x6f, 0x72, 0x65,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x22, 0x3a, 0x0a,
	0x07, 0x53, 0x65, 0x74, 0x73, 0x52, 0x65, 0x71, 0x12, 0x2f, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x6e, 0x63, 0x73, 0x2e, 0x67, 0x61,
	0x6d, 0x65, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x74, 0x52,
	0x65, 0x71, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x73, 0x22, 0x09, 0x0a, 0x07, 0x53, 0x65, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x22, 0x88, 0x01, 0x0a, 0x07, 0x49, 0x6e, 0x63, 0x72, 0x52, 0x65, 0x71,
	0x12, 0x14, 0x0a, 0x05, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x73, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x74, 0x61, 0x74,
	0x73, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12,
	0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x75, 0x69,
	0x64, 0x12, 0x1c, 0x0a, 0x09, 0x69, 0x6e, 0x63, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x02, 0x52, 0x09, 0x69, 0x6e, 0x63, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x22,
	0x3c, 0x0a, 0x08, 0x49, 0x6e, 0x63, 0x72, 0x73, 0x52, 0x65, 0x71, 0x12, 0x30, 0x0a, 0x05, 0x73,
	0x74, 0x61, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x6e, 0x63, 0x73,
	0x2e, 0x67, 0x61, 0x6d, 0x65, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x49,
	0x6e, 0x63, 0x72, 0x52, 0x65, 0x71, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x73, 0x22, 0x0a, 0x0a,
	0x08, 0x49, 0x6e, 0x63, 0x72, 0x52, 0x65, 0x73, 0x70, 0x32, 0xc0, 0x06, 0x0a, 0x05, 0x53, 0x74,
	0x61, 0x74, 0x73, 0x12, 0xac, 0x01, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x19, 0x2e, 0x6e, 0x63,
	0x73, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e,
	0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x1a, 0x2e, 0x6e, 0x63, 0x73, 0x2e, 0x67, 0x61, 0x6d,
	0x65, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x22, 0x6e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x68, 0x12, 0x2d, 0x2f, 0x67, 0x61, 0x6d,
	0x65, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x73, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x7b, 0x75, 0x69,
	0x64, 0x7d, 0x2f, 0x7b, 0x73, 0x74, 0x61, 0x74, 0x73, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x7d, 0x2f,
	0x7b, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x7d, 0x5a, 0x37, 0x12, 0x35, 0x2f, 0x67, 0x61,
	0x6d, 0x65, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x73, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x7b, 0x75,
	0x69, 0x64, 0x7d, 0x2f, 0x7b, 0x73, 0x74, 0x61, 0x74, 0x73, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x7d,
	0x2f, 0x7b, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x7d, 0x2f, 0x7b, 0x72, 0x61, 0x6e, 0x67,
	0x65, 0x7d, 0x12, 0x9f, 0x01, 0x0a, 0x06, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x12, 0x1c, 0x2e,
	0x6e, 0x63, 0x73, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x73, 0x2e, 0x76,
	0x31, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x1a, 0x1d, 0x2e, 0x6e, 0x63,
	0x73, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e,
	0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x22, 0x58, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x52, 0x12, 0x22, 0x2f, 0x67, 0x61, 0x6d, 0x65, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x73, 0x2f,
	0x7b, 0x73, 0x74, 0x61, 0x74, 0x73, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x7d, 0x2f, 0x7b, 0x76, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x7d, 0x5a, 0x2c, 0x12, 0x2a, 0x2f, 0x67, 0x61, 0x6d, 0x65, 0x2f,
	0x73, 0x74, 0x61, 0x74, 0x73, 0x2f, 0x7b, 0x73, 0x74, 0x61, 0x74, 0x73, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x7d, 0x2f, 0x7b, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x7d, 0x2f, 0x7b, 0x72, 0x61,
	0x6e, 0x67, 0x65, 0x7d, 0x12, 0x5c, 0x0a, 0x04, 0x47, 0x65, 0x74, 0x73, 0x12, 0x1a, 0x2e, 0x6e,
	0x63, 0x73, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x73, 0x2e, 0x76, 0x31,
	0x2e, 0x47, 0x65, 0x74, 0x73, 0x52, 0x65, 0x71, 0x1a, 0x1b, 0x2e, 0x6e, 0x63, 0x73, 0x2e, 0x67,
	0x61, 0x6d, 0x65, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x22, 0x1b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x15, 0x22, 0x10, 0x2f,
	0x67, 0x61, 0x6d, 0x65, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x73, 0x2f, 0x67, 0x65, 0x74, 0x73, 0x3a,
	0x01, 0x2a, 0x12, 0x5d, 0x0a, 0x03, 0x53, 0x65, 0x74, 0x12, 0x19, 0x2e, 0x6e, 0x63, 0x73, 0x2e,
	0x67, 0x61, 0x6d, 0x65, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65,
	0x74, 0x52, 0x65, 0x71, 0x1a, 0x1a, 0x2e, 0x6e, 0x63, 0x73, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x2e,
	0x73, 0x74, 0x61, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x22, 0x1f, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x19, 0x22, 0x14, 0x2f, 0x67, 0x61, 0x6d, 0x65, 0x2f,
	0x73, 0x74, 0x61, 0x74, 0x73, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x73, 0x65, 0x74, 0x3a, 0x01,
	0x2a, 0x12, 0x60, 0x0a, 0x04, 0x53, 0x65, 0x74, 0x73, 0x12, 0x1a, 0x2e, 0x6e, 0x63, 0x73, 0x2e,
	0x67, 0x61, 0x6d, 0x65, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65,
	0x74, 0x73, 0x52, 0x65, 0x71, 0x1a, 0x1a, 0x2e, 0x6e, 0x63, 0x73, 0x2e, 0x67, 0x61, 0x6d, 0x65,
	0x2e, 0x73, 0x74, 0x61, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x22, 0x20, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1a, 0x22, 0x15, 0x2f, 0x67, 0x61, 0x6d, 0x65,
	0x2f, 0x73, 0x74, 0x61, 0x74, 0x73, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x73, 0x65, 0x74, 0x73,
	0x3a, 0x01, 0x2a, 0x12, 0x61, 0x0a, 0x04, 0x49, 0x6e, 0x63, 0x72, 0x12, 0x1a, 0x2e, 0x6e, 0x63,
	0x73, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e,
	0x49, 0x6e, 0x63, 0x72, 0x52, 0x65, 0x71, 0x1a, 0x1b, 0x2e, 0x6e, 0x63, 0x73, 0x2e, 0x67, 0x61,
	0x6d, 0x65, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e, 0x63, 0x72,
	0x52, 0x65, 0x73, 0x70, 0x22, 0x20, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1a, 0x22, 0x15, 0x2f, 0x67,
	0x61, 0x6d, 0x65, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x73, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x69,
	0x6e, 0x63, 0x72, 0x3a, 0x01, 0x2a, 0x12, 0x64, 0x0a, 0x05, 0x49, 0x6e, 0x63, 0x72, 0x73, 0x12,
	0x1b, 0x2e, 0x6e, 0x63, 0x73, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x73,
	0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e, 0x63, 0x72, 0x73, 0x52, 0x65, 0x71, 0x1a, 0x1b, 0x2e, 0x6e,
	0x63, 0x73, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x73, 0x2e, 0x76, 0x31,
	0x2e, 0x49, 0x6e, 0x63, 0x72, 0x52, 0x65, 0x73, 0x70, 0x22, 0x21, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x1b, 0x22, 0x16, 0x2f, 0x67, 0x61, 0x6d, 0x65, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x73, 0x2f, 0x75,
	0x73, 0x65, 0x72, 0x2f, 0x69, 0x6e, 0x63, 0x72, 0x73, 0x3a, 0x01, 0x2a, 0x42, 0x24, 0x5a, 0x22,
	0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2f, 0x61, 0x70, 0x70, 0x2f, 0x67, 0x61, 0x6d, 0x65,
	0x2f, 0x73, 0x74, 0x61, 0x74, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f,
	0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_app_game_stats_api_grpc_v1_stats_proto_rawDescOnce sync.Once
	file_app_game_stats_api_grpc_v1_stats_proto_rawDescData = file_app_game_stats_api_grpc_v1_stats_proto_rawDesc
)

func file_app_game_stats_api_grpc_v1_stats_proto_rawDescGZIP() []byte {
	file_app_game_stats_api_grpc_v1_stats_proto_rawDescOnce.Do(func() {
		file_app_game_stats_api_grpc_v1_stats_proto_rawDescData = protoimpl.X.CompressGZIP(file_app_game_stats_api_grpc_v1_stats_proto_rawDescData)
	})
	return file_app_game_stats_api_grpc_v1_stats_proto_rawDescData
}

var file_app_game_stats_api_grpc_v1_stats_proto_msgTypes = make([]protoimpl.MessageInfo, 13)
var file_app_game_stats_api_grpc_v1_stats_proto_goTypes = []interface{}{
	(*GetReq)(nil),     // 0: ncs.game.stats.v1.GetReq
	(*GetResp)(nil),    // 1: ncs.game.stats.v1.GetResp
	(*GetAllReq)(nil),  // 2: ncs.game.stats.v1.GetAllReq
	(*GetAllResp)(nil), // 3: ncs.game.stats.v1.GetAllResp
	(*GetsReq)(nil),    // 4: ncs.game.stats.v1.GetsReq
	(*StatsInfo)(nil),  // 5: ncs.game.stats.v1.StatsInfo
	(*GetsResp)(nil),   // 6: ncs.game.stats.v1.GetsResp
	(*SetReq)(nil),     // 7: ncs.game.stats.v1.SetReq
	(*SetsReq)(nil),    // 8: ncs.game.stats.v1.SetsReq
	(*SetResp)(nil),    // 9: ncs.game.stats.v1.SetResp
	(*IncrReq)(nil),    // 10: ncs.game.stats.v1.IncrReq
	(*IncrsReq)(nil),   // 11: ncs.game.stats.v1.IncrsReq
	(*IncrResp)(nil),   // 12: ncs.game.stats.v1.IncrResp
}
var file_app_game_stats_api_grpc_v1_stats_proto_depIdxs = []int32{
	1,  // 0: ncs.game.stats.v1.GetAllResp.stats:type_name -> ncs.game.stats.v1.GetResp
	0,  // 1: ncs.game.stats.v1.GetsReq.stats:type_name -> ncs.game.stats.v1.GetReq
	5,  // 2: ncs.game.stats.v1.GetsResp.stats:type_name -> ncs.game.stats.v1.StatsInfo
	7,  // 3: ncs.game.stats.v1.SetsReq.stats:type_name -> ncs.game.stats.v1.SetReq
	10, // 4: ncs.game.stats.v1.IncrsReq.stats:type_name -> ncs.game.stats.v1.IncrReq
	0,  // 5: ncs.game.stats.v1.Stats.Get:input_type -> ncs.game.stats.v1.GetReq
	2,  // 6: ncs.game.stats.v1.Stats.GetAll:input_type -> ncs.game.stats.v1.GetAllReq
	4,  // 7: ncs.game.stats.v1.Stats.Gets:input_type -> ncs.game.stats.v1.GetsReq
	7,  // 8: ncs.game.stats.v1.Stats.Set:input_type -> ncs.game.stats.v1.SetReq
	8,  // 9: ncs.game.stats.v1.Stats.Sets:input_type -> ncs.game.stats.v1.SetsReq
	10, // 10: ncs.game.stats.v1.Stats.Incr:input_type -> ncs.game.stats.v1.IncrReq
	11, // 11: ncs.game.stats.v1.Stats.Incrs:input_type -> ncs.game.stats.v1.IncrsReq
	1,  // 12: ncs.game.stats.v1.Stats.Get:output_type -> ncs.game.stats.v1.GetResp
	3,  // 13: ncs.game.stats.v1.Stats.GetAll:output_type -> ncs.game.stats.v1.GetAllResp
	6,  // 14: ncs.game.stats.v1.Stats.Gets:output_type -> ncs.game.stats.v1.GetsResp
	9,  // 15: ncs.game.stats.v1.Stats.Set:output_type -> ncs.game.stats.v1.SetResp
	9,  // 16: ncs.game.stats.v1.Stats.Sets:output_type -> ncs.game.stats.v1.SetResp
	12, // 17: ncs.game.stats.v1.Stats.Incr:output_type -> ncs.game.stats.v1.IncrResp
	12, // 18: ncs.game.stats.v1.Stats.Incrs:output_type -> ncs.game.stats.v1.IncrResp
	12, // [12:19] is the sub-list for method output_type
	5,  // [5:12] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_app_game_stats_api_grpc_v1_stats_proto_init() }
func file_app_game_stats_api_grpc_v1_stats_proto_init() {
	if File_app_game_stats_api_grpc_v1_stats_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_app_game_stats_api_grpc_v1_stats_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_app_game_stats_api_grpc_v1_stats_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetResp); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_app_game_stats_api_grpc_v1_stats_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_app_game_stats_api_grpc_v1_stats_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllResp); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_app_game_stats_api_grpc_v1_stats_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetsReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_app_game_stats_api_grpc_v1_stats_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StatsInfo); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_app_game_stats_api_grpc_v1_stats_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetsResp); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_app_game_stats_api_grpc_v1_stats_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_app_game_stats_api_grpc_v1_stats_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetsReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_app_game_stats_api_grpc_v1_stats_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetResp); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_app_game_stats_api_grpc_v1_stats_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IncrReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_app_game_stats_api_grpc_v1_stats_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IncrsReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_app_game_stats_api_grpc_v1_stats_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IncrResp); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_app_game_stats_api_grpc_v1_stats_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   13,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_app_game_stats_api_grpc_v1_stats_proto_goTypes,
		DependencyIndexes: file_app_game_stats_api_grpc_v1_stats_proto_depIdxs,
		MessageInfos:      file_app_game_stats_api_grpc_v1_stats_proto_msgTypes,
	}.Build()
	File_app_game_stats_api_grpc_v1_stats_proto = out.File
	file_app_game_stats_api_grpc_v1_stats_proto_rawDesc = nil
	file_app_game_stats_api_grpc_v1_stats_proto_goTypes = nil
	file_app_game_stats_api_grpc_v1_stats_proto_depIdxs = nil
}