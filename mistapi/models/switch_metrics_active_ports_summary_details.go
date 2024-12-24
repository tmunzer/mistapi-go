package models

import (
    "encoding/json"
    "fmt"
)

// SwitchMetricsActivePortsSummaryDetails represents a SwitchMetricsActivePortsSummaryDetails struct.
type SwitchMetricsActivePortsSummaryDetails struct {
    ActivePortCount      *int                   `json:"active_port_count,omitempty"`
    TotalPortCount       *int                   `json:"total_port_count,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for SwitchMetricsActivePortsSummaryDetails,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s SwitchMetricsActivePortsSummaryDetails) String() string {
    return fmt.Sprintf(
    	"SwitchMetricsActivePortsSummaryDetails[ActivePortCount=%v, TotalPortCount=%v, AdditionalProperties=%v]",
    	s.ActivePortCount, s.TotalPortCount, s.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for SwitchMetricsActivePortsSummaryDetails.
// It customizes the JSON marshaling process for SwitchMetricsActivePortsSummaryDetails objects.
func (s SwitchMetricsActivePortsSummaryDetails) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(s.AdditionalProperties,
        "active_port_count", "total_port_count"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(s.toMap())
}

// toMap converts the SwitchMetricsActivePortsSummaryDetails object to a map representation for JSON marshaling.
func (s SwitchMetricsActivePortsSummaryDetails) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, s.AdditionalProperties)
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
    var temp tempSwitchMetricsActivePortsSummaryDetails
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "active_port_count", "total_port_count")
    if err != nil {
    	return err
    }
    s.AdditionalProperties = additionalProperties
    
    s.ActivePortCount = temp.ActivePortCount
    s.TotalPortCount = temp.TotalPortCount
    return nil
}

// tempSwitchMetricsActivePortsSummaryDetails is a temporary struct used for validating the fields of SwitchMetricsActivePortsSummaryDetails.
type tempSwitchMetricsActivePortsSummaryDetails  struct {
    ActivePortCount *int `json:"active_port_count,omitempty"`
    TotalPortCount  *int `json:"total_port_count,omitempty"`
}
