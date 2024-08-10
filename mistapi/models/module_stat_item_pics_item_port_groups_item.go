package models

import (
    "encoding/json"
)

// ModuleStatItemPicsItemPortGroupsItem represents a ModuleStatItemPicsItemPortGroupsItem struct.
type ModuleStatItemPicsItemPortGroupsItem struct {
    Count                *int           `json:"count,omitempty"`
    Type                 *string        `json:"type,omitempty"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for ModuleStatItemPicsItemPortGroupsItem.
// It customizes the JSON marshaling process for ModuleStatItemPicsItemPortGroupsItem objects.
func (m ModuleStatItemPicsItemPortGroupsItem) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(m.toMap())
}

// toMap converts the ModuleStatItemPicsItemPortGroupsItem object to a map representation for JSON marshaling.
func (m ModuleStatItemPicsItemPortGroupsItem) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, m.AdditionalProperties)
    if m.Count != nil {
        structMap["count"] = m.Count
    }
    if m.Type != nil {
        structMap["type"] = m.Type
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ModuleStatItemPicsItemPortGroupsItem.
// It customizes the JSON unmarshaling process for ModuleStatItemPicsItemPortGroupsItem objects.
func (m *ModuleStatItemPicsItemPortGroupsItem) UnmarshalJSON(input []byte) error {
    var temp tempModuleStatItemPicsItemPortGroupsItem
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "count", "type")
    if err != nil {
    	return err
    }
    
    m.AdditionalProperties = additionalProperties
    m.Count = temp.Count
    m.Type = temp.Type
    return nil
}

// tempModuleStatItemPicsItemPortGroupsItem is a temporary struct used for validating the fields of ModuleStatItemPicsItemPortGroupsItem.
type tempModuleStatItemPicsItemPortGroupsItem  struct {
    Count *int    `json:"count,omitempty"`
    Type  *string `json:"type,omitempty"`
}
