package models

import (
    "encoding/json"
)

// ProtectRe represents a ProtectRe struct.
// restrict inbound-traffic to host
// when enabled, all traffic that is not essential to our operation will be dropped 
// e.g. ntp / dns / traffic to mist will be allowed by default, if dhcpd is enabled, we'll make sure it works
type ProtectRe struct {
    // optionally, services we'll allow
    AllowedServices      []string          `json:"allowed_services,omitempty"`
    Custom               []ProtectReCustom `json:"custom,omitempty"`
    // when enabled, all traffic that is not essential to our operation will be dropped
    // e.g. ntp / dns / traffic to mist will be allowed by default
    // if dhcpd is enabled, we'll make sure it works
    Enabled              *bool             `json:"enabled,omitempty"`
    // host/subnets we'll allow traffic to/from
    TrustedHosts         []string          `json:"trusted_hosts,omitempty"`
    AdditionalProperties map[string]any    `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for ProtectRe.
// It customizes the JSON marshaling process for ProtectRe objects.
func (p ProtectRe) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(p.toMap())
}

// toMap converts the ProtectRe object to a map representation for JSON marshaling.
func (p ProtectRe) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, p.AdditionalProperties)
    if p.AllowedServices != nil {
        structMap["allowed_services"] = p.AllowedServices
    }
    if p.Custom != nil {
        structMap["custom"] = p.Custom
    }
    if p.Enabled != nil {
        structMap["enabled"] = p.Enabled
    }
    if p.TrustedHosts != nil {
        structMap["trusted_hosts"] = p.TrustedHosts
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ProtectRe.
// It customizes the JSON unmarshaling process for ProtectRe objects.
func (p *ProtectRe) UnmarshalJSON(input []byte) error {
    var temp tempProtectRe
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "allowed_services", "custom", "enabled", "trusted_hosts")
    if err != nil {
    	return err
    }
    
    p.AdditionalProperties = additionalProperties
    p.AllowedServices = temp.AllowedServices
    p.Custom = temp.Custom
    p.Enabled = temp.Enabled
    p.TrustedHosts = temp.TrustedHosts
    return nil
}

// tempProtectRe is a temporary struct used for validating the fields of ProtectRe.
type tempProtectRe  struct {
    AllowedServices []string          `json:"allowed_services,omitempty"`
    Custom          []ProtectReCustom `json:"custom,omitempty"`
    Enabled         *bool             `json:"enabled,omitempty"`
    TrustedHosts    []string          `json:"trusted_hosts,omitempty"`
}
