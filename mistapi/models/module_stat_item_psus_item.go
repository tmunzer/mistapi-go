package models

import (
    "encoding/json"
)

// ModuleStatItemPsusItem represents a ModuleStatItemPsusItem struct.
type ModuleStatItemPsusItem struct {
    Name                 *string                `json:"name,omitempty"`
    Status               *string                `json:"status,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for ModuleStatItemPsusItem.
// It customizes the JSON marshaling process for ModuleStatItemPsusItem objects.
func (m ModuleStatItemPsusItem) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(m.AdditionalProperties,
        "name", "status"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(m.toMap())
}

// toMap converts the ModuleStatItemPsusItem object to a map representation for JSON marshaling.
func (m ModuleStatItemPsusItem) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, m.AdditionalProperties)
    if m.Name != nil {
        structMap["name"] = m.Name
    }
    if m.Status != nil {
        structMap["status"] = m.Status
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ModuleStatItemPsusItem.
// It customizes the JSON unmarshaling process for ModuleStatItemPsusItem objects.
func (m *ModuleStatItemPsusItem) UnmarshalJSON(input []byte) error {
    var temp tempModuleStatItemPsusItem
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "name", "status")
    if err != nil {
    	return err
    }
    m.AdditionalProperties = additionalProperties
    
    m.Name = temp.Name
    m.Status = temp.Status
    return nil
}

// tempModuleStatItemPsusItem is a temporary struct used for validating the fields of ModuleStatItemPsusItem.
type tempModuleStatItemPsusItem  struct {
    Name   *string `json:"name,omitempty"`
    Status *string `json:"status,omitempty"`
}
