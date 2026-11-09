// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"errors"
	"fmt"
)

// SwitchPortUsageServerFailRetryInterval represents a SwitchPortUsageServerFailRetryInterval struct.
// This is a container for any-of cases.
type SwitchPortUsageServerFailRetryInterval struct {
	value    any
	isNumber bool
	isString bool
}

// String implements the fmt.Stringer interface for SwitchPortUsageServerFailRetryInterval,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s SwitchPortUsageServerFailRetryInterval) String() string {
	return fmt.Sprintf("%v", s.value)
}

// MarshalJSON implements the json.Marshaler interface for SwitchPortUsageServerFailRetryInterval.
// It customizes the JSON marshaling process for SwitchPortUsageServerFailRetryInterval objects.
func (s SwitchPortUsageServerFailRetryInterval) MarshalJSON() (
	[]byte,
	error) {
	if s.value == nil {
		return nil, errors.New("No underlying type is set. Please use any of the `models.SwitchPortUsageServerFailRetryIntervalContainer.From*` functions to initialize the SwitchPortUsageServerFailRetryInterval object.")
	}
	return json.Marshal(s.toMap())
}

// toMap converts the SwitchPortUsageServerFailRetryInterval object to a map representation for JSON marshaling.
func (s *SwitchPortUsageServerFailRetryInterval) toMap() any {
	switch obj := s.value.(type) {
	case *int:
		return *obj
	case *string:
		return *obj
	}
	return nil
}

// UnmarshalJSON implements the json.Unmarshaler interface for SwitchPortUsageServerFailRetryInterval.
// It customizes the JSON unmarshaling process for SwitchPortUsageServerFailRetryInterval objects.
func (s *SwitchPortUsageServerFailRetryInterval) UnmarshalJSON(input []byte) error {
	result, err := UnmarshallAnyOf(input,
		NewTypeHolder(new(int), false, &s.isNumber),
		NewTypeHolder(new(string), false, &s.isString),
	)

	s.value = result
	return err
}

func (s *SwitchPortUsageServerFailRetryInterval) AsNumber() (
	*int,
	bool) {
	if !s.isNumber {
		return nil, false
	}
	return s.value.(*int), true
}

func (s *SwitchPortUsageServerFailRetryInterval) AsString() (
	*string,
	bool) {
	if !s.isString {
		return nil, false
	}
	return s.value.(*string), true
}

// internalSwitchPortUsageServerFailRetryInterval represents a switchPortUsageServerFailRetryInterval struct.
// This is a container for any-of cases.
type internalSwitchPortUsageServerFailRetryInterval struct{}

var SwitchPortUsageServerFailRetryIntervalContainer internalSwitchPortUsageServerFailRetryInterval

// The internalSwitchPortUsageServerFailRetryInterval instance, wrapping the provided int value.
func (s *internalSwitchPortUsageServerFailRetryInterval) FromNumber(val int) SwitchPortUsageServerFailRetryInterval {
	return SwitchPortUsageServerFailRetryInterval{value: &val}
}

// The internalSwitchPortUsageServerFailRetryInterval instance, wrapping the provided string value.
func (s *internalSwitchPortUsageServerFailRetryInterval) FromString(val string) SwitchPortUsageServerFailRetryInterval {
	return SwitchPortUsageServerFailRetryInterval{value: &val}
}
