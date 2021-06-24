// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.17.3
// source: example/proto/hello/api1.proto

package hello

import (
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

//枚举消息类型
type PhoneType int32

const (
	PhoneType_MOBILE PhoneType = 0 //proto3版本中，首成员必须为0，成员不应有相同的值
	PhoneType_HOME   PhoneType = 1
	PhoneType_WORK   PhoneType = 2
)

// Enum value maps for PhoneType.
var (
	PhoneType_name = map[int32]string{
		0: "MOBILE",
		1: "HOME",
		2: "WORK",
	}
	PhoneType_value = map[string]int32{
		"MOBILE": 0,
		"HOME":   1,
		"WORK":   2,
	}
)

func (x PhoneType) Enum() *PhoneType {
	p := new(PhoneType)
	*p = x
	return p
}

func (x PhoneType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PhoneType) Descriptor() protoreflect.EnumDescriptor {
	return file_example_proto_hello_api1_proto_enumTypes[0].Descriptor()
}

func (PhoneType) Type() protoreflect.EnumType {
	return &file_example_proto_hello_api1_proto_enumTypes[0]
}

func (x PhoneType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PhoneType.Descriptor instead.
func (PhoneType) EnumDescriptor() ([]byte, []int) {
	return file_example_proto_hello_api1_proto_rawDescGZIP(), []int{0}
}

type TestReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Input string `protobuf:"bytes,1,opt,name=input,proto3" json:"input,omitempty"`
	Name  string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *TestReq) Reset() {
	*x = TestReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_example_proto_hello_api1_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TestReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TestReq) ProtoMessage() {}

func (x *TestReq) ProtoReflect() protoreflect.Message {
	mi := &file_example_proto_hello_api1_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TestReq.ProtoReflect.Descriptor instead.
func (*TestReq) Descriptor() ([]byte, []int) {
	return file_example_proto_hello_api1_proto_rawDescGZIP(), []int{0}
}

func (x *TestReq) GetInput() string {
	if x != nil {
		return x.Input
	}
	return ""
}

func (x *TestReq) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type TestApiData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Version    string `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"`
	SrvVersion string `protobuf:"bytes,2,opt,name=srvVersion,proto3" json:"srvVersion,omitempty"`
}

func (x *TestApiData) Reset() {
	*x = TestApiData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_example_proto_hello_api1_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TestApiData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TestApiData) ProtoMessage() {}

func (x *TestApiData) ProtoReflect() protoreflect.Message {
	mi := &file_example_proto_hello_api1_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TestApiData.ProtoReflect.Descriptor instead.
func (*TestApiData) Descriptor() ([]byte, []int) {
	return file_example_proto_hello_api1_proto_rawDescGZIP(), []int{1}
}

func (x *TestApiData) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *TestApiData) GetSrvVersion() string {
	if x != nil {
		return x.SrvVersion
	}
	return ""
}

type TestApiOutput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    int32        `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Msg     string       `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	NowTime int64        `protobuf:"varint,3,opt,name=nowTime,proto3" json:"nowTime,omitempty"`
	Data    *TestApiData `protobuf:"bytes,4,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *TestApiOutput) Reset() {
	*x = TestApiOutput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_example_proto_hello_api1_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TestApiOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TestApiOutput) ProtoMessage() {}

func (x *TestApiOutput) ProtoReflect() protoreflect.Message {
	mi := &file_example_proto_hello_api1_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TestApiOutput.ProtoReflect.Descriptor instead.
func (*TestApiOutput) Descriptor() ([]byte, []int) {
	return file_example_proto_hello_api1_proto_rawDescGZIP(), []int{2}
}

func (x *TestApiOutput) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *TestApiOutput) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *TestApiOutput) GetNowTime() int64 {
	if x != nil {
		return x.NowTime
	}
	return 0
}

func (x *TestApiOutput) GetData() *TestApiData {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_example_proto_hello_api1_proto protoreflect.FileDescriptor

var file_example_proto_hello_api1_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x05, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x22, 0x33, 0x0a, 0x07, 0x54, 0x65, 0x73, 0x74, 0x52,
	0x65, 0x71, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x47, 0x0a, 0x0b,
	0x54, 0x65, 0x73, 0x74, 0x41, 0x70, 0x69, 0x44, 0x61, 0x74, 0x61, 0x12, 0x18, 0x0a, 0x07, 0x76,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x72, 0x76, 0x56, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x72, 0x76, 0x56, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0x77, 0x0a, 0x0d, 0x54, 0x65, 0x73, 0x74, 0x41, 0x70, 0x69,
	0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73,
	0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x12, 0x18, 0x0a, 0x07,
	0x6e, 0x6f, 0x77, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x6e,
	0x6f, 0x77, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x26, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x2e, 0x54, 0x65, 0x73,
	0x74, 0x41, 0x70, 0x69, 0x44, 0x61, 0x74, 0x61, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x2a, 0x2b,
	0x0a, 0x09, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0a, 0x0a, 0x06, 0x4d,
	0x4f, 0x42, 0x49, 0x4c, 0x45, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x48, 0x4f, 0x4d, 0x45, 0x10,
	0x01, 0x12, 0x08, 0x0a, 0x04, 0x57, 0x4f, 0x52, 0x4b, 0x10, 0x02, 0x42, 0x0a, 0x5a, 0x08, 0x2e,
	0x2f, 0x3b, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_example_proto_hello_api1_proto_rawDescOnce sync.Once
	file_example_proto_hello_api1_proto_rawDescData = file_example_proto_hello_api1_proto_rawDesc
)

func file_example_proto_hello_api1_proto_rawDescGZIP() []byte {
	file_example_proto_hello_api1_proto_rawDescOnce.Do(func() {
		file_example_proto_hello_api1_proto_rawDescData = protoimpl.X.CompressGZIP(file_example_proto_hello_api1_proto_rawDescData)
	})
	return file_example_proto_hello_api1_proto_rawDescData
}

var file_example_proto_hello_api1_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_example_proto_hello_api1_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_example_proto_hello_api1_proto_goTypes = []interface{}{
	(PhoneType)(0),        // 0: hello.PhoneType
	(*TestReq)(nil),       // 1: hello.TestReq
	(*TestApiData)(nil),   // 2: hello.TestApiData
	(*TestApiOutput)(nil), // 3: hello.TestApiOutput
}
var file_example_proto_hello_api1_proto_depIdxs = []int32{
	2, // 0: hello.TestApiOutput.data:type_name -> hello.TestApiData
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_example_proto_hello_api1_proto_init() }
func file_example_proto_hello_api1_proto_init() {
	if File_example_proto_hello_api1_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_example_proto_hello_api1_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TestReq); i {
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
		file_example_proto_hello_api1_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TestApiData); i {
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
		file_example_proto_hello_api1_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TestApiOutput); i {
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
			RawDescriptor: file_example_proto_hello_api1_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_example_proto_hello_api1_proto_goTypes,
		DependencyIndexes: file_example_proto_hello_api1_proto_depIdxs,
		EnumInfos:         file_example_proto_hello_api1_proto_enumTypes,
		MessageInfos:      file_example_proto_hello_api1_proto_msgTypes,
	}.Build()
	File_example_proto_hello_api1_proto = out.File
	file_example_proto_hello_api1_proto_rawDesc = nil
	file_example_proto_hello_api1_proto_goTypes = nil
	file_example_proto_hello_api1_proto_depIdxs = nil
}
