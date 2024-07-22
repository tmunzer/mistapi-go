package models

import (
    "encoding/json"
    "errors"
    "strings"
)

// SwitchNetworkVlanId represents a SwitchNetworkVlanId struct.
// This is a container for one-of cases.
type SwitchNetworkVlanId struct {
    value    any
    isString bool
    isNumber bool
}

// String converts the SwitchNetworkVlanId object to a string representation.
func (s SwitchNetworkVlanId) String() string {
    if bytes, err := json.Marshal(s.value); err == nil {
         return strings.Trim(string(bytes), "\"")
    }
    return ""
}

// MarshalJSON implements the json.Marshaler interface for SwitchNetworkVlanId.
// It customizes the JSON marshaling process for SwitchNetworkVlanId objects.
func (s SwitchNetworkVlanId) MarshalJSON() (
    []byte,
    error) {
    if s.value == nil {
        return nil, errors.New("No underlying type is set. Please use any of the `models.SwitchNetworkVlanIdContainer.From*` functions to initialize the SwitchNetworkVlanId object.")
    }
    return json.Marshal(s.toMap())
}

// toMap converts the SwitchNetworkVlanId object to a map representation for JSON marshaling.
func (s *SwitchNetworkVlanId) toMap() any {
    switch obj := s.value.(type) {
    case *string:
        return *obj
    case *int:
        return *obj
    }
    return nil
}

// UnmarshalJSON implements the json.Unmarshaler interface for SwitchNetworkVlanId.
// It customizes the JSON unmarshaling process for SwitchNetworkVlanId objects.
func (s *SwitchNetworkVlanId) UnmarshalJSON(input []byte) error {
    result, err := UnmarshallOneOf(input,
        NewTypeHolder(new(string), false, &s.isString),
        NewTypeHolder(new(int), false, &s.isNumber),
    )
    
    s.value = result
    return err
}

func (s *SwitchNetworkVlanId) AsString() (
    *string,
    bool) {
    if !s.isString {
        return nil, false
    }
    return s.value.(*string), true
}

func (s *SwitchNetworkVlanId) AsNumber() (
    *int,
    bool) {
    if !s.isNumber {
        return nil, false
    }
    return s.value.(*int), true
}

// internalSwitchNetworkVlanId represents a switchNetworkVlanId struct.
// This is a container for one-of cases.
type internalSwitchNetworkVlanId struct {}

var SwitchNetworkVlanIdContainer internalSwitchNetworkVlanId

// The internalSwitchNetworkVlanId instance, wrapping the provided string value.
func (s *internalSwitchNetworkVlanId) FromString(val string) SwitchNetworkVlanId {
    return SwitchNetworkVlanId{value: &val}
}

// The internalSwitchNetworkVlanId instance, wrapping the provided int value.
func (s *internalSwitchNetworkVlanId) FromNumber(val int) SwitchNetworkVlanId {
    return SwitchNetworkVlanId{value: &val}
}
