package models

import (
    "encoding/json"
)

// MxedgeTuntermMulticastSsdp represents a MxedgeTuntermMulticastSsdp struct.
type MxedgeTuntermMulticastSsdp struct {
    Enabled              *bool          `json:"enabled,omitempty"`
    VlanIds              []string       `json:"vlan_ids,omitempty"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for MxedgeTuntermMulticastSsdp.
// It customizes the JSON marshaling process for MxedgeTuntermMulticastSsdp objects.
func (m MxedgeTuntermMulticastSsdp) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(m.toMap())
}

// toMap converts the MxedgeTuntermMulticastSsdp object to a map representation for JSON marshaling.
func (m MxedgeTuntermMulticastSsdp) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, m.AdditionalProperties)
    if m.Enabled != nil {
        structMap["enabled"] = m.Enabled
    }
    if m.VlanIds != nil {
        structMap["vlan_ids"] = m.VlanIds
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for MxedgeTuntermMulticastSsdp.
// It customizes the JSON unmarshaling process for MxedgeTuntermMulticastSsdp objects.
func (m *MxedgeTuntermMulticastSsdp) UnmarshalJSON(input []byte) error {
    var temp tempMxedgeTuntermMulticastSsdp
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "enabled", "vlan_ids")
    if err != nil {
    	return err
    }
    
    m.AdditionalProperties = additionalProperties
    m.Enabled = temp.Enabled
    m.VlanIds = temp.VlanIds
    return nil
}

// tempMxedgeTuntermMulticastSsdp is a temporary struct used for validating the fields of MxedgeTuntermMulticastSsdp.
type tempMxedgeTuntermMulticastSsdp  struct {
    Enabled *bool    `json:"enabled,omitempty"`
    VlanIds []string `json:"vlan_ids,omitempty"`
}
