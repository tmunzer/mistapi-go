package models

import (
	"encoding/json"
)

// VirtualChassisConfigMember represents a VirtualChassisConfigMember struct.
type VirtualChassisConfigMember struct {
	Locating *bool `json:"locating,omitempty"`
	// fpc0, same as the mac of device_id
	Mac *string `json:"mac,omitempty"`
	// to create a preprovisionned virtual chassis
	Member  *int     `json:"member,omitempty"`
	VcPorts []string `json:"vc_ports,omitempty"`
	// enum: `backup`, `linecard`, `master`
	VcRole               *VirtualChassisConfigMemberVcRoleEnum `json:"vc_role,omitempty"`
	AdditionalProperties map[string]any                        `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for VirtualChassisConfigMember.
// It customizes the JSON marshaling process for VirtualChassisConfigMember objects.
func (v VirtualChassisConfigMember) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(v.toMap())
}

// toMap converts the VirtualChassisConfigMember object to a map representation for JSON marshaling.
func (v VirtualChassisConfigMember) toMap() map[string]any {
	structMap := make(map[string]any)
	MapAdditionalProperties(structMap, v.AdditionalProperties)
	if v.Locating != nil {
		structMap["locating"] = v.Locating
	}
	if v.Mac != nil {
		structMap["mac"] = v.Mac
	}
	if v.Member != nil {
		structMap["member"] = v.Member
	}
	if v.VcPorts != nil {
		structMap["vc_ports"] = v.VcPorts
	}
	if v.VcRole != nil {
		structMap["vc_role"] = v.VcRole
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for VirtualChassisConfigMember.
// It customizes the JSON unmarshaling process for VirtualChassisConfigMember objects.
func (v *VirtualChassisConfigMember) UnmarshalJSON(input []byte) error {
	var temp tempVirtualChassisConfigMember
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := UnmarshalAdditionalProperties(input, "locating", "mac", "member", "vc_ports", "vc_role")
	if err != nil {
		return err
	}

	v.AdditionalProperties = additionalProperties
	v.Locating = temp.Locating
	v.Mac = temp.Mac
	v.Member = temp.Member
	v.VcPorts = temp.VcPorts
	v.VcRole = temp.VcRole
	return nil
}

// tempVirtualChassisConfigMember is a temporary struct used for validating the fields of VirtualChassisConfigMember.
type tempVirtualChassisConfigMember struct {
	Locating *bool                                 `json:"locating,omitempty"`
	Mac      *string                               `json:"mac,omitempty"`
	Member   *int                                  `json:"member,omitempty"`
	VcPorts  []string                              `json:"vc_ports,omitempty"`
	VcRole   *VirtualChassisConfigMemberVcRoleEnum `json:"vc_role,omitempty"`
}
