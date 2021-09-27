// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.12.4
// source: app/service/user/account/api/grpc/v1/account.proto

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

	SteamId   int64  `protobuf:"varint,1,opt,name=steam_id,json=steamId,proto3" json:"steam_id,omitempty"`
	Username  string `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	FirstJoin int64  `protobuf:"varint,3,opt,name=first_join,json=firstJoin,proto3" json:"first_join,omitempty"`
}

func (x *Info) Reset() {
	*x = Info{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_service_user_account_api_grpc_v1_account_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Info) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Info) ProtoMessage() {}

func (x *Info) ProtoReflect() protoreflect.Message {
	mi := &file_app_service_user_account_api_grpc_v1_account_proto_msgTypes[0]
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
	return file_app_service_user_account_api_grpc_v1_account_proto_rawDescGZIP(), []int{0}
}

func (x *Info) GetSteamId() int64 {
	if x != nil {
		return x.SteamId
	}
	return 0
}

func (x *Info) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *Info) GetFirstJoin() int64 {
	if x != nil {
		return x.FirstJoin
	}
	return 0
}

type UIDReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SteamId int64 `protobuf:"varint,1,opt,name=steam_id,json=steamId,proto3" json:"steam_id,omitempty"`
}

func (x *UIDReq) Reset() {
	*x = UIDReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_service_user_account_api_grpc_v1_account_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UIDReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UIDReq) ProtoMessage() {}

func (x *UIDReq) ProtoReflect() protoreflect.Message {
	mi := &file_app_service_user_account_api_grpc_v1_account_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UIDReq.ProtoReflect.Descriptor instead.
func (*UIDReq) Descriptor() ([]byte, []int) {
	return file_app_service_user_account_api_grpc_v1_account_proto_rawDescGZIP(), []int{1}
}

func (x *UIDReq) GetSteamId() int64 {
	if x != nil {
		return x.SteamId
	}
	return 0
}

type UIDResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid int64 `protobuf:"varint,1,opt,name=uid,proto3" json:"uid,omitempty"`
}

func (x *UIDResp) Reset() {
	*x = UIDResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_service_user_account_api_grpc_v1_account_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UIDResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UIDResp) ProtoMessage() {}

func (x *UIDResp) ProtoReflect() protoreflect.Message {
	mi := &file_app_service_user_account_api_grpc_v1_account_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UIDResp.ProtoReflect.Descriptor instead.
func (*UIDResp) Descriptor() ([]byte, []int) {
	return file_app_service_user_account_api_grpc_v1_account_proto_rawDescGZIP(), []int{2}
}

func (x *UIDResp) GetUid() int64 {
	if x != nil {
		return x.Uid
	}
	return 0
}

type InfoReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid int64 `protobuf:"varint,1,opt,name=uid,proto3" json:"uid,omitempty"`
}

func (x *InfoReq) Reset() {
	*x = InfoReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_service_user_account_api_grpc_v1_account_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InfoReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InfoReq) ProtoMessage() {}

func (x *InfoReq) ProtoReflect() protoreflect.Message {
	mi := &file_app_service_user_account_api_grpc_v1_account_proto_msgTypes[3]
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
	return file_app_service_user_account_api_grpc_v1_account_proto_rawDescGZIP(), []int{3}
}

func (x *InfoReq) GetUid() int64 {
	if x != nil {
		return x.Uid
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
		mi := &file_app_service_user_account_api_grpc_v1_account_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InfoResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InfoResp) ProtoMessage() {}

func (x *InfoResp) ProtoReflect() protoreflect.Message {
	mi := &file_app_service_user_account_api_grpc_v1_account_proto_msgTypes[4]
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
	return file_app_service_user_account_api_grpc_v1_account_proto_rawDescGZIP(), []int{4}
}

func (x *InfoResp) GetInfo() *Info {
	if x != nil {
		return x.Info
	}
	return nil
}

type RegisterReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SteamId int64 `protobuf:"varint,1,opt,name=steam_id,json=steamId,proto3" json:"steam_id,omitempty"`
}

func (x *RegisterReq) Reset() {
	*x = RegisterReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_service_user_account_api_grpc_v1_account_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterReq) ProtoMessage() {}

func (x *RegisterReq) ProtoReflect() protoreflect.Message {
	mi := &file_app_service_user_account_api_grpc_v1_account_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterReq.ProtoReflect.Descriptor instead.
func (*RegisterReq) Descriptor() ([]byte, []int) {
	return file_app_service_user_account_api_grpc_v1_account_proto_rawDescGZIP(), []int{5}
}

func (x *RegisterReq) GetSteamId() int64 {
	if x != nil {
		return x.SteamId
	}
	return 0
}

type RegisterResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid int64 `protobuf:"varint,1,opt,name=uid,proto3" json:"uid,omitempty"`
}

func (x *RegisterResp) Reset() {
	*x = RegisterResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_service_user_account_api_grpc_v1_account_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterResp) ProtoMessage() {}

func (x *RegisterResp) ProtoReflect() protoreflect.Message {
	mi := &file_app_service_user_account_api_grpc_v1_account_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterResp.ProtoReflect.Descriptor instead.
func (*RegisterResp) Descriptor() ([]byte, []int) {
	return file_app_service_user_account_api_grpc_v1_account_proto_rawDescGZIP(), []int{6}
}

func (x *RegisterResp) GetUid() int64 {
	if x != nil {
		return x.Uid
	}
	return 0
}

type ChangeNameReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid      int64  `protobuf:"varint,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Username string `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
}

func (x *ChangeNameReq) Reset() {
	*x = ChangeNameReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_service_user_account_api_grpc_v1_account_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChangeNameReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChangeNameReq) ProtoMessage() {}

func (x *ChangeNameReq) ProtoReflect() protoreflect.Message {
	mi := &file_app_service_user_account_api_grpc_v1_account_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChangeNameReq.ProtoReflect.Descriptor instead.
func (*ChangeNameReq) Descriptor() ([]byte, []int) {
	return file_app_service_user_account_api_grpc_v1_account_proto_rawDescGZIP(), []int{7}
}

func (x *ChangeNameReq) GetUid() int64 {
	if x != nil {
		return x.Uid
	}
	return 0
}

func (x *ChangeNameReq) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

type ChangeNameResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ChangeNameResp) Reset() {
	*x = ChangeNameResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_service_user_account_api_grpc_v1_account_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChangeNameResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChangeNameResp) ProtoMessage() {}

func (x *ChangeNameResp) ProtoReflect() protoreflect.Message {
	mi := &file_app_service_user_account_api_grpc_v1_account_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChangeNameResp.ProtoReflect.Descriptor instead.
func (*ChangeNameResp) Descriptor() ([]byte, []int) {
	return file_app_service_user_account_api_grpc_v1_account_proto_rawDescGZIP(), []int{8}
}

type GetAllUIDReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetAllUIDReq) Reset() {
	*x = GetAllUIDReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_service_user_account_api_grpc_v1_account_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllUIDReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllUIDReq) ProtoMessage() {}

func (x *GetAllUIDReq) ProtoReflect() protoreflect.Message {
	mi := &file_app_service_user_account_api_grpc_v1_account_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllUIDReq.ProtoReflect.Descriptor instead.
func (*GetAllUIDReq) Descriptor() ([]byte, []int) {
	return file_app_service_user_account_api_grpc_v1_account_proto_rawDescGZIP(), []int{9}
}

type GetAllUIDResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid []int64 `protobuf:"varint,1,rep,packed,name=uid,proto3" json:"uid,omitempty"`
}

func (x *GetAllUIDResp) Reset() {
	*x = GetAllUIDResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_service_user_account_api_grpc_v1_account_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllUIDResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllUIDResp) ProtoMessage() {}

func (x *GetAllUIDResp) ProtoReflect() protoreflect.Message {
	mi := &file_app_service_user_account_api_grpc_v1_account_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllUIDResp.ProtoReflect.Descriptor instead.
func (*GetAllUIDResp) Descriptor() ([]byte, []int) {
	return file_app_service_user_account_api_grpc_v1_account_proto_rawDescGZIP(), []int{10}
}

func (x *GetAllUIDResp) GetUid() []int64 {
	if x != nil {
		return x.Uid
	}
	return nil
}

var File_app_service_user_account_api_grpc_v1_account_proto protoreflect.FileDescriptor

var file_app_service_user_account_api_grpc_v1_account_proto_rawDesc = []byte{
	0x0a, 0x32, 0x61, 0x70, 0x70, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x75, 0x73,
	0x65, 0x72, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67,
	0x72, 0x70, 0x63, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1b, 0x6e, 0x63, 0x73, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x76,
	0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e,
	0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x5c, 0x0a, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x74, 0x65, 0x61, 0x6d,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x73, 0x74, 0x65, 0x61, 0x6d,
	0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1d,
	0x0a, 0x0a, 0x66, 0x69, 0x72, 0x73, 0x74, 0x5f, 0x6a, 0x6f, 0x69, 0x6e, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x4a, 0x6f, 0x69, 0x6e, 0x22, 0x23, 0x0a,
	0x06, 0x55, 0x49, 0x44, 0x52, 0x65, 0x71, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x74, 0x65, 0x61, 0x6d,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x73, 0x74, 0x65, 0x61, 0x6d,
	0x49, 0x64, 0x22, 0x1b, 0x0a, 0x07, 0x55, 0x49, 0x44, 0x52, 0x65, 0x73, 0x70, 0x12, 0x10, 0x0a,
	0x03, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x75, 0x69, 0x64, 0x22,
	0x1b, 0x0a, 0x07, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x75, 0x69, 0x64, 0x22, 0x41, 0x0a, 0x08,
	0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x12, 0x35, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x6e, 0x63, 0x73, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x22,
	0x28, 0x0a, 0x0b, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x12, 0x19,
	0x0a, 0x08, 0x73, 0x74, 0x65, 0x61, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x07, 0x73, 0x74, 0x65, 0x61, 0x6d, 0x49, 0x64, 0x22, 0x20, 0x0a, 0x0c, 0x52, 0x65, 0x67,
	0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x75, 0x69, 0x64, 0x22, 0x3d, 0x0a, 0x0d, 0x43,
	0x68, 0x61, 0x6e, 0x67, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x12, 0x10, 0x0a, 0x03,
	0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x75, 0x69, 0x64, 0x12, 0x1a,
	0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x10, 0x0a, 0x0e, 0x43, 0x68,
	0x61, 0x6e, 0x67, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x70, 0x22, 0x0e, 0x0a, 0x0c,
	0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x55, 0x49, 0x44, 0x52, 0x65, 0x71, 0x22, 0x21, 0x0a, 0x0d,
	0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x55, 0x49, 0x44, 0x52, 0x65, 0x73, 0x70, 0x12, 0x10, 0x0a,
	0x03, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x03, 0x28, 0x03, 0x52, 0x03, 0x75, 0x69, 0x64, 0x32,
	0xe4, 0x04, 0x0a, 0x07, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x76, 0x0a, 0x03, 0x55,
	0x49, 0x44, 0x12, 0x23, 0x2e, 0x6e, 0x63, 0x73, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x76, 0x31,
	0x2e, 0x55, 0x49, 0x44, 0x52, 0x65, 0x71, 0x1a, 0x24, 0x2e, 0x6e, 0x63, 0x73, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x49, 0x44, 0x52, 0x65, 0x73, 0x70, 0x22, 0x24, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x1e, 0x12, 0x1c, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x61, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x2f, 0x75, 0x69, 0x64, 0x2f, 0x7b, 0x73, 0x74, 0x65, 0x61, 0x6d, 0x5f,
	0x69, 0x64, 0x7d, 0x12, 0x70, 0x0a, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x24, 0x2e, 0x6e, 0x63,
	0x73, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x61,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65,
	0x71, 0x1a, 0x25, 0x2e, 0x6e, 0x63, 0x73, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x75, 0x73, 0x65, 0x72, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e,
	0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x22, 0x1b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x15,
	0x12, 0x13, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2f,
	0x7b, 0x75, 0x69, 0x64, 0x7d, 0x12, 0x82, 0x01, 0x0a, 0x08, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74,
	0x65, 0x72, 0x12, 0x28, 0x2e, 0x6e, 0x63, 0x73, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x76, 0x31,
	0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x1a, 0x29, 0x2e, 0x6e,
	0x63, 0x73, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e,
	0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73,
	0x74, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x22, 0x21, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1b, 0x22,
	0x16, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2f, 0x72,
	0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x3a, 0x01, 0x2a, 0x12, 0x85, 0x01, 0x0a, 0x0a, 0x43,
	0x68, 0x61, 0x6e, 0x67, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x2a, 0x2e, 0x6e, 0x63, 0x73, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x61, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x4e, 0x61,
	0x6d, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x2b, 0x2e, 0x6e, 0x63, 0x73, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x2e, 0x76, 0x31, 0x2e, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x22, 0x1e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x18, 0x1a, 0x13, 0x2f, 0x75, 0x73, 0x65,
	0x72, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2f, 0x7b, 0x75, 0x69, 0x64, 0x7d, 0x3a,
	0x01, 0x2a, 0x12, 0x62, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x55, 0x49, 0x44, 0x12,
	0x29, 0x2e, 0x6e, 0x63, 0x73, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x75, 0x73,
	0x65, 0x72, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65,
	0x74, 0x41, 0x6c, 0x6c, 0x55, 0x49, 0x44, 0x52, 0x65, 0x71, 0x1a, 0x2a, 0x2e, 0x6e, 0x63, 0x73,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x61, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x55,
	0x49, 0x44, 0x52, 0x65, 0x73, 0x70, 0x42, 0x2e, 0x5a, 0x2c, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e,
	0x64, 0x2f, 0x61, 0x70, 0x70, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x75, 0x73,
	0x65, 0x72, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67,
	0x72, 0x70, 0x63, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_app_service_user_account_api_grpc_v1_account_proto_rawDescOnce sync.Once
	file_app_service_user_account_api_grpc_v1_account_proto_rawDescData = file_app_service_user_account_api_grpc_v1_account_proto_rawDesc
)

func file_app_service_user_account_api_grpc_v1_account_proto_rawDescGZIP() []byte {
	file_app_service_user_account_api_grpc_v1_account_proto_rawDescOnce.Do(func() {
		file_app_service_user_account_api_grpc_v1_account_proto_rawDescData = protoimpl.X.CompressGZIP(file_app_service_user_account_api_grpc_v1_account_proto_rawDescData)
	})
	return file_app_service_user_account_api_grpc_v1_account_proto_rawDescData
}

var file_app_service_user_account_api_grpc_v1_account_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_app_service_user_account_api_grpc_v1_account_proto_goTypes = []interface{}{
	(*Info)(nil),           // 0: ncs.service.user.account.v1.Info
	(*UIDReq)(nil),         // 1: ncs.service.user.account.v1.UIDReq
	(*UIDResp)(nil),        // 2: ncs.service.user.account.v1.UIDResp
	(*InfoReq)(nil),        // 3: ncs.service.user.account.v1.InfoReq
	(*InfoResp)(nil),       // 4: ncs.service.user.account.v1.InfoResp
	(*RegisterReq)(nil),    // 5: ncs.service.user.account.v1.RegisterReq
	(*RegisterResp)(nil),   // 6: ncs.service.user.account.v1.RegisterResp
	(*ChangeNameReq)(nil),  // 7: ncs.service.user.account.v1.ChangeNameReq
	(*ChangeNameResp)(nil), // 8: ncs.service.user.account.v1.ChangeNameResp
	(*GetAllUIDReq)(nil),   // 9: ncs.service.user.account.v1.GetAllUIDReq
	(*GetAllUIDResp)(nil),  // 10: ncs.service.user.account.v1.GetAllUIDResp
}
var file_app_service_user_account_api_grpc_v1_account_proto_depIdxs = []int32{
	0,  // 0: ncs.service.user.account.v1.InfoResp.info:type_name -> ncs.service.user.account.v1.Info
	1,  // 1: ncs.service.user.account.v1.Account.UID:input_type -> ncs.service.user.account.v1.UIDReq
	3,  // 2: ncs.service.user.account.v1.Account.Info:input_type -> ncs.service.user.account.v1.InfoReq
	5,  // 3: ncs.service.user.account.v1.Account.Register:input_type -> ncs.service.user.account.v1.RegisterReq
	7,  // 4: ncs.service.user.account.v1.Account.ChangeName:input_type -> ncs.service.user.account.v1.ChangeNameReq
	9,  // 5: ncs.service.user.account.v1.Account.GetAllUID:input_type -> ncs.service.user.account.v1.GetAllUIDReq
	2,  // 6: ncs.service.user.account.v1.Account.UID:output_type -> ncs.service.user.account.v1.UIDResp
	4,  // 7: ncs.service.user.account.v1.Account.Info:output_type -> ncs.service.user.account.v1.InfoResp
	6,  // 8: ncs.service.user.account.v1.Account.Register:output_type -> ncs.service.user.account.v1.RegisterResp
	8,  // 9: ncs.service.user.account.v1.Account.ChangeName:output_type -> ncs.service.user.account.v1.ChangeNameResp
	10, // 10: ncs.service.user.account.v1.Account.GetAllUID:output_type -> ncs.service.user.account.v1.GetAllUIDResp
	6,  // [6:11] is the sub-list for method output_type
	1,  // [1:6] is the sub-list for method input_type
	1,  // [1:1] is the sub-list for extension type_name
	1,  // [1:1] is the sub-list for extension extendee
	0,  // [0:1] is the sub-list for field type_name
}

func init() { file_app_service_user_account_api_grpc_v1_account_proto_init() }
func file_app_service_user_account_api_grpc_v1_account_proto_init() {
	if File_app_service_user_account_api_grpc_v1_account_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_app_service_user_account_api_grpc_v1_account_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_app_service_user_account_api_grpc_v1_account_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UIDReq); i {
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
		file_app_service_user_account_api_grpc_v1_account_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UIDResp); i {
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
		file_app_service_user_account_api_grpc_v1_account_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
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
		file_app_service_user_account_api_grpc_v1_account_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
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
		file_app_service_user_account_api_grpc_v1_account_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterReq); i {
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
		file_app_service_user_account_api_grpc_v1_account_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterResp); i {
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
		file_app_service_user_account_api_grpc_v1_account_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChangeNameReq); i {
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
		file_app_service_user_account_api_grpc_v1_account_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChangeNameResp); i {
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
		file_app_service_user_account_api_grpc_v1_account_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllUIDReq); i {
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
		file_app_service_user_account_api_grpc_v1_account_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllUIDResp); i {
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
			RawDescriptor: file_app_service_user_account_api_grpc_v1_account_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_app_service_user_account_api_grpc_v1_account_proto_goTypes,
		DependencyIndexes: file_app_service_user_account_api_grpc_v1_account_proto_depIdxs,
		MessageInfos:      file_app_service_user_account_api_grpc_v1_account_proto_msgTypes,
	}.Build()
	File_app_service_user_account_api_grpc_v1_account_proto = out.File
	file_app_service_user_account_api_grpc_v1_account_proto_rawDesc = nil
	file_app_service_user_account_api_grpc_v1_account_proto_goTypes = nil
	file_app_service_user_account_api_grpc_v1_account_proto_depIdxs = nil
}
