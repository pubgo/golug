// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.17.3
// source: example/proto/hello/api.proto

package hello

import (
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/known/structpb"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_example_proto_hello_api_proto protoreflect.FileDescriptor

var file_example_proto_hello_api_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x05, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x31, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f,
	0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x32, 0xa5, 0x01, 0x0a, 0x07, 0x54, 0x65, 0x73, 0x74, 0x41, 0x70, 0x69, 0x12, 0x44,
	0x0a, 0x07, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x2e, 0x68, 0x65, 0x6c, 0x6c,
	0x6f, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x14, 0x2e, 0x68, 0x65, 0x6c, 0x6c,
	0x6f, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x41, 0x70, 0x69, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x22,
	0x13, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0d, 0x12, 0x0b, 0x2f, 0x76, 0x31, 0x2f, 0x76, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x12, 0x54, 0x0a, 0x0b, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x54,
	0x65, 0x73, 0x74, 0x12, 0x0e, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x2e, 0x54, 0x65, 0x73, 0x74,
	0x52, 0x65, 0x71, 0x1a, 0x14, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x2e, 0x54, 0x65, 0x73, 0x74,
	0x41, 0x70, 0x69, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x22, 0x1f, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x19, 0x12, 0x17, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2f, 0x76,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x74, 0x65, 0x73, 0x74, 0x32, 0xbe, 0x01, 0x0a, 0x09, 0x54,
	0x65, 0x73, 0x74, 0x41, 0x70, 0x69, 0x56, 0x32, 0x12, 0x57, 0x0a, 0x08, 0x56, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x31, 0x12, 0x0e, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x2e, 0x54, 0x65, 0x73,
	0x74, 0x52, 0x65, 0x71, 0x1a, 0x14, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x2e, 0x54, 0x65, 0x73,
	0x74, 0x41, 0x70, 0x69, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x22, 0x25, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x1f, 0x22, 0x1a, 0x2f, 0x76, 0x32, 0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2f,
	0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x2f, 0x7b, 0x6e, 0x61, 0x6d, 0x65, 0x7d, 0x3a, 0x01,
	0x2a, 0x12, 0x58, 0x0a, 0x0c, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x54, 0x65, 0x73, 0x74,
	0x31, 0x12, 0x0e, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x52, 0x65,
	0x71, 0x1a, 0x14, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x41, 0x70,
	0x69, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x22, 0x22, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1c, 0x22,
	0x17, 0x2f, 0x76, 0x32, 0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2f, 0x76, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x74, 0x65, 0x73, 0x74, 0x3a, 0x01, 0x2a, 0x42, 0x58, 0x5a, 0x08, 0x2e,
	0x2f, 0x3b, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x92, 0x41, 0x4b, 0x12, 0x05, 0x32, 0x03, 0x31, 0x2e,
	0x30, 0x2a, 0x01, 0x02, 0x72, 0x3f, 0x0a, 0x17, 0x67, 0x52, 0x50, 0x43, 0x20, 0x65, 0x78, 0x61,
	0x6d, 0x70, 0x6c, 0x65, 0x20, 0x72, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x12,
	0x24, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x67, 0x6f, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2d, 0x65, 0x78,
	0x61, 0x6d, 0x70, 0x6c, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_example_proto_hello_api_proto_goTypes = []interface{}{
	(*TestReq)(nil),       // 0: hello.TestReq
	(*TestApiOutput)(nil), // 1: hello.TestApiOutput
}
var file_example_proto_hello_api_proto_depIdxs = []int32{
	0, // 0: hello.TestApi.Version:input_type -> hello.TestReq
	0, // 1: hello.TestApi.VersionTest:input_type -> hello.TestReq
	0, // 2: hello.TestApiV2.Version1:input_type -> hello.TestReq
	0, // 3: hello.TestApiV2.VersionTest1:input_type -> hello.TestReq
	1, // 4: hello.TestApi.Version:output_type -> hello.TestApiOutput
	1, // 5: hello.TestApi.VersionTest:output_type -> hello.TestApiOutput
	1, // 6: hello.TestApiV2.Version1:output_type -> hello.TestApiOutput
	1, // 7: hello.TestApiV2.VersionTest1:output_type -> hello.TestApiOutput
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_example_proto_hello_api_proto_init() }
func file_example_proto_hello_api_proto_init() {
	if File_example_proto_hello_api_proto != nil {
		return
	}
	file_example_proto_hello_api1_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_example_proto_hello_api_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_example_proto_hello_api_proto_goTypes,
		DependencyIndexes: file_example_proto_hello_api_proto_depIdxs,
	}.Build()
	File_example_proto_hello_api_proto = out.File
	file_example_proto_hello_api_proto_rawDesc = nil
	file_example_proto_hello_api_proto_goTypes = nil
	file_example_proto_hello_api_proto_depIdxs = nil
}
