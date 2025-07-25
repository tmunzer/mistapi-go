// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
    "encoding/json"
    "fmt"
)

// ResponseSwitchMetricsConfigSuccess represents a ResponseSwitchMetricsConfigSuccess struct.
type ResponseSwitchMetricsConfigSuccess struct {
    Details              *ResponseSwitchMetricsConfigSuccessDetails `json:"details,omitempty"`
    Score                *int                                       `json:"score,omitempty"`
    TotalSwitchCount     *int                                       `json:"total_switch_count,omitempty"`
    AdditionalProperties map[string]interface{}                     `json:"_"`
}

// String implements the fmt.Stringer interface for ResponseSwitchMetricsConfigSuccess,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (r ResponseSwitchMetricsConfigSuccess) String() string {
    return fmt.Sprintf(
    	"ResponseSwitchMetricsConfigSuccess[Details=%v, Score=%v, TotalSwitchCount=%v, AdditionalProperties=%v]",
    	r.Details, r.Score, r.TotalSwitchCount, r.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for ResponseSwitchMetricsConfigSuccess.
// It customizes the JSON marshaling process for ResponseSwitchMetricsConfigSuccess objects.
func (r ResponseSwitchMetricsConfigSuccess) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(r.AdditionalProperties,
        "details", "score", "total_switch_count"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(r.toMap())
}

// toMap converts the ResponseSwitchMetricsConfigSuccess object to a map representation for JSON marshaling.
func (r ResponseSwitchMetricsConfigSuccess) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, r.AdditionalProperties)
    if r.Details != nil {
        structMap["details"] = r.Details.toMap()
    }
    if r.Score != nil {
        structMap["score"] = r.Score
    }
    if r.TotalSwitchCount != nil {
        structMap["total_switch_count"] = r.TotalSwitchCount
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ResponseSwitchMetricsConfigSuccess.
// It customizes the JSON unmarshaling process for ResponseSwitchMetricsConfigSuccess objects.
func (r *ResponseSwitchMetricsConfigSuccess) UnmarshalJSON(input []byte) error {
    var temp tempResponseSwitchMetricsConfigSuccess
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "details", "score", "total_switch_count")
    if err != nil {
    	return err
    }
    r.AdditionalProperties = additionalProperties
    
    r.Details = temp.Details
    r.Score = temp.Score
    r.TotalSwitchCount = temp.TotalSwitchCount
    return nil
}

// tempResponseSwitchMetricsConfigSuccess is a temporary struct used for validating the fields of ResponseSwitchMetricsConfigSuccess.
type tempResponseSwitchMetricsConfigSuccess  struct {
    Details          *ResponseSwitchMetricsConfigSuccessDetails `json:"details,omitempty"`
    Score            *int                                       `json:"score,omitempty"`
    TotalSwitchCount *int                                       `json:"total_switch_count,omitempty"`
}
