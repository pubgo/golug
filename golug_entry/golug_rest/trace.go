package golug_rest

import (
	"fmt"

	"github.com/pubgo/dix/dix_run"
	"github.com/pubgo/golug/golug_env"
	"github.com/pubgo/golug/internal/golug_util"
	"github.com/pubgo/xerror"
	"github.com/pubgo/xlog"
)

func (t *restEntry) trace() {
	xerror.Panic(dix_run.WithAfterStart(func(ctx *dix_run.AfterStartCtx) {
		if !golug_env.Trace || !t.Options().Initialized {
			return
		}

		xlog.Debug("rest server router trace")
		for _, stacks := range t.app.Stack() {
			for _, stack := range stacks {
				if stack.Path == "/" {
					continue
				}

				xlog.Debugf("%s %s", stack.Method, stack.Path)
			}
		}
		fmt.Println()

		xlog.Debugf("rest server config trace")
		fmt.Println(golug_util.MarshalIndent(cfg))
		fmt.Println()
	}))
}
