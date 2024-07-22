package models

import (
    "encoding/json"
    "errors"
    "strings"
)

// WlanVlanId represents a WlanVlanId struct.
// This is a container for one-of cases.
type WlanVlanId struct {
    value    any
    isString bool
    isNumber bool
}

// String converts the WlanVlanId object to a string representation.
func (w WlanVlanId) String() string {
    if bytes, err := json.Marshal(w.value); err == nil {
         return strings.Trim(string(bytes), "\"")
    }
    return ""
}

// MarshalJSON implements the json.Marshaler interface for WlanVlanId.
// It customizes the JSON marshaling process for WlanVlanId objects.
func (w WlanVlanId) MarshalJSON() (
    []byte,
    error) {
    if w.value == nil {
        return nil, errors.New("No underlying type is set. Please use any of the `models.WlanVlanIdContainer.From*` functions to initialize the WlanVlanId object.")
    }
    return json.Marshal(w.toMap())
}

// toMap converts the WlanVlanId object to a map representation for JSON marshaling.
func (w *WlanVlanId) toMap() any {
    switch obj := w.value.(type) {
    case *string:
        return *obj
    case *int:
        return *obj
    }
    return nil
}

// UnmarshalJSON implements the json.Unmarshaler interface for WlanVlanId.
// It customizes the JSON unmarshaling process for WlanVlanId objects.
func (w *WlanVlanId) UnmarshalJSON(input []byte) error {
    result, err := UnmarshallOneOf(input,
        NewTypeHolder(new(string), false, &w.isString),
        NewTypeHolder(new(int), false, &w.isNumber),
    )
    
    w.value = result
    return err
}

func (w *WlanVlanId) AsString() (
    *string,
    bool) {
    if !w.isString {
        return nil, false
    }
    return w.value.(*string), true
}

func (w *WlanVlanId) AsNumber() (
    *int,
    bool) {
    if !w.isNumber {
        return nil, false
    }
    return w.value.(*int), true
}

// internalWlanVlanId represents a wlanVlanId struct.
// This is a container for one-of cases.
type internalWlanVlanId struct {}

var WlanVlanIdContainer internalWlanVlanId

// The internalWlanVlanId instance, wrapping the provided string value.
func (w *internalWlanVlanId) FromString(val string) WlanVlanId {
    return WlanVlanId{value: &val}
}

// The internalWlanVlanId instance, wrapping the provided int value.
func (w *internalWlanVlanId) FromNumber(val int) WlanVlanId {
    return WlanVlanId{value: &val}
}
