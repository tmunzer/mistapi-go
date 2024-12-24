package models

import (
    "encoding/json"
    "fmt"
)

// SwitchVirtualChassisMember represents a SwitchVirtualChassisMember struct.
type SwitchVirtualChassisMember struct {
    // fpc0, same as the mac of device_id
    Mac                  *string                               `json:"mac,omitempty"`
    MemberId             *int                                  `json:"member_id,omitempty"`
    // Both vc_role master and backup will be matched to routing-engine role in Junos preprovisioned VC config. enum: `backup`, `linecard`, `master`
    VcRole               *SwitchVirtualChassisMemberVcRoleEnum `json:"vc_role,omitempty"`
    AdditionalProperties map[string]interface{}                `json:"_"`
}

// String implements the fmt.Stringer interface for SwitchVirtualChassisMember,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s SwitchVirtualChassisMember) String() string {
    return fmt.Sprintf(
    	"SwitchVirtualChassisMember[Mac=%v, MemberId=%v, VcRole=%v, AdditionalProperties=%v]",
    	s.Mac, s.MemberId, s.VcRole, s.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for SwitchVirtualChassisMember.
// It customizes the JSON marshaling process for SwitchVirtualChassisMember objects.
func (s SwitchVirtualChassisMember) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(s.AdditionalProperties,
        "mac", "member_id", "vc_role"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(s.toMap())
}

// toMap converts the SwitchVirtualChassisMember object to a map representation for JSON marshaling.
func (s SwitchVirtualChassisMember) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, s.AdditionalProperties)
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
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "mac", "member_id", "vc_role")
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
type tempSwitchVirtualChassisMember  struct {
    Mac      *string                               `json:"mac,omitempty"`
    MemberId *int                                  `json:"member_id,omitempty"`
    VcRole   *SwitchVirtualChassisMemberVcRoleEnum `json:"vc_role,omitempty"`
}
