// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"errors"
	"fmt"
)

// SwitchBgpConfigNeighborNeighborAs represents a SwitchBgpConfigNeighborNeighborAs struct.
// This is a container for any-of cases.
type SwitchBgpConfigNeighborNeighborAs struct {
	value    any
	isString bool
	isNumber bool
}

// String implements the fmt.Stringer interface for SwitchBgpConfigNeighborNeighborAs,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s SwitchBgpConfigNeighborNeighborAs) String() string {
	return fmt.Sprintf("%v", s.value)
}

// MarshalJSON implements the json.Marshaler interface for SwitchBgpConfigNeighborNeighborAs.
// It customizes the JSON marshaling process for SwitchBgpConfigNeighborNeighborAs objects.
func (s SwitchBgpConfigNeighborNeighborAs) MarshalJSON() (
	[]byte,
	error) {
	if s.value == nil {
		return nil, errors.New("No underlying type is set. Please use any of the `models.SwitchBgpConfigNeighborNeighborAsContainer.From*` functions to initialize the SwitchBgpConfigNeighborNeighborAs object.")
	}
	return json.Marshal(s.toMap())
}

// toMap converts the SwitchBgpConfigNeighborNeighborAs object to a map representation for JSON marshaling.
func (s *SwitchBgpConfigNeighborNeighborAs) toMap() any {
	switch obj := s.value.(type) {
	case *string:
		return *obj
	case *int:
		return *obj
	}
	return nil
}

// UnmarshalJSON implements the json.Unmarshaler interface for SwitchBgpConfigNeighborNeighborAs.
// It customizes the JSON unmarshaling process for SwitchBgpConfigNeighborNeighborAs objects.
func (s *SwitchBgpConfigNeighborNeighborAs) UnmarshalJSON(input []byte) error {
	result, err := UnmarshallAnyOf(input,
		NewTypeHolder(new(string), false, &s.isString),
		NewTypeHolder(new(int), false, &s.isNumber),
	)

	s.value = result
	return err
}

func (s *SwitchBgpConfigNeighborNeighborAs) AsString() (
	*string,
	bool) {
	if !s.isString {
		return nil, false
	}
	return s.value.(*string), true
}

func (s *SwitchBgpConfigNeighborNeighborAs) AsNumber() (
	*int,
	bool) {
	if !s.isNumber {
		return nil, false
	}
	return s.value.(*int), true
}

// internalSwitchBgpConfigNeighborNeighborAs represents a switchBgpConfigNeighborNeighborAs struct.
// This is a container for any-of cases.
type internalSwitchBgpConfigNeighborNeighborAs struct{}

var SwitchBgpConfigNeighborNeighborAsContainer internalSwitchBgpConfigNeighborNeighborAs

// The internalSwitchBgpConfigNeighborNeighborAs instance, wrapping the provided string value.
func (s *internalSwitchBgpConfigNeighborNeighborAs) FromString(val string) SwitchBgpConfigNeighborNeighborAs {
	return SwitchBgpConfigNeighborNeighborAs{value: &val}
}

// The internalSwitchBgpConfigNeighborNeighborAs instance, wrapping the provided int value.
func (s *internalSwitchBgpConfigNeighborNeighborAs) FromNumber(val int) SwitchBgpConfigNeighborNeighborAs {
	return SwitchBgpConfigNeighborNeighborAs{value: &val}
}
