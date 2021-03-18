package mux

import (
	"github.com/pubgo/golug/server/chi_srv"
)

const Name = "mux"

var cfg = GetDefaultCfg()

type Cfg struct {
	chi_srv.Cfg
	Port int
}

func GetDefaultCfg() Cfg {
	return Cfg{
		Cfg:  chi_srv.GetDefaultCfg(),
		Port: 8088,
	}
}
