package debug

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/pubgo/dix"
	"github.com/pubgo/lug/app"
	"github.com/pubgo/lug/config"
	"github.com/pubgo/lug/entry"
	"github.com/pubgo/lug/plugin"
	"github.com/pubgo/x/fx"
	"github.com/pubgo/xerror"
	"github.com/pubgo/xlog"
)

func On(fn func(mux *chi.Mux)) { xerror.Panic(dix.Dix(fn)) }
func init()                    { plugin.Register(plg) }

var plg = &plugin.Base{
	Name: Name,
	OnInit: func(ent interface{}) {
		_ = config.Decode(Name, &cfg)

		srv := cfg.Build()
		var addr = fmt.Sprintf(":%d", app.DebugPort)
		var server = &http.Server{Addr: addr, Handler: srv}
		xerror.Panic(dix.Dix(srv))

		entry.Parse(ent, func(ent entry.Entry) {
			ent.BeforeStart(func() {
				xerror.Exit(fx.GoDelay(time.Second, func() {
					if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
						xlog.Error("Server [debug] Listen Error", xlog.Any("err", err))
					}

					xlog.Info("Server [debug] Closed OK")
				}))
				xlog.Infof("Server [debug] Listening on http://localhost%s", addr)
			})

			ent.BeforeStop(func() {
				if err := server.Shutdown(context.Background()); err != nil {
					xlog.Error("Server Shutdown Error", xlog.Any("err", err))
				}
			})
		})
	},

	OnVars: func(w func(name string, data func() interface{})) {
		type Route struct {
			Pattern   string
			Handlers  map[string]bool
			SubRoutes []Route
		}

		var getRoutes func(routes []chi.Route) []Route
		getRoutes = func(routes []chi.Route) []Route {
			if len(routes) == 0 {
				return nil
			}

			var routeList []Route
			for _, r := range appMux.Routes() {
				rr := Route{Pattern: r.Pattern, Handlers: make(map[string]bool)}

				for k := range r.Handlers {
					rr.Handlers[k] = true
				}

				//if r.SubRoutes != nil {
				//	rr.SubRoutes = getRoutes(r.SubRoutes.Routes())
				//}

				routeList = append(routeList, rr)
			}
			return routeList
		}

		w(Name+"_rest_router", func() interface{} {
			if appMux == nil {
				return nil
			}

			return getRoutes(appMux.Routes())
		})
	},
}
