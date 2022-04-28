// Code generated by protoc-gen-lava. DO NOT EDIT.
// versions:
// - protoc-gen-lava v0.1.0
// - protoc         v3.19.4
// source: proto/hello/api.proto

package hello

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

func InitTestApiClient(addr string, alias ...string) {

	var name = ""
	if len(alias) > 0 {
		name = alias[0]
	}
	conn := grpcc_builder.NewClient(addr)

	module.Register(fx.Provide(fx.Annotated{
		Target: func() TestApiClient { return NewTestApiClient(conn) },
		Name:   name,
	}))
}

func RegisterTestApi(srv service.Service, impl TestApiServer) {
	srv.RegService(service.Desc{
		Handler:     impl,
		ServiceDesc: TestApi_ServiceDesc,
	})

	srv.RegGateway(func(ctx context.Context, mux *runtime.ServeMux, cc grpc.ClientConnInterface) error {
		return RegisterTestApiHandlerClient(ctx, mux, NewTestApiClient(cc))
	})

}

func InitTestApiV2Client(addr string, alias ...string) {

	var name = ""
	if len(alias) > 0 {
		name = alias[0]
	}
	conn := grpcc_builder.NewClient(addr)

	module.Register(fx.Provide(fx.Annotated{
		Target: func() TestApiV2Client { return NewTestApiV2Client(conn) },
		Name:   name,
	}))
}

func RegisterTestApiV2(srv service.Service, impl TestApiV2Server) {
	srv.RegService(service.Desc{
		Handler:     impl,
		ServiceDesc: TestApiV2_ServiceDesc,
	})

	srv.RegGateway(func(ctx context.Context, mux *runtime.ServeMux, cc grpc.ClientConnInterface) error {
		return RegisterTestApiV2HandlerClient(ctx, mux, NewTestApiV2Client(cc))
	})

}
