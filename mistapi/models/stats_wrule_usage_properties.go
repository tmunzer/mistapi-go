package models

import (
    "encoding/json"
)

// StatsWruleUsageProperties represents a StatsWruleUsageProperties struct.
type StatsWruleUsageProperties struct {
    NumFlows             *int           `json:"num_flows,omitempty"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for StatsWruleUsageProperties.
// It customizes the JSON marshaling process for StatsWruleUsageProperties objects.
func (s StatsWruleUsageProperties) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(s.toMap())
}

// toMap converts the StatsWruleUsageProperties object to a map representation for JSON marshaling.
func (s StatsWruleUsageProperties) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, s.AdditionalProperties)
    if s.NumFlows != nil {
        structMap["num_flows"] = s.NumFlows
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for StatsWruleUsageProperties.
// It customizes the JSON unmarshaling process for StatsWruleUsageProperties objects.
func (s *StatsWruleUsageProperties) UnmarshalJSON(input []byte) error {
    var temp tempStatsWruleUsageProperties
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

// tempStatsWruleUsageProperties is a temporary struct used for validating the fields of StatsWruleUsageProperties.
type tempStatsWruleUsageProperties  struct {
    NumFlows *int `json:"num_flows,omitempty"`
}
