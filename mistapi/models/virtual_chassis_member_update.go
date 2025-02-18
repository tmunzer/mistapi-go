package models

import (
    "encoding/json"
    "fmt"
)

// VirtualChassisMemberUpdate represents a VirtualChassisMemberUpdate struct.
type VirtualChassisMemberUpdate struct {
    // Required if `op`==`add` or `op`==`preprovision`.
    Mac                  *string                               `json:"mac,omitempty"`
    // Required if `op`==`remove`
    Member               *int                                  `json:"member,omitempty"`
    // Required if `op`==`preprovision`. Optional if `op`==`add`
    MemberId             *int                                  `json:"member_id,omitempty"`
    // Required if `op`==`add` or `op`==`preprovision`
    VcPorts              []string                              `json:"vc_ports,omitempty"`
    // Required if `op`==`add` or `op`==`preprovision`. enum: `backup`, `linecard`, `master`
    VcRole               *VirtualChassisMemberUpdateVcRoleEnum `json:"vc_role,omitempty"`
    AdditionalProperties map[string]interface{}                `json:"_"`
}

// String implements the fmt.Stringer interface for VirtualChassisMemberUpdate,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (v VirtualChassisMemberUpdate) String() string {
    return fmt.Sprintf(
    	"VirtualChassisMemberUpdate[Mac=%v, Member=%v, MemberId=%v, VcPorts=%v, VcRole=%v, AdditionalProperties=%v]",
    	v.Mac, v.Member, v.MemberId, v.VcPorts, v.VcRole, v.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for VirtualChassisMemberUpdate.
// It customizes the JSON marshaling process for VirtualChassisMemberUpdate objects.
func (v VirtualChassisMemberUpdate) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(v.AdditionalProperties,
        "mac", "member", "member_id", "vc_ports", "vc_role"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(v.toMap())
}

// toMap converts the VirtualChassisMemberUpdate object to a map representation for JSON marshaling.
func (v VirtualChassisMemberUpdate) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, v.AdditionalProperties)
    if v.Mac != nil {
        structMap["mac"] = v.Mac
    }
    if v.Member != nil {
        structMap["member"] = v.Member
    }
    if v.MemberId != nil {
        structMap["member_id"] = v.MemberId
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
    var temp tempVirtualChassisMemberUpdate
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "mac", "member", "member_id", "vc_ports", "vc_role")
    if err != nil {
    	return err
    }
    v.AdditionalProperties = additionalProperties
    
    v.Mac = temp.Mac
    v.Member = temp.Member
    v.MemberId = temp.MemberId
    v.VcPorts = temp.VcPorts
    v.VcRole = temp.VcRole
    return nil
}

// tempVirtualChassisMemberUpdate is a temporary struct used for validating the fields of VirtualChassisMemberUpdate.
type tempVirtualChassisMemberUpdate  struct {
    Mac      *string                               `json:"mac,omitempty"`
    Member   *int                                  `json:"member,omitempty"`
    MemberId *int                                  `json:"member_id,omitempty"`
    VcPorts  []string                              `json:"vc_ports,omitempty"`
    VcRole   *VirtualChassisMemberUpdateVcRoleEnum `json:"vc_role,omitempty"`
}
