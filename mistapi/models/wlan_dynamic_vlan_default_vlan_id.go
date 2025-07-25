// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
    "encoding/json"
    "errors"
    "fmt"
)

// WlanDynamicVlanDefaultVlanId represents a WlanDynamicVlanDefaultVlanId struct.
// VLAN ID, VLAN range or variable to use when there’s no match from RADIUS
type WlanDynamicVlanDefaultVlanId struct {
    value    any
    isString bool
    isNumber bool
}

// String implements the fmt.Stringer interface for WlanDynamicVlanDefaultVlanId,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (w WlanDynamicVlanDefaultVlanId) String() string {
    return fmt.Sprintf("%v", w.value)
}

// MarshalJSON implements the json.Marshaler interface for WlanDynamicVlanDefaultVlanId.
// It customizes the JSON marshaling process for WlanDynamicVlanDefaultVlanId objects.
func (w WlanDynamicVlanDefaultVlanId) MarshalJSON() (
    []byte,
    error) {
    if w.value == nil {
        return nil, errors.New("No underlying type is set. Please use any of the `models.WlanDynamicVlanDefaultVlanIdContainer.From*` functions to initialize the WlanDynamicVlanDefaultVlanId object.")
    }
    return json.Marshal(w.toMap())
}

// toMap converts the WlanDynamicVlanDefaultVlanId object to a map representation for JSON marshaling.
func (w *WlanDynamicVlanDefaultVlanId) toMap() any {
    switch obj := w.value.(type) {
    case *string:
        return *obj
    case *int:
        return *obj
    }
    return nil
}

// UnmarshalJSON implements the json.Unmarshaler interface for WlanDynamicVlanDefaultVlanId.
// It customizes the JSON unmarshaling process for WlanDynamicVlanDefaultVlanId objects.
func (w *WlanDynamicVlanDefaultVlanId) UnmarshalJSON(input []byte) error {
    result, err := UnmarshallOneOf(input,
        NewTypeHolder(new(string), false, &w.isString),
        NewTypeHolder(new(int), false, &w.isNumber),
    )
    
    w.value = result
    return err
}

func (w *WlanDynamicVlanDefaultVlanId) AsString() (
    *string,
    bool) {
    if !w.isString {
        return nil, false
    }
    return w.value.(*string), true
}

func (w *WlanDynamicVlanDefaultVlanId) AsNumber() (
    *int,
    bool) {
    if !w.isNumber {
        return nil, false
    }
    return w.value.(*int), true
}

// internalWlanDynamicVlanDefaultVlanId represents a wlanDynamicVlanDefaultVlanId struct.
// VLAN ID, VLAN range or variable to use when there’s no match from RADIUS
type internalWlanDynamicVlanDefaultVlanId struct {}

var WlanDynamicVlanDefaultVlanIdContainer internalWlanDynamicVlanDefaultVlanId

// The internalWlanDynamicVlanDefaultVlanId instance, wrapping the provided string value.
func (w *internalWlanDynamicVlanDefaultVlanId) FromString(val string) WlanDynamicVlanDefaultVlanId {
    return WlanDynamicVlanDefaultVlanId{value: &val}
}

// The internalWlanDynamicVlanDefaultVlanId instance, wrapping the provided int value.
func (w *internalWlanDynamicVlanDefaultVlanId) FromNumber(val int) WlanDynamicVlanDefaultVlanId {
    return WlanDynamicVlanDefaultVlanId{value: &val}
}
