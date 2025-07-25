// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
    "encoding/json"
    "errors"
    "fmt"
)

// VlanIdWithVariable represents a VlanIdWithVariable struct.
type VlanIdWithVariable struct {
    value    any
    isString bool
    isNumber bool
}

// String implements the fmt.Stringer interface for VlanIdWithVariable,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (v VlanIdWithVariable) String() string {
    return fmt.Sprintf("%v", v.value)
}

// MarshalJSON implements the json.Marshaler interface for VlanIdWithVariable.
// It customizes the JSON marshaling process for VlanIdWithVariable objects.
func (v VlanIdWithVariable) MarshalJSON() (
    []byte,
    error) {
    if v.value == nil {
        return nil, errors.New("No underlying type is set. Please use any of the `models.VlanIdWithVariableContainer.From*` functions to initialize the VlanIdWithVariable object.")
    }
    return json.Marshal(v.toMap())
}

// toMap converts the VlanIdWithVariable object to a map representation for JSON marshaling.
func (v *VlanIdWithVariable) toMap() any {
    switch obj := v.value.(type) {
    case *string:
        return *obj
    case *int:
        return *obj
    }
    return nil
}

// UnmarshalJSON implements the json.Unmarshaler interface for VlanIdWithVariable.
// It customizes the JSON unmarshaling process for VlanIdWithVariable objects.
func (v *VlanIdWithVariable) UnmarshalJSON(input []byte) error {
    result, err := UnmarshallOneOf(input,
        NewTypeHolder(new(string), false, &v.isString),
        NewTypeHolder(new(int), false, &v.isNumber),
    )
    
    v.value = result
    return err
}

func (v *VlanIdWithVariable) AsString() (
    *string,
    bool) {
    if !v.isString {
        return nil, false
    }
    return v.value.(*string), true
}

func (v *VlanIdWithVariable) AsNumber() (
    *int,
    bool) {
    if !v.isNumber {
        return nil, false
    }
    return v.value.(*int), true
}

// internalVlanIdWithVariable represents a vlanIdWithVariable struct.
type internalVlanIdWithVariable struct {}

var VlanIdWithVariableContainer internalVlanIdWithVariable

// The internalVlanIdWithVariable instance, wrapping the provided string value.
func (v *internalVlanIdWithVariable) FromString(val string) VlanIdWithVariable {
    return VlanIdWithVariable{value: &val}
}

// The internalVlanIdWithVariable instance, wrapping the provided int value.
func (v *internalVlanIdWithVariable) FromNumber(val int) VlanIdWithVariable {
    return VlanIdWithVariable{value: &val}
}
