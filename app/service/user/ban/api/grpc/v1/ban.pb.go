// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.6.1
// source: app/service/user/ban/api/grpc/v1/ban.proto

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

type Info struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Uid        uint64 `protobuf:"varint,2,opt,name=uid,proto3" json:"uid,omitempty"`
	Ip         string `protobuf:"bytes,3,opt,name=ip,proto3" json:"ip,omitempty"`
	CreateTime int64  `protobuf:"varint,4,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	ExpireTime int64  `protobuf:"varint,5,opt,name=expire_time,json=expireTime,proto3" json:"expire_time,omitempty"`
	Type       int32  `protobuf:"varint,6,opt,name=type,proto3" json:"type,omitempty"`
	ServerId   int64  `protobuf:"varint,7,opt,name=server_id,json=serverId,proto3" json:"server_id,omitempty"`
	ModId      int64  `protobuf:"varint,8,opt,name=mod_id,json=modId,proto3" json:"mod_id,omitempty"`
	GameId     int64  `protobuf:"varint,9,opt,name=game_id,json=gameId,proto3" json:"game_id,omitempty"`
	Reason     string `protobuf:"bytes,10,opt,name=reason,proto3" json:"reason,omitempty"`
	BlockIp    bool   `protobuf:"varint,11,opt,name=block_ip,json=blockIp,proto3" json:"block_ip,omitempty"`
	AppId      uint32 `protobuf:"varint,12,opt,name=app_id,json=appId,proto3" json:"app_id,omitempty"`
}

func (x *Info) Reset() {
	*x = Info{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_service_user_ban_api_grpc_v1_ban_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Info) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Info) ProtoMessage() {}

func (x *Info) ProtoReflect() protoreflect.Message {
	mi := &file_app_service_user_ban_api_grpc_v1_ban_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Info.ProtoReflect.Descriptor instead.
func (*Info) Descriptor() ([]byte, []int) {
	return file_app_service_user_ban_api_grpc_v1_ban_proto_rawDescGZIP(), []int{0}
}

func (x *Info) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Info) GetUid() uint64 {
	if x != nil {
		return x.Uid
	}
	return 0
}

func (x *Info) GetIp() string {
	if x != nil {
		return x.Ip
	}
	return ""
}

func (x *Info) GetCreateTime() int64 {
	if x != nil {
		return x.CreateTime
	}
	return 0
}

func (x *Info) GetExpireTime() int64 {
	if x != nil {
		return x.ExpireTime
	}
	return 0
}

func (x *Info) GetType() int32 {
	if x != nil {
		return x.Type
	}
	return 0
}

func (x *Info) GetServerId() int64 {
	if x != nil {
		return x.ServerId
	}
	return 0
}

func (x *Info) GetModId() int64 {
	if x != nil {
		return x.ModId
	}
	return 0
}

func (x *Info) GetGameId() int64 {
	if x != nil {
		return x.GameId
	}
	return 0
}

func (x *Info) GetReason() string {
	if x != nil {
		return x.Reason
	}
	return ""
}

func (x *Info) GetBlockIp() bool {
	if x != nil {
		return x.BlockIp
	}
	return false
}

func (x *Info) GetAppId() uint32 {
	if x != nil {
		return x.AppId
	}
	return 0
}

type InfoReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid uint64 `protobuf:"varint,1,opt,name=uid,proto3" json:"uid,omitempty"`
}

func (x *InfoReq) Reset() {
	*x = InfoReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_service_user_ban_api_grpc_v1_ban_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InfoReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InfoReq) ProtoMessage() {}

func (x *InfoReq) ProtoReflect() protoreflect.Message {
	mi := &file_app_service_user_ban_api_grpc_v1_ban_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InfoReq.ProtoReflect.Descriptor instead.
func (*InfoReq) Descriptor() ([]byte, []int) {
	return file_app_service_user_ban_api_grpc_v1_ban_proto_rawDescGZIP(), []int{1}
}

func (x *InfoReq) GetUid() uint64 {
	if x != nil {
		return x.Uid
	}
	return 0
}

type Info2Req struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid      uint64 `protobuf:"varint,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Ip       string `protobuf:"bytes,2,opt,name=ip,proto3" json:"ip,omitempty"`
	ServerId int64  `protobuf:"varint,3,opt,name=server_id,json=serverId,proto3" json:"server_id,omitempty"`
	ModId    int64  `protobuf:"varint,4,opt,name=mod_id,json=modId,proto3" json:"mod_id,omitempty"`
	GameId   int64  `protobuf:"varint,5,opt,name=game_id,json=gameId,proto3" json:"game_id,omitempty"`
	AppId    uint32 `protobuf:"varint,6,opt,name=app_id,json=appId,proto3" json:"app_id,omitempty"`
}

func (x *Info2Req) Reset() {
	*x = Info2Req{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_service_user_ban_api_grpc_v1_ban_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Info2Req) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Info2Req) ProtoMessage() {}

func (x *Info2Req) ProtoReflect() protoreflect.Message {
	mi := &file_app_service_user_ban_api_grpc_v1_ban_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Info2Req.ProtoReflect.Descriptor instead.
func (*Info2Req) Descriptor() ([]byte, []int) {
	return file_app_service_user_ban_api_grpc_v1_ban_proto_rawDescGZIP(), []int{2}
}

func (x *Info2Req) GetUid() uint64 {
	if x != nil {
		return x.Uid
	}
	return 0
}

func (x *Info2Req) GetIp() string {
	if x != nil {
		return x.Ip
	}
	return ""
}

func (x *Info2Req) GetServerId() int64 {
	if x != nil {
		return x.ServerId
	}
	return 0
}

func (x *Info2Req) GetModId() int64 {
	if x != nil {
		return x.ModId
	}
	return 0
}

func (x *Info2Req) GetGameId() int64 {
	if x != nil {
		return x.GameId
	}
	return 0
}

func (x *Info2Req) GetAppId() uint32 {
	if x != nil {
		return x.AppId
	}
	return 0
}

type InfoResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info *Info `protobuf:"bytes,1,opt,name=info,proto3" json:"info,omitempty"`
}

func (x *InfoResp) Reset() {
	*x = InfoResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_service_user_ban_api_grpc_v1_ban_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InfoResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InfoResp) ProtoMessage() {}

func (x *InfoResp) ProtoReflect() protoreflect.Message {
	mi := &file_app_service_user_ban_api_grpc_v1_ban_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InfoResp.ProtoReflect.Descriptor instead.
func (*InfoResp) Descriptor() ([]byte, []int) {
	return file_app_service_user_ban_api_grpc_v1_ban_proto_rawDescGZIP(), []int{3}
}

func (x *InfoResp) GetInfo() *Info {
	if x != nil {
		return x.Info
	}
	return nil
}

type AddReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info *Info `protobuf:"bytes,1,opt,name=info,proto3" json:"info,omitempty"`
}

func (x *AddReq) Reset() {
	*x = AddReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_service_user_ban_api_grpc_v1_ban_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddReq) ProtoMessage() {}

func (x *AddReq) ProtoReflect() protoreflect.Message {
	mi := &file_app_service_user_ban_api_grpc_v1_ban_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddReq.ProtoReflect.Descriptor instead.
func (*AddReq) Descriptor() ([]byte, []int) {
	return file_app_service_user_ban_api_grpc_v1_ban_proto_rawDescGZIP(), []int{4}
}

func (x *AddReq) GetInfo() *Info {
	if x != nil {
		return x.Info
	}
	return nil
}

type AddResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *AddResp) Reset() {
	*x = AddResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_service_user_ban_api_grpc_v1_ban_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddResp) ProtoMessage() {}

func (x *AddResp) ProtoReflect() protoreflect.Message {
	mi := &file_app_service_user_ban_api_grpc_v1_ban_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddResp.ProtoReflect.Descriptor instead.
func (*AddResp) Descriptor() ([]byte, []int) {
	return file_app_service_user_ban_api_grpc_v1_ban_proto_rawDescGZIP(), []int{5}
}

type RemoveReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *RemoveReq) Reset() {
	*x = RemoveReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_service_user_ban_api_grpc_v1_ban_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoveReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveReq) ProtoMessage() {}

func (x *RemoveReq) ProtoReflect() protoreflect.Message {
	mi := &file_app_service_user_ban_api_grpc_v1_ban_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveReq.ProtoReflect.Descriptor instead.
func (*RemoveReq) Descriptor() ([]byte, []int) {
	return file_app_service_user_ban_api_grpc_v1_ban_proto_rawDescGZIP(), []int{6}
}

func (x *RemoveReq) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type RemoveResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *RemoveResp) Reset() {
	*x = RemoveResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_service_user_ban_api_grpc_v1_ban_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoveResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveResp) ProtoMessage() {}

func (x *RemoveResp) ProtoReflect() protoreflect.Message {
	mi := &file_app_service_user_ban_api_grpc_v1_ban_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveResp.ProtoReflect.Descriptor instead.
func (*RemoveResp) Descriptor() ([]byte, []int) {
	return file_app_service_user_ban_api_grpc_v1_ban_proto_rawDescGZIP(), []int{7}
}

var File_app_service_user_ban_api_grpc_v1_ban_proto protoreflect.FileDescriptor

var file_app_service_user_ban_api_grpc_v1_ban_proto_rawDesc = []byte{
	0x0a, 0x2a, 0x61, 0x70, 0x70, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x75, 0x73,
	0x65, 0x72, 0x2f, 0x62, 0x61, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f,
	0x76, 0x31, 0x2f, 0x62, 0x61, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x17, 0x6e, 0x63,
	0x73, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x62,
	0x61, 0x6e, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xa5, 0x02, 0x0a, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03,
	0x75, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x03, 0x75, 0x69, 0x64, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x70, 0x12, 0x1f,
	0x0a, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12,
	0x1f, 0x0a, 0x0b, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x54, 0x69, 0x6d, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x49,
	0x64, 0x12, 0x15, 0x0a, 0x06, 0x6d, 0x6f, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x05, 0x6d, 0x6f, 0x64, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x67, 0x61, 0x6d, 0x65,
	0x5f, 0x69, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x67, 0x61, 0x6d, 0x65, 0x49,
	0x64, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x18, 0x0a, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x12, 0x19, 0x0a, 0x08, 0x62, 0x6c, 0x6f,
	0x63, 0x6b, 0x5f, 0x69, 0x70, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x62, 0x6c, 0x6f,
	0x63, 0x6b, 0x49, 0x70, 0x12, 0x15, 0x0a, 0x06, 0x61, 0x70, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x0c,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x61, 0x70, 0x70, 0x49, 0x64, 0x22, 0x1b, 0x0a, 0x07, 0x49,
	0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x03, 0x75, 0x69, 0x64, 0x22, 0x90, 0x01, 0x0a, 0x08, 0x49, 0x6e, 0x66,
	0x6f, 0x32, 0x52, 0x65, 0x71, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x03, 0x75, 0x69, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x70, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x70, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x49, 0x64, 0x12, 0x15, 0x0a, 0x06, 0x6d, 0x6f, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x6d, 0x6f, 0x64, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x67,
	0x61, 0x6d, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x67, 0x61,
	0x6d, 0x65, 0x49, 0x64, 0x12, 0x15, 0x0a, 0x06, 0x61, 0x70, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x61, 0x70, 0x70, 0x49, 0x64, 0x22, 0x3d, 0x0a, 0x08, 0x49,
	0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x12, 0x31, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x6e, 0x63, 0x73, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x62, 0x61, 0x6e, 0x2e, 0x76, 0x31, 0x2e,
	0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x22, 0x3b, 0x0a, 0x06, 0x41, 0x64,
	0x64, 0x52, 0x65, 0x71, 0x12, 0x31, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x6e, 0x63, 0x73, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x62, 0x61, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e, 0x66,
	0x6f, 0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x22, 0x09, 0x0a, 0x07, 0x41, 0x64, 0x64, 0x52, 0x65,
	0x73, 0x70, 0x22, 0x1b, 0x0a, 0x09, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x52, 0x65, 0x71, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x22,
	0x0c, 0x0a, 0x0a, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x52, 0x65, 0x73, 0x70, 0x32, 0xb5, 0x03,
	0x0a, 0x03, 0x42, 0x61, 0x6e, 0x12, 0x64, 0x0a, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x20, 0x2e,
	0x6e, 0x63, 0x73, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x75, 0x73, 0x65, 0x72,
	0x2e, 0x62, 0x61, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x1a,
	0x21, 0x2e, 0x6e, 0x63, 0x73, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x75, 0x73,
	0x65, 0x72, 0x2e, 0x62, 0x61, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65,
	0x73, 0x70, 0x22, 0x17, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x11, 0x12, 0x0f, 0x2f, 0x75, 0x73, 0x65,
	0x72, 0x2f, 0x62, 0x61, 0x6e, 0x2f, 0x7b, 0x75, 0x69, 0x64, 0x7d, 0x12, 0x62, 0x0a, 0x03, 0x41,
	0x64, 0x64, 0x12, 0x1f, 0x2e, 0x6e, 0x63, 0x73, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x62, 0x61, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x64, 0x64,
	0x52, 0x65, 0x71, 0x1a, 0x20, 0x2e, 0x6e, 0x63, 0x73, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x62, 0x61, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x64,
	0x64, 0x52, 0x65, 0x73, 0x70, 0x22, 0x18, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x12, 0x22, 0x0d, 0x2f,
	0x75, 0x73, 0x65, 0x72, 0x2f, 0x62, 0x61, 0x6e, 0x2f, 0x61, 0x64, 0x64, 0x3a, 0x01, 0x2a, 0x12,
	0x70, 0x0a, 0x06, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x12, 0x22, 0x2e, 0x6e, 0x63, 0x73, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x62, 0x61, 0x6e,
	0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x23, 0x2e,
	0x6e, 0x63, 0x73, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x75, 0x73, 0x65, 0x72,
	0x2e, 0x62, 0x61, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x22, 0x1d, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x17, 0x2a, 0x15, 0x2f, 0x75, 0x73, 0x65,
	0x72, 0x2f, 0x62, 0x61, 0x6e, 0x2f, 0x72, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x2f, 0x7b, 0x69, 0x64,
	0x7d, 0x12, 0x72, 0x0a, 0x08, 0x42, 0x61, 0x6e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x12, 0x21, 0x2e,
	0x6e, 0x63, 0x73, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x75, 0x73, 0x65, 0x72,
	0x2e, 0x62, 0x61, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e, 0x66, 0x6f, 0x32, 0x52, 0x65, 0x71,
	0x1a, 0x21, 0x2e, 0x6e, 0x63, 0x73, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x75,
	0x73, 0x65, 0x72, 0x2e, 0x62, 0x61, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e, 0x66, 0x6f, 0x52,
	0x65, 0x73, 0x70, 0x22, 0x20, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1a, 0x22, 0x15, 0x2f, 0x75, 0x73,
	0x65, 0x72, 0x2f, 0x62, 0x61, 0x6e, 0x2f, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x2f, 0x7b, 0x75, 0x69,
	0x64, 0x7d, 0x3a, 0x01, 0x2a, 0x42, 0x2a, 0x5a, 0x28, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64,
	0x2f, 0x61, 0x70, 0x70, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x75, 0x73, 0x65,
	0x72, 0x2f, 0x62, 0x61, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x76,
	0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_app_service_user_ban_api_grpc_v1_ban_proto_rawDescOnce sync.Once
	file_app_service_user_ban_api_grpc_v1_ban_proto_rawDescData = file_app_service_user_ban_api_grpc_v1_ban_proto_rawDesc
)

func file_app_service_user_ban_api_grpc_v1_ban_proto_rawDescGZIP() []byte {
	file_app_service_user_ban_api_grpc_v1_ban_proto_rawDescOnce.Do(func() {
		file_app_service_user_ban_api_grpc_v1_ban_proto_rawDescData = protoimpl.X.CompressGZIP(file_app_service_user_ban_api_grpc_v1_ban_proto_rawDescData)
	})
	return file_app_service_user_ban_api_grpc_v1_ban_proto_rawDescData
}

var file_app_service_user_ban_api_grpc_v1_ban_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_app_service_user_ban_api_grpc_v1_ban_proto_goTypes = []interface{}{
	(*Info)(nil),       // 0: ncs.service.user.ban.v1.Info
	(*InfoReq)(nil),    // 1: ncs.service.user.ban.v1.InfoReq
	(*Info2Req)(nil),   // 2: ncs.service.user.ban.v1.Info2Req
	(*InfoResp)(nil),   // 3: ncs.service.user.ban.v1.InfoResp
	(*AddReq)(nil),     // 4: ncs.service.user.ban.v1.AddReq
	(*AddResp)(nil),    // 5: ncs.service.user.ban.v1.AddResp
	(*RemoveReq)(nil),  // 6: ncs.service.user.ban.v1.RemoveReq
	(*RemoveResp)(nil), // 7: ncs.service.user.ban.v1.RemoveResp
}
var file_app_service_user_ban_api_grpc_v1_ban_proto_depIdxs = []int32{
	0, // 0: ncs.service.user.ban.v1.InfoResp.info:type_name -> ncs.service.user.ban.v1.Info
	0, // 1: ncs.service.user.ban.v1.AddReq.info:type_name -> ncs.service.user.ban.v1.Info
	1, // 2: ncs.service.user.ban.v1.Ban.Info:input_type -> ncs.service.user.ban.v1.InfoReq
	4, // 3: ncs.service.user.ban.v1.Ban.Add:input_type -> ncs.service.user.ban.v1.AddReq
	6, // 4: ncs.service.user.ban.v1.Ban.Remove:input_type -> ncs.service.user.ban.v1.RemoveReq
	2, // 5: ncs.service.user.ban.v1.Ban.BanCheck:input_type -> ncs.service.user.ban.v1.Info2Req
	3, // 6: ncs.service.user.ban.v1.Ban.Info:output_type -> ncs.service.user.ban.v1.InfoResp
	5, // 7: ncs.service.user.ban.v1.Ban.Add:output_type -> ncs.service.user.ban.v1.AddResp
	7, // 8: ncs.service.user.ban.v1.Ban.Remove:output_type -> ncs.service.user.ban.v1.RemoveResp
	3, // 9: ncs.service.user.ban.v1.Ban.BanCheck:output_type -> ncs.service.user.ban.v1.InfoResp
	6, // [6:10] is the sub-list for method output_type
	2, // [2:6] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_app_service_user_ban_api_grpc_v1_ban_proto_init() }
func file_app_service_user_ban_api_grpc_v1_ban_proto_init() {
	if File_app_service_user_ban_api_grpc_v1_ban_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_app_service_user_ban_api_grpc_v1_ban_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Info); i {
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
		file_app_service_user_ban_api_grpc_v1_ban_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InfoReq); i {
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
		file_app_service_user_ban_api_grpc_v1_ban_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Info2Req); i {
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
		file_app_service_user_ban_api_grpc_v1_ban_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InfoResp); i {
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
		file_app_service_user_ban_api_grpc_v1_ban_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddReq); i {
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
		file_app_service_user_ban_api_grpc_v1_ban_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddResp); i {
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
		file_app_service_user_ban_api_grpc_v1_ban_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoveReq); i {
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
		file_app_service_user_ban_api_grpc_v1_ban_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoveResp); i {
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
			RawDescriptor: file_app_service_user_ban_api_grpc_v1_ban_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_app_service_user_ban_api_grpc_v1_ban_proto_goTypes,
		DependencyIndexes: file_app_service_user_ban_api_grpc_v1_ban_proto_depIdxs,
		MessageInfos:      file_app_service_user_ban_api_grpc_v1_ban_proto_msgTypes,
	}.Build()
	File_app_service_user_ban_api_grpc_v1_ban_proto = out.File
	file_app_service_user_ban_api_grpc_v1_ban_proto_rawDesc = nil
	file_app_service_user_ban_api_grpc_v1_ban_proto_goTypes = nil
	file_app_service_user_ban_api_grpc_v1_ban_proto_depIdxs = nil
}
