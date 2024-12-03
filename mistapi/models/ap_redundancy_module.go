package models

import (
    "encoding/json"
)

// ApRedundancyModule represents a ApRedundancyModule struct.
type ApRedundancyModule struct {
    NumAps                     *int                   `json:"num_aps,omitempty"`
    NumApsWithSwitchRedundancy *int                   `json:"num_aps_with_switch_redundancy,omitempty"`
    AdditionalProperties       map[string]interface{} `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for ApRedundancyModule.
// It customizes the JSON marshaling process for ApRedundancyModule objects.
func (a ApRedundancyModule) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(a.AdditionalProperties,
        "num_aps", "num_aps_with_switch_redundancy"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(a.toMap())
}

// toMap converts the ApRedundancyModule object to a map representation for JSON marshaling.
func (a ApRedundancyModule) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, a.AdditionalProperties)
    if a.NumAps != nil {
        structMap["num_aps"] = a.NumAps
    }
    if a.NumApsWithSwitchRedundancy != nil {
        structMap["num_aps_with_switch_redundancy"] = a.NumApsWithSwitchRedundancy
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ApRedundancyModule.
// It customizes the JSON unmarshaling process for ApRedundancyModule objects.
func (a *ApRedundancyModule) UnmarshalJSON(input []byte) error {
    var temp tempApRedundancyModule
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "num_aps", "num_aps_with_switch_redundancy")
    if err != nil {
    	return err
    }
    a.AdditionalProperties = additionalProperties
    
    a.NumAps = temp.NumAps
    a.NumApsWithSwitchRedundancy = temp.NumApsWithSwitchRedundancy
    return nil
}

// tempApRedundancyModule is a temporary struct used for validating the fields of ApRedundancyModule.
type tempApRedundancyModule  struct {
    NumAps                     *int `json:"num_aps,omitempty"`
    NumApsWithSwitchRedundancy *int `json:"num_aps_with_switch_redundancy,omitempty"`
}
