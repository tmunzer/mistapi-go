// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"strings"
)

// DeviceprofileGateway represents a DeviceprofileGateway struct.
// Gateway Template is applied to a site for gateway(s) in a site.
type DeviceprofileGateway struct {
	// additional CLI commands to append to the generated Junos config. **Note**: no check is done
	AdditionalConfigCmds []string             `json:"additional_config_cmds,omitempty"`
	BgpConfig            map[string]BgpConfig `json:"bgp_config,omitempty"`
	// When the object has been created, in epoch
	CreatedTime *float64     `json:"created_time,omitempty"`
	DhcpdConfig *DhcpdConfig `json:"dhcpd_config,omitempty"`
	DnsOverride *bool        `json:"dnsOverride,omitempty"`
	// Global dns settings. To keep compatibility, dns settings in `ip_config` and `oob_ip_config` will overwrite this setting
	DnsServers []string `json:"dns_servers,omitempty"`
	// Global dns settings. To keep compatibility, dns settings in `ip_config` and `oob_ip_config` will overwrite this setting
	DnsSuffix []string `json:"dns_suffix,omitempty"`
	// Property key is the destination CIDR (e.g. "10.0.0.0/8"), the destination Network name or a variable (e.g. "{{myvar}}")
	ExtraRoutes map[string]GatewayExtraRoute `json:"extra_routes,omitempty"`
	// Property key is the destination CIDR (e.g. "2a02:1234:420a:10c9::/64"), the destination Network name or a variable (e.g. "{{myvar}}")
	ExtraRoutes6 map[string]GatewayExtraRoute `json:"extra_routes6,omitempty"`
	// Gateway matching
	GatewayMatching *GatewayMatching `json:"gateway_matching,omitempty"`
	// Unique ID of the object instance in the Mist Organization
	Id *uuid.UUID `json:"id,omitempty"`
	// Property key is the profile name
	IdpProfiles map[string]IdpProfile `json:"idp_profiles,omitempty"`
	// Property key is the network name
	IpConfigs map[string]GatewayIpConfigProperty `json:"ip_configs,omitempty"`
	// When the object has been modified for the last time, in epoch
	ModifiedTime *float64  `json:"modified_time,omitempty"`
	Name         string    `json:"name"`
	Networks     []Network `json:"networks,omitempty"`
	NtpOverride  *bool     `json:"ntpOverride,omitempty"`
	// List of NTP servers specific to this device. By default, those in Site Settings will be used
	NtpServers []string `json:"ntp_servers,omitempty"`
	// Out-of-band (vme/em0/fxp0) IP config
	OobIpConfig *GatewayOobIpConfig `json:"oob_ip_config,omitempty"`
	OrgId       *uuid.UUID          `json:"org_id,omitempty"`
	// Property key is the path name
	PathPreferences map[string]GatewayPathPreferences `json:"path_preferences,omitempty"`
	// Property key is the port(s) name or range (e.g. "ge-0/0/0-10")
	PortConfig map[string]GatewayPortConfig `json:"port_config,omitempty"`
	// Auto assigned if not set
	RouterId *string `json:"router_id,omitempty"`
	// Property key is the routing policy name
	RoutingPolicies map[string]RoutingPolicy `json:"routing_policies,omitempty"`
	ServicePolicies []ServicePolicy          `json:"service_policies,omitempty"`
	// Property key is the tunnel name
	TunnelConfigs         map[string]TunnelConfig `json:"tunnel_configs,omitempty"`
	TunnelProviderOptions *TunnelProviderOptions  `json:"tunnel_provider_options,omitempty"`
	// Device Type. enum: `gateway`
	Type string `json:"type"`
	// When a service policy denies a app_category, what message to show in user's browser
	UrlFilteringDenyMsg *string    `json:"url_filtering_deny_msg,omitempty"`
	VrfConfig           *VrfConfig `json:"vrf_config,omitempty"`
	// Property key is the network name
	VrfInstances map[string]GatewayVrfInstance `json:"vrf_instances,omitempty"`
	// additional CLI commands to append to the generated SSR config. **Note**: no check is done
	SsrAdditionalConfigCmds []string               `json:"ssr_additional_config_cmds,omitempty"`
	AdditionalProperties    map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for DeviceprofileGateway,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (d DeviceprofileGateway) String() string {
	return fmt.Sprintf(
		"DeviceprofileGateway[AdditionalConfigCmds=%v, BgpConfig=%v, CreatedTime=%v, DhcpdConfig=%v, DnsOverride=%v, DnsServers=%v, DnsSuffix=%v, ExtraRoutes=%v, ExtraRoutes6=%v, GatewayMatching=%v, Id=%v, IdpProfiles=%v, IpConfigs=%v, ModifiedTime=%v, Name=%v, Networks=%v, NtpOverride=%v, NtpServers=%v, OobIpConfig=%v, OrgId=%v, PathPreferences=%v, PortConfig=%v, RouterId=%v, RoutingPolicies=%v, ServicePolicies=%v, TunnelConfigs=%v, TunnelProviderOptions=%v, Type=%v, UrlFilteringDenyMsg=%v, VrfConfig=%v, VrfInstances=%v, SsrAdditionalConfigCmds=%v, AdditionalProperties=%v]",
		d.AdditionalConfigCmds, d.BgpConfig, d.CreatedTime, d.DhcpdConfig, d.DnsOverride, d.DnsServers, d.DnsSuffix, d.ExtraRoutes, d.ExtraRoutes6, d.GatewayMatching, d.Id, d.IdpProfiles, d.IpConfigs, d.ModifiedTime, d.Name, d.Networks, d.NtpOverride, d.NtpServers, d.OobIpConfig, d.OrgId, d.PathPreferences, d.PortConfig, d.RouterId, d.RoutingPolicies, d.ServicePolicies, d.TunnelConfigs, d.TunnelProviderOptions, d.Type, d.UrlFilteringDenyMsg, d.VrfConfig, d.VrfInstances, d.SsrAdditionalConfigCmds, d.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for DeviceprofileGateway.
// It customizes the JSON marshaling process for DeviceprofileGateway objects.
func (d DeviceprofileGateway) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(d.AdditionalProperties,
		"additional_config_cmds", "bgp_config", "created_time", "dhcpd_config", "dnsOverride", "dns_servers", "dns_suffix", "extra_routes", "extra_routes6", "gateway_matching", "id", "idp_profiles", "ip_configs", "modified_time", "name", "networks", "ntpOverride", "ntp_servers", "oob_ip_config", "org_id", "path_preferences", "port_config", "router_id", "routing_policies", "service_policies", "tunnel_configs", "tunnel_provider_options", "type", "url_filtering_deny_msg", "vrf_config", "vrf_instances", "ssr_additional_config_cmds"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(d.toMap())
}

// toMap converts the DeviceprofileGateway object to a map representation for JSON marshaling.
func (d DeviceprofileGateway) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, d.AdditionalProperties)
	if d.AdditionalConfigCmds != nil {
		structMap["additional_config_cmds"] = d.AdditionalConfigCmds
	}
	if d.BgpConfig != nil {
		structMap["bgp_config"] = d.BgpConfig
	}
	if d.CreatedTime != nil {
		structMap["created_time"] = d.CreatedTime
	}
	if d.DhcpdConfig != nil {
		structMap["dhcpd_config"] = d.DhcpdConfig.toMap()
	}
	if d.DnsOverride != nil {
		structMap["dnsOverride"] = d.DnsOverride
	}
	if d.DnsServers != nil {
		structMap["dns_servers"] = d.DnsServers
	}
	if d.DnsSuffix != nil {
		structMap["dns_suffix"] = d.DnsSuffix
	}
	if d.ExtraRoutes != nil {
		structMap["extra_routes"] = d.ExtraRoutes
	}
	if d.ExtraRoutes6 != nil {
		structMap["extra_routes6"] = d.ExtraRoutes6
	}
	if d.GatewayMatching != nil {
		structMap["gateway_matching"] = d.GatewayMatching.toMap()
	}
	if d.Id != nil {
		structMap["id"] = d.Id
	}
	if d.IdpProfiles != nil {
		structMap["idp_profiles"] = d.IdpProfiles
	}
	if d.IpConfigs != nil {
		structMap["ip_configs"] = d.IpConfigs
	}
	if d.ModifiedTime != nil {
		structMap["modified_time"] = d.ModifiedTime
	}
	structMap["name"] = d.Name
	if d.Networks != nil {
		structMap["networks"] = d.Networks
	}
	if d.NtpOverride != nil {
		structMap["ntpOverride"] = d.NtpOverride
	}
	if d.NtpServers != nil {
		structMap["ntp_servers"] = d.NtpServers
	}
	if d.OobIpConfig != nil {
		structMap["oob_ip_config"] = d.OobIpConfig.toMap()
	}
	if d.OrgId != nil {
		structMap["org_id"] = d.OrgId
	}
	if d.PathPreferences != nil {
		structMap["path_preferences"] = d.PathPreferences
	}
	if d.PortConfig != nil {
		structMap["port_config"] = d.PortConfig
	}
	if d.RouterId != nil {
		structMap["router_id"] = d.RouterId
	}
	if d.RoutingPolicies != nil {
		structMap["routing_policies"] = d.RoutingPolicies
	}
	if d.ServicePolicies != nil {
		structMap["service_policies"] = d.ServicePolicies
	}
	if d.TunnelConfigs != nil {
		structMap["tunnel_configs"] = d.TunnelConfigs
	}
	if d.TunnelProviderOptions != nil {
		structMap["tunnel_provider_options"] = d.TunnelProviderOptions.toMap()
	}
	structMap["type"] = d.Type
	if d.UrlFilteringDenyMsg != nil {
		structMap["url_filtering_deny_msg"] = d.UrlFilteringDenyMsg
	}
	if d.VrfConfig != nil {
		structMap["vrf_config"] = d.VrfConfig.toMap()
	}
	if d.VrfInstances != nil {
		structMap["vrf_instances"] = d.VrfInstances
	}
	if d.SsrAdditionalConfigCmds != nil {
		structMap["ssr_additional_config_cmds"] = d.SsrAdditionalConfigCmds
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for DeviceprofileGateway.
// It customizes the JSON unmarshaling process for DeviceprofileGateway objects.
func (d *DeviceprofileGateway) UnmarshalJSON(input []byte) error {
	var temp tempDeviceprofileGateway
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	err = temp.validate()
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "additional_config_cmds", "bgp_config", "created_time", "dhcpd_config", "dnsOverride", "dns_servers", "dns_suffix", "extra_routes", "extra_routes6", "gateway_matching", "id", "idp_profiles", "ip_configs", "modified_time", "name", "networks", "ntpOverride", "ntp_servers", "oob_ip_config", "org_id", "path_preferences", "port_config", "router_id", "routing_policies", "service_policies", "tunnel_configs", "tunnel_provider_options", "type", "url_filtering_deny_msg", "vrf_config", "vrf_instances", "ssr_additional_config_cmds")
	if err != nil {
		return err
	}
	d.AdditionalProperties = additionalProperties

	d.AdditionalConfigCmds = temp.AdditionalConfigCmds
	d.BgpConfig = temp.BgpConfig
	d.CreatedTime = temp.CreatedTime
	d.DhcpdConfig = temp.DhcpdConfig
	d.DnsOverride = temp.DnsOverride
	d.DnsServers = temp.DnsServers
	d.DnsSuffix = temp.DnsSuffix
	d.ExtraRoutes = temp.ExtraRoutes
	d.ExtraRoutes6 = temp.ExtraRoutes6
	d.GatewayMatching = temp.GatewayMatching
	d.Id = temp.Id
	d.IdpProfiles = temp.IdpProfiles
	d.IpConfigs = temp.IpConfigs
	d.ModifiedTime = temp.ModifiedTime
	d.Name = *temp.Name
	d.Networks = temp.Networks
	d.NtpOverride = temp.NtpOverride
	d.NtpServers = temp.NtpServers
	d.OobIpConfig = temp.OobIpConfig
	d.OrgId = temp.OrgId
	d.PathPreferences = temp.PathPreferences
	d.PortConfig = temp.PortConfig
	d.RouterId = temp.RouterId
	d.RoutingPolicies = temp.RoutingPolicies
	d.ServicePolicies = temp.ServicePolicies
	d.TunnelConfigs = temp.TunnelConfigs
	d.TunnelProviderOptions = temp.TunnelProviderOptions
	d.Type = *temp.Type
	d.UrlFilteringDenyMsg = temp.UrlFilteringDenyMsg
	d.VrfConfig = temp.VrfConfig
	d.VrfInstances = temp.VrfInstances
	d.SsrAdditionalConfigCmds = temp.SsrAdditionalConfigCmds
	return nil
}

// tempDeviceprofileGateway is a temporary struct used for validating the fields of DeviceprofileGateway.
type tempDeviceprofileGateway struct {
	AdditionalConfigCmds    []string                           `json:"additional_config_cmds,omitempty"`
	BgpConfig               map[string]BgpConfig               `json:"bgp_config,omitempty"`
	CreatedTime             *float64                           `json:"created_time,omitempty"`
	DhcpdConfig             *DhcpdConfig                       `json:"dhcpd_config,omitempty"`
	DnsOverride             *bool                              `json:"dnsOverride,omitempty"`
	DnsServers              []string                           `json:"dns_servers,omitempty"`
	DnsSuffix               []string                           `json:"dns_suffix,omitempty"`
	ExtraRoutes             map[string]GatewayExtraRoute       `json:"extra_routes,omitempty"`
	ExtraRoutes6            map[string]GatewayExtraRoute       `json:"extra_routes6,omitempty"`
	GatewayMatching         *GatewayMatching                   `json:"gateway_matching,omitempty"`
	Id                      *uuid.UUID                         `json:"id,omitempty"`
	IdpProfiles             map[string]IdpProfile              `json:"idp_profiles,omitempty"`
	IpConfigs               map[string]GatewayIpConfigProperty `json:"ip_configs,omitempty"`
	ModifiedTime            *float64                           `json:"modified_time,omitempty"`
	Name                    *string                            `json:"name"`
	Networks                []Network                          `json:"networks,omitempty"`
	NtpOverride             *bool                              `json:"ntpOverride,omitempty"`
	NtpServers              []string                           `json:"ntp_servers,omitempty"`
	OobIpConfig             *GatewayOobIpConfig                `json:"oob_ip_config,omitempty"`
	OrgId                   *uuid.UUID                         `json:"org_id,omitempty"`
	PathPreferences         map[string]GatewayPathPreferences  `json:"path_preferences,omitempty"`
	PortConfig              map[string]GatewayPortConfig       `json:"port_config,omitempty"`
	RouterId                *string                            `json:"router_id,omitempty"`
	RoutingPolicies         map[string]RoutingPolicy           `json:"routing_policies,omitempty"`
	ServicePolicies         []ServicePolicy                    `json:"service_policies,omitempty"`
	TunnelConfigs           map[string]TunnelConfig            `json:"tunnel_configs,omitempty"`
	TunnelProviderOptions   *TunnelProviderOptions             `json:"tunnel_provider_options,omitempty"`
	Type                    *string                            `json:"type"`
	UrlFilteringDenyMsg     *string                            `json:"url_filtering_deny_msg,omitempty"`
	VrfConfig               *VrfConfig                         `json:"vrf_config,omitempty"`
	VrfInstances            map[string]GatewayVrfInstance      `json:"vrf_instances,omitempty"`
	SsrAdditionalConfigCmds []string                           `json:"ssr_additional_config_cmds,omitempty"`
}

func (d *tempDeviceprofileGateway) validate() error {
	var errs []string
	if d.Name == nil {
		errs = append(errs, "required field `name` is missing for type `deviceprofile_gateway`")
	}
	if d.Type == nil {
		errs = append(errs, "required field `type` is missing for type `deviceprofile_gateway`")
	}
	if len(errs) == 0 {
		return nil
	}
	return errors.New(strings.Join(errs, "\n"))
}
