// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: example/proto/login/login.proto

// 统一登录入口

package login

import (
	fmt "fmt"
	math "math"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	_ "github.com/gogo/protobuf/gogoproto"
	github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *AuthenticateRequest) Validate() error {
	// Validation of proto3 map<> fields is unsupported.
	return nil
}
func (this *AuthenticateResponse) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
func (this *Data) Validate() error {
	return nil
}
func (this *LoginRequest) Validate() error {
	// Validation of proto3 map<> fields is unsupported.
	return nil
}
func (this *LoginResponse) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
func (this *Credentials) Validate() error {
	if this.PlatformInfo != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.PlatformInfo); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("PlatformInfo", err)
		}
	}
	return nil
}
func (this *PlatformInfo) Validate() error {
	return nil
}
