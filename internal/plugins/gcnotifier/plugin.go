package gcnotifier

import (
	"github.com/CAFxX/gcnotifier"

	"github.com/pubgo/lava/internal/logz"
	"github.com/pubgo/lava/pkg/syncx"
	"github.com/pubgo/lava/plugin"
	"github.com/pubgo/lava/runenv"
)

var Name = "gc"
var logs = logz.New(Name)

func init() {
	if runenv.IsProd() || runenv.IsRelease() {
		return
	}

	plugin.Register(&plugin.Base{
		Name: Name,
		OnInit: func() {
			syncx.GoSafe(func() {
				var gc = gcnotifier.New()
				defer gc.Close()

				for range gc.AfterGC() {
					logs.Infow("gc notify")
				}
			})
		},
	})
}
