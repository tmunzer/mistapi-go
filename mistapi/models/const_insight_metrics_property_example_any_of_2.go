// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"errors"
	"fmt"
)

// ConstInsightMetricsPropertyExampleAnyOf2 represents a ConstInsightMetricsPropertyExampleAnyOf2 struct.
type ConstInsightMetricsPropertyExampleAnyOf2 struct {
	value                                                   any
	isArrayOfConstInsightMetricsPropertyExample             bool
	isMapOfArrayOfConstInsightMetricsPropertyExamplesObject bool
}

// String implements the fmt.Stringer interface for ConstInsightMetricsPropertyExampleAnyOf2,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (c ConstInsightMetricsPropertyExampleAnyOf2) String() string {
	return fmt.Sprintf("%v", c.value)
}

// MarshalJSON implements the json.Marshaler interface for ConstInsightMetricsPropertyExampleAnyOf2.
// It customizes the JSON marshaling process for ConstInsightMetricsPropertyExampleAnyOf2 objects.
func (c ConstInsightMetricsPropertyExampleAnyOf2) MarshalJSON() (
	[]byte,
	error) {
	if c.value == nil {
		return nil, errors.New("No underlying type is set. Please use any of the `models.ConstInsightMetricsPropertyExampleAnyOf2Container.From*` functions to initialize the ConstInsightMetricsPropertyExampleAnyOf2 object.")
	}
	return json.Marshal(c.toMap())
}

// toMap converts the ConstInsightMetricsPropertyExampleAnyOf2 object to a map representation for JSON marshaling.
func (c *ConstInsightMetricsPropertyExampleAnyOf2) toMap() any {
	switch obj := c.value.(type) {
	case *[]ConstInsightMetricsPropertyExample:
		return *obj
	case *map[string][]ConstInsightMetricsPropertyExamplesObject:
		return *obj
	}
	return nil
}

// UnmarshalJSON implements the json.Unmarshaler interface for ConstInsightMetricsPropertyExampleAnyOf2.
// It customizes the JSON unmarshaling process for ConstInsightMetricsPropertyExampleAnyOf2 objects.
func (c *ConstInsightMetricsPropertyExampleAnyOf2) UnmarshalJSON(input []byte) error {
	result, err := UnmarshallAnyOf(input,
		NewTypeHolder(&[]ConstInsightMetricsPropertyExample{}, false, &c.isArrayOfConstInsightMetricsPropertyExample),
		NewTypeHolder(&map[string][]ConstInsightMetricsPropertyExamplesObject{}, false, &c.isMapOfArrayOfConstInsightMetricsPropertyExamplesObject),
	)

	c.value = result
	return err
}

func (c *ConstInsightMetricsPropertyExampleAnyOf2) AsArrayOfConstInsightMetricsPropertyExample() (
	*[]ConstInsightMetricsPropertyExample,
	bool) {
	if !c.isArrayOfConstInsightMetricsPropertyExample {
		return nil, false
	}
	return c.value.(*[]ConstInsightMetricsPropertyExample), true
}

func (c *ConstInsightMetricsPropertyExampleAnyOf2) AsMapOfArrayOfConstInsightMetricsPropertyExamplesObject() (
	*map[string][]ConstInsightMetricsPropertyExamplesObject,
	bool) {
	if !c.isMapOfArrayOfConstInsightMetricsPropertyExamplesObject {
		return nil, false
	}
	return c.value.(*map[string][]ConstInsightMetricsPropertyExamplesObject), true
}

// internalConstInsightMetricsPropertyExampleAnyOf2 represents a constInsightMetricsPropertyExampleAnyOf2 struct.
type internalConstInsightMetricsPropertyExampleAnyOf2 struct{}

var ConstInsightMetricsPropertyExampleAnyOf2Container internalConstInsightMetricsPropertyExampleAnyOf2

// The internalConstInsightMetricsPropertyExampleAnyOf2 instance, wrapping the provided []ConstInsightMetricsPropertyExample value.
func (c *internalConstInsightMetricsPropertyExampleAnyOf2) FromArrayOfConstInsightMetricsPropertyExample(val []ConstInsightMetricsPropertyExample) ConstInsightMetricsPropertyExampleAnyOf2 {
	return ConstInsightMetricsPropertyExampleAnyOf2{value: &val}
}

// The internalConstInsightMetricsPropertyExampleAnyOf2 instance, wrapping the provided map[string][]ConstInsightMetricsPropertyExamplesObject value.
func (c *internalConstInsightMetricsPropertyExampleAnyOf2) FromMapOfArrayOfConstInsightMetricsPropertyExamplesObject(val map[string][]ConstInsightMetricsPropertyExamplesObject) ConstInsightMetricsPropertyExampleAnyOf2 {
	return ConstInsightMetricsPropertyExampleAnyOf2{value: &val}
}
