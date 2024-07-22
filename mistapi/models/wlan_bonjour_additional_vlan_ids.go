package models

import (
    "encoding/json"
    "errors"
    "strings"
)

// WlanBonjourAdditionalVlanIds represents a WlanBonjourAdditionalVlanIds struct.
// This is Array of a container for one-of cases.
type WlanBonjourAdditionalVlanIds struct {
    value    any
    isString bool
    isNumber bool
}

// String converts the WlanBonjourAdditionalVlanIds object to a string representation.
func (w WlanBonjourAdditionalVlanIds) String() string {
    if bytes, err := json.Marshal(w.value); err == nil {
         return strings.Trim(string(bytes), "\"")
    }
    return ""
}

// MarshalJSON implements the json.Marshaler interface for WlanBonjourAdditionalVlanIds.
// It customizes the JSON marshaling process for WlanBonjourAdditionalVlanIds objects.
func (w WlanBonjourAdditionalVlanIds) MarshalJSON() (
    []byte,
    error) {
    if w.value == nil {
        return nil, errors.New("No underlying type is set. Please use any of the `models.WlanBonjourAdditionalVlanIdsContainer.From*` functions to initialize the WlanBonjourAdditionalVlanIds object.")
    }
    return json.Marshal(w.toMap())
}

// toMap converts the WlanBonjourAdditionalVlanIds object to a map representation for JSON marshaling.
func (w *WlanBonjourAdditionalVlanIds) toMap() any {
    switch obj := w.value.(type) {
    case *string:
        return *obj
    case *int:
        return *obj
    }
    return nil
}

// UnmarshalJSON implements the json.Unmarshaler interface for WlanBonjourAdditionalVlanIds.
// It customizes the JSON unmarshaling process for WlanBonjourAdditionalVlanIds objects.
func (w *WlanBonjourAdditionalVlanIds) UnmarshalJSON(input []byte) error {
    result, err := UnmarshallOneOf(input,
        NewTypeHolder(new(string), false, &w.isString),
        NewTypeHolder(new(int), false, &w.isNumber),
    )
    
    w.value = result
    return err
}

func (w *WlanBonjourAdditionalVlanIds) AsString() (
    *string,
    bool) {
    if !w.isString {
        return nil, false
    }
    return w.value.(*string), true
}

func (w *WlanBonjourAdditionalVlanIds) AsNumber() (
    *int,
    bool) {
    if !w.isNumber {
        return nil, false
    }
    return w.value.(*int), true
}

// internalWlanBonjourAdditionalVlanIds represents a wlanBonjourAdditionalVlanIds struct.
// This is Array of a container for one-of cases.
type internalWlanBonjourAdditionalVlanIds struct {}

var WlanBonjourAdditionalVlanIdsContainer internalWlanBonjourAdditionalVlanIds

// The internalWlanBonjourAdditionalVlanIds instance, wrapping the provided string value.
func (w *internalWlanBonjourAdditionalVlanIds) FromString(val string) WlanBonjourAdditionalVlanIds {
    return WlanBonjourAdditionalVlanIds{value: &val}
}

// The internalWlanBonjourAdditionalVlanIds instance, wrapping the provided int value.
func (w *internalWlanBonjourAdditionalVlanIds) FromNumber(val int) WlanBonjourAdditionalVlanIds {
    return WlanBonjourAdditionalVlanIds{value: &val}
}
