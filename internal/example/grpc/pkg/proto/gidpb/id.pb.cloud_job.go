// Code generated by protoc-gen-cloud-job. DO NOT EDIT.
// versions:
// - protoc-gen-cloud-job v0.0.1
// - protoc                 v5.28.0
// source: gid/id.proto

package gidpb

import (
	"context"
	cloudjobs "github.com/pubgo/lava/component/cloudjobs"
	cloudjobpb "github.com/pubgo/lava/pkg/proto/cloudjobpb"
)

const IdServiceJobKey = "gid"

// IdServiceProxyExecEventKey Id/ProxyExecEvent
const IdServiceProxyExecEventKey = "gid.proxy.exec"

// IdServiceEventChangedKey Id/EventChanged
const IdServiceEventChangedKey = "gid.event.update"

var _ = cloudjobs.RegisterSubject(IdServiceProxyExecEventKey, new(DoProxyEventReq))

func RegisterIdServiceProxyExecEventCloudJob(jobCli *cloudjobs.Client, handler func(ctx *cloudjobs.Context, req *DoProxyEventReq) error, opts ...*cloudjobpb.RegisterJobOptions) {
	cloudjobs.RegisterJobHandler(jobCli, IdServiceJobKey, IdServiceProxyExecEventKey, handler, opts...)
}

func PushIdServiceProxyExecEventCloudJob(jobCli *cloudjobs.Client, ctx context.Context, req *DoProxyEventReq, opts ...*cloudjobpb.PushEventOptions) error {
	return jobCli.Publish(ctx, IdServiceProxyExecEventKey, req, opts...)
}

var _ = cloudjobs.RegisterSubject(IdServiceEventChangedKey, new(DoProxyEventReq))

func RegisterIdServiceEventChangedCloudJob(jobCli *cloudjobs.Client, handler func(ctx *cloudjobs.Context, req *DoProxyEventReq) error, opts ...*cloudjobpb.RegisterJobOptions) {
	cloudjobs.RegisterJobHandler(jobCli, IdServiceJobKey, IdServiceEventChangedKey, handler, opts...)
}

func PushIdServiceEventChangedCloudJob(jobCli *cloudjobs.Client, ctx context.Context, req *DoProxyEventReq, opts ...*cloudjobpb.PushEventOptions) error {
	return jobCli.Publish(ctx, IdServiceEventChangedKey, req, opts...)
}
