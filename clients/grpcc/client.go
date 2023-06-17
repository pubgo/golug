package grpcc

import (
	"context"
	"fmt"
	"sync"

	"github.com/pubgo/funk/assert"
	"github.com/pubgo/funk/errors"
	"github.com/pubgo/funk/log"
	"github.com/pubgo/funk/recovery"
	"github.com/pubgo/funk/result"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health/grpc_health_v1"

	"github.com/pubgo/lava"
	"github.com/pubgo/lava/clients/grpcc/grpcc_config"
	"github.com/pubgo/lava/clients/grpcc/grpcc_resolver"
	"github.com/pubgo/lava/core/config"
	"github.com/pubgo/lava/core/logging/logkey"
	"github.com/pubgo/lava/internal/middlewares/middleware_accesslog"
	"github.com/pubgo/lava/internal/middlewares/middleware_metric"
	"github.com/pubgo/lava/internal/middlewares/middleware_recovery"
	"github.com/pubgo/lava/internal/middlewares/middleware_service_info"
)

type Params struct {
	Log       log.Logger
	ReqMetric *middleware_metric.MetricMiddleware
	AccessLog *middleware_accesslog.LogMiddleware
}

func New(cfg *grpcc_config.Cfg, p Params, middlewares ...lava.Middleware) Client {
	cfg = config.Merge(grpcc_config.DefaultCfg(), cfg)
	var defaultMiddlewares = []lava.Middleware{
		middleware_service_info.New(),
		p.ReqMetric,
		p.AccessLog,
		middleware_recovery.New(),
	}
	defaultMiddlewares = append(defaultMiddlewares, middlewares...)

	c := &clientImpl{
		cfg:         cfg,
		log:         p.Log,
		middlewares: defaultMiddlewares,
	}

	if cfg.Client.Block {
		c.Get().Unwrap()
	}

	return c
}

type clientImpl struct {
	log         log.Logger
	cfg         *grpcc_config.Cfg
	mu          sync.Mutex
	conn        grpc.ClientConnInterface
	middlewares []lava.Middleware
}

func (t *clientImpl) Invoke(ctx context.Context, method string, args interface{}, reply interface{}, opts ...grpc.CallOption) (err error) {
	defer recovery.Err(&err, func(err error) error {
		return errors.WrapTag(err, errors.T("method", method), errors.T("args", args))
	})

	conn := t.Get().Unwrap()
	assert.Must(conn.Invoke(ctx, method, args, reply, opts...))
	return
}

func (t *clientImpl) Healthy(ctx context.Context) error {
	conn := t.Get()
	if conn.IsErr() {
		return errors.Wrapf(conn.Err(), "get client failed, service=%s", t.cfg.Srv)
	}

	_, err := grpc_health_v1.NewHealthClient(conn.Unwrap()).Check(ctx, &grpc_health_v1.HealthCheckRequest{})
	if err != nil {
		return errors.Wrapf(err, "service %s heath check failed", t.cfg.Srv)
	}
	return nil
}

func (t *clientImpl) NewStream(ctx context.Context, desc *grpc.StreamDesc, method string, opts ...grpc.CallOption) (grpc.ClientStream, error) {
	conn := t.Get()
	if conn.IsErr() {
		return nil, errors.Wrapf(conn.Err(), "get client failed, service=%s method=%s", t.cfg.Srv, method)
	}

	c, err1 := conn.Unwrap().NewStream(ctx, desc, method, opts...)
	return c, errors.Wrap(err1, method)
}

// Get new grpc client
func (t *clientImpl) Get() (r result.Result[grpc.ClientConnInterface]) {
	defer recovery.Result(&r)

	if t.conn != nil {
		return r.WithVal(t.conn)
	}

	t.mu.Lock()
	defer t.mu.Unlock()

	// 双检, 避免多次创建
	if t.conn != nil {
		return r.WithVal(t.conn)
	}

	conn, err := createConn(t.cfg, t.log, t.middlewares)
	if err != nil {
		return r.WithErr(err)
	}

	t.conn = conn
	return r.WithVal(t.conn)
}

func buildTarget(cfg *grpcc_config.Cfg) string {
	addr := cfg.Addr
	scheme := grpcc_resolver.DirectScheme
	if cfg.Scheme != "" {
		scheme = cfg.Scheme
	}

	switch scheme {
	case grpcc_resolver.DiscovScheme:
		return grpcc_resolver.BuildDiscovTarget(addr)
	case grpcc_resolver.DirectScheme:
		return grpcc_resolver.BuildDirectTarget(addr)
	case grpcc_resolver.K8sScheme, grpcc_resolver.DnsScheme:
		return fmt.Sprintf("dns:///%s", addr)
	default:
		return addr
	}
}

func createConn(cfg *grpcc_config.Cfg, log log.Logger, mm []lava.Middleware) (grpc.ClientConnInterface, error) {
	// 创建grpc client
	ctx, cancel := context.WithTimeout(context.Background(), cfg.Client.DialTimeout)
	defer cancel()

	addr := buildTarget(cfg)

	ee := log.Info().
		Str(logkey.Service, cfg.Srv).
		Str("addr", addr)
	ee.Msg("grpc client init")

	conn, err := grpc.DialContext(ctx, addr, append(
		cfg.Client.ToOpts(),
		grpc.WithChainUnaryInterceptor(unaryInterceptor(mm)),
		grpc.WithChainStreamInterceptor(streamInterceptor(mm)))...)
	if err != nil {
		return nil, errors.Wrapf(err, "grpc dial failed, target=>%s", addr)
	}

	ee.Msg("grpc client init ok")
	return conn, nil
}
