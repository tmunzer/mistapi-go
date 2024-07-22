package models

import (
    "encoding/json"
)

// ApSwitchSetting represents a ApSwitchSetting struct.
type ApSwitchSetting struct {
    // additional VLAN IDs, only valid in mesh base mode
    AdditionalVlanIds    []int          `json:"additional_vlan_ids,omitempty"`
    EnableVlan           *bool          `json:"enable_vlan,omitempty"`
    // native VLAN id, optional
    PortVlanId           *int           `json:"port_vlan_id,omitempty"`
    // list of VLAN ids this
    VlanIds              []int          `json:"vlan_ids,omitempty"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for ApSwitchSetting.
// It customizes the JSON marshaling process for ApSwitchSetting objects.
func (a ApSwitchSetting) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(a.toMap())
}

// toMap converts the ApSwitchSetting object to a map representation for JSON marshaling.
func (a ApSwitchSetting) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, a.AdditionalProperties)
    if a.AdditionalVlanIds != nil {
        structMap["additional_vlan_ids"] = a.AdditionalVlanIds
    }
    if a.EnableVlan != nil {
        structMap["enable_vlan"] = a.EnableVlan
    }
    if a.PortVlanId != nil {
        structMap["port_vlan_id"] = a.PortVlanId
    }
    if a.VlanIds != nil {
        structMap["vlan_ids"] = a.VlanIds
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ApSwitchSetting.
// It customizes the JSON unmarshaling process for ApSwitchSetting objects.
func (a *ApSwitchSetting) UnmarshalJSON(input []byte) error {
    var temp apSwitchSetting
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "additional_vlan_ids", "enable_vlan", "port_vlan_id", "vlan_ids")
    if err != nil {
    	return err
    }
    
    a.AdditionalProperties = additionalProperties
    a.AdditionalVlanIds = temp.AdditionalVlanIds
    a.EnableVlan = temp.EnableVlan
    a.PortVlanId = temp.PortVlanId
    a.VlanIds = temp.VlanIds
    return nil
}

// apSwitchSetting is a temporary struct used for validating the fields of ApSwitchSetting.
type apSwitchSetting  struct {
    AdditionalVlanIds []int `json:"additional_vlan_ids,omitempty"`
    EnableVlan        *bool `json:"enable_vlan,omitempty"`
    PortVlanId        *int  `json:"port_vlan_id,omitempty"`
    VlanIds           []int `json:"vlan_ids,omitempty"`
}
