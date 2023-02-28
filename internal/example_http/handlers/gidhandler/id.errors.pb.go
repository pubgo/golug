package gidhandler

import (
	errors "github.com/pubgo/funk/errors"
	errorpb "github.com/pubgo/funk/proto/errorpb"
	grpc "google.golang.org/grpc"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

var ErrSrvErrCodeOK = errors.NewCode(errorpb.Code_OK).SetName("gid.srv.ok").SetReason("ok").SetStatus(0)
var ErrSrvErrCodeIDGenerateFailed = errors.NewCode(errorpb.Code_Internal).SetName("gid.srv.id_generate_failed").SetReason("id generate error").SetStatus(1)
