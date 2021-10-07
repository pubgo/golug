// Code generated by protoc-gen-lug. DO NOT EDIT.
// versions:
// - protoc-gen-lug v0.1.0
// - protoc         v3.17.3
// source: proto/gid/a_bit_of_everything.proto

package gid

import (
	fiber "github.com/pubgo/lug/pkg/builder/fiber"
	grpcc "github.com/pubgo/lug/plugins/grpcc"
	xgen "github.com/pubgo/lug/xgen"
	xerror "github.com/pubgo/xerror"
	grpc "google.golang.org/grpc"
	reflect "reflect"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

func GetLoginServiceClient(srv string, opts ...func(cfg *grpcc.Cfg)) func(func(cli LoginServiceClient)) error {
	client := grpcc.GetClient(srv, opts...)
	return func(fn func(cli LoginServiceClient)) (err error) {
		defer xerror.RespErr(&err)

		c, err := client.Get()
		if err != nil {
			return xerror.WrapF(err, "srv: %s", srv)
		}

		fn(&loginServiceClient{c})
		return
	}
}
func init() {
	var mthList []xgen.GrpcRestHandler
	mthList = append(mthList, xgen.GrpcRestHandler{
		Input:        &LoginRequest{},
		Output:       &LoginReply{},
		Service:      "gid.LoginService",
		Name:         "Login",
		Method:       "POST",
		Path:         "/v1/example/login",
		DefaultUrl:   false,
		ClientStream: false,
		ServerStream: false,
	})
	mthList = append(mthList, xgen.GrpcRestHandler{
		Input:        &LogoutRequest{},
		Output:       &LogoutReply{},
		Service:      "gid.LoginService",
		Name:         "Logout",
		Method:       "POST",
		Path:         "/v1/example/logout",
		DefaultUrl:   false,
		ClientStream: false,
		ServerStream: false,
	})
	xgen.Add(reflect.ValueOf(RegisterLoginServiceServer), mthList)
	xgen.Add(reflect.ValueOf(RegisterLoginServiceHandlerServer), nil)
}
func RegisterLoginServiceRestServer(app fiber.Router, server LoginServiceServer) {
	xerror.Assert(app == nil || server == nil, "app is nil or server is nil")
	app.Add("POST", "/v1/example/login", func(ctx *fiber.Ctx) error {
		var req = new(LoginRequest)
		xerror.Panic(ctx.BodyParser(req))
		var resp, err = server.Login(ctx.UserContext(), req)
		xerror.Panic(err)
		return xerror.Wrap(ctx.JSON(resp))
	})
	app.Add("POST", "/v1/example/logout", func(ctx *fiber.Ctx) error {
		var req = new(LogoutRequest)
		xerror.Panic(ctx.BodyParser(req))
		var resp, err = server.Logout(ctx.UserContext(), req)
		xerror.Panic(err)
		return xerror.Wrap(ctx.JSON(resp))
	})
}
