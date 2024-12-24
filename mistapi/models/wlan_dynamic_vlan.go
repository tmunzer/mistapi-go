package models

import (
    "encoding/json"
    "fmt"
)

// WlanDynamicVlan represents a WlanDynamicVlan struct.
// for 802.1x
type WlanDynamicVlan struct {
    // vlan_id to use when there’s no match from RADIUS
    DefaultVlanId        *WlanDynamicVlanDefaultVlanIdDeprecated `json:"default_vlan_id,omitempty"`  // Deprecated
    // Default VLAN ID(s) can be a number, a range of VLAN IDs, a variable or multiple numbers, ranges or variables as a VLAN pool. Default VLAN as a pool of VLANS requires 0.14.x or newer firmware
    DefaultVlanIds       []WlanDynamicVlanDefaultVlanId          `json:"default_vlan_ids,omitempty"`
    // Requires `vlan_enabled`==`true` to be set to `true`. Whether to enable dynamic vlan
    Enabled              *bool                                   `json:"enabled,omitempty"`
    // vlan_ids to be locally bridged
    LocalVlanIds         []VlanIdWithVariable                    `json:"local_vlan_ids,omitempty"`
    // standard (using Tunnel-Private-Group-ID, widely supported), airespace-interface-name (Airespace/Cisco). enum: `airespace-interface-name`, `standard`
    Type                 *WlanDynamicVlanTypeEnum                `json:"type,omitempty"`
    // map between vlan_id (as string) to airespace interface names (comma-separated) or null for stndard mapping
    // * if `dynamic_vlan.type`==`standard`, property key is the Vlan ID and property value is \"\"
    // * if `dynamic_vlan.type`==`airespace-interface-name`, property key is the Vlan ID and property value is the Airespace Interface Name
    Vlans                map[string]string                       `json:"vlans,omitempty"`
    AdditionalProperties map[string]interface{}                  `json:"_"`
}

// String implements the fmt.Stringer interface for WlanDynamicVlan,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (w WlanDynamicVlan) String() string {
    return fmt.Sprintf(
    	"WlanDynamicVlan[DefaultVlanId=%v, DefaultVlanIds=%v, Enabled=%v, LocalVlanIds=%v, Type=%v, Vlans=%v, AdditionalProperties=%v]",
    	w.DefaultVlanId, w.DefaultVlanIds, w.Enabled, w.LocalVlanIds, w.Type, w.Vlans, w.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for WlanDynamicVlan.
// It customizes the JSON marshaling process for WlanDynamicVlan objects.
func (w WlanDynamicVlan) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(w.AdditionalProperties,
        "default_vlan_id", "default_vlan_ids", "enabled", "local_vlan_ids", "type", "vlans"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(w.toMap())
}

// toMap converts the WlanDynamicVlan object to a map representation for JSON marshaling.
func (w WlanDynamicVlan) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, w.AdditionalProperties)
    if w.DefaultVlanId != nil {
        structMap["default_vlan_id"] = w.DefaultVlanId.toMap()
    }
    if w.DefaultVlanIds != nil {
        structMap["default_vlan_ids"] = w.DefaultVlanIds
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
    var temp tempWlanDynamicVlan
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "default_vlan_id", "default_vlan_ids", "enabled", "local_vlan_ids", "type", "vlans")
    if err != nil {
    	return err
    }
    w.AdditionalProperties = additionalProperties
    
    w.DefaultVlanId = temp.DefaultVlanId
    w.DefaultVlanIds = temp.DefaultVlanIds
    w.Enabled = temp.Enabled
    w.LocalVlanIds = temp.LocalVlanIds
    w.Type = temp.Type
    w.Vlans = temp.Vlans
    return nil
}

// tempWlanDynamicVlan is a temporary struct used for validating the fields of WlanDynamicVlan.
type tempWlanDynamicVlan  struct {
    DefaultVlanId  *WlanDynamicVlanDefaultVlanIdDeprecated `json:"default_vlan_id,omitempty"`
    DefaultVlanIds []WlanDynamicVlanDefaultVlanId          `json:"default_vlan_ids,omitempty"`
    Enabled        *bool                                   `json:"enabled,omitempty"`
    LocalVlanIds   []VlanIdWithVariable                    `json:"local_vlan_ids,omitempty"`
    Type           *WlanDynamicVlanTypeEnum                `json:"type,omitempty"`
    Vlans          map[string]string                       `json:"vlans,omitempty"`
}
