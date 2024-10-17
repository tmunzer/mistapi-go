package models

import (
	"encoding/json"
)

// DhcpdConfig represents a DhcpdConfig struct.
type DhcpdConfig struct {
	// if set to `true`, enable the DHCP server
	Enabled              *bool          `json:"enabled,omitempty"`
	AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for DhcpdConfig.
// It customizes the JSON marshaling process for DhcpdConfig objects.
func (d DhcpdConfig) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(d.toMap())
}

// toMap converts the DhcpdConfig object to a map representation for JSON marshaling.
func (d DhcpdConfig) toMap() map[string]any {
	structMap := make(map[string]any)
	MapAdditionalProperties(structMap, d.AdditionalProperties)
	if d.Enabled != nil {
		structMap["enabled"] = d.Enabled
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for DhcpdConfig.
// It customizes the JSON unmarshaling process for DhcpdConfig objects.
func (d *DhcpdConfig) UnmarshalJSON(input []byte) error {
	var temp tempDhcpdConfig
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := UnmarshalAdditionalProperties(input, "enabled")
	if err != nil {
		return err
	}

	d.AdditionalProperties = additionalProperties
	d.Enabled = temp.Enabled
	return nil
}

// tempDhcpdConfig is a temporary struct used for validating the fields of DhcpdConfig.
type tempDhcpdConfig struct {
	Enabled *bool `json:"enabled,omitempty"`
}
