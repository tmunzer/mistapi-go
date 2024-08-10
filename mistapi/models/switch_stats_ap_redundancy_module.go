package models

import (
    "encoding/json"
)

// SwitchStatsApRedundancyModule represents a SwitchStatsApRedundancyModule struct.
type SwitchStatsApRedundancyModule struct {
    NumAps                     *int           `json:"num_aps,omitempty"`
    NumApsWithSwitchRedundancy *int           `json:"num_aps_with_switch_redundancy,omitempty"`
    AdditionalProperties       map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for SwitchStatsApRedundancyModule.
// It customizes the JSON marshaling process for SwitchStatsApRedundancyModule objects.
func (s SwitchStatsApRedundancyModule) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(s.toMap())
}

// toMap converts the SwitchStatsApRedundancyModule object to a map representation for JSON marshaling.
func (s SwitchStatsApRedundancyModule) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, s.AdditionalProperties)
    if s.NumAps != nil {
        structMap["num_aps"] = s.NumAps
    }
    if s.NumApsWithSwitchRedundancy != nil {
        structMap["num_aps_with_switch_redundancy"] = s.NumApsWithSwitchRedundancy
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SwitchStatsApRedundancyModule.
// It customizes the JSON unmarshaling process for SwitchStatsApRedundancyModule objects.
func (s *SwitchStatsApRedundancyModule) UnmarshalJSON(input []byte) error {
    var temp tempSwitchStatsApRedundancyModule
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "num_aps", "num_aps_with_switch_redundancy")
    if err != nil {
    	return err
    }
    
    s.AdditionalProperties = additionalProperties
    s.NumAps = temp.NumAps
    s.NumApsWithSwitchRedundancy = temp.NumApsWithSwitchRedundancy
    return nil
}

// tempSwitchStatsApRedundancyModule is a temporary struct used for validating the fields of SwitchStatsApRedundancyModule.
type tempSwitchStatsApRedundancyModule  struct {
    NumAps                     *int `json:"num_aps,omitempty"`
    NumApsWithSwitchRedundancy *int `json:"num_aps_with_switch_redundancy,omitempty"`
}
