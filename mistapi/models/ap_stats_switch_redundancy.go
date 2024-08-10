package models

import (
    "encoding/json"
)

// ApStatsSwitchRedundancy represents a ApStatsSwitchRedundancy struct.
type ApStatsSwitchRedundancy struct {
    NumRedundantAps      Optional[int]  `json:"num_redundant_aps"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for ApStatsSwitchRedundancy.
// It customizes the JSON marshaling process for ApStatsSwitchRedundancy objects.
func (a ApStatsSwitchRedundancy) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(a.toMap())
}

// toMap converts the ApStatsSwitchRedundancy object to a map representation for JSON marshaling.
func (a ApStatsSwitchRedundancy) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, a.AdditionalProperties)
    if a.NumRedundantAps.IsValueSet() {
        if a.NumRedundantAps.Value() != nil {
            structMap["num_redundant_aps"] = a.NumRedundantAps.Value()
        } else {
            structMap["num_redundant_aps"] = nil
        }
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ApStatsSwitchRedundancy.
// It customizes the JSON unmarshaling process for ApStatsSwitchRedundancy objects.
func (a *ApStatsSwitchRedundancy) UnmarshalJSON(input []byte) error {
    var temp tempApStatsSwitchRedundancy
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "num_redundant_aps")
    if err != nil {
    	return err
    }
    
    a.AdditionalProperties = additionalProperties
    a.NumRedundantAps = temp.NumRedundantAps
    return nil
}

// tempApStatsSwitchRedundancy is a temporary struct used for validating the fields of ApStatsSwitchRedundancy.
type tempApStatsSwitchRedundancy  struct {
    NumRedundantAps Optional[int] `json:"num_redundant_aps"`
}
