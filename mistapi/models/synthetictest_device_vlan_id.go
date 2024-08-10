package models

import (
    "encoding/json"
    "errors"
    "strings"
)

// SynthetictestDeviceVlanId represents a SynthetictestDeviceVlanId struct.
// required for AP
type SynthetictestDeviceVlanId struct {
    value    any
    isString bool
    isNumber bool
}

// String converts the SynthetictestDeviceVlanId object to a string representation.
func (s SynthetictestDeviceVlanId) String() string {
    if bytes, err := json.Marshal(s.value); err == nil {
         return strings.Trim(string(bytes), "\"")
    }
    return ""
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
// required for AP
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
