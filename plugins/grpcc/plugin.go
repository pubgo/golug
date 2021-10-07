package grpcc

import (
	"github.com/pubgo/x/merge"
	"github.com/pubgo/xerror"

	"github.com/pubgo/lug/config"
	"github.com/pubgo/lug/plugin"
)

func init() {
	plugin.Register(&plugin.Base{
		Name: Name,
		OnInit: func(ent plugin.Entry) {
			_ = config.Decode(Name, &cfgMap)
			for k := range cfgMap {
				var cfg = cfgMap[k]
				var defCfg = DefaultCfg()
				xerror.Panic(merge.Copy(&defCfg, &cfg))
				cfgMap[k] = defCfg

			}
		},
	})
}
