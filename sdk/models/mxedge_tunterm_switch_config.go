package models

import (
    "encoding/json"
)

// MxedgeTuntermSwitchConfig represents a MxedgeTuntermSwitchConfig struct.
type MxedgeTuntermSwitchConfig struct {
    PortVlanId           *int                               `json:"port_vlan_id,omitempty"`
    VlanIds              []MxedgeTuntermSwitchConfigVlanIds `json:"vlan_ids,omitempty"`
    AdditionalProperties map[string]any                     `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for MxedgeTuntermSwitchConfig.
// It customizes the JSON marshaling process for MxedgeTuntermSwitchConfig objects.
func (m MxedgeTuntermSwitchConfig) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(m.toMap())
}

// toMap converts the MxedgeTuntermSwitchConfig object to a map representation for JSON marshaling.
func (m MxedgeTuntermSwitchConfig) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, m.AdditionalProperties)
    if m.PortVlanId != nil {
        structMap["port_vlan_id"] = m.PortVlanId
    }
    if m.VlanIds != nil {
        structMap["vlan_ids"] = m.VlanIds
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for MxedgeTuntermSwitchConfig.
// It customizes the JSON unmarshaling process for MxedgeTuntermSwitchConfig objects.
func (m *MxedgeTuntermSwitchConfig) UnmarshalJSON(input []byte) error {
    var temp mxedgeTuntermSwitchConfig
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "port_vlan_id", "vlan_ids")
    if err != nil {
    	return err
    }
    
    m.AdditionalProperties = additionalProperties
    m.PortVlanId = temp.PortVlanId
    m.VlanIds = temp.VlanIds
    return nil
}

// mxedgeTuntermSwitchConfig is a temporary struct used for validating the fields of MxedgeTuntermSwitchConfig.
type mxedgeTuntermSwitchConfig  struct {
    PortVlanId *int                               `json:"port_vlan_id,omitempty"`
    VlanIds    []MxedgeTuntermSwitchConfigVlanIds `json:"vlan_ids,omitempty"`
}
