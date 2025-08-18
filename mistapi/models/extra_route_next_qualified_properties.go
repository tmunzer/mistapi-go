// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"fmt"
)

// ExtraRouteNextQualifiedProperties represents a ExtraRouteNextQualifiedProperties struct.
type ExtraRouteNextQualifiedProperties struct {
	Metric               Optional[int]          `json:"metric"`
	Preference           Optional[int]          `json:"preference"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for ExtraRouteNextQualifiedProperties,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (e ExtraRouteNextQualifiedProperties) String() string {
	return fmt.Sprintf(
		"ExtraRouteNextQualifiedProperties[Metric=%v, Preference=%v, AdditionalProperties=%v]",
		e.Metric, e.Preference, e.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for ExtraRouteNextQualifiedProperties.
// It customizes the JSON marshaling process for ExtraRouteNextQualifiedProperties objects.
func (e ExtraRouteNextQualifiedProperties) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(e.AdditionalProperties,
		"metric", "preference"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(e.toMap())
}

// toMap converts the ExtraRouteNextQualifiedProperties object to a map representation for JSON marshaling.
func (e ExtraRouteNextQualifiedProperties) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, e.AdditionalProperties)
	if e.Metric.IsValueSet() {
		if e.Metric.Value() != nil {
			structMap["metric"] = e.Metric.Value()
		} else {
			structMap["metric"] = nil
		}
	}
	if e.Preference.IsValueSet() {
		if e.Preference.Value() != nil {
			structMap["preference"] = e.Preference.Value()
		} else {
			structMap["preference"] = nil
		}
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ExtraRouteNextQualifiedProperties.
// It customizes the JSON unmarshaling process for ExtraRouteNextQualifiedProperties objects.
func (e *ExtraRouteNextQualifiedProperties) UnmarshalJSON(input []byte) error {
	var temp tempExtraRouteNextQualifiedProperties
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "metric", "preference")
	if err != nil {
		return err
	}
	e.AdditionalProperties = additionalProperties

	e.Metric = temp.Metric
	e.Preference = temp.Preference
	return nil
}

// tempExtraRouteNextQualifiedProperties is a temporary struct used for validating the fields of ExtraRouteNextQualifiedProperties.
type tempExtraRouteNextQualifiedProperties struct {
	Metric     Optional[int] `json:"metric"`
	Preference Optional[int] `json:"preference"`
}
