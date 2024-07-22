package models

import (
    "encoding/json"
    "errors"
    "strings"
)

// WlanVlanIds represents a WlanVlanIds struct.
// This is Array of a container for one-of cases.
type WlanVlanIds struct {
    value    any
    isString bool
    isNumber bool
}

// String converts the WlanVlanIds object to a string representation.
func (w WlanVlanIds) String() string {
    if bytes, err := json.Marshal(w.value); err == nil {
         return strings.Trim(string(bytes), "\"")
    }
    return ""
}

// MarshalJSON implements the json.Marshaler interface for WlanVlanIds.
// It customizes the JSON marshaling process for WlanVlanIds objects.
func (w WlanVlanIds) MarshalJSON() (
    []byte,
    error) {
    if w.value == nil {
        return nil, errors.New("No underlying type is set. Please use any of the `models.WlanVlanIdsContainer.From*` functions to initialize the WlanVlanIds object.")
    }
    return json.Marshal(w.toMap())
}

// toMap converts the WlanVlanIds object to a map representation for JSON marshaling.
func (w *WlanVlanIds) toMap() any {
    switch obj := w.value.(type) {
    case *string:
        return *obj
    case *int:
        return *obj
    }
    return nil
}

// UnmarshalJSON implements the json.Unmarshaler interface for WlanVlanIds.
// It customizes the JSON unmarshaling process for WlanVlanIds objects.
func (w *WlanVlanIds) UnmarshalJSON(input []byte) error {
    result, err := UnmarshallOneOf(input,
        NewTypeHolder(new(string), false, &w.isString),
        NewTypeHolder(new(int), false, &w.isNumber),
    )
    
    w.value = result
    return err
}

func (w *WlanVlanIds) AsString() (
    *string,
    bool) {
    if !w.isString {
        return nil, false
    }
    return w.value.(*string), true
}

func (w *WlanVlanIds) AsNumber() (
    *int,
    bool) {
    if !w.isNumber {
        return nil, false
    }
    return w.value.(*int), true
}

// internalWlanVlanIds represents a wlanVlanIds struct.
// This is Array of a container for one-of cases.
type internalWlanVlanIds struct {}

var WlanVlanIdsContainer internalWlanVlanIds

// The internalWlanVlanIds instance, wrapping the provided string value.
func (w *internalWlanVlanIds) FromString(val string) WlanVlanIds {
    return WlanVlanIds{value: &val}
}

// The internalWlanVlanIds instance, wrapping the provided int value.
func (w *internalWlanVlanIds) FromNumber(val int) WlanVlanIds {
    return WlanVlanIds{value: &val}
}
