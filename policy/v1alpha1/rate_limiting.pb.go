// Code generated by protoc-gen-go. DO NOT EDIT.
// source: policy/v1alpha1/rate_limiting.proto

package v1alpha1

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
	math "math"
	v1alpha1 "opensergo.io/api/core/v1alpha1"
	_ "opensergo.io/api/policy/trafficcontrol/v1alpha1"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// <!-- crd generation tags
// +cue-gen:RateLimiting:groupName:policy.opensergo.io
// +cue-gen:RateLimiting:version:v1alpha1
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
	Selector *v1alpha1.Selector `protobuf:"bytes,1,opt,name=selector,proto3" json:"selector,omitempty"`
	// The config for rate limiting.
	Config *RateLimiting_Config `protobuf:"bytes,2,opt,name=config,proto3" json:"config,omitempty"`
	// The fallback action when the rate limiting works.
	Action               *FallbackAction `protobuf:"bytes,3,opt,name=action,proto3" json:"action,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *RateLimiting) Reset()         { *m = RateLimiting{} }
func (m *RateLimiting) String() string { return proto.CompactTextString(m) }
func (*RateLimiting) ProtoMessage()    {}
func (*RateLimiting) Descriptor() ([]byte, []int) {
	return fileDescriptor_52e97190637d71c6, []int{0}
}

func (m *RateLimiting) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RateLimiting.Unmarshal(m, b)
}
func (m *RateLimiting) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RateLimiting.Marshal(b, m, deterministic)
}
func (m *RateLimiting) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RateLimiting.Merge(m, src)
}
func (m *RateLimiting) XXX_Size() int {
	return xxx_messageInfo_RateLimiting.Size(m)
}
func (m *RateLimiting) XXX_DiscardUnknown() {
	xxx_messageInfo_RateLimiting.DiscardUnknown(m)
}

var xxx_messageInfo_RateLimiting proto.InternalMessageInfo

func (m *RateLimiting) GetSelector() *v1alpha1.Selector {
	if m != nil {
		return m.Selector
	}
	return nil
}

func (m *RateLimiting) GetConfig() *RateLimiting_Config {
	if m != nil {
		return m.Config
	}
	return nil
}

func (m *RateLimiting) GetAction() *FallbackAction {
	if m != nil {
		return m.Action
	}
	return nil
}

type RateLimiting_Config struct {
	Scope Scope `protobuf:"varint,1,opt,name=scope,proto3,enum=io.opensergo.policy.v1alpha1.Scope" json:"scope,omitempty"`
	// The max requests can pass within the window.
	Limit uint64 `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
	// The max requests can pass instantly.
	Burst uint64 `protobuf:"varint,3,opt,name=burst,proto3" json:"burst,omitempty"`
	// The window in which the max requests can pass
	Interval             *durationpb.Duration `protobuf:"bytes,4,opt,name=interval,proto3" json:"interval,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *RateLimiting_Config) Reset()         { *m = RateLimiting_Config{} }
func (m *RateLimiting_Config) String() string { return proto.CompactTextString(m) }
func (*RateLimiting_Config) ProtoMessage()    {}
func (*RateLimiting_Config) Descriptor() ([]byte, []int) {
	return fileDescriptor_52e97190637d71c6, []int{0, 0}
}

func (m *RateLimiting_Config) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RateLimiting_Config.Unmarshal(m, b)
}
func (m *RateLimiting_Config) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RateLimiting_Config.Marshal(b, m, deterministic)
}
func (m *RateLimiting_Config) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RateLimiting_Config.Merge(m, src)
}
func (m *RateLimiting_Config) XXX_Size() int {
	return xxx_messageInfo_RateLimiting_Config.Size(m)
}
func (m *RateLimiting_Config) XXX_DiscardUnknown() {
	xxx_messageInfo_RateLimiting_Config.DiscardUnknown(m)
}

var xxx_messageInfo_RateLimiting_Config proto.InternalMessageInfo

func (m *RateLimiting_Config) GetScope() Scope {
	if m != nil {
		return m.Scope
	}
	return Scope_LOCAL
}

func (m *RateLimiting_Config) GetLimit() uint64 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *RateLimiting_Config) GetBurst() uint64 {
	if m != nil {
		return m.Burst
	}
	return 0
}

func (m *RateLimiting_Config) GetInterval() *durationpb.Duration {
	if m != nil {
		return m.Interval
	}
	return nil
}

func init() {
	proto.RegisterType((*RateLimiting)(nil), "io.opensergo.policy.v1alpha1.RateLimiting")
	proto.RegisterType((*RateLimiting_Config)(nil), "io.opensergo.policy.v1alpha1.RateLimiting.Config")
}

func init() {
	proto.RegisterFile("policy/v1alpha1/rate_limiting.proto", fileDescriptor_52e97190637d71c6)
}

var fileDescriptor_52e97190637d71c6 = []byte{
	// 339 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x92, 0x4f, 0x4b, 0xf3, 0x40,
	0x10, 0xc6, 0xe9, 0xdb, 0x36, 0x94, 0x7d, 0xc5, 0xc3, 0xe2, 0x21, 0x96, 0x22, 0xa5, 0xf5, 0x20,
	0x28, 0x1b, 0x5a, 0xf1, 0xe0, 0xcd, 0x3f, 0x45, 0x10, 0x3c, 0x6d, 0x6f, 0x5e, 0x64, 0xb3, 0x6e,
	0xe2, 0xe2, 0x36, 0x13, 0x36, 0xdb, 0x82, 0x5f, 0xc1, 0x0f, 0xe2, 0xe7, 0x94, 0xcc, 0x6e, 0x5a,
	0x2d, 0x98, 0xe3, 0x64, 0xe6, 0xf7, 0xcc, 0x33, 0x79, 0x96, 0x4c, 0x4b, 0x30, 0x5a, 0x7e, 0x24,
	0x9b, 0x99, 0x30, 0xe5, 0x9b, 0x98, 0x25, 0x56, 0x38, 0xf5, 0x62, 0xf4, 0x4a, 0x3b, 0x5d, 0xe4,
	0xac, 0xb4, 0xe0, 0x80, 0x8e, 0x34, 0x30, 0x28, 0x55, 0x51, 0x29, 0x9b, 0x03, 0xf3, 0x04, 0x6b,
	0x88, 0x61, 0x2c, 0xc1, 0xaa, 0x9d, 0x40, 0x5d, 0x79, 0x6e, 0x38, 0xda, 0x17, 0x97, 0xb0, 0x5a,
	0x41, 0x11, 0xba, 0xe7, 0xa1, 0xeb, 0xac, 0xc8, 0x32, 0x2d, 0x25, 0x14, 0xce, 0x82, 0xf9, 0x63,
	0xf8, 0x24, 0x07, 0xc8, 0x8d, 0x4a, 0xb0, 0x4a, 0xd7, 0x59, 0xf2, 0xba, 0xb6, 0xc2, 0xe9, 0xa6,
	0x3f, 0xf9, 0xec, 0x92, 0x03, 0x2e, 0x9c, 0x7a, 0x0a, 0xce, 0xe9, 0x0d, 0x19, 0x54, 0xca, 0x28,
	0xe9, 0xc0, 0xc6, 0x9d, 0x71, 0xe7, 0xec, 0xff, 0xfc, 0x94, 0xfd, 0x3a, 0x03, 0x7d, 0x36, 0xcb,
	0xd8, 0x32, 0xcc, 0xf2, 0x2d, 0x45, 0x1f, 0x49, 0x24, 0xa1, 0xc8, 0x74, 0x1e, 0xff, 0x43, 0x7e,
	0xc6, 0xda, 0x7e, 0x03, 0xfb, 0xb9, 0x9d, 0xdd, 0x23, 0xc8, 0x83, 0x00, 0x5d, 0x90, 0x48, 0xc8,
	0xda, 0x6d, 0xdc, 0x45, 0xa9, 0x8b, 0x76, 0xa9, 0x07, 0x61, 0x4c, 0x2a, 0xe4, 0xfb, 0x2d, 0x32,
	0x3c, 0xb0, 0xc3, 0xaf, 0x0e, 0x89, 0xbc, 0x30, 0xbd, 0x26, 0xfd, 0x4a, 0x42, 0xa9, 0xf0, 0xb4,
	0xc3, 0xf9, 0xb4, 0x5d, 0x6f, 0x59, 0x8f, 0x72, 0x4f, 0xd0, 0x23, 0xd2, 0xc7, 0x78, 0xf1, 0xaa,
	0x1e, 0xf7, 0x45, 0xfd, 0x35, 0x5d, 0xdb, 0xca, 0xa1, 0xc1, 0x1e, 0xf7, 0x05, 0xbd, 0x22, 0x03,
	0x5d, 0x38, 0x65, 0x37, 0xc2, 0xc4, 0x3d, 0x74, 0x7e, 0xcc, 0x7c, 0x10, 0xac, 0x09, 0x82, 0x2d,
	0x42, 0x10, 0x7c, 0x3b, 0x7a, 0x37, 0x79, 0x1e, 0xef, 0xcc, 0x68, 0x48, 0x44, 0xa9, 0x93, 0xbd,
	0xa7, 0x90, 0x46, 0x28, 0x70, 0xf9, 0x1d, 0x00, 0x00, 0xff, 0xff, 0x6c, 0x2f, 0x8b, 0x17, 0x81,
	0x02, 0x00, 0x00,
}
