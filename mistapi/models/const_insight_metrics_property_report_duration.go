// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"fmt"
)

// ConstInsightMetricsPropertyReportDuration represents a ConstInsightMetricsPropertyReportDuration struct.
type ConstInsightMetricsPropertyReportDuration struct {
	Duration             *int                   `json:"duration,omitempty"`
	Interval             *int                   `json:"interval,omitempty"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for ConstInsightMetricsPropertyReportDuration,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (c ConstInsightMetricsPropertyReportDuration) String() string {
	return fmt.Sprintf(
		"ConstInsightMetricsPropertyReportDuration[Duration=%v, Interval=%v, AdditionalProperties=%v]",
		c.Duration, c.Interval, c.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for ConstInsightMetricsPropertyReportDuration.
// It customizes the JSON marshaling process for ConstInsightMetricsPropertyReportDuration objects.
func (c ConstInsightMetricsPropertyReportDuration) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(c.AdditionalProperties,
		"duration", "interval"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(c.toMap())
}

// toMap converts the ConstInsightMetricsPropertyReportDuration object to a map representation for JSON marshaling.
func (c ConstInsightMetricsPropertyReportDuration) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, c.AdditionalProperties)
	if c.Duration != nil {
		structMap["duration"] = c.Duration
	}
	if c.Interval != nil {
		structMap["interval"] = c.Interval
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ConstInsightMetricsPropertyReportDuration.
// It customizes the JSON unmarshaling process for ConstInsightMetricsPropertyReportDuration objects.
func (c *ConstInsightMetricsPropertyReportDuration) UnmarshalJSON(input []byte) error {
	var temp tempConstInsightMetricsPropertyReportDuration
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "duration", "interval")
	if err != nil {
		return err
	}
	c.AdditionalProperties = additionalProperties

	c.Duration = temp.Duration
	c.Interval = temp.Interval
	return nil
}

// tempConstInsightMetricsPropertyReportDuration is a temporary struct used for validating the fields of ConstInsightMetricsPropertyReportDuration.
type tempConstInsightMetricsPropertyReportDuration struct {
	Duration *int `json:"duration,omitempty"`
	Interval *int `json:"interval,omitempty"`
}
