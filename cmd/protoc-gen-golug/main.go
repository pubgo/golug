package main

import (
	"github.com/pubgo/golug/golug_entry"
	"github.com/pubgo/golug/golug_util"
	"github.com/pubgo/xerror"
	"github.com/pubgo/xprotogen/gen"
	"log"
)

func main() {
	m := gen.New("golug")
	m.Parameter(func(key, value string) {
		log.Println("params:", key, "=", value)
	})

	xerror.Exit(m.Gen(func(fd *gen.FileDescriptor) {
		fd.Set("fdName", fd.GetName())

		j := fd.Jen
		j.PackageComment("// Code generated by protoc-gen-golug. DO NOT EDIT.")
		if !fd.GetOptions().GetDeprecated() {
			j.PackageComment("// source: " + fd.GetName())
		} else {
			j.PackageComment("// " + fd.GetName() + " is a deprecated file.")
		}

		j.Id(`
import (
	"reflect"

	"github.com/pubgo/golug/golug_data"
)
`)

		for _, ss := range fd.GetService() {
			var mths []golug_entry.GrpcRestHandler
			for _, m := range ss.GetMethod() {
				mths = append(mths, golug_entry.GrpcRestHandler{
					Name:          m.GetName(),
					Method:        m.P("{{.http_method}}"),
					Path:          m.P("{{.http_path}}"),
					ClientStream:  m.P("{{.cs}}") == "true",
					ServerStreams: m.P("{{.ss}}") == "true",
				})
			}

			ss.Set("data", "`"+golug_util.Marshal(mths)+"`")
			j.Id(ss.P(`func init() {golug_data.Add(reflect.ValueOf(Register{{.srv}}Server),{{.data}})}`))
		}
	}))
}
