package models

import (
    "encoding/json"
    "errors"
    "fmt"
)

// WxlanTagVlanId represents a WxlanTagVlanId struct.
// if `type`==`vlan_id`, VLAN ID or variable
type WxlanTagVlanId struct {
    value    any
    isString bool
    isNumber bool
}

// String implements the fmt.Stringer interface for WxlanTagVlanId,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (w WxlanTagVlanId) String() string {
    return fmt.Sprintf("%v", w.value)
}

// MarshalJSON implements the json.Marshaler interface for WxlanTagVlanId.
// It customizes the JSON marshaling process for WxlanTagVlanId objects.
func (w WxlanTagVlanId) MarshalJSON() (
    []byte,
    error) {
    if w.value == nil {
        return nil, errors.New("No underlying type is set. Please use any of the `models.WxlanTagVlanIdContainer.From*` functions to initialize the WxlanTagVlanId object.")
    }
    return json.Marshal(w.toMap())
}

// toMap converts the WxlanTagVlanId object to a map representation for JSON marshaling.
func (w *WxlanTagVlanId) toMap() any {
    switch obj := w.value.(type) {
    case *string:
        return *obj
    case *int:
        return *obj
    }
    return nil
}

// UnmarshalJSON implements the json.Unmarshaler interface for WxlanTagVlanId.
// It customizes the JSON unmarshaling process for WxlanTagVlanId objects.
func (w *WxlanTagVlanId) UnmarshalJSON(input []byte) error {
    result, err := UnmarshallOneOf(input,
        NewTypeHolder(new(string), false, &w.isString),
        NewTypeHolder(new(int), false, &w.isNumber),
    )
    
    w.value = result
    return err
}

func (w *WxlanTagVlanId) AsString() (
    *string,
    bool) {
    if !w.isString {
        return nil, false
    }
    return w.value.(*string), true
}

func (w *WxlanTagVlanId) AsNumber() (
    *int,
    bool) {
    if !w.isNumber {
        return nil, false
    }
    return w.value.(*int), true
}

// internalWxlanTagVlanId represents a wxlanTagVlanId struct.
// if `type`==`vlan_id`, VLAN ID or variable
type internalWxlanTagVlanId struct {}

var WxlanTagVlanIdContainer internalWxlanTagVlanId

// The internalWxlanTagVlanId instance, wrapping the provided string value.
func (w *internalWxlanTagVlanId) FromString(val string) WxlanTagVlanId {
    return WxlanTagVlanId{value: &val}
}

// The internalWxlanTagVlanId instance, wrapping the provided int value.
func (w *internalWxlanTagVlanId) FromNumber(val int) WxlanTagVlanId {
    return WxlanTagVlanId{value: &val}
}
