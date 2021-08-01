package grpc_entry

import (
	"context"

	"go.uber.org/zap"
	"google.golang.org/grpc"

	"github.com/pubgo/lug"
	"github.com/pubgo/lug/entry"
	"github.com/pubgo/lug/example/grpc_entry/handler"
	"github.com/pubgo/lug/types"
)

var name = "test-grpc"

func GetEntry() entry.Entry {
	ent := lug.NewGrpc(name)
	ent.Description("entry grpc test")
	ent.Register(handler.NewTestAPIHandler())
	ent.UnaryInterceptor(func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		//q.Q(info)
		return handler(ctx, req)
	})

	ent.Middleware(func(next types.MiddleNext) types.MiddleNext {
		return func(ctx context.Context, req types.Request, resp func(rsp types.Response) error) error {
			zap.L().Info("test grpc entry")

			return next(ctx, req, resp)
		}
	})

	return ent
}
