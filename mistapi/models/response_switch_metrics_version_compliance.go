package models

import (
    "encoding/json"
)

// ResponseSwitchMetricsVersionCompliance represents a ResponseSwitchMetricsVersionCompliance struct.
type ResponseSwitchMetricsVersionCompliance struct {
    Details              *ResponseSwitchMetricsVersionComplianceDetails `json:"details,omitempty"`
    Score                *int                                           `json:"score,omitempty"`
    TotalSwitchCount     *int                                           `json:"total_switch_count,omitempty"`
    AdditionalProperties map[string]any                                 `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for ResponseSwitchMetricsVersionCompliance.
// It customizes the JSON marshaling process for ResponseSwitchMetricsVersionCompliance objects.
func (r ResponseSwitchMetricsVersionCompliance) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(r.toMap())
}

// toMap converts the ResponseSwitchMetricsVersionCompliance object to a map representation for JSON marshaling.
func (r ResponseSwitchMetricsVersionCompliance) toMap() map[string]any {
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

// UnmarshalJSON implements the json.Unmarshaler interface for ResponseSwitchMetricsVersionCompliance.
// It customizes the JSON unmarshaling process for ResponseSwitchMetricsVersionCompliance objects.
func (r *ResponseSwitchMetricsVersionCompliance) UnmarshalJSON(input []byte) error {
    var temp tempResponseSwitchMetricsVersionCompliance
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

// tempResponseSwitchMetricsVersionCompliance is a temporary struct used for validating the fields of ResponseSwitchMetricsVersionCompliance.
type tempResponseSwitchMetricsVersionCompliance  struct {
    Details          *ResponseSwitchMetricsVersionComplianceDetails `json:"details,omitempty"`
    Score            *int                                           `json:"score,omitempty"`
    TotalSwitchCount *int                                           `json:"total_switch_count,omitempty"`
}
