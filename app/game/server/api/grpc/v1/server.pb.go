// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.12.4
// source: app/game/server/api/grpc/v1/server.proto

package v1

import (
	v1 "backend/app/game/a2s/api/grpc/v1"
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

	ServerId  int32  `protobuf:"varint,1,opt,name=server_id,json=serverId,proto3" json:"server_id,omitempty"`
	ModId     int32  `protobuf:"varint,2,opt,name=mod_id,json=modId,proto3" json:"mod_id,omitempty"`
	GameId    int32  `protobuf:"varint,3,opt,name=game_id,json=gameId,proto3" json:"game_id,omitempty"`
	Rcon      string `protobuf:"bytes,4,opt,name=rcon,proto3" json:"rcon,omitempty"`
	Hostname  string `protobuf:"bytes,5,opt,name=hostname,proto3" json:"hostname,omitempty"`
	Address   string `protobuf:"bytes,6,opt,name=address,proto3" json:"address,omitempty"`
	ShortName string `protobuf:"bytes,7,opt,name=short_name,json=shortName,proto3" json:"short_name,omitempty"`
	// a2s
	A2SInfo   *v1.A2SInfo     `protobuf:"bytes,8,opt,name=a2s_info,json=a2sInfo,proto3" json:"a2s_info,omitempty"`
	A2SPlayer []*v1.A2SPlayer `protobuf:"bytes,9,rep,name=a2s_player,json=a2sPlayer,proto3" json:"a2s_player,omitempty"`
}

func (x *Info) Reset() {
	*x = Info{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_game_server_api_grpc_v1_server_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Info) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Info) ProtoMessage() {}

func (x *Info) ProtoReflect() protoreflect.Message {
	mi := &file_app_game_server_api_grpc_v1_server_proto_msgTypes[0]
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
	return file_app_game_server_api_grpc_v1_server_proto_rawDescGZIP(), []int{0}
}

func (x *Info) GetServerId() int32 {
	if x != nil {
		return x.ServerId
	}
	return 0
}

func (x *Info) GetModId() int32 {
	if x != nil {
		return x.ModId
	}
	return 0
}

func (x *Info) GetGameId() int32 {
	if x != nil {
		return x.GameId
	}
	return 0
}

func (x *Info) GetRcon() string {
	if x != nil {
		return x.Rcon
	}
	return ""
}

func (x *Info) GetHostname() string {
	if x != nil {
		return x.Hostname
	}
	return ""
}

func (x *Info) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *Info) GetShortName() string {
	if x != nil {
		return x.ShortName
	}
	return ""
}

func (x *Info) GetA2SInfo() *v1.A2SInfo {
	if x != nil {
		return x.A2SInfo
	}
	return nil
}

func (x *Info) GetA2SPlayer() []*v1.A2SPlayer {
	if x != nil {
		return x.A2SPlayer
	}
	return nil
}

type InitReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Port    int32  `protobuf:"varint,2,opt,name=port,proto3" json:"port,omitempty"`
}

func (x *InitReq) Reset() {
	*x = InitReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_game_server_api_grpc_v1_server_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InitReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InitReq) ProtoMessage() {}

func (x *InitReq) ProtoReflect() protoreflect.Message {
	mi := &file_app_game_server_api_grpc_v1_server_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InitReq.ProtoReflect.Descriptor instead.
func (*InitReq) Descriptor() ([]byte, []int) {
	return file_app_game_server_api_grpc_v1_server_proto_rawDescGZIP(), []int{1}
}

func (x *InitReq) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *InitReq) GetPort() int32 {
	if x != nil {
		return x.Port
	}
	return 0
}

type InitResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info *Info `protobuf:"bytes,1,opt,name=info,proto3" json:"info,omitempty"`
}

func (x *InitResp) Reset() {
	*x = InitResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_game_server_api_grpc_v1_server_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InitResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InitResp) ProtoMessage() {}

func (x *InitResp) ProtoReflect() protoreflect.Message {
	mi := &file_app_game_server_api_grpc_v1_server_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InitResp.ProtoReflect.Descriptor instead.
func (*InitResp) Descriptor() ([]byte, []int) {
	return file_app_game_server_api_grpc_v1_server_proto_rawDescGZIP(), []int{2}
}

func (x *InitResp) GetInfo() *Info {
	if x != nil {
		return x.Info
	}
	return nil
}

type InfoReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ServerId int32 `protobuf:"varint,1,opt,name=server_id,json=serverId,proto3" json:"server_id,omitempty"`
}

func (x *InfoReq) Reset() {
	*x = InfoReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_game_server_api_grpc_v1_server_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InfoReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InfoReq) ProtoMessage() {}

func (x *InfoReq) ProtoReflect() protoreflect.Message {
	mi := &file_app_game_server_api_grpc_v1_server_proto_msgTypes[3]
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
	return file_app_game_server_api_grpc_v1_server_proto_rawDescGZIP(), []int{3}
}

func (x *InfoReq) GetServerId() int32 {
	if x != nil {
		return x.ServerId
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
		mi := &file_app_game_server_api_grpc_v1_server_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InfoResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InfoResp) ProtoMessage() {}

func (x *InfoResp) ProtoReflect() protoreflect.Message {
	mi := &file_app_game_server_api_grpc_v1_server_proto_msgTypes[4]
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
	return file_app_game_server_api_grpc_v1_server_proto_rawDescGZIP(), []int{4}
}

func (x *InfoResp) GetInfo() *Info {
	if x != nil {
		return x.Info
	}
	return nil
}

type AllInfoReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	A2S bool `protobuf:"varint,1,opt,name=a2s,proto3" json:"a2s,omitempty"`
}

func (x *AllInfoReq) Reset() {
	*x = AllInfoReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_game_server_api_grpc_v1_server_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AllInfoReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AllInfoReq) ProtoMessage() {}

func (x *AllInfoReq) ProtoReflect() protoreflect.Message {
	mi := &file_app_game_server_api_grpc_v1_server_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AllInfoReq.ProtoReflect.Descriptor instead.
func (*AllInfoReq) Descriptor() ([]byte, []int) {
	return file_app_game_server_api_grpc_v1_server_proto_rawDescGZIP(), []int{5}
}

func (x *AllInfoReq) GetA2S() bool {
	if x != nil {
		return x.A2S
	}
	return false
}

type AllInfoResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info []*Info `protobuf:"bytes,1,rep,name=info,proto3" json:"info,omitempty"`
}

func (x *AllInfoResp) Reset() {
	*x = AllInfoResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_game_server_api_grpc_v1_server_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AllInfoResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AllInfoResp) ProtoMessage() {}

func (x *AllInfoResp) ProtoReflect() protoreflect.Message {
	mi := &file_app_game_server_api_grpc_v1_server_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AllInfoResp.ProtoReflect.Descriptor instead.
func (*AllInfoResp) Descriptor() ([]byte, []int) {
	return file_app_game_server_api_grpc_v1_server_proto_rawDescGZIP(), []int{6}
}

func (x *AllInfoResp) GetInfo() []*Info {
	if x != nil {
		return x.Info
	}
	return nil
}

type RconReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ServerId int32  `protobuf:"varint,1,opt,name=server_id,json=serverId,proto3" json:"server_id,omitempty"`
	Cmd      string `protobuf:"bytes,2,opt,name=cmd,proto3" json:"cmd,omitempty"`
}

func (x *RconReq) Reset() {
	*x = RconReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_game_server_api_grpc_v1_server_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RconReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RconReq) ProtoMessage() {}

func (x *RconReq) ProtoReflect() protoreflect.Message {
	mi := &file_app_game_server_api_grpc_v1_server_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RconReq.ProtoReflect.Descriptor instead.
func (*RconReq) Descriptor() ([]byte, []int) {
	return file_app_game_server_api_grpc_v1_server_proto_rawDescGZIP(), []int{7}
}

func (x *RconReq) GetServerId() int32 {
	if x != nil {
		return x.ServerId
	}
	return 0
}

func (x *RconReq) GetCmd() string {
	if x != nil {
		return x.Cmd
	}
	return ""
}

type RconResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Response string `protobuf:"bytes,1,opt,name=response,proto3" json:"response,omitempty"`
}

func (x *RconResp) Reset() {
	*x = RconResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_game_server_api_grpc_v1_server_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RconResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RconResp) ProtoMessage() {}

func (x *RconResp) ProtoReflect() protoreflect.Message {
	mi := &file_app_game_server_api_grpc_v1_server_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RconResp.ProtoReflect.Descriptor instead.
func (*RconResp) Descriptor() ([]byte, []int) {
	return file_app_game_server_api_grpc_v1_server_proto_rawDescGZIP(), []int{8}
}

func (x *RconResp) GetResponse() string {
	if x != nil {
		return x.Response
	}
	return ""
}

type RconAllReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Cmd string `protobuf:"bytes,1,opt,name=cmd,proto3" json:"cmd,omitempty"`
}

func (x *RconAllReq) Reset() {
	*x = RconAllReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_game_server_api_grpc_v1_server_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RconAllReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RconAllReq) ProtoMessage() {}

func (x *RconAllReq) ProtoReflect() protoreflect.Message {
	mi := &file_app_game_server_api_grpc_v1_server_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RconAllReq.ProtoReflect.Descriptor instead.
func (*RconAllReq) Descriptor() ([]byte, []int) {
	return file_app_game_server_api_grpc_v1_server_proto_rawDescGZIP(), []int{9}
}

func (x *RconAllReq) GetCmd() string {
	if x != nil {
		return x.Cmd
	}
	return ""
}

type RconAllResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success int32 `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *RconAllResp) Reset() {
	*x = RconAllResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_game_server_api_grpc_v1_server_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RconAllResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RconAllResp) ProtoMessage() {}

func (x *RconAllResp) ProtoReflect() protoreflect.Message {
	mi := &file_app_game_server_api_grpc_v1_server_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RconAllResp.ProtoReflect.Descriptor instead.
func (*RconAllResp) Descriptor() ([]byte, []int) {
	return file_app_game_server_api_grpc_v1_server_proto_rawDescGZIP(), []int{10}
}

func (x *RconAllResp) GetSuccess() int32 {
	if x != nil {
		return x.Success
	}
	return 0
}

type ChangeMapNotifyReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ServerId int32  `protobuf:"varint,1,opt,name=server_id,json=serverId,proto3" json:"server_id,omitempty"`
	Map      string `protobuf:"bytes,2,opt,name=map,proto3" json:"map,omitempty"`
	// MQ
	ServerName string `protobuf:"bytes,3,opt,name=server_name,json=serverName,proto3" json:"server_name,omitempty"`
}

func (x *ChangeMapNotifyReq) Reset() {
	*x = ChangeMapNotifyReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_game_server_api_grpc_v1_server_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChangeMapNotifyReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChangeMapNotifyReq) ProtoMessage() {}

func (x *ChangeMapNotifyReq) ProtoReflect() protoreflect.Message {
	mi := &file_app_game_server_api_grpc_v1_server_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChangeMapNotifyReq.ProtoReflect.Descriptor instead.
func (*ChangeMapNotifyReq) Descriptor() ([]byte, []int) {
	return file_app_game_server_api_grpc_v1_server_proto_rawDescGZIP(), []int{11}
}

func (x *ChangeMapNotifyReq) GetServerId() int32 {
	if x != nil {
		return x.ServerId
	}
	return 0
}

func (x *ChangeMapNotifyReq) GetMap() string {
	if x != nil {
		return x.Map
	}
	return ""
}

func (x *ChangeMapNotifyReq) GetServerName() string {
	if x != nil {
		return x.ServerName
	}
	return ""
}

type ChangeMapNotifyResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ChangeMapNotifyResp) Reset() {
	*x = ChangeMapNotifyResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_game_server_api_grpc_v1_server_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChangeMapNotifyResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChangeMapNotifyResp) ProtoMessage() {}

func (x *ChangeMapNotifyResp) ProtoReflect() protoreflect.Message {
	mi := &file_app_game_server_api_grpc_v1_server_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChangeMapNotifyResp.ProtoReflect.Descriptor instead.
func (*ChangeMapNotifyResp) Descriptor() ([]byte, []int) {
	return file_app_game_server_api_grpc_v1_server_proto_rawDescGZIP(), []int{12}
}

var File_app_game_server_api_grpc_v1_server_proto protoreflect.FileDescriptor

var file_app_game_server_api_grpc_v1_server_proto_rawDesc = []byte{
	0x0a, 0x28, 0x61, 0x70, 0x70, 0x2f, 0x67, 0x61, 0x6d, 0x65, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12, 0x6e, 0x63, 0x73, 0x2e,
	0x67, 0x61, 0x6d, 0x65, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x1a, 0x1c,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x22, 0x61, 0x70,
	0x70, 0x2f, 0x67, 0x61, 0x6d, 0x65, 0x2f, 0x61, 0x32, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67,
	0x72, 0x70, 0x63, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x32, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xac, 0x02, 0x0a, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x49, 0x64, 0x12, 0x15, 0x0a, 0x06, 0x6d, 0x6f, 0x64, 0x5f, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6d, 0x6f, 0x64, 0x49, 0x64, 0x12, 0x17, 0x0a,
	0x07, 0x67, 0x61, 0x6d, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06,
	0x67, 0x61, 0x6d, 0x65, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x63, 0x6f, 0x6e, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x72, 0x63, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x68, 0x6f,
	0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x68, 0x6f,
	0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x33, 0x0a, 0x08, 0x61, 0x32, 0x73, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x18, 0x2e, 0x6e, 0x63, 0x73, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x2e, 0x61, 0x32, 0x73,
	0x2e, 0x76, 0x31, 0x2e, 0x41, 0x32, 0x53, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x07, 0x61, 0x32, 0x73,
	0x49, 0x6e, 0x66, 0x6f, 0x12, 0x39, 0x0a, 0x0a, 0x61, 0x32, 0x73, 0x5f, 0x70, 0x6c, 0x61, 0x79,
	0x65, 0x72, 0x18, 0x09, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x6e, 0x63, 0x73, 0x2e, 0x67,
	0x61, 0x6d, 0x65, 0x2e, 0x61, 0x32, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x32, 0x53, 0x50, 0x6c,
	0x61, 0x79, 0x65, 0x72, 0x52, 0x09, 0x61, 0x32, 0x73, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x22,
	0x37, 0x0a, 0x07, 0x49, 0x6e, 0x69, 0x74, 0x52, 0x65, 0x71, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x22, 0x38, 0x0a, 0x08, 0x49, 0x6e, 0x69, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x12, 0x2c, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x18, 0x2e, 0x6e, 0x63, 0x73, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x69, 0x6e,
	0x66, 0x6f, 0x22, 0x26, 0x0a, 0x07, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x12, 0x1b, 0x0a,
	0x09, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x08, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x49, 0x64, 0x22, 0x38, 0x0a, 0x08, 0x49, 0x6e,
	0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x12, 0x2c, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x6e, 0x63, 0x73, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04,
	0x69, 0x6e, 0x66, 0x6f, 0x22, 0x1e, 0x0a, 0x0a, 0x41, 0x6c, 0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x52,
	0x65, 0x71, 0x12, 0x10, 0x0a, 0x03, 0x61, 0x32, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x03, 0x61, 0x32, 0x73, 0x22, 0x3b, 0x0a, 0x0b, 0x41, 0x6c, 0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x52,
	0x65, 0x73, 0x70, 0x12, 0x2c, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x18, 0x2e, 0x6e, 0x63, 0x73, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x69, 0x6e, 0x66,
	0x6f, 0x22, 0x38, 0x0a, 0x07, 0x52, 0x63, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x12, 0x1b, 0x0a, 0x09,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x08, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x49, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x63, 0x6d, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x63, 0x6d, 0x64, 0x22, 0x26, 0x0a, 0x08, 0x52,
	0x63, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x1e, 0x0a, 0x0a, 0x52, 0x63, 0x6f, 0x6e, 0x41, 0x6c, 0x6c, 0x52, 0x65,
	0x71, 0x12, 0x10, 0x0a, 0x03, 0x63, 0x6d, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x63, 0x6d, 0x64, 0x22, 0x27, 0x0a, 0x0b, 0x52, 0x63, 0x6f, 0x6e, 0x41, 0x6c, 0x6c, 0x52, 0x65,
	0x73, 0x70, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x22, 0x64, 0x0a, 0x12,
	0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x4d, 0x61, 0x70, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x52,
	0x65, 0x71, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x49, 0x64, 0x12,
	0x10, 0x0a, 0x03, 0x6d, 0x61, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x61,
	0x70, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x4e, 0x61,
	0x6d, 0x65, 0x22, 0x15, 0x0a, 0x13, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x4d, 0x61, 0x70, 0x4e,
	0x6f, 0x74, 0x69, 0x66, 0x79, 0x52, 0x65, 0x73, 0x70, 0x32, 0x94, 0x05, 0x0a, 0x06, 0x53, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x12, 0x68, 0x0a, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1b, 0x2e, 0x6e,
	0x63, 0x73, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x76,
	0x31, 0x2e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x1a, 0x1c, 0x2e, 0x6e, 0x63, 0x73, 0x2e,
	0x67, 0x61, 0x6d, 0x65, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x49,
	0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x22, 0x25, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1f, 0x12,
	0x1d, 0x2f, 0x67, 0x61, 0x6d, 0x65, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x69, 0x6e,
	0x66, 0x6f, 0x2f, 0x7b, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x7d, 0x12, 0x68,
	0x0a, 0x07, 0x41, 0x6c, 0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1e, 0x2e, 0x6e, 0x63, 0x73, 0x2e,
	0x67, 0x61, 0x6d, 0x65, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x41,
	0x6c, 0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x1a, 0x1f, 0x2e, 0x6e, 0x63, 0x73, 0x2e,
	0x67, 0x61, 0x6d, 0x65, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x41,
	0x6c, 0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x22, 0x1c, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x16, 0x12, 0x14, 0x2f, 0x67, 0x61, 0x6d, 0x65, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x2f, 0x61, 0x6c, 0x6c, 0x69, 0x6e, 0x66, 0x6f, 0x12, 0x5f, 0x0a, 0x04, 0x49, 0x6e, 0x69, 0x74,
	0x12, 0x1b, 0x2e, 0x6e, 0x63, 0x73, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e, 0x69, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x1c, 0x2e,
	0x6e, 0x63, 0x73, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e,
	0x76, 0x31, 0x2e, 0x49, 0x6e, 0x69, 0x74, 0x52, 0x65, 0x73, 0x70, 0x22, 0x1c, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x16, 0x22, 0x11, 0x2f, 0x67, 0x61, 0x6d, 0x65, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x2f, 0x69, 0x6e, 0x69, 0x74, 0x3a, 0x01, 0x2a, 0x12, 0x5f, 0x0a, 0x04, 0x52, 0x63, 0x6f,
	0x6e, 0x12, 0x1b, 0x2e, 0x6e, 0x63, 0x73, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x63, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x1a, 0x1c,
	0x2e, 0x6e, 0x63, 0x73, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x2e, 0x76, 0x31, 0x2e, 0x52, 0x63, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x22, 0x1c, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x16, 0x22, 0x11, 0x2f, 0x67, 0x61, 0x6d, 0x65, 0x2f, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x2f, 0x72, 0x63, 0x6f, 0x6e, 0x3a, 0x01, 0x2a, 0x12, 0x6c, 0x0a, 0x07, 0x52, 0x63,
	0x6f, 0x6e, 0x41, 0x6c, 0x6c, 0x12, 0x1e, 0x2e, 0x6e, 0x63, 0x73, 0x2e, 0x67, 0x61, 0x6d, 0x65,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x63, 0x6f, 0x6e, 0x41,
	0x6c, 0x6c, 0x52, 0x65, 0x71, 0x1a, 0x1f, 0x2e, 0x6e, 0x63, 0x73, 0x2e, 0x67, 0x61, 0x6d, 0x65,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x63, 0x6f, 0x6e, 0x41,
	0x6c, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x22, 0x20, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1a, 0x22, 0x15,
	0x2f, 0x67, 0x61, 0x6d, 0x65, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x72, 0x63, 0x6f,
	0x6e, 0x2f, 0x61, 0x6c, 0x6c, 0x3a, 0x01, 0x2a, 0x12, 0x85, 0x01, 0x0a, 0x0f, 0x43, 0x68, 0x61,
	0x6e, 0x67, 0x65, 0x4d, 0x61, 0x70, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x12, 0x26, 0x2e, 0x6e,
	0x63, 0x73, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x76,
	0x31, 0x2e, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x4d, 0x61, 0x70, 0x4e, 0x6f, 0x74, 0x69, 0x66,
	0x79, 0x52, 0x65, 0x71, 0x1a, 0x27, 0x2e, 0x6e, 0x63, 0x73, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65,
	0x4d, 0x61, 0x70, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x52, 0x65, 0x73, 0x70, 0x22, 0x21, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x1b, 0x22, 0x16, 0x2f, 0x67, 0x61, 0x6d, 0x65, 0x2f, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x2f, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x6d, 0x61, 0x70, 0x3a, 0x01, 0x2a,
	0x42, 0x25, 0x5a, 0x23, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2f, 0x61, 0x70, 0x70, 0x2f,
	0x67, 0x61, 0x6d, 0x65, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x67, 0x72, 0x70, 0x63, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_app_game_server_api_grpc_v1_server_proto_rawDescOnce sync.Once
	file_app_game_server_api_grpc_v1_server_proto_rawDescData = file_app_game_server_api_grpc_v1_server_proto_rawDesc
)

func file_app_game_server_api_grpc_v1_server_proto_rawDescGZIP() []byte {
	file_app_game_server_api_grpc_v1_server_proto_rawDescOnce.Do(func() {
		file_app_game_server_api_grpc_v1_server_proto_rawDescData = protoimpl.X.CompressGZIP(file_app_game_server_api_grpc_v1_server_proto_rawDescData)
	})
	return file_app_game_server_api_grpc_v1_server_proto_rawDescData
}

var file_app_game_server_api_grpc_v1_server_proto_msgTypes = make([]protoimpl.MessageInfo, 13)
var file_app_game_server_api_grpc_v1_server_proto_goTypes = []interface{}{
	(*Info)(nil),                // 0: ncs.game.server.v1.Info
	(*InitReq)(nil),             // 1: ncs.game.server.v1.InitReq
	(*InitResp)(nil),            // 2: ncs.game.server.v1.InitResp
	(*InfoReq)(nil),             // 3: ncs.game.server.v1.InfoReq
	(*InfoResp)(nil),            // 4: ncs.game.server.v1.InfoResp
	(*AllInfoReq)(nil),          // 5: ncs.game.server.v1.AllInfoReq
	(*AllInfoResp)(nil),         // 6: ncs.game.server.v1.AllInfoResp
	(*RconReq)(nil),             // 7: ncs.game.server.v1.RconReq
	(*RconResp)(nil),            // 8: ncs.game.server.v1.RconResp
	(*RconAllReq)(nil),          // 9: ncs.game.server.v1.RconAllReq
	(*RconAllResp)(nil),         // 10: ncs.game.server.v1.RconAllResp
	(*ChangeMapNotifyReq)(nil),  // 11: ncs.game.server.v1.ChangeMapNotifyReq
	(*ChangeMapNotifyResp)(nil), // 12: ncs.game.server.v1.ChangeMapNotifyResp
	(*v1.A2SInfo)(nil),          // 13: ncs.game.a2s.v1.A2SInfo
	(*v1.A2SPlayer)(nil),        // 14: ncs.game.a2s.v1.A2SPlayer
}
var file_app_game_server_api_grpc_v1_server_proto_depIdxs = []int32{
	13, // 0: ncs.game.server.v1.Info.a2s_info:type_name -> ncs.game.a2s.v1.A2SInfo
	14, // 1: ncs.game.server.v1.Info.a2s_player:type_name -> ncs.game.a2s.v1.A2SPlayer
	0,  // 2: ncs.game.server.v1.InitResp.info:type_name -> ncs.game.server.v1.Info
	0,  // 3: ncs.game.server.v1.InfoResp.info:type_name -> ncs.game.server.v1.Info
	0,  // 4: ncs.game.server.v1.AllInfoResp.info:type_name -> ncs.game.server.v1.Info
	3,  // 5: ncs.game.server.v1.Server.Info:input_type -> ncs.game.server.v1.InfoReq
	5,  // 6: ncs.game.server.v1.Server.AllInfo:input_type -> ncs.game.server.v1.AllInfoReq
	1,  // 7: ncs.game.server.v1.Server.Init:input_type -> ncs.game.server.v1.InitReq
	7,  // 8: ncs.game.server.v1.Server.Rcon:input_type -> ncs.game.server.v1.RconReq
	9,  // 9: ncs.game.server.v1.Server.RconAll:input_type -> ncs.game.server.v1.RconAllReq
	11, // 10: ncs.game.server.v1.Server.ChangeMapNotify:input_type -> ncs.game.server.v1.ChangeMapNotifyReq
	4,  // 11: ncs.game.server.v1.Server.Info:output_type -> ncs.game.server.v1.InfoResp
	6,  // 12: ncs.game.server.v1.Server.AllInfo:output_type -> ncs.game.server.v1.AllInfoResp
	2,  // 13: ncs.game.server.v1.Server.Init:output_type -> ncs.game.server.v1.InitResp
	8,  // 14: ncs.game.server.v1.Server.Rcon:output_type -> ncs.game.server.v1.RconResp
	10, // 15: ncs.game.server.v1.Server.RconAll:output_type -> ncs.game.server.v1.RconAllResp
	12, // 16: ncs.game.server.v1.Server.ChangeMapNotify:output_type -> ncs.game.server.v1.ChangeMapNotifyResp
	11, // [11:17] is the sub-list for method output_type
	5,  // [5:11] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_app_game_server_api_grpc_v1_server_proto_init() }
func file_app_game_server_api_grpc_v1_server_proto_init() {
	if File_app_game_server_api_grpc_v1_server_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_app_game_server_api_grpc_v1_server_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_app_game_server_api_grpc_v1_server_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InitReq); i {
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
		file_app_game_server_api_grpc_v1_server_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InitResp); i {
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
		file_app_game_server_api_grpc_v1_server_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
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
		file_app_game_server_api_grpc_v1_server_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
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
		file_app_game_server_api_grpc_v1_server_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AllInfoReq); i {
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
		file_app_game_server_api_grpc_v1_server_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AllInfoResp); i {
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
		file_app_game_server_api_grpc_v1_server_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RconReq); i {
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
		file_app_game_server_api_grpc_v1_server_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RconResp); i {
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
		file_app_game_server_api_grpc_v1_server_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RconAllReq); i {
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
		file_app_game_server_api_grpc_v1_server_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RconAllResp); i {
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
		file_app_game_server_api_grpc_v1_server_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChangeMapNotifyReq); i {
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
		file_app_game_server_api_grpc_v1_server_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChangeMapNotifyResp); i {
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
			RawDescriptor: file_app_game_server_api_grpc_v1_server_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   13,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_app_game_server_api_grpc_v1_server_proto_goTypes,
		DependencyIndexes: file_app_game_server_api_grpc_v1_server_proto_depIdxs,
		MessageInfos:      file_app_game_server_api_grpc_v1_server_proto_msgTypes,
	}.Build()
	File_app_game_server_api_grpc_v1_server_proto = out.File
	file_app_game_server_api_grpc_v1_server_proto_rawDesc = nil
	file_app_game_server_api_grpc_v1_server_proto_goTypes = nil
	file_app_game_server_api_grpc_v1_server_proto_depIdxs = nil
}
