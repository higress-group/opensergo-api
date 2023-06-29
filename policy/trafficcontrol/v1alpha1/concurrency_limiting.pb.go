// Code generated by protoc-gen-go. DO NOT EDIT.
// source: policy/trafficcontrol/v1alpha1/concurrency_limiting.proto

// $schema: io.opensergo.policy.trafficcontrol.v1alpha1.ConcurrencyLimiting
// $title: Concurrency Limiting
// $description: Concurrency limiting rule.

package v1alpha1

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
	v1alpha1 "opensergo.io/api/core/v1alpha1"
	v1alpha11 "opensergo.io/api/policy/v1alpha1"
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
// +cue-gen:ConcurrencyLimiting:groupName:trafficcontrol.policy.istio.io
// +cue-gen:ConcurrencyLimiting:version:v1alpha1
// +cue-gen:ConcurrencyLimiting:annotations:helm.sh/resource-policy=keep
// +cue-gen:ConcurrencyLimiting:subresource:status
// +cue-gen:ConcurrencyLimiting:scope:Namespaced
// +cue-gen:ConcurrencyLimiting:resource:categories=opensergo-io
// +cue-gen:ConcurrencyLimiting:preserveUnknownFields:false
// -->
//
// <!-- go code generation tags
// +kubetype-gen
// +kubetype-gen:groupVersion=trafficcontrol.policy.istio.io/v1alpha1
// +genclient
// +k8s:deepcopy-gen=true
// -->
type ConcurrencyLimiting struct {
	Selector *v1alpha1.Selector `protobuf:"bytes,1,opt,name=selector,proto3" json:"selector,omitempty"`
	// The config for concurrency limiting.
	Config *ConcurrencyLimiting_Config `protobuf:"bytes,2,opt,name=config,proto3" json:"config,omitempty"`
	// The fallback action when the concurrency limiting works.
	Action               *v1alpha11.FallbackAction `protobuf:"bytes,3,opt,name=action,proto3" json:"action,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *ConcurrencyLimiting) Reset()         { *m = ConcurrencyLimiting{} }
func (m *ConcurrencyLimiting) String() string { return proto.CompactTextString(m) }
func (*ConcurrencyLimiting) ProtoMessage()    {}
func (*ConcurrencyLimiting) Descriptor() ([]byte, []int) {
	return fileDescriptor_48cc870c2a6811dc, []int{0}
}

func (m *ConcurrencyLimiting) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConcurrencyLimiting.Unmarshal(m, b)
}
func (m *ConcurrencyLimiting) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConcurrencyLimiting.Marshal(b, m, deterministic)
}
func (m *ConcurrencyLimiting) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConcurrencyLimiting.Merge(m, src)
}
func (m *ConcurrencyLimiting) XXX_Size() int {
	return xxx_messageInfo_ConcurrencyLimiting.Size(m)
}
func (m *ConcurrencyLimiting) XXX_DiscardUnknown() {
	xxx_messageInfo_ConcurrencyLimiting.DiscardUnknown(m)
}

var xxx_messageInfo_ConcurrencyLimiting proto.InternalMessageInfo

func (m *ConcurrencyLimiting) GetSelector() *v1alpha1.Selector {
	if m != nil {
		return m.Selector
	}
	return nil
}

func (m *ConcurrencyLimiting) GetConfig() *ConcurrencyLimiting_Config {
	if m != nil {
		return m.Config
	}
	return nil
}

func (m *ConcurrencyLimiting) GetAction() *v1alpha11.FallbackAction {
	if m != nil {
		return m.Action
	}
	return nil
}

type ConcurrencyLimiting_Config struct {
	Scope                Scope    `protobuf:"varint,1,opt,name=scope,proto3,enum=io.opensergo.policy.trafficcontrol.v1alpha1.Scope" json:"scope,omitempty"`
	MaxConcurrency       uint64   `protobuf:"varint,2,opt,name=max_concurrency,json=maxConcurrency,proto3" json:"max_concurrency,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ConcurrencyLimiting_Config) Reset()         { *m = ConcurrencyLimiting_Config{} }
func (m *ConcurrencyLimiting_Config) String() string { return proto.CompactTextString(m) }
func (*ConcurrencyLimiting_Config) ProtoMessage()    {}
func (*ConcurrencyLimiting_Config) Descriptor() ([]byte, []int) {
	return fileDescriptor_48cc870c2a6811dc, []int{0, 0}
}

func (m *ConcurrencyLimiting_Config) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConcurrencyLimiting_Config.Unmarshal(m, b)
}
func (m *ConcurrencyLimiting_Config) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConcurrencyLimiting_Config.Marshal(b, m, deterministic)
}
func (m *ConcurrencyLimiting_Config) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConcurrencyLimiting_Config.Merge(m, src)
}
func (m *ConcurrencyLimiting_Config) XXX_Size() int {
	return xxx_messageInfo_ConcurrencyLimiting_Config.Size(m)
}
func (m *ConcurrencyLimiting_Config) XXX_DiscardUnknown() {
	xxx_messageInfo_ConcurrencyLimiting_Config.DiscardUnknown(m)
}

var xxx_messageInfo_ConcurrencyLimiting_Config proto.InternalMessageInfo

func (m *ConcurrencyLimiting_Config) GetScope() Scope {
	if m != nil {
		return m.Scope
	}
	return Scope_LOCAL
}

func (m *ConcurrencyLimiting_Config) GetMaxConcurrency() uint64 {
	if m != nil {
		return m.MaxConcurrency
	}
	return 0
}

func init() {
	proto.RegisterType((*ConcurrencyLimiting)(nil), "io.opensergo.policy.trafficcontrol.v1alpha1.ConcurrencyLimiting")
	proto.RegisterType((*ConcurrencyLimiting_Config)(nil), "io.opensergo.policy.trafficcontrol.v1alpha1.ConcurrencyLimiting.Config")
}

func init() {
	proto.RegisterFile("policy/trafficcontrol/v1alpha1/concurrency_limiting.proto", fileDescriptor_48cc870c2a6811dc)
}

var fileDescriptor_48cc870c2a6811dc = []byte{
	// 305 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0x4f, 0x4b, 0x33, 0x31,
	0x10, 0x87, 0x69, 0xdf, 0xd7, 0x45, 0x22, 0x54, 0x88, 0x97, 0xa5, 0x78, 0x10, 0x11, 0x14, 0x2a,
	0x09, 0xad, 0x27, 0x6f, 0xfe, 0x43, 0x3d, 0x78, 0xda, 0xde, 0xbc, 0x2c, 0xe9, 0x90, 0x5d, 0x83,
	0xd9, 0x4c, 0xc8, 0x46, 0x69, 0xf1, 0xa3, 0x7b, 0x91, 0x66, 0xd3, 0xee, 0x16, 0x8a, 0xd2, 0xeb,
	0xce, 0xfc, 0x9e, 0x79, 0x66, 0x27, 0xe4, 0xda, 0xa2, 0x56, 0xb0, 0xe0, 0xde, 0x89, 0xa2, 0x50,
	0x00, 0x68, 0xbc, 0x43, 0xcd, 0x3f, 0xc7, 0x42, 0xdb, 0x37, 0x31, 0xe6, 0x80, 0x06, 0x3e, 0x9c,
	0x93, 0x06, 0x16, 0xb9, 0x56, 0x95, 0xf2, 0xca, 0x94, 0xcc, 0x3a, 0xf4, 0x48, 0x47, 0x0a, 0x19,
	0x5a, 0x69, 0x6a, 0xe9, 0x4a, 0x64, 0x0d, 0x87, 0x6d, 0x72, 0xd8, 0x8a, 0x33, 0x4c, 0x01, 0x9d,
	0xec, 0x62, 0x9d, 0x6c, 0x30, 0xc3, 0xe3, 0x68, 0xd0, 0xa9, 0x55, 0x15, 0x9a, 0x58, 0x1d, 0xfd,
	0xe9, 0xd7, 0x36, 0x9f, 0x7e, 0xf7, 0xc9, 0xd1, 0x7d, 0x2b, 0xfc, 0x12, 0x7d, 0xe9, 0x0d, 0xd9,
	0xaf, 0xa5, 0x96, 0xe0, 0xd1, 0xa5, 0xbd, 0x93, 0xde, 0xc5, 0xc1, 0xe4, 0x8c, 0x6d, 0xc8, 0x07,
	0x9d, 0x15, 0x93, 0x4d, 0x63, 0x6f, 0xb6, 0x4e, 0xd1, 0x9c, 0x24, 0x80, 0xa6, 0x50, 0x65, 0xda,
	0x0f, 0xf9, 0x27, 0xb6, 0xc3, 0xf2, 0x6c, 0x8b, 0xd3, 0xf2, 0x5b, 0xa1, 0xca, 0x2c, 0x62, 0xe9,
	0x03, 0x49, 0x04, 0x78, 0x85, 0x26, 0xfd, 0x17, 0x06, 0x5c, 0x6e, 0x1d, 0xb0, 0x26, 0x3e, 0x0a,
	0xad, 0x67, 0x02, 0xde, 0x6f, 0x43, 0x26, 0x8b, 0xd9, 0xe1, 0x17, 0x49, 0x1a, 0x2e, 0x7d, 0x26,
	0x7b, 0x35, 0xa0, 0x95, 0x61, 0xdf, 0xc1, 0x64, 0xb2, 0x93, 0xef, 0x74, 0x99, 0xcc, 0x1a, 0x00,
	0x3d, 0x27, 0x87, 0x95, 0x98, 0xe7, 0x9d, 0x87, 0x10, 0xfe, 0xc1, 0xff, 0x6c, 0x50, 0x89, 0x79,
	0x67, 0xb3, 0xbb, 0xf1, 0x2b, 0x6f, 0x27, 0x28, 0xe4, 0xc2, 0x2a, 0xfe, 0xfb, 0xf5, 0x66, 0x49,
	0xb8, 0xdb, 0xd5, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xea, 0x55, 0x93, 0x40, 0x86, 0x02, 0x00,
	0x00,
}
