package models

import (
	"encoding/json"
)

// SwitchStpConfig represents a SwitchStpConfig struct.
type SwitchStpConfig struct {
	// ignored for switches participating in EVPN
	VstpEnabled          *bool          `json:"vstp_enabled,omitempty"`
	AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for SwitchStpConfig.
// It customizes the JSON marshaling process for SwitchStpConfig objects.
func (s SwitchStpConfig) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(s.toMap())
}

// toMap converts the SwitchStpConfig object to a map representation for JSON marshaling.
func (s SwitchStpConfig) toMap() map[string]any {
	structMap := make(map[string]any)
	MapAdditionalProperties(structMap, s.AdditionalProperties)
	if s.VstpEnabled != nil {
		structMap["vstp_enabled"] = s.VstpEnabled
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SwitchStpConfig.
// It customizes the JSON unmarshaling process for SwitchStpConfig objects.
func (s *SwitchStpConfig) UnmarshalJSON(input []byte) error {
	var temp tempSwitchStpConfig
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := UnmarshalAdditionalProperties(input, "vstp_enabled")
	if err != nil {
		return err
	}

	s.AdditionalProperties = additionalProperties
	s.VstpEnabled = temp.VstpEnabled
	return nil
}

// tempSwitchStpConfig is a temporary struct used for validating the fields of SwitchStpConfig.
type tempSwitchStpConfig struct {
	VstpEnabled *bool `json:"vstp_enabled,omitempty"`
}
