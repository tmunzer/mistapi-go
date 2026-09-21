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
	// When `true`, generates a shared anycast RP on all `is_l3_border` devices in EVPN (ERB/IPClos) topologies. Uses `rp_ip` as the shared RP address, or an internal default when `rp_ip` is omitted. Takes precedence over `rp_mac` and `rp_ip` when multiple RP options are set.
	AnycastRp *bool `json:"anycast_rp,omitempty"`
	// When `true`, enables the PIM EVPN Gateway on `is_l3_border` devices. Required for external sources or receivers in EVPN topologies.
	PegEnabled *bool `json:"peg_enabled,omitempty"`
	// RP address used for EVPN anycast RP when `anycast_rp` is true, or for an external RP when it is false. In non-EVPN topologies, a matching device router ID configures a local RP; otherwise a static RP is configured.
	RpIp *string `json:"rp_ip,omitempty"`
	// Device MAC address of a fabric RP in EVPN topologies. The RP address is the first usable IP of the VRF `evpn_auto_loopback_subnet`, not `rp_ip`; requires `evpn_auto_loopback_subnet`. Takes precedence over `rp_ip` when `anycast_rp` is false.
	RpMac *string `json:"rp_mac,omitempty"`
	// SBD IRB subnet; Mist auto-assigns per-device IPs from this range (EVPN eOISM only)
	SbdSubnet *string `json:"sbd_subnet,omitempty"`
	// Supplemental Bridge Domain VLAN ID (EVPN topology / eOISM only)
	SbdVlanId *int `json:"sbd_vlan_id,omitempty"`
	// When `true` on PEG borders, builds an eBGP mesh between PEG borders over SBD IRBs so WAN-learned routes can satisfy the PIM RPF check during a border WAN-uplink failure.
	SbdWanRpf            *bool                  `json:"sbd_wan_rpf,omitempty"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for SwitchMulticastConfig,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s SwitchMulticastConfig) String() string {
	return fmt.Sprintf(
		"SwitchMulticastConfig[AnycastRp=%v, PegEnabled=%v, RpIp=%v, RpMac=%v, SbdSubnet=%v, SbdVlanId=%v, SbdWanRpf=%v, AdditionalProperties=%v]",
		s.AnycastRp, s.PegEnabled, s.RpIp, s.RpMac, s.SbdSubnet, s.SbdVlanId, s.SbdWanRpf, s.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for SwitchMulticastConfig.
// It customizes the JSON marshaling process for SwitchMulticastConfig objects.
func (s SwitchMulticastConfig) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(s.AdditionalProperties,
		"anycast_rp", "peg_enabled", "rp_ip", "rp_mac", "sbd_subnet", "sbd_vlan_id", "sbd_wan_rpf"); err != nil {
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
	if s.PegEnabled != nil {
		structMap["peg_enabled"] = s.PegEnabled
	}
	if s.RpIp != nil {
		structMap["rp_ip"] = s.RpIp
	}
	if s.RpMac != nil {
		structMap["rp_mac"] = s.RpMac
	}
	if s.SbdSubnet != nil {
		structMap["sbd_subnet"] = s.SbdSubnet
	}
	if s.SbdVlanId != nil {
		structMap["sbd_vlan_id"] = s.SbdVlanId
	}
	if s.SbdWanRpf != nil {
		structMap["sbd_wan_rpf"] = s.SbdWanRpf
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
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "anycast_rp", "peg_enabled", "rp_ip", "rp_mac", "sbd_subnet", "sbd_vlan_id", "sbd_wan_rpf")
	if err != nil {
		return err
	}
	s.AdditionalProperties = additionalProperties

	s.AnycastRp = temp.AnycastRp
	s.PegEnabled = temp.PegEnabled
	s.RpIp = temp.RpIp
	s.RpMac = temp.RpMac
	s.SbdSubnet = temp.SbdSubnet
	s.SbdVlanId = temp.SbdVlanId
	s.SbdWanRpf = temp.SbdWanRpf
	return nil
}

// tempSwitchMulticastConfig is a temporary struct used for validating the fields of SwitchMulticastConfig.
type tempSwitchMulticastConfig struct {
	AnycastRp  *bool   `json:"anycast_rp,omitempty"`
	PegEnabled *bool   `json:"peg_enabled,omitempty"`
	RpIp       *string `json:"rp_ip,omitempty"`
	RpMac      *string `json:"rp_mac,omitempty"`
	SbdSubnet  *string `json:"sbd_subnet,omitempty"`
	SbdVlanId  *int    `json:"sbd_vlan_id,omitempty"`
	SbdWanRpf  *bool   `json:"sbd_wan_rpf,omitempty"`
}
