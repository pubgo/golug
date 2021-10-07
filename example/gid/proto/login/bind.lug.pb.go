// Code generated by protoc-gen-lug. DO NOT EDIT.
// versions:
// - protoc-gen-lug v0.1.0
// - protoc         v3.17.3
// source: proto/login/bind.proto

package login

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

func GetBindTelephoneClient(srv string, opts ...func(cfg *grpcc.Cfg)) func(func(cli BindTelephoneClient)) error {
	client := grpcc.GetClient(srv, opts...)
	return func(fn func(cli BindTelephoneClient)) (err error) {
		defer xerror.RespErr(&err)

		c, err := client.Get()
		if err != nil {
			return xerror.WrapF(err, "srv: %s", srv)
		}

		fn(&bindTelephoneClient{c})
		return
	}
}
func init() {
	var mthList []xgen.GrpcRestHandler
	mthList = append(mthList, xgen.GrpcRestHandler{
		Input:        &CheckRequest{},
		Output:       &CheckResponse{},
		Service:      "login.BindTelephone",
		Name:         "Check",
		Method:       "POST",
		Path:         "/user/bind-telephone/check",
		DefaultUrl:   false,
		ClientStream: false,
		ServerStream: false,
	})
	mthList = append(mthList, xgen.GrpcRestHandler{
		Input:        &BindVerifyRequest{},
		Output:       &BindVerifyResponse{},
		Service:      "login.BindTelephone",
		Name:         "BindVerify",
		Method:       "POST",
		Path:         "/user/bind-telephone/bind-verify",
		DefaultUrl:   false,
		ClientStream: false,
		ServerStream: false,
	})
	mthList = append(mthList, xgen.GrpcRestHandler{
		Input:        &BindChangeRequest{},
		Output:       &BindChangeResponse{},
		Service:      "login.BindTelephone",
		Name:         "BindChange",
		Method:       "POST",
		Path:         "/user/bind-telephone/bind-change",
		DefaultUrl:   false,
		ClientStream: false,
		ServerStream: false,
	})
	mthList = append(mthList, xgen.GrpcRestHandler{
		Input:        &AutomaticBindRequest{},
		Output:       &AutomaticBindResponse{},
		Service:      "login.BindTelephone",
		Name:         "AutomaticBind",
		Method:       "POST",
		Path:         "/user/bind-telephone/automatic-bind",
		DefaultUrl:   false,
		ClientStream: false,
		ServerStream: false,
	})
	mthList = append(mthList, xgen.GrpcRestHandler{
		Input:        &BindPhoneParseRequest{},
		Output:       &BindPhoneParseResponse{},
		Service:      "login.BindTelephone",
		Name:         "BindPhoneParse",
		Method:       "POST",
		Path:         "/user/bind-telephone/bind-phone-parse",
		DefaultUrl:   false,
		ClientStream: false,
		ServerStream: false,
	})
	mthList = append(mthList, xgen.GrpcRestHandler{
		Input:        &BindPhoneParseByOneClickRequest{},
		Output:       &BindPhoneParseByOneClickResponse{},
		Service:      "login.BindTelephone",
		Name:         "BindPhoneParseByOneClick",
		Method:       "POST",
		Path:         "/user/bind-telephone/bind-phone-parse-by-one-click",
		DefaultUrl:   false,
		ClientStream: false,
		ServerStream: false,
	})
	xgen.Add(reflect.ValueOf(RegisterBindTelephoneServer), mthList)
	xgen.Add(reflect.ValueOf(RegisterBindTelephoneHandlerServer), nil)
}
func RegisterBindTelephoneRestServer(app fiber.Router, server BindTelephoneServer) {
	xerror.Assert(app == nil || server == nil, "app is nil or server is nil")
	app.Add("POST", "/user/bind-telephone/check", func(ctx *fiber.Ctx) error {
		var req = new(CheckRequest)
		xerror.Panic(ctx.BodyParser(req))
		var resp, err = server.Check(ctx.UserContext(), req)
		xerror.Panic(err)
		return xerror.Wrap(ctx.JSON(resp))
	})
	app.Add("POST", "/user/bind-telephone/bind-verify", func(ctx *fiber.Ctx) error {
		var req = new(BindVerifyRequest)
		xerror.Panic(ctx.BodyParser(req))
		var resp, err = server.BindVerify(ctx.UserContext(), req)
		xerror.Panic(err)
		return xerror.Wrap(ctx.JSON(resp))
	})
	app.Add("POST", "/user/bind-telephone/bind-change", func(ctx *fiber.Ctx) error {
		var req = new(BindChangeRequest)
		xerror.Panic(ctx.BodyParser(req))
		var resp, err = server.BindChange(ctx.UserContext(), req)
		xerror.Panic(err)
		return xerror.Wrap(ctx.JSON(resp))
	})
	app.Add("POST", "/user/bind-telephone/automatic-bind", func(ctx *fiber.Ctx) error {
		var req = new(AutomaticBindRequest)
		xerror.Panic(ctx.BodyParser(req))
		var resp, err = server.AutomaticBind(ctx.UserContext(), req)
		xerror.Panic(err)
		return xerror.Wrap(ctx.JSON(resp))
	})
	app.Add("POST", "/user/bind-telephone/bind-phone-parse", func(ctx *fiber.Ctx) error {
		var req = new(BindPhoneParseRequest)
		xerror.Panic(ctx.BodyParser(req))
		var resp, err = server.BindPhoneParse(ctx.UserContext(), req)
		xerror.Panic(err)
		return xerror.Wrap(ctx.JSON(resp))
	})
	app.Add("POST", "/user/bind-telephone/bind-phone-parse-by-one-click", func(ctx *fiber.Ctx) error {
		var req = new(BindPhoneParseByOneClickRequest)
		xerror.Panic(ctx.BodyParser(req))
		var resp, err = server.BindPhoneParseByOneClick(ctx.UserContext(), req)
		xerror.Panic(err)
		return xerror.Wrap(ctx.JSON(resp))
	})
}
