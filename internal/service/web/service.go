package rests

import (
	"errors"
	"fmt"

	"net"
	"net/http"
	"strings"

	fiber2 "github.com/gofiber/fiber/v2"
	"github.com/pubgo/dix"
	"github.com/pubgo/funk"
	"github.com/pubgo/funk/xerr"
	"github.com/pubgo/x/stack"
	"github.com/pubgo/xerror"
	"github.com/urfave/cli/v2"
	"go.uber.org/zap"

	"github.com/pubgo/lava/core/flags"
	"github.com/pubgo/lava/core/lifecycle"
	"github.com/pubgo/lava/core/runmode"
	cmux2 "github.com/pubgo/lava/internal/cmux"
	fiber_builder2 "github.com/pubgo/lava/internal/pkg/fiber_builder"
	netutil2 "github.com/pubgo/lava/internal/pkg/netutil"
	"github.com/pubgo/lava/internal/pkg/syncx"
	"github.com/pubgo/lava/internal/pkg/utils"
	"github.com/pubgo/lava/logging/logutil"
	"github.com/pubgo/lava/service"
	"github.com/pubgo/lava/version"
)

func New(name string, desc ...string) service.Web {
	return newService(name, desc...)
}

func newService(name string, desc ...string) *serviceImpl {
	var g *serviceImpl
	g = &serviceImpl{
		httpSrv: fiber_builder2.New(),
		cmd: &cli.Command{
			Name:  name,
			Usage: utils.FirstNotEmpty(append(desc, fmt.Sprintf("%s service", name))...),
			Flags: flags.GetFlags(),
			Before: func(context *cli.Context) (gErr error) {
				defer funk.RecoverErr(&gErr, func(err xerr.XErr) xerr.XErr {
					fmt.Println(dix.Graph())
					return err
				})

				if runmode.Project == "" {
					runmode.Project = strings.Split(name, " ")[0]
				}
				funk.Assert(runmode.Project == "", "project is null")

				for i := range g.providerList {
					dix.Provider(g.providerList[i])
				}
				return
			},
		},
	}

	return g
}

var _ service.Web = (*serviceImpl)(nil)

type serviceImpl struct {
	getLifecycle lifecycle.GetLifecycle
	lc           lifecycle.Lifecycle
	log          *zap.Logger
	net          *cmux2.Mux
	app          *service.WebApp
	cfg          *Cfg

	cmd          *cli.Command
	httpSrv      fiber_builder2.Builder
	providerList []interface{}
	httpMiddle   func(_ *fiber2.Ctx) error
}

func (t *serviceImpl) dixInject(p struct {
	Middlewares  []service.Middleware
	GetLifecycle lifecycle.GetLifecycle
	Lifecycle    lifecycle.Lifecycle
	Log          *zap.Logger
	Net          *cmux2.Mux
	App          *service.WebApp
	Cfg          *Cfg
}) {
	t.getLifecycle = p.GetLifecycle
	t.lc = p.Lifecycle
	t.log = p.Log.Named(runmode.Project)
	t.net = p.Net
	t.app = p.App
	t.cfg = p.Cfg

	var middlewares []service.Middleware
	for _, m := range p.Middlewares {
		middlewares = append(middlewares, m)
	}

	t.httpMiddle = t.handlerHttpMiddle(middlewares)
}

func (t *serviceImpl) Start() error { return t.start() }
func (t *serviceImpl) Stop() error  { return t.stop() }

func (t *serviceImpl) init() (gErr error) {
	defer funk.RecoverErr(&gErr)

	dix.Inject(t)

	// 网关初始化
	funk.Must(t.httpSrv.Build(t.cfg.Api))
	t.httpSrv.Get().Use(t.httpMiddle)
	t.httpSrv.Get().Mount("/", t.app.App)

	if t.cfg.PrintRoute {
		for _, stacks := range t.httpSrv.Get().Stack() {
			for _, s := range stacks {
				t.log.Info("service route",
					zap.String("name", s.Name),
					zap.String("path", s.Path),
					zap.String("method", s.Method),
				)
			}
		}
	}

	return nil
}

func (t *serviceImpl) SubCmd(cmd *cli.Command) {
	t.cmd.Subcommands = append(t.cmd.Subcommands, cmd)
}

func (t *serviceImpl) Provider(provider interface{}) {
	t.providerList = append(t.providerList, provider)
}

func (t *serviceImpl) Command() *cli.Command { return t.cmd }

func (t *serviceImpl) Options() service.Options {
	return service.Options{
		Name:      runmode.Project,
		Id:        runmode.InstanceID,
		Version:   version.Version,
		Port:      netutil2.MustGetPort(t.net.Addr),
		Addr:      t.net.Addr,
		Advertise: "",
	}
}

func (t *serviceImpl) start() (gErr error) {
	defer funk.RecoverErr(&gErr)

	funk.Must(t.init())

	logutil.OkOrPanic(t.log, "service before-start", func() error {
		for _, run := range t.getLifecycle.GetBeforeStarts() {
			t.log.Sugar().Infof("before-start running %s", stack.Func(run))
			funk.MustF(xerror.Try(run), stack.Func(run))
		}
		return nil
	})

	var gwLn = t.net.HTTP1()

	logutil.OkOrPanic(t.log, "service start", func() error {
		t.log.Sugar().Infof("Server Listening on http://%s:%d", netutil2.GetLocalIP(), netutil2.MustGetPort(t.net.Addr))

		// 启动grpc网关
		syncx.GoDelay(func() {
			t.log.Info("[grpc-gw] Server Starting")
			logutil.LogOrErr(t.log, "[grpc-gw] Server Stop", func() error {
				if err := t.httpSrv.Get().Listener(<-gwLn); err != nil &&
					!errors.Is(err, cmux2.ErrListenerClosed) &&
					!errors.Is(err, http.ErrServerClosed) &&
					!errors.Is(err, net.ErrClosed) {
					return err
				}
				return nil
			})
		})

		// 启动net网络
		syncx.GoDelay(func() {
			t.log.Info("[cmux] Server Starting")
			logutil.LogOrErr(t.log, "[cmux] Server Stop", func() error {
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

	logutil.OkOrPanic(t.log, "service after-start", func() error {
		for _, run := range t.getLifecycle.GetAfterStarts() {
			t.log.Sugar().Infof("after-start running %s", stack.Func(run))
			xerror.PanicF(xerror.Try(run), stack.Func(run))
		}
		return nil
	})
	return nil
}

func (t *serviceImpl) stop() (err error) {
	defer xerror.RecoverErr(&err)

	logutil.OkOrErr(t.log, "service before-stop", func() error {
		for _, run := range t.getLifecycle.GetBeforeStops() {
			t.log.Sugar().Infof("before-stop running %s", stack.Func(run))
			xerror.PanicF(xerror.Try(run), stack.Func(run))
		}
		return nil
	})

	logutil.LogOrErr(t.log, "[grpc-gw] Shutdown", func() error {
		xerror.Panic(t.httpSrv.Get().Shutdown())
		return nil
	})

	logutil.OkOrErr(t.log, "service after-stop", func() error {
		for _, run := range t.getLifecycle.GetAfterStops() {
			t.log.Sugar().Infof("after-stop running %s", stack.Func(run))
			xerror.PanicF(xerror.Try(run), stack.Func(run))
		}
		return nil
	})

	return
}
