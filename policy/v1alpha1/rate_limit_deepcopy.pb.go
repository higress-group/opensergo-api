// Code generated by protoc-gen-deepcopy. DO NOT EDIT.

package v1alpha1

import (
	proto "google.golang.org/protobuf/proto"
)

// DeepCopyInto supports using RateLimit within kubernetes types, where deepcopy-gen is used.
func (in *RateLimit) DeepCopyInto(out *RateLimit) {
	p := proto.Clone(in).(*RateLimit)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RateLimit. Required by controller-gen.
func (in *RateLimit) DeepCopy() *RateLimit {
	if in == nil {
		return nil
	}
	out := new(RateLimit)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new RateLimit. Required by controller-gen.
func (in *RateLimit) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}
