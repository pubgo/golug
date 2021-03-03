package watcher

import (
	"github.com/pubgo/golug/types"
	"github.com/pubgo/xerror"
)

var watchers types.SMap

func List() (dt map[string]Watcher) { xerror.Panic(watchers.Map(&dt)); return dt }
func Register(name string, w Watcher) {
	xerror.Assert(name == "" || w == nil, "[watcher:%s] should not be null", name)
	xerror.Assert(watchers.Has(name), "[watcher:%s] already exists", name)

	watchers.Set(name, w)
}
