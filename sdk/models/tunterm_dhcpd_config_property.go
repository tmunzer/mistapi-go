package models

import (
    "encoding/json"
)

// TuntermDhcpdConfigProperty represents a TuntermDhcpdConfigProperty struct.
type TuntermDhcpdConfigProperty struct {
    Enabled              *bool                 `json:"enabled,omitempty"`
    Servers              []string              `json:"servers,omitempty"`
    Type                 *TuntermDhcpdTypeEnum `json:"type,omitempty"`
    AdditionalProperties map[string]any        `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for TuntermDhcpdConfigProperty.
// It customizes the JSON marshaling process for TuntermDhcpdConfigProperty objects.
func (t TuntermDhcpdConfigProperty) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(t.toMap())
}

// toMap converts the TuntermDhcpdConfigProperty object to a map representation for JSON marshaling.
func (t TuntermDhcpdConfigProperty) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, t.AdditionalProperties)
    if t.Enabled != nil {
        structMap["enabled"] = t.Enabled
    }
    if t.Servers != nil {
        structMap["servers"] = t.Servers
    }
    if t.Type != nil {
        structMap["type"] = t.Type
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for TuntermDhcpdConfigProperty.
// It customizes the JSON unmarshaling process for TuntermDhcpdConfigProperty objects.
func (t *TuntermDhcpdConfigProperty) UnmarshalJSON(input []byte) error {
    var temp tuntermDhcpdConfigProperty
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "enabled", "servers", "type")
    if err != nil {
    	return err
    }
    
    t.AdditionalProperties = additionalProperties
    t.Enabled = temp.Enabled
    t.Servers = temp.Servers
    t.Type = temp.Type
    return nil
}

// tuntermDhcpdConfigProperty is a temporary struct used for validating the fields of TuntermDhcpdConfigProperty.
type tuntermDhcpdConfigProperty  struct {
    Enabled *bool                 `json:"enabled,omitempty"`
    Servers []string              `json:"servers,omitempty"`
    Type    *TuntermDhcpdTypeEnum `json:"type,omitempty"`
}
