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
// source: policy/v1alpha1/rate_limiting.proto

package v1alpha1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
	v1alpha1 "opensergo.io/api/core/v1alpha1"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// <!-- crd generation tags
// +cue-gen:RateLimiting:groupName:policy.opensergo.io
// +cue-gen:RateLimiting:version:v1alpha1
// +cue-gen:RateLimiting:storageVersion
// +cue-gen:RateLimiting:annotations:helm.sh/resource-policy=keep
// +cue-gen:RateLimiting:subresource:status
// +cue-gen:RateLimiting:scope:Namespaced
// +cue-gen:RateLimiting:resource:categories=opensergo-io
// +cue-gen:RateLimiting:preserveUnknownFields:false
// -->
//
// <!-- go code generation tags
// +kubetype-gen
// +kubetype-gen:groupVersion=policy.opensergo.io/v1alpha1
// +genclient
// +k8s:deepcopy-gen=true
// -->
type RateLimiting struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Selector *v1alpha1.Selector `protobuf:"bytes,1,opt,name=selector,proto3" json:"selector,omitempty"`
	// The config for rate limiting.
	Config *RateLimiting_Config `protobuf:"bytes,2,opt,name=config,proto3" json:"config,omitempty"`
	// The fallback action when the rate limiting works.
	Action *FallbackAction `protobuf:"bytes,3,opt,name=action,proto3" json:"action,omitempty"`
}

func (x *RateLimiting) Reset() {
	*x = RateLimiting{}
	if protoimpl.UnsafeEnabled {
		mi := &file_policy_v1alpha1_rate_limiting_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RateLimiting) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RateLimiting) ProtoMessage() {}

func (x *RateLimiting) ProtoReflect() protoreflect.Message {
	mi := &file_policy_v1alpha1_rate_limiting_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RateLimiting.ProtoReflect.Descriptor instead.
func (*RateLimiting) Descriptor() ([]byte, []int) {
	return file_policy_v1alpha1_rate_limiting_proto_rawDescGZIP(), []int{0}
}

func (x *RateLimiting) GetSelector() *v1alpha1.Selector {
	if x != nil {
		return x.Selector
	}
	return nil
}

func (x *RateLimiting) GetConfig() *RateLimiting_Config {
	if x != nil {
		return x.Config
	}
	return nil
}

func (x *RateLimiting) GetAction() *FallbackAction {
	if x != nil {
		return x.Action
	}
	return nil
}

type RateLimiting_Config struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Scope Scope `protobuf:"varint,1,opt,name=scope,proto3,enum=io.opensergo.policy.v1alpha1.Scope" json:"scope,omitempty"`
	// The max requests can pass within the window.
	Limit uint64 `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
	// The max requests can pass instantly.
	Burst uint64 `protobuf:"varint,3,opt,name=burst,proto3" json:"burst,omitempty"`
	// The window in which the max requests can pass
	Interval *durationpb.Duration `protobuf:"bytes,4,opt,name=interval,proto3" json:"interval,omitempty"`
}

func (x *RateLimiting_Config) Reset() {
	*x = RateLimiting_Config{}
	if protoimpl.UnsafeEnabled {
		mi := &file_policy_v1alpha1_rate_limiting_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RateLimiting_Config) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RateLimiting_Config) ProtoMessage() {}

func (x *RateLimiting_Config) ProtoReflect() protoreflect.Message {
	mi := &file_policy_v1alpha1_rate_limiting_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RateLimiting_Config.ProtoReflect.Descriptor instead.
func (*RateLimiting_Config) Descriptor() ([]byte, []int) {
	return file_policy_v1alpha1_rate_limiting_proto_rawDescGZIP(), []int{0, 0}
}

func (x *RateLimiting_Config) GetScope() Scope {
	if x != nil {
		return x.Scope
	}
	return Scope_LOCAL
}

func (x *RateLimiting_Config) GetLimit() uint64 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *RateLimiting_Config) GetBurst() uint64 {
	if x != nil {
		return x.Burst
	}
	return 0
}

func (x *RateLimiting_Config) GetInterval() *durationpb.Duration {
	if x != nil {
		return x.Interval
	}
	return nil
}

var File_policy_v1alpha1_rate_limiting_proto protoreflect.FileDescriptor

var file_policy_v1alpha1_rate_limiting_proto_rawDesc = []byte{
	0x0a, 0x23, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x31, 0x2f, 0x72, 0x61, 0x74, 0x65, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x69, 0x6e, 0x67, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1c, 0x69, 0x6f, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x73, 0x65,
	0x72, 0x67, 0x6f, 0x2e, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0x1a, 0x18, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x31, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x70,
	0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x75, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8a, 0x03, 0x0a, 0x0c,
	0x52, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x40, 0x0a, 0x08,
	0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24,
	0x2e, 0x69, 0x6f, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x73, 0x65, 0x72, 0x67, 0x6f, 0x2e, 0x63, 0x6f,
	0x72, 0x65, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x53, 0x65, 0x6c, 0x65,
	0x63, 0x74, 0x6f, 0x72, 0x52, 0x08, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x12, 0x49,
	0x0a, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x31,
	0x2e, 0x69, 0x6f, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x73, 0x65, 0x72, 0x67, 0x6f, 0x2e, 0x70, 0x6f,
	0x6c, 0x69, 0x63, 0x79, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x52, 0x61,
	0x74, 0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x52, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x44, 0x0a, 0x06, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x69, 0x6f, 0x2e, 0x6f,
	0x70, 0x65, 0x6e, 0x73, 0x65, 0x72, 0x67, 0x6f, 0x2e, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x46, 0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63,
	0x6b, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x1a,
	0xa6, 0x01, 0x0a, 0x06, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x39, 0x0a, 0x05, 0x73, 0x63,
	0x6f, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x23, 0x2e, 0x69, 0x6f, 0x2e, 0x6f,
	0x70, 0x65, 0x6e, 0x73, 0x65, 0x72, 0x67, 0x6f, 0x2e, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x53, 0x63, 0x6f, 0x70, 0x65, 0x52, 0x05,
	0x73, 0x63, 0x6f, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x62,
	0x75, 0x72, 0x73, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x62, 0x75, 0x72, 0x73,
	0x74, 0x12, 0x35, 0x0a, 0x08, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x08,
	0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x42, 0x22, 0x5a, 0x20, 0x6f, 0x70, 0x65, 0x6e,
	0x73, 0x65, 0x72, 0x67, 0x6f, 0x2e, 0x69, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x6f, 0x6c,
	0x69, 0x63, 0x79, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_policy_v1alpha1_rate_limiting_proto_rawDescOnce sync.Once
	file_policy_v1alpha1_rate_limiting_proto_rawDescData = file_policy_v1alpha1_rate_limiting_proto_rawDesc
)

func file_policy_v1alpha1_rate_limiting_proto_rawDescGZIP() []byte {
	file_policy_v1alpha1_rate_limiting_proto_rawDescOnce.Do(func() {
		file_policy_v1alpha1_rate_limiting_proto_rawDescData = protoimpl.X.CompressGZIP(file_policy_v1alpha1_rate_limiting_proto_rawDescData)
	})
	return file_policy_v1alpha1_rate_limiting_proto_rawDescData
}

var file_policy_v1alpha1_rate_limiting_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_policy_v1alpha1_rate_limiting_proto_goTypes = []interface{}{
	(*RateLimiting)(nil),        // 0: io.opensergo.policy.v1alpha1.RateLimiting
	(*RateLimiting_Config)(nil), // 1: io.opensergo.policy.v1alpha1.RateLimiting.Config
	(*v1alpha1.Selector)(nil),   // 2: io.opensergo.core.v1alpha1.Selector
	(*FallbackAction)(nil),      // 3: io.opensergo.policy.v1alpha1.FallbackAction
	(Scope)(0),                  // 4: io.opensergo.policy.v1alpha1.Scope
	(*durationpb.Duration)(nil), // 5: google.protobuf.Duration
}
var file_policy_v1alpha1_rate_limiting_proto_depIdxs = []int32{
	2, // 0: io.opensergo.policy.v1alpha1.RateLimiting.selector:type_name -> io.opensergo.core.v1alpha1.Selector
	1, // 1: io.opensergo.policy.v1alpha1.RateLimiting.config:type_name -> io.opensergo.policy.v1alpha1.RateLimiting.Config
	3, // 2: io.opensergo.policy.v1alpha1.RateLimiting.action:type_name -> io.opensergo.policy.v1alpha1.FallbackAction
	4, // 3: io.opensergo.policy.v1alpha1.RateLimiting.Config.scope:type_name -> io.opensergo.policy.v1alpha1.Scope
	5, // 4: io.opensergo.policy.v1alpha1.RateLimiting.Config.interval:type_name -> google.protobuf.Duration
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_policy_v1alpha1_rate_limiting_proto_init() }
func file_policy_v1alpha1_rate_limiting_proto_init() {
	if File_policy_v1alpha1_rate_limiting_proto != nil {
		return
	}
	file_policy_v1alpha1_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_policy_v1alpha1_rate_limiting_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RateLimiting); i {
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
		file_policy_v1alpha1_rate_limiting_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RateLimiting_Config); i {
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
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_policy_v1alpha1_rate_limiting_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_policy_v1alpha1_rate_limiting_proto_goTypes,
		DependencyIndexes: file_policy_v1alpha1_rate_limiting_proto_depIdxs,
		MessageInfos:      file_policy_v1alpha1_rate_limiting_proto_msgTypes,
	}.Build()
	File_policy_v1alpha1_rate_limiting_proto = out.File
	file_policy_v1alpha1_rate_limiting_proto_rawDesc = nil
	file_policy_v1alpha1_rate_limiting_proto_goTypes = nil
	file_policy_v1alpha1_rate_limiting_proto_depIdxs = nil
}
