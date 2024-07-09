package models

import (
    "encoding/json"
)

// VpnPath represents a VpnPath struct.
type VpnPath struct {
    BfdProfile           *VpnPathBfdProfileEnum `json:"bfd_profile,omitempty"`
    // if different from the wan port
    Ip                   *string                `json:"ip,omitempty"`
    Pod                  *int                   `json:"pod,omitempty"`
    AdditionalProperties map[string]any         `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for VpnPath.
// It customizes the JSON marshaling process for VpnPath objects.
func (v VpnPath) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(v.toMap())
}

// toMap converts the VpnPath object to a map representation for JSON marshaling.
func (v VpnPath) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, v.AdditionalProperties)
    if v.BfdProfile != nil {
        structMap["bfd_profile"] = v.BfdProfile
    }
    if v.Ip != nil {
        structMap["ip"] = v.Ip
    }
    if v.Pod != nil {
        structMap["pod"] = v.Pod
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for VpnPath.
// It customizes the JSON unmarshaling process for VpnPath objects.
func (v *VpnPath) UnmarshalJSON(input []byte) error {
    var temp vpnPath
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "bfd_profile", "ip", "pod")
    if err != nil {
    	return err
    }
    
    v.AdditionalProperties = additionalProperties
    v.BfdProfile = temp.BfdProfile
    v.Ip = temp.Ip
    v.Pod = temp.Pod
    return nil
}

// vpnPath is a temporary struct used for validating the fields of VpnPath.
type vpnPath  struct {
    BfdProfile *VpnPathBfdProfileEnum `json:"bfd_profile,omitempty"`
    Ip         *string                `json:"ip,omitempty"`
    Pod        *int                   `json:"pod,omitempty"`
}
