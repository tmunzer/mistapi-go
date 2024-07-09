package models

import (
    "encoding/json"
)

// ApPortConfigDynamicVlan represents a ApPortConfigDynamicVlan struct.
// optional dynamic vlan
type ApPortConfigDynamicVlan struct {
    DefaultVlanId        *int              `json:"default_vlan_id,omitempty"`
    Enabled              *bool             `json:"enabled,omitempty"`
    Type                 *string           `json:"type,omitempty"`
    Vlans                map[string]string `json:"vlans,omitempty"`
    AdditionalProperties map[string]any    `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for ApPortConfigDynamicVlan.
// It customizes the JSON marshaling process for ApPortConfigDynamicVlan objects.
func (a ApPortConfigDynamicVlan) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(a.toMap())
}

// toMap converts the ApPortConfigDynamicVlan object to a map representation for JSON marshaling.
func (a ApPortConfigDynamicVlan) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, a.AdditionalProperties)
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
    var temp apPortConfigDynamicVlan
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "default_vlan_id", "enabled", "type", "vlans")
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

// apPortConfigDynamicVlan is a temporary struct used for validating the fields of ApPortConfigDynamicVlan.
type apPortConfigDynamicVlan  struct {
    DefaultVlanId *int              `json:"default_vlan_id,omitempty"`
    Enabled       *bool             `json:"enabled,omitempty"`
    Type          *string           `json:"type,omitempty"`
    Vlans         map[string]string `json:"vlans,omitempty"`
}
