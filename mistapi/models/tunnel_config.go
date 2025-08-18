// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"fmt"
)

// TunnelConfig represents a TunnelConfig struct.
type TunnelConfig struct {
	// Auto Provisioning configuration for the tunne. This takes precedence over the `primary` and `secondary` nodes.
	AutoProvision *TunnelConfigAutoProvision `json:"auto_provision,omitempty"`
	// Only if `provider`==`custom-ipsec`
	IkeLifetime *int `json:"ike_lifetime,omitempty"`
	// Only if `provider`==`custom-ipsec`. enum: `aggressive`, `main`
	IkeMode *TunnelConfigIkeModeEnum `json:"ike_mode,omitempty"`
	// If `provider`==`custom-ipsec`
	IkeProposals []TunnelConfigIkeProposal `json:"ike_proposals,omitempty"`
	// If `provider`==`custom-ipsec`
	IpsecLifetime *int `json:"ipsec_lifetime,omitempty"`
	// Only if  `provider`==`custom-ipsec`
	IpsecProposals []TunnelConfigIpsecProposal `json:"ipsec_proposals,omitempty"`
	// Required if `provider`==`zscaler-ipsec`, `provider`==`jse-ipsec` or `provider`==`custom-ipsec`
	LocalId *string `json:"local_id,omitempty"`
	// List of Local protected subnet for policy-based IPSec negotiation
	LocalSubnets []string `json:"local_subnets,omitempty"`
	// Required if `provider`==`zscaler-gre`, `provider`==`jse-ipsec`. enum: `active-active`, `active-standby`
	Mode *TunnelConfigTunnelModeEnum `json:"mode,omitempty"`
	// If `provider`==`custom-ipsec` or `provider`==`prisma-ipsec`, networks reachable via this tunnel
	Networks []string `json:"networks,omitempty"`
	// Only if `provider`==`zscaler-ipsec`, `provider`==`jse-ipsec` or `provider`==`custom-ipsec`
	Primary *TunnelConfigNode `json:"primary,omitempty"`
	// Only if `provider`==`custom-ipsec`
	Probe *TunnelConfigProbe `json:"probe,omitempty"`
	// Only if `provider`==`custom-ipsec`. enum: `gre`, `ipsec`
	Protocol *TunnelConfigProtocolEnum `json:"protocol,omitempty"`
	// Only if `auto_provision.enabled`==`false`. enum: `custom-ipsec`, `custom-gre`, `jse-ipsec`, `prisma-ipsec`, `zscaler-gre`, `zscaler-ipsec`
	Provider *TunnelConfigProviderEnum `json:"provider,omitempty"`
	// Required if `provider`==`zscaler-ipsec`, `provider`==`jse-ipsec` or `provider`==`custom-ipsec`
	Psk *string `json:"psk,omitempty"`
	// List of Remote protected subnet for policy-based IPSec negotiation
	RemoteSubnets []string `json:"remote_subnets,omitempty"`
	// Only if `provider`==`zscaler-ipsec`, `provider`==`jse-ipsec` or `provider`==`custom-ipsec`
	Secondary *TunnelConfigNode `json:"secondary,omitempty"`
	// Only if `provider`==`custom-gre` or `provider`==`custom-ipsec`. enum: `1`, `2`
	Version              *TunnelConfigVersionEnum `json:"version,omitempty"`
	AdditionalProperties map[string]interface{}   `json:"_"`
}

// String implements the fmt.Stringer interface for TunnelConfig,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (t TunnelConfig) String() string {
	return fmt.Sprintf(
		"TunnelConfig[AutoProvision=%v, IkeLifetime=%v, IkeMode=%v, IkeProposals=%v, IpsecLifetime=%v, IpsecProposals=%v, LocalId=%v, LocalSubnets=%v, Mode=%v, Networks=%v, Primary=%v, Probe=%v, Protocol=%v, Provider=%v, Psk=%v, RemoteSubnets=%v, Secondary=%v, Version=%v, AdditionalProperties=%v]",
		t.AutoProvision, t.IkeLifetime, t.IkeMode, t.IkeProposals, t.IpsecLifetime, t.IpsecProposals, t.LocalId, t.LocalSubnets, t.Mode, t.Networks, t.Primary, t.Probe, t.Protocol, t.Provider, t.Psk, t.RemoteSubnets, t.Secondary, t.Version, t.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for TunnelConfig.
// It customizes the JSON marshaling process for TunnelConfig objects.
func (t TunnelConfig) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(t.AdditionalProperties,
		"auto_provision", "ike_lifetime", "ike_mode", "ike_proposals", "ipsec_lifetime", "ipsec_proposals", "local_id", "local_subnets", "mode", "networks", "primary", "probe", "protocol", "provider", "psk", "remote_subnets", "secondary", "version"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(t.toMap())
}

// toMap converts the TunnelConfig object to a map representation for JSON marshaling.
func (t TunnelConfig) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, t.AdditionalProperties)
	if t.AutoProvision != nil {
		structMap["auto_provision"] = t.AutoProvision.toMap()
	}
	if t.IkeLifetime != nil {
		structMap["ike_lifetime"] = t.IkeLifetime
	}
	if t.IkeMode != nil {
		structMap["ike_mode"] = t.IkeMode
	}
	if t.IkeProposals != nil {
		structMap["ike_proposals"] = t.IkeProposals
	}
	if t.IpsecLifetime != nil {
		structMap["ipsec_lifetime"] = t.IpsecLifetime
	}
	if t.IpsecProposals != nil {
		structMap["ipsec_proposals"] = t.IpsecProposals
	}
	if t.LocalId != nil {
		structMap["local_id"] = t.LocalId
	}
	if t.LocalSubnets != nil {
		structMap["local_subnets"] = t.LocalSubnets
	}
	if t.Mode != nil {
		structMap["mode"] = t.Mode
	}
	if t.Networks != nil {
		structMap["networks"] = t.Networks
	}
	if t.Primary != nil {
		structMap["primary"] = t.Primary.toMap()
	}
	if t.Probe != nil {
		structMap["probe"] = t.Probe.toMap()
	}
	if t.Protocol != nil {
		structMap["protocol"] = t.Protocol
	}
	if t.Provider != nil {
		structMap["provider"] = t.Provider
	}
	if t.Psk != nil {
		structMap["psk"] = t.Psk
	}
	if t.RemoteSubnets != nil {
		structMap["remote_subnets"] = t.RemoteSubnets
	}
	if t.Secondary != nil {
		structMap["secondary"] = t.Secondary.toMap()
	}
	if t.Version != nil {
		structMap["version"] = t.Version
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for TunnelConfig.
// It customizes the JSON unmarshaling process for TunnelConfig objects.
func (t *TunnelConfig) UnmarshalJSON(input []byte) error {
	var temp tempTunnelConfig
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "auto_provision", "ike_lifetime", "ike_mode", "ike_proposals", "ipsec_lifetime", "ipsec_proposals", "local_id", "local_subnets", "mode", "networks", "primary", "probe", "protocol", "provider", "psk", "remote_subnets", "secondary", "version")
	if err != nil {
		return err
	}
	t.AdditionalProperties = additionalProperties

	t.AutoProvision = temp.AutoProvision
	t.IkeLifetime = temp.IkeLifetime
	t.IkeMode = temp.IkeMode
	t.IkeProposals = temp.IkeProposals
	t.IpsecLifetime = temp.IpsecLifetime
	t.IpsecProposals = temp.IpsecProposals
	t.LocalId = temp.LocalId
	t.LocalSubnets = temp.LocalSubnets
	t.Mode = temp.Mode
	t.Networks = temp.Networks
	t.Primary = temp.Primary
	t.Probe = temp.Probe
	t.Protocol = temp.Protocol
	t.Provider = temp.Provider
	t.Psk = temp.Psk
	t.RemoteSubnets = temp.RemoteSubnets
	t.Secondary = temp.Secondary
	t.Version = temp.Version
	return nil
}

// tempTunnelConfig is a temporary struct used for validating the fields of TunnelConfig.
type tempTunnelConfig struct {
	AutoProvision  *TunnelConfigAutoProvision  `json:"auto_provision,omitempty"`
	IkeLifetime    *int                        `json:"ike_lifetime,omitempty"`
	IkeMode        *TunnelConfigIkeModeEnum    `json:"ike_mode,omitempty"`
	IkeProposals   []TunnelConfigIkeProposal   `json:"ike_proposals,omitempty"`
	IpsecLifetime  *int                        `json:"ipsec_lifetime,omitempty"`
	IpsecProposals []TunnelConfigIpsecProposal `json:"ipsec_proposals,omitempty"`
	LocalId        *string                     `json:"local_id,omitempty"`
	LocalSubnets   []string                    `json:"local_subnets,omitempty"`
	Mode           *TunnelConfigTunnelModeEnum `json:"mode,omitempty"`
	Networks       []string                    `json:"networks,omitempty"`
	Primary        *TunnelConfigNode           `json:"primary,omitempty"`
	Probe          *TunnelConfigProbe          `json:"probe,omitempty"`
	Protocol       *TunnelConfigProtocolEnum   `json:"protocol,omitempty"`
	Provider       *TunnelConfigProviderEnum   `json:"provider,omitempty"`
	Psk            *string                     `json:"psk,omitempty"`
	RemoteSubnets  []string                    `json:"remote_subnets,omitempty"`
	Secondary      *TunnelConfigNode           `json:"secondary,omitempty"`
	Version        *TunnelConfigVersionEnum    `json:"version,omitempty"`
}
