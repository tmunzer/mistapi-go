package models

import (
    "encoding/json"
)

// MxedgeTuntermDhcpdConfigProperty represents a MxedgeTuntermDhcpdConfigProperty struct.
type MxedgeTuntermDhcpdConfigProperty struct {
    Enabled              *bool                             `json:"enabled,omitempty"`
    // list of DHCP servers; required if `type`==`relay`
    Servers              []string                          `json:"servers,omitempty"`
    // enum: `relay`
    Type                 *MxedgeTuntermDhcpdConfigTypeEnum `json:"type,omitempty"`
    AdditionalProperties map[string]interface{}            `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for MxedgeTuntermDhcpdConfigProperty.
// It customizes the JSON marshaling process for MxedgeTuntermDhcpdConfigProperty objects.
func (m MxedgeTuntermDhcpdConfigProperty) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(m.AdditionalProperties,
        "enabled", "servers", "type"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(m.toMap())
}

// toMap converts the MxedgeTuntermDhcpdConfigProperty object to a map representation for JSON marshaling.
func (m MxedgeTuntermDhcpdConfigProperty) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, m.AdditionalProperties)
    if m.Enabled != nil {
        structMap["enabled"] = m.Enabled
    }
    if m.Servers != nil {
        structMap["servers"] = m.Servers
    }
    if m.Type != nil {
        structMap["type"] = m.Type
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for MxedgeTuntermDhcpdConfigProperty.
// It customizes the JSON unmarshaling process for MxedgeTuntermDhcpdConfigProperty objects.
func (m *MxedgeTuntermDhcpdConfigProperty) UnmarshalJSON(input []byte) error {
    var temp tempMxedgeTuntermDhcpdConfigProperty
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "enabled", "servers", "type")
    if err != nil {
    	return err
    }
    m.AdditionalProperties = additionalProperties
    
    m.Enabled = temp.Enabled
    m.Servers = temp.Servers
    m.Type = temp.Type
    return nil
}

// tempMxedgeTuntermDhcpdConfigProperty is a temporary struct used for validating the fields of MxedgeTuntermDhcpdConfigProperty.
type tempMxedgeTuntermDhcpdConfigProperty  struct {
    Enabled *bool                             `json:"enabled,omitempty"`
    Servers []string                          `json:"servers,omitempty"`
    Type    *MxedgeTuntermDhcpdConfigTypeEnum `json:"type,omitempty"`
}
