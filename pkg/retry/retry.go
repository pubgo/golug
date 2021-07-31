package retry

import (
	"github.com/pubgo/xerror"

	"time"
)

type Handler func() Backoff

func (d Handler) Do(f func(i int) error) (err error) {
	var wrap = func(i int) (err error) {
		defer xerror.RespErr(&err)
		return f(i)
	}

	var b = d()
	for i := 0; ; i++ {
		if err = wrap(i); err == nil {
			return
		}

		dur, stop := b.Next()
		if stop {
			return
		}

		time.Sleep(dur)
	}
}

func New(bs ...Backoff) Handler {
	var b = WithMaxRetries(3, NewConstant(DefaultConstant))
	if len(bs) > 0 {
		b = bs[0]
	}

	return func() Backoff { return b }
}
