// Code generated by protoc-gen-deepcopy. DO NOT EDIT.

package v1alpha1

import (
	proto "google.golang.org/protobuf/proto"
)

// DeepCopyInto supports using FallbackAction within kubernetes types, where deepcopy-gen is used.
func (in *FallbackAction) DeepCopyInto(out *FallbackAction) {
	p := proto.Clone(in).(*FallbackAction)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FallbackAction. Required by controller-gen.
func (in *FallbackAction) DeepCopy() *FallbackAction {
	if in == nil {
		return nil
	}
	out := new(FallbackAction)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new FallbackAction. Required by controller-gen.
func (in *FallbackAction) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}
