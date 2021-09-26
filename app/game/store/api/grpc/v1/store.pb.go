// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.12.4
// source: app/game/store/api/grpc/v1/store.proto

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

type Item struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ItemId      int32   `protobuf:"varint,1,opt,name=item_id,json=itemId,proto3" json:"item_id,omitempty"`
	Name        string  `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description string  `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Type        int32   `protobuf:"varint,4,opt,name=type,proto3" json:"type,omitempty"`
	Price       int32   `protobuf:"varint,6,opt,name=price,proto3" json:"price,omitempty"`
	Discount    float32 `protobuf:"fixed32,7,opt,name=discount,proto3" json:"discount,omitempty"`
	Available   bool    `protobuf:"varint,8,opt,name=available,proto3" json:"available,omitempty"`
	AlreadyHave bool    `protobuf:"varint,9,opt,name=already_have,json=alreadyHave,proto3" json:"already_have,omitempty"`
}

func (x *Item) Reset() {
	*x = Item{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_game_store_api_grpc_v1_store_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Item) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Item) ProtoMessage() {}

func (x *Item) ProtoReflect() protoreflect.Message {
	mi := &file_app_game_store_api_grpc_v1_store_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Item.ProtoReflect.Descriptor instead.
func (*Item) Descriptor() ([]byte, []int) {
	return file_app_game_store_api_grpc_v1_store_proto_rawDescGZIP(), []int{0}
}

func (x *Item) GetItemId() int32 {
	if x != nil {
		return x.ItemId
	}
	return 0
}

func (x *Item) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Item) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Item) GetType() int32 {
	if x != nil {
		return x.Type
	}
	return 0
}

func (x *Item) GetPrice() int32 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *Item) GetDiscount() float32 {
	if x != nil {
		return x.Discount
	}
	return 0
}

func (x *Item) GetAvailable() bool {
	if x != nil {
		return x.Available
	}
	return false
}

func (x *Item) GetAlreadyHave() bool {
	if x != nil {
		return x.AlreadyHave
	}
	return false
}

type HotSaleListReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *HotSaleListReq) Reset() {
	*x = HotSaleListReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_game_store_api_grpc_v1_store_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HotSaleListReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HotSaleListReq) ProtoMessage() {}

func (x *HotSaleListReq) ProtoReflect() protoreflect.Message {
	mi := &file_app_game_store_api_grpc_v1_store_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HotSaleListReq.ProtoReflect.Descriptor instead.
func (*HotSaleListReq) Descriptor() ([]byte, []int) {
	return file_app_game_store_api_grpc_v1_store_proto_rawDescGZIP(), []int{1}
}

type HotSaleListResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ItemsId []int32 `protobuf:"varint,1,rep,packed,name=items_id,json=itemsId,proto3" json:"items_id,omitempty"`
}

func (x *HotSaleListResp) Reset() {
	*x = HotSaleListResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_game_store_api_grpc_v1_store_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HotSaleListResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HotSaleListResp) ProtoMessage() {}

func (x *HotSaleListResp) ProtoReflect() protoreflect.Message {
	mi := &file_app_game_store_api_grpc_v1_store_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HotSaleListResp.ProtoReflect.Descriptor instead.
func (*HotSaleListResp) Descriptor() ([]byte, []int) {
	return file_app_game_store_api_grpc_v1_store_proto_rawDescGZIP(), []int{2}
}

func (x *HotSaleListResp) GetItemsId() []int32 {
	if x != nil {
		return x.ItemsId
	}
	return nil
}

type BuyItemReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid    int64 `protobuf:"varint,1,opt,name=uid,proto3" json:"uid,omitempty"`
	ItemId int32 `protobuf:"varint,2,opt,name=item_id,json=itemId,proto3" json:"item_id,omitempty"`
	Price  int32 `protobuf:"varint,3,opt,name=price,proto3" json:"price,omitempty"`
}

func (x *BuyItemReq) Reset() {
	*x = BuyItemReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_game_store_api_grpc_v1_store_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BuyItemReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BuyItemReq) ProtoMessage() {}

func (x *BuyItemReq) ProtoReflect() protoreflect.Message {
	mi := &file_app_game_store_api_grpc_v1_store_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BuyItemReq.ProtoReflect.Descriptor instead.
func (*BuyItemReq) Descriptor() ([]byte, []int) {
	return file_app_game_store_api_grpc_v1_store_proto_rawDescGZIP(), []int{3}
}

func (x *BuyItemReq) GetUid() int64 {
	if x != nil {
		return x.Uid
	}
	return 0
}

func (x *BuyItemReq) GetItemId() int32 {
	if x != nil {
		return x.ItemId
	}
	return 0
}

func (x *BuyItemReq) GetPrice() int32 {
	if x != nil {
		return x.Price
	}
	return 0
}

type BuyItemResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *BuyItemResp) Reset() {
	*x = BuyItemResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_game_store_api_grpc_v1_store_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BuyItemResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BuyItemResp) ProtoMessage() {}

func (x *BuyItemResp) ProtoReflect() protoreflect.Message {
	mi := &file_app_game_store_api_grpc_v1_store_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BuyItemResp.ProtoReflect.Descriptor instead.
func (*BuyItemResp) Descriptor() ([]byte, []int) {
	return file_app_game_store_api_grpc_v1_store_proto_rawDescGZIP(), []int{4}
}

type SaleListReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid int64 `protobuf:"varint,1,opt,name=uid,proto3" json:"uid,omitempty"`
}

func (x *SaleListReq) Reset() {
	*x = SaleListReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_game_store_api_grpc_v1_store_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SaleListReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SaleListReq) ProtoMessage() {}

func (x *SaleListReq) ProtoReflect() protoreflect.Message {
	mi := &file_app_game_store_api_grpc_v1_store_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SaleListReq.ProtoReflect.Descriptor instead.
func (*SaleListReq) Descriptor() ([]byte, []int) {
	return file_app_game_store_api_grpc_v1_store_proto_rawDescGZIP(), []int{5}
}

func (x *SaleListReq) GetUid() int64 {
	if x != nil {
		return x.Uid
	}
	return 0
}

type SaleListResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items    []*Item `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
	Money    int32   `protobuf:"varint,2,opt,name=money,proto3" json:"money,omitempty"`
	PassType int32   `protobuf:"varint,3,opt,name=pass_type,json=passType,proto3" json:"pass_type,omitempty"`
}

func (x *SaleListResp) Reset() {
	*x = SaleListResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_game_store_api_grpc_v1_store_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SaleListResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SaleListResp) ProtoMessage() {}

func (x *SaleListResp) ProtoReflect() protoreflect.Message {
	mi := &file_app_game_store_api_grpc_v1_store_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SaleListResp.ProtoReflect.Descriptor instead.
func (*SaleListResp) Descriptor() ([]byte, []int) {
	return file_app_game_store_api_grpc_v1_store_proto_rawDescGZIP(), []int{6}
}

func (x *SaleListResp) GetItems() []*Item {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *SaleListResp) GetMoney() int32 {
	if x != nil {
		return x.Money
	}
	return 0
}

func (x *SaleListResp) GetPassType() int32 {
	if x != nil {
		return x.PassType
	}
	return 0
}

type BuyPassReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid  int64 `protobuf:"varint,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Type int32 `protobuf:"varint,2,opt,name=type,proto3" json:"type,omitempty"`
}

func (x *BuyPassReq) Reset() {
	*x = BuyPassReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_game_store_api_grpc_v1_store_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BuyPassReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BuyPassReq) ProtoMessage() {}

func (x *BuyPassReq) ProtoReflect() protoreflect.Message {
	mi := &file_app_game_store_api_grpc_v1_store_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BuyPassReq.ProtoReflect.Descriptor instead.
func (*BuyPassReq) Descriptor() ([]byte, []int) {
	return file_app_game_store_api_grpc_v1_store_proto_rawDescGZIP(), []int{7}
}

func (x *BuyPassReq) GetUid() int64 {
	if x != nil {
		return x.Uid
	}
	return 0
}

func (x *BuyPassReq) GetType() int32 {
	if x != nil {
		return x.Type
	}
	return 0
}

type BuyPassResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *BuyPassResp) Reset() {
	*x = BuyPassResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_game_store_api_grpc_v1_store_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BuyPassResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BuyPassResp) ProtoMessage() {}

func (x *BuyPassResp) ProtoReflect() protoreflect.Message {
	mi := &file_app_game_store_api_grpc_v1_store_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BuyPassResp.ProtoReflect.Descriptor instead.
func (*BuyPassResp) Descriptor() ([]byte, []int) {
	return file_app_game_store_api_grpc_v1_store_proto_rawDescGZIP(), []int{8}
}

type BuyVIPReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid   int64 `protobuf:"varint,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Month int32 `protobuf:"varint,2,opt,name=month,proto3" json:"month,omitempty"`
}

func (x *BuyVIPReq) Reset() {
	*x = BuyVIPReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_game_store_api_grpc_v1_store_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BuyVIPReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BuyVIPReq) ProtoMessage() {}

func (x *BuyVIPReq) ProtoReflect() protoreflect.Message {
	mi := &file_app_game_store_api_grpc_v1_store_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BuyVIPReq.ProtoReflect.Descriptor instead.
func (*BuyVIPReq) Descriptor() ([]byte, []int) {
	return file_app_game_store_api_grpc_v1_store_proto_rawDescGZIP(), []int{9}
}

func (x *BuyVIPReq) GetUid() int64 {
	if x != nil {
		return x.Uid
	}
	return 0
}

func (x *BuyVIPReq) GetMonth() int32 {
	if x != nil {
		return x.Month
	}
	return 0
}

type BuyVIPResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *BuyVIPResp) Reset() {
	*x = BuyVIPResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_game_store_api_grpc_v1_store_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BuyVIPResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BuyVIPResp) ProtoMessage() {}

func (x *BuyVIPResp) ProtoReflect() protoreflect.Message {
	mi := &file_app_game_store_api_grpc_v1_store_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BuyVIPResp.ProtoReflect.Descriptor instead.
func (*BuyVIPResp) Descriptor() ([]byte, []int) {
	return file_app_game_store_api_grpc_v1_store_proto_rawDescGZIP(), []int{10}
}

var File_app_game_store_api_grpc_v1_store_proto protoreflect.FileDescriptor

var file_app_game_store_api_grpc_v1_store_proto_rawDesc = []byte{
	0x0a, 0x26, 0x61, 0x70, 0x70, 0x2f, 0x67, 0x61, 0x6d, 0x65, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x74, 0x6f,
	0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x11, 0x6e, 0x63, 0x73, 0x2e, 0x67, 0x61,
	0x6d, 0x65, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xdc, 0x01, 0x0a, 0x04, 0x49, 0x74,
	0x65, 0x6d, 0x12, 0x17, 0x0a, 0x07, 0x69, 0x74, 0x65, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x06, 0x69, 0x74, 0x65, 0x6d, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x64,
	0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x02, 0x52, 0x08, 0x64,
	0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x61, 0x76, 0x61, 0x69, 0x6c,
	0x61, 0x62, 0x6c, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x61, 0x76, 0x61, 0x69,
	0x6c, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x61, 0x6c, 0x72, 0x65, 0x61, 0x64, 0x79,
	0x5f, 0x68, 0x61, 0x76, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x61, 0x6c, 0x72,
	0x65, 0x61, 0x64, 0x79, 0x48, 0x61, 0x76, 0x65, 0x22, 0x10, 0x0a, 0x0e, 0x48, 0x6f, 0x74, 0x53,
	0x61, 0x6c, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x22, 0x2c, 0x0a, 0x0f, 0x48, 0x6f,
	0x74, 0x53, 0x61, 0x6c, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x12, 0x19, 0x0a,
	0x08, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x03, 0x28, 0x05, 0x52,
	0x07, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x49, 0x64, 0x22, 0x4d, 0x0a, 0x0a, 0x42, 0x75, 0x79, 0x49,
	0x74, 0x65, 0x6d, 0x52, 0x65, 0x71, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x03, 0x75, 0x69, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x69, 0x74, 0x65, 0x6d,
	0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x69, 0x74, 0x65, 0x6d, 0x49,
	0x64, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x22, 0x0d, 0x0a, 0x0b, 0x42, 0x75, 0x79, 0x49, 0x74,
	0x65, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x22, 0x1f, 0x0a, 0x0b, 0x53, 0x61, 0x6c, 0x65, 0x4c, 0x69,
	0x73, 0x74, 0x52, 0x65, 0x71, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x03, 0x75, 0x69, 0x64, 0x22, 0x70, 0x0a, 0x0c, 0x53, 0x61, 0x6c, 0x65, 0x4c,
	0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x12, 0x2d, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x6e, 0x63, 0x73, 0x2e, 0x67, 0x61, 0x6d,
	0x65, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x52,
	0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x6d, 0x6f, 0x6e, 0x65, 0x79, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6d, 0x6f, 0x6e, 0x65, 0x79, 0x12, 0x1b, 0x0a, 0x09,
	0x70, 0x61, 0x73, 0x73, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x08, 0x70, 0x61, 0x73, 0x73, 0x54, 0x79, 0x70, 0x65, 0x22, 0x32, 0x0a, 0x0a, 0x42, 0x75, 0x79,
	0x50, 0x61, 0x73, 0x73, 0x52, 0x65, 0x71, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x75, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x22, 0x0d, 0x0a,
	0x0b, 0x42, 0x75, 0x79, 0x50, 0x61, 0x73, 0x73, 0x52, 0x65, 0x73, 0x70, 0x22, 0x33, 0x0a, 0x09,
	0x42, 0x75, 0x79, 0x56, 0x49, 0x50, 0x52, 0x65, 0x71, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x75, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x6d,
	0x6f, 0x6e, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6d, 0x6f, 0x6e, 0x74,
	0x68, 0x22, 0x0c, 0x0a, 0x0a, 0x42, 0x75, 0x79, 0x56, 0x49, 0x50, 0x52, 0x65, 0x73, 0x70, 0x32,
	0xb3, 0x04, 0x0a, 0x05, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x12, 0x64, 0x0a, 0x07, 0x42, 0x75, 0x79,
	0x49, 0x74, 0x65, 0x6d, 0x12, 0x1d, 0x2e, 0x6e, 0x63, 0x73, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x2e,
	0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x75, 0x79, 0x49, 0x74, 0x65, 0x6d,
	0x52, 0x65, 0x71, 0x1a, 0x1e, 0x2e, 0x6e, 0x63, 0x73, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x2e, 0x73,
	0x74, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x75, 0x79, 0x49, 0x74, 0x65, 0x6d, 0x52,
	0x65, 0x73, 0x70, 0x22, 0x1a, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x14, 0x22, 0x0f, 0x2f, 0x67, 0x61,
	0x6d, 0x65, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2f, 0x62, 0x75, 0x79, 0x3a, 0x01, 0x2a, 0x12,
	0x71, 0x0a, 0x0b, 0x48, 0x6f, 0x74, 0x53, 0x61, 0x6c, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x21,
	0x2e, 0x6e, 0x63, 0x73, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x48, 0x6f, 0x74, 0x53, 0x61, 0x6c, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65,
	0x71, 0x1a, 0x22, 0x2e, 0x6e, 0x63, 0x73, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x2e, 0x73, 0x74, 0x6f,
	0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x48, 0x6f, 0x74, 0x53, 0x61, 0x6c, 0x65, 0x4c, 0x69, 0x73,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x22, 0x1b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x15, 0x12, 0x13, 0x2f,
	0x67, 0x61, 0x6d, 0x65, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2f, 0x68, 0x6f, 0x74, 0x73, 0x61,
	0x6c, 0x65, 0x12, 0x7f, 0x0a, 0x08, 0x53, 0x61, 0x6c, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x1e,
	0x2e, 0x6e, 0x63, 0x73, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x53, 0x61, 0x6c, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x1f,
	0x2e, 0x6e, 0x63, 0x73, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x53, 0x61, 0x6c, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x22,
	0x32, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x2c, 0x12, 0x10, 0x2f, 0x67, 0x61, 0x6d, 0x65, 0x2f, 0x73,
	0x74, 0x6f, 0x72, 0x65, 0x2f, 0x6c, 0x69, 0x73, 0x74, 0x5a, 0x18, 0x12, 0x16, 0x2f, 0x67, 0x61,
	0x6d, 0x65, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2f, 0x6c, 0x69, 0x73, 0x74, 0x2f, 0x7b, 0x75,
	0x69, 0x64, 0x7d, 0x12, 0x69, 0x0a, 0x07, 0x42, 0x75, 0x79, 0x50, 0x61, 0x73, 0x73, 0x12, 0x1d,
	0x2e, 0x6e, 0x63, 0x73, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x42, 0x75, 0x79, 0x50, 0x61, 0x73, 0x73, 0x52, 0x65, 0x71, 0x1a, 0x1e, 0x2e,
	0x6e, 0x63, 0x73, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x42, 0x75, 0x79, 0x50, 0x61, 0x73, 0x73, 0x52, 0x65, 0x73, 0x70, 0x22, 0x1f, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x19, 0x22, 0x14, 0x2f, 0x67, 0x61, 0x6d, 0x65, 0x2f, 0x73, 0x74, 0x6f,
	0x72, 0x65, 0x2f, 0x62, 0x75, 0x79, 0x2f, 0x70, 0x61, 0x73, 0x73, 0x3a, 0x01, 0x2a, 0x12, 0x65,
	0x0a, 0x06, 0x42, 0x75, 0x79, 0x56, 0x49, 0x50, 0x12, 0x1c, 0x2e, 0x6e, 0x63, 0x73, 0x2e, 0x67,
	0x61, 0x6d, 0x65, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x75, 0x79,
	0x56, 0x49, 0x50, 0x52, 0x65, 0x71, 0x1a, 0x1d, 0x2e, 0x6e, 0x63, 0x73, 0x2e, 0x67, 0x61, 0x6d,
	0x65, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x75, 0x79, 0x56, 0x49,
	0x50, 0x52, 0x65, 0x73, 0x70, 0x22, 0x1e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x18, 0x22, 0x13, 0x2f,
	0x67, 0x61, 0x6d, 0x65, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2f, 0x62, 0x75, 0x79, 0x2f, 0x76,
	0x69, 0x70, 0x3a, 0x01, 0x2a, 0x42, 0x24, 0x5a, 0x22, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64,
	0x2f, 0x61, 0x70, 0x70, 0x2f, 0x67, 0x61, 0x6d, 0x65, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_app_game_store_api_grpc_v1_store_proto_rawDescOnce sync.Once
	file_app_game_store_api_grpc_v1_store_proto_rawDescData = file_app_game_store_api_grpc_v1_store_proto_rawDesc
)

func file_app_game_store_api_grpc_v1_store_proto_rawDescGZIP() []byte {
	file_app_game_store_api_grpc_v1_store_proto_rawDescOnce.Do(func() {
		file_app_game_store_api_grpc_v1_store_proto_rawDescData = protoimpl.X.CompressGZIP(file_app_game_store_api_grpc_v1_store_proto_rawDescData)
	})
	return file_app_game_store_api_grpc_v1_store_proto_rawDescData
}

var file_app_game_store_api_grpc_v1_store_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_app_game_store_api_grpc_v1_store_proto_goTypes = []interface{}{
	(*Item)(nil),            // 0: ncs.game.store.v1.Item
	(*HotSaleListReq)(nil),  // 1: ncs.game.store.v1.HotSaleListReq
	(*HotSaleListResp)(nil), // 2: ncs.game.store.v1.HotSaleListResp
	(*BuyItemReq)(nil),      // 3: ncs.game.store.v1.BuyItemReq
	(*BuyItemResp)(nil),     // 4: ncs.game.store.v1.BuyItemResp
	(*SaleListReq)(nil),     // 5: ncs.game.store.v1.SaleListReq
	(*SaleListResp)(nil),    // 6: ncs.game.store.v1.SaleListResp
	(*BuyPassReq)(nil),      // 7: ncs.game.store.v1.BuyPassReq
	(*BuyPassResp)(nil),     // 8: ncs.game.store.v1.BuyPassResp
	(*BuyVIPReq)(nil),       // 9: ncs.game.store.v1.BuyVIPReq
	(*BuyVIPResp)(nil),      // 10: ncs.game.store.v1.BuyVIPResp
}
var file_app_game_store_api_grpc_v1_store_proto_depIdxs = []int32{
	0,  // 0: ncs.game.store.v1.SaleListResp.items:type_name -> ncs.game.store.v1.Item
	3,  // 1: ncs.game.store.v1.Store.BuyItem:input_type -> ncs.game.store.v1.BuyItemReq
	1,  // 2: ncs.game.store.v1.Store.HotSaleList:input_type -> ncs.game.store.v1.HotSaleListReq
	5,  // 3: ncs.game.store.v1.Store.SaleList:input_type -> ncs.game.store.v1.SaleListReq
	7,  // 4: ncs.game.store.v1.Store.BuyPass:input_type -> ncs.game.store.v1.BuyPassReq
	9,  // 5: ncs.game.store.v1.Store.BuyVIP:input_type -> ncs.game.store.v1.BuyVIPReq
	4,  // 6: ncs.game.store.v1.Store.BuyItem:output_type -> ncs.game.store.v1.BuyItemResp
	2,  // 7: ncs.game.store.v1.Store.HotSaleList:output_type -> ncs.game.store.v1.HotSaleListResp
	6,  // 8: ncs.game.store.v1.Store.SaleList:output_type -> ncs.game.store.v1.SaleListResp
	8,  // 9: ncs.game.store.v1.Store.BuyPass:output_type -> ncs.game.store.v1.BuyPassResp
	10, // 10: ncs.game.store.v1.Store.BuyVIP:output_type -> ncs.game.store.v1.BuyVIPResp
	6,  // [6:11] is the sub-list for method output_type
	1,  // [1:6] is the sub-list for method input_type
	1,  // [1:1] is the sub-list for extension type_name
	1,  // [1:1] is the sub-list for extension extendee
	0,  // [0:1] is the sub-list for field type_name
}

func init() { file_app_game_store_api_grpc_v1_store_proto_init() }
func file_app_game_store_api_grpc_v1_store_proto_init() {
	if File_app_game_store_api_grpc_v1_store_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_app_game_store_api_grpc_v1_store_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Item); i {
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
		file_app_game_store_api_grpc_v1_store_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HotSaleListReq); i {
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
		file_app_game_store_api_grpc_v1_store_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HotSaleListResp); i {
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
		file_app_game_store_api_grpc_v1_store_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BuyItemReq); i {
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
		file_app_game_store_api_grpc_v1_store_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BuyItemResp); i {
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
		file_app_game_store_api_grpc_v1_store_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SaleListReq); i {
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
		file_app_game_store_api_grpc_v1_store_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SaleListResp); i {
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
		file_app_game_store_api_grpc_v1_store_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BuyPassReq); i {
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
		file_app_game_store_api_grpc_v1_store_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BuyPassResp); i {
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
		file_app_game_store_api_grpc_v1_store_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BuyVIPReq); i {
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
		file_app_game_store_api_grpc_v1_store_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BuyVIPResp); i {
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
			RawDescriptor: file_app_game_store_api_grpc_v1_store_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_app_game_store_api_grpc_v1_store_proto_goTypes,
		DependencyIndexes: file_app_game_store_api_grpc_v1_store_proto_depIdxs,
		MessageInfos:      file_app_game_store_api_grpc_v1_store_proto_msgTypes,
	}.Build()
	File_app_game_store_api_grpc_v1_store_proto = out.File
	file_app_game_store_api_grpc_v1_store_proto_rawDesc = nil
	file_app_game_store_api_grpc_v1_store_proto_goTypes = nil
	file_app_game_store_api_grpc_v1_store_proto_depIdxs = nil
}
