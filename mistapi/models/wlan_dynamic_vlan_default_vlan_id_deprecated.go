package models

import (
	"encoding/json"
	"errors"
	"strings"
)

// WlanDynamicVlanDefaultVlanIdDeprecated represents a WlanDynamicVlanDefaultVlanIdDeprecated struct.
// vlan_id to use when there’s no match from RADIUS
type WlanDynamicVlanDefaultVlanIdDeprecated struct {
	value    any
	isString bool
	isNumber bool
}

// String converts the WlanDynamicVlanDefaultVlanIdDeprecated object to a string representation.
func (w WlanDynamicVlanDefaultVlanIdDeprecated) String() string {
	if bytes, err := json.Marshal(w.value); err == nil {
		return strings.Trim(string(bytes), "\"")
	}
	return ""
}

// MarshalJSON implements the json.Marshaler interface for WlanDynamicVlanDefaultVlanIdDeprecated.
// It customizes the JSON marshaling process for WlanDynamicVlanDefaultVlanIdDeprecated objects.
func (w WlanDynamicVlanDefaultVlanIdDeprecated) MarshalJSON() (
	[]byte,
	error) {
	if w.value == nil {
		return nil, errors.New("No underlying type is set. Please use any of the `models.WlanDynamicVlanDefaultVlanIdDeprecatedContainer.From*` functions to initialize the WlanDynamicVlanDefaultVlanIdDeprecated object.")
	}
	return json.Marshal(w.toMap())
}

// toMap converts the WlanDynamicVlanDefaultVlanIdDeprecated object to a map representation for JSON marshaling.
func (w *WlanDynamicVlanDefaultVlanIdDeprecated) toMap() any {
	switch obj := w.value.(type) {
	case *string:
		return *obj
	case *int:
		return *obj
	}
	return nil
}

// UnmarshalJSON implements the json.Unmarshaler interface for WlanDynamicVlanDefaultVlanIdDeprecated.
// It customizes the JSON unmarshaling process for WlanDynamicVlanDefaultVlanIdDeprecated objects.
func (w *WlanDynamicVlanDefaultVlanIdDeprecated) UnmarshalJSON(input []byte) error {
	result, err := UnmarshallOneOf(input,
		NewTypeHolder(new(string), false, &w.isString),
		NewTypeHolder(new(int), false, &w.isNumber),
	)

	w.value = result
	return err
}

func (w *WlanDynamicVlanDefaultVlanIdDeprecated) AsString() (
	*string,
	bool) {
	if !w.isString {
		return nil, false
	}
	return w.value.(*string), true
}

func (w *WlanDynamicVlanDefaultVlanIdDeprecated) AsNumber() (
	*int,
	bool) {
	if !w.isNumber {
		return nil, false
	}
	return w.value.(*int), true
}

// internalWlanDynamicVlanDefaultVlanIdDeprecated represents a wlanDynamicVlanDefaultVlanIdDeprecated struct.
// vlan_id to use when there’s no match from RADIUS
type internalWlanDynamicVlanDefaultVlanIdDeprecated struct{}

var WlanDynamicVlanDefaultVlanIdDeprecatedContainer internalWlanDynamicVlanDefaultVlanIdDeprecated

// The internalWlanDynamicVlanDefaultVlanIdDeprecated instance, wrapping the provided string value.
func (w *internalWlanDynamicVlanDefaultVlanIdDeprecated) FromString(val string) WlanDynamicVlanDefaultVlanIdDeprecated {
	return WlanDynamicVlanDefaultVlanIdDeprecated{value: &val}
}

// The internalWlanDynamicVlanDefaultVlanIdDeprecated instance, wrapping the provided int value.
func (w *internalWlanDynamicVlanDefaultVlanIdDeprecated) FromNumber(val int) WlanDynamicVlanDefaultVlanIdDeprecated {
	return WlanDynamicVlanDefaultVlanIdDeprecated{value: &val}
}
