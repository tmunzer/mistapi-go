package models

import (
	"encoding/json"
)

// VrrpConfig represents a VrrpConfig struct.
// Junos VRRP config
type VrrpConfig struct {
	Enabled *bool `json:"enabled,omitempty"`
	// Property key is the VRRP name
	Groups               map[string]VrrpConfigGroup `json:"groups,omitempty"`
	AdditionalProperties map[string]any             `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for VrrpConfig.
// It customizes the JSON marshaling process for VrrpConfig objects.
func (v VrrpConfig) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(v.toMap())
}

// toMap converts the VrrpConfig object to a map representation for JSON marshaling.
func (v VrrpConfig) toMap() map[string]any {
	structMap := make(map[string]any)
	MapAdditionalProperties(structMap, v.AdditionalProperties)
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
	additionalProperties, err := UnmarshalAdditionalProperties(input, "enabled", "groups")
	if err != nil {
		return err
	}

	v.AdditionalProperties = additionalProperties
	v.Enabled = temp.Enabled
	v.Groups = temp.Groups
	return nil
}

// tempVrrpConfig is a temporary struct used for validating the fields of VrrpConfig.
type tempVrrpConfig struct {
	Enabled *bool                      `json:"enabled,omitempty"`
	Groups  map[string]VrrpConfigGroup `json:"groups,omitempty"`
}
