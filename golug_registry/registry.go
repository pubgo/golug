// Package registry is an interface for service discovery
package golug_registry

import "errors"

// The registry provides an interface for service discovery
// and an abstraction over varying implementations
// {consul, etcd, zookeeper, ...}
type Registry interface {
	Init(...Option) error
	Options() Options
	Register(*Service, ...RegisterOption) error
	Deregister(*Service) error
	GetService(string) ([]*Service, error)
	ListServices() ([]*Service, error)
	Watch(...WatchOption) (Watcher, error)
	String() string
}

type Option func(*Options)

type RegisterOption func(*RegisterOptions)

type WatchOption func(*WatchOptions)

// Not found error when GetService is called
var ErrNotFound = errors.New("not found")

// Watcher stopped error when watcher is stopped
var ErrWatcherStopped = errors.New("watcher stopped")
