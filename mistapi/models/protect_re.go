package models

import (
    "encoding/json"
    "fmt"
)

// ProtectRe represents a ProtectRe struct.
// Restrict inbound-traffic to host
// when enabled, all traffic that is not essential to our operation will be dropped 
// e.g. ntp / dns / traffic to mist will be allowed by default, if dhcpd is enabled, we'll make sure it works
type ProtectRe struct {
    // Optionally, services we'll allow
    AllowedServices      []ProtectReAllowedServiceEnum `json:"allowed_services,omitempty"`
    Custom               []ProtectReCustom             `json:"custom,omitempty"`
    // When enabled, all traffic that is not essential to our operation will be dropped
    // e.g. ntp / dns / traffic to mist will be allowed by default
    // if dhcpd is enabled, we'll make sure it works
    Enabled              *bool                         `json:"enabled,omitempty"`
    // host/subnets we'll allow traffic to/from
    TrustedHosts         []string                      `json:"trusted_hosts,omitempty"`
    AdditionalProperties map[string]interface{}        `json:"_"`
}

// String implements the fmt.Stringer interface for ProtectRe,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (p ProtectRe) String() string {
    return fmt.Sprintf(
    	"ProtectRe[AllowedServices=%v, Custom=%v, Enabled=%v, TrustedHosts=%v, AdditionalProperties=%v]",
    	p.AllowedServices, p.Custom, p.Enabled, p.TrustedHosts, p.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for ProtectRe.
// It customizes the JSON marshaling process for ProtectRe objects.
func (p ProtectRe) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(p.AdditionalProperties,
        "allowed_services", "custom", "enabled", "trusted_hosts"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(p.toMap())
}

// toMap converts the ProtectRe object to a map representation for JSON marshaling.
func (p ProtectRe) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, p.AdditionalProperties)
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
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "allowed_services", "custom", "enabled", "trusted_hosts")
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
    AllowedServices []ProtectReAllowedServiceEnum `json:"allowed_services,omitempty"`
    Custom          []ProtectReCustom             `json:"custom,omitempty"`
    Enabled         *bool                         `json:"enabled,omitempty"`
    TrustedHosts    []string                      `json:"trusted_hosts,omitempty"`
}
