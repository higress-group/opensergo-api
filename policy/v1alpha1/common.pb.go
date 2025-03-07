// Copyright 2023, OpenSergo Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        (unknown)
// source: policy/v1alpha1/common.proto

package v1alpha1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Scope int32

const (
	// Single instance.
	Scope_LOCAL Scope = 0
	// The whole cluster.
	Scope_GLOBAL          Scope = 1
	Scope_GLOBAL_TO_LOCAL Scope = 2
)

// Enum value maps for Scope.
var (
	Scope_name = map[int32]string{
		0: "LOCAL",
		1: "GLOBAL",
		2: "GLOBAL_TO_LOCAL",
	}
	Scope_value = map[string]int32{
		"LOCAL":           0,
		"GLOBAL":          1,
		"GLOBAL_TO_LOCAL": 2,
	}
)

func (x Scope) Enum() *Scope {
	p := new(Scope)
	*p = x
	return p
}

func (x Scope) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Scope) Descriptor() protoreflect.EnumDescriptor {
	return file_policy_v1alpha1_common_proto_enumTypes[0].Descriptor()
}

func (Scope) Type() protoreflect.EnumType {
	return &file_policy_v1alpha1_common_proto_enumTypes[0]
}

func (x Scope) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Scope.Descriptor instead.
func (Scope) EnumDescriptor() ([]byte, []int) {
	return file_policy_v1alpha1_common_proto_rawDescGZIP(), []int{0}
}

type FallbackAction_HTTPFallback_HTTPCustomResponse_BodyEncoding int32

const (
	FallbackAction_HTTPFallback_HTTPCustomResponse_TEXT FallbackAction_HTTPFallback_HTTPCustomResponse_BodyEncoding = 0
	FallbackAction_HTTPFallback_HTTPCustomResponse_JSON FallbackAction_HTTPFallback_HTTPCustomResponse_BodyEncoding = 1
)

// Enum value maps for FallbackAction_HTTPFallback_HTTPCustomResponse_BodyEncoding.
var (
	FallbackAction_HTTPFallback_HTTPCustomResponse_BodyEncoding_name = map[int32]string{
		0: "TEXT",
		1: "JSON",
	}
	FallbackAction_HTTPFallback_HTTPCustomResponse_BodyEncoding_value = map[string]int32{
		"TEXT": 0,
		"JSON": 1,
	}
)

func (x FallbackAction_HTTPFallback_HTTPCustomResponse_BodyEncoding) Enum() *FallbackAction_HTTPFallback_HTTPCustomResponse_BodyEncoding {
	p := new(FallbackAction_HTTPFallback_HTTPCustomResponse_BodyEncoding)
	*p = x
	return p
}

func (x FallbackAction_HTTPFallback_HTTPCustomResponse_BodyEncoding) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (FallbackAction_HTTPFallback_HTTPCustomResponse_BodyEncoding) Descriptor() protoreflect.EnumDescriptor {
	return file_policy_v1alpha1_common_proto_enumTypes[1].Descriptor()
}

func (FallbackAction_HTTPFallback_HTTPCustomResponse_BodyEncoding) Type() protoreflect.EnumType {
	return &file_policy_v1alpha1_common_proto_enumTypes[1]
}

func (x FallbackAction_HTTPFallback_HTTPCustomResponse_BodyEncoding) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use FallbackAction_HTTPFallback_HTTPCustomResponse_BodyEncoding.Descriptor instead.
func (FallbackAction_HTTPFallback_HTTPCustomResponse_BodyEncoding) EnumDescriptor() ([]byte, []int) {
	return file_policy_v1alpha1_common_proto_rawDescGZIP(), []int{0, 0, 0, 0}
}

type FallbackAction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to FallbackType:
	//
	//	*FallbackAction_Http
	FallbackType isFallbackAction_FallbackType `protobuf_oneof:"fallback_type"`
}

func (x *FallbackAction) Reset() {
	*x = FallbackAction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_policy_v1alpha1_common_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FallbackAction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FallbackAction) ProtoMessage() {}

func (x *FallbackAction) ProtoReflect() protoreflect.Message {
	mi := &file_policy_v1alpha1_common_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FallbackAction.ProtoReflect.Descriptor instead.
func (*FallbackAction) Descriptor() ([]byte, []int) {
	return file_policy_v1alpha1_common_proto_rawDescGZIP(), []int{0}
}

func (m *FallbackAction) GetFallbackType() isFallbackAction_FallbackType {
	if m != nil {
		return m.FallbackType
	}
	return nil
}

func (x *FallbackAction) GetHttp() *FallbackAction_HTTPFallback {
	if x, ok := x.GetFallbackType().(*FallbackAction_Http); ok {
		return x.Http
	}
	return nil
}

type isFallbackAction_FallbackType interface {
	isFallbackAction_FallbackType()
}

type FallbackAction_Http struct {
	// Fallback for http
	Http *FallbackAction_HTTPFallback `protobuf:"bytes,1,opt,name=http,proto3,oneof"`
}

func (*FallbackAction_Http) isFallbackAction_FallbackType() {}

type FallbackAction_HTTPFallback struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to ActionType:
	//
	//	*FallbackAction_HTTPFallback_CustomResponse
	//	*FallbackAction_HTTPFallback_Redirect
	ActionType isFallbackAction_HTTPFallback_ActionType `protobuf_oneof:"action_type"`
}

func (x *FallbackAction_HTTPFallback) Reset() {
	*x = FallbackAction_HTTPFallback{}
	if protoimpl.UnsafeEnabled {
		mi := &file_policy_v1alpha1_common_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FallbackAction_HTTPFallback) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FallbackAction_HTTPFallback) ProtoMessage() {}

func (x *FallbackAction_HTTPFallback) ProtoReflect() protoreflect.Message {
	mi := &file_policy_v1alpha1_common_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FallbackAction_HTTPFallback.ProtoReflect.Descriptor instead.
func (*FallbackAction_HTTPFallback) Descriptor() ([]byte, []int) {
	return file_policy_v1alpha1_common_proto_rawDescGZIP(), []int{0, 0}
}

func (m *FallbackAction_HTTPFallback) GetActionType() isFallbackAction_HTTPFallback_ActionType {
	if m != nil {
		return m.ActionType
	}
	return nil
}

func (x *FallbackAction_HTTPFallback) GetCustomResponse() *FallbackAction_HTTPFallback_HTTPCustomResponse {
	if x, ok := x.GetActionType().(*FallbackAction_HTTPFallback_CustomResponse); ok {
		return x.CustomResponse
	}
	return nil
}

func (x *FallbackAction_HTTPFallback) GetRedirect() *FallbackAction_HTTPFallback_HTTPRedirect {
	if x, ok := x.GetActionType().(*FallbackAction_HTTPFallback_Redirect); ok {
		return x.Redirect
	}
	return nil
}

type isFallbackAction_HTTPFallback_ActionType interface {
	isFallbackAction_HTTPFallback_ActionType()
}

type FallbackAction_HTTPFallback_CustomResponse struct {
	CustomResponse *FallbackAction_HTTPFallback_HTTPCustomResponse `protobuf:"bytes,1,opt,name=custom_response,json=customResponse,proto3,oneof"`
}

type FallbackAction_HTTPFallback_Redirect struct {
	Redirect *FallbackAction_HTTPFallback_HTTPRedirect `protobuf:"bytes,2,opt,name=redirect,proto3,oneof"`
}

func (*FallbackAction_HTTPFallback_CustomResponse) isFallbackAction_HTTPFallback_ActionType() {}

func (*FallbackAction_HTTPFallback_Redirect) isFallbackAction_HTTPFallback_ActionType() {}

type FallbackAction_HTTPFallback_HTTPCustomResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StatusCode           uint32                                                      `protobuf:"varint,1,opt,name=status_code,json=statusCode,proto3" json:"status_code,omitempty"`
	BodyEncoding         FallbackAction_HTTPFallback_HTTPCustomResponse_BodyEncoding `protobuf:"varint,2,opt,name=body_encoding,json=bodyEncoding,proto3,enum=io.opensergo.policy.v1alpha1.FallbackAction_HTTPFallback_HTTPCustomResponse_BodyEncoding" json:"body_encoding,omitempty"`
	Body                 string                                                      `protobuf:"bytes,3,opt,name=body,proto3" json:"body,omitempty"`
	ResponseHeadersToAdd map[string]string                                           `protobuf:"bytes,4,rep,name=response_headers_to_add,json=responseHeadersToAdd,proto3" json:"response_headers_to_add,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *FallbackAction_HTTPFallback_HTTPCustomResponse) Reset() {
	*x = FallbackAction_HTTPFallback_HTTPCustomResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_policy_v1alpha1_common_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FallbackAction_HTTPFallback_HTTPCustomResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FallbackAction_HTTPFallback_HTTPCustomResponse) ProtoMessage() {}

func (x *FallbackAction_HTTPFallback_HTTPCustomResponse) ProtoReflect() protoreflect.Message {
	mi := &file_policy_v1alpha1_common_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FallbackAction_HTTPFallback_HTTPCustomResponse.ProtoReflect.Descriptor instead.
func (*FallbackAction_HTTPFallback_HTTPCustomResponse) Descriptor() ([]byte, []int) {
	return file_policy_v1alpha1_common_proto_rawDescGZIP(), []int{0, 0, 0}
}

func (x *FallbackAction_HTTPFallback_HTTPCustomResponse) GetStatusCode() uint32 {
	if x != nil {
		return x.StatusCode
	}
	return 0
}

func (x *FallbackAction_HTTPFallback_HTTPCustomResponse) GetBodyEncoding() FallbackAction_HTTPFallback_HTTPCustomResponse_BodyEncoding {
	if x != nil {
		return x.BodyEncoding
	}
	return FallbackAction_HTTPFallback_HTTPCustomResponse_TEXT
}

func (x *FallbackAction_HTTPFallback_HTTPCustomResponse) GetBody() string {
	if x != nil {
		return x.Body
	}
	return ""
}

func (x *FallbackAction_HTTPFallback_HTTPCustomResponse) GetResponseHeadersToAdd() map[string]string {
	if x != nil {
		return x.ResponseHeadersToAdd
	}
	return nil
}

type FallbackAction_HTTPFallback_HTTPRedirect struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RedirectUrl string `protobuf:"bytes,1,opt,name=redirect_url,json=redirectUrl,proto3" json:"redirect_url,omitempty"`
}

func (x *FallbackAction_HTTPFallback_HTTPRedirect) Reset() {
	*x = FallbackAction_HTTPFallback_HTTPRedirect{}
	if protoimpl.UnsafeEnabled {
		mi := &file_policy_v1alpha1_common_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FallbackAction_HTTPFallback_HTTPRedirect) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FallbackAction_HTTPFallback_HTTPRedirect) ProtoMessage() {}

func (x *FallbackAction_HTTPFallback_HTTPRedirect) ProtoReflect() protoreflect.Message {
	mi := &file_policy_v1alpha1_common_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FallbackAction_HTTPFallback_HTTPRedirect.ProtoReflect.Descriptor instead.
func (*FallbackAction_HTTPFallback_HTTPRedirect) Descriptor() ([]byte, []int) {
	return file_policy_v1alpha1_common_proto_rawDescGZIP(), []int{0, 0, 1}
}

func (x *FallbackAction_HTTPFallback_HTTPRedirect) GetRedirectUrl() string {
	if x != nil {
		return x.RedirectUrl
	}
	return ""
}

var File_policy_v1alpha1_common_proto protoreflect.FileDescriptor

var file_policy_v1alpha1_common_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1c,
	0x69, 0x6f, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x73, 0x65, 0x72, 0x67, 0x6f, 0x2e, 0x70, 0x6f, 0x6c,
	0x69, 0x63, 0x79, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x22, 0xfd, 0x06, 0x0a,
	0x0e, 0x46, 0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x4f, 0x0a, 0x04, 0x68, 0x74, 0x74, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x39, 0x2e,
	0x69, 0x6f, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x73, 0x65, 0x72, 0x67, 0x6f, 0x2e, 0x70, 0x6f, 0x6c,
	0x69, 0x63, 0x79, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x46, 0x61, 0x6c,
	0x6c, 0x62, 0x61, 0x63, 0x6b, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x48, 0x54, 0x54, 0x50,
	0x46, 0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x48, 0x00, 0x52, 0x04, 0x68, 0x74, 0x74, 0x70,
	0x1a, 0x88, 0x06, 0x0a, 0x0c, 0x48, 0x54, 0x54, 0x50, 0x46, 0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63,
	0x6b, 0x12, 0x77, 0x0a, 0x0f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x5f, 0x72, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x4c, 0x2e, 0x69, 0x6f, 0x2e,
	0x6f, 0x70, 0x65, 0x6e, 0x73, 0x65, 0x72, 0x67, 0x6f, 0x2e, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79,
	0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x46, 0x61, 0x6c, 0x6c, 0x62, 0x61,
	0x63, 0x6b, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x48, 0x54, 0x54, 0x50, 0x46, 0x61, 0x6c,
	0x6c, 0x62, 0x61, 0x63, 0x6b, 0x2e, 0x48, 0x54, 0x54, 0x50, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x48, 0x00, 0x52, 0x0e, 0x63, 0x75, 0x73, 0x74,
	0x6f, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x64, 0x0a, 0x08, 0x72, 0x65,
	0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x46, 0x2e, 0x69,
	0x6f, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x73, 0x65, 0x72, 0x67, 0x6f, 0x2e, 0x70, 0x6f, 0x6c, 0x69,
	0x63, 0x79, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x46, 0x61, 0x6c, 0x6c,
	0x62, 0x61, 0x63, 0x6b, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x48, 0x54, 0x54, 0x50, 0x46,
	0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x2e, 0x48, 0x54, 0x54, 0x50, 0x52, 0x65, 0x64, 0x69,
	0x72, 0x65, 0x63, 0x74, 0x48, 0x00, 0x52, 0x08, 0x72, 0x65, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74,
	0x1a, 0xd6, 0x03, 0x0a, 0x12, 0x48, 0x54, 0x54, 0x50, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x7e, 0x0a, 0x0d, 0x62, 0x6f, 0x64, 0x79,
	0x5f, 0x65, 0x6e, 0x63, 0x6f, 0x64, 0x69, 0x6e, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x59, 0x2e, 0x69, 0x6f, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x73, 0x65, 0x72, 0x67, 0x6f, 0x2e, 0x70,
	0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x46,
	0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x48, 0x54,
	0x54, 0x50, 0x46, 0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x2e, 0x48, 0x54, 0x54, 0x50, 0x43,
	0x75, 0x73, 0x74, 0x6f, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x42, 0x6f,
	0x64, 0x79, 0x45, 0x6e, 0x63, 0x6f, 0x64, 0x69, 0x6e, 0x67, 0x52, 0x0c, 0x62, 0x6f, 0x64, 0x79,
	0x45, 0x6e, 0x63, 0x6f, 0x64, 0x69, 0x6e, 0x67, 0x12, 0x12, 0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x12, 0x9d, 0x01, 0x0a,
	0x17, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72,
	0x73, 0x5f, 0x74, 0x6f, 0x5f, 0x61, 0x64, 0x64, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x66,
	0x2e, 0x69, 0x6f, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x73, 0x65, 0x72, 0x67, 0x6f, 0x2e, 0x70, 0x6f,
	0x6c, 0x69, 0x63, 0x79, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x46, 0x61,
	0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x48, 0x54, 0x54,
	0x50, 0x46, 0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x2e, 0x48, 0x54, 0x54, 0x50, 0x43, 0x75,
	0x73, 0x74, 0x6f, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x54, 0x6f, 0x41, 0x64,
	0x64, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x14, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x54, 0x6f, 0x41, 0x64, 0x64, 0x1a, 0x47, 0x0a, 0x19,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x54,
	0x6f, 0x41, 0x64, 0x64, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x22, 0x0a, 0x0c, 0x42, 0x6f, 0x64, 0x79, 0x45, 0x6e, 0x63,
	0x6f, 0x64, 0x69, 0x6e, 0x67, 0x12, 0x08, 0x0a, 0x04, 0x54, 0x45, 0x58, 0x54, 0x10, 0x00, 0x12,
	0x08, 0x0a, 0x04, 0x4a, 0x53, 0x4f, 0x4e, 0x10, 0x01, 0x1a, 0x31, 0x0a, 0x0c, 0x48, 0x54, 0x54,
	0x50, 0x52, 0x65, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x72, 0x65, 0x64,
	0x69, 0x72, 0x65, 0x63, 0x74, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x72, 0x65, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x55, 0x72, 0x6c, 0x42, 0x0d, 0x0a, 0x0b,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x42, 0x0f, 0x0a, 0x0d, 0x66,
	0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x2a, 0x33, 0x0a, 0x05,
	0x53, 0x63, 0x6f, 0x70, 0x65, 0x12, 0x09, 0x0a, 0x05, 0x4c, 0x4f, 0x43, 0x41, 0x4c, 0x10, 0x00,
	0x12, 0x0a, 0x0a, 0x06, 0x47, 0x4c, 0x4f, 0x42, 0x41, 0x4c, 0x10, 0x01, 0x12, 0x13, 0x0a, 0x0f,
	0x47, 0x4c, 0x4f, 0x42, 0x41, 0x4c, 0x5f, 0x54, 0x4f, 0x5f, 0x4c, 0x4f, 0x43, 0x41, 0x4c, 0x10,
	0x02, 0x42, 0x22, 0x5a, 0x20, 0x6f, 0x70, 0x65, 0x6e, 0x73, 0x65, 0x72, 0x67, 0x6f, 0x2e, 0x69,
	0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2f, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_policy_v1alpha1_common_proto_rawDescOnce sync.Once
	file_policy_v1alpha1_common_proto_rawDescData = file_policy_v1alpha1_common_proto_rawDesc
)

func file_policy_v1alpha1_common_proto_rawDescGZIP() []byte {
	file_policy_v1alpha1_common_proto_rawDescOnce.Do(func() {
		file_policy_v1alpha1_common_proto_rawDescData = protoimpl.X.CompressGZIP(file_policy_v1alpha1_common_proto_rawDescData)
	})
	return file_policy_v1alpha1_common_proto_rawDescData
}

var file_policy_v1alpha1_common_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_policy_v1alpha1_common_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_policy_v1alpha1_common_proto_goTypes = []interface{}{
	(Scope)(0), // 0: io.opensergo.policy.v1alpha1.Scope
	(FallbackAction_HTTPFallback_HTTPCustomResponse_BodyEncoding)(0), // 1: io.opensergo.policy.v1alpha1.FallbackAction.HTTPFallback.HTTPCustomResponse.BodyEncoding
	(*FallbackAction)(nil),                                 // 2: io.opensergo.policy.v1alpha1.FallbackAction
	(*FallbackAction_HTTPFallback)(nil),                    // 3: io.opensergo.policy.v1alpha1.FallbackAction.HTTPFallback
	(*FallbackAction_HTTPFallback_HTTPCustomResponse)(nil), // 4: io.opensergo.policy.v1alpha1.FallbackAction.HTTPFallback.HTTPCustomResponse
	(*FallbackAction_HTTPFallback_HTTPRedirect)(nil),       // 5: io.opensergo.policy.v1alpha1.FallbackAction.HTTPFallback.HTTPRedirect
	nil, // 6: io.opensergo.policy.v1alpha1.FallbackAction.HTTPFallback.HTTPCustomResponse.ResponseHeadersToAddEntry
}
var file_policy_v1alpha1_common_proto_depIdxs = []int32{
	3, // 0: io.opensergo.policy.v1alpha1.FallbackAction.http:type_name -> io.opensergo.policy.v1alpha1.FallbackAction.HTTPFallback
	4, // 1: io.opensergo.policy.v1alpha1.FallbackAction.HTTPFallback.custom_response:type_name -> io.opensergo.policy.v1alpha1.FallbackAction.HTTPFallback.HTTPCustomResponse
	5, // 2: io.opensergo.policy.v1alpha1.FallbackAction.HTTPFallback.redirect:type_name -> io.opensergo.policy.v1alpha1.FallbackAction.HTTPFallback.HTTPRedirect
	1, // 3: io.opensergo.policy.v1alpha1.FallbackAction.HTTPFallback.HTTPCustomResponse.body_encoding:type_name -> io.opensergo.policy.v1alpha1.FallbackAction.HTTPFallback.HTTPCustomResponse.BodyEncoding
	6, // 4: io.opensergo.policy.v1alpha1.FallbackAction.HTTPFallback.HTTPCustomResponse.response_headers_to_add:type_name -> io.opensergo.policy.v1alpha1.FallbackAction.HTTPFallback.HTTPCustomResponse.ResponseHeadersToAddEntry
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_policy_v1alpha1_common_proto_init() }
func file_policy_v1alpha1_common_proto_init() {
	if File_policy_v1alpha1_common_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_policy_v1alpha1_common_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FallbackAction); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_policy_v1alpha1_common_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FallbackAction_HTTPFallback); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_policy_v1alpha1_common_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FallbackAction_HTTPFallback_HTTPCustomResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_policy_v1alpha1_common_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FallbackAction_HTTPFallback_HTTPRedirect); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_policy_v1alpha1_common_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*FallbackAction_Http)(nil),
	}
	file_policy_v1alpha1_common_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*FallbackAction_HTTPFallback_CustomResponse)(nil),
		(*FallbackAction_HTTPFallback_Redirect)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_policy_v1alpha1_common_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_policy_v1alpha1_common_proto_goTypes,
		DependencyIndexes: file_policy_v1alpha1_common_proto_depIdxs,
		EnumInfos:         file_policy_v1alpha1_common_proto_enumTypes,
		MessageInfos:      file_policy_v1alpha1_common_proto_msgTypes,
	}.Build()
	File_policy_v1alpha1_common_proto = out.File
	file_policy_v1alpha1_common_proto_rawDesc = nil
	file_policy_v1alpha1_common_proto_goTypes = nil
	file_policy_v1alpha1_common_proto_depIdxs = nil
}
