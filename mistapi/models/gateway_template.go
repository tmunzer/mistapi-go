package models

import (
    "encoding/json"
    "errors"
    "fmt"
    "github.com/google/uuid"
    "strings"
)

// GatewayTemplate represents a GatewayTemplate struct.
// Gateway Template is applied to a site for gateway(s) in a site.
type GatewayTemplate struct {
    // additional CLI commands to append to the generated Junos config. **Note**: no check is done
    AdditionalConfigCmds  []string                           `json:"additional_config_cmds,omitempty"`
    BgpConfig             map[string]BgpConfig               `json:"bgp_config,omitempty"`
    // When the object has been created, in epoch
    CreatedTime           *float64                           `json:"created_time,omitempty"`
    DhcpdConfig           *DhcpdConfig                       `json:"dhcpd_config,omitempty"`
    DnsOverride           *bool                              `json:"dnsOverride,omitempty"`
    // Global dns settings. To keep compatibility, dns settings in `ip_config` and `oob_ip_config` will overwrite this setting
    DnsServers            []string                           `json:"dns_servers,omitempty"`
    // Global dns settings. To keep compatibility, dns settings in `ip_config` and `oob_ip_config` will overwrite this setting
    DnsSuffix             []string                           `json:"dns_suffix,omitempty"`
    // Property key is the destination CIDR (e.g. "10.0.0.0/8")
    ExtraRoutes           map[string]GatewayExtraRoute       `json:"extra_routes,omitempty"`
    // Property key is the destination CIDR (e.g. "2a02:1234:420a:10c9::/64")
    ExtraRoutes6          map[string]GatewayExtraRoute       `json:"extra_routes6,omitempty"`
    // Gateway matching
    GatewayMatching       *GatewayMatching                   `json:"gateway_matching,omitempty"`
    // Unique ID of the object instance in the Mist Organnization
    Id                    *uuid.UUID                         `json:"id,omitempty"`
    // Property key is the profile name
    IdpProfiles           map[string]IdpProfile              `json:"idp_profiles,omitempty"`
    // Property key is the network name
    IpConfigs             map[string]GatewayIpConfigProperty `json:"ip_configs,omitempty"`
    // When the object has been modified for the last time, in epoch
    ModifiedTime          *float64                           `json:"modified_time,omitempty"`
    Name                  string                             `json:"name"`
    Networks              []Network                          `json:"networks,omitempty"`
    NtpOverride           *bool                              `json:"ntpOverride,omitempty"`
    // List of NTP servers specific to this device. By default, those in Site Settings will be used
    NtpServers            []string                           `json:"ntp_servers,omitempty"`
    // Out-of-band (vme/em0/fxp0) IP config
    OobIpConfig           *GatewayOobIpConfig                `json:"oob_ip_config,omitempty"`
    OrgId                 *uuid.UUID                         `json:"org_id,omitempty"`
    // Property key is the path name
    PathPreferences       map[string]GatewayPathPreferences  `json:"path_preferences,omitempty"`
    // Property key is the Port Name (i.e. "ge-0/0/0"), the Ports Range (i.e. "ge-0/0/0-10"), the List of Ports (i.e. "ge-0/0/0,ge-1/0/0", only allowed for Aggregated or Redundant interfaces) or a Variable (i.e. "{{myvar}}").
    PortConfig            map[string]GatewayPortConfig       `json:"port_config,omitempty"`
    // Auto assigned if not set
    RouterId              *string                            `json:"router_id,omitempty"`
    // Property key is the routing policy name
    RoutingPolicies       map[string]RoutingPolicy           `json:"routing_policies,omitempty"`
    ServicePolicies       []ServicePolicy                    `json:"service_policies,omitempty"`
    // Property key is the tunnel name
    TunnelConfigs         map[string]TunnelConfig            `json:"tunnel_configs,omitempty"`
    TunnelProviderOptions *TunnelProviderOptions             `json:"tunnel_provider_options,omitempty"`
    // enum: `spoke`, `standalone`
    Type                  *GatewayTemplateTypeEnum           `json:"type,omitempty"`
    VrfConfig             *VrfConfig                         `json:"vrf_config,omitempty"`
    // Property key is the network name
    VrfInstances          map[string]GatewayVrfInstance      `json:"vrf_instances,omitempty"`
    AdditionalProperties  map[string]interface{}             `json:"_"`
}

// String implements the fmt.Stringer interface for GatewayTemplate,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (g GatewayTemplate) String() string {
    return fmt.Sprintf(
    	"GatewayTemplate[AdditionalConfigCmds=%v, BgpConfig=%v, CreatedTime=%v, DhcpdConfig=%v, DnsOverride=%v, DnsServers=%v, DnsSuffix=%v, ExtraRoutes=%v, ExtraRoutes6=%v, GatewayMatching=%v, Id=%v, IdpProfiles=%v, IpConfigs=%v, ModifiedTime=%v, Name=%v, Networks=%v, NtpOverride=%v, NtpServers=%v, OobIpConfig=%v, OrgId=%v, PathPreferences=%v, PortConfig=%v, RouterId=%v, RoutingPolicies=%v, ServicePolicies=%v, TunnelConfigs=%v, TunnelProviderOptions=%v, Type=%v, VrfConfig=%v, VrfInstances=%v, AdditionalProperties=%v]",
    	g.AdditionalConfigCmds, g.BgpConfig, g.CreatedTime, g.DhcpdConfig, g.DnsOverride, g.DnsServers, g.DnsSuffix, g.ExtraRoutes, g.ExtraRoutes6, g.GatewayMatching, g.Id, g.IdpProfiles, g.IpConfigs, g.ModifiedTime, g.Name, g.Networks, g.NtpOverride, g.NtpServers, g.OobIpConfig, g.OrgId, g.PathPreferences, g.PortConfig, g.RouterId, g.RoutingPolicies, g.ServicePolicies, g.TunnelConfigs, g.TunnelProviderOptions, g.Type, g.VrfConfig, g.VrfInstances, g.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for GatewayTemplate.
// It customizes the JSON marshaling process for GatewayTemplate objects.
func (g GatewayTemplate) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(g.AdditionalProperties,
        "additional_config_cmds", "bgp_config", "created_time", "dhcpd_config", "dnsOverride", "dns_servers", "dns_suffix", "extra_routes", "extra_routes6", "gateway_matching", "id", "idp_profiles", "ip_configs", "modified_time", "name", "networks", "ntpOverride", "ntp_servers", "oob_ip_config", "org_id", "path_preferences", "port_config", "router_id", "routing_policies", "service_policies", "tunnel_configs", "tunnel_provider_options", "type", "vrf_config", "vrf_instances"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(g.toMap())
}

// toMap converts the GatewayTemplate object to a map representation for JSON marshaling.
func (g GatewayTemplate) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, g.AdditionalProperties)
    if g.AdditionalConfigCmds != nil {
        structMap["additional_config_cmds"] = g.AdditionalConfigCmds
    }
    if g.BgpConfig != nil {
        structMap["bgp_config"] = g.BgpConfig
    }
    if g.CreatedTime != nil {
        structMap["created_time"] = g.CreatedTime
    }
    if g.DhcpdConfig != nil {
        structMap["dhcpd_config"] = g.DhcpdConfig.toMap()
    }
    if g.DnsOverride != nil {
        structMap["dnsOverride"] = g.DnsOverride
    }
    if g.DnsServers != nil {
        structMap["dns_servers"] = g.DnsServers
    }
    if g.DnsSuffix != nil {
        structMap["dns_suffix"] = g.DnsSuffix
    }
    if g.ExtraRoutes != nil {
        structMap["extra_routes"] = g.ExtraRoutes
    }
    if g.ExtraRoutes6 != nil {
        structMap["extra_routes6"] = g.ExtraRoutes6
    }
    if g.GatewayMatching != nil {
        structMap["gateway_matching"] = g.GatewayMatching.toMap()
    }
    if g.Id != nil {
        structMap["id"] = g.Id
    }
    if g.IdpProfiles != nil {
        structMap["idp_profiles"] = g.IdpProfiles
    }
    if g.IpConfigs != nil {
        structMap["ip_configs"] = g.IpConfigs
    }
    if g.ModifiedTime != nil {
        structMap["modified_time"] = g.ModifiedTime
    }
    structMap["name"] = g.Name
    if g.Networks != nil {
        structMap["networks"] = g.Networks
    }
    if g.NtpOverride != nil {
        structMap["ntpOverride"] = g.NtpOverride
    }
    if g.NtpServers != nil {
        structMap["ntp_servers"] = g.NtpServers
    }
    if g.OobIpConfig != nil {
        structMap["oob_ip_config"] = g.OobIpConfig.toMap()
    }
    if g.OrgId != nil {
        structMap["org_id"] = g.OrgId
    }
    if g.PathPreferences != nil {
        structMap["path_preferences"] = g.PathPreferences
    }
    if g.PortConfig != nil {
        structMap["port_config"] = g.PortConfig
    }
    if g.RouterId != nil {
        structMap["router_id"] = g.RouterId
    }
    if g.RoutingPolicies != nil {
        structMap["routing_policies"] = g.RoutingPolicies
    }
    if g.ServicePolicies != nil {
        structMap["service_policies"] = g.ServicePolicies
    }
    if g.TunnelConfigs != nil {
        structMap["tunnel_configs"] = g.TunnelConfigs
    }
    if g.TunnelProviderOptions != nil {
        structMap["tunnel_provider_options"] = g.TunnelProviderOptions.toMap()
    }
    if g.Type != nil {
        structMap["type"] = g.Type
    }
    if g.VrfConfig != nil {
        structMap["vrf_config"] = g.VrfConfig.toMap()
    }
    if g.VrfInstances != nil {
        structMap["vrf_instances"] = g.VrfInstances
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for GatewayTemplate.
// It customizes the JSON unmarshaling process for GatewayTemplate objects.
func (g *GatewayTemplate) UnmarshalJSON(input []byte) error {
    var temp tempGatewayTemplate
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "additional_config_cmds", "bgp_config", "created_time", "dhcpd_config", "dnsOverride", "dns_servers", "dns_suffix", "extra_routes", "extra_routes6", "gateway_matching", "id", "idp_profiles", "ip_configs", "modified_time", "name", "networks", "ntpOverride", "ntp_servers", "oob_ip_config", "org_id", "path_preferences", "port_config", "router_id", "routing_policies", "service_policies", "tunnel_configs", "tunnel_provider_options", "type", "vrf_config", "vrf_instances")
    if err != nil {
    	return err
    }
    g.AdditionalProperties = additionalProperties
    
    g.AdditionalConfigCmds = temp.AdditionalConfigCmds
    g.BgpConfig = temp.BgpConfig
    g.CreatedTime = temp.CreatedTime
    g.DhcpdConfig = temp.DhcpdConfig
    g.DnsOverride = temp.DnsOverride
    g.DnsServers = temp.DnsServers
    g.DnsSuffix = temp.DnsSuffix
    g.ExtraRoutes = temp.ExtraRoutes
    g.ExtraRoutes6 = temp.ExtraRoutes6
    g.GatewayMatching = temp.GatewayMatching
    g.Id = temp.Id
    g.IdpProfiles = temp.IdpProfiles
    g.IpConfigs = temp.IpConfigs
    g.ModifiedTime = temp.ModifiedTime
    g.Name = *temp.Name
    g.Networks = temp.Networks
    g.NtpOverride = temp.NtpOverride
    g.NtpServers = temp.NtpServers
    g.OobIpConfig = temp.OobIpConfig
    g.OrgId = temp.OrgId
    g.PathPreferences = temp.PathPreferences
    g.PortConfig = temp.PortConfig
    g.RouterId = temp.RouterId
    g.RoutingPolicies = temp.RoutingPolicies
    g.ServicePolicies = temp.ServicePolicies
    g.TunnelConfigs = temp.TunnelConfigs
    g.TunnelProviderOptions = temp.TunnelProviderOptions
    g.Type = temp.Type
    g.VrfConfig = temp.VrfConfig
    g.VrfInstances = temp.VrfInstances
    return nil
}

// tempGatewayTemplate is a temporary struct used for validating the fields of GatewayTemplate.
type tempGatewayTemplate  struct {
    AdditionalConfigCmds  []string                           `json:"additional_config_cmds,omitempty"`
    BgpConfig             map[string]BgpConfig               `json:"bgp_config,omitempty"`
    CreatedTime           *float64                           `json:"created_time,omitempty"`
    DhcpdConfig           *DhcpdConfig                       `json:"dhcpd_config,omitempty"`
    DnsOverride           *bool                              `json:"dnsOverride,omitempty"`
    DnsServers            []string                           `json:"dns_servers,omitempty"`
    DnsSuffix             []string                           `json:"dns_suffix,omitempty"`
    ExtraRoutes           map[string]GatewayExtraRoute       `json:"extra_routes,omitempty"`
    ExtraRoutes6          map[string]GatewayExtraRoute       `json:"extra_routes6,omitempty"`
    GatewayMatching       *GatewayMatching                   `json:"gateway_matching,omitempty"`
    Id                    *uuid.UUID                         `json:"id,omitempty"`
    IdpProfiles           map[string]IdpProfile              `json:"idp_profiles,omitempty"`
    IpConfigs             map[string]GatewayIpConfigProperty `json:"ip_configs,omitempty"`
    ModifiedTime          *float64                           `json:"modified_time,omitempty"`
    Name                  *string                            `json:"name"`
    Networks              []Network                          `json:"networks,omitempty"`
    NtpOverride           *bool                              `json:"ntpOverride,omitempty"`
    NtpServers            []string                           `json:"ntp_servers,omitempty"`
    OobIpConfig           *GatewayOobIpConfig                `json:"oob_ip_config,omitempty"`
    OrgId                 *uuid.UUID                         `json:"org_id,omitempty"`
    PathPreferences       map[string]GatewayPathPreferences  `json:"path_preferences,omitempty"`
    PortConfig            map[string]GatewayPortConfig       `json:"port_config,omitempty"`
    RouterId              *string                            `json:"router_id,omitempty"`
    RoutingPolicies       map[string]RoutingPolicy           `json:"routing_policies,omitempty"`
    ServicePolicies       []ServicePolicy                    `json:"service_policies,omitempty"`
    TunnelConfigs         map[string]TunnelConfig            `json:"tunnel_configs,omitempty"`
    TunnelProviderOptions *TunnelProviderOptions             `json:"tunnel_provider_options,omitempty"`
    Type                  *GatewayTemplateTypeEnum           `json:"type,omitempty"`
    VrfConfig             *VrfConfig                         `json:"vrf_config,omitempty"`
    VrfInstances          map[string]GatewayVrfInstance      `json:"vrf_instances,omitempty"`
}

func (g *tempGatewayTemplate) validate() error {
    var errs []string
    if g.Name == nil {
        errs = append(errs, "required field `name` is missing for type `gateway_template`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
