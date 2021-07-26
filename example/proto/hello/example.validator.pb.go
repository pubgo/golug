// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: example/proto/hello/example.proto

// GoGo gRPC Example
//
// This example is used to show how to use gRPC and
// gRPC-Gateway with GoGo Protobuf.

package hello

import (
	fmt "fmt"
	math "math"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	_ "google.golang.org/protobuf/types/known/durationpb"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	_ "github.com/gogo/protobuf/gogoproto"
	_ "github.com/mwitkow/go-proto-validators"
	_ "google.golang.org/protobuf/types/known/emptypb"
	_ "google.golang.org/protobuf/types/known/fieldmaskpb"
	_ "google.golang.org/protobuf/types/known/structpb"
	time "time"
	github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

func (this *User) Validate() error {
	if !(this.Id > 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("Id", fmt.Errorf(`ID must be a positive integer`))
	}
	if this.CreateDate != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.CreateDate); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("CreateDate", err)
		}
	}
	return nil
}
func (this *UserRole) Validate() error {
	if this.Lists != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Lists); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Lists", err)
		}
	}
	// Validation of proto3 map<> fields is unsupported.
	return nil
}
func (this *UpdateUserRequest) Validate() error {
	if this.User != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.User); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("User", err)
		}
	}
	if this.UpdateMask != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.UpdateMask); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("UpdateMask", err)
		}
	}
	return nil
}
func (this *ListUsersRequest) Validate() error {
	if this.CreatedSince != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.CreatedSince); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("CreatedSince", err)
		}
	}
	if this.OlderThan != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.OlderThan); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("OlderThan", err)
		}
	}
	return nil
}
