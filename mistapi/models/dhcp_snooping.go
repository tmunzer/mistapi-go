package models

import (
    "encoding/json"
    "fmt"
)

// DhcpSnooping represents a DhcpSnooping struct.
type DhcpSnooping struct {
    AllNetworks          *bool                  `json:"all_networks,omitempty"`
    // Enable for dynamic ARP inspection check
    EnableArpSpoofCheck  *bool                  `json:"enable_arp_spoof_check,omitempty"`
    // Enable for check for forging source IP address
    EnableIpSourceGuard  *bool                  `json:"enable_ip_source_guard,omitempty"`
    Enabled              *bool                  `json:"enabled,omitempty"`
    // if `all_networks`==`false`, list of network with DHCP snooping enabled
    Networks             []string               `json:"networks,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for DhcpSnooping,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (d DhcpSnooping) String() string {
    return fmt.Sprintf(
    	"DhcpSnooping[AllNetworks=%v, EnableArpSpoofCheck=%v, EnableIpSourceGuard=%v, Enabled=%v, Networks=%v, AdditionalProperties=%v]",
    	d.AllNetworks, d.EnableArpSpoofCheck, d.EnableIpSourceGuard, d.Enabled, d.Networks, d.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for DhcpSnooping.
// It customizes the JSON marshaling process for DhcpSnooping objects.
func (d DhcpSnooping) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(d.AdditionalProperties,
        "all_networks", "enable_arp_spoof_check", "enable_ip_source_guard", "enabled", "networks"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(d.toMap())
}

// toMap converts the DhcpSnooping object to a map representation for JSON marshaling.
func (d DhcpSnooping) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, d.AdditionalProperties)
    if d.AllNetworks != nil {
        structMap["all_networks"] = d.AllNetworks
    }
    if d.EnableArpSpoofCheck != nil {
        structMap["enable_arp_spoof_check"] = d.EnableArpSpoofCheck
    }
    if d.EnableIpSourceGuard != nil {
        structMap["enable_ip_source_guard"] = d.EnableIpSourceGuard
    }
    if d.Enabled != nil {
        structMap["enabled"] = d.Enabled
    }
    if d.Networks != nil {
        structMap["networks"] = d.Networks
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for DhcpSnooping.
// It customizes the JSON unmarshaling process for DhcpSnooping objects.
func (d *DhcpSnooping) UnmarshalJSON(input []byte) error {
    var temp tempDhcpSnooping
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "all_networks", "enable_arp_spoof_check", "enable_ip_source_guard", "enabled", "networks")
    if err != nil {
    	return err
    }
    d.AdditionalProperties = additionalProperties
    
    d.AllNetworks = temp.AllNetworks
    d.EnableArpSpoofCheck = temp.EnableArpSpoofCheck
    d.EnableIpSourceGuard = temp.EnableIpSourceGuard
    d.Enabled = temp.Enabled
    d.Networks = temp.Networks
    return nil
}

// tempDhcpSnooping is a temporary struct used for validating the fields of DhcpSnooping.
type tempDhcpSnooping  struct {
    AllNetworks         *bool    `json:"all_networks,omitempty"`
    EnableArpSpoofCheck *bool    `json:"enable_arp_spoof_check,omitempty"`
    EnableIpSourceGuard *bool    `json:"enable_ip_source_guard,omitempty"`
    Enabled             *bool    `json:"enabled,omitempty"`
    Networks            []string `json:"networks,omitempty"`
}
