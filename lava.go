package lava

import (
	"github.com/pubgo/lava/service"
	"github.com/pubgo/lava/service/service_builder"
)

func Run(services ...service.Command) {
	service.Run(services...)
}

func NewSrv(name string, desc ...string) service.Service {
	return service_builder.New(name, desc...)
}
