// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
    "encoding/json"
    "fmt"
)

// ConstInsightMetricsPropertyReportDuration represents a ConstInsightMetricsPropertyReportDuration struct.
type ConstInsightMetricsPropertyReportDuration struct {
    Interval             *int                   `json:"interval,omitempty"`
    MaxAge               *int                   `json:"max_age,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for ConstInsightMetricsPropertyReportDuration,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (c ConstInsightMetricsPropertyReportDuration) String() string {
    return fmt.Sprintf(
    	"ConstInsightMetricsPropertyReportDuration[Interval=%v, MaxAge=%v, AdditionalProperties=%v]",
    	c.Interval, c.MaxAge, c.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for ConstInsightMetricsPropertyReportDuration.
// It customizes the JSON marshaling process for ConstInsightMetricsPropertyReportDuration objects.
func (c ConstInsightMetricsPropertyReportDuration) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(c.AdditionalProperties,
        "interval", "max_age"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(c.toMap())
}

// toMap converts the ConstInsightMetricsPropertyReportDuration object to a map representation for JSON marshaling.
func (c ConstInsightMetricsPropertyReportDuration) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, c.AdditionalProperties)
    if c.Interval != nil {
        structMap["interval"] = c.Interval
    }
    if c.MaxAge != nil {
        structMap["max_age"] = c.MaxAge
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
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "interval", "max_age")
    if err != nil {
    	return err
    }
    c.AdditionalProperties = additionalProperties
    
    c.Interval = temp.Interval
    c.MaxAge = temp.MaxAge
    return nil
}

// tempConstInsightMetricsPropertyReportDuration is a temporary struct used for validating the fields of ConstInsightMetricsPropertyReportDuration.
type tempConstInsightMetricsPropertyReportDuration  struct {
    Interval *int `json:"interval,omitempty"`
    MaxAge   *int `json:"max_age,omitempty"`
}
