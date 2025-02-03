package models

import (
    "encoding/json"
    "fmt"
)

// SwitchDhcpdConfigProperty represents a SwitchDhcpdConfigProperty struct.
type SwitchDhcpdConfigProperty struct {
    // If `type`==`server` or `type6`==`server` - optional, if not defined, system one will be used
    DnsServers           []string                           `json:"dns_servers,omitempty"`
    // If `type`==`server` or `type6`==`server` - optional, if not defined, system one will be used
    DnsSuffix            []string                           `json:"dns_suffix,omitempty"`
    // If `type`==`server` or `type6`==`server`. Property key is the MAC Address. Format is `[0-9a-f]{12}` (e.g "5684dae9ac8b")
    FixedBindings        map[string]DhcpdConfigFixedBinding `json:"fixed_bindings,omitempty"`
    // If `type`==`server`  - optional, `ip` will be used if not provided
    Gateway              *string                            `json:"gateway,omitempty"`
    // If `type`==`server`
    IpEnd                *string                            `json:"ip_end,omitempty"`
    // If `type6`==`server`
    IpEnd6               *string                            `json:"ip_end6,omitempty"`
    // If `type`==`server`
    IpStart              *string                            `json:"ip_start,omitempty"`
    // If `type6`==`server`
    IpStart6             *string                            `json:"ip_start6,omitempty"`
    // In seconds, lease time has to be between 3600 [1hr] - 604800 [1 week], default is 86400 [1 day]
    LeaseTime            *int                               `json:"lease_time,omitempty"`
    // If `type`==`server` or `type6`==`server`. Property key is the DHCP option number
    Options              map[string]DhcpdConfigOption       `json:"options,omitempty"`
    // `server_id_override`==`true` means the device, when acts as DHCP relay and forwards DHCP responses from DHCP server to clients,
    // should overwrite the Sever Identifier option (i.e. DHCP option 54) in DHCP responses with its own IP address.
    ServerIdOverride     *bool                              `json:"server_id_override,omitempty"`
    // If `type`==`relay`
    Servers              []string                           `json:"servers,omitempty"`
    // If `type6`==`relay`
    Servers6             []string                           `json:"servers6,omitempty"`
    // enum: `none`, `relay` (DHCP Relay), `server` (DHCP Server)
    Type                 *SwitchDhcpdConfigTypeEnum         `json:"type,omitempty"`
    // enum: `none`, `relay` (DHCP Relay), `server` (DHCP Server)
    Type6                *SwitchDhcpdConfigTypeEnum         `json:"type6,omitempty"`
    // If `type`==`server` or `type6`==`server`. Property key is <enterprise number>:<sub option code>, with
    // * enterprise number: 1-65535 (https://www.iana.org/assignments/enterprise-numbers/enterprise-numbers)
    // * sub option code: 1-255, sub-option code'
    VendorEncapsulated   map[string]DhcpdConfigVendorOption `json:"vendor_encapsulated,omitempty"`
    AdditionalProperties map[string]interface{}             `json:"_"`
}

// String implements the fmt.Stringer interface for SwitchDhcpdConfigProperty,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s SwitchDhcpdConfigProperty) String() string {
    return fmt.Sprintf(
    	"SwitchDhcpdConfigProperty[DnsServers=%v, DnsSuffix=%v, FixedBindings=%v, Gateway=%v, IpEnd=%v, IpEnd6=%v, IpStart=%v, IpStart6=%v, LeaseTime=%v, Options=%v, ServerIdOverride=%v, Servers=%v, Servers6=%v, Type=%v, Type6=%v, VendorEncapsulated=%v, AdditionalProperties=%v]",
    	s.DnsServers, s.DnsSuffix, s.FixedBindings, s.Gateway, s.IpEnd, s.IpEnd6, s.IpStart, s.IpStart6, s.LeaseTime, s.Options, s.ServerIdOverride, s.Servers, s.Servers6, s.Type, s.Type6, s.VendorEncapsulated, s.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for SwitchDhcpdConfigProperty.
// It customizes the JSON marshaling process for SwitchDhcpdConfigProperty objects.
func (s SwitchDhcpdConfigProperty) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(s.AdditionalProperties,
        "dns_servers", "dns_suffix", "fixed_bindings", "gateway", "ip_end", "ip_end6", "ip_start", "ip_start6", "lease_time", "options", "server_id_override", "servers", "servers6", "type", "type6", "vendor_encapsulated"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(s.toMap())
}

// toMap converts the SwitchDhcpdConfigProperty object to a map representation for JSON marshaling.
func (s SwitchDhcpdConfigProperty) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, s.AdditionalProperties)
    if s.DnsServers != nil {
        structMap["dns_servers"] = s.DnsServers
    }
    if s.DnsSuffix != nil {
        structMap["dns_suffix"] = s.DnsSuffix
    }
    if s.FixedBindings != nil {
        structMap["fixed_bindings"] = s.FixedBindings
    }
    if s.Gateway != nil {
        structMap["gateway"] = s.Gateway
    }
    if s.IpEnd != nil {
        structMap["ip_end"] = s.IpEnd
    }
    if s.IpEnd6 != nil {
        structMap["ip_end6"] = s.IpEnd6
    }
    if s.IpStart != nil {
        structMap["ip_start"] = s.IpStart
    }
    if s.IpStart6 != nil {
        structMap["ip_start6"] = s.IpStart6
    }
    if s.LeaseTime != nil {
        structMap["lease_time"] = s.LeaseTime
    }
    if s.Options != nil {
        structMap["options"] = s.Options
    }
    if s.ServerIdOverride != nil {
        structMap["server_id_override"] = s.ServerIdOverride
    }
    if s.Servers != nil {
        structMap["servers"] = s.Servers
    }
    if s.Servers6 != nil {
        structMap["servers6"] = s.Servers6
    }
    if s.Type != nil {
        structMap["type"] = s.Type
    }
    if s.Type6 != nil {
        structMap["type6"] = s.Type6
    }
    if s.VendorEncapsulated != nil {
        structMap["vendor_encapsulated"] = s.VendorEncapsulated
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SwitchDhcpdConfigProperty.
// It customizes the JSON unmarshaling process for SwitchDhcpdConfigProperty objects.
func (s *SwitchDhcpdConfigProperty) UnmarshalJSON(input []byte) error {
    var temp tempSwitchDhcpdConfigProperty
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "dns_servers", "dns_suffix", "fixed_bindings", "gateway", "ip_end", "ip_end6", "ip_start", "ip_start6", "lease_time", "options", "server_id_override", "servers", "servers6", "type", "type6", "vendor_encapsulated")
    if err != nil {
    	return err
    }
    s.AdditionalProperties = additionalProperties
    
    s.DnsServers = temp.DnsServers
    s.DnsSuffix = temp.DnsSuffix
    s.FixedBindings = temp.FixedBindings
    s.Gateway = temp.Gateway
    s.IpEnd = temp.IpEnd
    s.IpEnd6 = temp.IpEnd6
    s.IpStart = temp.IpStart
    s.IpStart6 = temp.IpStart6
    s.LeaseTime = temp.LeaseTime
    s.Options = temp.Options
    s.ServerIdOverride = temp.ServerIdOverride
    s.Servers = temp.Servers
    s.Servers6 = temp.Servers6
    s.Type = temp.Type
    s.Type6 = temp.Type6
    s.VendorEncapsulated = temp.VendorEncapsulated
    return nil
}

// tempSwitchDhcpdConfigProperty is a temporary struct used for validating the fields of SwitchDhcpdConfigProperty.
type tempSwitchDhcpdConfigProperty  struct {
    DnsServers         []string                           `json:"dns_servers,omitempty"`
    DnsSuffix          []string                           `json:"dns_suffix,omitempty"`
    FixedBindings      map[string]DhcpdConfigFixedBinding `json:"fixed_bindings,omitempty"`
    Gateway            *string                            `json:"gateway,omitempty"`
    IpEnd              *string                            `json:"ip_end,omitempty"`
    IpEnd6             *string                            `json:"ip_end6,omitempty"`
    IpStart            *string                            `json:"ip_start,omitempty"`
    IpStart6           *string                            `json:"ip_start6,omitempty"`
    LeaseTime          *int                               `json:"lease_time,omitempty"`
    Options            map[string]DhcpdConfigOption       `json:"options,omitempty"`
    ServerIdOverride   *bool                              `json:"server_id_override,omitempty"`
    Servers            []string                           `json:"servers,omitempty"`
    Servers6           []string                           `json:"servers6,omitempty"`
    Type               *SwitchDhcpdConfigTypeEnum         `json:"type,omitempty"`
    Type6              *SwitchDhcpdConfigTypeEnum         `json:"type6,omitempty"`
    VendorEncapsulated map[string]DhcpdConfigVendorOption `json:"vendor_encapsulated,omitempty"`
}
