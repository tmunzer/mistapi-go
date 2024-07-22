package models

import (
    "encoding/json"
    "errors"
    "strings"
)

// WlanDynamicPskDefaultVlanId represents a WlanDynamicPskDefaultVlanId struct.
// This is a container for one-of cases.
type WlanDynamicPskDefaultVlanId struct {
    value    any
    isString bool
    isNumber bool
}

// String converts the WlanDynamicPskDefaultVlanId object to a string representation.
func (w WlanDynamicPskDefaultVlanId) String() string {
    if bytes, err := json.Marshal(w.value); err == nil {
         return strings.Trim(string(bytes), "\"")
    }
    return ""
}

// MarshalJSON implements the json.Marshaler interface for WlanDynamicPskDefaultVlanId.
// It customizes the JSON marshaling process for WlanDynamicPskDefaultVlanId objects.
func (w WlanDynamicPskDefaultVlanId) MarshalJSON() (
    []byte,
    error) {
    if w.value == nil {
        return nil, errors.New("No underlying type is set. Please use any of the `models.WlanDynamicPskDefaultVlanIdContainer.From*` functions to initialize the WlanDynamicPskDefaultVlanId object.")
    }
    return json.Marshal(w.toMap())
}

// toMap converts the WlanDynamicPskDefaultVlanId object to a map representation for JSON marshaling.
func (w *WlanDynamicPskDefaultVlanId) toMap() any {
    switch obj := w.value.(type) {
    case *string:
        return *obj
    case *int:
        return *obj
    }
    return nil
}

// UnmarshalJSON implements the json.Unmarshaler interface for WlanDynamicPskDefaultVlanId.
// It customizes the JSON unmarshaling process for WlanDynamicPskDefaultVlanId objects.
func (w *WlanDynamicPskDefaultVlanId) UnmarshalJSON(input []byte) error {
    result, err := UnmarshallOneOf(input,
        NewTypeHolder(new(string), false, &w.isString),
        NewTypeHolder(new(int), false, &w.isNumber),
    )
    
    w.value = result
    return err
}

func (w *WlanDynamicPskDefaultVlanId) AsString() (
    *string,
    bool) {
    if !w.isString {
        return nil, false
    }
    return w.value.(*string), true
}

func (w *WlanDynamicPskDefaultVlanId) AsNumber() (
    *int,
    bool) {
    if !w.isNumber {
        return nil, false
    }
    return w.value.(*int), true
}

// internalWlanDynamicPskDefaultVlanId represents a wlanDynamicPskDefaultVlanId struct.
// This is a container for one-of cases.
type internalWlanDynamicPskDefaultVlanId struct {}

var WlanDynamicPskDefaultVlanIdContainer internalWlanDynamicPskDefaultVlanId

// The internalWlanDynamicPskDefaultVlanId instance, wrapping the provided string value.
func (w *internalWlanDynamicPskDefaultVlanId) FromString(val string) WlanDynamicPskDefaultVlanId {
    return WlanDynamicPskDefaultVlanId{value: &val}
}

// The internalWlanDynamicPskDefaultVlanId instance, wrapping the provided int value.
func (w *internalWlanDynamicPskDefaultVlanId) FromNumber(val int) WlanDynamicPskDefaultVlanId {
    return WlanDynamicPskDefaultVlanId{value: &val}
}
