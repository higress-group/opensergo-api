// Code generated by protoc-gen-deepcopy. DO NOT EDIT.

package v1alpha1

import (
	proto "google.golang.org/protobuf/proto"
)

// DeepCopyInto supports using CircuitBreaker within kubernetes types, where deepcopy-gen is used.
func (in *CircuitBreaker) DeepCopyInto(out *CircuitBreaker) {
	p := proto.Clone(in).(*CircuitBreaker)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CircuitBreaker. Required by controller-gen.
func (in *CircuitBreaker) DeepCopy() *CircuitBreaker {
	if in == nil {
		return nil
	}
	out := new(CircuitBreaker)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new CircuitBreaker. Required by controller-gen.
func (in *CircuitBreaker) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}
