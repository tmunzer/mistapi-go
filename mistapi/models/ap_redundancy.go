package models

import (
    "encoding/json"
)

// ApRedundancy represents a ApRedundancy struct.
type ApRedundancy struct {
    // Property key is the node id
    Modules                    map[string]ApRedundancyModule `json:"modules,omitempty"`
    NumAps                     *int                          `json:"num_aps,omitempty"`
    NumApsWithSwitchRedundancy *int                          `json:"num_aps_with_switch_redundancy,omitempty"`
    AdditionalProperties       map[string]any                `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for ApRedundancy.
// It customizes the JSON marshaling process for ApRedundancy objects.
func (a ApRedundancy) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(a.toMap())
}

// toMap converts the ApRedundancy object to a map representation for JSON marshaling.
func (a ApRedundancy) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, a.AdditionalProperties)
    if a.Modules != nil {
        structMap["modules"] = a.Modules
    }
    if a.NumAps != nil {
        structMap["num_aps"] = a.NumAps
    }
    if a.NumApsWithSwitchRedundancy != nil {
        structMap["num_aps_with_switch_redundancy"] = a.NumApsWithSwitchRedundancy
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ApRedundancy.
// It customizes the JSON unmarshaling process for ApRedundancy objects.
func (a *ApRedundancy) UnmarshalJSON(input []byte) error {
    var temp tempApRedundancy
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "modules", "num_aps", "num_aps_with_switch_redundancy")
    if err != nil {
    	return err
    }
    
    a.AdditionalProperties = additionalProperties
    a.Modules = temp.Modules
    a.NumAps = temp.NumAps
    a.NumApsWithSwitchRedundancy = temp.NumApsWithSwitchRedundancy
    return nil
}

// tempApRedundancy is a temporary struct used for validating the fields of ApRedundancy.
type tempApRedundancy  struct {
    Modules                    map[string]ApRedundancyModule `json:"modules,omitempty"`
    NumAps                     *int                          `json:"num_aps,omitempty"`
    NumApsWithSwitchRedundancy *int                          `json:"num_aps_with_switch_redundancy,omitempty"`
}
