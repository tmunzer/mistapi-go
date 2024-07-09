package models

import (
    "encoding/json"
)

// AutoOrient represents a AutoOrient struct.
type AutoOrient struct {
    // If `force_collection`==`false`, the API attempts to start auto orientation with existing BLE data. 
    // If `force_collection`==`true`, the API attempts to start BLE orchestration.
    ForceCollection      *bool          `json:"force_collection,omitempty"`
    // list of device macs
    Macs                 []string       `json:"macs,omitempty"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for AutoOrient.
// It customizes the JSON marshaling process for AutoOrient objects.
func (a AutoOrient) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(a.toMap())
}

// toMap converts the AutoOrient object to a map representation for JSON marshaling.
func (a AutoOrient) toMap() map[string]any {
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

// UnmarshalJSON implements the json.Unmarshaler interface for AutoOrient.
// It customizes the JSON unmarshaling process for AutoOrient objects.
func (a *AutoOrient) UnmarshalJSON(input []byte) error {
    var temp autoOrient
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

// autoOrient is a temporary struct used for validating the fields of AutoOrient.
type autoOrient  struct {
    ForceCollection *bool    `json:"force_collection,omitempty"`
    Macs            []string `json:"macs,omitempty"`
}