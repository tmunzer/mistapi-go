package models

import (
    "encoding/json"
)

// SwitchStatsApRedundancy represents a SwitchStatsApRedundancy struct.
type SwitchStatsApRedundancy struct {
    // for a VC / stacked switches.
    Modules                    map[string]SwitchStatsApRedundancyModule `json:"modules,omitempty"`
    NumAps                     *int                                     `json:"num_aps,omitempty"`
    NumApsWithSwitchRedundancy *int                                     `json:"num_aps_with_switch_redundancy,omitempty"`
    AdditionalProperties       map[string]any                           `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for SwitchStatsApRedundancy.
// It customizes the JSON marshaling process for SwitchStatsApRedundancy objects.
func (s SwitchStatsApRedundancy) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(s.toMap())
}

// toMap converts the SwitchStatsApRedundancy object to a map representation for JSON marshaling.
func (s SwitchStatsApRedundancy) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, s.AdditionalProperties)
    if s.Modules != nil {
        structMap["modules"] = s.Modules
    }
    if s.NumAps != nil {
        structMap["num_aps"] = s.NumAps
    }
    if s.NumApsWithSwitchRedundancy != nil {
        structMap["num_aps_with_switch_redundancy"] = s.NumApsWithSwitchRedundancy
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SwitchStatsApRedundancy.
// It customizes the JSON unmarshaling process for SwitchStatsApRedundancy objects.
func (s *SwitchStatsApRedundancy) UnmarshalJSON(input []byte) error {
    var temp tempSwitchStatsApRedundancy
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "modules", "num_aps", "num_aps_with_switch_redundancy")
    if err != nil {
    	return err
    }
    
    s.AdditionalProperties = additionalProperties
    s.Modules = temp.Modules
    s.NumAps = temp.NumAps
    s.NumApsWithSwitchRedundancy = temp.NumApsWithSwitchRedundancy
    return nil
}

// tempSwitchStatsApRedundancy is a temporary struct used for validating the fields of SwitchStatsApRedundancy.
type tempSwitchStatsApRedundancy  struct {
    Modules                    map[string]SwitchStatsApRedundancyModule `json:"modules,omitempty"`
    NumAps                     *int                                     `json:"num_aps,omitempty"`
    NumApsWithSwitchRedundancy *int                                     `json:"num_aps_with_switch_redundancy,omitempty"`
}
