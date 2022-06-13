package klog

import (
	"github.com/go-logr/zapr"
	"github.com/pubgo/dix"
	"k8s.io/klog/v2"

	"github.com/pubgo/lava/logging"
)

// 替换klog全局log
func init() {
	dix.Register(func(logger *logging.Logger) *logging.ExtLog {
		klog.SetLogger(zapr.NewLogger(logging.Component("klog").L()))
		return new(logging.ExtLog)
	})
}
