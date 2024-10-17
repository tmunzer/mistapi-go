package models

import (
    "encoding/json"
)

// StatsWxruleUsageProperties represents a StatsWxruleUsageProperties struct.
type StatsWxruleUsageProperties struct {
    NumFlows             *int           `json:"num_flows,omitempty"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for StatsWxruleUsageProperties.
// It customizes the JSON marshaling process for StatsWxruleUsageProperties objects.
func (s StatsWxruleUsageProperties) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(s.toMap())
}

// toMap converts the StatsWxruleUsageProperties object to a map representation for JSON marshaling.
func (s StatsWxruleUsageProperties) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, s.AdditionalProperties)
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
    additionalProperties, err := UnmarshalAdditionalProperties(input, "num_flows")
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
