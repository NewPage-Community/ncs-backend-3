// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.12.4
// source: app/game/chat/api/grpc/v1/chat.proto

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

type AllChatReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name     string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Message  string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	ServerId int32  `protobuf:"varint,3,opt,name=server_id,json=serverId,proto3" json:"server_id,omitempty"`
	Uid      int64  `protobuf:"varint,4,opt,name=uid,proto3" json:"uid,omitempty"`
}

func (x *AllChatReq) Reset() {
	*x = AllChatReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_game_chat_api_grpc_v1_chat_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AllChatReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AllChatReq) ProtoMessage() {}

func (x *AllChatReq) ProtoReflect() protoreflect.Message {
	mi := &file_app_game_chat_api_grpc_v1_chat_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AllChatReq.ProtoReflect.Descriptor instead.
func (*AllChatReq) Descriptor() ([]byte, []int) {
	return file_app_game_chat_api_grpc_v1_chat_proto_rawDescGZIP(), []int{0}
}

func (x *AllChatReq) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *AllChatReq) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *AllChatReq) GetServerId() int32 {
	if x != nil {
		return x.ServerId
	}
	return 0
}

func (x *AllChatReq) GetUid() int64 {
	if x != nil {
		return x.Uid
	}
	return 0
}

type AllChatResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *AllChatResp) Reset() {
	*x = AllChatResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_game_chat_api_grpc_v1_chat_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AllChatResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AllChatResp) ProtoMessage() {}

func (x *AllChatResp) ProtoReflect() protoreflect.Message {
	mi := &file_app_game_chat_api_grpc_v1_chat_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AllChatResp.ProtoReflect.Descriptor instead.
func (*AllChatResp) Descriptor() ([]byte, []int) {
	return file_app_game_chat_api_grpc_v1_chat_proto_rawDescGZIP(), []int{1}
}

type ChatNotifyReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid     int64  `protobuf:"varint,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Prefix  string `protobuf:"bytes,2,opt,name=prefix,proto3" json:"prefix,omitempty"`
	Message string `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *ChatNotifyReq) Reset() {
	*x = ChatNotifyReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_game_chat_api_grpc_v1_chat_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChatNotifyReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChatNotifyReq) ProtoMessage() {}

func (x *ChatNotifyReq) ProtoReflect() protoreflect.Message {
	mi := &file_app_game_chat_api_grpc_v1_chat_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChatNotifyReq.ProtoReflect.Descriptor instead.
func (*ChatNotifyReq) Descriptor() ([]byte, []int) {
	return file_app_game_chat_api_grpc_v1_chat_proto_rawDescGZIP(), []int{2}
}

func (x *ChatNotifyReq) GetUid() int64 {
	if x != nil {
		return x.Uid
	}
	return 0
}

func (x *ChatNotifyReq) GetPrefix() string {
	if x != nil {
		return x.Prefix
	}
	return ""
}

func (x *ChatNotifyReq) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type ChatNotifyResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ChatNotifyResp) Reset() {
	*x = ChatNotifyResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_game_chat_api_grpc_v1_chat_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChatNotifyResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChatNotifyResp) ProtoMessage() {}

func (x *ChatNotifyResp) ProtoReflect() protoreflect.Message {
	mi := &file_app_game_chat_api_grpc_v1_chat_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChatNotifyResp.ProtoReflect.Descriptor instead.
func (*ChatNotifyResp) Descriptor() ([]byte, []int) {
	return file_app_game_chat_api_grpc_v1_chat_proto_rawDescGZIP(), []int{3}
}

var File_app_game_chat_api_grpc_v1_chat_proto protoreflect.FileDescriptor

var file_app_game_chat_api_grpc_v1_chat_proto_rawDesc = []byte{
	0x0a, 0x24, 0x61, 0x70, 0x70, 0x2f, 0x67, 0x61, 0x6d, 0x65, 0x2f, 0x63, 0x68, 0x61, 0x74, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x68, 0x61, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x6e, 0x63, 0x73, 0x2e, 0x67, 0x61, 0x6d, 0x65,
	0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x69, 0x0a, 0x0a, 0x41, 0x6c, 0x6c, 0x43, 0x68, 0x61,
	0x74, 0x52, 0x65, 0x71, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x49, 0x64, 0x12,
	0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x75, 0x69,
	0x64, 0x22, 0x0d, 0x0a, 0x0b, 0x41, 0x6c, 0x6c, 0x43, 0x68, 0x61, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x22, 0x53, 0x0a, 0x0d, 0x43, 0x68, 0x61, 0x74, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x52, 0x65,
	0x71, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03,
	0x75, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78, 0x12, 0x18, 0x0a, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x10, 0x0a, 0x0e, 0x43, 0x68, 0x61, 0x74, 0x4e, 0x6f, 0x74,
	0x69, 0x66, 0x79, 0x52, 0x65, 0x73, 0x70, 0x32, 0xbe, 0x01, 0x0a, 0x04, 0x43, 0x68, 0x61, 0x74,
	0x12, 0x65, 0x0a, 0x07, 0x41, 0x6c, 0x6c, 0x43, 0x68, 0x61, 0x74, 0x12, 0x1c, 0x2e, 0x6e, 0x63,
	0x73, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x41,
	0x6c, 0x6c, 0x43, 0x68, 0x61, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x1d, 0x2e, 0x6e, 0x63, 0x73, 0x2e,
	0x67, 0x61, 0x6d, 0x65, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x6c, 0x6c,
	0x43, 0x68, 0x61, 0x74, 0x52, 0x65, 0x73, 0x70, 0x22, 0x1d, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x17,
	0x22, 0x12, 0x2f, 0x67, 0x61, 0x6d, 0x65, 0x2f, 0x63, 0x68, 0x61, 0x74, 0x2f, 0x61, 0x6c, 0x6c,
	0x63, 0x68, 0x61, 0x74, 0x3a, 0x01, 0x2a, 0x12, 0x4f, 0x0a, 0x0a, 0x43, 0x68, 0x61, 0x74, 0x4e,
	0x6f, 0x74, 0x69, 0x66, 0x79, 0x12, 0x1f, 0x2e, 0x6e, 0x63, 0x73, 0x2e, 0x67, 0x61, 0x6d, 0x65,
	0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x68, 0x61, 0x74, 0x4e, 0x6f, 0x74,
	0x69, 0x66, 0x79, 0x52, 0x65, 0x71, 0x1a, 0x20, 0x2e, 0x6e, 0x63, 0x73, 0x2e, 0x67, 0x61, 0x6d,
	0x65, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x68, 0x61, 0x74, 0x4e, 0x6f,
	0x74, 0x69, 0x66, 0x79, 0x52, 0x65, 0x73, 0x70, 0x42, 0x23, 0x5a, 0x21, 0x62, 0x61, 0x63, 0x6b,
	0x65, 0x6e, 0x64, 0x2f, 0x61, 0x70, 0x70, 0x2f, 0x67, 0x61, 0x6d, 0x65, 0x2f, 0x63, 0x68, 0x61,
	0x74, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_app_game_chat_api_grpc_v1_chat_proto_rawDescOnce sync.Once
	file_app_game_chat_api_grpc_v1_chat_proto_rawDescData = file_app_game_chat_api_grpc_v1_chat_proto_rawDesc
)

func file_app_game_chat_api_grpc_v1_chat_proto_rawDescGZIP() []byte {
	file_app_game_chat_api_grpc_v1_chat_proto_rawDescOnce.Do(func() {
		file_app_game_chat_api_grpc_v1_chat_proto_rawDescData = protoimpl.X.CompressGZIP(file_app_game_chat_api_grpc_v1_chat_proto_rawDescData)
	})
	return file_app_game_chat_api_grpc_v1_chat_proto_rawDescData
}

var file_app_game_chat_api_grpc_v1_chat_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_app_game_chat_api_grpc_v1_chat_proto_goTypes = []interface{}{
	(*AllChatReq)(nil),     // 0: ncs.game.chat.v1.AllChatReq
	(*AllChatResp)(nil),    // 1: ncs.game.chat.v1.AllChatResp
	(*ChatNotifyReq)(nil),  // 2: ncs.game.chat.v1.ChatNotifyReq
	(*ChatNotifyResp)(nil), // 3: ncs.game.chat.v1.ChatNotifyResp
}
var file_app_game_chat_api_grpc_v1_chat_proto_depIdxs = []int32{
	0, // 0: ncs.game.chat.v1.Chat.AllChat:input_type -> ncs.game.chat.v1.AllChatReq
	2, // 1: ncs.game.chat.v1.Chat.ChatNotify:input_type -> ncs.game.chat.v1.ChatNotifyReq
	1, // 2: ncs.game.chat.v1.Chat.AllChat:output_type -> ncs.game.chat.v1.AllChatResp
	3, // 3: ncs.game.chat.v1.Chat.ChatNotify:output_type -> ncs.game.chat.v1.ChatNotifyResp
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_app_game_chat_api_grpc_v1_chat_proto_init() }
func file_app_game_chat_api_grpc_v1_chat_proto_init() {
	if File_app_game_chat_api_grpc_v1_chat_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_app_game_chat_api_grpc_v1_chat_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AllChatReq); i {
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
		file_app_game_chat_api_grpc_v1_chat_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AllChatResp); i {
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
		file_app_game_chat_api_grpc_v1_chat_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChatNotifyReq); i {
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
		file_app_game_chat_api_grpc_v1_chat_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChatNotifyResp); i {
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
			RawDescriptor: file_app_game_chat_api_grpc_v1_chat_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_app_game_chat_api_grpc_v1_chat_proto_goTypes,
		DependencyIndexes: file_app_game_chat_api_grpc_v1_chat_proto_depIdxs,
		MessageInfos:      file_app_game_chat_api_grpc_v1_chat_proto_msgTypes,
	}.Build()
	File_app_game_chat_api_grpc_v1_chat_proto = out.File
	file_app_game_chat_api_grpc_v1_chat_proto_rawDesc = nil
	file_app_game_chat_api_grpc_v1_chat_proto_goTypes = nil
	file_app_game_chat_api_grpc_v1_chat_proto_depIdxs = nil
}
