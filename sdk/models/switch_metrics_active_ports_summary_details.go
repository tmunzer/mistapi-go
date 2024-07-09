package models

import (
    "encoding/json"
)

// SwitchMetricsActivePortsSummaryDetails represents a SwitchMetricsActivePortsSummaryDetails struct.
type SwitchMetricsActivePortsSummaryDetails struct {
    ActivePortCount      *int           `json:"active_port_count,omitempty"`
    TotalPortCount       *int           `json:"total_port_count,omitempty"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for SwitchMetricsActivePortsSummaryDetails.
// It customizes the JSON marshaling process for SwitchMetricsActivePortsSummaryDetails objects.
func (s SwitchMetricsActivePortsSummaryDetails) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(s.toMap())
}

// toMap converts the SwitchMetricsActivePortsSummaryDetails object to a map representation for JSON marshaling.
func (s SwitchMetricsActivePortsSummaryDetails) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, s.AdditionalProperties)
    if s.ActivePortCount != nil {
        structMap["active_port_count"] = s.ActivePortCount
    }
    if s.TotalPortCount != nil {
        structMap["total_port_count"] = s.TotalPortCount
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SwitchMetricsActivePortsSummaryDetails.
// It customizes the JSON unmarshaling process for SwitchMetricsActivePortsSummaryDetails objects.
func (s *SwitchMetricsActivePortsSummaryDetails) UnmarshalJSON(input []byte) error {
    var temp switchMetricsActivePortsSummaryDetails
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "active_port_count", "total_port_count")
    if err != nil {
    	return err
    }
    
    s.AdditionalProperties = additionalProperties
    s.ActivePortCount = temp.ActivePortCount
    s.TotalPortCount = temp.TotalPortCount
    return nil
}

// switchMetricsActivePortsSummaryDetails is a temporary struct used for validating the fields of SwitchMetricsActivePortsSummaryDetails.
type switchMetricsActivePortsSummaryDetails  struct {
    ActivePortCount *int `json:"active_port_count,omitempty"`
    TotalPortCount  *int `json:"total_port_count,omitempty"`
}
