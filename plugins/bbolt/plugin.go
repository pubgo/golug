package bbolt

import (
	"github.com/pubgo/lava/resource"
	"github.com/pubgo/xerror"

	"github.com/pubgo/lava/config"
	"github.com/pubgo/lava/plugin"
)

func init() {
	plugin.Register(&plugin.Base{
		Name: Name,
		OnInit: func(ent plugin.Entry) {
			xerror.PanicF(config.Decode(Name, &cfgMap), "config [%s] not found", Name)

			for k, v := range cfgMap {
				var db = v.Build()
				resource.Update(k, &Client{db})
			}
		},
	})
}
