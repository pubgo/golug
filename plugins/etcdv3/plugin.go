package etcdv3

import (
	"github.com/pubgo/xerror"

	"github.com/pubgo/lug/config"
	"github.com/pubgo/lug/consts"
	"github.com/pubgo/lug/entry"
	"github.com/pubgo/lug/logger"
	"github.com/pubgo/lug/plugin"
	"github.com/pubgo/lug/watcher"
)

func init() {
	plugin.Register(&plugin.Base{
		Name: Name,
		OnInit: func(ent entry.Entry) {
			_ = config.Decode(Name, &cfgList)
			for name, cfg := range cfgList {
				// etcd config处理
				cfg := xerror.PanicErr(cfgMerge(cfg)).(Cfg)
				xerror.Panic(Update(consts.GetDefault(name), cfg))
			}
		},
		OnWatch: func(name string, r *watcher.Response) {
			r.OnPut(func() {
				// 解析etcd配置
				var cfg Cfg
				xerror.PanicF(watcher.Decode(r.Value, &cfg), "etcd conf parse error, cfg: %s", r.Value)

				cfg = xerror.PanicErr(cfgMerge(cfg)).(Cfg)
				xerror.PanicF(Update(name, cfg), "client %s watcher update error", name)
			})

			r.OnDelete(func() {
				logs.Debug("delete client", logger.Name(name))
				Delete(name)
			})
		},
	})
}
