// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.21.12
// source: app/game/cookie/api/grpc/v1/cookie.proto

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

type GetCookieReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid int64  `protobuf:"varint,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Key string `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
}

func (x *GetCookieReq) Reset() {
	*x = GetCookieReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_game_cookie_api_grpc_v1_cookie_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCookieReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCookieReq) ProtoMessage() {}

func (x *GetCookieReq) ProtoReflect() protoreflect.Message {
	mi := &file_app_game_cookie_api_grpc_v1_cookie_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCookieReq.ProtoReflect.Descriptor instead.
func (*GetCookieReq) Descriptor() ([]byte, []int) {
	return file_app_game_cookie_api_grpc_v1_cookie_proto_rawDescGZIP(), []int{0}
}

func (x *GetCookieReq) GetUid() int64 {
	if x != nil {
		return x.Uid
	}
	return 0
}

func (x *GetCookieReq) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

type GetCookieResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value       string `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	Exist       bool   `protobuf:"varint,2,opt,name=exist,proto3" json:"exist,omitempty"`
	LastUpdated int64  `protobuf:"varint,3,opt,name=last_updated,json=lastUpdated,proto3" json:"last_updated,omitempty"`
}

func (x *GetCookieResp) Reset() {
	*x = GetCookieResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_game_cookie_api_grpc_v1_cookie_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCookieResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCookieResp) ProtoMessage() {}

func (x *GetCookieResp) ProtoReflect() protoreflect.Message {
	mi := &file_app_game_cookie_api_grpc_v1_cookie_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCookieResp.ProtoReflect.Descriptor instead.
func (*GetCookieResp) Descriptor() ([]byte, []int) {
	return file_app_game_cookie_api_grpc_v1_cookie_proto_rawDescGZIP(), []int{1}
}

func (x *GetCookieResp) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

func (x *GetCookieResp) GetExist() bool {
	if x != nil {
		return x.Exist
	}
	return false
}

func (x *GetCookieResp) GetLastUpdated() int64 {
	if x != nil {
		return x.LastUpdated
	}
	return 0
}

type GetAllCookieReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid int64 `protobuf:"varint,1,opt,name=uid,proto3" json:"uid,omitempty"`
}

func (x *GetAllCookieReq) Reset() {
	*x = GetAllCookieReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_game_cookie_api_grpc_v1_cookie_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllCookieReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllCookieReq) ProtoMessage() {}

func (x *GetAllCookieReq) ProtoReflect() protoreflect.Message {
	mi := &file_app_game_cookie_api_grpc_v1_cookie_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllCookieReq.ProtoReflect.Descriptor instead.
func (*GetAllCookieReq) Descriptor() ([]byte, []int) {
	return file_app_game_cookie_api_grpc_v1_cookie_proto_rawDescGZIP(), []int{2}
}

func (x *GetAllCookieReq) GetUid() int64 {
	if x != nil {
		return x.Uid
	}
	return 0
}

type GetAllCookieResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Cookie map[string]string `protobuf:"bytes,1,rep,name=cookie,proto3" json:"cookie,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *GetAllCookieResp) Reset() {
	*x = GetAllCookieResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_game_cookie_api_grpc_v1_cookie_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllCookieResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllCookieResp) ProtoMessage() {}

func (x *GetAllCookieResp) ProtoReflect() protoreflect.Message {
	mi := &file_app_game_cookie_api_grpc_v1_cookie_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllCookieResp.ProtoReflect.Descriptor instead.
func (*GetAllCookieResp) Descriptor() ([]byte, []int) {
	return file_app_game_cookie_api_grpc_v1_cookie_proto_rawDescGZIP(), []int{3}
}

func (x *GetAllCookieResp) GetCookie() map[string]string {
	if x != nil {
		return x.Cookie
	}
	return nil
}

type SetCookieReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid       int64  `protobuf:"varint,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Key       string `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
	Value     string `protobuf:"bytes,3,opt,name=value,proto3" json:"value,omitempty"`
	Notify    bool   `protobuf:"varint,4,opt,name=notify,proto3" json:"notify,omitempty"`
	Timestamp int64  `protobuf:"varint,5,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (x *SetCookieReq) Reset() {
	*x = SetCookieReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_game_cookie_api_grpc_v1_cookie_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetCookieReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetCookieReq) ProtoMessage() {}

func (x *SetCookieReq) ProtoReflect() protoreflect.Message {
	mi := &file_app_game_cookie_api_grpc_v1_cookie_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetCookieReq.ProtoReflect.Descriptor instead.
func (*SetCookieReq) Descriptor() ([]byte, []int) {
	return file_app_game_cookie_api_grpc_v1_cookie_proto_rawDescGZIP(), []int{4}
}

func (x *SetCookieReq) GetUid() int64 {
	if x != nil {
		return x.Uid
	}
	return 0
}

func (x *SetCookieReq) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *SetCookieReq) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

func (x *SetCookieReq) GetNotify() bool {
	if x != nil {
		return x.Notify
	}
	return false
}

func (x *SetCookieReq) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

type SetCookieResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SetCookieResp) Reset() {
	*x = SetCookieResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_game_cookie_api_grpc_v1_cookie_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetCookieResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetCookieResp) ProtoMessage() {}

func (x *SetCookieResp) ProtoReflect() protoreflect.Message {
	mi := &file_app_game_cookie_api_grpc_v1_cookie_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetCookieResp.ProtoReflect.Descriptor instead.
func (*SetCookieResp) Descriptor() ([]byte, []int) {
	return file_app_game_cookie_api_grpc_v1_cookie_proto_rawDescGZIP(), []int{5}
}

var File_app_game_cookie_api_grpc_v1_cookie_proto protoreflect.FileDescriptor

var file_app_game_cookie_api_grpc_v1_cookie_proto_rawDesc = []byte{
	0x0a, 0x28, 0x61, 0x70, 0x70, 0x2f, 0x67, 0x61, 0x6d, 0x65, 0x2f, 0x63, 0x6f, 0x6f, 0x6b, 0x69,
	0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f,
	0x6f, 0x6b, 0x69, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12, 0x6e, 0x63, 0x73, 0x2e,
	0x67, 0x61, 0x6d, 0x65, 0x2e, 0x63, 0x6f, 0x6f, 0x6b, 0x69, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x1c,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x32, 0x0a, 0x0c,
	0x47, 0x65, 0x74, 0x43, 0x6f, 0x6f, 0x6b, 0x69, 0x65, 0x52, 0x65, 0x71, 0x12, 0x10, 0x0a, 0x03,
	0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x75, 0x69, 0x64, 0x12, 0x10,
	0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79,
	0x22, 0x5e, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6f, 0x6b, 0x69, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x78, 0x69, 0x73, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x65, 0x78, 0x69, 0x73, 0x74, 0x12, 0x21, 0x0a,
	0x0c, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x0b, 0x6c, 0x61, 0x73, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64,
	0x22, 0x23, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x43, 0x6f, 0x6f, 0x6b, 0x69, 0x65,
	0x52, 0x65, 0x71, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x03, 0x75, 0x69, 0x64, 0x22, 0x97, 0x01, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c,
	0x43, 0x6f, 0x6f, 0x6b, 0x69, 0x65, 0x52, 0x65, 0x73, 0x70, 0x12, 0x48, 0x0a, 0x06, 0x63, 0x6f,
	0x6f, 0x6b, 0x69, 0x65, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x30, 0x2e, 0x6e, 0x63, 0x73,
	0x2e, 0x67, 0x61, 0x6d, 0x65, 0x2e, 0x63, 0x6f, 0x6f, 0x6b, 0x69, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x43, 0x6f, 0x6f, 0x6b, 0x69, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x2e, 0x43, 0x6f, 0x6f, 0x6b, 0x69, 0x65, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x63, 0x6f,
	0x6f, 0x6b, 0x69, 0x65, 0x1a, 0x39, 0x0a, 0x0b, 0x43, 0x6f, 0x6f, 0x6b, 0x69, 0x65, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22,
	0x7e, 0x0a, 0x0c, 0x53, 0x65, 0x74, 0x43, 0x6f, 0x6f, 0x6b, 0x69, 0x65, 0x52, 0x65, 0x71, 0x12,
	0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x75, 0x69,
	0x64, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6e, 0x6f, 0x74,
	0x69, 0x66, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x6e, 0x6f, 0x74, 0x69, 0x66,
	0x79, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x22,
	0x0f, 0x0a, 0x0d, 0x53, 0x65, 0x74, 0x43, 0x6f, 0x6f, 0x6b, 0x69, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x32, 0xea, 0x02, 0x0a, 0x06, 0x43, 0x6f, 0x6f, 0x6b, 0x69, 0x65, 0x12, 0x76, 0x0a, 0x09, 0x47,
	0x65, 0x74, 0x43, 0x6f, 0x6f, 0x6b, 0x69, 0x65, 0x12, 0x20, 0x2e, 0x6e, 0x63, 0x73, 0x2e, 0x67,
	0x61, 0x6d, 0x65, 0x2e, 0x63, 0x6f, 0x6f, 0x6b, 0x69, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65,
	0x74, 0x43, 0x6f, 0x6f, 0x6b, 0x69, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x21, 0x2e, 0x6e, 0x63, 0x73,
	0x2e, 0x67, 0x61, 0x6d, 0x65, 0x2e, 0x63, 0x6f, 0x6f, 0x6b, 0x69, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x47, 0x65, 0x74, 0x43, 0x6f, 0x6f, 0x6b, 0x69, 0x65, 0x52, 0x65, 0x73, 0x70, 0x22, 0x24, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x1e, 0x12, 0x1c, 0x2f, 0x67, 0x61, 0x6d, 0x65, 0x2f, 0x63, 0x6f, 0x6f,
	0x6b, 0x69, 0x65, 0x2f, 0x67, 0x65, 0x74, 0x2f, 0x7b, 0x75, 0x69, 0x64, 0x7d, 0x2f, 0x7b, 0x6b,
	0x65, 0x79, 0x7d, 0x12, 0x79, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x43, 0x6f, 0x6f,
	0x6b, 0x69, 0x65, 0x12, 0x23, 0x2e, 0x6e, 0x63, 0x73, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x2e, 0x63,
	0x6f, 0x6f, 0x6b, 0x69, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x43,
	0x6f, 0x6f, 0x6b, 0x69, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x24, 0x2e, 0x6e, 0x63, 0x73, 0x2e, 0x67,
	0x61, 0x6d, 0x65, 0x2e, 0x63, 0x6f, 0x6f, 0x6b, 0x69, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65,
	0x74, 0x41, 0x6c, 0x6c, 0x43, 0x6f, 0x6f, 0x6b, 0x69, 0x65, 0x52, 0x65, 0x73, 0x70, 0x22, 0x1e,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x18, 0x12, 0x16, 0x2f, 0x67, 0x61, 0x6d, 0x65, 0x2f, 0x63, 0x6f,
	0x6f, 0x6b, 0x69, 0x65, 0x2f, 0x61, 0x6c, 0x6c, 0x2f, 0x7b, 0x75, 0x69, 0x64, 0x7d, 0x12, 0x6d,
	0x0a, 0x09, 0x53, 0x65, 0x74, 0x43, 0x6f, 0x6f, 0x6b, 0x69, 0x65, 0x12, 0x20, 0x2e, 0x6e, 0x63,
	0x73, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x2e, 0x63, 0x6f, 0x6f, 0x6b, 0x69, 0x65, 0x2e, 0x76, 0x31,
	0x2e, 0x53, 0x65, 0x74, 0x43, 0x6f, 0x6f, 0x6b, 0x69, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x21, 0x2e,
	0x6e, 0x63, 0x73, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x2e, 0x63, 0x6f, 0x6f, 0x6b, 0x69, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x53, 0x65, 0x74, 0x43, 0x6f, 0x6f, 0x6b, 0x69, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x22, 0x1b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x15, 0x22, 0x10, 0x2f, 0x67, 0x61, 0x6d, 0x65, 0x2f,
	0x63, 0x6f, 0x6f, 0x6b, 0x69, 0x65, 0x2f, 0x73, 0x65, 0x74, 0x3a, 0x01, 0x2a, 0x42, 0x25, 0x5a,
	0x23, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2f, 0x61, 0x70, 0x70, 0x2f, 0x67, 0x61, 0x6d,
	0x65, 0x2f, 0x63, 0x6f, 0x6f, 0x6b, 0x69, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x72, 0x70,
	0x63, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_app_game_cookie_api_grpc_v1_cookie_proto_rawDescOnce sync.Once
	file_app_game_cookie_api_grpc_v1_cookie_proto_rawDescData = file_app_game_cookie_api_grpc_v1_cookie_proto_rawDesc
)

func file_app_game_cookie_api_grpc_v1_cookie_proto_rawDescGZIP() []byte {
	file_app_game_cookie_api_grpc_v1_cookie_proto_rawDescOnce.Do(func() {
		file_app_game_cookie_api_grpc_v1_cookie_proto_rawDescData = protoimpl.X.CompressGZIP(file_app_game_cookie_api_grpc_v1_cookie_proto_rawDescData)
	})
	return file_app_game_cookie_api_grpc_v1_cookie_proto_rawDescData
}

var file_app_game_cookie_api_grpc_v1_cookie_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_app_game_cookie_api_grpc_v1_cookie_proto_goTypes = []interface{}{
	(*GetCookieReq)(nil),     // 0: ncs.game.cookie.v1.GetCookieReq
	(*GetCookieResp)(nil),    // 1: ncs.game.cookie.v1.GetCookieResp
	(*GetAllCookieReq)(nil),  // 2: ncs.game.cookie.v1.GetAllCookieReq
	(*GetAllCookieResp)(nil), // 3: ncs.game.cookie.v1.GetAllCookieResp
	(*SetCookieReq)(nil),     // 4: ncs.game.cookie.v1.SetCookieReq
	(*SetCookieResp)(nil),    // 5: ncs.game.cookie.v1.SetCookieResp
	nil,                      // 6: ncs.game.cookie.v1.GetAllCookieResp.CookieEntry
}
var file_app_game_cookie_api_grpc_v1_cookie_proto_depIdxs = []int32{
	6, // 0: ncs.game.cookie.v1.GetAllCookieResp.cookie:type_name -> ncs.game.cookie.v1.GetAllCookieResp.CookieEntry
	0, // 1: ncs.game.cookie.v1.Cookie.GetCookie:input_type -> ncs.game.cookie.v1.GetCookieReq
	2, // 2: ncs.game.cookie.v1.Cookie.GetAllCookie:input_type -> ncs.game.cookie.v1.GetAllCookieReq
	4, // 3: ncs.game.cookie.v1.Cookie.SetCookie:input_type -> ncs.game.cookie.v1.SetCookieReq
	1, // 4: ncs.game.cookie.v1.Cookie.GetCookie:output_type -> ncs.game.cookie.v1.GetCookieResp
	3, // 5: ncs.game.cookie.v1.Cookie.GetAllCookie:output_type -> ncs.game.cookie.v1.GetAllCookieResp
	5, // 6: ncs.game.cookie.v1.Cookie.SetCookie:output_type -> ncs.game.cookie.v1.SetCookieResp
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_app_game_cookie_api_grpc_v1_cookie_proto_init() }
func file_app_game_cookie_api_grpc_v1_cookie_proto_init() {
	if File_app_game_cookie_api_grpc_v1_cookie_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_app_game_cookie_api_grpc_v1_cookie_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCookieReq); i {
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
		file_app_game_cookie_api_grpc_v1_cookie_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCookieResp); i {
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
		file_app_game_cookie_api_grpc_v1_cookie_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllCookieReq); i {
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
		file_app_game_cookie_api_grpc_v1_cookie_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllCookieResp); i {
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
		file_app_game_cookie_api_grpc_v1_cookie_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetCookieReq); i {
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
		file_app_game_cookie_api_grpc_v1_cookie_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetCookieResp); i {
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
			RawDescriptor: file_app_game_cookie_api_grpc_v1_cookie_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_app_game_cookie_api_grpc_v1_cookie_proto_goTypes,
		DependencyIndexes: file_app_game_cookie_api_grpc_v1_cookie_proto_depIdxs,
		MessageInfos:      file_app_game_cookie_api_grpc_v1_cookie_proto_msgTypes,
	}.Build()
	File_app_game_cookie_api_grpc_v1_cookie_proto = out.File
	file_app_game_cookie_api_grpc_v1_cookie_proto_rawDesc = nil
	file_app_game_cookie_api_grpc_v1_cookie_proto_goTypes = nil
	file_app_game_cookie_api_grpc_v1_cookie_proto_depIdxs = nil
}
