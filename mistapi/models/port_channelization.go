package models

import (
	"encoding/json"
)

// PortChannelization represents a PortChannelization struct.
type PortChannelization struct {
	Enabled              *bool          `json:"enabled,omitempty"`
	AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for PortChannelization.
// It customizes the JSON marshaling process for PortChannelization objects.
func (p PortChannelization) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(p.toMap())
}

// toMap converts the PortChannelization object to a map representation for JSON marshaling.
func (p PortChannelization) toMap() map[string]any {
	structMap := make(map[string]any)
	MapAdditionalProperties(structMap, p.AdditionalProperties)
	if p.Enabled != nil {
		structMap["enabled"] = p.Enabled
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for PortChannelization.
// It customizes the JSON unmarshaling process for PortChannelization objects.
func (p *PortChannelization) UnmarshalJSON(input []byte) error {
	var temp tempPortChannelization
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := UnmarshalAdditionalProperties(input, "enabled")
	if err != nil {
		return err
	}

	p.AdditionalProperties = additionalProperties
	p.Enabled = temp.Enabled
	return nil
}

// tempPortChannelization is a temporary struct used for validating the fields of PortChannelization.
type tempPortChannelization struct {
	Enabled *bool `json:"enabled,omitempty"`
}
