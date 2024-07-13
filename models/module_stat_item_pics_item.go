package models

import (
    "encoding/json"
)

// ModuleStatItemPicsItem represents a ModuleStatItemPicsItem struct.
type ModuleStatItemPicsItem struct {
    Index                *int                                   `json:"index,omitempty"`
    ModelNumber          *string                                `json:"model_number,omitempty"`
    PortGroups           []ModuleStatItemPicsItemPortGroupsItem `json:"port_groups,omitempty"`
    AdditionalProperties map[string]any                         `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for ModuleStatItemPicsItem.
// It customizes the JSON marshaling process for ModuleStatItemPicsItem objects.
func (m ModuleStatItemPicsItem) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(m.toMap())
}

// toMap converts the ModuleStatItemPicsItem object to a map representation for JSON marshaling.
func (m ModuleStatItemPicsItem) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, m.AdditionalProperties)
    if m.Index != nil {
        structMap["index"] = m.Index
    }
    if m.ModelNumber != nil {
        structMap["model_number"] = m.ModelNumber
    }
    if m.PortGroups != nil {
        structMap["port_groups"] = m.PortGroups
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ModuleStatItemPicsItem.
// It customizes the JSON unmarshaling process for ModuleStatItemPicsItem objects.
func (m *ModuleStatItemPicsItem) UnmarshalJSON(input []byte) error {
    var temp moduleStatItemPicsItem
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "index", "model_number", "port_groups")
    if err != nil {
    	return err
    }
    
    m.AdditionalProperties = additionalProperties
    m.Index = temp.Index
    m.ModelNumber = temp.ModelNumber
    m.PortGroups = temp.PortGroups
    return nil
}

// moduleStatItemPicsItem is a temporary struct used for validating the fields of ModuleStatItemPicsItem.
type moduleStatItemPicsItem  struct {
    Index       *int                                   `json:"index,omitempty"`
    ModelNumber *string                                `json:"model_number,omitempty"`
    PortGroups  []ModuleStatItemPicsItemPortGroupsItem `json:"port_groups,omitempty"`
}
