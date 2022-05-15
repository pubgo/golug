package pidfile

import (
	"go.uber.org/fx"
	"path/filepath"

	"github.com/pubgo/x/pathutil"
	"github.com/pubgo/xerror"

	"github.com/pubgo/lava/config"
	"github.com/pubgo/lava/core/lifecycle"
	"github.com/pubgo/lava/inject"
)

func init() {
	inject.Register(fx.Provide(func(r lifecycle.Lifecycle) {
		pidPath = filepath.Join(config.CfgDir, "pidfile")

		_ = pathutil.IsNotExistMkDir(pidPath)

		r.AfterStarts(func() {
			xerror.Panic(SavePid())
		})
	}))
}
