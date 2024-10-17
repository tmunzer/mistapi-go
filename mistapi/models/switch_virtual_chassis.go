package models

import (
	"encoding/json"
)

// SwitchVirtualChassis represents a SwitchVirtualChassis struct.
// required for preprovisioned Virtual Chassis
type SwitchVirtualChassis struct {
	// list of Virtual Chassis members
	Members []SwitchVirtualChassisMember `json:"members,omitempty"`
	// to configure whether the VC is preprovisioned or nonprovisioned
	Preprovisioned       *bool          `json:"preprovisioned,omitempty"`
	AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for SwitchVirtualChassis.
// It customizes the JSON marshaling process for SwitchVirtualChassis objects.
func (s SwitchVirtualChassis) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(s.toMap())
}

// toMap converts the SwitchVirtualChassis object to a map representation for JSON marshaling.
func (s SwitchVirtualChassis) toMap() map[string]any {
	structMap := make(map[string]any)
	MapAdditionalProperties(structMap, s.AdditionalProperties)
	if s.Members != nil {
		structMap["members"] = s.Members
	}
	if s.Preprovisioned != nil {
		structMap["preprovisioned"] = s.Preprovisioned
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SwitchVirtualChassis.
// It customizes the JSON unmarshaling process for SwitchVirtualChassis objects.
func (s *SwitchVirtualChassis) UnmarshalJSON(input []byte) error {
	var temp tempSwitchVirtualChassis
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := UnmarshalAdditionalProperties(input, "members", "preprovisioned")
	if err != nil {
		return err
	}

	s.AdditionalProperties = additionalProperties
	s.Members = temp.Members
	s.Preprovisioned = temp.Preprovisioned
	return nil
}

// tempSwitchVirtualChassis is a temporary struct used for validating the fields of SwitchVirtualChassis.
type tempSwitchVirtualChassis struct {
	Members        []SwitchVirtualChassisMember `json:"members,omitempty"`
	Preprovisioned *bool                        `json:"preprovisioned,omitempty"`
}
