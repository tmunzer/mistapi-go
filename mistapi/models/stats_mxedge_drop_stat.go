// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"fmt"
)

// StatsMxedgeDropStat represents a StatsMxedgeDropStat struct.
// Packet drop counters reported by the Mist Edge tunnel termination service. Counters not listed here may be reported as additional properties.
type StatsMxedgeDropStat struct {
	// Packets dropped because the IPv4 packet required fragmentation
	DropIp4MustFragment *int `json:"drop_ip4_must_fragment,omitempty"`
	// Packets dropped because no matching IPsec security association was found
	DropIpsecNoSa *int `json:"drop_ipsec_no_sa,omitempty"`
	// Packets dropped because the IPsec SPI was unknown
	DropIpsecUnknownSpi *int `json:"drop_ipsec_unknown_spi,omitempty"`
	// Packets dropped because no matching L2TP session was found
	DropL2tpNoSession *int `json:"drop_l2tp_no_session,omitempty"`
	// Packets dropped because the destination was a protected SVI
	DropProtectedSvi *int `json:"drop_protected_svi,omitempty"`
	// Packets dropped because the VLAN is not configured on the port
	DropVlanNotOnPort    *int                   `json:"drop_vlan_not_on_port,omitempty"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for StatsMxedgeDropStat,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s StatsMxedgeDropStat) String() string {
	return fmt.Sprintf(
		"StatsMxedgeDropStat[DropIp4MustFragment=%v, DropIpsecNoSa=%v, DropIpsecUnknownSpi=%v, DropL2tpNoSession=%v, DropProtectedSvi=%v, DropVlanNotOnPort=%v, AdditionalProperties=%v]",
		s.DropIp4MustFragment, s.DropIpsecNoSa, s.DropIpsecUnknownSpi, s.DropL2tpNoSession, s.DropProtectedSvi, s.DropVlanNotOnPort, s.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for StatsMxedgeDropStat.
// It customizes the JSON marshaling process for StatsMxedgeDropStat objects.
func (s StatsMxedgeDropStat) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(s.AdditionalProperties,
		"drop_ip4_must_fragment", "drop_ipsec_no_sa", "drop_ipsec_unknown_spi", "drop_l2tp_no_session", "drop_protected_svi", "drop_vlan_not_on_port"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(s.toMap())
}

// toMap converts the StatsMxedgeDropStat object to a map representation for JSON marshaling.
func (s StatsMxedgeDropStat) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, s.AdditionalProperties)
	if s.DropIp4MustFragment != nil {
		structMap["drop_ip4_must_fragment"] = s.DropIp4MustFragment
	}
	if s.DropIpsecNoSa != nil {
		structMap["drop_ipsec_no_sa"] = s.DropIpsecNoSa
	}
	if s.DropIpsecUnknownSpi != nil {
		structMap["drop_ipsec_unknown_spi"] = s.DropIpsecUnknownSpi
	}
	if s.DropL2tpNoSession != nil {
		structMap["drop_l2tp_no_session"] = s.DropL2tpNoSession
	}
	if s.DropProtectedSvi != nil {
		structMap["drop_protected_svi"] = s.DropProtectedSvi
	}
	if s.DropVlanNotOnPort != nil {
		structMap["drop_vlan_not_on_port"] = s.DropVlanNotOnPort
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for StatsMxedgeDropStat.
// It customizes the JSON unmarshaling process for StatsMxedgeDropStat objects.
func (s *StatsMxedgeDropStat) UnmarshalJSON(input []byte) error {
	var temp tempStatsMxedgeDropStat
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "drop_ip4_must_fragment", "drop_ipsec_no_sa", "drop_ipsec_unknown_spi", "drop_l2tp_no_session", "drop_protected_svi", "drop_vlan_not_on_port")
	if err != nil {
		return err
	}
	s.AdditionalProperties = additionalProperties

	s.DropIp4MustFragment = temp.DropIp4MustFragment
	s.DropIpsecNoSa = temp.DropIpsecNoSa
	s.DropIpsecUnknownSpi = temp.DropIpsecUnknownSpi
	s.DropL2tpNoSession = temp.DropL2tpNoSession
	s.DropProtectedSvi = temp.DropProtectedSvi
	s.DropVlanNotOnPort = temp.DropVlanNotOnPort
	return nil
}

// tempStatsMxedgeDropStat is a temporary struct used for validating the fields of StatsMxedgeDropStat.
type tempStatsMxedgeDropStat struct {
	DropIp4MustFragment *int `json:"drop_ip4_must_fragment,omitempty"`
	DropIpsecNoSa       *int `json:"drop_ipsec_no_sa,omitempty"`
	DropIpsecUnknownSpi *int `json:"drop_ipsec_unknown_spi,omitempty"`
	DropL2tpNoSession   *int `json:"drop_l2tp_no_session,omitempty"`
	DropProtectedSvi    *int `json:"drop_protected_svi,omitempty"`
	DropVlanNotOnPort   *int `json:"drop_vlan_not_on_port,omitempty"`
}
