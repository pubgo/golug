package grpc_entry

import (
	"github.com/pubgo/golug/golug_data"
	"github.com/pubgo/xerror"
	"reflect"
)

func register(server *entryServerWrapper, handler interface{}) (err error) {
	defer xerror.RespErr(&err)

	if handler == nil {
		return xerror.New("[handler] should not be nil")
	}

	if server == nil {
		return xerror.New("[server] should not be nil")
	}

	var vRegister reflect.Value
	hd := reflect.New(reflect.Indirect(reflect.ValueOf(handler)).Type()).Type()
	for v := range golug_data.List() {
		v, ok := v.(reflect.Value)
		if !ok {
			continue
		}

		v1 := v.Type()
		if v1.Kind() != reflect.Func || v1.NumIn() < 2 {
			continue
		}

		if hd.Implements(v1.In(1)) {
			vRegister = reflect.ValueOf(v)
			break
		}
	}

	if !vRegister.IsValid() || vRegister.IsNil() {
		return xerror.Fmt("[%#v, %#v] 没有找到匹配的interface", handler, vRegister.Interface())
	}

	vRegister.Call([]reflect.Value{reflect.ValueOf(server), reflect.ValueOf(handler)})
	return
}
