package models

import (
    "encoding/json"
)

// StatsSwitchApRedundancy represents a StatsSwitchApRedundancy struct.
type StatsSwitchApRedundancy struct {
    // for a VC / stacked switches.
    Modules                    map[string]StatsSwitchApRedundancyModule `json:"modules,omitempty"`
    NumAps                     *int                                     `json:"num_aps,omitempty"`
    NumApsWithSwitchRedundancy *int                                     `json:"num_aps_with_switch_redundancy,omitempty"`
    AdditionalProperties       map[string]any                           `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for StatsSwitchApRedundancy.
// It customizes the JSON marshaling process for StatsSwitchApRedundancy objects.
func (s StatsSwitchApRedundancy) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(s.toMap())
}

// toMap converts the StatsSwitchApRedundancy object to a map representation for JSON marshaling.
func (s StatsSwitchApRedundancy) toMap() map[string]any {
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

// UnmarshalJSON implements the json.Unmarshaler interface for StatsSwitchApRedundancy.
// It customizes the JSON unmarshaling process for StatsSwitchApRedundancy objects.
func (s *StatsSwitchApRedundancy) UnmarshalJSON(input []byte) error {
    var temp tempStatsSwitchApRedundancy
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

// tempStatsSwitchApRedundancy is a temporary struct used for validating the fields of StatsSwitchApRedundancy.
type tempStatsSwitchApRedundancy  struct {
    Modules                    map[string]StatsSwitchApRedundancyModule `json:"modules,omitempty"`
    NumAps                     *int                                     `json:"num_aps,omitempty"`
    NumApsWithSwitchRedundancy *int                                     `json:"num_aps_with_switch_redundancy,omitempty"`
}