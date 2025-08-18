// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"errors"
	"fmt"
)

// SwitchOspfConfigReferenceBandwidth represents a SwitchOspfConfigReferenceBandwidth struct.
// Reference bandwidth. Integer(100000) or String (10g)
type SwitchOspfConfigReferenceBandwidth struct {
	value    any
	isNumber bool
	isString bool
}

// String implements the fmt.Stringer interface for SwitchOspfConfigReferenceBandwidth,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s SwitchOspfConfigReferenceBandwidth) String() string {
	return fmt.Sprintf("%v", s.value)
}

// MarshalJSON implements the json.Marshaler interface for SwitchOspfConfigReferenceBandwidth.
// It customizes the JSON marshaling process for SwitchOspfConfigReferenceBandwidth objects.
func (s SwitchOspfConfigReferenceBandwidth) MarshalJSON() (
	[]byte,
	error) {
	if s.value == nil {
		return nil, errors.New("No underlying type is set. Please use any of the `models.SwitchOspfConfigReferenceBandwidthContainer.From*` functions to initialize the SwitchOspfConfigReferenceBandwidth object.")
	}
	return json.Marshal(s.toMap())
}

// toMap converts the SwitchOspfConfigReferenceBandwidth object to a map representation for JSON marshaling.
func (s *SwitchOspfConfigReferenceBandwidth) toMap() any {
	switch obj := s.value.(type) {
	case *int:
		return *obj
	case *string:
		return *obj
	}
	return nil
}

// UnmarshalJSON implements the json.Unmarshaler interface for SwitchOspfConfigReferenceBandwidth.
// It customizes the JSON unmarshaling process for SwitchOspfConfigReferenceBandwidth objects.
func (s *SwitchOspfConfigReferenceBandwidth) UnmarshalJSON(input []byte) error {
	result, err := UnmarshallAnyOf(input,
		NewTypeHolder(new(int), false, &s.isNumber),
		NewTypeHolder(new(string), false, &s.isString),
	)

	s.value = result
	return err
}

func (s *SwitchOspfConfigReferenceBandwidth) AsNumber() (
	*int,
	bool) {
	if !s.isNumber {
		return nil, false
	}
	return s.value.(*int), true
}

func (s *SwitchOspfConfigReferenceBandwidth) AsString() (
	*string,
	bool) {
	if !s.isString {
		return nil, false
	}
	return s.value.(*string), true
}

// internalSwitchOspfConfigReferenceBandwidth represents a switchOspfConfigReferenceBandwidth struct.
// Reference bandwidth. Integer(100000) or String (10g)
type internalSwitchOspfConfigReferenceBandwidth struct{}

var SwitchOspfConfigReferenceBandwidthContainer internalSwitchOspfConfigReferenceBandwidth

// The internalSwitchOspfConfigReferenceBandwidth instance, wrapping the provided int value.
func (s *internalSwitchOspfConfigReferenceBandwidth) FromNumber(val int) SwitchOspfConfigReferenceBandwidth {
	return SwitchOspfConfigReferenceBandwidth{value: &val}
}

// The internalSwitchOspfConfigReferenceBandwidth instance, wrapping the provided string value.
func (s *internalSwitchOspfConfigReferenceBandwidth) FromString(val string) SwitchOspfConfigReferenceBandwidth {
	return SwitchOspfConfigReferenceBandwidth{value: &val}
}
