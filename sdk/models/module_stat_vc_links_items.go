package models

import (
    "encoding/json"
)

// ModuleStatVcLinksItems represents a ModuleStatVcLinksItems struct.
type ModuleStatVcLinksItems struct {
    NeighborModuleIdx    *int           `json:"neighbor_module_idx,omitempty"`
    NeighborPortId       *string        `json:"neighbor_port_id,omitempty"`
    PortId               *string        `json:"port_id,omitempty"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for ModuleStatVcLinksItems.
// It customizes the JSON marshaling process for ModuleStatVcLinksItems objects.
func (m ModuleStatVcLinksItems) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(m.toMap())
}

// toMap converts the ModuleStatVcLinksItems object to a map representation for JSON marshaling.
func (m ModuleStatVcLinksItems) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, m.AdditionalProperties)
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

// UnmarshalJSON implements the json.Unmarshaler interface for ModuleStatVcLinksItems.
// It customizes the JSON unmarshaling process for ModuleStatVcLinksItems objects.
func (m *ModuleStatVcLinksItems) UnmarshalJSON(input []byte) error {
    var temp moduleStatVcLinksItems
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "neighbor_module_idx", "neighbor_port_id", "port_id")
    if err != nil {
    	return err
    }
    
    m.AdditionalProperties = additionalProperties
    m.NeighborModuleIdx = temp.NeighborModuleIdx
    m.NeighborPortId = temp.NeighborPortId
    m.PortId = temp.PortId
    return nil
}

// moduleStatVcLinksItems is a temporary struct used for validating the fields of ModuleStatVcLinksItems.
type moduleStatVcLinksItems  struct {
    NeighborModuleIdx *int    `json:"neighbor_module_idx,omitempty"`
    NeighborPortId    *string `json:"neighbor_port_id,omitempty"`
    PortId            *string `json:"port_id,omitempty"`
}
