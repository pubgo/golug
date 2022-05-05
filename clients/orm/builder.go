package orm

import (
	"time"

	"github.com/pubgo/xerror"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"gorm.io/gorm"
	gl "gorm.io/gorm/logger"
	opentracing "gorm.io/plugin/opentracing"

	"github.com/pubgo/lava/config"
	"github.com/pubgo/lava/core/tracing"
	"github.com/pubgo/lava/inject"
	"github.com/pubgo/lava/logging"
	"github.com/pubgo/lava/logging/logkey"
	"github.com/pubgo/lava/pkg/merge"
	"github.com/pubgo/lava/runtime"
	"github.com/pubgo/lava/vars"
)

func init() {
	defer xerror.RespExit()
	var cfgMap = make(map[string]*Cfg)
	xerror.Panic(config.Decode(Name, &cfgMap))
	for name := range cfgMap {
		cfg := cfgMap[name]
		xerror.Panic(cfg.Valid())
		inject.Register(fx.Provide(fx.Annotated{
			Name: inject.Name(name),
			Target: func(log *logging.Logger) *Client {
				return NewWithCfg(cfg, log)
			},
		}))
	}
}

func NewWithCfg(cfg *Cfg, log *logging.Logger) *Client {
	defer xerror.RespExit()

	var ormCfg = &gorm.Config{}
	xerror.Panic(merge.Struct(ormCfg, cfg))

	var level = gl.Info
	if runtime.IsProd() || runtime.IsRelease() {
		level = gl.Error
	}

	ormCfg.Logger = gl.New(
		logPrintf(zap.L().Named(logkey.Component).Named(Name).Sugar().Infof),
		gl.Config{
			SlowThreshold:             200 * time.Millisecond,
			LogLevel:                  level,
			IgnoreRecordNotFoundError: false,
			Colorful:                  true,
		},
	)

	var factory = Get(cfg.Driver)
	xerror.Assert(factory == nil, "factory[%s] not found", cfg.Driver)
	dialect := factory(cfg.DriverCfg)

	db, err := gorm.Open(dialect, ormCfg)
	xerror.Panic(err)

	// 添加链路追踪
	xerror.Panic(db.Use(opentracing.New(
		opentracing.WithErrorTagHook(tracing.SetIfErr),
	)))

	// 服务连接校验
	sqlDB, err := db.DB()
	xerror.Panic(err)
	xerror.Panic(sqlDB.Ping())

	if cfg.MaxConnTime != 0 {
		sqlDB.SetConnMaxLifetime(cfg.MaxConnTime)
	}

	if cfg.MaxConnIdle != 0 {
		sqlDB.SetMaxIdleConns(cfg.MaxConnIdle)
	}

	if cfg.MaxConnOpen != 0 {
		sqlDB.SetMaxOpenConns(cfg.MaxConnOpen)
	}

	var cli = &Client{DB: db}
	vars.Register(Name+"_stats", func() interface{} {
		var data = make(map[string]interface{})
		_db, err := cli.DB.DB()
		if err != nil {
			data["data"] = err.Error()
		} else {
			data["data"] = _db.Stats()
		}
		return data
	})
	return cli
}
