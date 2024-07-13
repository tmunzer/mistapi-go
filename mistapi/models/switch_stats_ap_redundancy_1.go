package models

import (
    "encoding/json"
)

// SwitchStatsApRedundancy1 represents a SwitchStatsApRedundancy1 struct.
type SwitchStatsApRedundancy1 struct {
    // for a VC / stacked switches.
    Modules                    *ApRedundancyModules `json:"modules,omitempty"`
    NumAps                     *int                 `json:"num_aps,omitempty"`
    NumApsWithSwitchRedundancy *int                 `json:"num_aps_with_switch_redundancy,omitempty"`
    AdditionalProperties       map[string]any       `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for SwitchStatsApRedundancy1.
// It customizes the JSON marshaling process for SwitchStatsApRedundancy1 objects.
func (s SwitchStatsApRedundancy1) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(s.toMap())
}

// toMap converts the SwitchStatsApRedundancy1 object to a map representation for JSON marshaling.
func (s SwitchStatsApRedundancy1) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, s.AdditionalProperties)
    if s.Modules != nil {
        structMap["modules"] = s.Modules.toMap()
    }
    if s.NumAps != nil {
        structMap["num_aps"] = s.NumAps
    }
    if s.NumApsWithSwitchRedundancy != nil {
        structMap["num_aps_with_switch_redundancy"] = s.NumApsWithSwitchRedundancy
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SwitchStatsApRedundancy1.
// It customizes the JSON unmarshaling process for SwitchStatsApRedundancy1 objects.
func (s *SwitchStatsApRedundancy1) UnmarshalJSON(input []byte) error {
    var temp switchStatsApRedundancy1
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

// switchStatsApRedundancy1 is a temporary struct used for validating the fields of SwitchStatsApRedundancy1.
type switchStatsApRedundancy1  struct {
    Modules                    *ApRedundancyModules `json:"modules,omitempty"`
    NumAps                     *int                 `json:"num_aps,omitempty"`
    NumApsWithSwitchRedundancy *int                 `json:"num_aps_with_switch_redundancy,omitempty"`
}
