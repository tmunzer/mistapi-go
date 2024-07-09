package models

import (
    "encoding/json"
)

// ResponseSwitchMetricsConfigSuccess represents a ResponseSwitchMetricsConfigSuccess struct.
type ResponseSwitchMetricsConfigSuccess struct {
    Details              *ResponseSwitchMetricsConfigSuccessDetails `json:"details,omitempty"`
    Score                *int                                       `json:"score,omitempty"`
    TotalSwitchCount     *int                                       `json:"total_switch_count,omitempty"`
    AdditionalProperties map[string]any                             `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for ResponseSwitchMetricsConfigSuccess.
// It customizes the JSON marshaling process for ResponseSwitchMetricsConfigSuccess objects.
func (r ResponseSwitchMetricsConfigSuccess) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(r.toMap())
}

// toMap converts the ResponseSwitchMetricsConfigSuccess object to a map representation for JSON marshaling.
func (r ResponseSwitchMetricsConfigSuccess) toMap() map[string]any {
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

// UnmarshalJSON implements the json.Unmarshaler interface for ResponseSwitchMetricsConfigSuccess.
// It customizes the JSON unmarshaling process for ResponseSwitchMetricsConfigSuccess objects.
func (r *ResponseSwitchMetricsConfigSuccess) UnmarshalJSON(input []byte) error {
    var temp responseSwitchMetricsConfigSuccess
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

// responseSwitchMetricsConfigSuccess is a temporary struct used for validating the fields of ResponseSwitchMetricsConfigSuccess.
type responseSwitchMetricsConfigSuccess  struct {
    Details          *ResponseSwitchMetricsConfigSuccessDetails `json:"details,omitempty"`
    Score            *int                                       `json:"score,omitempty"`
    TotalSwitchCount *int                                       `json:"total_switch_count,omitempty"`
}
