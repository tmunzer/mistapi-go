// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"errors"
	"fmt"
)

// AdditionalVlanIds represents a AdditionalVlanIds struct.
// List or Comma separated list of additional VLAN IDs (on the LAN side or from other WLANs) should we be forwarding bonjour queries/responses
type AdditionalVlanIds struct {
	value                        any
	isString                     bool
	isArrayOfVlanIdWithVariable7 bool
}

// String implements the fmt.Stringer interface for AdditionalVlanIds,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (a AdditionalVlanIds) String() string {
	return fmt.Sprintf("%v", a.value)
}

// MarshalJSON implements the json.Marshaler interface for AdditionalVlanIds.
// It customizes the JSON marshaling process for AdditionalVlanIds objects.
func (a AdditionalVlanIds) MarshalJSON() (
	[]byte,
	error) {
	if a.value == nil {
		return nil, errors.New("No underlying type is set. Please use any of the `models.AdditionalVlanIdsContainer.From*` functions to initialize the AdditionalVlanIds object.")
	}
	return json.Marshal(a.toMap())
}

// toMap converts the AdditionalVlanIds object to a map representation for JSON marshaling.
func (a *AdditionalVlanIds) toMap() any {
	switch obj := a.value.(type) {
	case *string:
		return *obj
	case *[]VlanIdWithVariable:
		return *obj
	}
	return nil
}

// UnmarshalJSON implements the json.Unmarshaler interface for AdditionalVlanIds.
// It customizes the JSON unmarshaling process for AdditionalVlanIds objects.
func (a *AdditionalVlanIds) UnmarshalJSON(input []byte) error {
	result, err := UnmarshallAnyOf(input,
		NewTypeHolder(new(string), false, &a.isString),
		NewTypeHolder(&[]VlanIdWithVariable{}, false, &a.isArrayOfVlanIdWithVariable7),
	)

	a.value = result
	return err
}

func (a *AdditionalVlanIds) AsString() (
	*string,
	bool) {
	if !a.isString {
		return nil, false
	}
	return a.value.(*string), true
}

func (a *AdditionalVlanIds) AsArrayOfVlanIdWithVariable7() (
	*[]VlanIdWithVariable,
	bool) {
	if !a.isArrayOfVlanIdWithVariable7 {
		return nil, false
	}
	return a.value.(*[]VlanIdWithVariable), true
}

// internalAdditionalVlanIds represents a additionalVlanIds struct.
// List or Comma separated list of additional VLAN IDs (on the LAN side or from other WLANs) should we be forwarding bonjour queries/responses
type internalAdditionalVlanIds struct{}

var AdditionalVlanIdsContainer internalAdditionalVlanIds

// The internalAdditionalVlanIds instance, wrapping the provided string value.
func (a *internalAdditionalVlanIds) FromString(val string) AdditionalVlanIds {
	return AdditionalVlanIds{value: &val}
}

// The internalAdditionalVlanIds instance, wrapping the provided []VlanIdWithVariable value.
func (a *internalAdditionalVlanIds) FromArrayOfVlanIdWithVariable7(val []VlanIdWithVariable) AdditionalVlanIds {
	return AdditionalVlanIds{value: &val}
}
