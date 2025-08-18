// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"fmt"
)

// RrmBandMetricInterference represents a RrmBandMetricInterference struct.
type RrmBandMetricInterference struct {
	Radar                *float64               `json:"radar,omitempty"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for RrmBandMetricInterference,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (r RrmBandMetricInterference) String() string {
	return fmt.Sprintf(
		"RrmBandMetricInterference[Radar=%v, AdditionalProperties=%v]",
		r.Radar, r.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for RrmBandMetricInterference.
// It customizes the JSON marshaling process for RrmBandMetricInterference objects.
func (r RrmBandMetricInterference) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(r.AdditionalProperties,
		"radar"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(r.toMap())
}

// toMap converts the RrmBandMetricInterference object to a map representation for JSON marshaling.
func (r RrmBandMetricInterference) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, r.AdditionalProperties)
	if r.Radar != nil {
		structMap["radar"] = r.Radar
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for RrmBandMetricInterference.
// It customizes the JSON unmarshaling process for RrmBandMetricInterference objects.
func (r *RrmBandMetricInterference) UnmarshalJSON(input []byte) error {
	var temp tempRrmBandMetricInterference
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "radar")
	if err != nil {
		return err
	}
	r.AdditionalProperties = additionalProperties

	r.Radar = temp.Radar
	return nil
}

// tempRrmBandMetricInterference is a temporary struct used for validating the fields of RrmBandMetricInterference.
type tempRrmBandMetricInterference struct {
	Radar *float64 `json:"radar,omitempty"`
}
