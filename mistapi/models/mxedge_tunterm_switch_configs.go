package models

import (
	"encoding/json"
)

// MxedgeTuntermSwitchConfigs represents a MxedgeTuntermSwitchConfigs struct.
// if custom vlan settings are desired
type MxedgeTuntermSwitchConfigs struct {
	Enabled              *bool          `json:"enabled,omitempty"`
	AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for MxedgeTuntermSwitchConfigs.
// It customizes the JSON marshaling process for MxedgeTuntermSwitchConfigs objects.
func (m MxedgeTuntermSwitchConfigs) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(m.toMap())
}

// toMap converts the MxedgeTuntermSwitchConfigs object to a map representation for JSON marshaling.
func (m MxedgeTuntermSwitchConfigs) toMap() map[string]any {
	structMap := make(map[string]any)
	MapAdditionalProperties(structMap, m.AdditionalProperties)
	if m.Enabled != nil {
		structMap["enabled"] = m.Enabled
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for MxedgeTuntermSwitchConfigs.
// It customizes the JSON unmarshaling process for MxedgeTuntermSwitchConfigs objects.
func (m *MxedgeTuntermSwitchConfigs) UnmarshalJSON(input []byte) error {
	var temp tempMxedgeTuntermSwitchConfigs
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := UnmarshalAdditionalProperties(input, "enabled")
	if err != nil {
		return err
	}

	m.AdditionalProperties = additionalProperties
	m.Enabled = temp.Enabled
	return nil
}

// tempMxedgeTuntermSwitchConfigs is a temporary struct used for validating the fields of MxedgeTuntermSwitchConfigs.
type tempMxedgeTuntermSwitchConfigs struct {
	Enabled *bool `json:"enabled,omitempty"`
}
