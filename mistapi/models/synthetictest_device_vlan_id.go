// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
    "encoding/json"
    "errors"
    "fmt"
)

// SynthetictestDeviceVlanId represents a SynthetictestDeviceVlanId struct.
// Required for AP
type SynthetictestDeviceVlanId struct {
    value    any
    isString bool
    isNumber bool
}

// String implements the fmt.Stringer interface for SynthetictestDeviceVlanId,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s SynthetictestDeviceVlanId) String() string {
    return fmt.Sprintf("%v", s.value)
}

// MarshalJSON implements the json.Marshaler interface for SynthetictestDeviceVlanId.
// It customizes the JSON marshaling process for SynthetictestDeviceVlanId objects.
func (s SynthetictestDeviceVlanId) MarshalJSON() (
    []byte,
    error) {
    if s.value == nil {
        return nil, errors.New("No underlying type is set. Please use any of the `models.SynthetictestDeviceVlanIdContainer.From*` functions to initialize the SynthetictestDeviceVlanId object.")
    }
    return json.Marshal(s.toMap())
}

// toMap converts the SynthetictestDeviceVlanId object to a map representation for JSON marshaling.
func (s *SynthetictestDeviceVlanId) toMap() any {
    switch obj := s.value.(type) {
    case *string:
        return *obj
    case *int:
        return *obj
    }
    return nil
}

// UnmarshalJSON implements the json.Unmarshaler interface for SynthetictestDeviceVlanId.
// It customizes the JSON unmarshaling process for SynthetictestDeviceVlanId objects.
func (s *SynthetictestDeviceVlanId) UnmarshalJSON(input []byte) error {
    result, err := UnmarshallOneOf(input,
        NewTypeHolder(new(string), false, &s.isString),
        NewTypeHolder(new(int), false, &s.isNumber),
    )
    
    s.value = result
    return err
}

func (s *SynthetictestDeviceVlanId) AsString() (
    *string,
    bool) {
    if !s.isString {
        return nil, false
    }
    return s.value.(*string), true
}

func (s *SynthetictestDeviceVlanId) AsNumber() (
    *int,
    bool) {
    if !s.isNumber {
        return nil, false
    }
    return s.value.(*int), true
}

// internalSynthetictestDeviceVlanId represents a synthetictestDeviceVlanId struct.
// Required for AP
type internalSynthetictestDeviceVlanId struct {}

var SynthetictestDeviceVlanIdContainer internalSynthetictestDeviceVlanId

// The internalSynthetictestDeviceVlanId instance, wrapping the provided string value.
func (s *internalSynthetictestDeviceVlanId) FromString(val string) SynthetictestDeviceVlanId {
    return SynthetictestDeviceVlanId{value: &val}
}

// The internalSynthetictestDeviceVlanId instance, wrapping the provided int value.
func (s *internalSynthetictestDeviceVlanId) FromNumber(val int) SynthetictestDeviceVlanId {
    return SynthetictestDeviceVlanId{value: &val}
}
