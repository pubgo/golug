package log

import (
	"github.com/pubgo/lug/runenv"
	"github.com/pubgo/xerror"
	"github.com/pubgo/xlog"
	"github.com/pubgo/xlog/xlog_config"
	"go.uber.org/zap"
)

// 默认logger初始化
func init() {
	cfg := xlog_config.NewDevConfig()
	cfg.EncoderConfig.EncodeCaller = "full"
	log := xlog.New(xerror.PanicErr(cfg.Build()).(*zap.Logger))

	// 全局log设置
	xerror.Panic(xlog.SetDefault(log.Named(runenv.Domain)))
}

func initLog(cfg xlog_config.Config) (err error) {
	defer xerror.RespErr(&err)

	log := xlog.New(xerror.PanicErr(cfg.Build()).(*zap.Logger))

	// 全局log设置
	xerror.Panic(xlog.SetDefault(log.Named(runenv.Domain)))

	return nil
}
