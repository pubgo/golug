package restc

import (
	"context"
	"net/url"
	"time"

	"github.com/pubgo/lug/pkg/retry"
	"github.com/pubgo/lug/tracing"
	"github.com/pubgo/lug/types"

	"github.com/opentracing/opentracing-go/ext"
	"github.com/pubgo/xerror"
	"github.com/valyala/fasthttp"
)

const (
	defaultRetryCount  = 1
	defaultHTTPTimeout = 2 * time.Second
	defaultContentType = "application/json"
)

var _ Client = (*client)(nil)

// client is the Client implementation
type client struct {
	client        *fasthttp.Client
	defaultHeader *fasthttp.RequestHeader
	cfg           Cfg
	do            types.MiddleNext
}

func (c *client) Do(ctx context.Context, req *Request) (resp *Response, err error) {
	return resp, xerror.Wrap(c.do(
		ctx,
		&request{
			req:    req,
			header: convertHeader(&req.Header),
		},
		func(res types.Response) error {
			resp = res.(*response).resp
			return nil
		},
	))
}

func (c *client) Get(ctx context.Context, url string, requests ...func(req *Request)) (*Response, error) {
	var resp, err = doUrl(ctx, c, fasthttp.MethodGet, url, requests...)
	return resp, xerror.Wrap(err)
}

func (c *client) Delete(ctx context.Context, url string, requests ...func(req *Request)) (*Response, error) {
	var resp, err = doUrl(ctx, c, fasthttp.MethodDelete, url, requests...)
	return resp, xerror.Wrap(err)
}

func (c *client) Post(ctx context.Context, url string, requests ...func(req *Request)) (*Response, error) {
	var resp, err = doUrl(ctx, c, fasthttp.MethodPost, url, requests...)
	return resp, xerror.Wrap(err)
}

func (c *client) PostForm(ctx context.Context, url string, val url.Values, requests ...func(req *Request)) (*Response, error) {
	var resp, err = doUrl(ctx, c, fasthttp.MethodPost, url, func(req *Request) {
		req.SetBodyString(val.Encode())
		req.Header.SetContentType("application/x-www-form-urlencoded")
		if len(requests) > 0 {
			requests[0](req)
		}
	})
	return resp, xerror.Wrap(err)
}

func (c *client) Put(ctx context.Context, url string, requests ...func(req *Request)) (*Response, error) {
	var resp, err = doUrl(ctx, c, fasthttp.MethodPut, url, requests...)
	return resp, xerror.Wrap(err)
}

func (c *client) Patch(ctx context.Context, url string, requests ...func(req *Request)) (*Response, error) {
	var resp, err = doUrl(ctx, c, fasthttp.MethodPatch, url, requests...)
	return resp, xerror.Wrap(err)
}

func doUrl(ctx context.Context, c *client, mth string, url string, requests ...func(req *Request)) (*Response, error) {
	var req = fasthttp.AcquireRequest()
	c.defaultHeader.CopyTo(&req.Header)

	req.SetRequestURI(url)
	req.Header.SetMethod(mth)

	if len(requests) > 0 {
		requests[0](req)
	}

	var resp, err = c.Do(ctx, req)
	fasthttp.ReleaseRequest(req)

	if err != nil {
		return nil, xerror.Wrap(err)
	}

	return resp, nil
}

func doFunc(c *client) types.MiddleNext {
	var backoff = retry.New(retry.WithMaxRetries(c.cfg.RetryCount, c.cfg.backoff))

	return func(ctx context.Context, req types.Request, callback func(rsp types.Response) error) error {
		var resp = fasthttp.AcquireResponse()

		defer func() {
			tracing.SpanFromCtx(ctx, func(span *tracing.Span) {
				ext.HTTPStatusCode.Set(span, uint16(resp.StatusCode()))
			})
		}()

		xerror.Panic(backoff.Do(func(i int) error {
			if c.cfg.Timeout > 0 {
				return xerror.Wrap(c.client.DoTimeout(req.(*request).req, resp, c.cfg.Timeout))
			}

			return xerror.Wrap(c.client.Do(req.(*request).req, resp))
		}))

		return xerror.Wrap(callback(&response{resp: resp}))
	}
}
