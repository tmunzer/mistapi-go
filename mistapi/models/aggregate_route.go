// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"fmt"
)

// AggregateRoute represents a AggregateRoute struct.
type AggregateRoute struct {
	Discard              *bool                  `json:"discard,omitempty"`
	Metric               Optional[int]          `json:"metric"`
	Preference           Optional[int]          `json:"preference"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for AggregateRoute,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (a AggregateRoute) String() string {
	return fmt.Sprintf(
		"AggregateRoute[Discard=%v, Metric=%v, Preference=%v, AdditionalProperties=%v]",
		a.Discard, a.Metric, a.Preference, a.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for AggregateRoute.
// It customizes the JSON marshaling process for AggregateRoute objects.
func (a AggregateRoute) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(a.AdditionalProperties,
		"discard", "metric", "preference"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(a.toMap())
}

// toMap converts the AggregateRoute object to a map representation for JSON marshaling.
func (a AggregateRoute) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, a.AdditionalProperties)
	if a.Discard != nil {
		structMap["discard"] = a.Discard
	}
	if a.Metric.IsValueSet() {
		if a.Metric.Value() != nil {
			structMap["metric"] = a.Metric.Value()
		} else {
			structMap["metric"] = nil
		}
	}
	if a.Preference.IsValueSet() {
		if a.Preference.Value() != nil {
			structMap["preference"] = a.Preference.Value()
		} else {
			structMap["preference"] = nil
		}
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for AggregateRoute.
// It customizes the JSON unmarshaling process for AggregateRoute objects.
func (a *AggregateRoute) UnmarshalJSON(input []byte) error {
	var temp tempAggregateRoute
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "discard", "metric", "preference")
	if err != nil {
		return err
	}
	a.AdditionalProperties = additionalProperties

	a.Discard = temp.Discard
	a.Metric = temp.Metric
	a.Preference = temp.Preference
	return nil
}

// tempAggregateRoute is a temporary struct used for validating the fields of AggregateRoute.
type tempAggregateRoute struct {
	Discard    *bool         `json:"discard,omitempty"`
	Metric     Optional[int] `json:"metric"`
	Preference Optional[int] `json:"preference"`
}
