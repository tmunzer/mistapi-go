// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
    "encoding/json"
    "errors"
    "fmt"
)

// ApSwitchSettingPortVlanId represents a ApSwitchSettingPortVlanId struct.
// Native VLAN id, optional
type ApSwitchSettingPortVlanId struct {
    value    any
    isNumber bool
    isString bool
}

// String implements the fmt.Stringer interface for ApSwitchSettingPortVlanId,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (a ApSwitchSettingPortVlanId) String() string {
    return fmt.Sprintf("%v", a.value)
}

// MarshalJSON implements the json.Marshaler interface for ApSwitchSettingPortVlanId.
// It customizes the JSON marshaling process for ApSwitchSettingPortVlanId objects.
func (a ApSwitchSettingPortVlanId) MarshalJSON() (
    []byte,
    error) {
    if a.value == nil {
        return nil, errors.New("No underlying type is set. Please use any of the `models.ApSwitchSettingPortVlanIdContainer.From*` functions to initialize the ApSwitchSettingPortVlanId object.")
    }
    return json.Marshal(a.toMap())
}

// toMap converts the ApSwitchSettingPortVlanId object to a map representation for JSON marshaling.
func (a *ApSwitchSettingPortVlanId) toMap() any {
    switch obj := a.value.(type) {
    case *int:
        return *obj
    case *string:
        return *obj
    }
    return nil
}

// UnmarshalJSON implements the json.Unmarshaler interface for ApSwitchSettingPortVlanId.
// It customizes the JSON unmarshaling process for ApSwitchSettingPortVlanId objects.
func (a *ApSwitchSettingPortVlanId) UnmarshalJSON(input []byte) error {
    result, err := UnmarshallAnyOf(input,
        NewTypeHolder(new(int), false, &a.isNumber),
        NewTypeHolder(new(string), false, &a.isString),
    )
    
    a.value = result
    return err
}

func (a *ApSwitchSettingPortVlanId) AsNumber() (
    *int,
    bool) {
    if !a.isNumber {
        return nil, false
    }
    return a.value.(*int), true
}

func (a *ApSwitchSettingPortVlanId) AsString() (
    *string,
    bool) {
    if !a.isString {
        return nil, false
    }
    return a.value.(*string), true
}

// internalApSwitchSettingPortVlanId represents a apSwitchSettingPortVlanId struct.
// Native VLAN id, optional
type internalApSwitchSettingPortVlanId struct {}

var ApSwitchSettingPortVlanIdContainer internalApSwitchSettingPortVlanId

// The internalApSwitchSettingPortVlanId instance, wrapping the provided int value.
func (a *internalApSwitchSettingPortVlanId) FromNumber(val int) ApSwitchSettingPortVlanId {
    return ApSwitchSettingPortVlanId{value: &val}
}

// The internalApSwitchSettingPortVlanId instance, wrapping the provided string value.
func (a *internalApSwitchSettingPortVlanId) FromString(val string) ApSwitchSettingPortVlanId {
    return ApSwitchSettingPortVlanId{value: &val}
}
