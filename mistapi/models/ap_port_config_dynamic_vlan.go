package models

import (
    "encoding/json"
    "fmt"
)

// ApPortConfigDynamicVlan represents a ApPortConfigDynamicVlan struct.
// Optional dynamic vlan
type ApPortConfigDynamicVlan struct {
    DefaultVlanId        *int                   `json:"default_vlan_id,omitempty"`
    Enabled              *bool                  `json:"enabled,omitempty"`
    Type                 *string                `json:"type,omitempty"`
    Vlans                map[string]string      `json:"vlans,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for ApPortConfigDynamicVlan,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (a ApPortConfigDynamicVlan) String() string {
    return fmt.Sprintf(
    	"ApPortConfigDynamicVlan[DefaultVlanId=%v, Enabled=%v, Type=%v, Vlans=%v, AdditionalProperties=%v]",
    	a.DefaultVlanId, a.Enabled, a.Type, a.Vlans, a.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for ApPortConfigDynamicVlan.
// It customizes the JSON marshaling process for ApPortConfigDynamicVlan objects.
func (a ApPortConfigDynamicVlan) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(a.AdditionalProperties,
        "default_vlan_id", "enabled", "type", "vlans"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(a.toMap())
}

// toMap converts the ApPortConfigDynamicVlan object to a map representation for JSON marshaling.
func (a ApPortConfigDynamicVlan) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, a.AdditionalProperties)
    if a.DefaultVlanId != nil {
        structMap["default_vlan_id"] = a.DefaultVlanId
    }
    if a.Enabled != nil {
        structMap["enabled"] = a.Enabled
    }
    if a.Type != nil {
        structMap["type"] = a.Type
    }
    if a.Vlans != nil {
        structMap["vlans"] = a.Vlans
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ApPortConfigDynamicVlan.
// It customizes the JSON unmarshaling process for ApPortConfigDynamicVlan objects.
func (a *ApPortConfigDynamicVlan) UnmarshalJSON(input []byte) error {
    var temp tempApPortConfigDynamicVlan
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "default_vlan_id", "enabled", "type", "vlans")
    if err != nil {
    	return err
    }
    a.AdditionalProperties = additionalProperties
    
    a.DefaultVlanId = temp.DefaultVlanId
    a.Enabled = temp.Enabled
    a.Type = temp.Type
    a.Vlans = temp.Vlans
    return nil
}

// tempApPortConfigDynamicVlan is a temporary struct used for validating the fields of ApPortConfigDynamicVlan.
type tempApPortConfigDynamicVlan  struct {
    DefaultVlanId *int              `json:"default_vlan_id,omitempty"`
    Enabled       *bool             `json:"enabled,omitempty"`
    Type          *string           `json:"type,omitempty"`
    Vlans         map[string]string `json:"vlans,omitempty"`
}
