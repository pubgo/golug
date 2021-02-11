// Code generated by protoc-gen-go. DO NOT EDIT.
// source: example/proto/login/merge.proto

// 账户合并相关

package login

import (
	context "context"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type TelephoneRequest struct {
	// 登陆用户
	Uid int64 `protobuf:"varint,1,opt,name=uid,proto3" json:"uid,omitempty"`
	// 新手机号
	TargetTelephone string `protobuf:"bytes,2,opt,name=targetTelephone,proto3" json:"targetTelephone,omitempty"`
	// 是否走新流程
	IsNewProcess         bool     `protobuf:"varint,3,opt,name=isNewProcess,proto3" json:"isNewProcess,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TelephoneRequest) Reset()         { *m = TelephoneRequest{} }
func (m *TelephoneRequest) String() string { return proto.CompactTextString(m) }
func (*TelephoneRequest) ProtoMessage()    {}
func (*TelephoneRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_81493a7b4908a7a9, []int{0}
}

func (m *TelephoneRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TelephoneRequest.Unmarshal(m, b)
}
func (m *TelephoneRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TelephoneRequest.Marshal(b, m, deterministic)
}
func (m *TelephoneRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TelephoneRequest.Merge(m, src)
}
func (m *TelephoneRequest) XXX_Size() int {
	return xxx_messageInfo_TelephoneRequest.Size(m)
}
func (m *TelephoneRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_TelephoneRequest.DiscardUnknown(m)
}

var xxx_messageInfo_TelephoneRequest proto.InternalMessageInfo

func (m *TelephoneRequest) GetUid() int64 {
	if m != nil {
		return m.Uid
	}
	return 0
}

func (m *TelephoneRequest) GetTargetTelephone() string {
	if m != nil {
		return m.TargetTelephone
	}
	return ""
}

func (m *TelephoneRequest) GetIsNewProcess() bool {
	if m != nil {
		return m.IsNewProcess
	}
	return false
}

type WeChatRequest struct {
	// 登陆用户
	Uid int64 `protobuf:"varint,1,opt,name=uid,proto3" json:"uid,omitempty"`
	// 要合并的用户
	TargetUid            int64    `protobuf:"varint,2,opt,name=targetUid,proto3" json:"targetUid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WeChatRequest) Reset()         { *m = WeChatRequest{} }
func (m *WeChatRequest) String() string { return proto.CompactTextString(m) }
func (*WeChatRequest) ProtoMessage()    {}
func (*WeChatRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_81493a7b4908a7a9, []int{1}
}

func (m *WeChatRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WeChatRequest.Unmarshal(m, b)
}
func (m *WeChatRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WeChatRequest.Marshal(b, m, deterministic)
}
func (m *WeChatRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WeChatRequest.Merge(m, src)
}
func (m *WeChatRequest) XXX_Size() int {
	return xxx_messageInfo_WeChatRequest.Size(m)
}
func (m *WeChatRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_WeChatRequest.DiscardUnknown(m)
}

var xxx_messageInfo_WeChatRequest proto.InternalMessageInfo

func (m *WeChatRequest) GetUid() int64 {
	if m != nil {
		return m.Uid
	}
	return 0
}

func (m *WeChatRequest) GetTargetUid() int64 {
	if m != nil {
		return m.TargetUid
	}
	return 0
}

type WeChatUnMergeRequest struct {
	// 登陆用户
	Uid                  int64    `protobuf:"varint,1,opt,name=uid,proto3" json:"uid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WeChatUnMergeRequest) Reset()         { *m = WeChatUnMergeRequest{} }
func (m *WeChatUnMergeRequest) String() string { return proto.CompactTextString(m) }
func (*WeChatUnMergeRequest) ProtoMessage()    {}
func (*WeChatUnMergeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_81493a7b4908a7a9, []int{2}
}

func (m *WeChatUnMergeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WeChatUnMergeRequest.Unmarshal(m, b)
}
func (m *WeChatUnMergeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WeChatUnMergeRequest.Marshal(b, m, deterministic)
}
func (m *WeChatUnMergeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WeChatUnMergeRequest.Merge(m, src)
}
func (m *WeChatUnMergeRequest) XXX_Size() int {
	return xxx_messageInfo_WeChatUnMergeRequest.Size(m)
}
func (m *WeChatUnMergeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_WeChatUnMergeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_WeChatUnMergeRequest proto.InternalMessageInfo

func (m *WeChatUnMergeRequest) GetUid() int64 {
	if m != nil {
		return m.Uid
	}
	return 0
}

type Reply struct {
	// code
	Code int64 `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	// msg
	Msg string `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	// 时间戳
	NowTime int64 `protobuf:"varint,3,opt,name=nowTime,proto3" json:"nowTime,omitempty"`
	// 数据
	Data                 map[string]string `protobuf:"bytes,4,rep,name=data,proto3" json:"data,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Reply) Reset()         { *m = Reply{} }
func (m *Reply) String() string { return proto.CompactTextString(m) }
func (*Reply) ProtoMessage()    {}
func (*Reply) Descriptor() ([]byte, []int) {
	return fileDescriptor_81493a7b4908a7a9, []int{3}
}

func (m *Reply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Reply.Unmarshal(m, b)
}
func (m *Reply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Reply.Marshal(b, m, deterministic)
}
func (m *Reply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Reply.Merge(m, src)
}
func (m *Reply) XXX_Size() int {
	return xxx_messageInfo_Reply.Size(m)
}
func (m *Reply) XXX_DiscardUnknown() {
	xxx_messageInfo_Reply.DiscardUnknown(m)
}

var xxx_messageInfo_Reply proto.InternalMessageInfo

func (m *Reply) GetCode() int64 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *Reply) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func (m *Reply) GetNowTime() int64 {
	if m != nil {
		return m.NowTime
	}
	return 0
}

func (m *Reply) GetData() map[string]string {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*TelephoneRequest)(nil), "login.TelephoneRequest")
	proto.RegisterType((*WeChatRequest)(nil), "login.WeChatRequest")
	proto.RegisterType((*WeChatUnMergeRequest)(nil), "login.WeChatUnMergeRequest")
	proto.RegisterType((*Reply)(nil), "login.Reply")
	proto.RegisterMapType((map[string]string)(nil), "login.Reply.DataEntry")
}

func init() { proto.RegisterFile("example/proto/login/merge.proto", fileDescriptor_81493a7b4908a7a9) }

var fileDescriptor_81493a7b4908a7a9 = []byte{
	// 544 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x93, 0xcd, 0x6e, 0xd3, 0x40,
	0x10, 0xc7, 0xe5, 0x38, 0x69, 0x9b, 0x4d, 0x0b, 0xd1, 0x36, 0x40, 0x9a, 0x46, 0xb5, 0xb1, 0xf8,
	0x88, 0x8a, 0x12, 0x8b, 0x52, 0x09, 0x14, 0x89, 0x4b, 0x0a, 0x37, 0x40, 0x68, 0x95, 0x2a, 0x47,
	0xb4, 0x71, 0x06, 0xc7, 0xaa, 0xed, 0x0d, 0xf6, 0x9a, 0x90, 0x2b, 0xaf, 0xc0, 0x1b, 0xf0, 0x38,
	0x5c, 0x91, 0x38, 0xfa, 0xc4, 0xc9, 0x4f, 0x81, 0x76, 0xd7, 0x49, 0x1a, 0x2b, 0x12, 0xb9, 0x8c,
	0x66, 0x66, 0x77, 0x7e, 0xff, 0xf1, 0xce, 0x18, 0x19, 0xf0, 0x8d, 0x06, 0x33, 0x1f, 0xec, 0x59,
	0xc4, 0x38, 0xb3, 0x7d, 0xe6, 0x7a, 0xa1, 0x1d, 0x40, 0xe4, 0x42, 0x4f, 0x66, 0x70, 0x45, 0xa6,
	0x5a, 0x5d, 0xd7, 0xe3, 0xd3, 0x64, 0xdc, 0x73, 0x58, 0x60, 0xbb, 0xcc, 0x65, 0xea, 0xfe, 0x38,
	0xf9, 0x2c, 0x23, 0x55, 0x2c, 0x3c, 0x55, 0xd5, 0x6a, 0xbb, 0x8c, 0xb9, 0x3e, 0xd8, 0x74, 0xe6,
	0xd9, 0x34, 0x0c, 0x19, 0xa7, 0xdc, 0x63, 0x61, 0xac, 0x4e, 0xad, 0x9f, 0x1a, 0xaa, 0x0f, 0xc1,
	0x87, 0xd9, 0x94, 0x85, 0x40, 0xe0, 0x4b, 0x02, 0x31, 0xc7, 0x27, 0x48, 0x4f, 0xbc, 0x49, 0x53,
	0x33, 0xb5, 0x8e, 0x3e, 0xd8, 0xcf, 0x52, 0x43, 0x84, 0x44, 0x18, 0xfc, 0x1a, 0xdd, 0xe5, 0x34,
	0x72, 0x81, 0xaf, 0x8a, 0x9a, 0x25, 0x53, 0xeb, 0x54, 0x07, 0xc7, 0x59, 0x6a, 0x14, 0x8f, 0x48,
	0x31, 0x81, 0x2f, 0xd1, 0xa1, 0x17, 0x7f, 0x80, 0xf9, 0xc7, 0x88, 0x39, 0x10, 0xc7, 0x4d, 0xdd,
	0xd4, 0x3a, 0x07, 0x83, 0x7a, 0x96, 0x1a, 0x1b, 0x79, 0xb2, 0x11, 0x59, 0x23, 0x74, 0x34, 0x82,
	0xab, 0x29, 0xe5, 0x3b, 0x34, 0xf8, 0x0c, 0x55, 0x95, 0xe8, 0xb5, 0x37, 0x91, 0xad, 0xe9, 0x83,
	0xa3, 0x2c, 0x35, 0xd6, 0x49, 0xb2, 0x76, 0xad, 0xe7, 0xa8, 0xa1, 0xc0, 0xd7, 0xe1, 0x7b, 0xf1,
	0xd0, 0xff, 0xe7, 0x5b, 0x7f, 0x34, 0x54, 0x21, 0x30, 0xf3, 0x17, 0xb8, 0x8d, 0xca, 0x0e, 0x9b,
	0x40, 0x7e, 0xeb, 0x20, 0x4b, 0x0d, 0x19, 0x13, 0x69, 0x05, 0x22, 0x88, 0xdd, 0xfc, 0x71, 0x24,
	0x22, 0x88, 0x5d, 0x22, 0x0c, 0x7e, 0x8c, 0xf6, 0x43, 0x36, 0x1f, 0x7a, 0x01, 0xc8, 0xef, 0xd7,
	0x07, 0xb5, 0x2c, 0x35, 0x96, 0x29, 0xb2, 0x74, 0xf0, 0x25, 0x2a, 0x4f, 0x28, 0xa7, 0xcd, 0xb2,
	0xa9, 0x77, 0x6a, 0x17, 0xf7, 0x7b, 0x72, 0xfa, 0x3d, 0xa9, 0xdd, 0x7b, 0x43, 0x39, 0x7d, 0x1b,
	0xf2, 0x68, 0xa1, 0x74, 0xc5, 0x3d, 0x22, 0x6d, 0xeb, 0x25, 0xaa, 0xae, 0x0e, 0x71, 0x1d, 0xe9,
	0x37, 0xb0, 0x90, 0x1d, 0x56, 0x89, 0x70, 0x71, 0x03, 0x55, 0xbe, 0x52, 0x3f, 0xc9, 0xa7, 0x46,
	0x54, 0xd0, 0x2f, 0xbd, 0xd2, 0x2e, 0x7e, 0xe9, 0xa8, 0x22, 0x1f, 0x01, 0x0f, 0x51, 0x75, 0x3d,
	0xb1, 0x07, 0xb9, 0x6e, 0x71, 0x49, 0x5a, 0x87, 0xb7, 0x1b, 0xb2, 0xcc, 0xef, 0xbf, 0xff, 0xfe,
	0x28, 0xb5, 0xac, 0x7b, 0x76, 0x12, 0x43, 0xa4, 0xd6, 0xd6, 0xe6, 0xcb, 0x9a, 0xbe, 0x76, 0x8e,
	0x3f, 0xa1, 0x3b, 0x2b, 0xc6, 0xd5, 0x14, 0x9c, 0x9b, 0x5d, 0xd1, 0x4f, 0x24, 0xda, 0xb4, 0x4e,
	0xb7, 0xa2, 0xbb, 0x8e, 0x60, 0x09, 0x81, 0x77, 0x68, 0x4f, 0x0d, 0x13, 0x37, 0xf2, 0xfa, 0x8d,
	0xa5, 0x29, 0x50, 0xcf, 0x24, 0xb5, 0x69, 0x1d, 0xdf, 0xa6, 0xce, 0x05, 0x8e, 0x72, 0x41, 0x1b,
	0xa1, 0x9a, 0x2a, 0x57, 0xbd, 0xee, 0x82, 0x7c, 0x24, 0x91, 0x67, 0xd6, 0xc9, 0x16, 0xe4, 0xba,
	0x4d, 0x67, 0xb9, 0xcc, 0xf9, 0xce, 0xe1, 0xd3, 0x0d, 0xf4, 0xe6, 0x26, 0x16, 0x14, 0x9e, 0x4a,
	0x85, 0x87, 0x56, 0x7b, 0x9b, 0x42, 0x12, 0x76, 0x65, 0xa2, 0xaf, 0x9d, 0x8f, 0xf7, 0xe4, 0xdf,
	0xfd, 0xe2, 0x5f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x3c, 0xbf, 0xc1, 0x7c, 0x54, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MergeClient is the client API for Merge service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MergeClient interface {
	// 手机号合并,换绑,手机号更换
	Telephone(ctx context.Context, in *TelephoneRequest, opts ...grpc.CallOption) (*Reply, error)
	// 手机号账户合并检查
	TelephoneCheck(ctx context.Context, in *TelephoneRequest, opts ...grpc.CallOption) (*Reply, error)
	// 微信账户绑定
	WeChat(ctx context.Context, in *WeChatRequest, opts ...grpc.CallOption) (*Reply, error)
	// 微信合并检查
	WeChatCheck(ctx context.Context, in *WeChatRequest, opts ...grpc.CallOption) (*Reply, error)
	// 解除微信绑定, 必须拥有手机号
	WeChatUnMerge(ctx context.Context, in *WeChatUnMergeRequest, opts ...grpc.CallOption) (*Reply, error)
}

type mergeClient struct {
	cc *grpc.ClientConn
}

func NewMergeClient(cc *grpc.ClientConn) MergeClient {
	return &mergeClient{cc}
}

func (c *mergeClient) Telephone(ctx context.Context, in *TelephoneRequest, opts ...grpc.CallOption) (*Reply, error) {
	out := new(Reply)
	err := c.cc.Invoke(ctx, "/login.Merge/Telephone", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mergeClient) TelephoneCheck(ctx context.Context, in *TelephoneRequest, opts ...grpc.CallOption) (*Reply, error) {
	out := new(Reply)
	err := c.cc.Invoke(ctx, "/login.Merge/TelephoneCheck", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mergeClient) WeChat(ctx context.Context, in *WeChatRequest, opts ...grpc.CallOption) (*Reply, error) {
	out := new(Reply)
	err := c.cc.Invoke(ctx, "/login.Merge/WeChat", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mergeClient) WeChatCheck(ctx context.Context, in *WeChatRequest, opts ...grpc.CallOption) (*Reply, error) {
	out := new(Reply)
	err := c.cc.Invoke(ctx, "/login.Merge/WeChatCheck", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mergeClient) WeChatUnMerge(ctx context.Context, in *WeChatUnMergeRequest, opts ...grpc.CallOption) (*Reply, error) {
	out := new(Reply)
	err := c.cc.Invoke(ctx, "/login.Merge/WeChatUnMerge", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MergeServer is the server API for Merge service.
type MergeServer interface {
	// 手机号合并,换绑,手机号更换
	Telephone(context.Context, *TelephoneRequest) (*Reply, error)
	// 手机号账户合并检查
	TelephoneCheck(context.Context, *TelephoneRequest) (*Reply, error)
	// 微信账户绑定
	WeChat(context.Context, *WeChatRequest) (*Reply, error)
	// 微信合并检查
	WeChatCheck(context.Context, *WeChatRequest) (*Reply, error)
	// 解除微信绑定, 必须拥有手机号
	WeChatUnMerge(context.Context, *WeChatUnMergeRequest) (*Reply, error)
}

// UnimplementedMergeServer can be embedded to have forward compatible implementations.
type UnimplementedMergeServer struct {
}

func (*UnimplementedMergeServer) Telephone(ctx context.Context, req *TelephoneRequest) (*Reply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Telephone not implemented")
}
func (*UnimplementedMergeServer) TelephoneCheck(ctx context.Context, req *TelephoneRequest) (*Reply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TelephoneCheck not implemented")
}
func (*UnimplementedMergeServer) WeChat(ctx context.Context, req *WeChatRequest) (*Reply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method WeChat not implemented")
}
func (*UnimplementedMergeServer) WeChatCheck(ctx context.Context, req *WeChatRequest) (*Reply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method WeChatCheck not implemented")
}
func (*UnimplementedMergeServer) WeChatUnMerge(ctx context.Context, req *WeChatUnMergeRequest) (*Reply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method WeChatUnMerge not implemented")
}

func RegisterMergeServer(s *grpc.Server, srv MergeServer) {
	s.RegisterService(&_Merge_serviceDesc, srv)
}

func _Merge_Telephone_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TelephoneRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MergeServer).Telephone(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/login.Merge/Telephone",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MergeServer).Telephone(ctx, req.(*TelephoneRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Merge_TelephoneCheck_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TelephoneRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MergeServer).TelephoneCheck(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/login.Merge/TelephoneCheck",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MergeServer).TelephoneCheck(ctx, req.(*TelephoneRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Merge_WeChat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WeChatRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MergeServer).WeChat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/login.Merge/WeChat",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MergeServer).WeChat(ctx, req.(*WeChatRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Merge_WeChatCheck_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WeChatRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MergeServer).WeChatCheck(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/login.Merge/WeChatCheck",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MergeServer).WeChatCheck(ctx, req.(*WeChatRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Merge_WeChatUnMerge_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WeChatUnMergeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MergeServer).WeChatUnMerge(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/login.Merge/WeChatUnMerge",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MergeServer).WeChatUnMerge(ctx, req.(*WeChatUnMergeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Merge_serviceDesc = grpc.ServiceDesc{
	ServiceName: "login.Merge",
	HandlerType: (*MergeServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Telephone",
			Handler:    _Merge_Telephone_Handler,
		},
		{
			MethodName: "TelephoneCheck",
			Handler:    _Merge_TelephoneCheck_Handler,
		},
		{
			MethodName: "WeChat",
			Handler:    _Merge_WeChat_Handler,
		},
		{
			MethodName: "WeChatCheck",
			Handler:    _Merge_WeChatCheck_Handler,
		},
		{
			MethodName: "WeChatUnMerge",
			Handler:    _Merge_WeChatUnMerge_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "example/proto/login/merge.proto",
}
