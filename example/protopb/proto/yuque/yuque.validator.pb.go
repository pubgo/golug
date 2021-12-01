// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: proto/yuque/yuque.proto

package yuque_pb

import (
	fmt "fmt"
	math "math"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	_ "github.com/pubgo/lava/proto/lava"
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	_ "google.golang.org/protobuf/types/known/emptypb"
	_ "google.golang.org/protobuf/types/known/wrapperspb"
	github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *CreateGroupResp) Validate() error {
	for _, item := range this.Data {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
			}
		}
	}
	if this.Response != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Response); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Response", err)
		}
	}
	return nil
}
func (this *CreateGroupResp_Data) Validate() error {
	return nil
}
func (this *CreateGroupReq) Validate() error {
	return nil
}
func (this *UserInfoResp) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	if this.Response != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Response); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Response", err)
		}
	}
	return nil
}
func (this *UserInfoResp_Data) Validate() error {
	return nil
}
func (this *UserInfoReq) Validate() error {
	return nil
}
