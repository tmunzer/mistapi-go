package models

import (
    "encoding/json"
)

// AutoPlacement represents a AutoPlacement struct.
type AutoPlacement struct {
    // * If `force_collection`==`false`: the API Iattempts to start localization with existing data. 
    // * If `force_collection`==`true`: maintenance the API attempts to start orchestration.
    ForceCollection      *bool          `json:"force_collection,omitempty"`
    // list of device macs
    Macs                 []string       `json:"macs,omitempty"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for AutoPlacement.
// It customizes the JSON marshaling process for AutoPlacement objects.
func (a AutoPlacement) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(a.toMap())
}

// toMap converts the AutoPlacement object to a map representation for JSON marshaling.
func (a AutoPlacement) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, a.AdditionalProperties)
    if a.ForceCollection != nil {
        structMap["force_collection"] = a.ForceCollection
    }
    if a.Macs != nil {
        structMap["macs"] = a.Macs
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for AutoPlacement.
// It customizes the JSON unmarshaling process for AutoPlacement objects.
func (a *AutoPlacement) UnmarshalJSON(input []byte) error {
    var temp autoPlacement
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "force_collection", "macs")
    if err != nil {
    	return err
    }
    
    a.AdditionalProperties = additionalProperties
    a.ForceCollection = temp.ForceCollection
    a.Macs = temp.Macs
    return nil
}

// autoPlacement is a temporary struct used for validating the fields of AutoPlacement.
type autoPlacement  struct {
    ForceCollection *bool    `json:"force_collection,omitempty"`
    Macs            []string `json:"macs,omitempty"`
}
