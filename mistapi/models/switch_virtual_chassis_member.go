package models

import (
	"encoding/json"
)

// SwitchVirtualChassisMember represents a SwitchVirtualChassisMember struct.
type SwitchVirtualChassisMember struct {
	// fpc0, same as the mac of device_id
	Mac      *string `json:"mac,omitempty"`
	MemberId *int    `json:"member_id,omitempty"`
	// Both vc_role master and backup will be matched to routing-engine role in Junos preprovisioned VC config. enum: `backup`, `linecard`, `master`
	VcRole               *SwitchVirtualChassisMemberVcRoleEnum `json:"vc_role,omitempty"`
	AdditionalProperties map[string]any                        `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for SwitchVirtualChassisMember.
// It customizes the JSON marshaling process for SwitchVirtualChassisMember objects.
func (s SwitchVirtualChassisMember) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(s.toMap())
}

// toMap converts the SwitchVirtualChassisMember object to a map representation for JSON marshaling.
func (s SwitchVirtualChassisMember) toMap() map[string]any {
	structMap := make(map[string]any)
	MapAdditionalProperties(structMap, s.AdditionalProperties)
	if s.Mac != nil {
		structMap["mac"] = s.Mac
	}
	if s.MemberId != nil {
		structMap["member_id"] = s.MemberId
	}
	if s.VcRole != nil {
		structMap["vc_role"] = s.VcRole
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SwitchVirtualChassisMember.
// It customizes the JSON unmarshaling process for SwitchVirtualChassisMember objects.
func (s *SwitchVirtualChassisMember) UnmarshalJSON(input []byte) error {
	var temp tempSwitchVirtualChassisMember
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := UnmarshalAdditionalProperties(input, "mac", "member_id", "vc_role")
	if err != nil {
		return err
	}

	s.AdditionalProperties = additionalProperties
	s.Mac = temp.Mac
	s.MemberId = temp.MemberId
	s.VcRole = temp.VcRole
	return nil
}

// tempSwitchVirtualChassisMember is a temporary struct used for validating the fields of SwitchVirtualChassisMember.
type tempSwitchVirtualChassisMember struct {
	Mac      *string                               `json:"mac,omitempty"`
	MemberId *int                                  `json:"member_id,omitempty"`
	VcRole   *SwitchVirtualChassisMemberVcRoleEnum `json:"vc_role,omitempty"`
}
