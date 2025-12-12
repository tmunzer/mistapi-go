// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"errors"
	"fmt"
)

// SwitchBgpConfigHoldTime represents a SwitchBgpConfigHoldTime struct.
// Hold time is three times the interval at which keepalive messages are sent. It indicates to the peer the length of time that it should consider the sender valid. Must be 0 or a number in the range 3-65535.
type SwitchBgpConfigHoldTime struct {
	value                     any
	isSwitchBgpConfigHoldTime bool
	isNumber                  bool
}

// String implements the fmt.Stringer interface for SwitchBgpConfigHoldTime,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s SwitchBgpConfigHoldTime) String() string {
	return fmt.Sprintf("%v", s.value)
}

// MarshalJSON implements the json.Marshaler interface for SwitchBgpConfigHoldTime.
// It customizes the JSON marshaling process for SwitchBgpConfigHoldTime objects.
func (s SwitchBgpConfigHoldTime) MarshalJSON() (
	[]byte,
	error) {
	if s.value == nil {
		return nil, errors.New("No underlying type is set. Please use any of the `models.SwitchBgpConfigHoldTimeContainer.From*` functions to initialize the SwitchBgpConfigHoldTime object.")
	}
	return json.Marshal(s.toMap())
}

// toMap converts the SwitchBgpConfigHoldTime object to a map representation for JSON marshaling.
func (s *SwitchBgpConfigHoldTime) toMap() any {
	switch obj := s.value.(type) {
	case *SwitchBgpConfigHoldTimeEnum:
		return *obj
	case *int:
		return *obj
	}
	return nil
}

// UnmarshalJSON implements the json.Unmarshaler interface for SwitchBgpConfigHoldTime.
// It customizes the JSON unmarshaling process for SwitchBgpConfigHoldTime objects.
func (s *SwitchBgpConfigHoldTime) UnmarshalJSON(input []byte) error {
	result, err := UnmarshallOneOf(input,
		NewTypeHolder(new(SwitchBgpConfigHoldTimeEnum), false, &s.isSwitchBgpConfigHoldTime),
		NewTypeHolder(new(int), false, &s.isNumber),
	)

	s.value = result
	return err
}

func (s *SwitchBgpConfigHoldTime) AsSwitchBgpConfigHoldTime() (
	*SwitchBgpConfigHoldTimeEnum,
	bool) {
	if !s.isSwitchBgpConfigHoldTime {
		return nil, false
	}
	return s.value.(*SwitchBgpConfigHoldTimeEnum), true
}

func (s *SwitchBgpConfigHoldTime) AsNumber() (
	*int,
	bool) {
	if !s.isNumber {
		return nil, false
	}
	return s.value.(*int), true
}

// internalSwitchBgpConfigHoldTime represents a switchBgpConfigHoldTime struct.
// Hold time is three times the interval at which keepalive messages are sent. It indicates to the peer the length of time that it should consider the sender valid. Must be 0 or a number in the range 3-65535.
type internalSwitchBgpConfigHoldTime struct{}

var SwitchBgpConfigHoldTimeContainer internalSwitchBgpConfigHoldTime

// The internalSwitchBgpConfigHoldTime instance, wrapping the provided SwitchBgpConfigHoldTimeEnum value.
func (s *internalSwitchBgpConfigHoldTime) FromSwitchBgpConfigHoldTime(val SwitchBgpConfigHoldTimeEnum) SwitchBgpConfigHoldTime {
	return SwitchBgpConfigHoldTime{value: &val}
}

// The internalSwitchBgpConfigHoldTime instance, wrapping the provided int value.
func (s *internalSwitchBgpConfigHoldTime) FromNumber(val int) SwitchBgpConfigHoldTime {
	return SwitchBgpConfigHoldTime{value: &val}
}
