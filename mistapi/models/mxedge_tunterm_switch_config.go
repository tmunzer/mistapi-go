// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
    "encoding/json"
    "fmt"
)

// MxedgeTuntermSwitchConfig represents a MxedgeTuntermSwitchConfig struct.
type MxedgeTuntermSwitchConfig struct {
    PortVlanId           *int                   `json:"port_vlan_id,omitempty"`
    VlanIds              []VlanIdWithVariable   `json:"vlan_ids,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for MxedgeTuntermSwitchConfig,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (m MxedgeTuntermSwitchConfig) String() string {
    return fmt.Sprintf(
    	"MxedgeTuntermSwitchConfig[PortVlanId=%v, VlanIds=%v, AdditionalProperties=%v]",
    	m.PortVlanId, m.VlanIds, m.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for MxedgeTuntermSwitchConfig.
// It customizes the JSON marshaling process for MxedgeTuntermSwitchConfig objects.
func (m MxedgeTuntermSwitchConfig) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(m.AdditionalProperties,
        "port_vlan_id", "vlan_ids"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(m.toMap())
}

// toMap converts the MxedgeTuntermSwitchConfig object to a map representation for JSON marshaling.
func (m MxedgeTuntermSwitchConfig) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, m.AdditionalProperties)
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
    var temp tempMxedgeTuntermSwitchConfig
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "port_vlan_id", "vlan_ids")
    if err != nil {
    	return err
    }
    m.AdditionalProperties = additionalProperties
    
    m.PortVlanId = temp.PortVlanId
    m.VlanIds = temp.VlanIds
    return nil
}

// tempMxedgeTuntermSwitchConfig is a temporary struct used for validating the fields of MxedgeTuntermSwitchConfig.
type tempMxedgeTuntermSwitchConfig  struct {
    PortVlanId *int                 `json:"port_vlan_id,omitempty"`
    VlanIds    []VlanIdWithVariable `json:"vlan_ids,omitempty"`
}
