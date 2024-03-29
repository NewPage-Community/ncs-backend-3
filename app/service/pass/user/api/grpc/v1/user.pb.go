// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.21.12
// source: app/service/pass/user/api/grpc/v1/user.proto

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

	Uid      int64 `protobuf:"varint,1,opt,name=uid,proto3" json:"uid,omitempty"`
	PassType int32 `protobuf:"varint,2,opt,name=pass_type,json=passType,proto3" json:"pass_type,omitempty"`
	Point    int32 `protobuf:"varint,3,opt,name=point,proto3" json:"point,omitempty"`
}

func (x *Info) Reset() {
	*x = Info{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_service_pass_user_api_grpc_v1_user_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Info) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Info) ProtoMessage() {}

func (x *Info) ProtoReflect() protoreflect.Message {
	mi := &file_app_service_pass_user_api_grpc_v1_user_proto_msgTypes[0]
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
	return file_app_service_pass_user_api_grpc_v1_user_proto_rawDescGZIP(), []int{0}
}

func (x *Info) GetUid() int64 {
	if x != nil {
		return x.Uid
	}
	return 0
}

func (x *Info) GetPassType() int32 {
	if x != nil {
		return x.PassType
	}
	return 0
}

func (x *Info) GetPoint() int32 {
	if x != nil {
		return x.Point
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
		mi := &file_app_service_pass_user_api_grpc_v1_user_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InfoReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InfoReq) ProtoMessage() {}

func (x *InfoReq) ProtoReflect() protoreflect.Message {
	mi := &file_app_service_pass_user_api_grpc_v1_user_proto_msgTypes[1]
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
	return file_app_service_pass_user_api_grpc_v1_user_proto_rawDescGZIP(), []int{1}
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
		mi := &file_app_service_pass_user_api_grpc_v1_user_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InfoResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InfoResp) ProtoMessage() {}

func (x *InfoResp) ProtoReflect() protoreflect.Message {
	mi := &file_app_service_pass_user_api_grpc_v1_user_proto_msgTypes[2]
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
	return file_app_service_pass_user_api_grpc_v1_user_proto_rawDescGZIP(), []int{2}
}

func (x *InfoResp) GetInfo() *Info {
	if x != nil {
		return x.Info
	}
	return nil
}

type AddPoints struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid   int64 `protobuf:"varint,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Point int32 `protobuf:"varint,2,opt,name=point,proto3" json:"point,omitempty"`
}

func (x *AddPoints) Reset() {
	*x = AddPoints{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_service_pass_user_api_grpc_v1_user_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddPoints) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddPoints) ProtoMessage() {}

func (x *AddPoints) ProtoReflect() protoreflect.Message {
	mi := &file_app_service_pass_user_api_grpc_v1_user_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddPoints.ProtoReflect.Descriptor instead.
func (*AddPoints) Descriptor() ([]byte, []int) {
	return file_app_service_pass_user_api_grpc_v1_user_proto_rawDescGZIP(), []int{3}
}

func (x *AddPoints) GetUid() int64 {
	if x != nil {
		return x.Uid
	}
	return 0
}

func (x *AddPoints) GetPoint() int32 {
	if x != nil {
		return x.Point
	}
	return 0
}

type AddPointsReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Add []*AddPoints `protobuf:"bytes,1,rep,name=add,proto3" json:"add,omitempty"`
}

func (x *AddPointsReq) Reset() {
	*x = AddPointsReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_service_pass_user_api_grpc_v1_user_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddPointsReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddPointsReq) ProtoMessage() {}

func (x *AddPointsReq) ProtoReflect() protoreflect.Message {
	mi := &file_app_service_pass_user_api_grpc_v1_user_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddPointsReq.ProtoReflect.Descriptor instead.
func (*AddPointsReq) Descriptor() ([]byte, []int) {
	return file_app_service_pass_user_api_grpc_v1_user_proto_rawDescGZIP(), []int{4}
}

func (x *AddPointsReq) GetAdd() []*AddPoints {
	if x != nil {
		return x.Add
	}
	return nil
}

type AddPointsResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *AddPointsResp) Reset() {
	*x = AddPointsResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_service_pass_user_api_grpc_v1_user_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddPointsResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddPointsResp) ProtoMessage() {}

func (x *AddPointsResp) ProtoReflect() protoreflect.Message {
	mi := &file_app_service_pass_user_api_grpc_v1_user_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddPointsResp.ProtoReflect.Descriptor instead.
func (*AddPointsResp) Descriptor() ([]byte, []int) {
	return file_app_service_pass_user_api_grpc_v1_user_proto_rawDescGZIP(), []int{5}
}

type UpgradePassReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid  int64 `protobuf:"varint,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Type int32 `protobuf:"varint,2,opt,name=type,proto3" json:"type,omitempty"`
}

func (x *UpgradePassReq) Reset() {
	*x = UpgradePassReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_service_pass_user_api_grpc_v1_user_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpgradePassReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpgradePassReq) ProtoMessage() {}

func (x *UpgradePassReq) ProtoReflect() protoreflect.Message {
	mi := &file_app_service_pass_user_api_grpc_v1_user_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpgradePassReq.ProtoReflect.Descriptor instead.
func (*UpgradePassReq) Descriptor() ([]byte, []int) {
	return file_app_service_pass_user_api_grpc_v1_user_proto_rawDescGZIP(), []int{6}
}

func (x *UpgradePassReq) GetUid() int64 {
	if x != nil {
		return x.Uid
	}
	return 0
}

func (x *UpgradePassReq) GetType() int32 {
	if x != nil {
		return x.Type
	}
	return 0
}

type UpgradePassResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpgradePassResp) Reset() {
	*x = UpgradePassResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_service_pass_user_api_grpc_v1_user_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpgradePassResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpgradePassResp) ProtoMessage() {}

func (x *UpgradePassResp) ProtoReflect() protoreflect.Message {
	mi := &file_app_service_pass_user_api_grpc_v1_user_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpgradePassResp.ProtoReflect.Descriptor instead.
func (*UpgradePassResp) Descriptor() ([]byte, []int) {
	return file_app_service_pass_user_api_grpc_v1_user_proto_rawDescGZIP(), []int{7}
}

var File_app_service_pass_user_api_grpc_v1_user_proto protoreflect.FileDescriptor

var file_app_service_pass_user_api_grpc_v1_user_proto_rawDesc = []byte{
	0x0a, 0x2c, 0x61, 0x70, 0x70, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x70, 0x61,
	0x73, 0x73, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x72, 0x70, 0x63,
	0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x18,
	0x6e, 0x63, 0x73, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x61, 0x73, 0x73,
	0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x4b, 0x0a, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x10,
	0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x75, 0x69, 0x64,
	0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x73, 0x73, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x54, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x70, 0x6f,
	0x69, 0x6e, 0x74, 0x22, 0x1b, 0x0a, 0x07, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x12, 0x10,
	0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x75, 0x69, 0x64,
	0x22, 0x3e, 0x0a, 0x08, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x12, 0x32, 0x0a, 0x04,
	0x69, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x6e, 0x63, 0x73,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x61, 0x73, 0x73, 0x2e, 0x75, 0x73,
	0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f,
	0x22, 0x33, 0x0a, 0x09, 0x41, 0x64, 0x64, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x12, 0x10, 0x0a,
	0x03, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x75, 0x69, 0x64, 0x12,
	0x14, 0x0a, 0x05, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05,
	0x70, 0x6f, 0x69, 0x6e, 0x74, 0x22, 0x45, 0x0a, 0x0c, 0x41, 0x64, 0x64, 0x50, 0x6f, 0x69, 0x6e,
	0x74, 0x73, 0x52, 0x65, 0x71, 0x12, 0x35, 0x0a, 0x03, 0x61, 0x64, 0x64, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x23, 0x2e, 0x6e, 0x63, 0x73, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x70, 0x61, 0x73, 0x73, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x64,
	0x64, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x52, 0x03, 0x61, 0x64, 0x64, 0x22, 0x0f, 0x0a, 0x0d,
	0x41, 0x64, 0x64, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x22, 0x36, 0x0a,
	0x0e, 0x55, 0x70, 0x67, 0x72, 0x61, 0x64, 0x65, 0x50, 0x61, 0x73, 0x73, 0x52, 0x65, 0x71, 0x12,
	0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x75, 0x69,
	0x64, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x22, 0x11, 0x0a, 0x0f, 0x55, 0x70, 0x67, 0x72, 0x61, 0x64, 0x65,
	0x50, 0x61, 0x73, 0x73, 0x52, 0x65, 0x73, 0x70, 0x32, 0xf8, 0x02, 0x0a, 0x04, 0x55, 0x73, 0x65,
	0x72, 0x12, 0x67, 0x0a, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x21, 0x2e, 0x6e, 0x63, 0x73, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x61, 0x73, 0x73, 0x2e, 0x75, 0x73, 0x65,
	0x72, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x1a, 0x22, 0x2e, 0x6e,
	0x63, 0x73, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x61, 0x73, 0x73, 0x2e,
	0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70,
	0x22, 0x18, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x12, 0x12, 0x10, 0x2f, 0x70, 0x61, 0x73, 0x73, 0x2f,
	0x75, 0x73, 0x65, 0x72, 0x2f, 0x7b, 0x75, 0x69, 0x64, 0x7d, 0x12, 0x7d, 0x0a, 0x09, 0x41, 0x64,
	0x64, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x12, 0x26, 0x2e, 0x6e, 0x63, 0x73, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x61, 0x73, 0x73, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e,
	0x76, 0x31, 0x2e, 0x41, 0x64, 0x64, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x1a,
	0x27, 0x2e, 0x6e, 0x63, 0x73, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x61,
	0x73, 0x73, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x64, 0x64, 0x50, 0x6f,
	0x69, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x22, 0x1f, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x19,
	0x22, 0x14, 0x2f, 0x70, 0x61, 0x73, 0x73, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x61, 0x64, 0x64,
	0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x3a, 0x01, 0x2a, 0x12, 0x87, 0x01, 0x0a, 0x0b, 0x55, 0x70,
	0x67, 0x72, 0x61, 0x64, 0x65, 0x50, 0x61, 0x73, 0x73, 0x12, 0x28, 0x2e, 0x6e, 0x63, 0x73, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x61, 0x73, 0x73, 0x2e, 0x75, 0x73, 0x65,
	0x72, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x67, 0x72, 0x61, 0x64, 0x65, 0x50, 0x61, 0x73, 0x73,
	0x52, 0x65, 0x71, 0x1a, 0x29, 0x2e, 0x6e, 0x63, 0x73, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x70, 0x61, 0x73, 0x73, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x55,
	0x70, 0x67, 0x72, 0x61, 0x64, 0x65, 0x50, 0x61, 0x73, 0x73, 0x52, 0x65, 0x73, 0x70, 0x22, 0x23,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1d, 0x22, 0x18, 0x2f, 0x70, 0x61, 0x73, 0x73, 0x2f, 0x75, 0x73,
	0x65, 0x72, 0x2f, 0x7b, 0x75, 0x69, 0x64, 0x7d, 0x2f, 0x75, 0x70, 0x67, 0x72, 0x61, 0x64, 0x65,
	0x3a, 0x01, 0x2a, 0x42, 0x2b, 0x5a, 0x29, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2f, 0x61,
	0x70, 0x70, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x70, 0x61, 0x73, 0x73, 0x2f,
	0x75, 0x73, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x76, 0x31,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_app_service_pass_user_api_grpc_v1_user_proto_rawDescOnce sync.Once
	file_app_service_pass_user_api_grpc_v1_user_proto_rawDescData = file_app_service_pass_user_api_grpc_v1_user_proto_rawDesc
)

func file_app_service_pass_user_api_grpc_v1_user_proto_rawDescGZIP() []byte {
	file_app_service_pass_user_api_grpc_v1_user_proto_rawDescOnce.Do(func() {
		file_app_service_pass_user_api_grpc_v1_user_proto_rawDescData = protoimpl.X.CompressGZIP(file_app_service_pass_user_api_grpc_v1_user_proto_rawDescData)
	})
	return file_app_service_pass_user_api_grpc_v1_user_proto_rawDescData
}

var file_app_service_pass_user_api_grpc_v1_user_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_app_service_pass_user_api_grpc_v1_user_proto_goTypes = []interface{}{
	(*Info)(nil),            // 0: ncs.service.pass.user.v1.Info
	(*InfoReq)(nil),         // 1: ncs.service.pass.user.v1.InfoReq
	(*InfoResp)(nil),        // 2: ncs.service.pass.user.v1.InfoResp
	(*AddPoints)(nil),       // 3: ncs.service.pass.user.v1.AddPoints
	(*AddPointsReq)(nil),    // 4: ncs.service.pass.user.v1.AddPointsReq
	(*AddPointsResp)(nil),   // 5: ncs.service.pass.user.v1.AddPointsResp
	(*UpgradePassReq)(nil),  // 6: ncs.service.pass.user.v1.UpgradePassReq
	(*UpgradePassResp)(nil), // 7: ncs.service.pass.user.v1.UpgradePassResp
}
var file_app_service_pass_user_api_grpc_v1_user_proto_depIdxs = []int32{
	0, // 0: ncs.service.pass.user.v1.InfoResp.info:type_name -> ncs.service.pass.user.v1.Info
	3, // 1: ncs.service.pass.user.v1.AddPointsReq.add:type_name -> ncs.service.pass.user.v1.AddPoints
	1, // 2: ncs.service.pass.user.v1.User.Info:input_type -> ncs.service.pass.user.v1.InfoReq
	4, // 3: ncs.service.pass.user.v1.User.AddPoints:input_type -> ncs.service.pass.user.v1.AddPointsReq
	6, // 4: ncs.service.pass.user.v1.User.UpgradePass:input_type -> ncs.service.pass.user.v1.UpgradePassReq
	2, // 5: ncs.service.pass.user.v1.User.Info:output_type -> ncs.service.pass.user.v1.InfoResp
	5, // 6: ncs.service.pass.user.v1.User.AddPoints:output_type -> ncs.service.pass.user.v1.AddPointsResp
	7, // 7: ncs.service.pass.user.v1.User.UpgradePass:output_type -> ncs.service.pass.user.v1.UpgradePassResp
	5, // [5:8] is the sub-list for method output_type
	2, // [2:5] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_app_service_pass_user_api_grpc_v1_user_proto_init() }
func file_app_service_pass_user_api_grpc_v1_user_proto_init() {
	if File_app_service_pass_user_api_grpc_v1_user_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_app_service_pass_user_api_grpc_v1_user_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_app_service_pass_user_api_grpc_v1_user_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
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
		file_app_service_pass_user_api_grpc_v1_user_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
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
		file_app_service_pass_user_api_grpc_v1_user_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddPoints); i {
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
		file_app_service_pass_user_api_grpc_v1_user_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddPointsReq); i {
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
		file_app_service_pass_user_api_grpc_v1_user_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddPointsResp); i {
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
		file_app_service_pass_user_api_grpc_v1_user_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpgradePassReq); i {
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
		file_app_service_pass_user_api_grpc_v1_user_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpgradePassResp); i {
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
			RawDescriptor: file_app_service_pass_user_api_grpc_v1_user_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_app_service_pass_user_api_grpc_v1_user_proto_goTypes,
		DependencyIndexes: file_app_service_pass_user_api_grpc_v1_user_proto_depIdxs,
		MessageInfos:      file_app_service_pass_user_api_grpc_v1_user_proto_msgTypes,
	}.Build()
	File_app_service_pass_user_api_grpc_v1_user_proto = out.File
	file_app_service_pass_user_api_grpc_v1_user_proto_rawDesc = nil
	file_app_service_pass_user_api_grpc_v1_user_proto_goTypes = nil
	file_app_service_pass_user_api_grpc_v1_user_proto_depIdxs = nil
}
