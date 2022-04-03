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
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// CodeClient is the client API for Code service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CodeClient interface {
	// 发送
	SendCode(ctx context.Context, in *SendCodeRequest, opts ...grpc.CallOption) (*SendCodeResponse, error)
	// 校验
	Verify(ctx context.Context, in *VerifyRequest, opts ...grpc.CallOption) (*VerifyResponse, error)
	// 是否校验图片验证码
	IsCheckImageCode(ctx context.Context, in *IsCheckImageCodeRequest, opts ...grpc.CallOption) (*IsCheckImageCodeResponse, error)
	// 校验图片验证码
	VerifyImageCode(ctx context.Context, in *VerifyImageCodeRequest, opts ...grpc.CallOption) (*VerifyImageCodeResponse, error)
	// 获取发送状态
	GetSendStatus(ctx context.Context, in *GetSendStatusRequest, opts ...grpc.CallOption) (*GetSendStatusResponse, error)
}

type codeClient struct {
	cc grpc.ClientConnInterface
}

func NewCodeClient(cc grpc.ClientConnInterface) CodeClient {
	return &codeClient{cc}
}

func (c *codeClient) SendCode(ctx context.Context, in *SendCodeRequest, opts ...grpc.CallOption) (*SendCodeResponse, error) {
	out := new(SendCodeResponse)
	err := c.cc.Invoke(ctx, "/login.Code/SendCode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *codeClient) Verify(ctx context.Context, in *VerifyRequest, opts ...grpc.CallOption) (*VerifyResponse, error) {
	out := new(VerifyResponse)
	err := c.cc.Invoke(ctx, "/login.Code/Verify", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *codeClient) IsCheckImageCode(ctx context.Context, in *IsCheckImageCodeRequest, opts ...grpc.CallOption) (*IsCheckImageCodeResponse, error) {
	out := new(IsCheckImageCodeResponse)
	err := c.cc.Invoke(ctx, "/login.Code/IsCheckImageCode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *codeClient) VerifyImageCode(ctx context.Context, in *VerifyImageCodeRequest, opts ...grpc.CallOption) (*VerifyImageCodeResponse, error) {
	out := new(VerifyImageCodeResponse)
	err := c.cc.Invoke(ctx, "/login.Code/VerifyImageCode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *codeClient) GetSendStatus(ctx context.Context, in *GetSendStatusRequest, opts ...grpc.CallOption) (*GetSendStatusResponse, error) {
	out := new(GetSendStatusResponse)
	err := c.cc.Invoke(ctx, "/login.Code/GetSendStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CodeServer is the server API for Code service.
// All implementations must embed UnimplementedCodeServer
// for forward compatibility
type CodeServer interface {
	// 发送
	SendCode(context.Context, *SendCodeRequest) (*SendCodeResponse, error)
	// 校验
	Verify(context.Context, *VerifyRequest) (*VerifyResponse, error)
	// 是否校验图片验证码
	IsCheckImageCode(context.Context, *IsCheckImageCodeRequest) (*IsCheckImageCodeResponse, error)
	// 校验图片验证码
	VerifyImageCode(context.Context, *VerifyImageCodeRequest) (*VerifyImageCodeResponse, error)
	// 获取发送状态
	GetSendStatus(context.Context, *GetSendStatusRequest) (*GetSendStatusResponse, error)
	mustEmbedUnimplementedCodeServer()
}

// UnimplementedCodeServer must be embedded to have forward compatible implementations.
type UnimplementedCodeServer struct {
}

func (UnimplementedCodeServer) SendCode(context.Context, *SendCodeRequest) (*SendCodeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendCode not implemented")
}
func (UnimplementedCodeServer) Verify(context.Context, *VerifyRequest) (*VerifyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Verify not implemented")
}
func (UnimplementedCodeServer) IsCheckImageCode(context.Context, *IsCheckImageCodeRequest) (*IsCheckImageCodeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IsCheckImageCode not implemented")
}
func (UnimplementedCodeServer) VerifyImageCode(context.Context, *VerifyImageCodeRequest) (*VerifyImageCodeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VerifyImageCode not implemented")
}
func (UnimplementedCodeServer) GetSendStatus(context.Context, *GetSendStatusRequest) (*GetSendStatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSendStatus not implemented")
}
func (UnimplementedCodeServer) mustEmbedUnimplementedCodeServer() {}

// UnsafeCodeServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CodeServer will
// result in compilation errors.
type UnsafeCodeServer interface {
	mustEmbedUnimplementedCodeServer()
}

func RegisterCodeServer(s grpc.ServiceRegistrar, srv CodeServer) {
	s.RegisterService(&Code_ServiceDesc, srv)
}

func _Code_SendCode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendCodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CodeServer).SendCode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/login.Code/SendCode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CodeServer).SendCode(ctx, req.(*SendCodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Code_Verify_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VerifyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CodeServer).Verify(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/login.Code/Verify",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CodeServer).Verify(ctx, req.(*VerifyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Code_IsCheckImageCode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IsCheckImageCodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CodeServer).IsCheckImageCode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/login.Code/IsCheckImageCode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CodeServer).IsCheckImageCode(ctx, req.(*IsCheckImageCodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Code_VerifyImageCode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VerifyImageCodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CodeServer).VerifyImageCode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/login.Code/VerifyImageCode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CodeServer).VerifyImageCode(ctx, req.(*VerifyImageCodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Code_GetSendStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSendStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CodeServer).GetSendStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/login.Code/GetSendStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CodeServer).GetSendStatus(ctx, req.(*GetSendStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Code_ServiceDesc is the grpc.ServiceDesc for Code service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Code_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "login.Code",
	HandlerType: (*CodeServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendCode",
			Handler:    _Code_SendCode_Handler,
		},
		{
			MethodName: "Verify",
			Handler:    _Code_Verify_Handler,
		},
		{
			MethodName: "IsCheckImageCode",
			Handler:    _Code_IsCheckImageCode_Handler,
		},
		{
			MethodName: "VerifyImageCode",
			Handler:    _Code_VerifyImageCode_Handler,
		},
		{
			MethodName: "GetSendStatus",
			Handler:    _Code_GetSendStatus_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/login/code.proto",
}
