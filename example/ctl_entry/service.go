package ctl_entry

import (
	"fmt"
	db2 "github.com/pubgo/lug/plugins/db"
	"time"

	"github.com/pubgo/lug/consts"
	"github.com/pubgo/lug/entry/ctl"

	"github.com/pubgo/x/fx"
	"go.uber.org/zap"
)

var _ ctl.Service = (*Service)(nil)

type Service struct {
	Db *db2.Client `dix:""`
}

func (t *Service) Run() map[string]ctl.Handler {
	return map[string]ctl.Handler{
		consts.Default: func(ctx fx.Ctx) {
			fmt.Println("db ping:", t.Db.Get().Ping())
			zap.L().Info("ctl hello once")
		},
	}
}

func (t *Service) RunLoop() map[string]ctl.Handler {
	return map[string]ctl.Handler{
		"hello": func(ctx fx.Ctx) {
			fmt.Println("db ping:", t.Db.Get().Ping())
			zap.L().Info("ctl hello forever")
			time.Sleep(time.Second)
		},
	}
}
