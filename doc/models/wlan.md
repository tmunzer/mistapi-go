
# Wlan

**Note**: portal_template will be forked out of wlan objects soon. To fetch portal_template, please query portal_template_url. To update portal_template, use Wlan Portal Template.

*This model accepts additional fields of type interface{}.*

## Structure

`Wlan`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AcctImmediateUpdate` | `*bool` | Optional | Enable coa-immediate-update and address-change-immediate-update on the access profile.<br><br>**Default**: `false` |
| `AcctInterimInterval` | `*int` | Optional | How frequently should interim accounting be reported, 60-65535. default is 0 (use one specified in Access-Accept request from RADIUS Server). Very frequent messages can affect the performance of the RADIUS server, 600 and up is recommended when enabled<br><br>**Default**: `0`<br><br>**Constraints**: `>= 0`, `<= 65535` |
| `AcctServers` | [`[]models.RadiusAcctServer`](../../doc/models/radius-acct-server.md) | Optional | List of RADIUS accounting servers, optional, order matters where the first one is treated as primary |
| `Airwatch` | [`*models.WlanAirwatch`](../../doc/models/wlan-airwatch.md) | Optional | AirWatch integration settings for the WLAN |
| `AllowIpv6Ndp` | `*bool` | Optional | Only applicable when `limit_bcast`==`true`, which allows or disallows ipv6 Neighbor Discovery packets to go through<br><br>**Default**: `true` |
| `AllowMdns` | `*bool` | Optional | Only applicable when `limit_bcast`==`true`, which allows mDNS / Bonjour packets to go through<br><br>**Default**: `false` |
| `AllowSsdp` | `*bool` | Optional | Only applicable when `limit_bcast`==`true`, which allows SSDP<br><br>**Default**: `false` |
| `ApIds` | `models.Optional[[]uuid.UUID]` | Optional | List of device ids |
| `AppLimit` | [`*models.WlanAppLimit`](../../doc/models/wlan-app-limit.md) | Optional | Bandwidth limiting for apps (applies to up/down) |
| `AppQos` | [`*models.WlanAppQos`](../../doc/models/wlan-app-qos.md) | Optional | APP qos wlan settings |
| `ApplyTo` | [`*models.WlanApplyToEnum`](../../doc/models/wlan-apply-to-enum.md) | Optional | enum: `aps`, `site`, `wxtags` |
| `ArpFilter` | `*bool` | Optional | Whether to enable smart arp filter<br><br>**Default**: `false` |
| `Auth` | [`*models.WlanAuth`](../../doc/models/wlan-auth.md) | Optional | WLAN client authentication settings |
| `AuthServerSelection` | [`*models.WlanAuthServerSelectionEnum`](../../doc/models/wlan-auth-server-selection-enum.md) | Optional | When ordered, AP will prefer and go back to the first server if possible. enum: `ordered`, `unordered`<br><br>**Default**: `"ordered"` |
| `AuthServers` | [`[]models.RadiusAuthServer`](../../doc/models/radius-auth-server.md) | Optional | List of RADIUS authentication servers, at least one is needed if `auth type`==`eap`, order matters where the first one is treated as primary |
| `AuthServersNasId` | `models.Optional[string]` | Optional | Optional, up to 48 bytes, will be dynamically generated if not provided. used only for authentication servers |
| `AuthServersNasIp` | `models.Optional[string]` | Optional | Optional, NAS-IP-ADDRESS to use |
| `AuthServersRetries` | `*int` | Optional | RADIUS auth session retries. Following fast timers are set if "fast_dot1x_timers" knob is enabled. ‘retries’ are set to value of auth_servers_retries. ‘max-requests’ is also set when setting auth_servers_retries and is set to default value to 3.<br><br>**Default**: `2` |
| `AuthServersTimeout` | `*int` | Optional | RADIUS auth session timeout. Following fast timers are set if "fast_dot1x_timers" knob is enabled. ‘quite-period’  and ‘transmit-period’ are set to half the value of auth_servers_timeout. ‘supplicant-timeout’ is also set when setting auth_servers_timeout and is set to default value of 10.<br><br>**Default**: `5` |
| `Band` | `*string` | Optional | `band` is deprecated and kept for backward compatibility. Use `bands` instead |
| `BandSteer` | `*bool` | Optional | Whether to enable band_steering, this works only when band==both<br><br>**Default**: `false` |
| `BandSteerForceBand5` | `*bool` | Optional | Force dual_band capable client to connect to 5G<br><br>**Default**: `false` |
| `Bands` | [`[]models.Dot11BandEnum`](../../doc/models/dot-11-band-enum.md) | Optional | List of radios that the wlan should apply to. enum: `24`, `5`, `5-dedicated`, `5-selectable`, `6`, `6-dedicated`, `6-selectable` |
| `BlockBlacklistClients` | `*bool` | Optional | Whether to block the clients in the blacklist (up to first 256 macs)<br><br>**Default**: `false` |
| `Bonjour` | [`*models.WlanBonjour`](../../doc/models/wlan-bonjour.md) | Optional | Bonjour gateway wlan settings |
| `CiscoCwa` | [`*models.WlanCiscoCwa`](../../doc/models/wlan-cisco-cwa.md) | Optional | Cisco CWA (central web authentication) required RADIUS with COA in order to work. See CWA: https://www.cisco.com/c/en/us/support/docs/security/identity-services-engine/115732-central-web-auth-00.html |
| `ClientLimitDown` | [`*models.WlanLimit`](../../doc/models/containers/wlan-limit.md) | Optional | In kbps, value from 1 to 999000 |
| `ClientLimitDownEnabled` | `*bool` | Optional | If downlink limiting per-client is enabled<br><br>**Default**: `false` |
| `ClientLimitUp` | [`*models.WlanLimit`](../../doc/models/containers/wlan-limit.md) | Optional | In kbps, value from 1 to 999000 |
| `ClientLimitUpEnabled` | `*bool` | Optional | If uplink limiting per-client is enabled<br><br>**Default**: `false` |
| `CoaServers` | [`[]models.CoaServer`](../../doc/models/coa-server.md) | Optional | List of COA (change of authorization) servers, optional |
| `CreatedTime` | `*float64` | Optional, Read-only | When the object has been created, in epoch |
| `Disable11ax` | `*bool` | Optional | Some old WLAN drivers may not be compatible<br><br>**Default**: `false` |
| `Disable11be` | `*bool` | Optional | To disable Wi-Fi 7 EHT IEs<br><br>**Default**: `false` |
| `DisableHtVhtRates` | `*bool` | Optional | To disable ht or vht rates<br><br>**Default**: `false` |
| `DisableMessageAuthenticatorCheck` | `*bool` | Optional | whether to disable Message-Authenticator Check, which is used to verify the integrity of RADIUS messages, default is false (i.e. for better security)<br><br>**Default**: `false` |
| `DisableUapsd` | `*bool` | Optional | Whether to disable U-APSD<br><br>**Default**: `false` |
| `DisableV1RoamNotify` | `*bool` | Optional | Disable sending v2 roam notification messages<br><br>**Default**: `false` |
| `DisableV2RoamNotify` | `*bool` | Optional | Disable sending v2 roam notification messages<br><br>**Default**: `false` |
| `DisableWhenGatewayUnreachable` | `*bool` | Optional | When any of the following is true, this WLAN will be disabled<br><br>* cannot get IP<br>* cannot obtain default gateway<br>* cannot reach default gateway<br><br>**Default**: `false` |
| `DisableWhenMxtunnelDown` | `*bool` | Optional | Whether to disable this WLAN when the configured Mist tunnel is down<br><br>**Default**: `false` |
| `DisableWmm` | `*bool` | Optional | Whether to disable WMM<br><br>**Default**: `false` |
| `DnsServerRewrite` | [`models.Optional[models.WlanDnsServerRewrite]`](../../doc/models/wlan-dns-server-rewrite.md) | Optional | For radius_group-based DNS server (rewrite DNS request depending on the Group RADIUS server returns) |
| `Dtim` | `*int` | Optional | Delivery Traffic Indication Message interval for this WLAN<br><br>**Default**: `2` |
| `DynamicPsk` | [`models.Optional[models.WlanDynamicPsk]`](../../doc/models/wlan-dynamic-psk.md) | Optional | For dynamic PSK where we get per_user PSK from RADIUS. dynamic_psk allows PSK to be selected at runtime depending on context (wlan/site/user/...) thus following configurations are assumed (currently)<br><br>* PSK will come from RADIUS server<br>* AP sends client MAC as username and password (i.e. `enable_mac_auth` is assumed)<br>* AP sends BSSID:SSID as Caller-Station-ID<br>* `auth_servers` is required<br>* PSK will come from cloud WLC if source is cloud_psks<br>* default_psk will be used if cloud WLC is not available<br>* `multi_psk_only` and `psk` is ignored<br>* `pairwise` can only be wpa2-ccmp (for now, wpa3 support on the roadmap) |
| `DynamicVlan` | [`models.Optional[models.WlanDynamicVlan]`](../../doc/models/wlan-dynamic-vlan.md) | Optional | Dynamic VLAN assignment settings for 802.1X WLAN authentication |
| `EnableFtm` | `*bool` | Optional | Enable FTM (Fine-Time Measurement, 802.11mc); configures the AP as an FTM Responder (target), allowing clients to perform ranging requests against it<br><br>**Default**: `false` |
| `EnableLocalKeycaching` | `*bool` | Optional | Enable AP-AP keycaching via multicast<br><br>**Default**: `false` |
| `EnableWirelessBridging` | `*bool` | Optional | By default, we'd inspect all DHCP packets and drop those unrelated to the wireless client itself in the case where client is a wireless bridge (DHCP packets for other MACs will need to be forwarded), wireless_bridging can be enabled<br><br>**Default**: `false` |
| `EnableWirelessBridgingDhcpTracking` | `*bool` | Optional | If the client bridge is doing DHCP on behalf of other devices (L2-NAT), enable dhcp_tracking will cut down DHCP response packets to be forwarded to wireless<br><br>**Default**: `false` |
| `Enabled` | `*bool` | Optional | If this wlan is enabled<br><br>**Default**: `true` |
| `FastDot1xTimers` | `*bool` | Optional | If set to true, sets default fast-timers with values calculated from ‘auth_servers_timeout’ and ‘auth_server_retries’ .<br><br>**Default**: `false` |
| `ForSite` | `*bool` | Optional, Read-only | Whether this WLAN record is scoped to a site |
| `HideSsid` | `*bool` | Optional | Whether to hide SSID in beacon<br><br>**Default**: `false` |
| `HostnameIe` | `*bool` | Optional | Include hostname inside IE in AP beacons / probe responses<br><br>**Default**: `false` |
| `Hotspot20` | [`*models.WlanHotspot20`](../../doc/models/wlan-hotspot-20.md) | Optional | Hotspot 2.0 WLAN settings |
| `Id` | `*uuid.UUID` | Optional, Read-only | Unique ID of the object instance in the Mist Organization |
| `InjectDhcpOption82` | [`*models.WlanInjectDhcpOption82`](../../doc/models/wlan-inject-dhcp-option-82.md) | Optional | DHCP Option 82 injection settings for a WLAN |
| `Interface` | [`*models.WlanInterfaceEnum`](../../doc/models/wlan-interface-enum.md) | Optional | where this WLAN will be connected to. enum: `all`, `eth0`, `eth1`, `eth2`, `eth3`, `mxtunnel`, `site_mxedge`, `wxtunnel`<br><br>**Default**: `"all"` |
| `Isolation` | `*bool` | Optional | Whether to stop clients to talk to each other<br><br>**Default**: `false` |
| `L2Isolation` | `*bool` | Optional | If isolation is enabled, whether to deny clients to talk to L2 on the LAN<br><br>**Default**: `false` |
| `LegacyOverds` | `*bool` | Optional | Legacy devices requires the Over-DS (for Fast BSS Transition) bit set (while our chip doesn’t support it). Warning! Enabling this will cause problem for iOS devices.<br><br>**Default**: `false` |
| `LimitBcast` | `*bool` | Optional | Whether to limit broadcast packets going to wireless (i.e. only allow certain bcast packets to go through)<br><br>**Default**: `false` |
| `LimitProbeResponse` | `*bool` | Optional | Limit probe response base on some heuristic rules<br><br>**Default**: `false` |
| `MaxIdletime` | `*int` | Optional | Max idle time in seconds<br><br>**Default**: `1800`<br><br>**Constraints**: `>= 60`, `<= 86400` |
| `MaxNumClients` | `*int` | Optional | Maximum number of client connected to the SSID. `0` means unlimited<br><br>**Default**: `0`<br><br>**Constraints**: `>= 0`, `<= 128` |
| `MistNac` | [`*models.WlanMistNac`](../../doc/models/wlan-mist-nac.md) | Optional | Mist NAC RADIUS settings for a WLAN |
| `ModifiedTime` | `*float64` | Optional, Read-only | When the object has been modified for the last time, in epoch |
| `MspId` | `*uuid.UUID` | Optional, Read-only | Managed service provider identifier |
| `MxtunnelId` | `*uuid.UUID` | Optional | (deprecated, use mxtunnel_ids instead) when `interface`==`mxtunnel`, id of the Mist Tunnel |
| `MxtunnelIds` | `[]string` | Optional | When `interface`=`mxtunnel`, id of the Mist Tunnel |
| `MxtunnelName` | `[]string` | Optional | When `interface`=`site_mxedge`, name of the mxtunnel that in mxtunnels under Site Setting |
| `NoStaticDns` | `*bool` | Optional | Whether to only allow client to use DNS that we’ve learned from DHCP response<br><br>**Default**: `false` |
| `NoStaticIp` | `*bool` | Optional | Whether to only allow client that we’ve learned from DHCP exchange to talk<br><br>**Default**: `false` |
| `OrgId` | `*uuid.UUID` | Optional, Read-only | Unique identifier of a Mist organization |
| `Portal` | [`*models.WlanPortal`](../../doc/models/wlan-portal.md) | Optional | Guest portal settings for the WLAN |
| `PortalAllowedHostnames` | `[]string` | Optional | List of hostnames without http(s):// (matched by substring) |
| `PortalAllowedSubnets` | `[]string` | Optional | Guest portal CIDR subnets that clients may reach before authorization |
| `PortalApiSecret` | `models.Optional[string]` | Optional | API secret (auto-generated) that can be used to sign guest authorization requests, only generated when auth is set to `external` |
| `PortalDeniedHostnames` | `[]string` | Optional | List of hostnames without http(s):// (matched by substring), this takes precedence over portal_allowed_hostnames |
| `PortalImage` | `models.Optional[string]` | Optional, Read-only | Url of portal background image |
| `PortalSsoUrl` | `models.Optional[string]` | Optional, Read-only | URL used in the SSO process, auto-generated when auth is set to `sso` |
| `PortalTemplateUrl` | `models.Optional[string]` | Optional, Read-only | N.B portal_template will be forked out of wlan objects soon. To fetch portal_template, please query portal_template_url. To update portal_template, use Wlan Portal Template. |
| `Qos` | [`*models.WlanQos`](../../doc/models/wlan-qos.md) | Optional | QoS override settings for WLAN client traffic |
| `Radsec` | [`*models.Radsec`](../../doc/models/radsec.md) | Optional | RadSec settings for sending RADIUS traffic over TLS |
| `Rateset` | [`map[string]models.WlanDatarates`](../../doc/models/wlan-datarates.md) | Optional | Property key is the RF band. enum: `24`, `5`, `6` |
| `ReconnectClientsWhenRoamingMxcluster` | `*bool` | Optional | When different mxcluster is on different subnet, we'd want to disconnect clients (so they'll reconnect and get new IPs)<br><br>**Default**: `false` |
| `RoamMode` | [`*models.WlanRoamModeEnum`](../../doc/models/wlan-roam-mode-enum.md) | Optional | enum: `11r`, `OKC`, `NONE`<br><br>**Default**: `"NONE"` |
| `Schedule` | [`*models.WlanSchedule`](../../doc/models/wlan-schedule.md) | Optional | WLAN operating schedule, default is disabled |
| `SiteId` | `*uuid.UUID` | Optional, Read-only | Unique identifier of a Mist site |
| `SleExcluded` | `*bool` | Optional | Whether to exclude this WLAN from SLE metrics<br><br>**Default**: `false` |
| `Ssid` | `string` | Required | Name of the SSID |
| `TemplateId` | `models.Optional[uuid.UUID]` | Optional | Identifier of the WLAN template associated with this WLAN |
| `Thumbnail` | `models.Optional[string]` | Optional, Read-only | Url of portal background image thumbnail |
| `UseEapolV1` | `*bool` | Optional | If `auth.type`==`eap` or `auth.type`==`psk`, should only be set for legacy client, such as pre-2004, 802.11b devices<br><br>**Default**: `false` |
| `VlanEnabled` | `*bool` | Optional | If vlan tagging is enabled<br><br>**Default**: `false` |
| `VlanId` | [`models.Optional[models.WlanVlanIdWithVariable]`](../../doc/models/containers/wlan-vlan-id-with-variable.md) | Optional | WLAN VLAN ID, either numeric, a variable string, or null |
| `VlanIds` | [`*models.WlanVlanIds`](../../doc/models/containers/wlan-vlan-ids.md) | Optional | WLAN VLAN pool IDs represented as either a comma-separated string or a list |
| `VlanPooling` | `*bool` | Optional | Requires `vlan_enabled`==`true` to be set to `true`. Vlan pooling allows AP to place client on different VLAN using a deterministic algorithm<br><br>**Default**: `false` |
| `WlanLimitDown` | [`*models.WlanLimit`](../../doc/models/containers/wlan-limit.md) | Optional | In kbps, value from 1 to 999000 |
| `WlanLimitDownEnabled` | `*bool` | Optional | If downlink limiting for whole wlan is enabled<br><br>**Default**: `false` |
| `WlanLimitUp` | [`*models.WlanLimit`](../../doc/models/containers/wlan-limit.md) | Optional | In kbps, value from 1 to 999000 |
| `WlanLimitUpEnabled` | `*bool` | Optional | If uplink limiting for whole wlan is enabled<br><br>**Default**: `false` |
| `WxtagIds` | `models.Optional[[]uuid.UUID]` | Optional | Identifiers of WxLAN tags used when `apply_to`==`wxtags` |
| `WxtunnelId` | `models.Optional[string]` | Optional | When `interface`=`wxtunnel`, id of the WXLAN Tunnel |
| `WxtunnelRemoteId` | `models.Optional[string]` | Optional | When `interface`=`wxtunnel`, remote tunnel identifier |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    wlan := models.Wlan{
        AcctImmediateUpdate:                  models.ToPointer(false),
        AcctInterimInterval:                  models.ToPointer(0),
        AcctServers:                          []models.RadiusAcctServer{
            models.RadiusAcctServer{
                Host:                 "host4",
                KeywrapEnabled:       models.ToPointer(false),
                KeywrapFormat:        models.ToPointer(models.RadiusKeywrapFormatEnum_ASCII),
                KeywrapKek:           models.ToPointer("keywrap_kek0"),
                KeywrapMack:          models.ToPointer("keywrap_mack2"),
                Port:                 models.ToPointer(models.RadiusAcctPortContainer.FromNumber(176)),
                Secret:               "secret0",
            },
        },
        Airwatch:                             models.ToPointer(models.WlanAirwatch{
            ApiKey:               models.ToPointer("api_key8"),
            ConsoleUrl:           models.ToPointer("console_url2"),
            Enabled:              models.ToPointer(false),
            Password:             models.ToPointer("password8"),
            Username:             models.ToPointer("username4"),
        }),
        AllowIpv6Ndp:                         models.ToPointer(true),
        AllowMdns:                            models.ToPointer(false),
        AllowSsdp:                            models.ToPointer(false),
        ArpFilter:                            models.ToPointer(false),
        AuthServerSelection:                  models.ToPointer(models.WlanAuthServerSelectionEnum_ORDERED),
        AuthServersNasId:                     models.NewOptional(models.ToPointer("5c5b350e0101-nas")),
        AuthServersNasIp:                     models.NewOptional(models.ToPointer("15.3.1.5")),
        AuthServersRetries:                   models.ToPointer(5),
        AuthServersTimeout:                   models.ToPointer(5),
        BandSteer:                            models.ToPointer(false),
        BandSteerForceBand5:                  models.ToPointer(false),
        BlockBlacklistClients:                models.ToPointer(false),
        ClientLimitDownEnabled:               models.ToPointer(false),
        ClientLimitUpEnabled:                 models.ToPointer(false),
        Disable11ax:                          models.ToPointer(false),
        Disable11be:                          models.ToPointer(false),
        DisableHtVhtRates:                    models.ToPointer(false),
        DisableMessageAuthenticatorCheck:     models.ToPointer(false),
        DisableUapsd:                         models.ToPointer(false),
        DisableV1RoamNotify:                  models.ToPointer(false),
        DisableV2RoamNotify:                  models.ToPointer(false),
        DisableWhenGatewayUnreachable:        models.ToPointer(false),
        DisableWhenMxtunnelDown:              models.ToPointer(false),
        DisableWmm:                           models.ToPointer(false),
        Dtim:                                 models.ToPointer(2),
        EnableFtm:                            models.ToPointer(false),
        EnableLocalKeycaching:                models.ToPointer(false),
        EnableWirelessBridging:               models.ToPointer(false),
        EnableWirelessBridgingDhcpTracking:   models.ToPointer(false),
        Enabled:                              models.ToPointer(true),
        FastDot1xTimers:                      models.ToPointer(false),
        HideSsid:                             models.ToPointer(false),
        HostnameIe:                           models.ToPointer(false),
        Id:                                   models.ToPointer(uuid.MustParse("53f10664-3ce8-4c27-b382-0ef66432349f")),
        Interface:                            models.ToPointer(models.WlanInterfaceEnum_ALL),
        Isolation:                            models.ToPointer(false),
        L2Isolation:                          models.ToPointer(false),
        LegacyOverds:                         models.ToPointer(false),
        LimitBcast:                           models.ToPointer(false),
        LimitProbeResponse:                   models.ToPointer(false),
        MaxIdletime:                          models.ToPointer(1800),
        MaxNumClients:                        models.ToPointer(0),
        MspId:                                models.ToPointer(uuid.MustParse("b9d42c2e-88ee-41f8-b798-f009ce7fe909")),
        NoStaticDns:                          models.ToPointer(false),
        NoStaticIp:                           models.ToPointer(false),
        OrgId:                                models.ToPointer(uuid.MustParse("a97c1b22-a4e9-411e-9bfd-d8695a0f9e61")),
        PortalAllowedHostnames:               []string{
            "snapchat.com",
            "ibm.com",
        },
        PortalAllowedSubnets:                 []string{
            "63.5.3.0/24",
        },
        PortalApiSecret:                      models.NewOptional(models.ToPointer("EIfPMOykI3lMlDdNPub2WcbqT6dNOtWwmYHAd6bY")),
        PortalDeniedHostnames:                []string{
            "msg.snapchat.com",
        },
        PortalImage:                          models.NewOptional(models.ToPointer("https://url/to/image.png")),
        ReconnectClientsWhenRoamingMxcluster: models.ToPointer(false),
        RoamMode:                             models.ToPointer(models.WlanRoamModeEnum_NONE),
        SiteId:                               models.ToPointer(uuid.MustParse("441a1214-6928-442a-8e92-e1d34b8ec6a6")),
        SleExcluded:                          models.ToPointer(false),
        Ssid:                                 "corporate",
        UseEapolV1:                           models.ToPointer(false),
        VlanEnabled:                          models.ToPointer(false),
        VlanPooling:                          models.ToPointer(false),
        WlanLimitDownEnabled:                 models.ToPointer(false),
        WlanLimitUpEnabled:                   models.ToPointer(false),
        AdditionalProperties:                 map[string]interface{}{
            "exampleAdditionalProperty": interface{}("[key1, val1][key2, val2]"),
        },
    }

}
```

