package models

import (
    "encoding/json"
    "fmt"
)

// MxedgeTuntermSwitchConfigs represents a MxedgeTuntermSwitchConfigs struct.
// If custom vlan settings are desired
type MxedgeTuntermSwitchConfigs struct {
    Enabled              *bool                                `json:"enabled,omitempty"`
    AdditionalProperties map[string]MxedgeTuntermSwitchConfig `json:"_"`
}

// String implements the fmt.Stringer interface for MxedgeTuntermSwitchConfigs,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (m MxedgeTuntermSwitchConfigs) String() string {
    return fmt.Sprintf(
    	"MxedgeTuntermSwitchConfigs[Enabled=%v, AdditionalProperties=%v]",
    	m.Enabled, m.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for MxedgeTuntermSwitchConfigs.
// It customizes the JSON marshaling process for MxedgeTuntermSwitchConfigs objects.
func (m MxedgeTuntermSwitchConfigs) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(m.AdditionalProperties,
        "enabled"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(m.toMap())
}

// toMap converts the MxedgeTuntermSwitchConfigs object to a map representation for JSON marshaling.
func (m MxedgeTuntermSwitchConfigs) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, m.AdditionalProperties)
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
    additionalProperties, err := ExtractAdditionalProperties[MxedgeTuntermSwitchConfig](input, "enabled")
    if err != nil {
    	return err
    }
    m.AdditionalProperties = additionalProperties
    
    m.Enabled = temp.Enabled
    return nil
}

// tempMxedgeTuntermSwitchConfigs is a temporary struct used for validating the fields of MxedgeTuntermSwitchConfigs.
type tempMxedgeTuntermSwitchConfigs  struct {
    Enabled *bool `json:"enabled,omitempty"`
}
