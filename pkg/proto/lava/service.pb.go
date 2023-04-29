// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: proto/lava/service.proto

package lavapbv1

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

type ServiceInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name      string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Version   string `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
	Path      string `protobuf:"bytes,3,opt,name=path,proto3" json:"path,omitempty"`
	Hostname  string `protobuf:"bytes,4,opt,name=hostname,proto3" json:"hostname,omitempty"`
	Ip        string `protobuf:"bytes,5,opt,name=ip,proto3" json:"ip,omitempty"`
	RequestId string `protobuf:"bytes,6,opt,name=request_id,json=requestId,proto3" json:"request_id,omitempty"`
}

func (x *ServiceInfo) Reset() {
	*x = ServiceInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_lava_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServiceInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServiceInfo) ProtoMessage() {}

func (x *ServiceInfo) ProtoReflect() protoreflect.Message {
	mi := &file_proto_lava_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServiceInfo.ProtoReflect.Descriptor instead.
func (*ServiceInfo) Descriptor() ([]byte, []int) {
	return file_proto_lava_service_proto_rawDescGZIP(), []int{0}
}

func (x *ServiceInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ServiceInfo) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *ServiceInfo) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *ServiceInfo) GetHostname() string {
	if x != nil {
		return x.Hostname
	}
	return ""
}

func (x *ServiceInfo) GetIp() string {
	if x != nil {
		return x.Ip
	}
	return ""
}

func (x *ServiceInfo) GetRequestId() string {
	if x != nil {
		return x.RequestId
	}
	return ""
}

var File_proto_lava_service_proto protoreflect.FileDescriptor

var file_proto_lava_service_proto_rawDesc = []byte{
	0x0a, 0x18, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6c, 0x61, 0x76, 0x61, 0x2f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x6c, 0x61, 0x76, 0x61,
	0x2e, 0x76, 0x31, 0x22, 0x9a, 0x01, 0x0a, 0x0b, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x49,
	0x6e, 0x66, 0x6f, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x70, 0x61, 0x74, 0x68, 0x12, 0x1a, 0x0a, 0x08, 0x68, 0x6f, 0x73, 0x74, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x68, 0x6f, 0x73, 0x74, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x70, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x70, 0x12, 0x1d, 0x0a, 0x0a, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64,
	0x42, 0x33, 0x5a, 0x31, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70,
	0x75, 0x62, 0x67, 0x6f, 0x2f, 0x6c, 0x61, 0x76, 0x61, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x6c, 0x61, 0x76, 0x61, 0x70, 0x62, 0x76, 0x31, 0x3b, 0x6c, 0x61, 0x76,
	0x61, 0x70, 0x62, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_lava_service_proto_rawDescOnce sync.Once
	file_proto_lava_service_proto_rawDescData = file_proto_lava_service_proto_rawDesc
)

func file_proto_lava_service_proto_rawDescGZIP() []byte {
	file_proto_lava_service_proto_rawDescOnce.Do(func() {
		file_proto_lava_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_lava_service_proto_rawDescData)
	})
	return file_proto_lava_service_proto_rawDescData
}

var file_proto_lava_service_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_proto_lava_service_proto_goTypes = []interface{}{
	(*ServiceInfo)(nil), // 0: lava.v1.ServiceInfo
}
var file_proto_lava_service_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_lava_service_proto_init() }
func file_proto_lava_service_proto_init() {
	if File_proto_lava_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_lava_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServiceInfo); i {
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
			RawDescriptor: file_proto_lava_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_lava_service_proto_goTypes,
		DependencyIndexes: file_proto_lava_service_proto_depIdxs,
		MessageInfos:      file_proto_lava_service_proto_msgTypes,
	}.Build()
	File_proto_lava_service_proto = out.File
	file_proto_lava_service_proto_rawDesc = nil
	file_proto_lava_service_proto_goTypes = nil
	file_proto_lava_service_proto_depIdxs = nil
}
