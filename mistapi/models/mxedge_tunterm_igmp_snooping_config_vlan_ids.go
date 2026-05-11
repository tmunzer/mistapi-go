// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"errors"
	"fmt"
)

// MxedgeTuntermIgmpSnoopingConfigVlanIds represents a MxedgeTuntermIgmpSnoopingConfigVlanIds struct.
// List of vlans on which tunterm performs IGMP snooping
type MxedgeTuntermIgmpSnoopingConfigVlanIds struct {
	value           any
	isArrayOfNumber bool
	isString        bool
}

// String implements the fmt.Stringer interface for MxedgeTuntermIgmpSnoopingConfigVlanIds,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (m MxedgeTuntermIgmpSnoopingConfigVlanIds) String() string {
	return fmt.Sprintf("%v", m.value)
}

// MarshalJSON implements the json.Marshaler interface for MxedgeTuntermIgmpSnoopingConfigVlanIds.
// It customizes the JSON marshaling process for MxedgeTuntermIgmpSnoopingConfigVlanIds objects.
func (m MxedgeTuntermIgmpSnoopingConfigVlanIds) MarshalJSON() (
	[]byte,
	error) {
	if m.value == nil {
		return nil, errors.New("No underlying type is set. Please use any of the `models.MxedgeTuntermIgmpSnoopingConfigVlanIdsContainer.From*` functions to initialize the MxedgeTuntermIgmpSnoopingConfigVlanIds object.")
	}
	return json.Marshal(m.toMap())
}

// toMap converts the MxedgeTuntermIgmpSnoopingConfigVlanIds object to a map representation for JSON marshaling.
func (m *MxedgeTuntermIgmpSnoopingConfigVlanIds) toMap() any {
	switch obj := m.value.(type) {
	case *[]int:
		return *obj
	case *string:
		return *obj
	}
	return nil
}

// UnmarshalJSON implements the json.Unmarshaler interface for MxedgeTuntermIgmpSnoopingConfigVlanIds.
// It customizes the JSON unmarshaling process for MxedgeTuntermIgmpSnoopingConfigVlanIds objects.
func (m *MxedgeTuntermIgmpSnoopingConfigVlanIds) UnmarshalJSON(input []byte) error {
	result, err := UnmarshallAnyOf(input,
		NewTypeHolder(new([]int), false, &m.isArrayOfNumber),
		NewTypeHolder(new(string), false, &m.isString),
	)

	m.value = result
	return err
}

func (m *MxedgeTuntermIgmpSnoopingConfigVlanIds) AsArrayOfNumber() (
	*[]int,
	bool) {
	if !m.isArrayOfNumber {
		return nil, false
	}
	return m.value.(*[]int), true
}

func (m *MxedgeTuntermIgmpSnoopingConfigVlanIds) AsString() (
	*string,
	bool) {
	if !m.isString {
		return nil, false
	}
	return m.value.(*string), true
}

// internalMxedgeTuntermIgmpSnoopingConfigVlanIds represents a mxedgeTuntermIgmpSnoopingConfigVlanIds struct.
// List of vlans on which tunterm performs IGMP snooping
type internalMxedgeTuntermIgmpSnoopingConfigVlanIds struct{}

var MxedgeTuntermIgmpSnoopingConfigVlanIdsContainer internalMxedgeTuntermIgmpSnoopingConfigVlanIds

// The internalMxedgeTuntermIgmpSnoopingConfigVlanIds instance, wrapping the provided []int value.
func (m *internalMxedgeTuntermIgmpSnoopingConfigVlanIds) FromArrayOfNumber(val []int) MxedgeTuntermIgmpSnoopingConfigVlanIds {
	return MxedgeTuntermIgmpSnoopingConfigVlanIds{value: &val}
}

// The internalMxedgeTuntermIgmpSnoopingConfigVlanIds instance, wrapping the provided string value.
func (m *internalMxedgeTuntermIgmpSnoopingConfigVlanIds) FromString(val string) MxedgeTuntermIgmpSnoopingConfigVlanIds {
	return MxedgeTuntermIgmpSnoopingConfigVlanIds{value: &val}
}
