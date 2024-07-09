package models

import (
    "encoding/json"
)

// WlanDynamicVlan represents a WlanDynamicVlan struct.
// for 802.1x
type WlanDynamicVlan struct {
    // vlan_id to use when there’s no match from RADIUS
    DefaultVlanId        Optional[int]            `json:"default_vlan_id"`
    // whether to enable dynamic vlan
    Enabled              *bool                    `json:"enabled,omitempty"`
    // vlan_ids to be locally bridged
    LocalVlanIds         []int                    `json:"local_vlan_ids,omitempty"`
    // standard (using Tunnel-Private-Group-ID, widely supported), airespace-interface-name (Airespace/Cisco)
    Type                 *WlanDynamicVlanTypeEnum `json:"type,omitempty"`
    // map between vlan_id (as string) to airespace interface names (comma-separated) or null for stndard mapping
    //   * if `dynamic_vlan.type`==`standard`, property key is the Vlan ID and property value is ""
    //   * if `dynamic_vlan.type`==`airespace-interface-name`, property key is the Vlan ID and property value is the Airespace Interface Name
    Vlans                map[string]string        `json:"vlans,omitempty"`
    AdditionalProperties map[string]any           `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for WlanDynamicVlan.
// It customizes the JSON marshaling process for WlanDynamicVlan objects.
func (w WlanDynamicVlan) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(w.toMap())
}

// toMap converts the WlanDynamicVlan object to a map representation for JSON marshaling.
func (w WlanDynamicVlan) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, w.AdditionalProperties)
    if w.DefaultVlanId.IsValueSet() {
        if w.DefaultVlanId.Value() != nil {
            structMap["default_vlan_id"] = w.DefaultVlanId.Value()
        } else {
            structMap["default_vlan_id"] = nil
        }
    }
    if w.Enabled != nil {
        structMap["enabled"] = w.Enabled
    }
    if w.LocalVlanIds != nil {
        structMap["local_vlan_ids"] = w.LocalVlanIds
    }
    if w.Type != nil {
        structMap["type"] = w.Type
    }
    if w.Vlans != nil {
        structMap["vlans"] = w.Vlans
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for WlanDynamicVlan.
// It customizes the JSON unmarshaling process for WlanDynamicVlan objects.
func (w *WlanDynamicVlan) UnmarshalJSON(input []byte) error {
    var temp wlanDynamicVlan
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "default_vlan_id", "enabled", "local_vlan_ids", "type", "vlans")
    if err != nil {
    	return err
    }
    
    w.AdditionalProperties = additionalProperties
    w.DefaultVlanId = temp.DefaultVlanId
    w.Enabled = temp.Enabled
    w.LocalVlanIds = temp.LocalVlanIds
    w.Type = temp.Type
    w.Vlans = temp.Vlans
    return nil
}

// wlanDynamicVlan is a temporary struct used for validating the fields of WlanDynamicVlan.
type wlanDynamicVlan  struct {
    DefaultVlanId Optional[int]            `json:"default_vlan_id"`
    Enabled       *bool                    `json:"enabled,omitempty"`
    LocalVlanIds  []int                    `json:"local_vlan_ids,omitempty"`
    Type          *WlanDynamicVlanTypeEnum `json:"type,omitempty"`
    Vlans         map[string]string        `json:"vlans,omitempty"`
}
