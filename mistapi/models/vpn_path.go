package models

import (
    "encoding/json"
    "fmt"
)

// VpnPath represents a VpnPath struct.
type VpnPath struct {
    // enum: `broadband`, `lte`
    BfdProfile           *VpnPathBfdProfileEnum `json:"bfd_profile,omitempty"`
    // If different from the wan port
    Ip                   *string                `json:"ip,omitempty"`
    Pod                  *int                   `json:"pod,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for VpnPath,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (v VpnPath) String() string {
    return fmt.Sprintf(
    	"VpnPath[BfdProfile=%v, Ip=%v, Pod=%v, AdditionalProperties=%v]",
    	v.BfdProfile, v.Ip, v.Pod, v.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for VpnPath.
// It customizes the JSON marshaling process for VpnPath objects.
func (v VpnPath) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(v.AdditionalProperties,
        "bfd_profile", "ip", "pod"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(v.toMap())
}

// toMap converts the VpnPath object to a map representation for JSON marshaling.
func (v VpnPath) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, v.AdditionalProperties)
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
    var temp tempVpnPath
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "bfd_profile", "ip", "pod")
    if err != nil {
    	return err
    }
    v.AdditionalProperties = additionalProperties
    
    v.BfdProfile = temp.BfdProfile
    v.Ip = temp.Ip
    v.Pod = temp.Pod
    return nil
}

// tempVpnPath is a temporary struct used for validating the fields of VpnPath.
type tempVpnPath  struct {
    BfdProfile *VpnPathBfdProfileEnum `json:"bfd_profile,omitempty"`
    Ip         *string                `json:"ip,omitempty"`
    Pod        *int                   `json:"pod,omitempty"`
}
