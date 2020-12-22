// Code generated by protoc-gen-golug. DO NOT EDIT.
// source: example/proto/login/login.proto

package login

import (
	"reflect"

	"github.com/pubgo/golug/golug_client/grpclient"
	"github.com/pubgo/golug/golug_xgen"
)

func init() {
	var mthList []golug_xgen.GrpcRestHandler
	mthList = append(mthList, golug_xgen.GrpcRestHandler{
		Name:          "Login",
		Method:        "POST",
		Path:          "/user/login/login",
		ClientStream:  "False" == "True",
		ServerStreams: "False" == "True",
	})

	mthList = append(mthList, golug_xgen.GrpcRestHandler{
		Name:          "Authenticate",
		Method:        "POST",
		Path:          "/user/login/authenticate",
		ClientStream:  "False" == "True",
		ServerStreams: "False" == "True",
	})

	golug_xgen.Add(reflect.ValueOf(RegisterLoginServer), mthList)
}

func GetLoginClient(srv grpclient.Client) LoginClient {
	return &loginClient{grpclient.GetClient(srv.Name())}
}
