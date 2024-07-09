package models

import (
    "encoding/json"
)

// ResponseSwitchMetricsActivePortsSummary represents a ResponseSwitchMetricsActivePortsSummary struct.
type ResponseSwitchMetricsActivePortsSummary struct {
    Details              *SwitchMetricsActivePortsSummaryDetails `json:"details,omitempty"`
    Score                *int                                    `json:"score,omitempty"`
    TotalSwitchCount     *int                                    `json:"total_switch_count,omitempty"`
    AdditionalProperties map[string]any                          `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for ResponseSwitchMetricsActivePortsSummary.
// It customizes the JSON marshaling process for ResponseSwitchMetricsActivePortsSummary objects.
func (r ResponseSwitchMetricsActivePortsSummary) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(r.toMap())
}

// toMap converts the ResponseSwitchMetricsActivePortsSummary object to a map representation for JSON marshaling.
func (r ResponseSwitchMetricsActivePortsSummary) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, r.AdditionalProperties)
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

// UnmarshalJSON implements the json.Unmarshaler interface for ResponseSwitchMetricsActivePortsSummary.
// It customizes the JSON unmarshaling process for ResponseSwitchMetricsActivePortsSummary objects.
func (r *ResponseSwitchMetricsActivePortsSummary) UnmarshalJSON(input []byte) error {
    var temp responseSwitchMetricsActivePortsSummary
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "details", "score", "total_switch_count")
    if err != nil {
    	return err
    }
    
    r.AdditionalProperties = additionalProperties
    r.Details = temp.Details
    r.Score = temp.Score
    r.TotalSwitchCount = temp.TotalSwitchCount
    return nil
}

// responseSwitchMetricsActivePortsSummary is a temporary struct used for validating the fields of ResponseSwitchMetricsActivePortsSummary.
type responseSwitchMetricsActivePortsSummary  struct {
    Details          *SwitchMetricsActivePortsSummaryDetails `json:"details,omitempty"`
    Score            *int                                    `json:"score,omitempty"`
    TotalSwitchCount *int                                    `json:"total_switch_count,omitempty"`
}