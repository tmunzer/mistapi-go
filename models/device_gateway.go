package models

import (
    "encoding/json"
    "github.com/google/uuid"
)

// DeviceGateway represents a DeviceGateway struct.
// device gateway
type DeviceGateway struct {
    // additional CLI commands to append to the generated Junos config
    // **Note**: no check is done
    AdditionalConfigCmds  []string                           `json:"additional_config_cmds,omitempty"`
    BgpConfig             map[string]BgpConfig               `json:"bgp_config,omitempty"`
    CreatedTime           *float64                           `json:"created_time,omitempty"`
    DeviceprofileId       *uuid.UUID                         `json:"deviceprofile_id,omitempty"`
    DhcpdConfig           *DhcpdConfig                       `json:"dhcpd_config,omitempty"`
    // Global dns settings. To keep compatibility, dns settings in `ip_config` and `oob_ip_config` will overwrite this setting
    DnsServers            []string                           `json:"dns_servers,omitempty"`
    // Global dns settings. To keep compatibility, dns settings in `ip_config` and `oob_ip_config` will overwrite this setting
    DnsSuffix             []string                           `json:"dns_suffix,omitempty"`
    ExtraRoutes           map[string]GatewayExtraRoute       `json:"extra_routes,omitempty"`
    ForSite               *bool                              `json:"for_site,omitempty"`
    Id                    *uuid.UUID                         `json:"id,omitempty"`
    // Property key is the profile name
    IdpProfiles           map[string]IdpProfile              `json:"idp_profiles,omitempty"`
    Image1Url             Optional[string]                   `json:"image1_url"`
    Image2Url             Optional[string]                   `json:"image2_url"`
    Image3Url             Optional[string]                   `json:"image3_url"`
    // Property key is the network name
    IpConfigs             map[string]GatewayIpConfigProperty `json:"ip_configs,omitempty"`
    // device MAC address
    Mac                   *string                            `json:"mac,omitempty"`
    Managed               *bool                              `json:"managed,omitempty"`
    // map where the device belongs to
    MapId                 *uuid.UUID                         `json:"map_id,omitempty"`
    // device Model
    Model                 *string                            `json:"model,omitempty"`
    ModifiedTime          *float64                           `json:"modified_time,omitempty"`
    MspId                 *uuid.UUID                         `json:"msp_id,omitempty"`
    Name                  *string                            `json:"name,omitempty"`
    Networks              []Network                          `json:"networks,omitempty"`
    Notes                 *string                            `json:"notes,omitempty"`
    NtpServers            []string                           `json:"ntp_servers,omitempty"`
    // out-of-band (vme/em0/fxp0) IP config
    OobIpConfig           *GatewayOobIpConfig                `json:"oob_ip_config,omitempty"`
    OrgId                 *uuid.UUID                         `json:"org_id,omitempty"`
    // Property key is the path name
    PathPreferences       map[string]GatewayPathPreferences  `json:"path_preferences,omitempty"`
    // Property key is the port name or range (e.g. "ge-0/0/0-10")
    PortConfig            map[string]GatewayPortConfig       `json:"port_config,omitempty"`
    PortMirroring         *GatewayPortMirroring              `json:"port_mirroring,omitempty"`
    // auto assigned if not set
    RouterId              *string                            `json:"router_id,omitempty"`
    // Property key is the routing policy name
    RoutingPolicies       map[string]RoutingPolicy           `json:"routing_policies,omitempty"`
    // device Serial
    Serial                *string                            `json:"serial,omitempty"`
    ServicePolicies       []ServicePolicy                    `json:"service_policies,omitempty"`
    SiteId                *uuid.UUID                         `json:"site_id,omitempty"`
    // Property key is the tunnel name
    TunnelConfigs         map[string]TunnelConfigs           `json:"tunnel_configs,omitempty"`
    TunnelProviderOptions *TunnelProviderOptions             `json:"tunnel_provider_options,omitempty"`
    // Device Type
    Type                  *DeviceTypeGatewayEnum             `json:"type,omitempty"`
    // a dictionary of name->value, the vars can then be used in Wlans. This can overwrite those from Site Vars
    Vars                  map[string]string                  `json:"vars,omitempty"`
    // x in pixel
    X                     *float64                           `json:"x,omitempty"`
    // y in pixel
    Y                     *float64                           `json:"y,omitempty"`
    AdditionalProperties  map[string]any                     `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for DeviceGateway.
// It customizes the JSON marshaling process for DeviceGateway objects.
func (d DeviceGateway) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(d.toMap())
}

// toMap converts the DeviceGateway object to a map representation for JSON marshaling.
func (d DeviceGateway) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, d.AdditionalProperties)
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
    if d.Type != nil {
        structMap["type"] = d.Type
    }
    if d.Vars != nil {
        structMap["vars"] = d.Vars
    }
    if d.X != nil {
        structMap["x"] = d.X
    }
    if d.Y != nil {
        structMap["y"] = d.Y
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for DeviceGateway.
// It customizes the JSON unmarshaling process for DeviceGateway objects.
func (d *DeviceGateway) UnmarshalJSON(input []byte) error {
    var temp deviceGateway
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "additional_config_cmds", "bgp_config", "created_time", "deviceprofile_id", "dhcpd_config", "dns_servers", "dns_suffix", "extra_routes", "for_site", "id", "idp_profiles", "image1_url", "image2_url", "image3_url", "ip_configs", "mac", "managed", "map_id", "model", "modified_time", "msp_id", "name", "networks", "notes", "ntp_servers", "oob_ip_config", "org_id", "path_preferences", "port_config", "port_mirroring", "router_id", "routing_policies", "serial", "service_policies", "site_id", "tunnel_configs", "tunnel_provider_options", "type", "vars", "x", "y")
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
    d.Type = temp.Type
    d.Vars = temp.Vars
    d.X = temp.X
    d.Y = temp.Y
    return nil
}

// deviceGateway is a temporary struct used for validating the fields of DeviceGateway.
type deviceGateway  struct {
    AdditionalConfigCmds  []string                           `json:"additional_config_cmds,omitempty"`
    BgpConfig             map[string]BgpConfig               `json:"bgp_config,omitempty"`
    CreatedTime           *float64                           `json:"created_time,omitempty"`
    DeviceprofileId       *uuid.UUID                         `json:"deviceprofile_id,omitempty"`
    DhcpdConfig           *DhcpdConfig                       `json:"dhcpd_config,omitempty"`
    DnsServers            []string                           `json:"dns_servers,omitempty"`
    DnsSuffix             []string                           `json:"dns_suffix,omitempty"`
    ExtraRoutes           map[string]GatewayExtraRoute       `json:"extra_routes,omitempty"`
    ForSite               *bool                              `json:"for_site,omitempty"`
    Id                    *uuid.UUID                         `json:"id,omitempty"`
    IdpProfiles           map[string]IdpProfile              `json:"idp_profiles,omitempty"`
    Image1Url             Optional[string]                   `json:"image1_url"`
    Image2Url             Optional[string]                   `json:"image2_url"`
    Image3Url             Optional[string]                   `json:"image3_url"`
    IpConfigs             map[string]GatewayIpConfigProperty `json:"ip_configs,omitempty"`
    Mac                   *string                            `json:"mac,omitempty"`
    Managed               *bool                              `json:"managed,omitempty"`
    MapId                 *uuid.UUID                         `json:"map_id,omitempty"`
    Model                 *string                            `json:"model,omitempty"`
    ModifiedTime          *float64                           `json:"modified_time,omitempty"`
    MspId                 *uuid.UUID                         `json:"msp_id,omitempty"`
    Name                  *string                            `json:"name,omitempty"`
    Networks              []Network                          `json:"networks,omitempty"`
    Notes                 *string                            `json:"notes,omitempty"`
    NtpServers            []string                           `json:"ntp_servers,omitempty"`
    OobIpConfig           *GatewayOobIpConfig                `json:"oob_ip_config,omitempty"`
    OrgId                 *uuid.UUID                         `json:"org_id,omitempty"`
    PathPreferences       map[string]GatewayPathPreferences  `json:"path_preferences,omitempty"`
    PortConfig            map[string]GatewayPortConfig       `json:"port_config,omitempty"`
    PortMirroring         *GatewayPortMirroring              `json:"port_mirroring,omitempty"`
    RouterId              *string                            `json:"router_id,omitempty"`
    RoutingPolicies       map[string]RoutingPolicy           `json:"routing_policies,omitempty"`
    Serial                *string                            `json:"serial,omitempty"`
    ServicePolicies       []ServicePolicy                    `json:"service_policies,omitempty"`
    SiteId                *uuid.UUID                         `json:"site_id,omitempty"`
    TunnelConfigs         map[string]TunnelConfigs           `json:"tunnel_configs,omitempty"`
    TunnelProviderOptions *TunnelProviderOptions             `json:"tunnel_provider_options,omitempty"`
    Type                  *DeviceTypeGatewayEnum             `json:"type,omitempty"`
    Vars                  map[string]string                  `json:"vars,omitempty"`
    X                     *float64                           `json:"x,omitempty"`
    Y                     *float64                           `json:"y,omitempty"`
}
