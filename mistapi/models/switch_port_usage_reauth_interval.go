// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
    "encoding/json"
    "errors"
    "fmt"
)

// SwitchPortUsageReauthInterval represents a SwitchPortUsageReauthInterval struct.
// Only if `mode`!=`dynamic` and `port_auth`=`dot1x` reauthentication interval range (min: 10, max: 65535, default: 3600)
type SwitchPortUsageReauthInterval struct {
    value    any
    isNumber bool
    isString bool
}

// String implements the fmt.Stringer interface for SwitchPortUsageReauthInterval,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s SwitchPortUsageReauthInterval) String() string {
    return fmt.Sprintf("%v", s.value)
}

// MarshalJSON implements the json.Marshaler interface for SwitchPortUsageReauthInterval.
// It customizes the JSON marshaling process for SwitchPortUsageReauthInterval objects.
func (s SwitchPortUsageReauthInterval) MarshalJSON() (
    []byte,
    error) {
    if s.value == nil {
        return nil, errors.New("No underlying type is set. Please use any of the `models.SwitchPortUsageReauthIntervalContainer.From*` functions to initialize the SwitchPortUsageReauthInterval object.")
    }
    return json.Marshal(s.toMap())
}

// toMap converts the SwitchPortUsageReauthInterval object to a map representation for JSON marshaling.
func (s *SwitchPortUsageReauthInterval) toMap() any {
    switch obj := s.value.(type) {
    case *int:
        return *obj
    case *string:
        return *obj
    }
    return nil
}

// UnmarshalJSON implements the json.Unmarshaler interface for SwitchPortUsageReauthInterval.
// It customizes the JSON unmarshaling process for SwitchPortUsageReauthInterval objects.
func (s *SwitchPortUsageReauthInterval) UnmarshalJSON(input []byte) error {
    result, err := UnmarshallAnyOf(input,
        NewTypeHolder(new(int), false, &s.isNumber),
        NewTypeHolder(new(string), false, &s.isString),
    )
    
    s.value = result
    return err
}

func (s *SwitchPortUsageReauthInterval) AsNumber() (
    *int,
    bool) {
    if !s.isNumber {
        return nil, false
    }
    return s.value.(*int), true
}

func (s *SwitchPortUsageReauthInterval) AsString() (
    *string,
    bool) {
    if !s.isString {
        return nil, false
    }
    return s.value.(*string), true
}

// internalSwitchPortUsageReauthInterval represents a switchPortUsageReauthInterval struct.
// Only if `mode`!=`dynamic` and `port_auth`=`dot1x` reauthentication interval range (min: 10, max: 65535, default: 3600)
type internalSwitchPortUsageReauthInterval struct {}

var SwitchPortUsageReauthIntervalContainer internalSwitchPortUsageReauthInterval

// The internalSwitchPortUsageReauthInterval instance, wrapping the provided int value.
func (s *internalSwitchPortUsageReauthInterval) FromNumber(val int) SwitchPortUsageReauthInterval {
    return SwitchPortUsageReauthInterval{value: &val}
}

// The internalSwitchPortUsageReauthInterval instance, wrapping the provided string value.
func (s *internalSwitchPortUsageReauthInterval) FromString(val string) SwitchPortUsageReauthInterval {
    return SwitchPortUsageReauthInterval{value: &val}
}
