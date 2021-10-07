// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package login

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// BindTelephoneClient is the client API for BindTelephone service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BindTelephoneClient interface {
	// 检查是否可以绑定
	Check(ctx context.Context, in *CheckRequest, opts ...grpc.CallOption) (*CheckResponse, error)
	// 通过验证码,校验手机号是否可以接收验证码
	BindVerify(ctx context.Context, in *BindVerifyRequest, opts ...grpc.CallOption) (*BindVerifyResponse, error)
	// 通过验证码,进行手机号绑定,换绑
	BindChange(ctx context.Context, in *BindChangeRequest, opts ...grpc.CallOption) (*BindChangeResponse, error)
	// 手机号绑定,不通过验证码
	AutomaticBind(ctx context.Context, in *AutomaticBindRequest, opts ...grpc.CallOption) (*AutomaticBindResponse, error)
	// 绑定手机号解析，通过第三方小程序code换取手机号
	BindPhoneParse(ctx context.Context, in *BindPhoneParseRequest, opts ...grpc.CallOption) (*BindPhoneParseResponse, error)
	// 绑定手机号解析，通过阿里一键
	BindPhoneParseByOneClick(ctx context.Context, in *BindPhoneParseByOneClickRequest, opts ...grpc.CallOption) (*BindPhoneParseByOneClickResponse, error)
}

type bindTelephoneClient struct {
	cc grpc.ClientConnInterface
}

func NewBindTelephoneClient(cc grpc.ClientConnInterface) BindTelephoneClient {
	return &bindTelephoneClient{cc}
}

func (c *bindTelephoneClient) Check(ctx context.Context, in *CheckRequest, opts ...grpc.CallOption) (*CheckResponse, error) {
	out := new(CheckResponse)
	err := c.cc.Invoke(ctx, "/login.BindTelephone/CheckHealth", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bindTelephoneClient) BindVerify(ctx context.Context, in *BindVerifyRequest, opts ...grpc.CallOption) (*BindVerifyResponse, error) {
	out := new(BindVerifyResponse)
	err := c.cc.Invoke(ctx, "/login.BindTelephone/BindVerify", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bindTelephoneClient) BindChange(ctx context.Context, in *BindChangeRequest, opts ...grpc.CallOption) (*BindChangeResponse, error) {
	out := new(BindChangeResponse)
	err := c.cc.Invoke(ctx, "/login.BindTelephone/BindChange", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bindTelephoneClient) AutomaticBind(ctx context.Context, in *AutomaticBindRequest, opts ...grpc.CallOption) (*AutomaticBindResponse, error) {
	out := new(AutomaticBindResponse)
	err := c.cc.Invoke(ctx, "/login.BindTelephone/AutomaticBind", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bindTelephoneClient) BindPhoneParse(ctx context.Context, in *BindPhoneParseRequest, opts ...grpc.CallOption) (*BindPhoneParseResponse, error) {
	out := new(BindPhoneParseResponse)
	err := c.cc.Invoke(ctx, "/login.BindTelephone/BindPhoneParse", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bindTelephoneClient) BindPhoneParseByOneClick(ctx context.Context, in *BindPhoneParseByOneClickRequest, opts ...grpc.CallOption) (*BindPhoneParseByOneClickResponse, error) {
	out := new(BindPhoneParseByOneClickResponse)
	err := c.cc.Invoke(ctx, "/login.BindTelephone/BindPhoneParseByOneClick", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BindTelephoneServer is the server API for BindTelephone service.
// All implementations should embed UnimplementedBindTelephoneServer
// for forward compatibility
type BindTelephoneServer interface {
	// 检查是否可以绑定
	Check(context.Context, *CheckRequest) (*CheckResponse, error)
	// 通过验证码,校验手机号是否可以接收验证码
	BindVerify(context.Context, *BindVerifyRequest) (*BindVerifyResponse, error)
	// 通过验证码,进行手机号绑定,换绑
	BindChange(context.Context, *BindChangeRequest) (*BindChangeResponse, error)
	// 手机号绑定,不通过验证码
	AutomaticBind(context.Context, *AutomaticBindRequest) (*AutomaticBindResponse, error)
	// 绑定手机号解析，通过第三方小程序code换取手机号
	BindPhoneParse(context.Context, *BindPhoneParseRequest) (*BindPhoneParseResponse, error)
	// 绑定手机号解析，通过阿里一键
	BindPhoneParseByOneClick(context.Context, *BindPhoneParseByOneClickRequest) (*BindPhoneParseByOneClickResponse, error)
}

// UnimplementedBindTelephoneServer should be embedded to have forward compatible implementations.
type UnimplementedBindTelephoneServer struct {
}

func (UnimplementedBindTelephoneServer) Check(context.Context, *CheckRequest) (*CheckResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckHealth not implemented")
}
func (UnimplementedBindTelephoneServer) BindVerify(context.Context, *BindVerifyRequest) (*BindVerifyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BindVerify not implemented")
}
func (UnimplementedBindTelephoneServer) BindChange(context.Context, *BindChangeRequest) (*BindChangeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BindChange not implemented")
}
func (UnimplementedBindTelephoneServer) AutomaticBind(context.Context, *AutomaticBindRequest) (*AutomaticBindResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AutomaticBind not implemented")
}
func (UnimplementedBindTelephoneServer) BindPhoneParse(context.Context, *BindPhoneParseRequest) (*BindPhoneParseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BindPhoneParse not implemented")
}
func (UnimplementedBindTelephoneServer) BindPhoneParseByOneClick(context.Context, *BindPhoneParseByOneClickRequest) (*BindPhoneParseByOneClickResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BindPhoneParseByOneClick not implemented")
}

// UnsafeBindTelephoneServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BindTelephoneServer will
// result in compilation errors.
type UnsafeBindTelephoneServer interface {
	mustEmbedUnimplementedBindTelephoneServer()
}

func RegisterBindTelephoneServer(s grpc.ServiceRegistrar, srv BindTelephoneServer) {
	s.RegisterService(&_BindTelephone_serviceDesc, srv)
}

func _BindTelephone_Check_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BindTelephoneServer).Check(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/login.BindTelephone/CheckHealth",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BindTelephoneServer).Check(ctx, req.(*CheckRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BindTelephone_BindVerify_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BindVerifyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BindTelephoneServer).BindVerify(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/login.BindTelephone/BindVerify",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BindTelephoneServer).BindVerify(ctx, req.(*BindVerifyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BindTelephone_BindChange_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BindChangeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BindTelephoneServer).BindChange(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/login.BindTelephone/BindChange",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BindTelephoneServer).BindChange(ctx, req.(*BindChangeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BindTelephone_AutomaticBind_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AutomaticBindRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BindTelephoneServer).AutomaticBind(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/login.BindTelephone/AutomaticBind",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BindTelephoneServer).AutomaticBind(ctx, req.(*AutomaticBindRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BindTelephone_BindPhoneParse_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BindPhoneParseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BindTelephoneServer).BindPhoneParse(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/login.BindTelephone/BindPhoneParse",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BindTelephoneServer).BindPhoneParse(ctx, req.(*BindPhoneParseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BindTelephone_BindPhoneParseByOneClick_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BindPhoneParseByOneClickRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BindTelephoneServer).BindPhoneParseByOneClick(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/login.BindTelephone/BindPhoneParseByOneClick",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BindTelephoneServer).BindPhoneParseByOneClick(ctx, req.(*BindPhoneParseByOneClickRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _BindTelephone_serviceDesc = grpc.ServiceDesc{
	ServiceName: "login.BindTelephone",
	HandlerType: (*BindTelephoneServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CheckHealth",
			Handler:    _BindTelephone_Check_Handler,
		},
		{
			MethodName: "BindVerify",
			Handler:    _BindTelephone_BindVerify_Handler,
		},
		{
			MethodName: "BindChange",
			Handler:    _BindTelephone_BindChange_Handler,
		},
		{
			MethodName: "AutomaticBind",
			Handler:    _BindTelephone_AutomaticBind_Handler,
		},
		{
			MethodName: "BindPhoneParse",
			Handler:    _BindTelephone_BindPhoneParse_Handler,
		},
		{
			MethodName: "BindPhoneParseByOneClick",
			Handler:    _BindTelephone_BindPhoneParseByOneClick_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "example/proto/login/bind.proto",
}
