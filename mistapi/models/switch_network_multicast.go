// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"fmt"
)

// SwitchNetworkMulticast represents a SwitchNetworkMulticast struct.
// Multicast settings for a switch network (VLAN)
type SwitchNetworkMulticast struct {
	// Whether to enable IGMP snooping on this VLAN
	Enabled *bool `json:"enabled,omitempty"`
	// IGMP version. '2' (default, ASM/IGMPv2) / '3' (SSM/IGMPv3)
	IgmpVersion          *IgmpVersionEnum       `json:"igmp_version,omitempty"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for SwitchNetworkMulticast,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s SwitchNetworkMulticast) String() string {
	return fmt.Sprintf(
		"SwitchNetworkMulticast[Enabled=%v, IgmpVersion=%v, AdditionalProperties=%v]",
		s.Enabled, s.IgmpVersion, s.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for SwitchNetworkMulticast.
// It customizes the JSON marshaling process for SwitchNetworkMulticast objects.
func (s SwitchNetworkMulticast) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(s.AdditionalProperties,
		"enabled", "igmp_version"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(s.toMap())
}

// toMap converts the SwitchNetworkMulticast object to a map representation for JSON marshaling.
func (s SwitchNetworkMulticast) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, s.AdditionalProperties)
	if s.Enabled != nil {
		structMap["enabled"] = s.Enabled
	}
	if s.IgmpVersion != nil {
		structMap["igmp_version"] = s.IgmpVersion
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SwitchNetworkMulticast.
// It customizes the JSON unmarshaling process for SwitchNetworkMulticast objects.
func (s *SwitchNetworkMulticast) UnmarshalJSON(input []byte) error {
	var temp tempSwitchNetworkMulticast
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "enabled", "igmp_version")
	if err != nil {
		return err
	}
	s.AdditionalProperties = additionalProperties

	s.Enabled = temp.Enabled
	s.IgmpVersion = temp.IgmpVersion
	return nil
}

// tempSwitchNetworkMulticast is a temporary struct used for validating the fields of SwitchNetworkMulticast.
type tempSwitchNetworkMulticast struct {
	Enabled     *bool            `json:"enabled,omitempty"`
	IgmpVersion *IgmpVersionEnum `json:"igmp_version,omitempty"`
}
