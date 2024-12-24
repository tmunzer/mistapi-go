package models

import (
    "encoding/json"
    "fmt"
)

// StatsWxruleUsageProperties represents a StatsWxruleUsageProperties struct.
type StatsWxruleUsageProperties struct {
    NumFlows             *int                   `json:"num_flows,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for StatsWxruleUsageProperties,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s StatsWxruleUsageProperties) String() string {
    return fmt.Sprintf(
    	"StatsWxruleUsageProperties[NumFlows=%v, AdditionalProperties=%v]",
    	s.NumFlows, s.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for StatsWxruleUsageProperties.
// It customizes the JSON marshaling process for StatsWxruleUsageProperties objects.
func (s StatsWxruleUsageProperties) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(s.AdditionalProperties,
        "num_flows"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(s.toMap())
}

// toMap converts the StatsWxruleUsageProperties object to a map representation for JSON marshaling.
func (s StatsWxruleUsageProperties) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, s.AdditionalProperties)
    if s.NumFlows != nil {
        structMap["num_flows"] = s.NumFlows
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for StatsWxruleUsageProperties.
// It customizes the JSON unmarshaling process for StatsWxruleUsageProperties objects.
func (s *StatsWxruleUsageProperties) UnmarshalJSON(input []byte) error {
    var temp tempStatsWxruleUsageProperties
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "num_flows")
    if err != nil {
    	return err
    }
    s.AdditionalProperties = additionalProperties
    
    s.NumFlows = temp.NumFlows
    return nil
}

// tempStatsWxruleUsageProperties is a temporary struct used for validating the fields of StatsWxruleUsageProperties.
type tempStatsWxruleUsageProperties  struct {
    NumFlows *int `json:"num_flows,omitempty"`
}
