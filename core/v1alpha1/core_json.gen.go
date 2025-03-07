// Code generated by protoc-gen-jsonshim. DO NOT EDIT.
package v1alpha1

import (
	bytes "bytes"
	jsonpb "github.com/golang/protobuf/jsonpb"
)

// MarshalJSON is a custom marshaler for Selector
func (this *Selector) MarshalJSON() ([]byte, error) {
	str, err := CoreMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for Selector
func (this *Selector) UnmarshalJSON(b []byte) error {
	return CoreUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for RouteSelector
func (this *RouteSelector) MarshalJSON() ([]byte, error) {
	str, err := CoreMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for RouteSelector
func (this *RouteSelector) UnmarshalJSON(b []byte) error {
	return CoreUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for WorkloadSelector
func (this *WorkloadSelector) MarshalJSON() ([]byte, error) {
	str, err := CoreMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for WorkloadSelector
func (this *WorkloadSelector) UnmarshalJSON(b []byte) error {
	return CoreUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

var (
	CoreMarshaler   = &jsonpb.Marshaler{}
	CoreUnmarshaler = &jsonpb.Unmarshaler{AllowUnknownFields: true}
)
