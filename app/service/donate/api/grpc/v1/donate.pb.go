// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.21.12
// source: app/service/donate/api/grpc/v1/donate.proto

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

type PAY_TYPE int32

const (
	PAY_TYPE_NULL   PAY_TYPE = 0
	PAY_TYPE_ALIPAY PAY_TYPE = 1
	PAY_TYPE_WEPAY  PAY_TYPE = 2
)

// Enum value maps for PAY_TYPE.
var (
	PAY_TYPE_name = map[int32]string{
		0: "NULL",
		1: "ALIPAY",
		2: "WEPAY",
	}
	PAY_TYPE_value = map[string]int32{
		"NULL":   0,
		"ALIPAY": 1,
		"WEPAY":  2,
	}
)

func (x PAY_TYPE) Enum() *PAY_TYPE {
	p := new(PAY_TYPE)
	*p = x
	return p
}

func (x PAY_TYPE) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PAY_TYPE) Descriptor() protoreflect.EnumDescriptor {
	return file_app_service_donate_api_grpc_v1_donate_proto_enumTypes[0].Descriptor()
}

func (PAY_TYPE) Type() protoreflect.EnumType {
	return &file_app_service_donate_api_grpc_v1_donate_proto_enumTypes[0]
}

func (x PAY_TYPE) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PAY_TYPE.Descriptor instead.
func (PAY_TYPE) EnumDescriptor() ([]byte, []int) {
	return file_app_service_donate_api_grpc_v1_donate_proto_rawDescGZIP(), []int{0}
}

type CreateDonateReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SteamId int64    `protobuf:"varint,1,opt,name=steam_id,json=steamId,proto3" json:"steam_id,omitempty"`
	Amount  int32    `protobuf:"varint,2,opt,name=amount,proto3" json:"amount,omitempty"`
	Payment PAY_TYPE `protobuf:"varint,3,opt,name=payment,proto3,enum=ncs.service.donate.v1.PAY_TYPE" json:"payment,omitempty"`
}

func (x *CreateDonateReq) Reset() {
	*x = CreateDonateReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_service_donate_api_grpc_v1_donate_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateDonateReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateDonateReq) ProtoMessage() {}

func (x *CreateDonateReq) ProtoReflect() protoreflect.Message {
	mi := &file_app_service_donate_api_grpc_v1_donate_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateDonateReq.ProtoReflect.Descriptor instead.
func (*CreateDonateReq) Descriptor() ([]byte, []int) {
	return file_app_service_donate_api_grpc_v1_donate_proto_rawDescGZIP(), []int{0}
}

func (x *CreateDonateReq) GetSteamId() int64 {
	if x != nil {
		return x.SteamId
	}
	return 0
}

func (x *CreateDonateReq) GetAmount() int32 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *CreateDonateReq) GetPayment() PAY_TYPE {
	if x != nil {
		return x.Payment
	}
	return PAY_TYPE_NULL
}

type CreateDonateResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OutTradeNo string   `protobuf:"bytes,1,opt,name=out_trade_no,json=outTradeNo,proto3" json:"out_trade_no,omitempty"`
	QrCode     string   `protobuf:"bytes,2,opt,name=qr_code,json=qrCode,proto3" json:"qr_code,omitempty"`
	Payment    PAY_TYPE `protobuf:"varint,3,opt,name=payment,proto3,enum=ncs.service.donate.v1.PAY_TYPE" json:"payment,omitempty"`
}

func (x *CreateDonateResp) Reset() {
	*x = CreateDonateResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_service_donate_api_grpc_v1_donate_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateDonateResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateDonateResp) ProtoMessage() {}

func (x *CreateDonateResp) ProtoReflect() protoreflect.Message {
	mi := &file_app_service_donate_api_grpc_v1_donate_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateDonateResp.ProtoReflect.Descriptor instead.
func (*CreateDonateResp) Descriptor() ([]byte, []int) {
	return file_app_service_donate_api_grpc_v1_donate_proto_rawDescGZIP(), []int{1}
}

func (x *CreateDonateResp) GetOutTradeNo() string {
	if x != nil {
		return x.OutTradeNo
	}
	return ""
}

func (x *CreateDonateResp) GetQrCode() string {
	if x != nil {
		return x.QrCode
	}
	return ""
}

func (x *CreateDonateResp) GetPayment() PAY_TYPE {
	if x != nil {
		return x.Payment
	}
	return PAY_TYPE_NULL
}

type QueryDonateReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OutTradeNo string `protobuf:"bytes,1,opt,name=out_trade_no,json=outTradeNo,proto3" json:"out_trade_no,omitempty"`
}

func (x *QueryDonateReq) Reset() {
	*x = QueryDonateReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_service_donate_api_grpc_v1_donate_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryDonateReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryDonateReq) ProtoMessage() {}

func (x *QueryDonateReq) ProtoReflect() protoreflect.Message {
	mi := &file_app_service_donate_api_grpc_v1_donate_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryDonateReq.ProtoReflect.Descriptor instead.
func (*QueryDonateReq) Descriptor() ([]byte, []int) {
	return file_app_service_donate_api_grpc_v1_donate_proto_rawDescGZIP(), []int{2}
}

func (x *QueryDonateReq) GetOutTradeNo() string {
	if x != nil {
		return x.OutTradeNo
	}
	return ""
}

type QueryDonateResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid      int64 `protobuf:"varint,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Amount   int32 `protobuf:"varint,2,opt,name=amount,proto3" json:"amount,omitempty"`
	Status   int32 `protobuf:"varint,3,opt,name=status,proto3" json:"status,omitempty"`
	CreateAt int64 `protobuf:"varint,4,opt,name=create_at,json=createAt,proto3" json:"create_at,omitempty"`
}

func (x *QueryDonateResp) Reset() {
	*x = QueryDonateResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_service_donate_api_grpc_v1_donate_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryDonateResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryDonateResp) ProtoMessage() {}

func (x *QueryDonateResp) ProtoReflect() protoreflect.Message {
	mi := &file_app_service_donate_api_grpc_v1_donate_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryDonateResp.ProtoReflect.Descriptor instead.
func (*QueryDonateResp) Descriptor() ([]byte, []int) {
	return file_app_service_donate_api_grpc_v1_donate_proto_rawDescGZIP(), []int{3}
}

func (x *QueryDonateResp) GetUid() int64 {
	if x != nil {
		return x.Uid
	}
	return 0
}

func (x *QueryDonateResp) GetAmount() int32 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *QueryDonateResp) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *QueryDonateResp) GetCreateAt() int64 {
	if x != nil {
		return x.CreateAt
	}
	return 0
}

type GetDonateListReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StartTime int64 `protobuf:"varint,1,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	EndTime   int64 `protobuf:"varint,2,opt,name=end_time,json=endTime,proto3" json:"end_time,omitempty"`
	SteamId   int64 `protobuf:"varint,3,opt,name=steam_id,json=steamId,proto3" json:"steam_id,omitempty"`
}

func (x *GetDonateListReq) Reset() {
	*x = GetDonateListReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_service_donate_api_grpc_v1_donate_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDonateListReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDonateListReq) ProtoMessage() {}

func (x *GetDonateListReq) ProtoReflect() protoreflect.Message {
	mi := &file_app_service_donate_api_grpc_v1_donate_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDonateListReq.ProtoReflect.Descriptor instead.
func (*GetDonateListReq) Descriptor() ([]byte, []int) {
	return file_app_service_donate_api_grpc_v1_donate_proto_rawDescGZIP(), []int{4}
}

func (x *GetDonateListReq) GetStartTime() int64 {
	if x != nil {
		return x.StartTime
	}
	return 0
}

func (x *GetDonateListReq) GetEndTime() int64 {
	if x != nil {
		return x.EndTime
	}
	return 0
}

func (x *GetDonateListReq) GetSteamId() int64 {
	if x != nil {
		return x.SteamId
	}
	return 0
}

type GetDonateListResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	List []*QueryDonateResp `protobuf:"bytes,1,rep,name=list,proto3" json:"list,omitempty"`
}

func (x *GetDonateListResp) Reset() {
	*x = GetDonateListResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_service_donate_api_grpc_v1_donate_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDonateListResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDonateListResp) ProtoMessage() {}

func (x *GetDonateListResp) ProtoReflect() protoreflect.Message {
	mi := &file_app_service_donate_api_grpc_v1_donate_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDonateListResp.ProtoReflect.Descriptor instead.
func (*GetDonateListResp) Descriptor() ([]byte, []int) {
	return file_app_service_donate_api_grpc_v1_donate_proto_rawDescGZIP(), []int{5}
}

func (x *GetDonateListResp) GetList() []*QueryDonateResp {
	if x != nil {
		return x.List
	}
	return nil
}

type AddDonateReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SteamId int64 `protobuf:"varint,1,opt,name=steam_id,json=steamId,proto3" json:"steam_id,omitempty"`
	Amount  int32 `protobuf:"varint,2,opt,name=amount,proto3" json:"amount,omitempty"`
	// MQ
	Username string `protobuf:"bytes,3,opt,name=username,proto3" json:"username,omitempty"`
	// MQ
	Uid int64 `protobuf:"varint,4,opt,name=uid,proto3" json:"uid,omitempty"`
}

func (x *AddDonateReq) Reset() {
	*x = AddDonateReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_service_donate_api_grpc_v1_donate_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddDonateReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddDonateReq) ProtoMessage() {}

func (x *AddDonateReq) ProtoReflect() protoreflect.Message {
	mi := &file_app_service_donate_api_grpc_v1_donate_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddDonateReq.ProtoReflect.Descriptor instead.
func (*AddDonateReq) Descriptor() ([]byte, []int) {
	return file_app_service_donate_api_grpc_v1_donate_proto_rawDescGZIP(), []int{6}
}

func (x *AddDonateReq) GetSteamId() int64 {
	if x != nil {
		return x.SteamId
	}
	return 0
}

func (x *AddDonateReq) GetAmount() int32 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *AddDonateReq) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *AddDonateReq) GetUid() int64 {
	if x != nil {
		return x.Uid
	}
	return 0
}

type AddDonateResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *AddDonateResp) Reset() {
	*x = AddDonateResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_service_donate_api_grpc_v1_donate_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddDonateResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddDonateResp) ProtoMessage() {}

func (x *AddDonateResp) ProtoReflect() protoreflect.Message {
	mi := &file_app_service_donate_api_grpc_v1_donate_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddDonateResp.ProtoReflect.Descriptor instead.
func (*AddDonateResp) Descriptor() ([]byte, []int) {
	return file_app_service_donate_api_grpc_v1_donate_proto_rawDescGZIP(), []int{7}
}

var File_app_service_donate_api_grpc_v1_donate_proto protoreflect.FileDescriptor

var file_app_service_donate_api_grpc_v1_donate_proto_rawDesc = []byte{
	0x0a, 0x2b, 0x61, 0x70, 0x70, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x64, 0x6f,
	0x6e, 0x61, 0x74, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x76, 0x31,
	0x2f, 0x64, 0x6f, 0x6e, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x15, 0x6e,
	0x63, 0x73, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x64, 0x6f, 0x6e, 0x61, 0x74,
	0x65, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x7f, 0x0a, 0x0f, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44, 0x6f, 0x6e, 0x61,
	0x74, 0x65, 0x52, 0x65, 0x71, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x74, 0x65, 0x61, 0x6d, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x73, 0x74, 0x65, 0x61, 0x6d, 0x49, 0x64,
	0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x39, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6d,
	0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1f, 0x2e, 0x6e, 0x63, 0x73, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x64, 0x6f, 0x6e, 0x61, 0x74, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x50, 0x41, 0x59, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x52, 0x07, 0x70, 0x61, 0x79, 0x6d,
	0x65, 0x6e, 0x74, 0x22, 0x88, 0x01, 0x0a, 0x10, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44, 0x6f,
	0x6e, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x12, 0x20, 0x0a, 0x0c, 0x6f, 0x75, 0x74, 0x5f,
	0x74, 0x72, 0x61, 0x64, 0x65, 0x5f, 0x6e, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x6f, 0x75, 0x74, 0x54, 0x72, 0x61, 0x64, 0x65, 0x4e, 0x6f, 0x12, 0x17, 0x0a, 0x07, 0x71, 0x72,
	0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x71, 0x72, 0x43,
	0x6f, 0x64, 0x65, 0x12, 0x39, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x1f, 0x2e, 0x6e, 0x63, 0x73, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x64, 0x6f, 0x6e, 0x61, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x41, 0x59,
	0x5f, 0x54, 0x59, 0x50, 0x45, 0x52, 0x07, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x22, 0x32,
	0x0a, 0x0e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x44, 0x6f, 0x6e, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71,
	0x12, 0x20, 0x0a, 0x0c, 0x6f, 0x75, 0x74, 0x5f, 0x74, 0x72, 0x61, 0x64, 0x65, 0x5f, 0x6e, 0x6f,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6f, 0x75, 0x74, 0x54, 0x72, 0x61, 0x64, 0x65,
	0x4e, 0x6f, 0x22, 0x70, 0x0a, 0x0f, 0x51, 0x75, 0x65, 0x72, 0x79, 0x44, 0x6f, 0x6e, 0x61, 0x74,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x03, 0x75, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12,
	0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x5f, 0x61, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x41, 0x74, 0x22, 0x67, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x44, 0x6f, 0x6e, 0x61, 0x74,
	0x65, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72,
	0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x73, 0x74,
	0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x65, 0x6e, 0x64, 0x5f, 0x74,
	0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x65, 0x6e, 0x64, 0x54, 0x69,
	0x6d, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x74, 0x65, 0x61, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x73, 0x74, 0x65, 0x61, 0x6d, 0x49, 0x64, 0x22, 0x4f, 0x0a,
	0x11, 0x47, 0x65, 0x74, 0x44, 0x6f, 0x6e, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x12, 0x3a, 0x0a, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x26, 0x2e, 0x6e, 0x63, 0x73, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x64,
	0x6f, 0x6e, 0x61, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x44, 0x6f,
	0x6e, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x52, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x22, 0x6f,
	0x0a, 0x0c, 0x41, 0x64, 0x64, 0x44, 0x6f, 0x6e, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x12, 0x19,
	0x0a, 0x08, 0x73, 0x74, 0x65, 0x61, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x07, 0x73, 0x74, 0x65, 0x61, 0x6d, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f,
	0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e,
	0x74, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a,
	0x03, 0x75, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x75, 0x69, 0x64, 0x22,
	0x0f, 0x0a, 0x0d, 0x41, 0x64, 0x64, 0x44, 0x6f, 0x6e, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x2a, 0x2b, 0x0a, 0x08, 0x50, 0x41, 0x59, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x12, 0x08, 0x0a, 0x04,
	0x4e, 0x55, 0x4c, 0x4c, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x41, 0x4c, 0x49, 0x50, 0x41, 0x59,
	0x10, 0x01, 0x12, 0x09, 0x0a, 0x05, 0x57, 0x45, 0x50, 0x41, 0x59, 0x10, 0x02, 0x32, 0xa1, 0x04,
	0x0a, 0x06, 0x44, 0x6f, 0x6e, 0x61, 0x74, 0x65, 0x12, 0x7a, 0x0a, 0x0c, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x44, 0x6f, 0x6e, 0x61, 0x74, 0x65, 0x12, 0x26, 0x2e, 0x6e, 0x63, 0x73, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x64, 0x6f, 0x6e, 0x61, 0x74, 0x65, 0x2e, 0x76, 0x31,
	0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44, 0x6f, 0x6e, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71,
	0x1a, 0x27, 0x2e, 0x6e, 0x63, 0x73, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x64,
	0x6f, 0x6e, 0x61, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44,
	0x6f, 0x6e, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x22, 0x19, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x13, 0x22, 0x0e, 0x2f, 0x64, 0x6f, 0x6e, 0x61, 0x74, 0x65, 0x2f, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x3a, 0x01, 0x2a, 0x12, 0x82, 0x01, 0x0a, 0x0b, 0x51, 0x75, 0x65, 0x72, 0x79, 0x44, 0x6f,
	0x6e, 0x61, 0x74, 0x65, 0x12, 0x25, 0x2e, 0x6e, 0x63, 0x73, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x64, 0x6f, 0x6e, 0x61, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x51, 0x75, 0x65,
	0x72, 0x79, 0x44, 0x6f, 0x6e, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x26, 0x2e, 0x6e, 0x63,
	0x73, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x64, 0x6f, 0x6e, 0x61, 0x74, 0x65,
	0x2e, 0x76, 0x31, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x44, 0x6f, 0x6e, 0x61, 0x74, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x22, 0x24, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1e, 0x12, 0x1c, 0x2f, 0x64, 0x6f,
	0x6e, 0x61, 0x74, 0x65, 0x2f, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2f, 0x7b, 0x6f, 0x75, 0x74, 0x5f,
	0x74, 0x72, 0x61, 0x64, 0x65, 0x5f, 0x6e, 0x6f, 0x7d, 0x12, 0xbd, 0x01, 0x0a, 0x0d, 0x47, 0x65,
	0x74, 0x44, 0x6f, 0x6e, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x27, 0x2e, 0x6e, 0x63,
	0x73, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x64, 0x6f, 0x6e, 0x61, 0x74, 0x65,
	0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x44, 0x6f, 0x6e, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x73,
	0x74, 0x52, 0x65, 0x71, 0x1a, 0x28, 0x2e, 0x6e, 0x63, 0x73, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x64, 0x6f, 0x6e, 0x61, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74,
	0x44, 0x6f, 0x6e, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x22, 0x59,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x53, 0x12, 0x0c, 0x2f, 0x64, 0x6f, 0x6e, 0x61, 0x74, 0x65, 0x2f,
	0x6c, 0x69, 0x73, 0x74, 0x5a, 0x1b, 0x12, 0x19, 0x2f, 0x64, 0x6f, 0x6e, 0x61, 0x74, 0x65, 0x2f,
	0x6c, 0x69, 0x73, 0x74, 0x2f, 0x7b, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65,
	0x7d, 0x5a, 0x26, 0x12, 0x24, 0x2f, 0x64, 0x6f, 0x6e, 0x61, 0x74, 0x65, 0x2f, 0x6c, 0x69, 0x73,
	0x74, 0x2f, 0x7b, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x7d, 0x2f, 0x7b,
	0x65, 0x6e, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x7d, 0x12, 0x56, 0x0a, 0x09, 0x41, 0x64, 0x64,
	0x44, 0x6f, 0x6e, 0x61, 0x74, 0x65, 0x12, 0x23, 0x2e, 0x6e, 0x63, 0x73, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x64, 0x6f, 0x6e, 0x61, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x41,
	0x64, 0x64, 0x44, 0x6f, 0x6e, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x24, 0x2e, 0x6e, 0x63,
	0x73, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x64, 0x6f, 0x6e, 0x61, 0x74, 0x65,
	0x2e, 0x76, 0x31, 0x2e, 0x41, 0x64, 0x64, 0x44, 0x6f, 0x6e, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x42, 0x28, 0x5a, 0x26, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2f, 0x61, 0x70, 0x70,
	0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x64, 0x6f, 0x6e, 0x61, 0x74, 0x65, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_app_service_donate_api_grpc_v1_donate_proto_rawDescOnce sync.Once
	file_app_service_donate_api_grpc_v1_donate_proto_rawDescData = file_app_service_donate_api_grpc_v1_donate_proto_rawDesc
)

func file_app_service_donate_api_grpc_v1_donate_proto_rawDescGZIP() []byte {
	file_app_service_donate_api_grpc_v1_donate_proto_rawDescOnce.Do(func() {
		file_app_service_donate_api_grpc_v1_donate_proto_rawDescData = protoimpl.X.CompressGZIP(file_app_service_donate_api_grpc_v1_donate_proto_rawDescData)
	})
	return file_app_service_donate_api_grpc_v1_donate_proto_rawDescData
}

var file_app_service_donate_api_grpc_v1_donate_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_app_service_donate_api_grpc_v1_donate_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_app_service_donate_api_grpc_v1_donate_proto_goTypes = []interface{}{
	(PAY_TYPE)(0),             // 0: ncs.service.donate.v1.PAY_TYPE
	(*CreateDonateReq)(nil),   // 1: ncs.service.donate.v1.CreateDonateReq
	(*CreateDonateResp)(nil),  // 2: ncs.service.donate.v1.CreateDonateResp
	(*QueryDonateReq)(nil),    // 3: ncs.service.donate.v1.QueryDonateReq
	(*QueryDonateResp)(nil),   // 4: ncs.service.donate.v1.QueryDonateResp
	(*GetDonateListReq)(nil),  // 5: ncs.service.donate.v1.GetDonateListReq
	(*GetDonateListResp)(nil), // 6: ncs.service.donate.v1.GetDonateListResp
	(*AddDonateReq)(nil),      // 7: ncs.service.donate.v1.AddDonateReq
	(*AddDonateResp)(nil),     // 8: ncs.service.donate.v1.AddDonateResp
}
var file_app_service_donate_api_grpc_v1_donate_proto_depIdxs = []int32{
	0, // 0: ncs.service.donate.v1.CreateDonateReq.payment:type_name -> ncs.service.donate.v1.PAY_TYPE
	0, // 1: ncs.service.donate.v1.CreateDonateResp.payment:type_name -> ncs.service.donate.v1.PAY_TYPE
	4, // 2: ncs.service.donate.v1.GetDonateListResp.list:type_name -> ncs.service.donate.v1.QueryDonateResp
	1, // 3: ncs.service.donate.v1.Donate.CreateDonate:input_type -> ncs.service.donate.v1.CreateDonateReq
	3, // 4: ncs.service.donate.v1.Donate.QueryDonate:input_type -> ncs.service.donate.v1.QueryDonateReq
	5, // 5: ncs.service.donate.v1.Donate.GetDonateList:input_type -> ncs.service.donate.v1.GetDonateListReq
	7, // 6: ncs.service.donate.v1.Donate.AddDonate:input_type -> ncs.service.donate.v1.AddDonateReq
	2, // 7: ncs.service.donate.v1.Donate.CreateDonate:output_type -> ncs.service.donate.v1.CreateDonateResp
	4, // 8: ncs.service.donate.v1.Donate.QueryDonate:output_type -> ncs.service.donate.v1.QueryDonateResp
	6, // 9: ncs.service.donate.v1.Donate.GetDonateList:output_type -> ncs.service.donate.v1.GetDonateListResp
	8, // 10: ncs.service.donate.v1.Donate.AddDonate:output_type -> ncs.service.donate.v1.AddDonateResp
	7, // [7:11] is the sub-list for method output_type
	3, // [3:7] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_app_service_donate_api_grpc_v1_donate_proto_init() }
func file_app_service_donate_api_grpc_v1_donate_proto_init() {
	if File_app_service_donate_api_grpc_v1_donate_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_app_service_donate_api_grpc_v1_donate_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateDonateReq); i {
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
		file_app_service_donate_api_grpc_v1_donate_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateDonateResp); i {
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
		file_app_service_donate_api_grpc_v1_donate_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryDonateReq); i {
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
		file_app_service_donate_api_grpc_v1_donate_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryDonateResp); i {
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
		file_app_service_donate_api_grpc_v1_donate_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetDonateListReq); i {
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
		file_app_service_donate_api_grpc_v1_donate_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetDonateListResp); i {
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
		file_app_service_donate_api_grpc_v1_donate_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddDonateReq); i {
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
		file_app_service_donate_api_grpc_v1_donate_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddDonateResp); i {
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
			RawDescriptor: file_app_service_donate_api_grpc_v1_donate_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_app_service_donate_api_grpc_v1_donate_proto_goTypes,
		DependencyIndexes: file_app_service_donate_api_grpc_v1_donate_proto_depIdxs,
		EnumInfos:         file_app_service_donate_api_grpc_v1_donate_proto_enumTypes,
		MessageInfos:      file_app_service_donate_api_grpc_v1_donate_proto_msgTypes,
	}.Build()
	File_app_service_donate_api_grpc_v1_donate_proto = out.File
	file_app_service_donate_api_grpc_v1_donate_proto_rawDesc = nil
	file_app_service_donate_api_grpc_v1_donate_proto_goTypes = nil
	file_app_service_donate_api_grpc_v1_donate_proto_depIdxs = nil
}
