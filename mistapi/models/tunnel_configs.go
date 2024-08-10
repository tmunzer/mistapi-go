package models

import (
    "encoding/json"
)

// TunnelConfigs represents a TunnelConfigs struct.
type TunnelConfigs struct {
    AutoProvision        *TunnelConfigsAutoProvision          `json:"auto_provision,omitempty"`
    // Only if `provider`== `custom-ipsec`
    IkeLifetime          *int                                 `json:"ike_lifetime,omitempty"`
    // Only if `provider`== `custom-ipsec`. enum: `aggressive`, `main`
    IkeMode              *GatewayTemplateTunnelIkeModeEnum    `json:"ike_mode,omitempty"`
    // if `provider`== `custom-ipsec`
    IkeProposals         []GatewayTemplateTunnelIkeProposal   `json:"ike_proposals,omitempty"`
    // if `provider`== `custom-ipsec`
    IpsecLifetime        *int                                 `json:"ipsec_lifetime,omitempty"`
    // Only if  `provider`== `custom-ipsec`
    IpsecProposals       []GatewayTemplateTunnelIpsecProposal `json:"ipsec_proposals,omitempty"`
    // Only if:
    //   * `provider`== `zscaler-ipsec`
    //   * `provider`==`jse-ipsec`
    //   * `provider`== `custom-ipsec`
    LocalId              *string                              `json:"local_id,omitempty"`
    // enum: `active-active`, `active-standby`
    Mode                 *GatewayTemplateTunnelModeEnum       `json:"mode,omitempty"`
    // networks reachable via this tunnel
    Networks             []string                             `json:"networks,omitempty"`
    Primary              *GatewayTemplateTunnelNode           `json:"primary,omitempty"`
    // Only if `provider`== `custom-ipsec`
    Probe                *GatewayTemplateTunnelProbe          `json:"probe,omitempty"`
    // Only if `provider`== `custom-ipsec`. enum: `gre`, `ipsec`
    Protocol             *GatewayTemplateTunnelProtocolEnum   `json:"protocol,omitempty"`
    // enum: `custom-ipsec`, `customer-gre`, `jse-ipsec`, `zscaler-gre`, `zscaler-ipsec`
    Provider             *TunnelProviderOptionsNameEnum       `json:"provider,omitempty"`
    // Only if:
    //   * `provider`== `zscaler-ipsec`
    //   * `provider`==`jse-ipsec`
    //   * `provider`== `custom-ipsec`
    Psk                  *string                              `json:"psk,omitempty"`
    Secondary            *GatewayTemplateTunnelNode           `json:"secondary,omitempty"`
    // Only if `provider`== `custom-gre` or `provider`== `custom-ipsec`. enum: `1`, `2`
    Version              *GatewayTemplateTunnelVersionEnum    `json:"version,omitempty"`
    AdditionalProperties map[string]any                       `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for TunnelConfigs.
// It customizes the JSON marshaling process for TunnelConfigs objects.
func (t TunnelConfigs) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(t.toMap())
}

// toMap converts the TunnelConfigs object to a map representation for JSON marshaling.
func (t TunnelConfigs) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, t.AdditionalProperties)
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
    if t.Secondary != nil {
        structMap["secondary"] = t.Secondary.toMap()
    }
    if t.Version != nil {
        structMap["version"] = t.Version
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for TunnelConfigs.
// It customizes the JSON unmarshaling process for TunnelConfigs objects.
func (t *TunnelConfigs) UnmarshalJSON(input []byte) error {
    var temp tempTunnelConfigs
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "auto_provision", "ike_lifetime", "ike_mode", "ike_proposals", "ipsec_lifetime", "ipsec_proposals", "local_id", "mode", "networks", "primary", "probe", "protocol", "provider", "psk", "secondary", "version")
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
    t.Mode = temp.Mode
    t.Networks = temp.Networks
    t.Primary = temp.Primary
    t.Probe = temp.Probe
    t.Protocol = temp.Protocol
    t.Provider = temp.Provider
    t.Psk = temp.Psk
    t.Secondary = temp.Secondary
    t.Version = temp.Version
    return nil
}

// tempTunnelConfigs is a temporary struct used for validating the fields of TunnelConfigs.
type tempTunnelConfigs  struct {
    AutoProvision  *TunnelConfigsAutoProvision          `json:"auto_provision,omitempty"`
    IkeLifetime    *int                                 `json:"ike_lifetime,omitempty"`
    IkeMode        *GatewayTemplateTunnelIkeModeEnum    `json:"ike_mode,omitempty"`
    IkeProposals   []GatewayTemplateTunnelIkeProposal   `json:"ike_proposals,omitempty"`
    IpsecLifetime  *int                                 `json:"ipsec_lifetime,omitempty"`
    IpsecProposals []GatewayTemplateTunnelIpsecProposal `json:"ipsec_proposals,omitempty"`
    LocalId        *string                              `json:"local_id,omitempty"`
    Mode           *GatewayTemplateTunnelModeEnum       `json:"mode,omitempty"`
    Networks       []string                             `json:"networks,omitempty"`
    Primary        *GatewayTemplateTunnelNode           `json:"primary,omitempty"`
    Probe          *GatewayTemplateTunnelProbe          `json:"probe,omitempty"`
    Protocol       *GatewayTemplateTunnelProtocolEnum   `json:"protocol,omitempty"`
    Provider       *TunnelProviderOptionsNameEnum       `json:"provider,omitempty"`
    Psk            *string                              `json:"psk,omitempty"`
    Secondary      *GatewayTemplateTunnelNode           `json:"secondary,omitempty"`
    Version        *GatewayTemplateTunnelVersionEnum    `json:"version,omitempty"`
}
