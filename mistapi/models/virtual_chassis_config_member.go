package models

import (
    "encoding/json"
    "errors"
    "fmt"
    "strings"
)

// VirtualChassisConfigMember represents a VirtualChassisConfigMember struct.
type VirtualChassisConfigMember struct {
    Locating             *bool                                `json:"locating,omitempty"`
    // fpc0, same as the mac of device_id
    Mac                  string                               `json:"mac"`
    // For preprovisionned virtual chassis
    MemberId             *int                                 `json:"member_id,omitempty"`
    VcPorts              []string                             `json:"vc_ports,omitempty"`
    // enum: `backup`, `linecard`, `master`
    VcRole               VirtualChassisConfigMemberVcRoleEnum `json:"vc_role"`
    AdditionalProperties map[string]interface{}               `json:"_"`
}

// String implements the fmt.Stringer interface for VirtualChassisConfigMember,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (v VirtualChassisConfigMember) String() string {
    return fmt.Sprintf(
    	"VirtualChassisConfigMember[Locating=%v, Mac=%v, MemberId=%v, VcPorts=%v, VcRole=%v, AdditionalProperties=%v]",
    	v.Locating, v.Mac, v.MemberId, v.VcPorts, v.VcRole, v.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for VirtualChassisConfigMember.
// It customizes the JSON marshaling process for VirtualChassisConfigMember objects.
func (v VirtualChassisConfigMember) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(v.AdditionalProperties,
        "locating", "mac", "member_id", "vc_ports", "vc_role"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(v.toMap())
}

// toMap converts the VirtualChassisConfigMember object to a map representation for JSON marshaling.
func (v VirtualChassisConfigMember) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, v.AdditionalProperties)
    if v.Locating != nil {
        structMap["locating"] = v.Locating
    }
    structMap["mac"] = v.Mac
    if v.MemberId != nil {
        structMap["member_id"] = v.MemberId
    }
    if v.VcPorts != nil {
        structMap["vc_ports"] = v.VcPorts
    }
    structMap["vc_role"] = v.VcRole
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
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "locating", "mac", "member_id", "vc_ports", "vc_role")
    if err != nil {
    	return err
    }
    v.AdditionalProperties = additionalProperties
    
    v.Locating = temp.Locating
    v.Mac = *temp.Mac
    v.MemberId = temp.MemberId
    v.VcPorts = temp.VcPorts
    v.VcRole = *temp.VcRole
    return nil
}

// tempVirtualChassisConfigMember is a temporary struct used for validating the fields of VirtualChassisConfigMember.
type tempVirtualChassisConfigMember  struct {
    Locating *bool                                 `json:"locating,omitempty"`
    Mac      *string                               `json:"mac"`
    MemberId *int                                  `json:"member_id,omitempty"`
    VcPorts  []string                              `json:"vc_ports,omitempty"`
    VcRole   *VirtualChassisConfigMemberVcRoleEnum `json:"vc_role"`
}

func (v *tempVirtualChassisConfigMember) validate() error {
    var errs []string
    if v.Mac == nil {
        errs = append(errs, "required field `mac` is missing for type `virtual_chassis_config_member`")
    }
    if v.VcRole == nil {
        errs = append(errs, "required field `vc_role` is missing for type `virtual_chassis_config_member`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
