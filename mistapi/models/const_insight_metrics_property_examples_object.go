// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"errors"
	"fmt"
)

// ConstInsightMetricsPropertyExamplesObject represents a ConstInsightMetricsPropertyExamplesObject struct.
// Object containing named fields, each with array of example values
type ConstInsightMetricsPropertyExamplesObject struct {
	value       any
	isNumber    bool
	isPrecision bool
	isString    bool
	isBoolean   bool
	isObject    bool
}

// String implements the fmt.Stringer interface for ConstInsightMetricsPropertyExamplesObject,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (c ConstInsightMetricsPropertyExamplesObject) String() string {
	return fmt.Sprintf("%v", c.value)
}

// MarshalJSON implements the json.Marshaler interface for ConstInsightMetricsPropertyExamplesObject.
// It customizes the JSON marshaling process for ConstInsightMetricsPropertyExamplesObject objects.
func (c ConstInsightMetricsPropertyExamplesObject) MarshalJSON() (
	[]byte,
	error) {
	if c.value == nil {
		return nil, errors.New("No underlying type is set. Please use any of the `models.ConstInsightMetricsPropertyExamplesObjectContainer.From*` functions to initialize the ConstInsightMetricsPropertyExamplesObject object.")
	}
	return json.Marshal(c.toMap())
}

// toMap converts the ConstInsightMetricsPropertyExamplesObject object to a map representation for JSON marshaling.
func (c *ConstInsightMetricsPropertyExamplesObject) toMap() any {
	switch obj := c.value.(type) {
	case *int:
		return *obj
	case *float64:
		return *obj
	case *string:
		return *obj
	case *bool:
		return *obj
	case *interface{}:
		return *obj
	}
	return nil
}

// UnmarshalJSON implements the json.Unmarshaler interface for ConstInsightMetricsPropertyExamplesObject.
// It customizes the JSON unmarshaling process for ConstInsightMetricsPropertyExamplesObject objects.
func (c *ConstInsightMetricsPropertyExamplesObject) UnmarshalJSON(input []byte) error {
	result, err := UnmarshallAnyOf(input,
		NewTypeHolder(new(int), false, &c.isNumber),
		NewTypeHolder(new(float64), false, &c.isPrecision),
		NewTypeHolder(new(string), false, &c.isString),
		NewTypeHolder(new(bool), false, &c.isBoolean),
		NewTypeHolder(new(interface{}), false, &c.isObject),
	)

	c.value = result
	return err
}

func (c *ConstInsightMetricsPropertyExamplesObject) AsNumber() (
	*int,
	bool) {
	if !c.isNumber {
		return nil, false
	}
	return c.value.(*int), true
}

func (c *ConstInsightMetricsPropertyExamplesObject) AsPrecision() (
	*float64,
	bool) {
	if !c.isPrecision {
		return nil, false
	}
	return c.value.(*float64), true
}

func (c *ConstInsightMetricsPropertyExamplesObject) AsString() (
	*string,
	bool) {
	if !c.isString {
		return nil, false
	}
	return c.value.(*string), true
}

func (c *ConstInsightMetricsPropertyExamplesObject) AsBoolean() (
	*bool,
	bool) {
	if !c.isBoolean {
		return nil, false
	}
	return c.value.(*bool), true
}

func (c *ConstInsightMetricsPropertyExamplesObject) AsObject() (
	*interface{},
	bool) {
	if !c.isObject {
		return nil, false
	}
	return c.value.(*interface{}), true
}

// internalConstInsightMetricsPropertyExamplesObject represents a constInsightMetricsPropertyExamplesObject struct.
// Object containing named fields, each with array of example values
type internalConstInsightMetricsPropertyExamplesObject struct{}

var ConstInsightMetricsPropertyExamplesObjectContainer internalConstInsightMetricsPropertyExamplesObject

// The internalConstInsightMetricsPropertyExamplesObject instance, wrapping the provided int value.
func (c *internalConstInsightMetricsPropertyExamplesObject) FromNumber(val int) ConstInsightMetricsPropertyExamplesObject {
	return ConstInsightMetricsPropertyExamplesObject{value: &val}
}

// The internalConstInsightMetricsPropertyExamplesObject instance, wrapping the provided float64 value.
func (c *internalConstInsightMetricsPropertyExamplesObject) FromPrecision(val float64) ConstInsightMetricsPropertyExamplesObject {
	return ConstInsightMetricsPropertyExamplesObject{value: &val}
}

// The internalConstInsightMetricsPropertyExamplesObject instance, wrapping the provided string value.
func (c *internalConstInsightMetricsPropertyExamplesObject) FromString(val string) ConstInsightMetricsPropertyExamplesObject {
	return ConstInsightMetricsPropertyExamplesObject{value: &val}
}

// The internalConstInsightMetricsPropertyExamplesObject instance, wrapping the provided bool value.
func (c *internalConstInsightMetricsPropertyExamplesObject) FromBoolean(val bool) ConstInsightMetricsPropertyExamplesObject {
	return ConstInsightMetricsPropertyExamplesObject{value: &val}
}

// The internalConstInsightMetricsPropertyExamplesObject instance, wrapping the provided interface{} value.
func (c *internalConstInsightMetricsPropertyExamplesObject) FromObject(val interface{}) ConstInsightMetricsPropertyExamplesObject {
	return ConstInsightMetricsPropertyExamplesObject{value: &val}
}
