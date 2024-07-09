package models

import (
    "encoding/json"
)

// NameString represents a NameString struct.
type NameString struct {
    Name                 *string        `json:"name,omitempty"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for NameString.
// It customizes the JSON marshaling process for NameString objects.
func (n NameString) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(n.toMap())
}

// toMap converts the NameString object to a map representation for JSON marshaling.
func (n NameString) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, n.AdditionalProperties)
    if n.Name != nil {
        structMap["name"] = n.Name
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for NameString.
// It customizes the JSON unmarshaling process for NameString objects.
func (n *NameString) UnmarshalJSON(input []byte) error {
    var temp nameString
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "name")
    if err != nil {
    	return err
    }
    
    n.AdditionalProperties = additionalProperties
    n.Name = temp.Name
    return nil
}

// nameString is a temporary struct used for validating the fields of NameString.
type nameString  struct {
    Name *string `json:"name,omitempty"`
}
