package models

import (
    "encoding/json"
)

// InventorySearchResultMember represents a InventorySearchResultMember struct.
type InventorySearchResultMember struct {
    Mac                  *string        `json:"mac,omitempty"`
    Model                *string        `json:"model,omitempty"`
    Serial               *string        `json:"serial,omitempty"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for InventorySearchResultMember.
// It customizes the JSON marshaling process for InventorySearchResultMember objects.
func (i InventorySearchResultMember) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(i.toMap())
}

// toMap converts the InventorySearchResultMember object to a map representation for JSON marshaling.
func (i InventorySearchResultMember) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, i.AdditionalProperties)
    if i.Mac != nil {
        structMap["mac"] = i.Mac
    }
    if i.Model != nil {
        structMap["model"] = i.Model
    }
    if i.Serial != nil {
        structMap["serial"] = i.Serial
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for InventorySearchResultMember.
// It customizes the JSON unmarshaling process for InventorySearchResultMember objects.
func (i *InventorySearchResultMember) UnmarshalJSON(input []byte) error {
    var temp tempInventorySearchResultMember
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "mac", "model", "serial")
    if err != nil {
    	return err
    }
    
    i.AdditionalProperties = additionalProperties
    i.Mac = temp.Mac
    i.Model = temp.Model
    i.Serial = temp.Serial
    return nil
}

// tempInventorySearchResultMember is a temporary struct used for validating the fields of InventorySearchResultMember.
type tempInventorySearchResultMember  struct {
    Mac    *string `json:"mac,omitempty"`
    Model  *string `json:"model,omitempty"`
    Serial *string `json:"serial,omitempty"`
}
