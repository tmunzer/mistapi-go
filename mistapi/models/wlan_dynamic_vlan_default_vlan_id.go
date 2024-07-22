package models

import (
    "encoding/json"
    "errors"
    "strings"
)

// WlanDynamicVlanDefaultVlanId represents a WlanDynamicVlanDefaultVlanId struct.
// This is a container for one-of cases.
type WlanDynamicVlanDefaultVlanId struct {
    value    any
    isString bool
    isNumber bool
}

// String converts the WlanDynamicVlanDefaultVlanId object to a string representation.
func (w WlanDynamicVlanDefaultVlanId) String() string {
    if bytes, err := json.Marshal(w.value); err == nil {
         return strings.Trim(string(bytes), "\"")
    }
    return ""
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
// This is a container for one-of cases.
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
