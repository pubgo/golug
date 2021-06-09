package main

import (
	"github.com/pubgo/xerror"
	"github.com/pubgo/xprotogen/gen"

	"log"
)

func main() {
	defer xerror.RespDebug()

	m := gen.New("lug", gen.OnlyService())
	m.Parameter(func(key, value string) {
		log.Println("params:", key, "=", value)
	})

	xerror.Panic(m.GenWithTpl(
		func(fd *gen.FileDescriptor) string {
			return `
// Code generated by protoc-gen-lug. DO NOT EDIT.
{%- if !fd.GetOptions().GetDeprecated() %}
// source: {{fd.GetName()}}
{%- else %}
// {{fd.GetName()}} is a deprecated file.
{%- endif %}

package {{pkg}}
import (
	"reflect"
	"strings"

	"github.com/pubgo/xerror"
	"google.golang.org/grpc"
	"github.com/pubgo/lug/xgen"
	"github.com/pubgo/lug/client/grpcc"
	"github.com/gofiber/fiber/v2"
	"github.com/pubgo/lug/pkg/gutil"
	"github.com/gofiber/fiber/v2/utils"
	fb "github.com/pubgo/lug/builder/fiber"
)

var _ = strings.Trim
var _ = gutil.EqualFieldType
var _ = utils.ByteSize
`
		},

		func(fd *gen.FileDescriptor) string {
			return `
{% for ss in fd.GetService() %}
	func Get{{ss.Srv}}Client(srv string, optFns ...func(service string) []grpc.DialOption) func() ({{ss.Srv}}Client,error) {
		client := grpcc.GetClient(srv, optFns...)
		return func() ({{ss.Srv}}Client,error) {
			c, err := client.Get()
			return &{{unExport(ss.Srv)}}Client{c},xerror.WrapF(err, "srv: %s", srv)
		}
	}
{% endfor %}
`
		},

		func(fd *gen.FileDescriptor) string {
			return `
{% for ss in fd.GetService() %}
	func init(){
		var mthList []xgen.GrpcRestHandler
		{% for m in ss.GetMethod() %}
			mthList = append(mthList, xgen.GrpcRestHandler{
				Service:      "{{pkg}}.{{ss.Name}}",
				Name:         "{{m.GetName()}}",
				Method:       "{{m.HttpMethod}}",
				Path:          "{{m.HttpPath}}",
				ClientStream:  "{{m.CS}}"=="True",
				ServerStreams: "{{m.SS}}"=="True",
			})
		{% endfor %}
		xgen.Add(reflect.ValueOf(Register{{ss.Srv}}Server),mthList)
	}
{% endfor %}
`
		},

		func(fd *gen.FileDescriptor) string {
			return `
{% for ss in fd.GetService() %}
	func Register{{ss.Srv}}RestServer(app fiber.Router, server {{ss.Srv}}Server) {
		if app == nil || server == nil {
			panic("app is nil or server is nil")
		}

		{% for m in ss.GetMethod() %}
			{%- if !m.CS && !m.SS %}
			app.Add("{{m.HttpMethod}}","{{m.HttpPath}}", func(ctx *fiber.Ctx) error {
				var req = new({{m.GetInputType()}})				
				{%- if m.HttpMethod=="GET" %}
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
				{%- else %}
					xerror.Panic(ctx.BodyParser(req))
				{%- endif %}

				var resp,err=server.{{m.GetName()}}(ctx.Context(),req)
				if err!=nil{
					return err
				}

				return ctx.JSON(resp)
			})
			{%- else %}

			app.Get("{{m.HttpPath}}", fb.NewWs(func(ctx *fiber.Ctx,c *fb.Conn) {
			defer c.Close()
		
			{%- if m.CS %}
				if err := server.{{m.GetName()}}(&{{unExport(ss.Srv)}}{{m.GetName()}}Server{fb.NewWsStream(ctx,c)}); err != nil {
					c.WriteMessage(fb.TextMessage, []byte(err.Error()))
				}
			{%- else %}
				var req = new({{m.GetInputType()}})				
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
				if err := server.{{m.GetName()}}(req,&{{unExport(ss.Srv)}}{{m.GetName()}}Server{fb.NewWsStream(ctx,c)}); err != nil {
					c.WriteMessage(fb.TextMessage, []byte(err.Error()))
				}
			{%- endif %}
			}))
			{%- endif %}
		{% endfor %}
	}
{% endfor %}
`
		},
	))
}
