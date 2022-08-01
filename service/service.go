package service

import (
	_ "github.com/go-chi/chi/v5"
	"github.com/urfave/cli/v2"
	"google.golang.org/grpc"
)

type Init interface {
	Init()
}

type Close interface {
	Close()
}

type Flags interface {
	Flags() []cli.Flag
}

type Options struct {
	Id        string            `json:"id,omitempty"`
	Name      string            `json:"name,omitempty"`
	Version   string            `json:"version,omitempty"`
	Port      int               `json:"port,omitempty"`
	Addr      string            `json:"addr,omitempty"`
	Advertise string            `json:"advertise,omitempty"`
	Metadata  map[string]string `json:"metadata,omitempty"`
}

type Runtime interface {
	Start()
	Stop()
	Run()
}

type Service interface {
	Runtime
	grpc.ServiceRegistrar
	Provider(provider interface{})
	RegisterGateway(register ...GatewayRegister)
}
