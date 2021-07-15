// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: example/proto/hello/proto3.proto

package hello

import (
	fmt "fmt"
	math "math"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	_ "google.golang.org/protobuf/types/known/wrapperspb"
	_ "google.golang.org/protobuf/types/known/durationpb"
	_ "google.golang.org/protobuf/types/known/anypb"
	_ "google.golang.org/protobuf/types/known/fieldmaskpb"
	time "time"
	github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

func (this *Proto3Message) Validate() error {
	if this.Nested != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Nested); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Nested", err)
		}
	}
	for _, item := range this.RepeatedMessage {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("RepeatedMessage", err)
			}
		}
	}
	if this.TimestampValue != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.TimestampValue); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("TimestampValue", err)
		}
	}
	if this.DurationValue != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.DurationValue); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("DurationValue", err)
		}
	}
	if this.FieldmaskValue != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.FieldmaskValue); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("FieldmaskValue", err)
		}
	}
	if this.WrapperDoubleValue != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.WrapperDoubleValue); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("WrapperDoubleValue", err)
		}
	}
	if this.WrapperFloatValue != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.WrapperFloatValue); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("WrapperFloatValue", err)
		}
	}
	if this.WrapperInt64Value != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.WrapperInt64Value); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("WrapperInt64Value", err)
		}
	}
	if this.WrapperInt32Value != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.WrapperInt32Value); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("WrapperInt32Value", err)
		}
	}
	if this.WrapperUInt64Value != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.WrapperUInt64Value); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("WrapperUInt64Value", err)
		}
	}
	if this.WrapperUInt32Value != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.WrapperUInt32Value); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("WrapperUInt32Value", err)
		}
	}
	if this.WrapperBoolValue != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.WrapperBoolValue); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("WrapperBoolValue", err)
		}
	}
	if this.WrapperStringValue != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.WrapperStringValue); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("WrapperStringValue", err)
		}
	}
	if this.WrapperBytesValue != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.WrapperBytesValue); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("WrapperBytesValue", err)
		}
	}
	// Validation of proto3 map<> fields is unsupported.
	// Validation of proto3 map<> fields is unsupported.
	// Validation of proto3 map<> fields is unsupported.
	// Validation of proto3 map<> fields is unsupported.
	// Validation of proto3 map<> fields is unsupported.
	// Validation of proto3 map<> fields is unsupported.
	// Validation of proto3 map<> fields is unsupported.
	// Validation of proto3 map<> fields is unsupported.
	// Validation of proto3 map<> fields is unsupported.
	// Validation of proto3 map<> fields is unsupported.
	// Validation of proto3 map<> fields is unsupported.
	// Validation of proto3 map<> fields is unsupported.
	// Validation of proto3 map<> fields is unsupported.
	// Validation of proto3 map<> fields is unsupported.
	for _, item := range this.Details {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Details", err)
			}
		}
	}
	return nil
}
