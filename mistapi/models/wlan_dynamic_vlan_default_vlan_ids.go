package models

import (
    "encoding/json"
    "errors"
    "strings"
)

// WlanDynamicVlanDefaultVlanIds represents a WlanDynamicVlanDefaultVlanIds struct.
// This is Array of a container for one-of cases.
type WlanDynamicVlanDefaultVlanIds struct {
    value    any
    isString bool
    isNumber bool
}

// String converts the WlanDynamicVlanDefaultVlanIds object to a string representation.
func (w WlanDynamicVlanDefaultVlanIds) String() string {
    if bytes, err := json.Marshal(w.value); err == nil {
         return strings.Trim(string(bytes), "\"")
    }
    return ""
}

// MarshalJSON implements the json.Marshaler interface for WlanDynamicVlanDefaultVlanIds.
// It customizes the JSON marshaling process for WlanDynamicVlanDefaultVlanIds objects.
func (w WlanDynamicVlanDefaultVlanIds) MarshalJSON() (
    []byte,
    error) {
    if w.value == nil {
        return nil, errors.New("No underlying type is set. Please use any of the `models.WlanDynamicVlanDefaultVlanIdsContainer.From*` functions to initialize the WlanDynamicVlanDefaultVlanIds object.")
    }
    return json.Marshal(w.toMap())
}

// toMap converts the WlanDynamicVlanDefaultVlanIds object to a map representation for JSON marshaling.
func (w *WlanDynamicVlanDefaultVlanIds) toMap() any {
    switch obj := w.value.(type) {
    case *string:
        return *obj
    case *int:
        return *obj
    }
    return nil
}

// UnmarshalJSON implements the json.Unmarshaler interface for WlanDynamicVlanDefaultVlanIds.
// It customizes the JSON unmarshaling process for WlanDynamicVlanDefaultVlanIds objects.
func (w *WlanDynamicVlanDefaultVlanIds) UnmarshalJSON(input []byte) error {
    result, err := UnmarshallOneOf(input,
        NewTypeHolder(new(string), false, &w.isString),
        NewTypeHolder(new(int), false, &w.isNumber),
    )
    
    w.value = result
    return err
}

func (w *WlanDynamicVlanDefaultVlanIds) AsString() (
    *string,
    bool) {
    if !w.isString {
        return nil, false
    }
    return w.value.(*string), true
}

func (w *WlanDynamicVlanDefaultVlanIds) AsNumber() (
    *int,
    bool) {
    if !w.isNumber {
        return nil, false
    }
    return w.value.(*int), true
}

// internalWlanDynamicVlanDefaultVlanIds represents a wlanDynamicVlanDefaultVlanIds struct.
// This is Array of a container for one-of cases.
type internalWlanDynamicVlanDefaultVlanIds struct {}

var WlanDynamicVlanDefaultVlanIdsContainer internalWlanDynamicVlanDefaultVlanIds

// The internalWlanDynamicVlanDefaultVlanIds instance, wrapping the provided string value.
func (w *internalWlanDynamicVlanDefaultVlanIds) FromString(val string) WlanDynamicVlanDefaultVlanIds {
    return WlanDynamicVlanDefaultVlanIds{value: &val}
}

// The internalWlanDynamicVlanDefaultVlanIds instance, wrapping the provided int value.
func (w *internalWlanDynamicVlanDefaultVlanIds) FromNumber(val int) WlanDynamicVlanDefaultVlanIds {
    return WlanDynamicVlanDefaultVlanIds{value: &val}
}
