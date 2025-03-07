// Code generated by protoc-gen-jsonshim. DO NOT EDIT.
package v1alpha1

import (
	bytes "bytes"
	jsonpb "github.com/golang/protobuf/jsonpb"
)

// MarshalJSON is a custom marshaler for CircuitBreaker
func (this *CircuitBreaker) MarshalJSON() ([]byte, error) {
	str, err := CircuitBreakerMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for CircuitBreaker
func (this *CircuitBreaker) UnmarshalJSON(b []byte) error {
	return CircuitBreakerUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for CircuitBreaker_Config
func (this *CircuitBreaker_Config) MarshalJSON() ([]byte, error) {
	str, err := CircuitBreakerMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for CircuitBreaker_Config
func (this *CircuitBreaker_Config) UnmarshalJSON(b []byte) error {
	return CircuitBreakerUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for CircuitBreaker_Config_SlowRequest
func (this *CircuitBreaker_Config_SlowRequest) MarshalJSON() ([]byte, error) {
	str, err := CircuitBreakerMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for CircuitBreaker_Config_SlowRequest
func (this *CircuitBreaker_Config_SlowRequest) UnmarshalJSON(b []byte) error {
	return CircuitBreakerUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for CircuitBreaker_Config_ErrorRequest
func (this *CircuitBreaker_Config_ErrorRequest) MarshalJSON() ([]byte, error) {
	str, err := CircuitBreakerMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for CircuitBreaker_Config_ErrorRequest
func (this *CircuitBreaker_Config_ErrorRequest) UnmarshalJSON(b []byte) error {
	return CircuitBreakerUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for CircuitBreaker_Config_Recovery
func (this *CircuitBreaker_Config_Recovery) MarshalJSON() ([]byte, error) {
	str, err := CircuitBreakerMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for CircuitBreaker_Config_Recovery
func (this *CircuitBreaker_Config_Recovery) UnmarshalJSON(b []byte) error {
	return CircuitBreakerUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

var (
	CircuitBreakerMarshaler   = &jsonpb.Marshaler{}
	CircuitBreakerUnmarshaler = &jsonpb.Unmarshaler{AllowUnknownFields: true}
)
