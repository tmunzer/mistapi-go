package models

import (
    "encoding/json"
    "errors"
    "fmt"
)

// WlanVlanIds represents a WlanVlanIds struct.
type WlanVlanIds struct {
    value                        any
    isString                     bool
    isArrayOfVlanIdWithVariable5 bool
}

// String implements the fmt.Stringer interface for WlanVlanIds,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (w WlanVlanIds) String() string {
    return fmt.Sprintf("%v", w.value)
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
    case *[]VlanIdWithVariable:
        return *obj
    }
    return nil
}

// UnmarshalJSON implements the json.Unmarshaler interface for WlanVlanIds.
// It customizes the JSON unmarshaling process for WlanVlanIds objects.
func (w *WlanVlanIds) UnmarshalJSON(input []byte) error {
    result, err := UnmarshallOneOf(input,
        NewTypeHolder(new(string), false, &w.isString),
        NewTypeHolder(&[]VlanIdWithVariable{}, false, &w.isArrayOfVlanIdWithVariable5),
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

func (w *WlanVlanIds) AsArrayOfVlanIdWithVariable5() (
    *[]VlanIdWithVariable,
    bool) {
    if !w.isArrayOfVlanIdWithVariable5 {
        return nil, false
    }
    return w.value.(*[]VlanIdWithVariable), true
}

// internalWlanVlanIds represents a wlanVlanIds struct.
type internalWlanVlanIds struct {}

var WlanVlanIdsContainer internalWlanVlanIds

// The internalWlanVlanIds instance, wrapping the provided string value.
func (w *internalWlanVlanIds) FromString(val string) WlanVlanIds {
    return WlanVlanIds{value: &val}
}

// The internalWlanVlanIds instance, wrapping the provided []VlanIdWithVariable value.
func (w *internalWlanVlanIds) FromArrayOfVlanIdWithVariable5(val []VlanIdWithVariable) WlanVlanIds {
    return WlanVlanIds{value: &val}
}
