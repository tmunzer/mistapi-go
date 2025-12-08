// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"errors"
	"fmt"
)

// InsightMetricsResultsItem represents a InsightMetricsResultsItem struct.
type InsightMetricsResultsItem struct {
	value       any
	isPrecision bool
	isObject    bool
}

// String implements the fmt.Stringer interface for InsightMetricsResultsItem,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (i InsightMetricsResultsItem) String() string {
	return fmt.Sprintf("%v", i.value)
}

// MarshalJSON implements the json.Marshaler interface for InsightMetricsResultsItem.
// It customizes the JSON marshaling process for InsightMetricsResultsItem objects.
func (i InsightMetricsResultsItem) MarshalJSON() (
	[]byte,
	error) {
	if i.value == nil {
		return nil, errors.New("No underlying type is set. Please use any of the `models.InsightMetricsResultsItemContainer.From*` functions to initialize the InsightMetricsResultsItem object.")
	}
	return json.Marshal(i.toMap())
}

// toMap converts the InsightMetricsResultsItem object to a map representation for JSON marshaling.
func (i *InsightMetricsResultsItem) toMap() any {
	switch obj := i.value.(type) {
	case *float64:
		return *obj
	case *interface{}:
		return *obj
	}
	return nil
}

// UnmarshalJSON implements the json.Unmarshaler interface for InsightMetricsResultsItem.
// It customizes the JSON unmarshaling process for InsightMetricsResultsItem objects.
func (i *InsightMetricsResultsItem) UnmarshalJSON(input []byte) error {
	result, err := UnmarshallAnyOf(input,
		NewTypeHolder(new(float64), false, &i.isPrecision),
		NewTypeHolder(new(interface{}), false, &i.isObject),
	)

	i.value = result
	return err
}

func (i *InsightMetricsResultsItem) AsPrecision() (
	*float64,
	bool) {
	if !i.isPrecision {
		return nil, false
	}
	return i.value.(*float64), true
}

func (i *InsightMetricsResultsItem) AsObject() (
	*interface{},
	bool) {
	if !i.isObject {
		return nil, false
	}
	return i.value.(*interface{}), true
}

// internalInsightMetricsResultsItem represents a insightMetricsResultsItem struct.
type internalInsightMetricsResultsItem struct{}

var InsightMetricsResultsItemContainer internalInsightMetricsResultsItem

// The internalInsightMetricsResultsItem instance, wrapping the provided float64 value.
func (i *internalInsightMetricsResultsItem) FromPrecision(val float64) InsightMetricsResultsItem {
	return InsightMetricsResultsItem{value: &val}
}

// The internalInsightMetricsResultsItem instance, wrapping the provided interface{} value.
func (i *internalInsightMetricsResultsItem) FromObject(val interface{}) InsightMetricsResultsItem {
	return InsightMetricsResultsItem{value: &val}
}
