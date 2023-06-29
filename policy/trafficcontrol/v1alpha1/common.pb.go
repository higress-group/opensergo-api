// Code generated by protoc-gen-go. DO NOT EDIT.
// source: policy/trafficcontrol/v1alpha1/common.proto

// $description: xxx.

package v1alpha1

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
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

type Scope int32

const (
	// Single instance.
	Scope_LOCAL Scope = 0
	// The whole cluster.
	Scope_GLOBAL          Scope = 1
	Scope_GLOBAL_TO_LOCAL Scope = 2
)

var Scope_name = map[int32]string{
	0: "LOCAL",
	1: "GLOBAL",
	2: "GLOBAL_TO_LOCAL",
}

var Scope_value = map[string]int32{
	"LOCAL":           0,
	"GLOBAL":          1,
	"GLOBAL_TO_LOCAL": 2,
}

func (x Scope) String() string {
	return proto.EnumName(Scope_name, int32(x))
}

func (Scope) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_963368394dca5c6b, []int{0}
}

func init() {
	proto.RegisterEnum("io.opensergo.policy.trafficcontrol.v1alpha1.Scope", Scope_name, Scope_value)
}

func init() {
	proto.RegisterFile("policy/trafficcontrol/v1alpha1/common.proto", fileDescriptor_963368394dca5c6b)
}

var fileDescriptor_963368394dca5c6b = []byte{
	// 158 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0x2e, 0xc8, 0xcf, 0xc9,
	0x4c, 0xae, 0xd4, 0x2f, 0x29, 0x4a, 0x4c, 0x4b, 0xcb, 0x4c, 0x4e, 0xce, 0xcf, 0x2b, 0x29, 0xca,
	0xcf, 0xd1, 0x2f, 0x33, 0x4c, 0xcc, 0x29, 0xc8, 0x48, 0x34, 0xd4, 0x4f, 0xce, 0xcf, 0xcd, 0xcd,
	0xcf, 0xd3, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xd2, 0xce, 0xcc, 0xd7, 0xcb, 0x2f, 0x48, 0xcd,
	0x2b, 0x4e, 0x2d, 0x4a, 0xcf, 0xd7, 0x83, 0xe8, 0xd4, 0x43, 0xd5, 0xa9, 0x07, 0xd3, 0xa9, 0x65,
	0xcc, 0xc5, 0x1a, 0x9c, 0x9c, 0x5f, 0x90, 0x2a, 0xc4, 0xc9, 0xc5, 0xea, 0xe3, 0xef, 0xec, 0xe8,
	0x23, 0xc0, 0x20, 0xc4, 0xc5, 0xc5, 0xe6, 0xee, 0xe3, 0xef, 0xe4, 0xe8, 0x23, 0xc0, 0x28, 0x24,
	0xcc, 0xc5, 0x0f, 0x61, 0xc7, 0x87, 0xf8, 0xc7, 0x43, 0x14, 0x30, 0x39, 0x19, 0x46, 0xe9, 0x23,
	0x2c, 0xc8, 0xcc, 0xd7, 0x4f, 0x2c, 0xc8, 0xd4, 0xc7, 0xef, 0xc2, 0x24, 0x36, 0xb0, 0xdb, 0x8c,
	0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0xea, 0x8b, 0x6e, 0xf3, 0xca, 0x00, 0x00, 0x00,
}
