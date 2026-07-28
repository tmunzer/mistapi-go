// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"fmt"
)

// SwitchMulticastConfig represents a SwitchMulticastConfig struct.
// Multicast configuration for a VRF. When set at the network template level it applies to networks in the master VRF (not assigned to any vrf_instances). PIM is automatically enabled when any network in the VRF has `multicast.enabled`==`true`.
type SwitchMulticastConfig struct {
	// When `true`, auto-generates a shared RP on `is_l3_border` devices (ERB/IPClos topologies only)
	AnycastRp *bool `json:"anycast_rp,omitempty"`
	// RP address used when `anycast_rp`==`false`. If the address matches a device SVI, it is configured as a local RP; otherwise a static RP is configured
	RpIp *string `json:"rp_ip,omitempty"`
	// SBD IRB subnet; Mist auto-assigns per-device IPs from this range (EVPN eOISM only)
	SbdSubnet *string `json:"sbd_subnet,omitempty"`
	// Supplemental Bridge Domain VLAN ID (EVPN topology / eOISM only)
	SbdVlanId            *int                   `json:"sbd_vlan_id,omitempty"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for SwitchMulticastConfig,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s SwitchMulticastConfig) String() string {
	return fmt.Sprintf(
		"SwitchMulticastConfig[AnycastRp=%v, RpIp=%v, SbdSubnet=%v, SbdVlanId=%v, AdditionalProperties=%v]",
		s.AnycastRp, s.RpIp, s.SbdSubnet, s.SbdVlanId, s.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for SwitchMulticastConfig.
// It customizes the JSON marshaling process for SwitchMulticastConfig objects.
func (s SwitchMulticastConfig) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(s.AdditionalProperties,
		"anycast_rp", "rp_ip", "sbd_subnet", "sbd_vlan_id"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(s.toMap())
}

// toMap converts the SwitchMulticastConfig object to a map representation for JSON marshaling.
func (s SwitchMulticastConfig) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, s.AdditionalProperties)
	if s.AnycastRp != nil {
		structMap["anycast_rp"] = s.AnycastRp
	}
	if s.RpIp != nil {
		structMap["rp_ip"] = s.RpIp
	}
	if s.SbdSubnet != nil {
		structMap["sbd_subnet"] = s.SbdSubnet
	}
	if s.SbdVlanId != nil {
		structMap["sbd_vlan_id"] = s.SbdVlanId
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SwitchMulticastConfig.
// It customizes the JSON unmarshaling process for SwitchMulticastConfig objects.
func (s *SwitchMulticastConfig) UnmarshalJSON(input []byte) error {
	var temp tempSwitchMulticastConfig
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "anycast_rp", "rp_ip", "sbd_subnet", "sbd_vlan_id")
	if err != nil {
		return err
	}
	s.AdditionalProperties = additionalProperties

	s.AnycastRp = temp.AnycastRp
	s.RpIp = temp.RpIp
	s.SbdSubnet = temp.SbdSubnet
	s.SbdVlanId = temp.SbdVlanId
	return nil
}

// tempSwitchMulticastConfig is a temporary struct used for validating the fields of SwitchMulticastConfig.
type tempSwitchMulticastConfig struct {
	AnycastRp *bool   `json:"anycast_rp,omitempty"`
	RpIp      *string `json:"rp_ip,omitempty"`
	SbdSubnet *string `json:"sbd_subnet,omitempty"`
	SbdVlanId *int    `json:"sbd_vlan_id,omitempty"`
}
