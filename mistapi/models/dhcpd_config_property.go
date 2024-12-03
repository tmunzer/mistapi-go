package models

import (
    "encoding/json"
)

// DhcpdConfigProperty represents a DhcpdConfigProperty struct.
type DhcpdConfigProperty struct {
    // if `type`==`local` or `type6`==`local` - optional, if not defined, system one will be used
    DnsServers           []string                           `json:"dns_servers,omitempty"`
    // if `type`==`local` or `type6`==`local` - optional, if not defined, system one will be used
    DnsSuffix            []string                           `json:"dns_suffix,omitempty"`
    // if `type`==`local` or `type6`==`local`. Property key is the MAC Address. Format is `[0-9a-f]{12}` (e.g "5684dae9ac8b")
    FixedBindings        map[string]DhcpdConfigFixedBinding `json:"fixed_bindings,omitempty"`
    // if `type`==`local` - optional, `ip` will be used if not provided
    Gateway              *string                            `json:"gateway,omitempty"`
    // if `type`==`local`
    IpEnd                *string                            `json:"ip_end,omitempty"`
    // if `type6`==`local`
    IpEnd6               *string                            `json:"ip_end6,omitempty"`
    // if `type`==`local`
    IpStart              *string                            `json:"ip_start,omitempty"`
    // if `type6`==`local`
    IpStart6             *string                            `json:"ip_start6,omitempty"`
    // in seconds, lease time has to be between 3600 [1hr] - 604800 [1 week], default is 86400 [1 day]
    LeaseTime            *int                               `json:"lease_time,omitempty"`
    // if `type`==`local` or `type6`==`local`. Property key is the DHCP option number
    Options              map[string]DhcpdConfigOption       `json:"options,omitempty"`
    // `server_id_override`==`true` means the device, when acts as DHCP relay and forwards DHCP responses from DHCP server to clients,
    // should overwrite the Sever Identifier option (i.e. DHCP option 54) in DHCP responses with its own IP address.
    ServerIdOverride     *bool                              `json:"server_id_override,omitempty"`
    // if `type`==`relay`
    Servers              []string                           `json:"servers,omitempty"`
    // if `type6`==`relay`
    Servers6             []string                           `json:"servers6,omitempty"`
    // enum: `local` (DHCP Server), `none`, `relay` (DHCP Relay)
    Type                 *DhcpdConfigTypeEnum               `json:"type,omitempty"`
    // enum: `local` (DHCP Server), `none`, `relay` (DHCP Relay)
    Type6                *DhcpdConfigTypeEnum               `json:"type6,omitempty"`
    // if `type`==`local` or `type6`==`local`. Property key is <enterprise number>:<sub option code>, with
    // * enterprise number: 1-65535 (https://www.iana.org/assignments/enterprise-numbers/enterprise-numbers)
    // * sub option code: 1-255, sub-option code'
    VendorEncapulated    map[string]DhcpdConfigVendorOption `json:"vendor_encapulated,omitempty"`
    AdditionalProperties map[string]interface{}             `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for DhcpdConfigProperty.
// It customizes the JSON marshaling process for DhcpdConfigProperty objects.
func (d DhcpdConfigProperty) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(d.AdditionalProperties,
        "dns_servers", "dns_suffix", "fixed_bindings", "gateway", "ip_end", "ip_end6", "ip_start", "ip_start6", "lease_time", "options", "server_id_override", "servers", "servers6", "type", "type6", "vendor_encapulated"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(d.toMap())
}

// toMap converts the DhcpdConfigProperty object to a map representation for JSON marshaling.
func (d DhcpdConfigProperty) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, d.AdditionalProperties)
    if d.DnsServers != nil {
        structMap["dns_servers"] = d.DnsServers
    }
    if d.DnsSuffix != nil {
        structMap["dns_suffix"] = d.DnsSuffix
    }
    if d.FixedBindings != nil {
        structMap["fixed_bindings"] = d.FixedBindings
    }
    if d.Gateway != nil {
        structMap["gateway"] = d.Gateway
    }
    if d.IpEnd != nil {
        structMap["ip_end"] = d.IpEnd
    }
    if d.IpEnd6 != nil {
        structMap["ip_end6"] = d.IpEnd6
    }
    if d.IpStart != nil {
        structMap["ip_start"] = d.IpStart
    }
    if d.IpStart6 != nil {
        structMap["ip_start6"] = d.IpStart6
    }
    if d.LeaseTime != nil {
        structMap["lease_time"] = d.LeaseTime
    }
    if d.Options != nil {
        structMap["options"] = d.Options
    }
    if d.ServerIdOverride != nil {
        structMap["server_id_override"] = d.ServerIdOverride
    }
    if d.Servers != nil {
        structMap["servers"] = d.Servers
    }
    if d.Servers6 != nil {
        structMap["servers6"] = d.Servers6
    }
    if d.Type != nil {
        structMap["type"] = d.Type
    }
    if d.Type6 != nil {
        structMap["type6"] = d.Type6
    }
    if d.VendorEncapulated != nil {
        structMap["vendor_encapulated"] = d.VendorEncapulated
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for DhcpdConfigProperty.
// It customizes the JSON unmarshaling process for DhcpdConfigProperty objects.
func (d *DhcpdConfigProperty) UnmarshalJSON(input []byte) error {
    var temp tempDhcpdConfigProperty
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "dns_servers", "dns_suffix", "fixed_bindings", "gateway", "ip_end", "ip_end6", "ip_start", "ip_start6", "lease_time", "options", "server_id_override", "servers", "servers6", "type", "type6", "vendor_encapulated")
    if err != nil {
    	return err
    }
    d.AdditionalProperties = additionalProperties
    
    d.DnsServers = temp.DnsServers
    d.DnsSuffix = temp.DnsSuffix
    d.FixedBindings = temp.FixedBindings
    d.Gateway = temp.Gateway
    d.IpEnd = temp.IpEnd
    d.IpEnd6 = temp.IpEnd6
    d.IpStart = temp.IpStart
    d.IpStart6 = temp.IpStart6
    d.LeaseTime = temp.LeaseTime
    d.Options = temp.Options
    d.ServerIdOverride = temp.ServerIdOverride
    d.Servers = temp.Servers
    d.Servers6 = temp.Servers6
    d.Type = temp.Type
    d.Type6 = temp.Type6
    d.VendorEncapulated = temp.VendorEncapulated
    return nil
}

// tempDhcpdConfigProperty is a temporary struct used for validating the fields of DhcpdConfigProperty.
type tempDhcpdConfigProperty  struct {
    DnsServers        []string                           `json:"dns_servers,omitempty"`
    DnsSuffix         []string                           `json:"dns_suffix,omitempty"`
    FixedBindings     map[string]DhcpdConfigFixedBinding `json:"fixed_bindings,omitempty"`
    Gateway           *string                            `json:"gateway,omitempty"`
    IpEnd             *string                            `json:"ip_end,omitempty"`
    IpEnd6            *string                            `json:"ip_end6,omitempty"`
    IpStart           *string                            `json:"ip_start,omitempty"`
    IpStart6          *string                            `json:"ip_start6,omitempty"`
    LeaseTime         *int                               `json:"lease_time,omitempty"`
    Options           map[string]DhcpdConfigOption       `json:"options,omitempty"`
    ServerIdOverride  *bool                              `json:"server_id_override,omitempty"`
    Servers           []string                           `json:"servers,omitempty"`
    Servers6          []string                           `json:"servers6,omitempty"`
    Type              *DhcpdConfigTypeEnum               `json:"type,omitempty"`
    Type6             *DhcpdConfigTypeEnum               `json:"type6,omitempty"`
    VendorEncapulated map[string]DhcpdConfigVendorOption `json:"vendor_encapulated,omitempty"`
}
