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

// DeviceGateway represents a DeviceGateway struct.
// Device gateway
type DeviceGateway struct {
	// additional CLI commands to append to the generated Junos config. **Note**: no check is done
	AdditionalConfigCmds []string             `json:"additional_config_cmds,omitempty"`
	BgpConfig            map[string]BgpConfig `json:"bgp_config,omitempty"`
	// When the object has been created, in epoch
	CreatedTime     *float64     `json:"created_time,omitempty"`
	DeviceprofileId *uuid.UUID   `json:"deviceprofile_id,omitempty"`
	DhcpdConfig     *DhcpdConfig `json:"dhcpd_config,omitempty"`
	// Global dns settings. To keep compatibility, dns settings in `ip_config` and `oob_ip_config` will overwrite this setting
	DnsServers []string `json:"dns_servers,omitempty"`
	// Global dns settings. To keep compatibility, dns settings in `ip_config` and `oob_ip_config` will overwrite this setting
	DnsSuffix []string `json:"dns_suffix,omitempty"`
	// Property key is the destination CIDR (e.g. "10.0.0.0/8"), the destination Network name or a variable (e.g. "{{myvar}}")
	ExtraRoutes map[string]GatewayExtraRoute `json:"extra_routes,omitempty"`
	// Property key is the destination CIDR (e.g. "2a02:1234:420a:10c9::/64"), the destination Network name or a variable (e.g. "{{myvar}}")
	ExtraRoutes6 map[string]GatewayExtraRoute `json:"extra_routes6,omitempty"`
	ForSite      *bool                        `json:"for_site,omitempty"`
	// Unique ID of the object instance in the Mist Organization
	Id *uuid.UUID `json:"id,omitempty"`
	// Property key is the profile name
	IdpProfiles map[string]IdpProfile `json:"idp_profiles,omitempty"`
	Image1Url   Optional[string]      `json:"image1_url"`
	Image2Url   Optional[string]      `json:"image2_url"`
	Image3Url   Optional[string]      `json:"image3_url"`
	// Property key is the network name
	IpConfigs map[string]GatewayIpConfigProperty `json:"ip_configs,omitempty"`
	// Device MAC address
	Mac     *string `json:"mac,omitempty"`
	Managed *bool   `json:"managed,omitempty"`
	// Map where the device belongs to
	MapId *uuid.UUID `json:"map_id,omitempty"`
	// Device Model
	Model *string `json:"model,omitempty"`
	// When the object has been modified for the last time, in epoch
	ModifiedTime *float64   `json:"modified_time,omitempty"`
	MspId        *uuid.UUID `json:"msp_id,omitempty"`
	Name         *string    `json:"name,omitempty"`
	Networks     []Network  `json:"networks,omitempty"`
	Notes        *string    `json:"notes,omitempty"`
	NtpServers   []string   `json:"ntp_servers,omitempty"`
	// Out-of-band (vme/em0/fxp0) IP config
	OobIpConfig *GatewayOobIpConfig `json:"oob_ip_config,omitempty"`
	OrgId       *uuid.UUID          `json:"org_id,omitempty"`
	// Property key is the path name
	PathPreferences map[string]GatewayPathPreferences `json:"path_preferences,omitempty"`
	// Property key is the port name or range (e.g. "ge-0/0/0-10")
	PortConfig    map[string]GatewayPortConfig `json:"port_config,omitempty"`
	PortMirroring *GatewayPortMirroring        `json:"port_mirroring,omitempty"`
	// Auto assigned if not set
	RouterId *string `json:"router_id,omitempty"`
	// Property key is the routing policy name
	RoutingPolicies map[string]RoutingPolicy `json:"routing_policies,omitempty"`
	// Device Serial
	Serial          *string         `json:"serial,omitempty"`
	ServicePolicies []ServicePolicy `json:"service_policies,omitempty"`
	SiteId          *uuid.UUID      `json:"site_id,omitempty"`
	// Property key is the tunnel name
	TunnelConfigs         map[string]TunnelConfig `json:"tunnel_configs,omitempty"`
	TunnelProviderOptions *TunnelProviderOptions  `json:"tunnel_provider_options,omitempty"`
	// Device Type. enum: `gateway`
	Type string `json:"type"`
	// Dictionary of name->value, the vars can then be used in Wlans. This can overwrite those from Site Vars
	Vars      map[string]string `json:"vars,omitempty"`
	VrfConfig *VrfConfig        `json:"vrf_config,omitempty"`
	// Property key is the network name
	VrfInstances map[string]GatewayVrfInstance `json:"vrf_instances,omitempty"`
	// X in pixel
	X *float64 `json:"x,omitempty"`
	// Y in pixel
	Y *float64 `json:"y,omitempty"`
	// additional CLI commands to append to the generated SSR config. **Note**: no check is done
	SsrAdditionalConfigCmds []string               `json:"ssr_additional_config_cmds,omitempty"`
	AdditionalProperties    map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for DeviceGateway,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (d DeviceGateway) String() string {
	return fmt.Sprintf(
		"DeviceGateway[AdditionalConfigCmds=%v, BgpConfig=%v, CreatedTime=%v, DeviceprofileId=%v, DhcpdConfig=%v, DnsServers=%v, DnsSuffix=%v, ExtraRoutes=%v, ExtraRoutes6=%v, ForSite=%v, Id=%v, IdpProfiles=%v, Image1Url=%v, Image2Url=%v, Image3Url=%v, IpConfigs=%v, Mac=%v, Managed=%v, MapId=%v, Model=%v, ModifiedTime=%v, MspId=%v, Name=%v, Networks=%v, Notes=%v, NtpServers=%v, OobIpConfig=%v, OrgId=%v, PathPreferences=%v, PortConfig=%v, PortMirroring=%v, RouterId=%v, RoutingPolicies=%v, Serial=%v, ServicePolicies=%v, SiteId=%v, TunnelConfigs=%v, TunnelProviderOptions=%v, Type=%v, Vars=%v, VrfConfig=%v, VrfInstances=%v, X=%v, Y=%v, SsrAdditionalConfigCmds=%v, AdditionalProperties=%v]",
		d.AdditionalConfigCmds, d.BgpConfig, d.CreatedTime, d.DeviceprofileId, d.DhcpdConfig, d.DnsServers, d.DnsSuffix, d.ExtraRoutes, d.ExtraRoutes6, d.ForSite, d.Id, d.IdpProfiles, d.Image1Url, d.Image2Url, d.Image3Url, d.IpConfigs, d.Mac, d.Managed, d.MapId, d.Model, d.ModifiedTime, d.MspId, d.Name, d.Networks, d.Notes, d.NtpServers, d.OobIpConfig, d.OrgId, d.PathPreferences, d.PortConfig, d.PortMirroring, d.RouterId, d.RoutingPolicies, d.Serial, d.ServicePolicies, d.SiteId, d.TunnelConfigs, d.TunnelProviderOptions, d.Type, d.Vars, d.VrfConfig, d.VrfInstances, d.X, d.Y, d.SsrAdditionalConfigCmds, d.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for DeviceGateway.
// It customizes the JSON marshaling process for DeviceGateway objects.
func (d DeviceGateway) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(d.AdditionalProperties,
		"additional_config_cmds", "bgp_config", "created_time", "deviceprofile_id", "dhcpd_config", "dns_servers", "dns_suffix", "extra_routes", "extra_routes6", "for_site", "id", "idp_profiles", "image1_url", "image2_url", "image3_url", "ip_configs", "mac", "managed", "map_id", "model", "modified_time", "msp_id", "name", "networks", "notes", "ntp_servers", "oob_ip_config", "org_id", "path_preferences", "port_config", "port_mirroring", "router_id", "routing_policies", "serial", "service_policies", "site_id", "tunnel_configs", "tunnel_provider_options", "type", "vars", "vrf_config", "vrf_instances", "x", "y", "ssr_additional_config_cmds"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(d.toMap())
}

// toMap converts the DeviceGateway object to a map representation for JSON marshaling.
func (d DeviceGateway) toMap() map[string]any {
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
	if d.DeviceprofileId != nil {
		structMap["deviceprofile_id"] = d.DeviceprofileId
	}
	if d.DhcpdConfig != nil {
		structMap["dhcpd_config"] = d.DhcpdConfig.toMap()
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
	if d.ForSite != nil {
		structMap["for_site"] = d.ForSite
	}
	if d.Id != nil {
		structMap["id"] = d.Id
	}
	if d.IdpProfiles != nil {
		structMap["idp_profiles"] = d.IdpProfiles
	}
	if d.Image1Url.IsValueSet() {
		if d.Image1Url.Value() != nil {
			structMap["image1_url"] = d.Image1Url.Value()
		} else {
			structMap["image1_url"] = nil
		}
	}
	if d.Image2Url.IsValueSet() {
		if d.Image2Url.Value() != nil {
			structMap["image2_url"] = d.Image2Url.Value()
		} else {
			structMap["image2_url"] = nil
		}
	}
	if d.Image3Url.IsValueSet() {
		if d.Image3Url.Value() != nil {
			structMap["image3_url"] = d.Image3Url.Value()
		} else {
			structMap["image3_url"] = nil
		}
	}
	if d.IpConfigs != nil {
		structMap["ip_configs"] = d.IpConfigs
	}
	if d.Mac != nil {
		structMap["mac"] = d.Mac
	}
	if d.Managed != nil {
		structMap["managed"] = d.Managed
	}
	if d.MapId != nil {
		structMap["map_id"] = d.MapId
	}
	if d.Model != nil {
		structMap["model"] = d.Model
	}
	if d.ModifiedTime != nil {
		structMap["modified_time"] = d.ModifiedTime
	}
	if d.MspId != nil {
		structMap["msp_id"] = d.MspId
	}
	if d.Name != nil {
		structMap["name"] = d.Name
	}
	if d.Networks != nil {
		structMap["networks"] = d.Networks
	}
	if d.Notes != nil {
		structMap["notes"] = d.Notes
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
	if d.PortMirroring != nil {
		structMap["port_mirroring"] = d.PortMirroring.toMap()
	}
	if d.RouterId != nil {
		structMap["router_id"] = d.RouterId
	}
	if d.RoutingPolicies != nil {
		structMap["routing_policies"] = d.RoutingPolicies
	}
	if d.Serial != nil {
		structMap["serial"] = d.Serial
	}
	if d.ServicePolicies != nil {
		structMap["service_policies"] = d.ServicePolicies
	}
	if d.SiteId != nil {
		structMap["site_id"] = d.SiteId
	}
	if d.TunnelConfigs != nil {
		structMap["tunnel_configs"] = d.TunnelConfigs
	}
	if d.TunnelProviderOptions != nil {
		structMap["tunnel_provider_options"] = d.TunnelProviderOptions.toMap()
	}
	structMap["type"] = d.Type
	if d.Vars != nil {
		structMap["vars"] = d.Vars
	}
	if d.VrfConfig != nil {
		structMap["vrf_config"] = d.VrfConfig.toMap()
	}
	if d.VrfInstances != nil {
		structMap["vrf_instances"] = d.VrfInstances
	}
	if d.X != nil {
		structMap["x"] = d.X
	}
	if d.Y != nil {
		structMap["y"] = d.Y
	}
	if d.SsrAdditionalConfigCmds != nil {
		structMap["ssr_additional_config_cmds"] = d.SsrAdditionalConfigCmds
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for DeviceGateway.
// It customizes the JSON unmarshaling process for DeviceGateway objects.
func (d *DeviceGateway) UnmarshalJSON(input []byte) error {
	var temp tempDeviceGateway
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	err = temp.validate()
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "additional_config_cmds", "bgp_config", "created_time", "deviceprofile_id", "dhcpd_config", "dns_servers", "dns_suffix", "extra_routes", "extra_routes6", "for_site", "id", "idp_profiles", "image1_url", "image2_url", "image3_url", "ip_configs", "mac", "managed", "map_id", "model", "modified_time", "msp_id", "name", "networks", "notes", "ntp_servers", "oob_ip_config", "org_id", "path_preferences", "port_config", "port_mirroring", "router_id", "routing_policies", "serial", "service_policies", "site_id", "tunnel_configs", "tunnel_provider_options", "type", "vars", "vrf_config", "vrf_instances", "x", "y", "ssr_additional_config_cmds")
	if err != nil {
		return err
	}
	d.AdditionalProperties = additionalProperties

	d.AdditionalConfigCmds = temp.AdditionalConfigCmds
	d.BgpConfig = temp.BgpConfig
	d.CreatedTime = temp.CreatedTime
	d.DeviceprofileId = temp.DeviceprofileId
	d.DhcpdConfig = temp.DhcpdConfig
	d.DnsServers = temp.DnsServers
	d.DnsSuffix = temp.DnsSuffix
	d.ExtraRoutes = temp.ExtraRoutes
	d.ExtraRoutes6 = temp.ExtraRoutes6
	d.ForSite = temp.ForSite
	d.Id = temp.Id
	d.IdpProfiles = temp.IdpProfiles
	d.Image1Url = temp.Image1Url
	d.Image2Url = temp.Image2Url
	d.Image3Url = temp.Image3Url
	d.IpConfigs = temp.IpConfigs
	d.Mac = temp.Mac
	d.Managed = temp.Managed
	d.MapId = temp.MapId
	d.Model = temp.Model
	d.ModifiedTime = temp.ModifiedTime
	d.MspId = temp.MspId
	d.Name = temp.Name
	d.Networks = temp.Networks
	d.Notes = temp.Notes
	d.NtpServers = temp.NtpServers
	d.OobIpConfig = temp.OobIpConfig
	d.OrgId = temp.OrgId
	d.PathPreferences = temp.PathPreferences
	d.PortConfig = temp.PortConfig
	d.PortMirroring = temp.PortMirroring
	d.RouterId = temp.RouterId
	d.RoutingPolicies = temp.RoutingPolicies
	d.Serial = temp.Serial
	d.ServicePolicies = temp.ServicePolicies
	d.SiteId = temp.SiteId
	d.TunnelConfigs = temp.TunnelConfigs
	d.TunnelProviderOptions = temp.TunnelProviderOptions
	d.Type = *temp.Type
	d.Vars = temp.Vars
	d.VrfConfig = temp.VrfConfig
	d.VrfInstances = temp.VrfInstances
	d.X = temp.X
	d.Y = temp.Y
	d.SsrAdditionalConfigCmds = temp.SsrAdditionalConfigCmds
	return nil
}

// tempDeviceGateway is a temporary struct used for validating the fields of DeviceGateway.
type tempDeviceGateway struct {
	AdditionalConfigCmds    []string                           `json:"additional_config_cmds,omitempty"`
	BgpConfig               map[string]BgpConfig               `json:"bgp_config,omitempty"`
	CreatedTime             *float64                           `json:"created_time,omitempty"`
	DeviceprofileId         *uuid.UUID                         `json:"deviceprofile_id,omitempty"`
	DhcpdConfig             *DhcpdConfig                       `json:"dhcpd_config,omitempty"`
	DnsServers              []string                           `json:"dns_servers,omitempty"`
	DnsSuffix               []string                           `json:"dns_suffix,omitempty"`
	ExtraRoutes             map[string]GatewayExtraRoute       `json:"extra_routes,omitempty"`
	ExtraRoutes6            map[string]GatewayExtraRoute       `json:"extra_routes6,omitempty"`
	ForSite                 *bool                              `json:"for_site,omitempty"`
	Id                      *uuid.UUID                         `json:"id,omitempty"`
	IdpProfiles             map[string]IdpProfile              `json:"idp_profiles,omitempty"`
	Image1Url               Optional[string]                   `json:"image1_url"`
	Image2Url               Optional[string]                   `json:"image2_url"`
	Image3Url               Optional[string]                   `json:"image3_url"`
	IpConfigs               map[string]GatewayIpConfigProperty `json:"ip_configs,omitempty"`
	Mac                     *string                            `json:"mac,omitempty"`
	Managed                 *bool                              `json:"managed,omitempty"`
	MapId                   *uuid.UUID                         `json:"map_id,omitempty"`
	Model                   *string                            `json:"model,omitempty"`
	ModifiedTime            *float64                           `json:"modified_time,omitempty"`
	MspId                   *uuid.UUID                         `json:"msp_id,omitempty"`
	Name                    *string                            `json:"name,omitempty"`
	Networks                []Network                          `json:"networks,omitempty"`
	Notes                   *string                            `json:"notes,omitempty"`
	NtpServers              []string                           `json:"ntp_servers,omitempty"`
	OobIpConfig             *GatewayOobIpConfig                `json:"oob_ip_config,omitempty"`
	OrgId                   *uuid.UUID                         `json:"org_id,omitempty"`
	PathPreferences         map[string]GatewayPathPreferences  `json:"path_preferences,omitempty"`
	PortConfig              map[string]GatewayPortConfig       `json:"port_config,omitempty"`
	PortMirroring           *GatewayPortMirroring              `json:"port_mirroring,omitempty"`
	RouterId                *string                            `json:"router_id,omitempty"`
	RoutingPolicies         map[string]RoutingPolicy           `json:"routing_policies,omitempty"`
	Serial                  *string                            `json:"serial,omitempty"`
	ServicePolicies         []ServicePolicy                    `json:"service_policies,omitempty"`
	SiteId                  *uuid.UUID                         `json:"site_id,omitempty"`
	TunnelConfigs           map[string]TunnelConfig            `json:"tunnel_configs,omitempty"`
	TunnelProviderOptions   *TunnelProviderOptions             `json:"tunnel_provider_options,omitempty"`
	Type                    *string                            `json:"type"`
	Vars                    map[string]string                  `json:"vars,omitempty"`
	VrfConfig               *VrfConfig                         `json:"vrf_config,omitempty"`
	VrfInstances            map[string]GatewayVrfInstance      `json:"vrf_instances,omitempty"`
	X                       *float64                           `json:"x,omitempty"`
	Y                       *float64                           `json:"y,omitempty"`
	SsrAdditionalConfigCmds []string                           `json:"ssr_additional_config_cmds,omitempty"`
}

func (d *tempDeviceGateway) validate() error {
	var errs []string
	if d.Type == nil {
		errs = append(errs, "required field `type` is missing for type `device_gateway`")
	}
	if len(errs) == 0 {
		return nil
	}
	return errors.New(strings.Join(errs, "\n"))
}
