// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"errors"
	"fmt"
)

// MxedgeTuntermIgmpSnoopingConfigEnabled represents a MxedgeTuntermIgmpSnoopingConfigEnabled struct.
// This is a container for any-of cases.
type MxedgeTuntermIgmpSnoopingConfigEnabled struct {
	value     any
	isBoolean bool
	isString  bool
}

// String implements the fmt.Stringer interface for MxedgeTuntermIgmpSnoopingConfigEnabled,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (m MxedgeTuntermIgmpSnoopingConfigEnabled) String() string {
	return fmt.Sprintf("%v", m.value)
}

// MarshalJSON implements the json.Marshaler interface for MxedgeTuntermIgmpSnoopingConfigEnabled.
// It customizes the JSON marshaling process for MxedgeTuntermIgmpSnoopingConfigEnabled objects.
func (m MxedgeTuntermIgmpSnoopingConfigEnabled) MarshalJSON() (
	[]byte,
	error) {
	if m.value == nil {
		return nil, errors.New("No underlying type is set. Please use any of the `models.MxedgeTuntermIgmpSnoopingConfigEnabledContainer.From*` functions to initialize the MxedgeTuntermIgmpSnoopingConfigEnabled object.")
	}
	return json.Marshal(m.toMap())
}

// toMap converts the MxedgeTuntermIgmpSnoopingConfigEnabled object to a map representation for JSON marshaling.
func (m *MxedgeTuntermIgmpSnoopingConfigEnabled) toMap() any {
	switch obj := m.value.(type) {
	case *bool:
		return *obj
	case *string:
		return *obj
	}
	return nil
}

// UnmarshalJSON implements the json.Unmarshaler interface for MxedgeTuntermIgmpSnoopingConfigEnabled.
// It customizes the JSON unmarshaling process for MxedgeTuntermIgmpSnoopingConfigEnabled objects.
func (m *MxedgeTuntermIgmpSnoopingConfigEnabled) UnmarshalJSON(input []byte) error {
	result, err := UnmarshallAnyOf(input,
		NewTypeHolder(new(bool), false, &m.isBoolean),
		NewTypeHolder(new(string), false, &m.isString),
	)

	m.value = result
	return err
}

func (m *MxedgeTuntermIgmpSnoopingConfigEnabled) AsBoolean() (
	*bool,
	bool) {
	if !m.isBoolean {
		return nil, false
	}
	return m.value.(*bool), true
}

func (m *MxedgeTuntermIgmpSnoopingConfigEnabled) AsString() (
	*string,
	bool) {
	if !m.isString {
		return nil, false
	}
	return m.value.(*string), true
}

// internalMxedgeTuntermIgmpSnoopingConfigEnabled represents a mxedgeTuntermIgmpSnoopingConfigEnabled struct.
// This is a container for any-of cases.
type internalMxedgeTuntermIgmpSnoopingConfigEnabled struct{}

var MxedgeTuntermIgmpSnoopingConfigEnabledContainer internalMxedgeTuntermIgmpSnoopingConfigEnabled

// The internalMxedgeTuntermIgmpSnoopingConfigEnabled instance, wrapping the provided bool value.
func (m *internalMxedgeTuntermIgmpSnoopingConfigEnabled) FromBoolean(val bool) MxedgeTuntermIgmpSnoopingConfigEnabled {
	return MxedgeTuntermIgmpSnoopingConfigEnabled{value: &val}
}

// The internalMxedgeTuntermIgmpSnoopingConfigEnabled instance, wrapping the provided string value.
func (m *internalMxedgeTuntermIgmpSnoopingConfigEnabled) FromString(val string) MxedgeTuntermIgmpSnoopingConfigEnabled {
	return MxedgeTuntermIgmpSnoopingConfigEnabled{value: &val}
}
