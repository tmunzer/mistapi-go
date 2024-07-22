package models

import (
    "encoding/json"
    "errors"
    "strings"
)

// WlanDynamicPskVlanIds represents a WlanDynamicPskVlanIds struct.
// This is Array of a container for one-of cases.
type WlanDynamicPskVlanIds struct {
    value    any
    isString bool
    isNumber bool
}

// String converts the WlanDynamicPskVlanIds object to a string representation.
func (w WlanDynamicPskVlanIds) String() string {
    if bytes, err := json.Marshal(w.value); err == nil {
         return strings.Trim(string(bytes), "\"")
    }
    return ""
}

// MarshalJSON implements the json.Marshaler interface for WlanDynamicPskVlanIds.
// It customizes the JSON marshaling process for WlanDynamicPskVlanIds objects.
func (w WlanDynamicPskVlanIds) MarshalJSON() (
    []byte,
    error) {
    if w.value == nil {
        return nil, errors.New("No underlying type is set. Please use any of the `models.WlanDynamicPskVlanIdsContainer.From*` functions to initialize the WlanDynamicPskVlanIds object.")
    }
    return json.Marshal(w.toMap())
}

// toMap converts the WlanDynamicPskVlanIds object to a map representation for JSON marshaling.
func (w *WlanDynamicPskVlanIds) toMap() any {
    switch obj := w.value.(type) {
    case *string:
        return *obj
    case *int:
        return *obj
    }
    return nil
}

// UnmarshalJSON implements the json.Unmarshaler interface for WlanDynamicPskVlanIds.
// It customizes the JSON unmarshaling process for WlanDynamicPskVlanIds objects.
func (w *WlanDynamicPskVlanIds) UnmarshalJSON(input []byte) error {
    result, err := UnmarshallOneOf(input,
        NewTypeHolder(new(string), false, &w.isString),
        NewTypeHolder(new(int), false, &w.isNumber),
    )
    
    w.value = result
    return err
}

func (w *WlanDynamicPskVlanIds) AsString() (
    *string,
    bool) {
    if !w.isString {
        return nil, false
    }
    return w.value.(*string), true
}

func (w *WlanDynamicPskVlanIds) AsNumber() (
    *int,
    bool) {
    if !w.isNumber {
        return nil, false
    }
    return w.value.(*int), true
}

// internalWlanDynamicPskVlanIds represents a wlanDynamicPskVlanIds struct.
// This is Array of a container for one-of cases.
type internalWlanDynamicPskVlanIds struct {}

var WlanDynamicPskVlanIdsContainer internalWlanDynamicPskVlanIds

// The internalWlanDynamicPskVlanIds instance, wrapping the provided string value.
func (w *internalWlanDynamicPskVlanIds) FromString(val string) WlanDynamicPskVlanIds {
    return WlanDynamicPskVlanIds{value: &val}
}

// The internalWlanDynamicPskVlanIds instance, wrapping the provided int value.
func (w *internalWlanDynamicPskVlanIds) FromNumber(val int) WlanDynamicPskVlanIds {
    return WlanDynamicPskVlanIds{value: &val}
}
