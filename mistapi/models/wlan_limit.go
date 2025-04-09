package models

import (
    "encoding/json"
    "errors"
    "fmt"
)

// WlanLimit represents a WlanLimit struct.
// In kbps, value from 1 to 999000
type WlanLimit struct {
    value    any
    isNumber bool
    isString bool
}

// String implements the fmt.Stringer interface for WlanLimit,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (w WlanLimit) String() string {
    return fmt.Sprintf("%v", w.value)
}

// MarshalJSON implements the json.Marshaler interface for WlanLimit.
// It customizes the JSON marshaling process for WlanLimit objects.
func (w WlanLimit) MarshalJSON() (
    []byte,
    error) {
    if w.value == nil {
        return nil, errors.New("No underlying type is set. Please use any of the `models.WlanLimitContainer.From*` functions to initialize the WlanLimit object.")
    }
    return json.Marshal(w.toMap())
}

// toMap converts the WlanLimit object to a map representation for JSON marshaling.
func (w *WlanLimit) toMap() any {
    switch obj := w.value.(type) {
    case *int:
        return *obj
    case *string:
        return *obj
    }
    return nil
}

// UnmarshalJSON implements the json.Unmarshaler interface for WlanLimit.
// It customizes the JSON unmarshaling process for WlanLimit objects.
func (w *WlanLimit) UnmarshalJSON(input []byte) error {
    result, err := UnmarshallAnyOf(input,
        NewTypeHolder(new(int), false, &w.isNumber),
        NewTypeHolder(new(string), false, &w.isString),
    )
    
    w.value = result
    return err
}

func (w *WlanLimit) AsNumber() (
    *int,
    bool) {
    if !w.isNumber {
        return nil, false
    }
    return w.value.(*int), true
}

func (w *WlanLimit) AsString() (
    *string,
    bool) {
    if !w.isString {
        return nil, false
    }
    return w.value.(*string), true
}

// internalWlanLimit represents a wlanLimit struct.
// In kbps, value from 1 to 999000
type internalWlanLimit struct {}

var WlanLimitContainer internalWlanLimit

// The internalWlanLimit instance, wrapping the provided int value.
func (w *internalWlanLimit) FromNumber(val int) WlanLimit {
    return WlanLimit{value: &val}
}

// The internalWlanLimit instance, wrapping the provided string value.
func (w *internalWlanLimit) FromString(val string) WlanLimit {
    return WlanLimit{value: &val}
}
