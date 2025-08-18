// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"errors"
	"fmt"
)

// SwitchPortUsageMacLimitOverwrite represents a SwitchPortUsageMacLimitOverwrite struct.
// Max number of mac addresses, default is 0 for unlimited, otherwise range is 1 to 16383 (upper bound constrained by platform)
type SwitchPortUsageMacLimitOverwrite struct {
	value    any
	isNumber bool
	isString bool
}

// String implements the fmt.Stringer interface for SwitchPortUsageMacLimitOverwrite,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s SwitchPortUsageMacLimitOverwrite) String() string {
	return fmt.Sprintf("%v", s.value)
}

// MarshalJSON implements the json.Marshaler interface for SwitchPortUsageMacLimitOverwrite.
// It customizes the JSON marshaling process for SwitchPortUsageMacLimitOverwrite objects.
func (s SwitchPortUsageMacLimitOverwrite) MarshalJSON() (
	[]byte,
	error) {
	if s.value == nil {
		return nil, errors.New("No underlying type is set. Please use any of the `models.SwitchPortUsageMacLimitOverwriteContainer.From*` functions to initialize the SwitchPortUsageMacLimitOverwrite object.")
	}
	return json.Marshal(s.toMap())
}

// toMap converts the SwitchPortUsageMacLimitOverwrite object to a map representation for JSON marshaling.
func (s *SwitchPortUsageMacLimitOverwrite) toMap() any {
	switch obj := s.value.(type) {
	case *int:
		return *obj
	case *string:
		return *obj
	}
	return nil
}

// UnmarshalJSON implements the json.Unmarshaler interface for SwitchPortUsageMacLimitOverwrite.
// It customizes the JSON unmarshaling process for SwitchPortUsageMacLimitOverwrite objects.
func (s *SwitchPortUsageMacLimitOverwrite) UnmarshalJSON(input []byte) error {
	result, err := UnmarshallAnyOf(input,
		NewTypeHolder(new(int), false, &s.isNumber),
		NewTypeHolder(new(string), false, &s.isString),
	)

	s.value = result
	return err
}

func (s *SwitchPortUsageMacLimitOverwrite) AsNumber() (
	*int,
	bool) {
	if !s.isNumber {
		return nil, false
	}
	return s.value.(*int), true
}

func (s *SwitchPortUsageMacLimitOverwrite) AsString() (
	*string,
	bool) {
	if !s.isString {
		return nil, false
	}
	return s.value.(*string), true
}

// internalSwitchPortUsageMacLimitOverwrite represents a switchPortUsageMacLimitOverwrite struct.
// Max number of mac addresses, default is 0 for unlimited, otherwise range is 1 to 16383 (upper bound constrained by platform)
type internalSwitchPortUsageMacLimitOverwrite struct{}

var SwitchPortUsageMacLimitOverwriteContainer internalSwitchPortUsageMacLimitOverwrite

// The internalSwitchPortUsageMacLimitOverwrite instance, wrapping the provided int value.
func (s *internalSwitchPortUsageMacLimitOverwrite) FromNumber(val int) SwitchPortUsageMacLimitOverwrite {
	return SwitchPortUsageMacLimitOverwrite{value: &val}
}

// The internalSwitchPortUsageMacLimitOverwrite instance, wrapping the provided string value.
func (s *internalSwitchPortUsageMacLimitOverwrite) FromString(val string) SwitchPortUsageMacLimitOverwrite {
	return SwitchPortUsageMacLimitOverwrite{value: &val}
}
