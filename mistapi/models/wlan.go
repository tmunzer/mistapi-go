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

// Wlan represents a Wlan struct.
// **Note**: portal_template will be forked out of wlan objects soon. To fetch portal_template, please query portal_template_url. To update portal_template, use Wlan Portal Template.
type Wlan struct {
    // Enable coa-immediate-update and address-change-immediate-update on the access profile.
    AcctImmediateUpdate                  *bool                            `json:"acct_immediate_update,omitempty"`
    // How frequently should interim accounting be reported, 60-65535. default is 0 (use one specified in Access-Accept request from RADIUS Server). Very frequent messages can affect the performance of the radius server, 600 and up is recommended when enabled
    AcctInterimInterval                  *int                             `json:"acct_interim_interval,omitempty"`
    // List of RADIUS accounting servers, optional, order matters where the first one is treated as primary
    AcctServers                          []RadiusAcctServer               `json:"acct_servers,omitempty"`
    // Airwatch wlan settings
    Airwatch                             *WlanAirwatch                    `json:"airwatch,omitempty"`
    // Only applicable when limit_bcast==true, which allows or disallows ipv6 Neighbor Discovery packets to go through
    AllowIpv6Ndp                         *bool                            `json:"allow_ipv6_ndp,omitempty"`
    // Only applicable when limit_bcast==true, which allows mDNS / Bonjour packets to go through
    AllowMdns                            *bool                            `json:"allow_mdns,omitempty"`
    // Only applicable when `limit_bcast`==`true`, which allows SSDP
    AllowSsdp                            *bool                            `json:"allow_ssdp,omitempty"`
    // List of device ids
    ApIds                                Optional[[]uuid.UUID]            `json:"ap_ids"`
    // Bandwidth limiting for apps (applies to up/down)
    AppLimit                             *WlanAppLimit                    `json:"app_limit,omitempty"`
    // APp qos wlan settings
    AppQos                               *WlanAppQos                      `json:"app_qos,omitempty"`
    // enum: `aps`, `site`, `wxtags`
    ApplyTo                              *WlanApplyToEnum                 `json:"apply_to,omitempty"`
    // Whether to enable smart arp filter
    ArpFilter                            *bool                            `json:"arp_filter,omitempty"`
    // Authentication wlan settings
    Auth                                 *WlanAuth                        `json:"auth,omitempty"`
    // When ordered, AP will prefer and go back to the first server if possible. enum: `ordered`, `unordered`
    AuthServerSelection                  *WlanAuthServerSelectionEnum     `json:"auth_server_selection,omitempty"`
    // List of RADIUS authentication servers, at least one is needed if `auth type`==`eap`, order matters where the first one is treated as primary
    AuthServers                          []RadiusAuthServer               `json:"auth_servers,omitempty"`
    // Optional, up to 48 bytes, will be dynamically generated if not provided. used only for authentication servers
    AuthServersNasId                     Optional[string]                 `json:"auth_servers_nas_id"`
    // Optional, NAS-IP-ADDRESS to use
    AuthServersNasIp                     Optional[string]                 `json:"auth_servers_nas_ip"`
    // Radius auth session retries. Following fast timers are set if "fast_dot1x_timers" knob is enabled. ‘retries’  are set to value of auth_servers_retries. ‘max-requests’ is also set when setting auth_servers_retries and is set to default value to 3.
    AuthServersRetries                   *int                             `json:"auth_servers_retries,omitempty"`
    // Radius auth session timeout. Following fast timers are set if "fast_dot1x_timers" knob is enabled. ‘quite-period’  and ‘transmit-period’ are set to half the value of auth_servers_timeout. ‘supplicant-timeout’ is also set when setting auth_servers_timeout and is set to default value of 10.
    AuthServersTimeout                   *int                             `json:"auth_servers_timeout,omitempty"`
    // `band` is deprecated and kept for backward compatibility. Use bands instead
    Band                                 *string                          `json:"band,omitempty"`                                     // Deprecated
    // Whether to enable band_steering, this works only when band==both
    BandSteer                            *bool                            `json:"band_steer,omitempty"`
    // Force dual_band capable client to connect to 5G
    BandSteerForceBand5                  *bool                            `json:"band_steer_force_band5,omitempty"`
    // List of radios that the wlan should apply to.
    Bands                                []Dot11BandEnum                  `json:"bands,omitempty"`
    // Whether to block the clients in the blacklist (up to first 256 macs)
    BlockBlacklistClients                *bool                            `json:"block_blacklist_clients,omitempty"`
    // Bonjour gateway wlan settings
    Bonjour                              *WlanBonjour                     `json:"bonjour,omitempty"`
    // Cisco CWA (central web authentication) required RADIUS with COA in order to work. See CWA: https://www.cisco.com/c/en/us/support/docs/security/identity-services-engine/115732-central-web-auth-00.html
    CiscoCwa                             *WlanCiscoCwa                    `json:"cisco_cwa,omitempty"`
    // In kbps, value from 1 to 999000
    ClientLimitDown                      *WlanLimit                       `json:"client_limit_down,omitempty"`
    // If downlink limiting per-client is enabled
    ClientLimitDownEnabled               *bool                            `json:"client_limit_down_enabled,omitempty"`
    // In kbps, value from 1 to 999000
    ClientLimitUp                        *WlanLimit                       `json:"client_limit_up,omitempty"`
    // If uplink limiting per-client is enabled
    ClientLimitUpEnabled                 *bool                            `json:"client_limit_up_enabled,omitempty"`
    // List of COA (change of authorization) servers, optional
    CoaServers                           []CoaServer                      `json:"coa_servers,omitempty"`
    // When the object has been created, in epoch
    CreatedTime                          *float64                         `json:"created_time,omitempty"`
    // Some old WLAN drivers may not be compatible
    Disable11ax                          *bool                            `json:"disable_11ax,omitempty"`
    // To disable Wi-Fi 7 EHT IEs
    Disable11be                          *bool                            `json:"disable_11be,omitempty"`
    // To disable ht or vht rates
    DisableHtVhtRates                    *bool                            `json:"disable_ht_vht_rates,omitempty"`
    // Whether to disable U-APSD
    DisableUapsd                         *bool                            `json:"disable_uapsd,omitempty"`
    // Disable sending v2 roam notification messages
    DisableV1RoamNotify                  *bool                            `json:"disable_v1_roam_notify,omitempty"`
    // Disable sending v2 roam notification messages
    DisableV2RoamNotify                  *bool                            `json:"disable_v2_roam_notify,omitempty"`
    // When any of the following is true, this WLAN will be disabled
    // * cannot get IP
    // * cannot obtain default gateway
    // * cannot reach default gateway
    DisableWhenGatewayUnreachable        *bool                            `json:"disable_when_gateway_unreachable,omitempty"`
    DisableWhenMxtunnelDown              *bool                            `json:"disable_when_mxtunnel_down,omitempty"`
    // Whether to disable WMM
    DisableWmm                           *bool                            `json:"disable_wmm,omitempty"`
    // For radius_group-based DNS server (rewrite DNS request depending on the Group RADIUS server returns)
    DnsServerRewrite                     Optional[WlanDnsServerRewrite]   `json:"dns_server_rewrite"`
    Dtim                                 *int                             `json:"dtim,omitempty"`
    // For dynamic PSK where we get per_user PSK from Radius. dynamic_psk allows PSK to be selected at runtime depending on context (wlan/site/user/...) thus following configurations are assumed (currently)
    // * PSK will come from RADIUS server
    // * AP sends client MAC as username and password (i.e. `enable_mac_auth` is assumed)
    // * AP sends BSSID:SSID as Caller-Station-ID
    // * `auth_servers` is required
    // * PSK will come from cloud WLC if source is cloud_psks
    // * default_psk will be used if cloud WLC is not available
    // * `multi_psk_only` and `psk` is ignored
    // * `pairwise` can only be wpa2-ccmp (for now, wpa3 support on the roadmap)
    DynamicPsk                           Optional[WlanDynamicPsk]         `json:"dynamic_psk"`
    // For 802.1x
    DynamicVlan                          Optional[WlanDynamicVlan]        `json:"dynamic_vlan"`
    // Enable AP-AP keycaching via multicast
    EnableLocalKeycaching                *bool                            `json:"enable_local_keycaching,omitempty"`
    // By default, we'd inspect all DHCP packets and drop those unrelated to the wireless client itself in the case where client is a wireless bridge (DHCP packets for other MACs will need to be forwarded), wireless_bridging can be enabled
    EnableWirelessBridging               *bool                            `json:"enable_wireless_bridging,omitempty"`
    // If the client bridge is doing DHCP on behalf of other devices (L2-NAT), enable dhcp_tracking will cut down DHCP response packets to be forwarded to wireless
    EnableWirelessBridgingDhcpTracking   *bool                            `json:"enable_wireless_bridging_dhcp_tracking,omitempty"`
    // If this wlan is enabled
    Enabled                              *bool                            `json:"enabled,omitempty"`
    // If set to true, sets default fast-timers with values calculated from ‘auth_servers_timeout’ and ‘auth_server_retries’ .
    FastDot1xTimers                      *bool                            `json:"fast_dot1x_timers,omitempty"`
    ForSite                              *bool                            `json:"for_site,omitempty"`
    // Whether to hide SSID in beacon
    HideSsid                             *bool                            `json:"hide_ssid,omitempty"`
    // Include hostname inside IE in AP beacons / probe responses
    HostnameIe                           *bool                            `json:"hostname_ie,omitempty"`
    // Hostspot 2.0 wlan settings
    Hotspot20                            *WlanHotspot20                   `json:"hotspot20,omitempty"`
    // Unique ID of the object instance in the Mist Organization
    Id                                   *uuid.UUID                       `json:"id,omitempty"`
    InjectDhcpOption82                   *WlanInjectDhcpOption82          `json:"inject_dhcp_option_82,omitempty"`
    // where this WLAN will be connected to. enum: `all`, `eth0`, `eth1`, `eth2`, `eth3`, `mxtunnel`, `site_mxedge`, `wxtunnel`
    Interface                            *WlanInterfaceEnum               `json:"interface,omitempty"`
    // Whether to stop clients to talk to each other
    Isolation                            *bool                            `json:"isolation,omitempty"`
    // If isolation is enabled, whether to deny clients to talk to L2 on the LAN
    L2Isolation                          *bool                            `json:"l2_isolation,omitempty"`
    // Legacy devices requires the Over-DS (for Fast BSS Transition) bit set (while our chip doesn’t support it). Warning! Enabling this will cause problem for iOS devices.
    LegacyOverds                         *bool                            `json:"legacy_overds,omitempty"`
    // Whether to limit broadcast packets going to wireless (i.e. only allow certain bcast packets to go through)
    LimitBcast                           *bool                            `json:"limit_bcast,omitempty"`
    // Limit probe response base on some heuristic rules
    LimitProbeResponse                   *bool                            `json:"limit_probe_response,omitempty"`
    // Max idle time in seconds
    MaxIdletime                          *int                             `json:"max_idletime,omitempty"`
    // Maximum number of client connected to the SSID. `0` means unlimited
    MaxNumClients                        *int                             `json:"max_num_clients,omitempty"`
    MistNac                              *WlanMistNac                     `json:"mist_nac,omitempty"`
    // When the object has been modified for the last time, in epoch
    ModifiedTime                         *float64                         `json:"modified_time,omitempty"`
    MspId                                *uuid.UUID                       `json:"msp_id,omitempty"`
    // (deprecated, use mxtunnel_ids instead) when `interface`==`mxtunnel`, id of the Mist Tunnel
    MxtunnelId                           *uuid.UUID                       `json:"mxtunnel_id,omitempty"`                              // Deprecated
    // When `interface`=`mxtunnel`, id of the Mist Tunnel
    MxtunnelIds                          []string                         `json:"mxtunnel_ids,omitempty"`
    // When `interface`=`site_mxedge`, name of the mxtunnel that in mxtunnels under Site Setting
    MxtunnelName                         []string                         `json:"mxtunnel_name,omitempty"`
    // Whether to only allow client to use DNS that we’ve learned from DHCP response
    NoStaticDns                          *bool                            `json:"no_static_dns,omitempty"`
    // Whether to only allow client that we’ve learned from DHCP exchange to talk
    NoStaticIp                           *bool                            `json:"no_static_ip,omitempty"`
    OrgId                                *uuid.UUID                       `json:"org_id,omitempty"`
    // Portal wlan settings
    Portal                               *WlanPortal                      `json:"portal,omitempty"`
    // List of hostnames without http(s):// (matched by substring)
    PortalAllowedHostnames               []string                         `json:"portal_allowed_hostnames,omitempty"`
    // List of CIDRs
    PortalAllowedSubnets                 []string                         `json:"portal_allowed_subnets,omitempty"`
    // APi secret (auto-generated) that can be used to sign guest authorization requests
    PortalApiSecret                      Optional[string]                 `json:"portal_api_secret"`
    // List of hostnames without http(s):// (matched by substring), this takes precedence over portal_allowed_hostnames
    PortalDeniedHostnames                []string                         `json:"portal_denied_hostnames,omitempty"`
    // Url of portal background image
    PortalImage                          Optional[string]                 `json:"portal_image"`
    PortalSsoUrl                         Optional[string]                 `json:"portal_sso_url"`
    // N.B portal_template will be forked out of wlan objects soon. To fetch portal_template, please query portal_template_url. To update portal_template, use Wlan Portal Template.
    PortalTemplateUrl                    Optional[string]                 `json:"portal_template_url"`
    Qos                                  *WlanQos                         `json:"qos,omitempty"`
    // RadSec settings
    Radsec                               *Radsec                          `json:"radsec,omitempty"`
    // Property key is the RF band. enum: `24`, `5`, `6`
    Rateset                              map[string]WlanDatarates         `json:"rateset,omitempty"`
    // When different mxcluster is on different subnet, we'd want to disconnect clients (so they'll reconnect and get new IPs)
    ReconnectClientsWhenRoamingMxcluster *bool                            `json:"reconnect_clients_when_roaming_mxcluster,omitempty"`
    // enum: `11r`, `OKC`, `NONE`
    RoamMode                             *WlanRoamModeEnum                `json:"roam_mode,omitempty"`
    // WLAN operating schedule, default is disabled
    Schedule                             *WlanSchedule                    `json:"schedule,omitempty"`
    SiteId                               *uuid.UUID                       `json:"site_id,omitempty"`
    // Whether to exclude this WLAN from SLE metrics
    SleExcluded                          *bool                            `json:"sle_excluded,omitempty"`
    // Name of the SSID
    Ssid                                 string                           `json:"ssid"`
    TemplateId                           Optional[uuid.UUID]              `json:"template_id"`
    // Url of portal background image thumbnail
    Thumbnail                            Optional[string]                 `json:"thumbnail"`
    // If `auth.type`==`eap` or `auth.type`==`psk`, should only be set for legacy client, such as pre-2004, 802.11b devices
    UseEapolV1                           *bool                            `json:"use_eapol_v1,omitempty"`
    // If vlan tagging is enabled
    VlanEnabled                          *bool                            `json:"vlan_enabled,omitempty"`
    VlanId                               Optional[WlanVlanIdWithVariable] `json:"vlan_id"`
    VlanIds                              *WlanVlanIds                     `json:"vlan_ids,omitempty"`
    // Requires `vlan_enabled`==`true` to be set to `true`. Vlan pooling allows AP to place client on different VLAN using a deterministic algorithm
    VlanPooling                          *bool                            `json:"vlan_pooling,omitempty"`
    // In kbps, value from 1 to 999000
    WlanLimitDown                        *WlanLimit                       `json:"wlan_limit_down,omitempty"`
    // If downlink limiting for whole wlan is enabled
    WlanLimitDownEnabled                 *bool                            `json:"wlan_limit_down_enabled,omitempty"`
    // In kbps, value from 1 to 999000
    WlanLimitUp                          *WlanLimit                       `json:"wlan_limit_up,omitempty"`
    // If uplink limiting for whole wlan is enabled
    WlanLimitUpEnabled                   *bool                            `json:"wlan_limit_up_enabled,omitempty"`
    // List of wxtag_ids
    WxtagIds                             Optional[[]uuid.UUID]            `json:"wxtag_ids"`
    // When `interface`=`wxtunnel`, id of the WXLAN Tunnel
    WxtunnelId                           Optional[string]                 `json:"wxtunnel_id"`
    // When `interface`=`wxtunnel`, remote tunnel identifier
    WxtunnelRemoteId                     Optional[string]                 `json:"wxtunnel_remote_id"`
    AdditionalProperties                 map[string]interface{}           `json:"_"`
}

// String implements the fmt.Stringer interface for Wlan,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (w Wlan) String() string {
    return fmt.Sprintf(
    	"Wlan[AcctImmediateUpdate=%v, AcctInterimInterval=%v, AcctServers=%v, Airwatch=%v, AllowIpv6Ndp=%v, AllowMdns=%v, AllowSsdp=%v, ApIds=%v, AppLimit=%v, AppQos=%v, ApplyTo=%v, ArpFilter=%v, Auth=%v, AuthServerSelection=%v, AuthServers=%v, AuthServersNasId=%v, AuthServersNasIp=%v, AuthServersRetries=%v, AuthServersTimeout=%v, Band=%v, BandSteer=%v, BandSteerForceBand5=%v, Bands=%v, BlockBlacklistClients=%v, Bonjour=%v, CiscoCwa=%v, ClientLimitDown=%v, ClientLimitDownEnabled=%v, ClientLimitUp=%v, ClientLimitUpEnabled=%v, CoaServers=%v, CreatedTime=%v, Disable11ax=%v, Disable11be=%v, DisableHtVhtRates=%v, DisableUapsd=%v, DisableV1RoamNotify=%v, DisableV2RoamNotify=%v, DisableWhenGatewayUnreachable=%v, DisableWhenMxtunnelDown=%v, DisableWmm=%v, DnsServerRewrite=%v, Dtim=%v, DynamicPsk=%v, DynamicVlan=%v, EnableLocalKeycaching=%v, EnableWirelessBridging=%v, EnableWirelessBridgingDhcpTracking=%v, Enabled=%v, FastDot1xTimers=%v, ForSite=%v, HideSsid=%v, HostnameIe=%v, Hotspot20=%v, Id=%v, InjectDhcpOption82=%v, Interface=%v, Isolation=%v, L2Isolation=%v, LegacyOverds=%v, LimitBcast=%v, LimitProbeResponse=%v, MaxIdletime=%v, MaxNumClients=%v, MistNac=%v, ModifiedTime=%v, MspId=%v, MxtunnelId=%v, MxtunnelIds=%v, MxtunnelName=%v, NoStaticDns=%v, NoStaticIp=%v, OrgId=%v, Portal=%v, PortalAllowedHostnames=%v, PortalAllowedSubnets=%v, PortalApiSecret=%v, PortalDeniedHostnames=%v, PortalImage=%v, PortalSsoUrl=%v, PortalTemplateUrl=%v, Qos=%v, Radsec=%v, Rateset=%v, ReconnectClientsWhenRoamingMxcluster=%v, RoamMode=%v, Schedule=%v, SiteId=%v, SleExcluded=%v, Ssid=%v, TemplateId=%v, Thumbnail=%v, UseEapolV1=%v, VlanEnabled=%v, VlanId=%v, VlanIds=%v, VlanPooling=%v, WlanLimitDown=%v, WlanLimitDownEnabled=%v, WlanLimitUp=%v, WlanLimitUpEnabled=%v, WxtagIds=%v, WxtunnelId=%v, WxtunnelRemoteId=%v, AdditionalProperties=%v]",
    	w.AcctImmediateUpdate, w.AcctInterimInterval, w.AcctServers, w.Airwatch, w.AllowIpv6Ndp, w.AllowMdns, w.AllowSsdp, w.ApIds, w.AppLimit, w.AppQos, w.ApplyTo, w.ArpFilter, w.Auth, w.AuthServerSelection, w.AuthServers, w.AuthServersNasId, w.AuthServersNasIp, w.AuthServersRetries, w.AuthServersTimeout, w.Band, w.BandSteer, w.BandSteerForceBand5, w.Bands, w.BlockBlacklistClients, w.Bonjour, w.CiscoCwa, w.ClientLimitDown, w.ClientLimitDownEnabled, w.ClientLimitUp, w.ClientLimitUpEnabled, w.CoaServers, w.CreatedTime, w.Disable11ax, w.Disable11be, w.DisableHtVhtRates, w.DisableUapsd, w.DisableV1RoamNotify, w.DisableV2RoamNotify, w.DisableWhenGatewayUnreachable, w.DisableWhenMxtunnelDown, w.DisableWmm, w.DnsServerRewrite, w.Dtim, w.DynamicPsk, w.DynamicVlan, w.EnableLocalKeycaching, w.EnableWirelessBridging, w.EnableWirelessBridgingDhcpTracking, w.Enabled, w.FastDot1xTimers, w.ForSite, w.HideSsid, w.HostnameIe, w.Hotspot20, w.Id, w.InjectDhcpOption82, w.Interface, w.Isolation, w.L2Isolation, w.LegacyOverds, w.LimitBcast, w.LimitProbeResponse, w.MaxIdletime, w.MaxNumClients, w.MistNac, w.ModifiedTime, w.MspId, w.MxtunnelId, w.MxtunnelIds, w.MxtunnelName, w.NoStaticDns, w.NoStaticIp, w.OrgId, w.Portal, w.PortalAllowedHostnames, w.PortalAllowedSubnets, w.PortalApiSecret, w.PortalDeniedHostnames, w.PortalImage, w.PortalSsoUrl, w.PortalTemplateUrl, w.Qos, w.Radsec, w.Rateset, w.ReconnectClientsWhenRoamingMxcluster, w.RoamMode, w.Schedule, w.SiteId, w.SleExcluded, w.Ssid, w.TemplateId, w.Thumbnail, w.UseEapolV1, w.VlanEnabled, w.VlanId, w.VlanIds, w.VlanPooling, w.WlanLimitDown, w.WlanLimitDownEnabled, w.WlanLimitUp, w.WlanLimitUpEnabled, w.WxtagIds, w.WxtunnelId, w.WxtunnelRemoteId, w.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for Wlan.
// It customizes the JSON marshaling process for Wlan objects.
func (w Wlan) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(w.AdditionalProperties,
        "acct_immediate_update", "acct_interim_interval", "acct_servers", "airwatch", "allow_ipv6_ndp", "allow_mdns", "allow_ssdp", "ap_ids", "app_limit", "app_qos", "apply_to", "arp_filter", "auth", "auth_server_selection", "auth_servers", "auth_servers_nas_id", "auth_servers_nas_ip", "auth_servers_retries", "auth_servers_timeout", "band", "band_steer", "band_steer_force_band5", "bands", "block_blacklist_clients", "bonjour", "cisco_cwa", "client_limit_down", "client_limit_down_enabled", "client_limit_up", "client_limit_up_enabled", "coa_servers", "created_time", "disable_11ax", "disable_11be", "disable_ht_vht_rates", "disable_uapsd", "disable_v1_roam_notify", "disable_v2_roam_notify", "disable_when_gateway_unreachable", "disable_when_mxtunnel_down", "disable_wmm", "dns_server_rewrite", "dtim", "dynamic_psk", "dynamic_vlan", "enable_local_keycaching", "enable_wireless_bridging", "enable_wireless_bridging_dhcp_tracking", "enabled", "fast_dot1x_timers", "for_site", "hide_ssid", "hostname_ie", "hotspot20", "id", "inject_dhcp_option_82", "interface", "isolation", "l2_isolation", "legacy_overds", "limit_bcast", "limit_probe_response", "max_idletime", "max_num_clients", "mist_nac", "modified_time", "msp_id", "mxtunnel_id", "mxtunnel_ids", "mxtunnel_name", "no_static_dns", "no_static_ip", "org_id", "portal", "portal_allowed_hostnames", "portal_allowed_subnets", "portal_api_secret", "portal_denied_hostnames", "portal_image", "portal_sso_url", "portal_template_url", "qos", "radsec", "rateset", "reconnect_clients_when_roaming_mxcluster", "roam_mode", "schedule", "site_id", "sle_excluded", "ssid", "template_id", "thumbnail", "use_eapol_v1", "vlan_enabled", "vlan_id", "vlan_ids", "vlan_pooling", "wlan_limit_down", "wlan_limit_down_enabled", "wlan_limit_up", "wlan_limit_up_enabled", "wxtag_ids", "wxtunnel_id", "wxtunnel_remote_id"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(w.toMap())
}

// toMap converts the Wlan object to a map representation for JSON marshaling.
func (w Wlan) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, w.AdditionalProperties)
    if w.AcctImmediateUpdate != nil {
        structMap["acct_immediate_update"] = w.AcctImmediateUpdate
    }
    if w.AcctInterimInterval != nil {
        structMap["acct_interim_interval"] = w.AcctInterimInterval
    }
    if w.AcctServers != nil {
        structMap["acct_servers"] = w.AcctServers
    }
    if w.Airwatch != nil {
        structMap["airwatch"] = w.Airwatch.toMap()
    }
    if w.AllowIpv6Ndp != nil {
        structMap["allow_ipv6_ndp"] = w.AllowIpv6Ndp
    }
    if w.AllowMdns != nil {
        structMap["allow_mdns"] = w.AllowMdns
    }
    if w.AllowSsdp != nil {
        structMap["allow_ssdp"] = w.AllowSsdp
    }
    if w.ApIds.IsValueSet() {
        if w.ApIds.Value() != nil {
            structMap["ap_ids"] = w.ApIds.Value()
        } else {
            structMap["ap_ids"] = nil
        }
    }
    if w.AppLimit != nil {
        structMap["app_limit"] = w.AppLimit.toMap()
    }
    if w.AppQos != nil {
        structMap["app_qos"] = w.AppQos.toMap()
    }
    if w.ApplyTo != nil {
        structMap["apply_to"] = w.ApplyTo
    }
    if w.ArpFilter != nil {
        structMap["arp_filter"] = w.ArpFilter
    }
    if w.Auth != nil {
        structMap["auth"] = w.Auth.toMap()
    }
    if w.AuthServerSelection != nil {
        structMap["auth_server_selection"] = w.AuthServerSelection
    }
    if w.AuthServers != nil {
        structMap["auth_servers"] = w.AuthServers
    }
    if w.AuthServersNasId.IsValueSet() {
        if w.AuthServersNasId.Value() != nil {
            structMap["auth_servers_nas_id"] = w.AuthServersNasId.Value()
        } else {
            structMap["auth_servers_nas_id"] = nil
        }
    }
    if w.AuthServersNasIp.IsValueSet() {
        if w.AuthServersNasIp.Value() != nil {
            structMap["auth_servers_nas_ip"] = w.AuthServersNasIp.Value()
        } else {
            structMap["auth_servers_nas_ip"] = nil
        }
    }
    if w.AuthServersRetries != nil {
        structMap["auth_servers_retries"] = w.AuthServersRetries
    }
    if w.AuthServersTimeout != nil {
        structMap["auth_servers_timeout"] = w.AuthServersTimeout
    }
    if w.Band != nil {
        structMap["band"] = w.Band
    }
    if w.BandSteer != nil {
        structMap["band_steer"] = w.BandSteer
    }
    if w.BandSteerForceBand5 != nil {
        structMap["band_steer_force_band5"] = w.BandSteerForceBand5
    }
    if w.Bands != nil {
        structMap["bands"] = w.Bands
    }
    if w.BlockBlacklistClients != nil {
        structMap["block_blacklist_clients"] = w.BlockBlacklistClients
    }
    if w.Bonjour != nil {
        structMap["bonjour"] = w.Bonjour.toMap()
    }
    if w.CiscoCwa != nil {
        structMap["cisco_cwa"] = w.CiscoCwa.toMap()
    }
    if w.ClientLimitDown != nil {
        structMap["client_limit_down"] = w.ClientLimitDown.toMap()
    }
    if w.ClientLimitDownEnabled != nil {
        structMap["client_limit_down_enabled"] = w.ClientLimitDownEnabled
    }
    if w.ClientLimitUp != nil {
        structMap["client_limit_up"] = w.ClientLimitUp.toMap()
    }
    if w.ClientLimitUpEnabled != nil {
        structMap["client_limit_up_enabled"] = w.ClientLimitUpEnabled
    }
    if w.CoaServers != nil {
        structMap["coa_servers"] = w.CoaServers
    }
    if w.CreatedTime != nil {
        structMap["created_time"] = w.CreatedTime
    }
    if w.Disable11ax != nil {
        structMap["disable_11ax"] = w.Disable11ax
    }
    if w.Disable11be != nil {
        structMap["disable_11be"] = w.Disable11be
    }
    if w.DisableHtVhtRates != nil {
        structMap["disable_ht_vht_rates"] = w.DisableHtVhtRates
    }
    if w.DisableUapsd != nil {
        structMap["disable_uapsd"] = w.DisableUapsd
    }
    if w.DisableV1RoamNotify != nil {
        structMap["disable_v1_roam_notify"] = w.DisableV1RoamNotify
    }
    if w.DisableV2RoamNotify != nil {
        structMap["disable_v2_roam_notify"] = w.DisableV2RoamNotify
    }
    if w.DisableWhenGatewayUnreachable != nil {
        structMap["disable_when_gateway_unreachable"] = w.DisableWhenGatewayUnreachable
    }
    if w.DisableWhenMxtunnelDown != nil {
        structMap["disable_when_mxtunnel_down"] = w.DisableWhenMxtunnelDown
    }
    if w.DisableWmm != nil {
        structMap["disable_wmm"] = w.DisableWmm
    }
    if w.DnsServerRewrite.IsValueSet() {
        if w.DnsServerRewrite.Value() != nil {
            structMap["dns_server_rewrite"] = w.DnsServerRewrite.Value().toMap()
        } else {
            structMap["dns_server_rewrite"] = nil
        }
    }
    if w.Dtim != nil {
        structMap["dtim"] = w.Dtim
    }
    if w.DynamicPsk.IsValueSet() {
        if w.DynamicPsk.Value() != nil {
            structMap["dynamic_psk"] = w.DynamicPsk.Value().toMap()
        } else {
            structMap["dynamic_psk"] = nil
        }
    }
    if w.DynamicVlan.IsValueSet() {
        if w.DynamicVlan.Value() != nil {
            structMap["dynamic_vlan"] = w.DynamicVlan.Value().toMap()
        } else {
            structMap["dynamic_vlan"] = nil
        }
    }
    if w.EnableLocalKeycaching != nil {
        structMap["enable_local_keycaching"] = w.EnableLocalKeycaching
    }
    if w.EnableWirelessBridging != nil {
        structMap["enable_wireless_bridging"] = w.EnableWirelessBridging
    }
    if w.EnableWirelessBridgingDhcpTracking != nil {
        structMap["enable_wireless_bridging_dhcp_tracking"] = w.EnableWirelessBridgingDhcpTracking
    }
    if w.Enabled != nil {
        structMap["enabled"] = w.Enabled
    }
    if w.FastDot1xTimers != nil {
        structMap["fast_dot1x_timers"] = w.FastDot1xTimers
    }
    if w.ForSite != nil {
        structMap["for_site"] = w.ForSite
    }
    if w.HideSsid != nil {
        structMap["hide_ssid"] = w.HideSsid
    }
    if w.HostnameIe != nil {
        structMap["hostname_ie"] = w.HostnameIe
    }
    if w.Hotspot20 != nil {
        structMap["hotspot20"] = w.Hotspot20.toMap()
    }
    if w.Id != nil {
        structMap["id"] = w.Id
    }
    if w.InjectDhcpOption82 != nil {
        structMap["inject_dhcp_option_82"] = w.InjectDhcpOption82.toMap()
    }
    if w.Interface != nil {
        structMap["interface"] = w.Interface
    }
    if w.Isolation != nil {
        structMap["isolation"] = w.Isolation
    }
    if w.L2Isolation != nil {
        structMap["l2_isolation"] = w.L2Isolation
    }
    if w.LegacyOverds != nil {
        structMap["legacy_overds"] = w.LegacyOverds
    }
    if w.LimitBcast != nil {
        structMap["limit_bcast"] = w.LimitBcast
    }
    if w.LimitProbeResponse != nil {
        structMap["limit_probe_response"] = w.LimitProbeResponse
    }
    if w.MaxIdletime != nil {
        structMap["max_idletime"] = w.MaxIdletime
    }
    if w.MaxNumClients != nil {
        structMap["max_num_clients"] = w.MaxNumClients
    }
    if w.MistNac != nil {
        structMap["mist_nac"] = w.MistNac.toMap()
    }
    if w.ModifiedTime != nil {
        structMap["modified_time"] = w.ModifiedTime
    }
    if w.MspId != nil {
        structMap["msp_id"] = w.MspId
    }
    if w.MxtunnelId != nil {
        structMap["mxtunnel_id"] = w.MxtunnelId
    }
    if w.MxtunnelIds != nil {
        structMap["mxtunnel_ids"] = w.MxtunnelIds
    }
    if w.MxtunnelName != nil {
        structMap["mxtunnel_name"] = w.MxtunnelName
    }
    if w.NoStaticDns != nil {
        structMap["no_static_dns"] = w.NoStaticDns
    }
    if w.NoStaticIp != nil {
        structMap["no_static_ip"] = w.NoStaticIp
    }
    if w.OrgId != nil {
        structMap["org_id"] = w.OrgId
    }
    if w.Portal != nil {
        structMap["portal"] = w.Portal.toMap()
    }
    if w.PortalAllowedHostnames != nil {
        structMap["portal_allowed_hostnames"] = w.PortalAllowedHostnames
    }
    if w.PortalAllowedSubnets != nil {
        structMap["portal_allowed_subnets"] = w.PortalAllowedSubnets
    }
    if w.PortalApiSecret.IsValueSet() {
        if w.PortalApiSecret.Value() != nil {
            structMap["portal_api_secret"] = w.PortalApiSecret.Value()
        } else {
            structMap["portal_api_secret"] = nil
        }
    }
    if w.PortalDeniedHostnames != nil {
        structMap["portal_denied_hostnames"] = w.PortalDeniedHostnames
    }
    if w.PortalImage.IsValueSet() {
        if w.PortalImage.Value() != nil {
            structMap["portal_image"] = w.PortalImage.Value()
        } else {
            structMap["portal_image"] = nil
        }
    }
    if w.PortalSsoUrl.IsValueSet() {
        if w.PortalSsoUrl.Value() != nil {
            structMap["portal_sso_url"] = w.PortalSsoUrl.Value()
        } else {
            structMap["portal_sso_url"] = nil
        }
    }
    if w.PortalTemplateUrl.IsValueSet() {
        if w.PortalTemplateUrl.Value() != nil {
            structMap["portal_template_url"] = w.PortalTemplateUrl.Value()
        } else {
            structMap["portal_template_url"] = nil
        }
    }
    if w.Qos != nil {
        structMap["qos"] = w.Qos.toMap()
    }
    if w.Radsec != nil {
        structMap["radsec"] = w.Radsec.toMap()
    }
    if w.Rateset != nil {
        structMap["rateset"] = w.Rateset
    }
    if w.ReconnectClientsWhenRoamingMxcluster != nil {
        structMap["reconnect_clients_when_roaming_mxcluster"] = w.ReconnectClientsWhenRoamingMxcluster
    }
    if w.RoamMode != nil {
        structMap["roam_mode"] = w.RoamMode
    }
    if w.Schedule != nil {
        structMap["schedule"] = w.Schedule.toMap()
    }
    if w.SiteId != nil {
        structMap["site_id"] = w.SiteId
    }
    if w.SleExcluded != nil {
        structMap["sle_excluded"] = w.SleExcluded
    }
    structMap["ssid"] = w.Ssid
    if w.TemplateId.IsValueSet() {
        if w.TemplateId.Value() != nil {
            structMap["template_id"] = w.TemplateId.Value()
        } else {
            structMap["template_id"] = nil
        }
    }
    if w.Thumbnail.IsValueSet() {
        if w.Thumbnail.Value() != nil {
            structMap["thumbnail"] = w.Thumbnail.Value()
        } else {
            structMap["thumbnail"] = nil
        }
    }
    if w.UseEapolV1 != nil {
        structMap["use_eapol_v1"] = w.UseEapolV1
    }
    if w.VlanEnabled != nil {
        structMap["vlan_enabled"] = w.VlanEnabled
    }
    if w.VlanId.IsValueSet() {
        if w.VlanId.Value() != nil {
            structMap["vlan_id"] = w.VlanId.Value().toMap()
        } else {
            structMap["vlan_id"] = nil
        }
    }
    if w.VlanIds != nil {
        structMap["vlan_ids"] = w.VlanIds.toMap()
    }
    if w.VlanPooling != nil {
        structMap["vlan_pooling"] = w.VlanPooling
    }
    if w.WlanLimitDown != nil {
        structMap["wlan_limit_down"] = w.WlanLimitDown.toMap()
    }
    if w.WlanLimitDownEnabled != nil {
        structMap["wlan_limit_down_enabled"] = w.WlanLimitDownEnabled
    }
    if w.WlanLimitUp != nil {
        structMap["wlan_limit_up"] = w.WlanLimitUp.toMap()
    }
    if w.WlanLimitUpEnabled != nil {
        structMap["wlan_limit_up_enabled"] = w.WlanLimitUpEnabled
    }
    if w.WxtagIds.IsValueSet() {
        if w.WxtagIds.Value() != nil {
            structMap["wxtag_ids"] = w.WxtagIds.Value()
        } else {
            structMap["wxtag_ids"] = nil
        }
    }
    if w.WxtunnelId.IsValueSet() {
        if w.WxtunnelId.Value() != nil {
            structMap["wxtunnel_id"] = w.WxtunnelId.Value()
        } else {
            structMap["wxtunnel_id"] = nil
        }
    }
    if w.WxtunnelRemoteId.IsValueSet() {
        if w.WxtunnelRemoteId.Value() != nil {
            structMap["wxtunnel_remote_id"] = w.WxtunnelRemoteId.Value()
        } else {
            structMap["wxtunnel_remote_id"] = nil
        }
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for Wlan.
// It customizes the JSON unmarshaling process for Wlan objects.
func (w *Wlan) UnmarshalJSON(input []byte) error {
    var temp tempWlan
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "acct_immediate_update", "acct_interim_interval", "acct_servers", "airwatch", "allow_ipv6_ndp", "allow_mdns", "allow_ssdp", "ap_ids", "app_limit", "app_qos", "apply_to", "arp_filter", "auth", "auth_server_selection", "auth_servers", "auth_servers_nas_id", "auth_servers_nas_ip", "auth_servers_retries", "auth_servers_timeout", "band", "band_steer", "band_steer_force_band5", "bands", "block_blacklist_clients", "bonjour", "cisco_cwa", "client_limit_down", "client_limit_down_enabled", "client_limit_up", "client_limit_up_enabled", "coa_servers", "created_time", "disable_11ax", "disable_11be", "disable_ht_vht_rates", "disable_uapsd", "disable_v1_roam_notify", "disable_v2_roam_notify", "disable_when_gateway_unreachable", "disable_when_mxtunnel_down", "disable_wmm", "dns_server_rewrite", "dtim", "dynamic_psk", "dynamic_vlan", "enable_local_keycaching", "enable_wireless_bridging", "enable_wireless_bridging_dhcp_tracking", "enabled", "fast_dot1x_timers", "for_site", "hide_ssid", "hostname_ie", "hotspot20", "id", "inject_dhcp_option_82", "interface", "isolation", "l2_isolation", "legacy_overds", "limit_bcast", "limit_probe_response", "max_idletime", "max_num_clients", "mist_nac", "modified_time", "msp_id", "mxtunnel_id", "mxtunnel_ids", "mxtunnel_name", "no_static_dns", "no_static_ip", "org_id", "portal", "portal_allowed_hostnames", "portal_allowed_subnets", "portal_api_secret", "portal_denied_hostnames", "portal_image", "portal_sso_url", "portal_template_url", "qos", "radsec", "rateset", "reconnect_clients_when_roaming_mxcluster", "roam_mode", "schedule", "site_id", "sle_excluded", "ssid", "template_id", "thumbnail", "use_eapol_v1", "vlan_enabled", "vlan_id", "vlan_ids", "vlan_pooling", "wlan_limit_down", "wlan_limit_down_enabled", "wlan_limit_up", "wlan_limit_up_enabled", "wxtag_ids", "wxtunnel_id", "wxtunnel_remote_id")
    if err != nil {
    	return err
    }
    w.AdditionalProperties = additionalProperties
    
    w.AcctImmediateUpdate = temp.AcctImmediateUpdate
    w.AcctInterimInterval = temp.AcctInterimInterval
    w.AcctServers = temp.AcctServers
    w.Airwatch = temp.Airwatch
    w.AllowIpv6Ndp = temp.AllowIpv6Ndp
    w.AllowMdns = temp.AllowMdns
    w.AllowSsdp = temp.AllowSsdp
    w.ApIds = temp.ApIds
    w.AppLimit = temp.AppLimit
    w.AppQos = temp.AppQos
    w.ApplyTo = temp.ApplyTo
    w.ArpFilter = temp.ArpFilter
    w.Auth = temp.Auth
    w.AuthServerSelection = temp.AuthServerSelection
    w.AuthServers = temp.AuthServers
    w.AuthServersNasId = temp.AuthServersNasId
    w.AuthServersNasIp = temp.AuthServersNasIp
    w.AuthServersRetries = temp.AuthServersRetries
    w.AuthServersTimeout = temp.AuthServersTimeout
    w.Band = temp.Band
    w.BandSteer = temp.BandSteer
    w.BandSteerForceBand5 = temp.BandSteerForceBand5
    w.Bands = temp.Bands
    w.BlockBlacklistClients = temp.BlockBlacklistClients
    w.Bonjour = temp.Bonjour
    w.CiscoCwa = temp.CiscoCwa
    w.ClientLimitDown = temp.ClientLimitDown
    w.ClientLimitDownEnabled = temp.ClientLimitDownEnabled
    w.ClientLimitUp = temp.ClientLimitUp
    w.ClientLimitUpEnabled = temp.ClientLimitUpEnabled
    w.CoaServers = temp.CoaServers
    w.CreatedTime = temp.CreatedTime
    w.Disable11ax = temp.Disable11ax
    w.Disable11be = temp.Disable11be
    w.DisableHtVhtRates = temp.DisableHtVhtRates
    w.DisableUapsd = temp.DisableUapsd
    w.DisableV1RoamNotify = temp.DisableV1RoamNotify
    w.DisableV2RoamNotify = temp.DisableV2RoamNotify
    w.DisableWhenGatewayUnreachable = temp.DisableWhenGatewayUnreachable
    w.DisableWhenMxtunnelDown = temp.DisableWhenMxtunnelDown
    w.DisableWmm = temp.DisableWmm
    w.DnsServerRewrite = temp.DnsServerRewrite
    w.Dtim = temp.Dtim
    w.DynamicPsk = temp.DynamicPsk
    w.DynamicVlan = temp.DynamicVlan
    w.EnableLocalKeycaching = temp.EnableLocalKeycaching
    w.EnableWirelessBridging = temp.EnableWirelessBridging
    w.EnableWirelessBridgingDhcpTracking = temp.EnableWirelessBridgingDhcpTracking
    w.Enabled = temp.Enabled
    w.FastDot1xTimers = temp.FastDot1xTimers
    w.ForSite = temp.ForSite
    w.HideSsid = temp.HideSsid
    w.HostnameIe = temp.HostnameIe
    w.Hotspot20 = temp.Hotspot20
    w.Id = temp.Id
    w.InjectDhcpOption82 = temp.InjectDhcpOption82
    w.Interface = temp.Interface
    w.Isolation = temp.Isolation
    w.L2Isolation = temp.L2Isolation
    w.LegacyOverds = temp.LegacyOverds
    w.LimitBcast = temp.LimitBcast
    w.LimitProbeResponse = temp.LimitProbeResponse
    w.MaxIdletime = temp.MaxIdletime
    w.MaxNumClients = temp.MaxNumClients
    w.MistNac = temp.MistNac
    w.ModifiedTime = temp.ModifiedTime
    w.MspId = temp.MspId
    w.MxtunnelId = temp.MxtunnelId
    w.MxtunnelIds = temp.MxtunnelIds
    w.MxtunnelName = temp.MxtunnelName
    w.NoStaticDns = temp.NoStaticDns
    w.NoStaticIp = temp.NoStaticIp
    w.OrgId = temp.OrgId
    w.Portal = temp.Portal
    w.PortalAllowedHostnames = temp.PortalAllowedHostnames
    w.PortalAllowedSubnets = temp.PortalAllowedSubnets
    w.PortalApiSecret = temp.PortalApiSecret
    w.PortalDeniedHostnames = temp.PortalDeniedHostnames
    w.PortalImage = temp.PortalImage
    w.PortalSsoUrl = temp.PortalSsoUrl
    w.PortalTemplateUrl = temp.PortalTemplateUrl
    w.Qos = temp.Qos
    w.Radsec = temp.Radsec
    w.Rateset = temp.Rateset
    w.ReconnectClientsWhenRoamingMxcluster = temp.ReconnectClientsWhenRoamingMxcluster
    w.RoamMode = temp.RoamMode
    w.Schedule = temp.Schedule
    w.SiteId = temp.SiteId
    w.SleExcluded = temp.SleExcluded
    w.Ssid = *temp.Ssid
    w.TemplateId = temp.TemplateId
    w.Thumbnail = temp.Thumbnail
    w.UseEapolV1 = temp.UseEapolV1
    w.VlanEnabled = temp.VlanEnabled
    w.VlanId = temp.VlanId
    w.VlanIds = temp.VlanIds
    w.VlanPooling = temp.VlanPooling
    w.WlanLimitDown = temp.WlanLimitDown
    w.WlanLimitDownEnabled = temp.WlanLimitDownEnabled
    w.WlanLimitUp = temp.WlanLimitUp
    w.WlanLimitUpEnabled = temp.WlanLimitUpEnabled
    w.WxtagIds = temp.WxtagIds
    w.WxtunnelId = temp.WxtunnelId
    w.WxtunnelRemoteId = temp.WxtunnelRemoteId
    return nil
}

// tempWlan is a temporary struct used for validating the fields of Wlan.
type tempWlan  struct {
    AcctImmediateUpdate                  *bool                            `json:"acct_immediate_update,omitempty"`
    AcctInterimInterval                  *int                             `json:"acct_interim_interval,omitempty"`
    AcctServers                          []RadiusAcctServer               `json:"acct_servers,omitempty"`
    Airwatch                             *WlanAirwatch                    `json:"airwatch,omitempty"`
    AllowIpv6Ndp                         *bool                            `json:"allow_ipv6_ndp,omitempty"`
    AllowMdns                            *bool                            `json:"allow_mdns,omitempty"`
    AllowSsdp                            *bool                            `json:"allow_ssdp,omitempty"`
    ApIds                                Optional[[]uuid.UUID]            `json:"ap_ids"`
    AppLimit                             *WlanAppLimit                    `json:"app_limit,omitempty"`
    AppQos                               *WlanAppQos                      `json:"app_qos,omitempty"`
    ApplyTo                              *WlanApplyToEnum                 `json:"apply_to,omitempty"`
    ArpFilter                            *bool                            `json:"arp_filter,omitempty"`
    Auth                                 *WlanAuth                        `json:"auth,omitempty"`
    AuthServerSelection                  *WlanAuthServerSelectionEnum     `json:"auth_server_selection,omitempty"`
    AuthServers                          []RadiusAuthServer               `json:"auth_servers,omitempty"`
    AuthServersNasId                     Optional[string]                 `json:"auth_servers_nas_id"`
    AuthServersNasIp                     Optional[string]                 `json:"auth_servers_nas_ip"`
    AuthServersRetries                   *int                             `json:"auth_servers_retries,omitempty"`
    AuthServersTimeout                   *int                             `json:"auth_servers_timeout,omitempty"`
    Band                                 *string                          `json:"band,omitempty"`
    BandSteer                            *bool                            `json:"band_steer,omitempty"`
    BandSteerForceBand5                  *bool                            `json:"band_steer_force_band5,omitempty"`
    Bands                                []Dot11BandEnum                  `json:"bands,omitempty"`
    BlockBlacklistClients                *bool                            `json:"block_blacklist_clients,omitempty"`
    Bonjour                              *WlanBonjour                     `json:"bonjour,omitempty"`
    CiscoCwa                             *WlanCiscoCwa                    `json:"cisco_cwa,omitempty"`
    ClientLimitDown                      *WlanLimit                       `json:"client_limit_down,omitempty"`
    ClientLimitDownEnabled               *bool                            `json:"client_limit_down_enabled,omitempty"`
    ClientLimitUp                        *WlanLimit                       `json:"client_limit_up,omitempty"`
    ClientLimitUpEnabled                 *bool                            `json:"client_limit_up_enabled,omitempty"`
    CoaServers                           []CoaServer                      `json:"coa_servers,omitempty"`
    CreatedTime                          *float64                         `json:"created_time,omitempty"`
    Disable11ax                          *bool                            `json:"disable_11ax,omitempty"`
    Disable11be                          *bool                            `json:"disable_11be,omitempty"`
    DisableHtVhtRates                    *bool                            `json:"disable_ht_vht_rates,omitempty"`
    DisableUapsd                         *bool                            `json:"disable_uapsd,omitempty"`
    DisableV1RoamNotify                  *bool                            `json:"disable_v1_roam_notify,omitempty"`
    DisableV2RoamNotify                  *bool                            `json:"disable_v2_roam_notify,omitempty"`
    DisableWhenGatewayUnreachable        *bool                            `json:"disable_when_gateway_unreachable,omitempty"`
    DisableWhenMxtunnelDown              *bool                            `json:"disable_when_mxtunnel_down,omitempty"`
    DisableWmm                           *bool                            `json:"disable_wmm,omitempty"`
    DnsServerRewrite                     Optional[WlanDnsServerRewrite]   `json:"dns_server_rewrite"`
    Dtim                                 *int                             `json:"dtim,omitempty"`
    DynamicPsk                           Optional[WlanDynamicPsk]         `json:"dynamic_psk"`
    DynamicVlan                          Optional[WlanDynamicVlan]        `json:"dynamic_vlan"`
    EnableLocalKeycaching                *bool                            `json:"enable_local_keycaching,omitempty"`
    EnableWirelessBridging               *bool                            `json:"enable_wireless_bridging,omitempty"`
    EnableWirelessBridgingDhcpTracking   *bool                            `json:"enable_wireless_bridging_dhcp_tracking,omitempty"`
    Enabled                              *bool                            `json:"enabled,omitempty"`
    FastDot1xTimers                      *bool                            `json:"fast_dot1x_timers,omitempty"`
    ForSite                              *bool                            `json:"for_site,omitempty"`
    HideSsid                             *bool                            `json:"hide_ssid,omitempty"`
    HostnameIe                           *bool                            `json:"hostname_ie,omitempty"`
    Hotspot20                            *WlanHotspot20                   `json:"hotspot20,omitempty"`
    Id                                   *uuid.UUID                       `json:"id,omitempty"`
    InjectDhcpOption82                   *WlanInjectDhcpOption82          `json:"inject_dhcp_option_82,omitempty"`
    Interface                            *WlanInterfaceEnum               `json:"interface,omitempty"`
    Isolation                            *bool                            `json:"isolation,omitempty"`
    L2Isolation                          *bool                            `json:"l2_isolation,omitempty"`
    LegacyOverds                         *bool                            `json:"legacy_overds,omitempty"`
    LimitBcast                           *bool                            `json:"limit_bcast,omitempty"`
    LimitProbeResponse                   *bool                            `json:"limit_probe_response,omitempty"`
    MaxIdletime                          *int                             `json:"max_idletime,omitempty"`
    MaxNumClients                        *int                             `json:"max_num_clients,omitempty"`
    MistNac                              *WlanMistNac                     `json:"mist_nac,omitempty"`
    ModifiedTime                         *float64                         `json:"modified_time,omitempty"`
    MspId                                *uuid.UUID                       `json:"msp_id,omitempty"`
    MxtunnelId                           *uuid.UUID                       `json:"mxtunnel_id,omitempty"`
    MxtunnelIds                          []string                         `json:"mxtunnel_ids,omitempty"`
    MxtunnelName                         []string                         `json:"mxtunnel_name,omitempty"`
    NoStaticDns                          *bool                            `json:"no_static_dns,omitempty"`
    NoStaticIp                           *bool                            `json:"no_static_ip,omitempty"`
    OrgId                                *uuid.UUID                       `json:"org_id,omitempty"`
    Portal                               *WlanPortal                      `json:"portal,omitempty"`
    PortalAllowedHostnames               []string                         `json:"portal_allowed_hostnames,omitempty"`
    PortalAllowedSubnets                 []string                         `json:"portal_allowed_subnets,omitempty"`
    PortalApiSecret                      Optional[string]                 `json:"portal_api_secret"`
    PortalDeniedHostnames                []string                         `json:"portal_denied_hostnames,omitempty"`
    PortalImage                          Optional[string]                 `json:"portal_image"`
    PortalSsoUrl                         Optional[string]                 `json:"portal_sso_url"`
    PortalTemplateUrl                    Optional[string]                 `json:"portal_template_url"`
    Qos                                  *WlanQos                         `json:"qos,omitempty"`
    Radsec                               *Radsec                          `json:"radsec,omitempty"`
    Rateset                              map[string]WlanDatarates         `json:"rateset,omitempty"`
    ReconnectClientsWhenRoamingMxcluster *bool                            `json:"reconnect_clients_when_roaming_mxcluster,omitempty"`
    RoamMode                             *WlanRoamModeEnum                `json:"roam_mode,omitempty"`
    Schedule                             *WlanSchedule                    `json:"schedule,omitempty"`
    SiteId                               *uuid.UUID                       `json:"site_id,omitempty"`
    SleExcluded                          *bool                            `json:"sle_excluded,omitempty"`
    Ssid                                 *string                          `json:"ssid"`
    TemplateId                           Optional[uuid.UUID]              `json:"template_id"`
    Thumbnail                            Optional[string]                 `json:"thumbnail"`
    UseEapolV1                           *bool                            `json:"use_eapol_v1,omitempty"`
    VlanEnabled                          *bool                            `json:"vlan_enabled,omitempty"`
    VlanId                               Optional[WlanVlanIdWithVariable] `json:"vlan_id"`
    VlanIds                              *WlanVlanIds                     `json:"vlan_ids,omitempty"`
    VlanPooling                          *bool                            `json:"vlan_pooling,omitempty"`
    WlanLimitDown                        *WlanLimit                       `json:"wlan_limit_down,omitempty"`
    WlanLimitDownEnabled                 *bool                            `json:"wlan_limit_down_enabled,omitempty"`
    WlanLimitUp                          *WlanLimit                       `json:"wlan_limit_up,omitempty"`
    WlanLimitUpEnabled                   *bool                            `json:"wlan_limit_up_enabled,omitempty"`
    WxtagIds                             Optional[[]uuid.UUID]            `json:"wxtag_ids"`
    WxtunnelId                           Optional[string]                 `json:"wxtunnel_id"`
    WxtunnelRemoteId                     Optional[string]                 `json:"wxtunnel_remote_id"`
}

func (w *tempWlan) validate() error {
    var errs []string
    if w.Ssid == nil {
        errs = append(errs, "required field `ssid` is missing for type `wlan`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
