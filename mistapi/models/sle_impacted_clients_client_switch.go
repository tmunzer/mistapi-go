package models

import (
	"encoding/json"
)

// SleImpactedClientsClientSwitch represents a SleImpactedClientsClientSwitch struct.
type SleImpactedClientsClientSwitch struct {
	Interfaces           []string       `json:"interfaces,omitempty"`
	SwitchMac            *string        `json:"switch_mac,omitempty"`
	SwitchName           *string        `json:"switch_name,omitempty"`
	AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for SleImpactedClientsClientSwitch.
// It customizes the JSON marshaling process for SleImpactedClientsClientSwitch objects.
func (s SleImpactedClientsClientSwitch) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(s.toMap())
}

// toMap converts the SleImpactedClientsClientSwitch object to a map representation for JSON marshaling.
func (s SleImpactedClientsClientSwitch) toMap() map[string]any {
	structMap := make(map[string]any)
	MapAdditionalProperties(structMap, s.AdditionalProperties)
	if s.Interfaces != nil {
		structMap["interfaces"] = s.Interfaces
	}
	if s.SwitchMac != nil {
		structMap["switch_mac"] = s.SwitchMac
	}
	if s.SwitchName != nil {
		structMap["switch_name"] = s.SwitchName
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SleImpactedClientsClientSwitch.
// It customizes the JSON unmarshaling process for SleImpactedClientsClientSwitch objects.
func (s *SleImpactedClientsClientSwitch) UnmarshalJSON(input []byte) error {
	var temp tempSleImpactedClientsClientSwitch
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := UnmarshalAdditionalProperties(input, "interfaces", "switch_mac", "switch_name")
	if err != nil {
		return err
	}

	s.AdditionalProperties = additionalProperties
	s.Interfaces = temp.Interfaces
	s.SwitchMac = temp.SwitchMac
	s.SwitchName = temp.SwitchName
	return nil
}

// tempSleImpactedClientsClientSwitch is a temporary struct used for validating the fields of SleImpactedClientsClientSwitch.
type tempSleImpactedClientsClientSwitch struct {
	Interfaces []string `json:"interfaces,omitempty"`
	SwitchMac  *string  `json:"switch_mac,omitempty"`
	SwitchName *string  `json:"switch_name,omitempty"`
}
