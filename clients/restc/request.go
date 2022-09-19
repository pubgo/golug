package restc

import (
	"github.com/gofiber/utils"
	"github.com/pubgo/lava/service"
	"github.com/valyala/fasthttp"
)

var _ service.Request = (*Request)(nil)

type Request struct {
	req     *fasthttp.Request
	service string
	ct      string
	data    []byte
}

func (r *Request) Operation() string              { return utils.UnsafeString(r.req.Header.Method()) }
func (r *Request) Kind() string                   { return Name }
func (r *Request) Client() bool                   { return true }
func (r *Request) Service() string                { return r.service }
func (r *Request) Endpoint() string               { return utils.UnsafeString(r.req.URI().Path()) }
func (r *Request) ContentType() string            { return r.ct }
func (r *Request) Header() *service.RequestHeader { return &r.req.Header }
func (r *Request) Payload() interface{}           { return r.data }
func (r *Request) Stream() bool                   { return false }
