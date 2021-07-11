// Package mdns is a multicast dns registry
package mdns

import (
	"context"
	"fmt"
	"go.uber.org/atomic"
	"sync"
	"time"

	"github.com/pubgo/lug/pkg/typex"
	"github.com/pubgo/lug/registry"
	"github.com/pubgo/lug/runenv"

	"github.com/grandcat/zeroconf"
	"github.com/pubgo/x/fx"
	"github.com/pubgo/x/merge"
	"github.com/pubgo/x/xutil"
	"github.com/pubgo/xerror"
)

func init() {
	registry.Register(Name, NewWithMap)
}

func NewWithMap(m map[string]interface{}) (registry.Registry, error) {
	resolver, err := zeroconf.NewResolver()
	xerror.Panic(err, "Failed to initialize zeroconf resolver")

	var r = &mdnsRegistry{resolver: resolver}

	xerror.Panic(merge.MapStruct(&r.cfg, m))
	return r, nil
}

var _ registry.Registry = (*mdnsRegistry)(nil)
var _ registry.Watcher = (*mdnsRegistry)(nil)

type mdnsRegistry struct {
	closed   atomic.Bool
	cfg      Cfg
	services sync.Map
	results  chan *registry.Result
	resolver *zeroconf.Resolver
	cancel   context.CancelFunc
}

func (m *mdnsRegistry) Next() (*registry.Result, error) {
	result, ok := <-m.results
	if !ok {
		return nil, registry.ErrWatcherStopped
	}

	return result, nil
}

func (m *mdnsRegistry) Stop() {
	if m.closed.Load() {
		return
	}
	m.closed.Store(true)

	close(m.results)
	if m.cancel != nil {
		m.cancel()
	}
}

func (m *mdnsRegistry) Register(service *registry.Service, optList ...registry.RegOpt) (err error) {
	defer xerror.RespErr(&err)

	xerror.Assert(service == nil, "[service] should not be nil")
	xerror.Assert(len(service.Nodes) == 0, "[service] nodes should not be zero")

	node := service.Nodes[0]
	server, err := zeroconf.Register(service.Name, runenv.Domain, "local", node.GetPort(), []string{node.Id}, nil)
	xerror.PanicF(err, "[mdns] service %s register error", service.Name)

	var opts registry.RegOpts
	for i := range optList {
		optList[i](&opts)
	}

	if opts.TTL != 0 {
		server.TTL(uint32(opts.TTL.Seconds()))
	}

	m.services.Store(node.Id, server)
	return nil
}

func (m *mdnsRegistry) DeRegister(service *registry.Service, opt ...registry.DeRegOpt) (err error) {
	defer xerror.RespErr(&err)

	xerror.Assert(service == nil, "[service] should not be nil")
	xerror.Assert(len(service.Nodes) == 0, "[service] nodes should not be zero")

	node := service.Nodes[0]
	var val, ok = m.services.LoadAndDelete(node.Id)
	if !ok {
		return nil
	}

	val.(*zeroconf.Server).Shutdown()

	return nil
}

func (m *mdnsRegistry) GetService(name string, opts ...registry.GetOpt) (services []*registry.Service, _ error) {
	return services, xutil.Try(func() {
		entries := make(chan *zeroconf.ServiceEntry)
		_ = fx.Go(func(ctx context.Context) {
			for s := range entries {
				services = append(services, &registry.Service{
					Name: s.Instance,
					Nodes: registry.Nodes{{
						Id:      s.Text[0],
						Port:    s.Port,
						Address: fmt.Sprintf("%s:%d", s.AddrIPv4[0].String(), s.Port),
					}},
				})
			}
		})

		var gOpts registry.GetOpts
		for i := range opts {
			opts[i](&gOpts)
		}

		if gOpts.Timeout == 0 {
			gOpts.Timeout = time.Second * 5
		}

		ctx, cancel := context.WithTimeout(context.Background(), gOpts.Timeout)
		defer cancel()

		xerror.PanicF(m.resolver.Lookup(ctx, name, runenv.Domain, "local", entries), "Failed to Lookup Service %s", name)
		<-ctx.Done()
	})
}

func (m *mdnsRegistry) ListServices(opts ...registry.ListOpt) (services []*registry.Service, _ error) {
	return services, xutil.Try(func() {
		entries := make(chan *zeroconf.ServiceEntry)
		_ = fx.GoLoop(func(ctx fx.Ctx) {
			var s, ok = <-entries
			if !ok {
				ctx.Break()
			}

			services = append(services, &registry.Service{
				Name: s.Instance,
				Nodes: registry.Nodes{{
					Id:      s.Text[0],
					Port:    s.Port,
					Address: fmt.Sprintf("%s:%d", s.AddrIPv4[0].String(), s.Port),
				}},
			})
		})

		ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
		defer cancel()

		xerror.PanicF(m.resolver.Browse(ctx, runenv.Domain, "local", entries), "Failed to Browse Service")
		<-ctx.Done()
	})
}

func (m *mdnsRegistry) Watch(service string, opt ...registry.WatchOpt) (registry.Watcher, error) {
	var watcher = &mdnsRegistry{results: make(chan *registry.Result)}

	return watcher, xutil.Try(func() {
		xerror.Assert(service == "", "[service] should not be null")

		var allNodes typex.SMap
		services, err := m.GetService(service)
		xerror.Panic(err)
		for i := range services {
			for _, n := range services[i].Nodes {
				allNodes.Set(n.Id, n)
			}
		}

		watcher.cancel = fx.Tick(func(_ctx fx.Ctx) {
			var nodes typex.SMap
			services, err := m.GetService(service)
			xerror.PanicF(err, "Watch Service %s Error", service)
			for i := range services {
				for _, n := range services[i].Nodes {
					nodes.Set(n.Id, n)
				}
			}

			xerror.Panic(nodes.Each(func(id string, n *registry.Node) {
				if allNodes.Has(id) {
					return
				}

				allNodes.Set(id, n)
				watcher.results <- &registry.Result{
					Action:  registry.Update.String(),
					Service: &registry.Service{Name: service, Nodes: registry.Nodes{n}},
				}
			}))

			xerror.Panic(allNodes.Each(func(id string, n *registry.Node) {
				if nodes.Has(id) {
					return
				}

				allNodes.Delete(id)
				watcher.results <- &registry.Result{
					Action:  registry.Delete.String(),
					Service: &registry.Service{Name: service, Nodes: registry.Nodes{n}},
				}
			}))
		}, m.cfg.TTL)
	})
}

func (m *mdnsRegistry) String() string { return Name }
