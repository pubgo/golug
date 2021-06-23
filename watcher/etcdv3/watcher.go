package etcdv3

import (
	"context"
	"sync"

	"github.com/pubgo/lug/pkg/typex"
	etcdv32 "github.com/pubgo/lug/plugins/etcdv3"
	"github.com/pubgo/lug/watcher"
	"go.etcd.io/etcd/client/v3"
)

var Name = "etcd"

func init() {
	watcher.Register(Name, func(cfg typex.M) (watcher.Watcher, error) {
		return newWatcher("", ""), nil
	})
}

var _ watcher.Watcher = (*etcdWatcher)(nil)

func newWatcher(prefix string, name string) watcher.Watcher {
	ctx, cancel := context.WithCancel(context.Background())
	return &etcdWatcher{
		name:   name,
		prefix: prefix,
		ctx:    ctx,
		cancel: cancel,
		exitCh: make(chan struct{}, 1),
	}
}

type etcdWatcher struct {
	cancel context.CancelFunc

	name string

	mu     sync.Mutex
	client *clientv3.Client

	ctx context.Context

	closed   bool
	prefix   string
	revision int64
	exitCh   chan struct{}
}

func (w *etcdWatcher) getEtcd() *etcdv32.Client                       { return etcdv32.Get(w.name) }
func (w *etcdWatcher) Close(ctx context.Context, opts ...watcher.Opt) {}
func (w *etcdWatcher) Get(ctx context.Context, key string, opts ...watcher.Opt) ([]*watcher.Response, error) {
	w.getEtcd().Get(ctx, key)
	return nil, nil
}

func (w *etcdWatcher) GetCallback(ctx context.Context, key string, fn func(resp *watcher.Response), opts ...watcher.Opt) error {
	return nil
}

func (w *etcdWatcher) WatchCallback(ctx context.Context, key string, fn func(resp *watcher.Response), opts ...watcher.Opt) {
	go func() {
		for w := range w.getEtcd().Watch(ctx, key) {
			for i := range w.Events {
				var e = w.Events[i]
				fn(&watcher.Response{
					Event:   e.Type.String(),
					Key:     string(e.Kv.Key),
					Value:   e.Kv.Value,
					Version: e.Kv.Version,
				})
			}
		}
	}()
}

func (w *etcdWatcher) Watch(ctx context.Context, key string, opts ...watcher.Opt) <-chan *watcher.Response {
	var resp = make(chan *watcher.Response)
	go func() {
		for w := range etcdv32.Get().Watch(ctx, key) {
			for i := range w.Events {
				var e = w.Events[i]
				resp <- &watcher.Response{
					Event:   e.Type.String(),
					Key:     string(e.Kv.Key),
					Value:   e.Kv.Value,
					Version: e.Kv.Version,
				}
			}
		}
	}()

	return resp
}

func (w *etcdWatcher) Name() string {
	return w.prefix
}
