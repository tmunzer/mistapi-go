package models

import (
    "encoding/json"
    "errors"
    "strings"
)

// WlanDynamicVlanLocalVlanIds represents a WlanDynamicVlanLocalVlanIds struct.
// This is Array of a container for one-of cases.
type WlanDynamicVlanLocalVlanIds struct {
    value    any
    isString bool
    isNumber bool
}

// String converts the WlanDynamicVlanLocalVlanIds object to a string representation.
func (w WlanDynamicVlanLocalVlanIds) String() string {
    if bytes, err := json.Marshal(w.value); err == nil {
         return strings.Trim(string(bytes), "\"")
    }
    return ""
}

// MarshalJSON implements the json.Marshaler interface for WlanDynamicVlanLocalVlanIds.
// It customizes the JSON marshaling process for WlanDynamicVlanLocalVlanIds objects.
func (w WlanDynamicVlanLocalVlanIds) MarshalJSON() (
    []byte,
    error) {
    if w.value == nil {
        return nil, errors.New("No underlying type is set. Please use any of the `models.WlanDynamicVlanLocalVlanIdsContainer.From*` functions to initialize the WlanDynamicVlanLocalVlanIds object.")
    }
    return json.Marshal(w.toMap())
}

// toMap converts the WlanDynamicVlanLocalVlanIds object to a map representation for JSON marshaling.
func (w *WlanDynamicVlanLocalVlanIds) toMap() any {
    switch obj := w.value.(type) {
    case *string:
        return *obj
    case *int:
        return *obj
    }
    return nil
}

// UnmarshalJSON implements the json.Unmarshaler interface for WlanDynamicVlanLocalVlanIds.
// It customizes the JSON unmarshaling process for WlanDynamicVlanLocalVlanIds objects.
func (w *WlanDynamicVlanLocalVlanIds) UnmarshalJSON(input []byte) error {
    result, err := UnmarshallOneOf(input,
        NewTypeHolder(new(string), false, &w.isString),
        NewTypeHolder(new(int), false, &w.isNumber),
    )
    
    w.value = result
    return err
}

func (w *WlanDynamicVlanLocalVlanIds) AsString() (
    *string,
    bool) {
    if !w.isString {
        return nil, false
    }
    return w.value.(*string), true
}

func (w *WlanDynamicVlanLocalVlanIds) AsNumber() (
    *int,
    bool) {
    if !w.isNumber {
        return nil, false
    }
    return w.value.(*int), true
}

// internalWlanDynamicVlanLocalVlanIds represents a wlanDynamicVlanLocalVlanIds struct.
// This is Array of a container for one-of cases.
type internalWlanDynamicVlanLocalVlanIds struct {}

var WlanDynamicVlanLocalVlanIdsContainer internalWlanDynamicVlanLocalVlanIds

// The internalWlanDynamicVlanLocalVlanIds instance, wrapping the provided string value.
func (w *internalWlanDynamicVlanLocalVlanIds) FromString(val string) WlanDynamicVlanLocalVlanIds {
    return WlanDynamicVlanLocalVlanIds{value: &val}
}

// The internalWlanDynamicVlanLocalVlanIds instance, wrapping the provided int value.
func (w *internalWlanDynamicVlanLocalVlanIds) FromNumber(val int) WlanDynamicVlanLocalVlanIds {
    return WlanDynamicVlanLocalVlanIds{value: &val}
}
