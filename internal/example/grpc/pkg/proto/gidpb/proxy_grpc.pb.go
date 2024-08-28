// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v5.27.1
// source: gid/proxy.proto

package gidpb

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

const (
	IdProxy_Echo_FullMethodName = "/gid.IdProxy/Echo"
)

// IdProxyClient is the client API for IdProxy service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type IdProxyClient interface {
	Echo(ctx context.Context, in *EchoReq, opts ...grpc.CallOption) (*EchoRsp, error)
}

type idProxyClient struct {
	cc grpc.ClientConnInterface
}

func NewIdProxyClient(cc grpc.ClientConnInterface) IdProxyClient {
	return &idProxyClient{cc}
}

func (c *idProxyClient) Echo(ctx context.Context, in *EchoReq, opts ...grpc.CallOption) (*EchoRsp, error) {
	out := new(EchoRsp)
	err := c.cc.Invoke(ctx, IdProxy_Echo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// IdProxyServer is the server API for IdProxy service.
// All implementations should embed UnimplementedIdProxyServer
// for forward compatibility
type IdProxyServer interface {
	Echo(context.Context, *EchoReq) (*EchoRsp, error)
}

// UnimplementedIdProxyServer should be embedded to have forward compatible implementations.
type UnimplementedIdProxyServer struct {
}

func (UnimplementedIdProxyServer) Echo(context.Context, *EchoReq) (*EchoRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Echo not implemented")
}

// UnsafeIdProxyServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to IdProxyServer will
// result in compilation errors.
type UnsafeIdProxyServer interface {
	mustEmbedUnimplementedIdProxyServer()
}

func RegisterIdProxyServer(s grpc.ServiceRegistrar, srv IdProxyServer) {
	s.RegisterService(&IdProxy_ServiceDesc, srv)
}

func _IdProxy_Echo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EchoReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IdProxyServer).Echo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: IdProxy_Echo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IdProxyServer).Echo(ctx, req.(*EchoReq))
	}
	return interceptor(ctx, in, info, handler)
}

// IdProxy_ServiceDesc is the grpc.ServiceDesc for IdProxy service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var IdProxy_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "gid.IdProxy",
	HandlerType: (*IdProxyServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Echo",
			Handler:    _IdProxy_Echo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "gid/proxy.proto",
}
