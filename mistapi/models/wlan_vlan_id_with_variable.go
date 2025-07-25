// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
    "encoding/json"
    "errors"
    "fmt"
)

// WlanVlanIdWithVariable represents a WlanVlanIdWithVariable struct.
type WlanVlanIdWithVariable struct {
    value    any
    isString bool
    isNumber bool
}

// String implements the fmt.Stringer interface for WlanVlanIdWithVariable,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (w WlanVlanIdWithVariable) String() string {
    return fmt.Sprintf("%v", w.value)
}

// MarshalJSON implements the json.Marshaler interface for WlanVlanIdWithVariable.
// It customizes the JSON marshaling process for WlanVlanIdWithVariable objects.
func (w WlanVlanIdWithVariable) MarshalJSON() (
    []byte,
    error) {
    if w.value == nil {
        return nil, errors.New("No underlying type is set. Please use any of the `models.WlanVlanIdWithVariableContainer.From*` functions to initialize the WlanVlanIdWithVariable object.")
    }
    return json.Marshal(w.toMap())
}

// toMap converts the WlanVlanIdWithVariable object to a map representation for JSON marshaling.
func (w *WlanVlanIdWithVariable) toMap() any {
    switch obj := w.value.(type) {
    case *string:
        return *obj
    case *int:
        return *obj
    }
    return nil
}

// UnmarshalJSON implements the json.Unmarshaler interface for WlanVlanIdWithVariable.
// It customizes the JSON unmarshaling process for WlanVlanIdWithVariable objects.
func (w *WlanVlanIdWithVariable) UnmarshalJSON(input []byte) error {
    result, err := UnmarshallOneOf(input,
        NewTypeHolder(new(string), false, &w.isString),
        NewTypeHolder(new(int), false, &w.isNumber),
    )
    
    w.value = result
    return err
}

func (w *WlanVlanIdWithVariable) AsString() (
    *string,
    bool) {
    if !w.isString {
        return nil, false
    }
    return w.value.(*string), true
}

func (w *WlanVlanIdWithVariable) AsNumber() (
    *int,
    bool) {
    if !w.isNumber {
        return nil, false
    }
    return w.value.(*int), true
}

// internalWlanVlanIdWithVariable represents a wlanVlanIdWithVariable struct.
type internalWlanVlanIdWithVariable struct {}

var WlanVlanIdWithVariableContainer internalWlanVlanIdWithVariable

// The internalWlanVlanIdWithVariable instance, wrapping the provided string value.
func (w *internalWlanVlanIdWithVariable) FromString(val string) WlanVlanIdWithVariable {
    return WlanVlanIdWithVariable{value: &val}
}

// The internalWlanVlanIdWithVariable instance, wrapping the provided int value.
func (w *internalWlanVlanIdWithVariable) FromNumber(val int) WlanVlanIdWithVariable {
    return WlanVlanIdWithVariable{value: &val}
}
