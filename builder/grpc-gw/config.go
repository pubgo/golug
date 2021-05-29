package grpc_gw

import (
	"time"

	gw "github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
)

func init() {
	gw.DefaultContextTimeout = time.Second * 2
}

type ServeMux = gw.ServeMux

type Cfg struct {
	Timeout time.Duration `json:"timeout"`
}

func GetDefaultCfg() Cfg {
	return Cfg{
		Timeout: time.Second * 2,
	}
}
