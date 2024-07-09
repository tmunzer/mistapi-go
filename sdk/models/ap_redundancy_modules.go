package models

import (
    "encoding/json"
)

// ApRedundancyModules represents a ApRedundancyModules struct.
// for a VC / stacked switches.
type ApRedundancyModules struct {
    NumAps                     *int           `json:"num_aps,omitempty"`
    NumApsWithSwitchRedundancy *int           `json:"num_aps_with_switch_redundancy,omitempty"`
    AdditionalProperties       map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for ApRedundancyModules.
// It customizes the JSON marshaling process for ApRedundancyModules objects.
func (a ApRedundancyModules) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(a.toMap())
}

// toMap converts the ApRedundancyModules object to a map representation for JSON marshaling.
func (a ApRedundancyModules) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, a.AdditionalProperties)
    if a.NumAps != nil {
        structMap["num_aps"] = a.NumAps
    }
    if a.NumApsWithSwitchRedundancy != nil {
        structMap["num_aps_with_switch_redundancy"] = a.NumApsWithSwitchRedundancy
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ApRedundancyModules.
// It customizes the JSON unmarshaling process for ApRedundancyModules objects.
func (a *ApRedundancyModules) UnmarshalJSON(input []byte) error {
    var temp apRedundancyModules
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "num_aps", "num_aps_with_switch_redundancy")
    if err != nil {
    	return err
    }
    
    a.AdditionalProperties = additionalProperties
    a.NumAps = temp.NumAps
    a.NumApsWithSwitchRedundancy = temp.NumApsWithSwitchRedundancy
    return nil
}

// apRedundancyModules is a temporary struct used for validating the fields of ApRedundancyModules.
type apRedundancyModules  struct {
    NumAps                     *int `json:"num_aps,omitempty"`
    NumApsWithSwitchRedundancy *int `json:"num_aps_with_switch_redundancy,omitempty"`
}
