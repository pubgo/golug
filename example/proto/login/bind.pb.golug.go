// Code generated by protoc-gen-golug. DO NOT EDIT.
// source: example/proto/login/bind.proto

package login

import (
	"reflect"

	"github.com/gofiber/fiber/v2"
	"github.com/pubgo/golug/client/grpclient"
	"github.com/pubgo/golug/golug_xgen"
	"github.com/pubgo/golug/pkg/golug_utils"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

var _ = golug_utils.Decode

func init() {
	var mthList []golug_xgen.GrpcRestHandler
	mthList = append(mthList, golug_xgen.GrpcRestHandler{
		Service:       "login.BindTelephone",
		Name:          "Check",
		Method:        "POST",
		Path:          "/user/bind-telephone/check",
		ClientStream:  "False" == "True",
		ServerStreams: "False" == "True",
	})

	mthList = append(mthList, golug_xgen.GrpcRestHandler{
		Service:       "login.BindTelephone",
		Name:          "BindVerify",
		Method:        "POST",
		Path:          "/user/bind-telephone/bind-verify",
		ClientStream:  "False" == "True",
		ServerStreams: "False" == "True",
	})

	mthList = append(mthList, golug_xgen.GrpcRestHandler{
		Service:       "login.BindTelephone",
		Name:          "BindChange",
		Method:        "POST",
		Path:          "/user/bind-telephone/bind-change",
		ClientStream:  "False" == "True",
		ServerStreams: "False" == "True",
	})

	mthList = append(mthList, golug_xgen.GrpcRestHandler{
		Service:       "login.BindTelephone",
		Name:          "AutomaticBind",
		Method:        "POST",
		Path:          "/user/bind-telephone/automatic-bind",
		ClientStream:  "False" == "True",
		ServerStreams: "False" == "True",
	})

	mthList = append(mthList, golug_xgen.GrpcRestHandler{
		Service:       "login.BindTelephone",
		Name:          "BindPhoneParse",
		Method:        "POST",
		Path:          "/user/bind-telephone/bind-phone-parse",
		ClientStream:  "False" == "True",
		ServerStreams: "False" == "True",
	})

	mthList = append(mthList, golug_xgen.GrpcRestHandler{
		Service:       "login.BindTelephone",
		Name:          "BindPhoneParseByOneClick",
		Method:        "POST",
		Path:          "/user/bind-telephone/bind-phone-parse-by-one-click",
		ClientStream:  "False" == "True",
		ServerStreams: "False" == "True",
	})

	golug_xgen.Add(reflect.ValueOf(RegisterBindTelephoneServer), mthList)
	golug_xgen.Add(reflect.ValueOf(RegisterBindTelephoneGateway), nil)
}

func GetBindTelephoneClient(srv string, opts ...grpc.DialOption) (BindTelephoneClient, error) {
	c, err := grpclient.Get(srv, opts...)
	return &bindTelephoneClient{c}, err
}

func RegisterBindTelephoneGateway(srv string, g fiber.Router, opts ...grpc.DialOption) error {
	c, err := GetBindTelephoneClient(srv, opts...)
	if err != nil {
		return err
	}
	g.Add("POST", "/user/bind-telephone/check", func(ctx *fiber.Ctx) error {
		p := metadata.Pairs()
		ctx.Request().Header.VisitAll(func(key, value []byte) { p.Set(string(key), string(value)) })

		var req CheckRequest
		if err := ctx.BodyParser(&req); err != nil {
			return err
		}

		resp, err := c.Check(metadata.NewIncomingContext(ctx.Context(), p), &req)
		if err != nil {
			return err
		}
		return ctx.JSON(resp)
	})

	g.Add("POST", "/user/bind-telephone/bind-verify", func(ctx *fiber.Ctx) error {
		p := metadata.Pairs()
		ctx.Request().Header.VisitAll(func(key, value []byte) { p.Set(string(key), string(value)) })

		var req BindVerifyRequest
		if err := ctx.BodyParser(&req); err != nil {
			return err
		}

		resp, err := c.BindVerify(metadata.NewIncomingContext(ctx.Context(), p), &req)
		if err != nil {
			return err
		}
		return ctx.JSON(resp)
	})

	g.Add("POST", "/user/bind-telephone/bind-change", func(ctx *fiber.Ctx) error {
		p := metadata.Pairs()
		ctx.Request().Header.VisitAll(func(key, value []byte) { p.Set(string(key), string(value)) })

		var req BindChangeRequest
		if err := ctx.BodyParser(&req); err != nil {
			return err
		}

		resp, err := c.BindChange(metadata.NewIncomingContext(ctx.Context(), p), &req)
		if err != nil {
			return err
		}
		return ctx.JSON(resp)
	})

	g.Add("POST", "/user/bind-telephone/automatic-bind", func(ctx *fiber.Ctx) error {
		p := metadata.Pairs()
		ctx.Request().Header.VisitAll(func(key, value []byte) { p.Set(string(key), string(value)) })

		var req AutomaticBindRequest
		if err := ctx.BodyParser(&req); err != nil {
			return err
		}

		resp, err := c.AutomaticBind(metadata.NewIncomingContext(ctx.Context(), p), &req)
		if err != nil {
			return err
		}
		return ctx.JSON(resp)
	})

	g.Add("POST", "/user/bind-telephone/bind-phone-parse", func(ctx *fiber.Ctx) error {
		p := metadata.Pairs()
		ctx.Request().Header.VisitAll(func(key, value []byte) { p.Set(string(key), string(value)) })

		var req BindPhoneParseRequest
		if err := ctx.BodyParser(&req); err != nil {
			return err
		}

		resp, err := c.BindPhoneParse(metadata.NewIncomingContext(ctx.Context(), p), &req)
		if err != nil {
			return err
		}
		return ctx.JSON(resp)
	})

	g.Add("POST", "/user/bind-telephone/bind-phone-parse-by-one-click", func(ctx *fiber.Ctx) error {
		p := metadata.Pairs()
		ctx.Request().Header.VisitAll(func(key, value []byte) { p.Set(string(key), string(value)) })

		var req BindPhoneParseByOneClickRequest
		if err := ctx.BodyParser(&req); err != nil {
			return err
		}

		resp, err := c.BindPhoneParseByOneClick(metadata.NewIncomingContext(ctx.Context(), p), &req)
		if err != nil {
			return err
		}
		return ctx.JSON(resp)
	})

	return nil
}
