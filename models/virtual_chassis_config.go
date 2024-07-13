package models

import (
    "encoding/json"
)

// VirtualChassisConfig represents a VirtualChassisConfig struct.
// Virtual Chassis
type VirtualChassisConfig struct {
    Locating             *bool                        `json:"locating,omitempty"`
    Members              []VirtualChassisConfigMember `json:"members,omitempty"`
    // To create the Virtual Chassis in Pre-Provisioned mode
    Preprovisioned       *bool                        `json:"preprovisioned,omitempty"`
    AdditionalProperties map[string]any               `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for VirtualChassisConfig.
// It customizes the JSON marshaling process for VirtualChassisConfig objects.
func (v VirtualChassisConfig) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(v.toMap())
}

// toMap converts the VirtualChassisConfig object to a map representation for JSON marshaling.
func (v VirtualChassisConfig) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, v.AdditionalProperties)
    if v.Locating != nil {
        structMap["locating"] = v.Locating
    }
    if v.Members != nil {
        structMap["members"] = v.Members
    }
    if v.Preprovisioned != nil {
        structMap["preprovisioned"] = v.Preprovisioned
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for VirtualChassisConfig.
// It customizes the JSON unmarshaling process for VirtualChassisConfig objects.
func (v *VirtualChassisConfig) UnmarshalJSON(input []byte) error {
    var temp virtualChassisConfig
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "locating", "members", "preprovisioned")
    if err != nil {
    	return err
    }
    
    v.AdditionalProperties = additionalProperties
    v.Locating = temp.Locating
    v.Members = temp.Members
    v.Preprovisioned = temp.Preprovisioned
    return nil
}

// virtualChassisConfig is a temporary struct used for validating the fields of VirtualChassisConfig.
type virtualChassisConfig  struct {
    Locating       *bool                        `json:"locating,omitempty"`
    Members        []VirtualChassisConfigMember `json:"members,omitempty"`
    Preprovisioned *bool                        `json:"preprovisioned,omitempty"`
}