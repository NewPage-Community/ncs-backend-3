// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.21.12
// source: app/hello/api/grpc/v1/hello.proto

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

type HelloReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *HelloReq) Reset() {
	*x = HelloReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_hello_api_grpc_v1_hello_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HelloReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HelloReq) ProtoMessage() {}

func (x *HelloReq) ProtoReflect() protoreflect.Message {
	mi := &file_app_hello_api_grpc_v1_hello_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HelloReq.ProtoReflect.Descriptor instead.
func (*HelloReq) Descriptor() ([]byte, []int) {
	return file_app_hello_api_grpc_v1_hello_proto_rawDescGZIP(), []int{0}
}

type HelloResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *HelloResp) Reset() {
	*x = HelloResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_hello_api_grpc_v1_hello_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HelloResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HelloResp) ProtoMessage() {}

func (x *HelloResp) ProtoReflect() protoreflect.Message {
	mi := &file_app_hello_api_grpc_v1_hello_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HelloResp.ProtoReflect.Descriptor instead.
func (*HelloResp) Descriptor() ([]byte, []int) {
	return file_app_hello_api_grpc_v1_hello_proto_rawDescGZIP(), []int{1}
}

func (x *HelloResp) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_app_hello_api_grpc_v1_hello_proto protoreflect.FileDescriptor

var file_app_hello_api_grpc_v1_hello_proto_rawDesc = []byte{
	0x0a, 0x21, 0x61, 0x70, 0x70, 0x2f, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x67, 0x72, 0x70, 0x63, 0x2f, 0x76, 0x31, 0x2f, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x6e, 0x63, 0x73, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x2e, 0x76,
	0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e,
	0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x0a, 0x0a, 0x08, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x71, 0x22, 0x25, 0x0a, 0x09, 0x48,
	0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x32, 0x51, 0x0a, 0x05, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x12, 0x48, 0x0a, 0x05, 0x48,
	0x65, 0x6c, 0x6c, 0x6f, 0x12, 0x16, 0x2e, 0x6e, 0x63, 0x73, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f,
	0x2e, 0x76, 0x31, 0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x71, 0x1a, 0x17, 0x2e, 0x6e,
	0x63, 0x73, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x48, 0x65, 0x6c, 0x6c,
	0x6f, 0x52, 0x65, 0x73, 0x70, 0x22, 0x0e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x08, 0x12, 0x06, 0x2f,
	0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x42, 0x1f, 0x5a, 0x1d, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64,
	0x2f, 0x61, 0x70, 0x70, 0x2f, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67,
	0x72, 0x70, 0x63, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_app_hello_api_grpc_v1_hello_proto_rawDescOnce sync.Once
	file_app_hello_api_grpc_v1_hello_proto_rawDescData = file_app_hello_api_grpc_v1_hello_proto_rawDesc
)

func file_app_hello_api_grpc_v1_hello_proto_rawDescGZIP() []byte {
	file_app_hello_api_grpc_v1_hello_proto_rawDescOnce.Do(func() {
		file_app_hello_api_grpc_v1_hello_proto_rawDescData = protoimpl.X.CompressGZIP(file_app_hello_api_grpc_v1_hello_proto_rawDescData)
	})
	return file_app_hello_api_grpc_v1_hello_proto_rawDescData
}

var file_app_hello_api_grpc_v1_hello_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_app_hello_api_grpc_v1_hello_proto_goTypes = []interface{}{
	(*HelloReq)(nil),  // 0: ncs.hello.v1.HelloReq
	(*HelloResp)(nil), // 1: ncs.hello.v1.HelloResp
}
var file_app_hello_api_grpc_v1_hello_proto_depIdxs = []int32{
	0, // 0: ncs.hello.v1.Hello.Hello:input_type -> ncs.hello.v1.HelloReq
	1, // 1: ncs.hello.v1.Hello.Hello:output_type -> ncs.hello.v1.HelloResp
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_app_hello_api_grpc_v1_hello_proto_init() }
func file_app_hello_api_grpc_v1_hello_proto_init() {
	if File_app_hello_api_grpc_v1_hello_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_app_hello_api_grpc_v1_hello_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HelloReq); i {
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
		file_app_hello_api_grpc_v1_hello_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HelloResp); i {
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
			RawDescriptor: file_app_hello_api_grpc_v1_hello_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_app_hello_api_grpc_v1_hello_proto_goTypes,
		DependencyIndexes: file_app_hello_api_grpc_v1_hello_proto_depIdxs,
		MessageInfos:      file_app_hello_api_grpc_v1_hello_proto_msgTypes,
	}.Build()
	File_app_hello_api_grpc_v1_hello_proto = out.File
	file_app_hello_api_grpc_v1_hello_proto_rawDesc = nil
	file_app_hello_api_grpc_v1_hello_proto_goTypes = nil
	file_app_hello_api_grpc_v1_hello_proto_depIdxs = nil
}
