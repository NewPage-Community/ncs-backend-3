// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.21.12
// source: app/game/skin/api/grpc/v1/skin.proto

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

type SkinInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name     string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	SkinPath string `protobuf:"bytes,4,opt,name=skin_path,json=skinPath,proto3" json:"skin_path,omitempty"`
	ArmPath  string `protobuf:"bytes,5,opt,name=arm_path,json=armPath,proto3" json:"arm_path,omitempty"`
	Price    int32  `protobuf:"varint,6,opt,name=price,proto3" json:"price,omitempty"`
}

func (x *SkinInfo) Reset() {
	*x = SkinInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_game_skin_api_grpc_v1_skin_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SkinInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SkinInfo) ProtoMessage() {}

func (x *SkinInfo) ProtoReflect() protoreflect.Message {
	mi := &file_app_game_skin_api_grpc_v1_skin_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SkinInfo.ProtoReflect.Descriptor instead.
func (*SkinInfo) Descriptor() ([]byte, []int) {
	return file_app_game_skin_api_grpc_v1_skin_proto_rawDescGZIP(), []int{0}
}

func (x *SkinInfo) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *SkinInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *SkinInfo) GetSkinPath() string {
	if x != nil {
		return x.SkinPath
	}
	return ""
}

func (x *SkinInfo) GetArmPath() string {
	if x != nil {
		return x.ArmPath
	}
	return ""
}

func (x *SkinInfo) GetPrice() int32 {
	if x != nil {
		return x.Price
	}
	return 0
}

type GetInfoReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid int64 `protobuf:"varint,1,opt,name=uid,proto3" json:"uid,omitempty"`
}

func (x *GetInfoReq) Reset() {
	*x = GetInfoReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_game_skin_api_grpc_v1_skin_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetInfoReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetInfoReq) ProtoMessage() {}

func (x *GetInfoReq) ProtoReflect() protoreflect.Message {
	mi := &file_app_game_skin_api_grpc_v1_skin_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetInfoReq.ProtoReflect.Descriptor instead.
func (*GetInfoReq) Descriptor() ([]byte, []int) {
	return file_app_game_skin_api_grpc_v1_skin_proto_rawDescGZIP(), []int{1}
}

func (x *GetInfoReq) GetUid() int64 {
	if x != nil {
		return x.Uid
	}
	return 0
}

type GetInfoResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid      int64 `protobuf:"varint,1,opt,name=uid,proto3" json:"uid,omitempty"`
	UsedSkin int32 `protobuf:"varint,2,opt,name=used_skin,json=usedSkin,proto3" json:"used_skin,omitempty"`
}

func (x *GetInfoResp) Reset() {
	*x = GetInfoResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_game_skin_api_grpc_v1_skin_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetInfoResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetInfoResp) ProtoMessage() {}

func (x *GetInfoResp) ProtoReflect() protoreflect.Message {
	mi := &file_app_game_skin_api_grpc_v1_skin_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetInfoResp.ProtoReflect.Descriptor instead.
func (*GetInfoResp) Descriptor() ([]byte, []int) {
	return file_app_game_skin_api_grpc_v1_skin_proto_rawDescGZIP(), []int{2}
}

func (x *GetInfoResp) GetUid() int64 {
	if x != nil {
		return x.Uid
	}
	return 0
}

func (x *GetInfoResp) GetUsedSkin() int32 {
	if x != nil {
		return x.UsedSkin
	}
	return 0
}

type GetSkinsReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetSkinsReq) Reset() {
	*x = GetSkinsReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_game_skin_api_grpc_v1_skin_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSkinsReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSkinsReq) ProtoMessage() {}

func (x *GetSkinsReq) ProtoReflect() protoreflect.Message {
	mi := &file_app_game_skin_api_grpc_v1_skin_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSkinsReq.ProtoReflect.Descriptor instead.
func (*GetSkinsReq) Descriptor() ([]byte, []int) {
	return file_app_game_skin_api_grpc_v1_skin_proto_rawDescGZIP(), []int{3}
}

type GetSkinsResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info []*SkinInfo `protobuf:"bytes,1,rep,name=info,proto3" json:"info,omitempty"`
}

func (x *GetSkinsResp) Reset() {
	*x = GetSkinsResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_game_skin_api_grpc_v1_skin_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSkinsResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSkinsResp) ProtoMessage() {}

func (x *GetSkinsResp) ProtoReflect() protoreflect.Message {
	mi := &file_app_game_skin_api_grpc_v1_skin_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSkinsResp.ProtoReflect.Descriptor instead.
func (*GetSkinsResp) Descriptor() ([]byte, []int) {
	return file_app_game_skin_api_grpc_v1_skin_proto_rawDescGZIP(), []int{4}
}

func (x *GetSkinsResp) GetInfo() []*SkinInfo {
	if x != nil {
		return x.Info
	}
	return nil
}

type SetUsedSkinReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid      int64 `protobuf:"varint,1,opt,name=uid,proto3" json:"uid,omitempty"`
	UsedSkin int32 `protobuf:"varint,2,opt,name=used_skin,json=usedSkin,proto3" json:"used_skin,omitempty"`
}

func (x *SetUsedSkinReq) Reset() {
	*x = SetUsedSkinReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_game_skin_api_grpc_v1_skin_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetUsedSkinReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetUsedSkinReq) ProtoMessage() {}

func (x *SetUsedSkinReq) ProtoReflect() protoreflect.Message {
	mi := &file_app_game_skin_api_grpc_v1_skin_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetUsedSkinReq.ProtoReflect.Descriptor instead.
func (*SetUsedSkinReq) Descriptor() ([]byte, []int) {
	return file_app_game_skin_api_grpc_v1_skin_proto_rawDescGZIP(), []int{5}
}

func (x *SetUsedSkinReq) GetUid() int64 {
	if x != nil {
		return x.Uid
	}
	return 0
}

func (x *SetUsedSkinReq) GetUsedSkin() int32 {
	if x != nil {
		return x.UsedSkin
	}
	return 0
}

type SetUsedSkinResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SetUsedSkinResp) Reset() {
	*x = SetUsedSkinResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_game_skin_api_grpc_v1_skin_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetUsedSkinResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetUsedSkinResp) ProtoMessage() {}

func (x *SetUsedSkinResp) ProtoReflect() protoreflect.Message {
	mi := &file_app_game_skin_api_grpc_v1_skin_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetUsedSkinResp.ProtoReflect.Descriptor instead.
func (*SetUsedSkinResp) Descriptor() ([]byte, []int) {
	return file_app_game_skin_api_grpc_v1_skin_proto_rawDescGZIP(), []int{6}
}

var File_app_game_skin_api_grpc_v1_skin_proto protoreflect.FileDescriptor

var file_app_game_skin_api_grpc_v1_skin_proto_rawDesc = []byte{
	0x0a, 0x24, 0x61, 0x70, 0x70, 0x2f, 0x67, 0x61, 0x6d, 0x65, 0x2f, 0x73, 0x6b, 0x69, 0x6e, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x6b, 0x69, 0x6e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x6e, 0x63, 0x73, 0x2e, 0x67, 0x61, 0x6d, 0x65,
	0x2e, 0x73, 0x6b, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x7c, 0x0a, 0x08, 0x53, 0x6b, 0x69, 0x6e, 0x49, 0x6e,
	0x66, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x6b, 0x69, 0x6e, 0x5f, 0x70,
	0x61, 0x74, 0x68, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x6b, 0x69, 0x6e, 0x50,
	0x61, 0x74, 0x68, 0x12, 0x19, 0x0a, 0x08, 0x61, 0x72, 0x6d, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x72, 0x6d, 0x50, 0x61, 0x74, 0x68, 0x12, 0x14,
	0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x70,
	0x72, 0x69, 0x63, 0x65, 0x22, 0x1e, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52,
	0x65, 0x71, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x03, 0x75, 0x69, 0x64, 0x22, 0x3c, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52,
	0x65, 0x73, 0x70, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x03, 0x75, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x75, 0x73, 0x65, 0x64, 0x5f, 0x73, 0x6b,
	0x69, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x75, 0x73, 0x65, 0x64, 0x53, 0x6b,
	0x69, 0x6e, 0x22, 0x0d, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x53, 0x6b, 0x69, 0x6e, 0x73, 0x52, 0x65,
	0x71, 0x22, 0x3e, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x53, 0x6b, 0x69, 0x6e, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x12, 0x2e, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x6e, 0x63, 0x73, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x2e, 0x73, 0x6b, 0x69, 0x6e, 0x2e,
	0x76, 0x31, 0x2e, 0x53, 0x6b, 0x69, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x69, 0x6e, 0x66,
	0x6f, 0x22, 0x3f, 0x0a, 0x0e, 0x53, 0x65, 0x74, 0x55, 0x73, 0x65, 0x64, 0x53, 0x6b, 0x69, 0x6e,
	0x52, 0x65, 0x71, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x03, 0x75, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x75, 0x73, 0x65, 0x64, 0x5f, 0x73, 0x6b,
	0x69, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x75, 0x73, 0x65, 0x64, 0x53, 0x6b,
	0x69, 0x6e, 0x22, 0x11, 0x0a, 0x0f, 0x53, 0x65, 0x74, 0x55, 0x73, 0x65, 0x64, 0x53, 0x6b, 0x69,
	0x6e, 0x52, 0x65, 0x73, 0x70, 0x32, 0xbd, 0x02, 0x0a, 0x04, 0x53, 0x6b, 0x69, 0x6e, 0x12, 0x5d,
	0x0a, 0x08, 0x47, 0x65, 0x74, 0x53, 0x6b, 0x69, 0x6e, 0x73, 0x12, 0x1d, 0x2e, 0x6e, 0x63, 0x73,
	0x2e, 0x67, 0x61, 0x6d, 0x65, 0x2e, 0x73, 0x6b, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65,
	0x74, 0x53, 0x6b, 0x69, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x1a, 0x1e, 0x2e, 0x6e, 0x63, 0x73, 0x2e,
	0x67, 0x61, 0x6d, 0x65, 0x2e, 0x73, 0x6b, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74,
	0x53, 0x6b, 0x69, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x22, 0x12, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x0c, 0x12, 0x0a, 0x2f, 0x67, 0x61, 0x6d, 0x65, 0x2f, 0x73, 0x6b, 0x69, 0x6e, 0x12, 0x60, 0x0a,
	0x07, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1c, 0x2e, 0x6e, 0x63, 0x73, 0x2e, 0x67,
	0x61, 0x6d, 0x65, 0x2e, 0x73, 0x6b, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x49,
	0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x1a, 0x1d, 0x2e, 0x6e, 0x63, 0x73, 0x2e, 0x67, 0x61, 0x6d,
	0x65, 0x2e, 0x73, 0x6b, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x66,
	0x6f, 0x52, 0x65, 0x73, 0x70, 0x22, 0x18, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x12, 0x12, 0x10, 0x2f,
	0x67, 0x61, 0x6d, 0x65, 0x2f, 0x73, 0x6b, 0x69, 0x6e, 0x2f, 0x7b, 0x75, 0x69, 0x64, 0x7d, 0x12,
	0x74, 0x0a, 0x0b, 0x53, 0x65, 0x74, 0x55, 0x73, 0x65, 0x64, 0x53, 0x6b, 0x69, 0x6e, 0x12, 0x20,
	0x2e, 0x6e, 0x63, 0x73, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x2e, 0x73, 0x6b, 0x69, 0x6e, 0x2e, 0x76,
	0x31, 0x2e, 0x53, 0x65, 0x74, 0x55, 0x73, 0x65, 0x64, 0x53, 0x6b, 0x69, 0x6e, 0x52, 0x65, 0x71,
	0x1a, 0x21, 0x2e, 0x6e, 0x63, 0x73, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x2e, 0x73, 0x6b, 0x69, 0x6e,
	0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x74, 0x55, 0x73, 0x65, 0x64, 0x53, 0x6b, 0x69, 0x6e, 0x52,
	0x65, 0x73, 0x70, 0x22, 0x20, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1a, 0x22, 0x15, 0x2f, 0x67, 0x61,
	0x6d, 0x65, 0x2f, 0x73, 0x6b, 0x69, 0x6e, 0x2f, 0x7b, 0x75, 0x69, 0x64, 0x7d, 0x2f, 0x75, 0x73,
	0x65, 0x64, 0x3a, 0x01, 0x2a, 0x42, 0x23, 0x5a, 0x21, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64,
	0x2f, 0x61, 0x70, 0x70, 0x2f, 0x67, 0x61, 0x6d, 0x65, 0x2f, 0x73, 0x6b, 0x69, 0x6e, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_app_game_skin_api_grpc_v1_skin_proto_rawDescOnce sync.Once
	file_app_game_skin_api_grpc_v1_skin_proto_rawDescData = file_app_game_skin_api_grpc_v1_skin_proto_rawDesc
)

func file_app_game_skin_api_grpc_v1_skin_proto_rawDescGZIP() []byte {
	file_app_game_skin_api_grpc_v1_skin_proto_rawDescOnce.Do(func() {
		file_app_game_skin_api_grpc_v1_skin_proto_rawDescData = protoimpl.X.CompressGZIP(file_app_game_skin_api_grpc_v1_skin_proto_rawDescData)
	})
	return file_app_game_skin_api_grpc_v1_skin_proto_rawDescData
}

var file_app_game_skin_api_grpc_v1_skin_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_app_game_skin_api_grpc_v1_skin_proto_goTypes = []interface{}{
	(*SkinInfo)(nil),        // 0: ncs.game.skin.v1.SkinInfo
	(*GetInfoReq)(nil),      // 1: ncs.game.skin.v1.GetInfoReq
	(*GetInfoResp)(nil),     // 2: ncs.game.skin.v1.GetInfoResp
	(*GetSkinsReq)(nil),     // 3: ncs.game.skin.v1.GetSkinsReq
	(*GetSkinsResp)(nil),    // 4: ncs.game.skin.v1.GetSkinsResp
	(*SetUsedSkinReq)(nil),  // 5: ncs.game.skin.v1.SetUsedSkinReq
	(*SetUsedSkinResp)(nil), // 6: ncs.game.skin.v1.SetUsedSkinResp
}
var file_app_game_skin_api_grpc_v1_skin_proto_depIdxs = []int32{
	0, // 0: ncs.game.skin.v1.GetSkinsResp.info:type_name -> ncs.game.skin.v1.SkinInfo
	3, // 1: ncs.game.skin.v1.Skin.GetSkins:input_type -> ncs.game.skin.v1.GetSkinsReq
	1, // 2: ncs.game.skin.v1.Skin.GetInfo:input_type -> ncs.game.skin.v1.GetInfoReq
	5, // 3: ncs.game.skin.v1.Skin.SetUsedSkin:input_type -> ncs.game.skin.v1.SetUsedSkinReq
	4, // 4: ncs.game.skin.v1.Skin.GetSkins:output_type -> ncs.game.skin.v1.GetSkinsResp
	2, // 5: ncs.game.skin.v1.Skin.GetInfo:output_type -> ncs.game.skin.v1.GetInfoResp
	6, // 6: ncs.game.skin.v1.Skin.SetUsedSkin:output_type -> ncs.game.skin.v1.SetUsedSkinResp
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_app_game_skin_api_grpc_v1_skin_proto_init() }
func file_app_game_skin_api_grpc_v1_skin_proto_init() {
	if File_app_game_skin_api_grpc_v1_skin_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_app_game_skin_api_grpc_v1_skin_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SkinInfo); i {
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
		file_app_game_skin_api_grpc_v1_skin_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetInfoReq); i {
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
		file_app_game_skin_api_grpc_v1_skin_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetInfoResp); i {
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
		file_app_game_skin_api_grpc_v1_skin_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSkinsReq); i {
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
		file_app_game_skin_api_grpc_v1_skin_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSkinsResp); i {
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
		file_app_game_skin_api_grpc_v1_skin_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetUsedSkinReq); i {
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
		file_app_game_skin_api_grpc_v1_skin_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetUsedSkinResp); i {
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
			RawDescriptor: file_app_game_skin_api_grpc_v1_skin_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_app_game_skin_api_grpc_v1_skin_proto_goTypes,
		DependencyIndexes: file_app_game_skin_api_grpc_v1_skin_proto_depIdxs,
		MessageInfos:      file_app_game_skin_api_grpc_v1_skin_proto_msgTypes,
	}.Build()
	File_app_game_skin_api_grpc_v1_skin_proto = out.File
	file_app_game_skin_api_grpc_v1_skin_proto_rawDesc = nil
	file_app_game_skin_api_grpc_v1_skin_proto_goTypes = nil
	file_app_game_skin_api_grpc_v1_skin_proto_depIdxs = nil
}
