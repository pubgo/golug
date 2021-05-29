package main

import (
	"fmt"
	"log"

	"github.com/pubgo/xerror"
	"github.com/pubgo/xprotogen/gen"
)

func main() {
	defer xerror.RespDebug()

	m := gen.New("lug",gen.OnlyService())
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

	"github.com/pubgo/xerror"
	"google.golang.org/grpc"
	"github.com/pubgo/lug/xgen"
	"github.com/pubgo/lug/client/grpcc"
)`
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
			var tpl = ""
			gen.Append(&tpl, `func init() {`)
			gen.Append(&tpl, `var mthList []xgen.GrpcRestHandler`)
			for _, ss := range fd.GetService() {
				for _, m := range ss.GetMethod() {
					gen.Append(&tpl, gen.Template(`
					mthList = append(mthList, xgen.GrpcRestHandler{
						Service:      "{{pkg}}.{{ss.Name}}",
						Name:         "{{m.GetName()}}",
						Method:       "{{m.HttpMethod}}",
						Path:          "{{m.HttpPath}}",
						ClientStream:  "{{m.CS}}"=="True",
						ServerStreams: "{{m.SS}}"=="True",
					})`, gen.Context{"pkg": fd.Pkg, "m": m, "ss": ss}))
				}
				gen.Append(&tpl, fmt.Sprintf(`xgen.Add(reflect.ValueOf(Register%sServer),mthList)`, ss.Srv))

				var isStream bool
				for _, m := range ss.GetMethod() {
					if m.CS || m.SS {
						isStream = true
						break
					}
				}

				if !isStream {
					gen.Append(&tpl, fmt.Sprintf(`xgen.Add(reflect.ValueOf(Register%sHandlerFromEndpoint), nil)`, ss.Srv))
				}
			}

			gen.Append(&tpl, `}`)
			return tpl
		},
	))
}
