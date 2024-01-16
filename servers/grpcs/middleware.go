package grpcs

import (
	"context"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/pubgo/funk/log"
	"reflect"
	"time"

	grpcMiddle "github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/pubgo/funk/convert"
	"github.com/pubgo/funk/errors/errutil"
	"github.com/pubgo/funk/proto/errorpb"
	"github.com/pubgo/funk/strutil"
	"github.com/pubgo/funk/version"
	"github.com/rs/xid"
	"github.com/valyala/fasthttp"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/peer"

	"github.com/pubgo/lava/lava"
	"github.com/pubgo/lava/pkg/grpcutil"
	"github.com/pubgo/lava/pkg/httputil"
	pbv1 "github.com/pubgo/lava/pkg/proto/lava"
)

var parserTypes []fiber.ParserType

func RegParser(customType interface{}, converter func(string) reflect.Value) {
	parserTypes = append(parserTypes, fiber.ParserType{
		Customtype: customType,
		Converter:  converter,
	})
}

func handlerUnaryMiddle(middlewares map[string][]lava.Middleware) grpc.UnaryServerInterceptor {
	unaryWrapper := func(ctx context.Context, req lava.Request) (rsp lava.Response, gErr error) {
		dt, err := req.(*rpcRequest).handler(ctx, req.Payload())
		if err != nil {
			return nil, err
		}

		return &rpcResponse{header: req.(*rpcRequest).rspHeader, dt: dt}, nil
	}

	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		rspMetadata, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			rspMetadata = make(metadata.MD)
		}

		// get content type
		ct := defaultContentType
		if c := rspMetadata.Get("x-content-type"); len(c) != 0 && c[0] != "" {
			ct = c[0]
		}

		if c := rspMetadata.Get("content-type"); len(c) != 0 && c[0] != "" {
			ct = c[0]
		}

		delete(rspMetadata, "x-content-type")

		var clientInfo = new(pbv1.ServiceInfo)

		// get peer from context
		if p := grpcutil.ClientIP(rspMetadata); p != "" {
			clientInfo.Ip = p
		}

		if p := grpcutil.ClientName(rspMetadata); p != "" {
			clientInfo.Name = p
		}

		// timeout for server deadline
		to := rspMetadata.Get("timeout")
		delete(rspMetadata, "timeout")

		// set the timeout if we have it
		if len(to) != 0 && to[0] != "" {
			if dur, err := time.ParseDuration(to[0]); err == nil {
				var cancel context.CancelFunc
				ctx, cancel = context.WithTimeout(ctx, dur)
				_ = cancel
			}
		}

		// 从gateway获取url
		url := info.FullMethod
		if _url, ok := rspMetadata["url"]; ok {
			url = _url[0]
		}

		header := &fasthttp.RequestHeader{}
		for k, v := range rspMetadata {
			for i := range v {
				header.Add(k, v[i])
			}
		}

		srvName := serviceFromMethod(info.FullMethod)
		rpcReq := &rpcRequest{
			service:     srvName,
			method:      info.FullMethod,
			url:         url,
			handler:     handler,
			contentType: ct,
			payload:     req,
			header:      header,
			rspHeader:   new(fasthttp.ResponseHeader),
		}

		reqId := strutil.FirstFnNotEmpty(
			func() string { return lava.GetReqID(ctx) },
			func() string { return string(rpcReq.Header().Peek(httputil.HeaderXRequestID)) },
			func() string { return xid.New().String() },
		)
		ctx = lava.CreateCtxWithReqID(ctx, reqId)

		header.Set(httputil.HeaderXRequestID, reqId)
		header.Set(httputil.HeaderXRequestVersion, version.Version())

		rspMetadata = make(metadata.MD)
		rspMetadata.Set(httputil.HeaderXRequestID, reqId)
		rspMetadata.Set(httputil.HeaderXRequestVersion, version.Version())

		ctx = lava.CreateReqHeader(ctx, header)
		ctx = lava.CreateRspHeader(ctx, rpcReq.rspHeader)
		rsp, err := lava.Chain(middlewares[srvName]...).Middleware(unaryWrapper)(ctx, rpcReq)
		if err != nil {
			pb := errutil.ParseError(err)
			if pb.Trace == nil {
				pb.Trace = new(errorpb.ErrTrace)
			}
			pb.Trace.Operation = rpcReq.Operation()
			pb.Trace.Service = rpcReq.Service()
			pb.Trace.Version = version.Version()

			if pb.Msg != nil {
				pb.Msg = new(errorpb.ErrMsg)
			}
			pb.Msg.Msg = err.Error()
			pb.Msg.Detail = fmt.Sprintf("%#v", err)
			if pb.Msg.Tags == nil {
				pb.Msg.Tags = make(map[string]string)
			}
			pb.Msg.Tags["header"] = string(rpcReq.Header().Header())

			if pb.Code.Reason == "" {
				pb.Code.Reason = err.Error()
			}

			if pb.Code.Code == 0 {
				pb.Code.StatusCode = errorpb.Code_Internal
			}

			if err = grpc.SetTrailer(ctx, rspMetadata); err != nil {
				return nil, err
			}

			return nil, errutil.ConvertErr2Status(pb).Err()
		}

		rsp.Header().VisitAll(func(key, value []byte) {
			rspMetadata.Set(convert.BtoS(key), convert.BtoS(value))
		})

		if err = grpc.SetHeader(ctx, rspMetadata); err != nil {
			log.Err(err, ctx).Msg("grpc set header failed")
		}

		if err = grpc.SetTrailer(ctx, rspMetadata); err != nil {
			log.Err(err, ctx).Msg("grpc set trailer failed")
		}

		return rsp.(*rpcResponse).dt, nil
	}
}

func handlerStreamMiddle(middlewares map[string][]lava.Middleware) grpc.StreamServerInterceptor {
	streamWrapper := func(ctx context.Context, req lava.Request) (lava.Response, error) {
		reqCtx := req.(*rpcRequest)
		wrap := &grpcMiddle.WrappedServerStream{
			WrappedContext: ctx,
			ServerStream:   reqCtx.stream,
		}
		if err := reqCtx.handlerStream(reqCtx.srv, wrap); err != nil {
			return nil, err
		}

		return &rpcResponse{stream: reqCtx.stream, header: new(lava.ResponseHeader)}, nil
	}

	return func(srv interface{}, stream grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
		ctx := stream.Context()
		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			md = make(metadata.MD)
		}

		ct := defaultContentType
		if c := md.Get("x-content-type"); len(c) != 0 && c[0] != "" {
			ct = c[0]
		}

		if c := md.Get("content-type"); len(c) != 0 && c[0] != "" {
			ct = c[0]
		}

		delete(md, "x-content-type")

		// get peer from context
		if p, ok := peer.FromContext(ctx); ok {
			md.Set("remote", p.Addr.String())
		}

		// timeout for server deadline
		to := md.Get("timeout")
		delete(md, "timeout")

		// set the timeout if we have it
		if len(to) != 0 {
			if dur, err := time.ParseDuration(to[0]); err == nil {
				var cancel context.CancelFunc
				ctx, cancel = context.WithTimeout(ctx, dur)
				defer cancel()
			}
		}

		header := new(fasthttp.RequestHeader)
		for k, v := range md {
			for i := range v {
				header.Add(k, v[i])
			}
		}

		srvName := serviceFromMethod(info.FullMethod)
		rpcReq := &rpcRequest{
			stream:        stream,
			srv:           srv,
			handlerStream: handler,
			header:        header,
			method:        info.FullMethod,
			service:       srvName,
			contentType:   ct,
		}

		reqId := strutil.FirstFnNotEmpty(
			func() string { return lava.GetReqID(ctx) },
			func() string { return string(rpcReq.Header().Peek(httputil.HeaderXRequestID)) },
			func() string { return xid.New().String() },
		)
		rpcReq.Header().Set(httputil.HeaderXRequestID, reqId)
		ctx = lava.CreateCtxWithReqID(ctx, reqId)

		rsp, err := lava.Chain(middlewares[srvName]...).Middleware(streamWrapper)(ctx, rpcReq)
		if err != nil {
			pb := errutil.ParseError(err)
			pb.Trace.Operation = rpcReq.Operation()
			pb.Trace.Service = rpcReq.Service()
			pb.Trace.Version = version.Version()
			pb.Msg.Msg = err.Error()
			pb.Msg.Detail = fmt.Sprintf("%#v", err)
			if pb.Msg.Tags == nil {
				pb.Msg.Tags = make(map[string]string)
			}
			pb.Msg.Tags["header"] = string(rpcReq.Header().Header())

			if pb.Code.Reason == "" {
				pb.Code.Reason = err.Error()
			}

			if pb.Code.Code == 0 {
				pb.Code.StatusCode = errorpb.Code_Internal
			}

			return errutil.ConvertErr2Status(pb).Err()
		}

		h := rsp.Header()
		md = make(metadata.MD)
		h.VisitAll(func(key, value []byte) {
			md.Append(convert.BtoS(key), convert.BtoS(value))
		})
		return grpc.SetTrailer(ctx, md)
	}
}

func handlerHttpMiddle(middlewares []lava.Middleware) func(fbCtx *fiber.Ctx) error {
	h := func(ctx context.Context, req lava.Request) (lava.Response, error) {
		reqCtx := req.(*httpRequest)
		reqCtx.ctx.SetUserContext(ctx)
		return &httpResponse{ctx: reqCtx.ctx}, reqCtx.ctx.Next()
	}

	h = lava.Chain(middlewares...).Middleware(h)
	return func(ctx *fiber.Ctx) error {
		_, err := h(ctx.Context(), &httpRequest{ctx: ctx})
		return err
	}
}
