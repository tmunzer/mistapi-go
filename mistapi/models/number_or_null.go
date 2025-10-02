// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"errors"
	"fmt"
)

// NumberOrNull represents a NumberOrNull struct.
type NumberOrNull struct {
	value       any
	isPrecision bool
}

// String implements the fmt.Stringer interface for NumberOrNull,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (n NumberOrNull) String() string {
	return fmt.Sprintf("%v", n.value)
}

// MarshalJSON implements the json.Marshaler interface for NumberOrNull.
// It customizes the JSON marshaling process for NumberOrNull objects.
func (n NumberOrNull) MarshalJSON() (
	[]byte,
	error) {
	if n.value == nil {
		return nil, errors.New("No underlying type is set. Please use any of the `models.NumberOrNullContainer.From*` functions to initialize the NumberOrNull object.")
	}
	return json.Marshal(n.toMap())
}

// toMap converts the NumberOrNull object to a map representation for JSON marshaling.
func (n *NumberOrNull) toMap() any {
	switch obj := n.value.(type) {
	case *float64:
		return *obj
	}
	return nil
}

// UnmarshalJSON implements the json.Unmarshaler interface for NumberOrNull.
// It customizes the JSON unmarshaling process for NumberOrNull objects.
func (n *NumberOrNull) UnmarshalJSON(input []byte) error {
	result, err := UnmarshallOneOf(input,
		NewTypeHolder(new(float64), false, &n.isPrecision),
	)

	n.value = result
	return err
}

func (n *NumberOrNull) AsPrecision() (
	*float64,
	bool) {
	if !n.isPrecision {
		return nil, false
	}
	return n.value.(*float64), true
}

// internalNumberOrNull represents a numberOrNull struct.
type internalNumberOrNull struct{}

var NumberOrNullContainer internalNumberOrNull

// The internalNumberOrNull instance, wrapping the provided float64 value.
func (n *internalNumberOrNull) FromPrecision(val float64) NumberOrNull {
	return NumberOrNull{value: &val}
}
