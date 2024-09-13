package models

import (
    "encoding/json"
)

// PortChannelization represents a PortChannelization struct.
type PortChannelization struct {
    // Property key is the interface name or range (e.g. `et-0/0/47`, `et-0/0/48-49`), Property value is the interface speed (e.g. `25g`, `50g`)
    AdditionalProperties  *string        `json:"additionalProperties,omitempty"`
    Enabled               *bool          `json:"enabled,omitempty"`
    AdditionalProperties2 map[string]any `json:"_"`
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
    MapAdditionalProperties(structMap, p.AdditionalProperties2)
    if p.AdditionalProperties != nil {
        structMap["additionalProperties"] = p.AdditionalProperties
    }
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
    additionalProperties2, err := UnmarshalAdditionalProperties(input, "additionalProperties", "enabled")
    if err != nil {
    	return err
    }
    
    p.AdditionalProperties2 = additionalProperties2
    p.AdditionalProperties = temp.AdditionalProperties
    p.Enabled = temp.Enabled
    return nil
}

// tempPortChannelization is a temporary struct used for validating the fields of PortChannelization.
type tempPortChannelization  struct {
    AdditionalProperties *string `json:"additionalProperties,omitempty"`
    Enabled              *bool   `json:"enabled,omitempty"`
}
