package restc

import (
	"context"
	"crypto/tls"

	"github.com/pubgo/lug/pkg/retry"
	"github.com/valyala/fasthttp"
)

type Option func(opts *Cfg)

func WithRetryIf(rf fasthttp.RetryIfFunc) Option {
	return func(c *Cfg) { c.retryIf = rf }
}

func WithTLS(tc *tls.Config) Option {
	return func(c *Cfg) { c.tlsConfig = tc }
}

func WithDial(dial fasthttp.DialFunc) Option {
	return func(c *Cfg) { c.dial = dial }
}

func WithBackoff(bk retry.Backoff) Option {
	return func(c *Cfg) { c.backoff = bk }
}

func WithMiddle(middles ...Middleware) Option {
	return func(c *Cfg) { c.middles = middles }
}

func Context(ctx context.Context) func(req *Request) {
	return func(req *Request) {
		req.Context = ctx
	}
}
