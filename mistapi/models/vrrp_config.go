package models

import (
    "encoding/json"
    "fmt"
)

// VrrpConfig represents a VrrpConfig struct.
// Junos VRRP config
type VrrpConfig struct {
    Enabled              *bool                      `json:"enabled,omitempty"`
    // Property key is the VRRP name
    Groups               map[string]VrrpConfigGroup `json:"groups,omitempty"`
    AdditionalProperties map[string]interface{}     `json:"_"`
}

// String implements the fmt.Stringer interface for VrrpConfig,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (v VrrpConfig) String() string {
    return fmt.Sprintf(
    	"VrrpConfig[Enabled=%v, Groups=%v, AdditionalProperties=%v]",
    	v.Enabled, v.Groups, v.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for VrrpConfig.
// It customizes the JSON marshaling process for VrrpConfig objects.
func (v VrrpConfig) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(v.AdditionalProperties,
        "enabled", "groups"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(v.toMap())
}

// toMap converts the VrrpConfig object to a map representation for JSON marshaling.
func (v VrrpConfig) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, v.AdditionalProperties)
    if v.Enabled != nil {
        structMap["enabled"] = v.Enabled
    }
    if v.Groups != nil {
        structMap["groups"] = v.Groups
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for VrrpConfig.
// It customizes the JSON unmarshaling process for VrrpConfig objects.
func (v *VrrpConfig) UnmarshalJSON(input []byte) error {
    var temp tempVrrpConfig
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "enabled", "groups")
    if err != nil {
    	return err
    }
    v.AdditionalProperties = additionalProperties
    
    v.Enabled = temp.Enabled
    v.Groups = temp.Groups
    return nil
}

// tempVrrpConfig is a temporary struct used for validating the fields of VrrpConfig.
type tempVrrpConfig  struct {
    Enabled *bool                      `json:"enabled,omitempty"`
    Groups  map[string]VrrpConfigGroup `json:"groups,omitempty"`
}
