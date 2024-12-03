package models

import (
    "encoding/json"
)

// ModuleStatItemVcLinksItem represents a ModuleStatItemVcLinksItem struct.
type ModuleStatItemVcLinksItem struct {
    NeighborModuleIdx    *int                   `json:"neighbor_module_idx,omitempty"`
    NeighborPortId       *string                `json:"neighbor_port_id,omitempty"`
    PortId               *string                `json:"port_id,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for ModuleStatItemVcLinksItem.
// It customizes the JSON marshaling process for ModuleStatItemVcLinksItem objects.
func (m ModuleStatItemVcLinksItem) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(m.AdditionalProperties,
        "neighbor_module_idx", "neighbor_port_id", "port_id"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(m.toMap())
}

// toMap converts the ModuleStatItemVcLinksItem object to a map representation for JSON marshaling.
func (m ModuleStatItemVcLinksItem) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, m.AdditionalProperties)
    if m.NeighborModuleIdx != nil {
        structMap["neighbor_module_idx"] = m.NeighborModuleIdx
    }
    if m.NeighborPortId != nil {
        structMap["neighbor_port_id"] = m.NeighborPortId
    }
    if m.PortId != nil {
        structMap["port_id"] = m.PortId
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ModuleStatItemVcLinksItem.
// It customizes the JSON unmarshaling process for ModuleStatItemVcLinksItem objects.
func (m *ModuleStatItemVcLinksItem) UnmarshalJSON(input []byte) error {
    var temp tempModuleStatItemVcLinksItem
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "neighbor_module_idx", "neighbor_port_id", "port_id")
    if err != nil {
    	return err
    }
    m.AdditionalProperties = additionalProperties
    
    m.NeighborModuleIdx = temp.NeighborModuleIdx
    m.NeighborPortId = temp.NeighborPortId
    m.PortId = temp.PortId
    return nil
}

// tempModuleStatItemVcLinksItem is a temporary struct used for validating the fields of ModuleStatItemVcLinksItem.
type tempModuleStatItemVcLinksItem  struct {
    NeighborModuleIdx *int    `json:"neighbor_module_idx,omitempty"`
    NeighborPortId    *string `json:"neighbor_port_id,omitempty"`
    PortId            *string `json:"port_id,omitempty"`
}
