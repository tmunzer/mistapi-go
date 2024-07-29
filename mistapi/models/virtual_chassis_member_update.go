package models

import (
    "encoding/json"
)

// VirtualChassisMemberUpdate represents a VirtualChassisMemberUpdate struct.
type VirtualChassisMemberUpdate struct {
    // Required if `op`==`add` or `op`==`preprovision`.
    Mac                  *string                               `json:"mac,omitempty"`
    // Required if `op`==`remove` or `op`==`preprovision`. Optional if `op`==`add`
    Member               *int                                  `json:"member,omitempty"`
    // Required if `op`==`add` or `op`==`preprovision`
    VcPorts              []string                              `json:"vc_ports,omitempty"`
    // Required if `op`==`add` or `op`==`preprovision`. enum: `backup`, `linecard`, `master`
    VcRole               *VirtualChassisMemberUpdateVcRoleEnum `json:"vc_role,omitempty"`
    AdditionalProperties map[string]any                        `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for VirtualChassisMemberUpdate.
// It customizes the JSON marshaling process for VirtualChassisMemberUpdate objects.
func (v VirtualChassisMemberUpdate) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(v.toMap())
}

// toMap converts the VirtualChassisMemberUpdate object to a map representation for JSON marshaling.
func (v VirtualChassisMemberUpdate) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, v.AdditionalProperties)
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

// UnmarshalJSON implements the json.Unmarshaler interface for VirtualChassisMemberUpdate.
// It customizes the JSON unmarshaling process for VirtualChassisMemberUpdate objects.
func (v *VirtualChassisMemberUpdate) UnmarshalJSON(input []byte) error {
    var temp virtualChassisMemberUpdate
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "mac", "member", "vc_ports", "vc_role")
    if err != nil {
    	return err
    }
    
    v.AdditionalProperties = additionalProperties
    v.Mac = temp.Mac
    v.Member = temp.Member
    v.VcPorts = temp.VcPorts
    v.VcRole = temp.VcRole
    return nil
}

// virtualChassisMemberUpdate is a temporary struct used for validating the fields of VirtualChassisMemberUpdate.
type virtualChassisMemberUpdate  struct {
    Mac     *string                               `json:"mac,omitempty"`
    Member  *int                                  `json:"member,omitempty"`
    VcPorts []string                              `json:"vc_ports,omitempty"`
    VcRole  *VirtualChassisMemberUpdateVcRoleEnum `json:"vc_role,omitempty"`
}
