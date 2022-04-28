// Code generated by protoc-gen-lava. DO NOT EDIT.
// versions:
// - protoc-gen-lava v0.1.0
// - protoc         v3.19.4
// source: proto/gid/echo_service.proto

package gid

import (
	context "context"
	runtime "github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	grpcc_builder "github.com/pubgo/lava/clients/grpcc/grpcc_builder"
	module "github.com/pubgo/lava/module"
	service "github.com/pubgo/lava/service"
	fx "go.uber.org/fx"
	grpc "google.golang.org/grpc"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

func InitEchoServiceClient(addr string, alias ...string) {

	var name = ""
	if len(alias) > 0 {
		name = alias[0]
	}
	conn := grpcc_builder.NewClient(addr)

	module.Register(fx.Provide(fx.Annotated{
		Target: func() EchoServiceClient { return NewEchoServiceClient(conn) },
		Name:   name,
	}))
}

func RegisterEchoService(srv service.Service, impl EchoServiceServer) {
	srv.RegService(service.Desc{
		Handler:     impl,
		ServiceDesc: EchoService_ServiceDesc,
	})

	srv.RegGateway(func(ctx context.Context, mux *runtime.ServeMux, cc grpc.ClientConnInterface) error {
		return RegisterEchoServiceHandlerClient(ctx, mux, NewEchoServiceClient(cc))
	})

}
