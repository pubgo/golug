package ctl_entry

import (
	"github.com/pubgo/golug/golug_log"
	"github.com/pubgo/xlog"
)

const Name = "ctl_entry"

var log xlog.XLog

func init() {
	golug_log.Watch(func(logs xlog.XLog) {
		log = logs.Named(Name)
	})
}
