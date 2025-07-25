// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
    "encoding/json"
    "errors"
    "fmt"
)

// SwitchMgmtMxedgeProxyPort represents a SwitchMgmtMxedgeProxyPort struct.
// Mist Edge port used to proxy the switch management traffic to the Mist Cloud. Value in range 1-65535
type SwitchMgmtMxedgeProxyPort struct {
    value    any
    isNumber bool
    isString bool
}

// String implements the fmt.Stringer interface for SwitchMgmtMxedgeProxyPort,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s SwitchMgmtMxedgeProxyPort) String() string {
    return fmt.Sprintf("%v", s.value)
}

// MarshalJSON implements the json.Marshaler interface for SwitchMgmtMxedgeProxyPort.
// It customizes the JSON marshaling process for SwitchMgmtMxedgeProxyPort objects.
func (s SwitchMgmtMxedgeProxyPort) MarshalJSON() (
    []byte,
    error) {
    if s.value == nil {
        return nil, errors.New("No underlying type is set. Please use any of the `models.SwitchMgmtMxedgeProxyPortContainer.From*` functions to initialize the SwitchMgmtMxedgeProxyPort object.")
    }
    return json.Marshal(s.toMap())
}

// toMap converts the SwitchMgmtMxedgeProxyPort object to a map representation for JSON marshaling.
func (s *SwitchMgmtMxedgeProxyPort) toMap() any {
    switch obj := s.value.(type) {
    case *int:
        return *obj
    case *string:
        return *obj
    }
    return nil
}

// UnmarshalJSON implements the json.Unmarshaler interface for SwitchMgmtMxedgeProxyPort.
// It customizes the JSON unmarshaling process for SwitchMgmtMxedgeProxyPort objects.
func (s *SwitchMgmtMxedgeProxyPort) UnmarshalJSON(input []byte) error {
    result, err := UnmarshallAnyOf(input,
        NewTypeHolder(new(int), false, &s.isNumber),
        NewTypeHolder(new(string), false, &s.isString),
    )
    
    s.value = result
    return err
}

func (s *SwitchMgmtMxedgeProxyPort) AsNumber() (
    *int,
    bool) {
    if !s.isNumber {
        return nil, false
    }
    return s.value.(*int), true
}

func (s *SwitchMgmtMxedgeProxyPort) AsString() (
    *string,
    bool) {
    if !s.isString {
        return nil, false
    }
    return s.value.(*string), true
}

// internalSwitchMgmtMxedgeProxyPort represents a switchMgmtMxedgeProxyPort struct.
// Mist Edge port used to proxy the switch management traffic to the Mist Cloud. Value in range 1-65535
type internalSwitchMgmtMxedgeProxyPort struct {}

var SwitchMgmtMxedgeProxyPortContainer internalSwitchMgmtMxedgeProxyPort

// The internalSwitchMgmtMxedgeProxyPort instance, wrapping the provided int value.
func (s *internalSwitchMgmtMxedgeProxyPort) FromNumber(val int) SwitchMgmtMxedgeProxyPort {
    return SwitchMgmtMxedgeProxyPort{value: &val}
}

// The internalSwitchMgmtMxedgeProxyPort instance, wrapping the provided string value.
func (s *internalSwitchMgmtMxedgeProxyPort) FromString(val string) SwitchMgmtMxedgeProxyPort {
    return SwitchMgmtMxedgeProxyPort{value: &val}
}
