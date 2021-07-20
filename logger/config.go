package logger

import (
	"github.com/pubgo/lug/consts"
	"github.com/pubgo/lug/runenv"
	"github.com/pubgo/lug/vars"

	"github.com/pubgo/xerror"
	"github.com/pubgo/xlog"
	"github.com/pubgo/xlog/xlog_config"
	"go.uber.org/zap"
)

var Name = "logger"

var cfg = xlog_config.NewProdConfig()

func init() {
	if runenv.IsDev() || runenv.IsTest() {
		cfg = xlog_config.NewDevConfig()
		cfg.EncoderConfig.EncodeCaller = consts.Default
		cfg.EncoderConfig.EncodeTime = consts.DefaultTimeFormat
	}

	// 全局log设置
	var log = xerror.PanicErr(cfg.Build()).(*zap.Logger)
	xerror.Panic(xlog.SetDefault(log.Named(runenv.Domain)))

	vars.Watch(Name, func() interface{} { return cfg })
}
