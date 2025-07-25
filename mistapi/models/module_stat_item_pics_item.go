// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
    "encoding/json"
    "fmt"
)

// ModuleStatItemPicsItem represents a ModuleStatItemPicsItem struct.
type ModuleStatItemPicsItem struct {
    Index                *int                                   `json:"index,omitempty"`
    ModelNumber          *string                                `json:"model_number,omitempty"`
    PortGroups           []ModuleStatItemPicsItemPortGroupsItem `json:"port_groups,omitempty"`
    AdditionalProperties map[string]interface{}                 `json:"_"`
}

// String implements the fmt.Stringer interface for ModuleStatItemPicsItem,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (m ModuleStatItemPicsItem) String() string {
    return fmt.Sprintf(
    	"ModuleStatItemPicsItem[Index=%v, ModelNumber=%v, PortGroups=%v, AdditionalProperties=%v]",
    	m.Index, m.ModelNumber, m.PortGroups, m.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for ModuleStatItemPicsItem.
// It customizes the JSON marshaling process for ModuleStatItemPicsItem objects.
func (m ModuleStatItemPicsItem) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(m.AdditionalProperties,
        "index", "model_number", "port_groups"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(m.toMap())
}

// toMap converts the ModuleStatItemPicsItem object to a map representation for JSON marshaling.
func (m ModuleStatItemPicsItem) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, m.AdditionalProperties)
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
    var temp tempModuleStatItemPicsItem
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "index", "model_number", "port_groups")
    if err != nil {
    	return err
    }
    m.AdditionalProperties = additionalProperties
    
    m.Index = temp.Index
    m.ModelNumber = temp.ModelNumber
    m.PortGroups = temp.PortGroups
    return nil
}

// tempModuleStatItemPicsItem is a temporary struct used for validating the fields of ModuleStatItemPicsItem.
type tempModuleStatItemPicsItem  struct {
    Index       *int                                   `json:"index,omitempty"`
    ModelNumber *string                                `json:"model_number,omitempty"`
    PortGroups  []ModuleStatItemPicsItemPortGroupsItem `json:"port_groups,omitempty"`
}
