package models

import (
    "encoding/json"
    "fmt"
)

// ApSwitchSetting represents a ApSwitchSetting struct.
type ApSwitchSetting struct {
    EnableVlan           *bool                  `json:"enable_vlan,omitempty"`
    // Native VLAN id, optional
    PortVlanId           *int                   `json:"port_vlan_id,omitempty"`
    // List of VLAN ids this
    VlanIds              []int                  `json:"vlan_ids,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for ApSwitchSetting,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (a ApSwitchSetting) String() string {
    return fmt.Sprintf(
    	"ApSwitchSetting[EnableVlan=%v, PortVlanId=%v, VlanIds=%v, AdditionalProperties=%v]",
    	a.EnableVlan, a.PortVlanId, a.VlanIds, a.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for ApSwitchSetting.
// It customizes the JSON marshaling process for ApSwitchSetting objects.
func (a ApSwitchSetting) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(a.AdditionalProperties,
        "enable_vlan", "port_vlan_id", "vlan_ids"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(a.toMap())
}

// toMap converts the ApSwitchSetting object to a map representation for JSON marshaling.
func (a ApSwitchSetting) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, a.AdditionalProperties)
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
    var temp tempApSwitchSetting
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "enable_vlan", "port_vlan_id", "vlan_ids")
    if err != nil {
    	return err
    }
    a.AdditionalProperties = additionalProperties
    
    a.EnableVlan = temp.EnableVlan
    a.PortVlanId = temp.PortVlanId
    a.VlanIds = temp.VlanIds
    return nil
}

// tempApSwitchSetting is a temporary struct used for validating the fields of ApSwitchSetting.
type tempApSwitchSetting  struct {
    EnableVlan *bool `json:"enable_vlan,omitempty"`
    PortVlanId *int  `json:"port_vlan_id,omitempty"`
    VlanIds    []int `json:"vlan_ids,omitempty"`
}
