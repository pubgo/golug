// Code generated by protoc-gen-lug. DO NOT EDIT.
// source: example/proto/hello/helloworld.proto

package hello

import (
	"reflect"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
	fb "github.com/pubgo/lug/builder/fiber"
	"github.com/pubgo/lug/pkg/gutil"
	"github.com/pubgo/lug/plugins/grpcc"
	"github.com/pubgo/lug/xgen"
	"github.com/pubgo/xerror"
	"google.golang.org/grpc"
)

var _ = strings.Trim
var _ = utils.UnsafeString
var _ fiber.Router = nil
var _ = gutil.MapFormByTag
var _ = fb.Cfg{}

func GetGreeterClient(srv string, optFns ...func(service string) []grpc.DialOption) func() (GreeterClient, error) {
	client := grpcc.GetClient(srv, optFns...)
	return func() (GreeterClient, error) {
		c, err := client.Get()
		return &greeterClient{c}, xerror.WrapF(err, "srv: %s", srv)
	}
}

func init() {
	var mthList []xgen.GrpcRestHandler

	mthList = append(mthList, xgen.GrpcRestHandler{
		Service:      "hello.Greeter",
		Name:         "SayHello",
		Method:       "GET",
		Path:         "/say/{name}",
		ClientStream: "False" == "True",
		ServerStream: "False" == "True",
		DefaultUrl:   "False" == "True",
	})

	xgen.Add(reflect.ValueOf(RegisterGreeterServer), mthList)
	xgen.Add(reflect.ValueOf(RegisterGreeterRestServer), nil)
	xgen.Add(reflect.ValueOf(RegisterGreeterHandler), nil)
}

func RegisterGreeterRestServer(app fiber.Router, server GreeterServer) {
	xerror.Assert(app == nil || server == nil, "app is nil or server is nil")

	// restful
	app.Add("GET", "/say/{name}", func(ctx *fiber.Ctx) error {
		var req = new(HelloRequest)
		data := make(map[string][]string)
		ctx.Context().QueryArgs().VisitAll(func(key []byte, val []byte) {
			k := utils.UnsafeString(key)
			v := utils.UnsafeString(val)
			if strings.Contains(v, ",") && gutil.EqualFieldType(req, reflect.Slice, k) {
				values := strings.Split(v, ",")
				for i := 0; i < len(values); i++ {
					data[k] = append(data[k], values[i])
				}
			} else {
				data[k] = append(data[k], v)
			}
		})
		xerror.Panic(gutil.MapFormByTag(req, data, "json"))
		var resp, err = server.SayHello(ctx.UserContext(), req)
		if err != nil {
			return xerror.Wrap(err)
		}

		return xerror.Wrap(ctx.JSON(resp))
	})

}
