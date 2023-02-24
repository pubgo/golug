package resty

import (
	"github.com/pubgo/lava/lava"
	"github.com/valyala/fasthttp"
)

var _ lava.Response = (*Response)(nil)

type Response struct {
	resp *fasthttp.Response
}

func (r *Response) Header() *lava.ResponseHeader { return &r.resp.Header }
func (r *Response) Payload() interface{}         { return nil }
func (r *Response) Stream() bool                    { return false }
