package grpc

import (
	"context"
	"fmt"
	"net"
	"strconv"
	"strings"
	"time"

	grpcGw "github.com/pubgo/lug/builder/grpc-gw"
	"github.com/pubgo/lug/builder/grpcs"
	"github.com/pubgo/lug/config"
	"github.com/pubgo/lug/entry/base"
	"github.com/pubgo/lug/pkg/ctxutil"
	"github.com/pubgo/lug/pkg/logutil"
	"github.com/pubgo/lug/pkg/netutil"
	"github.com/pubgo/lug/plugins/grpcc"
	"github.com/pubgo/lug/registry"
	"github.com/pubgo/lug/runenv"
	"github.com/pubgo/lug/version"

	"github.com/google/uuid"
	"github.com/pubgo/x/fx"
	"github.com/pubgo/xerror"
	"github.com/soheilhy/cmux"
	"go.uber.org/atomic"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

var _ Entry = (*grpcEntry)(nil)

type grpcEntry struct {
	*base.Entry
	cfg Cfg
	mux cmux.CMux
	rpc grpcs.Builder
	gw  grpcGw.Builder

	registry    registry.Registry
	registered  atomic.Bool
	registryMap map[string][]*registry.Endpoint

	handlers  []interface{}
	endpoints []*registry.Endpoint
	client    *grpc.ClientConn
}

func (g *grpcEntry) Init(opts ...grpc.ServerOption) { g.rpc.Init(opts...) }
func (g *grpcEntry) UnaryInterceptor(interceptors ...grpc.UnaryServerInterceptor) {
	g.rpc.UnaryInterceptor(interceptors...)
}

func (g *grpcEntry) StreamInterceptor(interceptors ...grpc.StreamServerInterceptor) {
	g.rpc.StreamInterceptor(interceptors...)
}

func (g *grpcEntry) serve() error { return g.mux.Serve() }
func (g *grpcEntry) handleError() {
	g.mux.HandleError(func(err error) bool {
		logs.Error("grpcEntry mux handleError", logutil.Err(err), logutil.Name(g.cfg.name))
		return false
	})
}
func (g *grpcEntry) matchAny() net.Listener   { return g.mux.Match(cmux.Any()) }
func (g *grpcEntry) matchHttp1() net.Listener { return g.mux.Match(cmux.HTTP1()) }
func (g *grpcEntry) matchHttp2() net.Listener {
	return g.mux.Match(
		cmux.HTTP2(),
		cmux.HTTP2HeaderFieldPrefix("content-type", "application/grpc"),
		cmux.HTTP2HeaderFieldPrefix("Content-Type", "application/grpc"),
	)
}

func (g *grpcEntry) register() (err error) {
	defer xerror.RespErr(&err)

	if g.registry == nil {
		return nil
	}

	// parse address for host, port
	var advt, host string
	var port int

	// check the advertise address first
	// if it exists then use it, otherwise
	// use the address
	if len(g.cfg.Advertise) > 0 {
		advt = g.cfg.Advertise
	} else {
		advt = g.cfg.Address
	}

	parts := strings.Split(advt, ":")
	if len(parts) > 1 {
		host = strings.Join(parts[:len(parts)-1], ":")
		port, _ = strconv.Atoi(parts[len(parts)-1])
	} else {
		host = parts[0]
	}

	// register service
	node := &registry.Node{
		Id:      g.cfg.name + "-" + g.cfg.hostname + "-" + g.cfg.id,
		Address: fmt.Sprintf("%s:%d", host, port),
		Port:    port,
	}

	node.Metadata["registry"] = g.registry.String()
	node.Metadata["transport"] = "grpc"

	services := &registry.Service{
		Name:      g.cfg.name,
		Version:   version.Version,
		Nodes:     []*registry.Node{node},
		Endpoints: g.endpoints,
	}

	if !g.registered.Load() {
		logs.Info("Registering node", logutil.Id(node.Id), logutil.Name(g.cfg.name))
	}

	// registry options
	opts := []registry.RegOpt{registry.TTL(g.cfg.RegisterTTL)}
	xerror.Panic(g.registry.Register(services, opts...), "[grpc] register error")

	// already registered? don't need to register subscribers
	if g.registered.Load() {
		return nil
	}

	g.registered.Store(true)
	return nil
}

func (g *grpcEntry) deRegister() (err error) {
	defer xerror.RespErr(&err)

	if g.registry == nil {
		return nil
	}

	var advt, host string
	var port int

	// check the advertise address first
	// if it exists then use it, otherwise
	// use the address
	if len(g.cfg.Advertise) > 0 {
		advt = g.cfg.Advertise
	} else {
		advt = g.cfg.Address
	}

	parts := strings.Split(advt, ":")
	if len(parts) > 1 {
		host = strings.Join(parts[:len(parts)-1], ":")
		port, _ = strconv.Atoi(parts[len(parts)-1])
	} else {
		host = parts[0]
	}

	node := &registry.Node{
		Id:      g.cfg.name + "-" + g.cfg.hostname + "-" + g.cfg.id,
		Address: fmt.Sprintf("%s:%d", host, port),
		Port:    port,
	}

	services := &registry.Service{
		Name:    g.cfg.name,
		Version: version.Version,
		Nodes:   []*registry.Node{node},
	}

	logs.Infof("DeRegistering node: %s", node.Id)
	xerror.Panic(g.registry.DeRegister(services), "[grpc] registry deRegister error")

	if !g.registered.Load() {
		return nil
	}

	g.registered.Store(false)
	return nil
}

func (g *grpcEntry) Stop() (err error) {
	defer xerror.RespErr(&err)

	// deRegister self
	if err := g.deRegister(); err != nil {
		logs.Info("[grpc] server deRegister error", logutil.Err(err))
	}

	// Add sleep for those requests which have selected this port.
	time.Sleep(g.cfg.SleepAfterDeRegister)

	logs.Info("[ExitProgress] Start Shutdown.")
	if err := g.gw.Get().Shutdown(ctxutil.Default()); err != nil && !strings.Contains(err.Error(), net.ErrClosed.Error()) {
		logs.Error("[grpc-gw] Shutdown Error", logutil.Err(err))
	} else {
		logs.Info("[ExitProgress] Shutdown Ok.")
	}

	// stop the grpc server
	logs.Info("[ExitProgress] Start GracefulStop.")
	g.rpc.Get().GracefulStop()
	logs.Info("[ExitProgress] GracefulStop Ok.")

	return
}

func (g *grpcEntry) initHandler() {
	xerror.RespExit()

	// 初始化routes
	for i := range g.handlers {
		xerror.PanicF(grpcs.Register(g.rpc.Get(), g.handlers[i]), "register handler error")
	}
}

// 初始化gw routes
func (g *grpcEntry) initGwHandler() {
	xerror.ExitF(g.gw.Register(g.client), "gw register handler error")
}

func (g *grpcEntry) Register(handler interface{}, opts ...Opt) {
	defer xerror.RespExit()

	xerror.Assert(handler == nil, "[handler] should not be nil")
	xerror.Assert(!grpcs.FindHandle(handler).IsValid(), "register [%#v] 没有找到匹配的interface", handler)

	g.handlers = append(g.handlers, handler)
	g.endpoints = append(g.endpoints, newRpcHandler(handler)...)
}

func (g *grpcEntry) Start(args ...string) (gErr error) {
	defer xerror.RespErr(&gErr)

	logs.Info("Server Listening", logutil.Name(g.cfg.name), zap.String("addr", runenv.Addr))
	ln := xerror.PanicErr(netutil.Listen(runenv.Addr)).(net.Listener)
	g.mux = cmux.New(ln)
	g.handleError()

	// 启动grpc服务
	fx.GoDelay(func() {
		logs.Info("[grpc] Server Starting")
		if err := g.rpc.Get().Serve(g.matchHttp2()); err != nil && err != cmux.ErrListenerClosed {
			logs.Error("[grpc] Server Stop", logutil.Err(err))
		}
	})

	// 启动grpc网关
	fx.GoDelay(func() {
		logs.Info("[grpc-gw] Server Staring")
		if err := g.gw.Get().Serve(g.matchHttp1()); err != nil && err != cmux.ErrListenerClosed {
			logs.Error(" [grpc-gw] Server Stop", logutil.Err(err))
		}
	})

	// 启动net网络
	fx.GoDelay(func() {
		logs.Info("[mux] Server Staring")
		if err := g.serve(); err != nil && !strings.Contains(err.Error(), net.ErrClosed.Error()) {
			logs.Error(" [mux] Server Stop", logutil.Err(err))
		}
	})

	// 启动本地grpc客户端
	fx.GoDelay(func() {
		logs.Info("[grpc] Client Connecting")
		conn, err := grpcc.NewDirect(runenv.Addr)
		xerror.Panic(err, "[grpc] Client Connecting Error")
		g.client = conn
		g.initGwHandler()
	})

	// register self
	xerror.Panic(g.register(), "[grpc] try to register self")

	_ = fx.Go(func(ctx context.Context) {
		var interval = DefaultRegisterInterval

		// only process if it exists
		if g.cfg.RegisterInterval > time.Duration(0) {
			interval = g.cfg.RegisterInterval
		}

		// register self on interval
		for range time.NewTicker(interval).C {
			if err := g.register(); err != nil {
				logs.Error("[grpc] server register on interval", logutil.Err(err))
			}
		}
	})

	return nil
}

func newEntry(name string) *grpcEntry {
	var g = &grpcEntry{
		Entry: base.New(name),
		rpc:   grpcs.New(name),
		gw:    grpcGw.New(name),
		cfg: Cfg{
			name:                 name,
			hostname:             getHostname(),
			id:                   uuid.New().String(),
			Rpc:                  grpcs.GetDefaultCfg(),
			Gw:                   grpcGw.GetDefaultCfg(),
			RegisterTTL:          time.Minute,
			RegisterInterval:     time.Second * 30,
			SleepAfterDeRegister: time.Second * 2,
		},
	}

	g.OnInit(func() {
		grpcs.InitEncoding()
		_ = config.Decode(Name, &g.cfg)

		// 注册中心校验
		g.registry = registry.Default()

		xerror.Panic(g.gw.Build(g.cfg.Gw))

		xerror.Panic(g.rpc.Build(g.cfg.Rpc, func() {
			g.initHandler()
		}))
	})

	return g
}

func New(name string) Entry { return newEntry(name) }
