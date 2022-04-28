package service_builder

import (
	"context"
	"errors"
	"fmt"
	"github.com/pubgo/lava/module"
	"go.uber.org/fx"
	"net"
	"net/http"
	"plugin"

	"github.com/fullstorydev/grpchan/inprocgrpc"
	fiber2 "github.com/gofiber/fiber/v2"
	gw "github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/pubgo/x/stack"
	"github.com/pubgo/xerror"
	"github.com/urfave/cli/v2"
	"google.golang.org/grpc"

	"github.com/pubgo/lava/abc"
	"github.com/pubgo/lava/config"
	"github.com/pubgo/lava/core/cmux"
	"github.com/pubgo/lava/core/flags"
	"github.com/pubgo/lava/internal/envs"
	"github.com/pubgo/lava/logging/logutil"
	"github.com/pubgo/lava/middleware"
	"github.com/pubgo/lava/pkg/fiber_builder"
	"github.com/pubgo/lava/pkg/grpc_builder"
	"github.com/pubgo/lava/pkg/netutil"
	"github.com/pubgo/lava/pkg/syncx"
	"github.com/pubgo/lava/plugins/signal"
	"github.com/pubgo/lava/runtime"
	"github.com/pubgo/lava/service"
	"github.com/pubgo/lava/version"
)

func New(name string, desc string, plugins ...plugin.Plugin) service.Service {
	return newService(name, desc, plugins...)
}

func newService(name string, desc string, plugins ...plugin.Plugin) *serviceImpl {
	var g = &serviceImpl{
		ctx:        context.Background(),
		pluginList: plugins,
		cmd: &cli.Command{
			Name:  name,
			Usage: desc,
			Flags: flags.GetFlags(),
		},
		srv:     grpc_builder.New(),
		gw:      fiber_builder.New(),
		inproc:  &inprocgrpc.Channel{},
		httpSrv: fiber2.New(),
		net:     cmux.DefaultCfg(),
		cfg: Cfg{
			name:     name,
			hostname: runtime.Hostname,
			id:       runtime.AppID,
			Grpc:     grpc_builder.GetDefaultCfg(),
			Gw:       fiber_builder.Cfg{},
		},
	}

	g.cmd.Action = func(ctx *cli.Context) error {
		defer xerror.RespExit()

		// 项目名初始化
		runtime.Project = name
		envs.SetName(version.Domain, runtime.Project)

		// 运行环境检查
		if _, ok := runtime.RunModeValue[runtime.Mode.String()]; !ok {
			panic(fmt.Sprintf("mode(%s) not match in (%v)", runtime.Mode, runtime.RunModeValue))
		}

		xerror.Panic(g.init())
		xerror.Panic(g.start())
		signal.Block()
		xerror.Panic(g.stop())
		return nil
	}

	return g
}

var _ service.Service = (*serviceImpl)(nil)

type serviceImpl struct {
	beforeStarts []func()
	afterStarts  []func()
	beforeStops  []func()
	afterStops   []func()
	pluginList   []plugin.Plugin
	middlewares  []middleware.Middleware
	services     []service.Desc

	app *fx.App
	cmd *cli.Command

	net *cmux.Mux

	cfg     Cfg
	srv     grpc_builder.Builder
	gw      fiber_builder.Builder
	httpSrv *fiber2.App
	opts    []fx.Option

	// inproc Channel is used to serve grpc gateway
	inproc *inprocgrpc.Channel

	wrapperUnary  middleware.HandlerFunc
	wrapperStream middleware.HandlerFunc

	ctx        context.Context
	gwHandlers []func(ctx context.Context, mux *gw.ServeMux, cc grpc.ClientConnInterface) error
}

func (t *serviceImpl) Provide(constructors ...interface{}) {
	t.opts = append(t.opts, fx.Provide(constructors...))
}

func (t *serviceImpl) Invoke(funcs ...interface{}) {
	t.opts = append(t.opts, fx.Invoke(funcs...))
}

func (t *serviceImpl) Start() error          { return t.start() }
func (t *serviceImpl) Stop() error           { return t.stop() }
func (t *serviceImpl) Command() *cli.Command { return t.cmd }

func (t *serviceImpl) RegService(desc service.Desc) {
	xerror.Assert(desc.Handler == nil, "[handler] is nil")

	t.srv.RegisterService(&desc.ServiceDesc, desc.Handler)
	t.inproc.RegisterService(&desc.ServiceDesc, desc.Handler)
	t.services = append(t.services, desc)

	if h, ok := desc.Handler.(abc.Flags); ok {
		t.Flags(h.Flags()...)
	}

	t.opts = append(t.opts, fx.Populate(desc.Handler))
}

func (t *serviceImpl) RegRouter(prefix string, fn func(r fiber2.Router)) {
	t.httpSrv.Route(prefix, fn)
}

func (t *serviceImpl) RegGateway(fn func(ctx context.Context, mux *gw.ServeMux, cc grpc.ClientConnInterface) error) {
	t.gwHandlers = append(t.gwHandlers, fn)
}

func (t *serviceImpl) RegApp(prefix string, r *fiber2.App) {
	t.httpSrv.Mount(prefix, r)
}

func (t *serviceImpl) Middleware(mid middleware.Middleware) {
	xerror.Assert(mid == nil, "[mid] is nil")
	t.middlewares = append(t.middlewares, mid)
}

func (t *serviceImpl) BeforeStarts(f ...func()) { t.beforeStarts = append(t.beforeStarts, f...) }
func (t *serviceImpl) BeforeStops(f ...func())  { t.beforeStops = append(t.beforeStops, f...) }
func (t *serviceImpl) AfterStarts(f ...func())  { t.afterStarts = append(t.afterStarts, f...) }
func (t *serviceImpl) AfterStops(f ...func())   { t.afterStops = append(t.afterStops, f...) }

func (t *serviceImpl) init() error {
	defer xerror.RespExit()

	t.app = fx.New(append(module.List(), t.opts...)...)
	t.app.Run()

	t.net.Addr = runtime.Addr

	// 配置解析
	xerror.Panic(config.Decode(Name, &t.cfg))

	// 网关初始化
	xerror.Panic(t.gw.Build(t.cfg.Gw))
	t.gw.Get().Use(t.handlerHttpMiddle(t.middlewares))
	t.gw.Get().Mount("/", t.httpSrv)

	// 注册系统middleware
	t.srv.UnaryInterceptor(t.handlerUnaryMiddle(t.middlewares))
	t.srv.StreamInterceptor(t.handlerStreamMiddle(t.middlewares))

	// grpc serve初始化
	xerror.Panic(t.srv.Build(t.cfg.Grpc))

	// 加载inproc的middleware
	t.inproc.WithServerUnaryInterceptor(t.handlerUnaryMiddle(t.middlewares))
	t.inproc.WithServerStreamInterceptor(t.handlerStreamMiddle(t.middlewares))

	// 初始化 handlers
	for _, desc := range t.services() {
		// service handler依赖对象注入
		if h, ok := desc.Handler.(service.Handler); ok {
			t.AfterStops(h.Close)
			logutil.LogOrPanic(t.L, "Service Handler Init", func() error {
				return xerror.Try(func() {
					// register router
					h.Router(t.gw.Get())
					// service handler init
					h.Init()
				})
			})
		}
	}
	return nil
}

func (t *serviceImpl) Flags(flags ...cli.Flag) {
	if len(flags) == 0 {
		return
	}

	t.cmd.Flags = append(t.cmd.Flags, flags...)
}

func (t *serviceImpl) Options() service.Options {
	return service.Options{
		Name:      t.cfg.name,
		Id:        t.cfg.id,
		Version:   version.Version,
		Port:      netutil.MustGetPort(t.net.Addr),
		Address:   t.net.Addr,
		Advertise: t.cfg.Advertise,
	}
}

func (t *serviceImpl) start() (gErr error) {
	defer xerror.RespErr(&gErr)

	logutil.OkOrPanic(t.L, "service before-start", func() error {
		var beforeList []func()
		for _, p := range plugin.All() {
			beforeList = append(beforeList, p.BeforeStarts()...)
		}
		beforeList = append(beforeList, t.beforeStarts...)
		for i := range beforeList {
			t.L.Sugar().Infof("running %s", stack.Func(beforeList[i]))
			xerror.PanicF(xerror.Try(beforeList[i]), stack.Func(beforeList[i]))
		}
		return nil
	})

	var grpcLn = t.net.HTTP2()
	var gwLn = t.net.HTTP1Fast()

	// 启动grpc网关
	syncx.GoDelay(func() {
		t.L.Info("[grpc-gw] Server Starting")
		logutil.LogOrErr(t.L, "[grpc-gw] Server Stop", func() error {
			if err := t.gw.Get().Listener(<-gwLn); err != nil &&
				!errors.Is(err, cmux.ErrListenerClosed) &&
				!errors.Is(err, http.ErrServerClosed) &&
				!errors.Is(err, net.ErrClosed) {
				return err
			}
			return nil
		})
	})

	logutil.OkOrPanic(t.L, "service start", func() error {
		t.L.Sugar().Infof("Server Listening on http://%s:%d", netutil.GetLocalIP(), netutil.MustGetPort(runtime.Addr))

		// 启动grpc服务
		syncx.GoDelay(func() {
			t.L.Info("[grpc] Server Starting")
			logutil.LogOrErr(t.L, "[grpc] Server Stop", func() error {
				if err := t.srv.Get().Serve(<-grpcLn); err != nil &&
					err != cmux.ErrListenerClosed &&
					!errors.Is(err, http.ErrServerClosed) &&
					!errors.Is(err, net.ErrClosed) {
					return err
				}
				return nil
			})
		})

		// 启动net网络
		syncx.GoDelay(func() {
			t.L.Info("[cmux] Server Starting")
			logutil.LogOrErr(t.L, "[cmux] Server Stop", func() error {
				if err := t.net.Serve(); err != nil &&
					!errors.Is(err, http.ErrServerClosed) &&
					!errors.Is(err, net.ErrClosed) {
					return err
				}
				return nil
			})
		})
		return nil
	})

	logutil.OkOrPanic(t.L, "service after-start", func() error {
		var afterList []func()
		for _, p := range plugin.All() {
			afterList = append(afterList, p.AfterStarts()...)
		}
		afterList = append(afterList, t.afterStarts...)
		for i := range afterList {
			t.L.Sugar().Infof("running %s", stack.Func(afterList[i]))
			xerror.PanicF(xerror.Try(afterList[i]), stack.Func(afterList[i]))
		}
		return nil
	})
	return nil
}

func (t *serviceImpl) stop() (err error) {
	defer xerror.RespErr(&err)

	logutil.OkOrErr(t.L, "service before-stop", func() error {
		var beforeList []func()
		for _, p := range plugin.All() {
			beforeList = append(beforeList, p.BeforeStops()...)
		}
		beforeList = append(beforeList, t.beforeStops...)
		for i := range beforeList {
			t.L.Sugar().Infof("running %s", stack.Func(beforeList[i]))
			xerror.PanicF(xerror.Try(beforeList[i]), stack.Func(beforeList[i]))
		}
		return nil
	})

	logutil.LogOrErr(t.L, "[grpc] GracefulStop", func() error {
		t.srv.Get().GracefulStop()
		xerror.Panic(t.gw.Get().Shutdown())
		xerror.Panic(t.net.Close())
		return nil
	})

	logutil.OkOrErr(t.L, "service after-stop", func() error {
		var afterList []func()
		for _, p := range plugin.All() {
			afterList = append(afterList, p.AfterStops()...)
		}
		afterList = append(afterList, t.afterStops...)
		for i := range afterList {
			t.L.Sugar().Infof("running %s", stack.Func(afterList[i]))
			xerror.PanicF(xerror.Try(afterList[i]), stack.Func(afterList[i]))
		}
		return nil
	})

	return
}
