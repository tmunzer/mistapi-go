// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"fmt"
)

// VirtualChassisConfig represents a VirtualChassisConfig struct.
// Virtual Chassis creation or configuration request
type VirtualChassisConfig struct {
	// Whether the Virtual Chassis is currently in locate mode
	Locating *bool `json:"locating,omitempty"`
	// Virtual Chassis member configurations
	Members []VirtualChassisConfigMember `json:"members,omitempty"`
	// Whether to create the Virtual Chassis in pre-provisioned mode
	Preprovisioned       *bool                  `json:"preprovisioned,omitempty"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for VirtualChassisConfig,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (v VirtualChassisConfig) String() string {
	return fmt.Sprintf(
		"VirtualChassisConfig[Locating=%v, Members=%v, Preprovisioned=%v, AdditionalProperties=%v]",
		v.Locating, v.Members, v.Preprovisioned, v.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for VirtualChassisConfig.
// It customizes the JSON marshaling process for VirtualChassisConfig objects.
func (v VirtualChassisConfig) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(v.AdditionalProperties,
		"locating", "members", "preprovisioned"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(v.toMap())
}

// toMap converts the VirtualChassisConfig object to a map representation for JSON marshaling.
func (v VirtualChassisConfig) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, v.AdditionalProperties)
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
	var temp tempVirtualChassisConfig
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "locating", "members", "preprovisioned")
	if err != nil {
		return err
	}
	v.AdditionalProperties = additionalProperties

	v.Locating = temp.Locating
	v.Members = temp.Members
	v.Preprovisioned = temp.Preprovisioned
	return nil
}

// tempVirtualChassisConfig is a temporary struct used for validating the fields of VirtualChassisConfig.
type tempVirtualChassisConfig struct {
	Locating       *bool                        `json:"locating,omitempty"`
	Members        []VirtualChassisConfigMember `json:"members,omitempty"`
	Preprovisioned *bool                        `json:"preprovisioned,omitempty"`
}
