// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"errors"
	"fmt"
)

// AdditionalVlanIds2 represents a AdditionalVlanIds2 struct.
// List or Comma separated list of additional VLAN IDs (on the LAN side or from other WLANs) should we be forwarding bonjour queries/responses
type AdditionalVlanIds2 struct {
	value                        any
	isString                     bool
	isArrayOfVlanIdWithVariable8 bool
}

// String implements the fmt.Stringer interface for AdditionalVlanIds2,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (a AdditionalVlanIds2) String() string {
	return fmt.Sprintf("%v", a.value)
}

// MarshalJSON implements the json.Marshaler interface for AdditionalVlanIds2.
// It customizes the JSON marshaling process for AdditionalVlanIds2 objects.
func (a AdditionalVlanIds2) MarshalJSON() (
	[]byte,
	error) {
	if a.value == nil {
		return nil, errors.New("No underlying type is set. Please use any of the `models.AdditionalVlanIds2Container.From*` functions to initialize the AdditionalVlanIds2 object.")
	}
	return json.Marshal(a.toMap())
}

// toMap converts the AdditionalVlanIds2 object to a map representation for JSON marshaling.
func (a *AdditionalVlanIds2) toMap() any {
	switch obj := a.value.(type) {
	case *string:
		return *obj
	case *[]VlanIdWithVariable:
		return *obj
	}
	return nil
}

// UnmarshalJSON implements the json.Unmarshaler interface for AdditionalVlanIds2.
// It customizes the JSON unmarshaling process for AdditionalVlanIds2 objects.
func (a *AdditionalVlanIds2) UnmarshalJSON(input []byte) error {
	result, err := UnmarshallAnyOf(input,
		NewTypeHolder(new(string), false, &a.isString),
		NewTypeHolder(&[]VlanIdWithVariable{}, false, &a.isArrayOfVlanIdWithVariable8),
	)

	a.value = result
	return err
}

func (a *AdditionalVlanIds2) AsString() (
	*string,
	bool) {
	if !a.isString {
		return nil, false
	}
	return a.value.(*string), true
}

func (a *AdditionalVlanIds2) AsArrayOfVlanIdWithVariable8() (
	*[]VlanIdWithVariable,
	bool) {
	if !a.isArrayOfVlanIdWithVariable8 {
		return nil, false
	}
	return a.value.(*[]VlanIdWithVariable), true
}

// internalAdditionalVlanIds2 represents a additionalVlanIds2 struct.
// List or Comma separated list of additional VLAN IDs (on the LAN side or from other WLANs) should we be forwarding bonjour queries/responses
type internalAdditionalVlanIds2 struct{}

var AdditionalVlanIds2Container internalAdditionalVlanIds2

// The internalAdditionalVlanIds2 instance, wrapping the provided string value.
func (a *internalAdditionalVlanIds2) FromString(val string) AdditionalVlanIds2 {
	return AdditionalVlanIds2{value: &val}
}

// The internalAdditionalVlanIds2 instance, wrapping the provided []VlanIdWithVariable value.
func (a *internalAdditionalVlanIds2) FromArrayOfVlanIdWithVariable8(val []VlanIdWithVariable) AdditionalVlanIds2 {
	return AdditionalVlanIds2{value: &val}
}
