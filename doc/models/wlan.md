
# Wlan

**Note**: portal_template will be forked out of wlan objects soon. To fetch portal_template, please query portal_template_url. To update portal_template, use Wlan Portal Template.

*This model accepts additional fields of type interface{}.*

## Structure

`Wlan`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AcctImmediateUpdate` | `*bool` | Optional | Enable coa-immediate-update and address-change-immediate-update on the access profile.<br>**Default**: `false` |
| `AcctInterimInterval` | `*int` | Optional | How frequently should interim accounting be reported, 60-65535. default is 0 (use one specified in Access-Accept request from RADIUS Server). Very frequent messages can affect the performance of the radius server, 600 and up is recommended when enabled<br>**Default**: `0`<br>**Constraints**: `>= 0`, `<= 65535` |
| `AcctServers` | [`[]models.RadiusAcctServer`](../../doc/models/radius-acct-server.md) | Optional | List of RADIUS accounting servers, optional, order matters where the first one is treated as primary |
| `Airwatch` | [`*models.WlanAirwatch`](../../doc/models/wlan-airwatch.md) | Optional | Airwatch wlan settings |
| `AllowIpv6Ndp` | `*bool` | Optional | Only applicable when limit_bcast==true, which allows or disallows ipv6 Neighbor Discovery packets to go through<br>**Default**: `true` |
| `AllowMdns` | `*bool` | Optional | Only applicable when limit_bcast==true, which allows mDNS / Bonjour packets to go through<br>**Default**: `false` |
| `AllowSsdp` | `*bool` | Optional | Only applicable when `limit_bcast`==`true`, which allows SSDP<br>**Default**: `false` |
| `ApIds` | `models.Optional[[]uuid.UUID]` | Optional | List of device ids |
| `AppLimit` | [`*models.WlanAppLimit`](../../doc/models/wlan-app-limit.md) | Optional | Bandwidth limiting for apps (applies to up/down) |
| `AppQos` | [`*models.WlanAppQos`](../../doc/models/wlan-app-qos.md) | Optional | APp qos wlan settings |
| `ApplyTo` | [`*models.WlanApplyToEnum`](../../doc/models/wlan-apply-to-enum.md) | Optional | enum: `aps`, `site`, `wxtags` |
| `ArpFilter` | `*bool` | Optional | Whether to enable smart arp filter<br>**Default**: `false` |
| `Auth` | [`*models.WlanAuth`](../../doc/models/wlan-auth.md) | Optional | Authentication wlan settings |
| `AuthServerSelection` | [`*models.WlanAuthServerSelectionEnum`](../../doc/models/wlan-auth-server-selection-enum.md) | Optional | When ordered, AP will prefer and go back to the first server if possible. enum: `ordered`, `unordered`<br>**Default**: `"ordered"` |
| `AuthServers` | [`[]models.RadiusAuthServer`](../../doc/models/radius-auth-server.md) | Optional | List of RADIUS authentication servers, at least one is needed if `auth type`==`eap`, order matters where the first one is treated as primary |
| `AuthServersNasId` | `models.Optional[string]` | Optional | Optional, up to 48 bytes, will be dynamically generated if not provided. used only for authentication servers |
| `AuthServersNasIp` | `models.Optional[string]` | Optional | Optional, NAS-IP-ADDRESS to use |
| `AuthServersRetries` | `*int` | Optional | Radius auth session retries. Following fast timers are set if "fast_dot1x_timers" knob is enabled. ‘retries’  are set to value of auth_servers_retries. ‘max-requests’ is also set when setting auth_servers_retries and is set to default value to 3.<br>**Default**: `2` |
| `AuthServersTimeout` | `*int` | Optional | Radius auth session timeout. Following fast timers are set if "fast_dot1x_timers" knob is enabled. ‘quite-period’  and ‘transmit-period’ are set to half the value of auth_servers_timeout. ‘supplicant-timeout’ is also set when setting auth_servers_timeout and is set to default value of 10.<br>**Default**: `5` |
| `Band` | `*string` | Optional | `band` is deprecated and kept for backward compatibility. Use bands instead |
| `BandSteer` | `*bool` | Optional | Whether to enable band_steering, this works only when band==both<br>**Default**: `false` |
| `BandSteerForceBand5` | `*bool` | Optional | Force dual_band capable client to connect to 5G<br>**Default**: `false` |
| `Bands` | [`[]models.Dot11BandEnum`](../../doc/models/dot-11-band-enum.md) | Optional | List of radios that the wlan should apply to. |
| `BlockBlacklistClients` | `*bool` | Optional | Whether to block the clients in the blacklist (up to first 256 macs)<br>**Default**: `false` |
| `Bonjour` | [`*models.WlanBonjour`](../../doc/models/wlan-bonjour.md) | Optional | Bonjour gateway wlan settings |
| `CiscoCwa` | [`*models.WlanCiscoCwa`](../../doc/models/wlan-cisco-cwa.md) | Optional | Cisco CWA (central web authentication) required RADIUS with COA in order to work. See CWA: https://www.cisco.com/c/en/us/support/docs/security/identity-services-engine/115732-central-web-auth-00.html |
| `ClientLimitDown` | `*int` | Optional | In kbps<br>**Constraints**: `>= 1`, `<= 999000` |
| `ClientLimitDownEnabled` | `*bool` | Optional | If downlink limiting per-client is enabled<br>**Default**: `false` |
| `ClientLimitUp` | `*int` | Optional | In kbps<br>**Constraints**: `>= 1`, `<= 999000` |
| `ClientLimitUpEnabled` | `*bool` | Optional | If uplink limiting per-client is enabled<br>**Default**: `false` |
| `CoaServers` | [`[]models.CoaServer`](../../doc/models/coa-server.md) | Optional | List of COA (change of authorization) servers, optional |
| `CreatedTime` | `*float64` | Optional | When the object has been created, in epoch |
| `Disable11ax` | `*bool` | Optional | Some old WLAN drivers may not be compatible<br>**Default**: `false` |
| `Disable11be` | `*bool` | Optional | To disable Wi-Fi 7 EHT IEs<br>**Default**: `false` |
| `DisableHtVhtRates` | `*bool` | Optional | To disable ht or vht rates<br>**Default**: `false` |
| `DisableUapsd` | `*bool` | Optional | Whether to disable U-APSD<br>**Default**: `false` |
| `DisableV1RoamNotify` | `*bool` | Optional | Disable sending v2 roam notification messages<br>**Default**: `false` |
| `DisableV2RoamNotify` | `*bool` | Optional | Disable sending v2 roam notification messages<br>**Default**: `false` |
| `DisableWhenGatewayUnreachable` | `*bool` | Optional | When any of the following is true, this WLAN will be disabled<br><br>* cannot get IP<br>* cannot obtain default gateway<br>* cannot reach default gateway<br>**Default**: `false` |
| `DisableWhenMxtunnelDown` | `*bool` | Optional | **Default**: `false` |
| `DisableWmm` | `*bool` | Optional | Whether to disable WMM<br>**Default**: `false` |
| `DnsServerRewrite` | [`models.Optional[models.WlanDnsServerRewrite]`](../../doc/models/wlan-dns-server-rewrite.md) | Optional | For radius_group-based DNS server (rewrite DNS request depending on the Group RADIUS server returns) |
| `Dtim` | `*int` | Optional | **Default**: `2` |
| `DynamicPsk` | [`models.Optional[models.WlanDynamicPsk]`](../../doc/models/wlan-dynamic-psk.md) | Optional | For dynamic PSK where we get per_user PSK from Radius. dynamic_psk allows PSK to be selected at runtime depending on context (wlan/site/user/...) thus following configurations are assumed (currently)<br><br>* PSK will come from RADIUS server<br>* AP sends client MAC as username and password (i.e. `enable_mac_auth` is assumed)<br>* AP sends BSSID:SSID as Caller-Station-ID<br>* `auth_servers` is required<br>* PSK will come from cloud WLC if source is cloud_psks<br>* default_psk will be used if cloud WLC is not available<br>* `multi_psk_only` and `psk` is ignored<br>* `pairwise` can only be wpa2-ccmp (for now, wpa3 support on the roadmap) |
| `DynamicVlan` | [`models.Optional[models.WlanDynamicVlan]`](../../doc/models/wlan-dynamic-vlan.md) | Optional | For 802.1x |
| `EnableLocalKeycaching` | `*bool` | Optional | Enable AP-AP keycaching via multicast<br>**Default**: `false` |
| `EnableWirelessBridging` | `*bool` | Optional | By default, we'd inspect all DHCP packets and drop those unrelated to the wireless client itself in the case where client is a wireless bridge (DHCP packets for other MACs will need to be forwarded), wireless_bridging can be enabled<br>**Default**: `false` |
| `EnableWirelessBridgingDhcpTracking` | `*bool` | Optional | If the client bridge is doing DHCP on behalf of other devices (L2-NAT), enable dhcp_tracking will cut down DHCP response packets to be forwarded to wireless<br>**Default**: `false` |
| `Enabled` | `*bool` | Optional | If this wlan is enabled<br>**Default**: `true` |
| `FastDot1xTimers` | `*bool` | Optional | If set to true, sets default fast-timers with values calculated from ‘auth_servers_timeout’ and ‘auth_server_retries’ .<br>**Default**: `false` |
| `ForSite` | `*bool` | Optional | - |
| `HideSsid` | `*bool` | Optional | Whether to hide SSID in beacon<br>**Default**: `false` |
| `HostnameIe` | `*bool` | Optional | Include hostname inside IE in AP beacons / probe responses<br>**Default**: `false` |
| `Hotspot20` | [`*models.WlanHotspot20`](../../doc/models/wlan-hotspot-20.md) | Optional | Hostspot 2.0 wlan settings |
| `Id` | `*uuid.UUID` | Optional | Unique ID of the object instance in the Mist Organization |
| `InjectDhcpOption82` | [`*models.WlanInjectDhcpOption82`](../../doc/models/wlan-inject-dhcp-option-82.md) | Optional | - |
| `Interface` | [`*models.WlanInterfaceEnum`](../../doc/models/wlan-interface-enum.md) | Optional | where this WLAN will be connected to. enum: `all`, `eth0`, `eth1`, `eth2`, `eth3`, `mxtunnel`, `site_mxedge`, `wxtunnel`<br>**Default**: `"all"` |
| `Isolation` | `*bool` | Optional | Whether to stop clients to talk to each other<br>**Default**: `false` |
| `L2Isolation` | `*bool` | Optional | If isolation is enabled, whether to deny clients to talk to L2 on the LAN<br>**Default**: `false` |
| `LegacyOverds` | `*bool` | Optional | Legacy devices requires the Over-DS (for Fast BSS Transition) bit set (while our chip doesn’t support it). Warning! Enabling this will cause problem for iOS devices.<br>**Default**: `false` |
| `LimitBcast` | `*bool` | Optional | Whether to limit broadcast packets going to wireless (i.e. only allow certain bcast packets to go through)<br>**Default**: `false` |
| `LimitProbeResponse` | `*bool` | Optional | Limit probe response base on some heuristic rules<br>**Default**: `false` |
| `MaxIdletime` | `*int` | Optional | Max idle time in seconds<br>**Default**: `1800`<br>**Constraints**: `>= 60`, `<= 86400` |
| `MaxNumClients` | `*int` | Optional | Maximum number of client connected to the SSID. `0` means unlimited<br>**Default**: `0`<br>**Constraints**: `>= 0`, `<= 128` |
| `MistNac` | [`*models.WlanMistNac`](../../doc/models/wlan-mist-nac.md) | Optional | - |
| `ModifiedTime` | `*float64` | Optional | When the object has been modified for the last time, in epoch |
| `MspId` | `*uuid.UUID` | Optional | - |
| `MxtunnelId` | `*uuid.UUID` | Optional | (deprecated, use mxtunnel_ids instead) when `interface`==`mxtunnel`, id of the Mist Tunnel |
| `MxtunnelIds` | `[]string` | Optional | When `interface`=`mxtunnel`, id of the Mist Tunnel |
| `MxtunnelName` | `[]string` | Optional | When `interface`=`site_mxedge`, name of the mxtunnel that in mxtunnels under Site Setting |
| `NoStaticDns` | `*bool` | Optional | Whether to only allow client to use DNS that we’ve learned from DHCP response<br>**Default**: `false` |
| `NoStaticIp` | `*bool` | Optional | Whether to only allow client that we’ve learned from DHCP exchange to talk<br>**Default**: `false` |
| `OrgId` | `*uuid.UUID` | Optional | - |
| `Portal` | [`*models.WlanPortal`](../../doc/models/wlan-portal.md) | Optional | Portal wlan settings |
| `PortalAllowedHostnames` | `[]string` | Optional | List of hostnames without http(s):// (matched by substring) |
| `PortalAllowedSubnets` | `[]string` | Optional | List of CIDRs |
| `PortalApiSecret` | `models.Optional[string]` | Optional | APi secret (auto-generated) that can be used to sign guest authorization requests |
| `PortalDeniedHostnames` | `[]string` | Optional | List of hostnames without http(s):// (matched by substring), this takes precedence over portal_allowed_hostnames |
| `PortalImage` | `models.Optional[string]` | Optional | Url of portal background image |
| `PortalSsoUrl` | `models.Optional[string]` | Optional | - |
| `PortalTemplateUrl` | `models.Optional[string]` | Optional | N.B portal_template will be forked out of wlan objects soon. To fetch portal_template, please query portal_template_url. To update portal_template, use Wlan Portal Template. |
| `Qos` | [`*models.WlanQos`](../../doc/models/wlan-qos.md) | Optional | - |
| `Radsec` | [`*models.Radsec`](../../doc/models/radsec.md) | Optional | RadSec settings |
| `Rateset` | [`map[string]models.WlanDatarates`](../../doc/models/wlan-datarates.md) | Optional | Property key is the RF band. enum: `24`, `5`, `6` |
| `ReconnectClientsWhenRoamingMxcluster` | `*bool` | Optional | When different mxcluster is on different subnet, we'd want to disconnect clients (so they'll reconnect and get new IPs)<br>**Default**: `false` |
| `RoamMode` | [`*models.WlanRoamModeEnum`](../../doc/models/wlan-roam-mode-enum.md) | Optional | enum: `11r`, `OKC`, `NONE`<br>**Default**: `"NONE"` |
| `Schedule` | [`*models.WlanSchedule`](../../doc/models/wlan-schedule.md) | Optional | WLAN operating schedule, default is disabled |
| `SiteId` | `*uuid.UUID` | Optional | - |
| `SleExcluded` | `*bool` | Optional | Whether to exclude this WLAN from SLE metrics<br>**Default**: `false` |
| `Ssid` | `string` | Required | Name of the SSID |
| `TemplateId` | `models.Optional[uuid.UUID]` | Optional | - |
| `Thumbnail` | `models.Optional[string]` | Optional | Url of portal background image thumbnail |
| `UseEapolV1` | `*bool` | Optional | If `auth.type`==`eap` or `auth.type`==`psk`, should only be set for legacy client, such as pre-2004, 802.11b devices<br>**Default**: `false` |
| `VlanEnabled` | `*bool` | Optional | If vlan tagging is enabled<br>**Default**: `false` |
| `VlanId` | [`models.Optional[models.WlanVlanIdWithVariable]`](../../doc/models/containers/wlan-vlan-id-with-variable.md) | Optional | - |
| `VlanIds` | [`*models.WlanVlanIds`](../../doc/models/containers/wlan-vlan-ids.md) | Optional | - |
| `VlanPooling` | `*bool` | Optional | Requires `vlan_enabled`==`true` to be set to `true`. Vlan pooling allows AP to place client on different VLAN using a deterministic algorithm<br>**Default**: `false` |
| `WlanLimitDown` | `models.Optional[int]` | Optional | In kbps |
| `WlanLimitDownEnabled` | `*bool` | Optional | If downlink limiting for whole wlan is enabled<br>**Default**: `false` |
| `WlanLimitUp` | `models.Optional[int]` | Optional | In kbps |
| `WlanLimitUpEnabled` | `*bool` | Optional | If uplink limiting for whole wlan is enabled<br>**Default**: `false` |
| `WxtagIds` | `models.Optional[[]uuid.UUID]` | Optional | List of wxtag_ids |
| `WxtunnelId` | `models.Optional[string]` | Optional | When `interface`=`wxtunnel`, id of the WXLAN Tunnel |
| `WxtunnelRemoteId` | `models.Optional[string]` | Optional | When `interface`=`wxtunnel`, remote tunnel identifier |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "acct_immediate_update": false,
  "acct_interim_interval": 0,
  "allow_ipv6_ndp": true,
  "allow_mdns": false,
  "allow_ssdp": false,
  "arp_filter": false,
  "auth_server_selection": "ordered",
  "auth_servers_nas_id": "5c5b350e0101-nas",
  "auth_servers_nas_ip": "15.3.1.5",
  "auth_servers_retries": 5,
  "auth_servers_timeout": 5,
  "band_steer": false,
  "band_steer_force_band5": false,
  "block_blacklist_clients": false,
  "client_limit_down_enabled": false,
  "client_limit_up_enabled": false,
  "disable_11ax": false,
  "disable_11be": false,
  "disable_ht_vht_rates": false,
  "disable_uapsd": false,
  "disable_v1_roam_notify": false,
  "disable_v2_roam_notify": false,
  "disable_when_gateway_unreachable": false,
  "disable_when_mxtunnel_down": false,
  "disable_wmm": false,
  "dtim": 2,
  "enable_local_keycaching": false,
  "enable_wireless_bridging": false,
  "enable_wireless_bridging_dhcp_tracking": false,
  "enabled": true,
  "fast_dot1x_timers": false,
  "hide_ssid": false,
  "hostname_ie": false,
  "id": "53f10664-3ce8-4c27-b382-0ef66432349f",
  "interface": "all",
  "isolation": false,
  "l2_isolation": false,
  "legacy_overds": false,
  "limit_bcast": false,
  "limit_probe_response": false,
  "max_idletime": 1800,
  "max_num_clients": 0,
  "msp_id": "b9d42c2e-88ee-41f8-b798-f009ce7fe909",
  "no_static_dns": false,
  "no_static_ip": false,
  "org_id": "a97c1b22-a4e9-411e-9bfd-d8695a0f9e61",
  "portal_allowed_hostnames": [
    "snapchat.com",
    "ibm.com"
  ],
  "portal_allowed_subnets": [
    "63.5.3.0/24"
  ],
  "portal_api_secret": "EIfPMOykI3lMlDdNPub2WcbqT6dNOtWwmYHAd6bY",
  "portal_denied_hostnames": [
    "msg.snapchat.com"
  ],
  "portal_image": "https://url/to/image.png",
  "reconnect_clients_when_roaming_mxcluster": false,
  "roam_mode": "NONE",
  "site_id": "441a1214-6928-442a-8e92-e1d34b8ec6a6",
  "sle_excluded": false,
  "ssid": "corporate",
  "use_eapol_v1": false,
  "vlan_enabled": false,
  "vlan_pooling": false,
  "wlan_limit_down_enabled": false,
  "wlan_limit_up_enabled": false,
  "acct_servers": [
    {
      "host": "host4",
      "keywrap_enabled": false,
      "keywrap_format": "ascii",
      "keywrap_kek": "keywrap_kek0",
      "keywrap_mack": "keywrap_mack2",
      "port": 254,
      "secret": "secret0",
      "exampleAdditionalProperty": {
        "key1": "val1",
        "key2": "val2"
      }
    }
  ],
  "airwatch": {
    "api_key": "api_key8",
    "console_url": "console_url2",
    "enabled": false,
    "password": "password8",
    "username": "username4",
    "exampleAdditionalProperty": {
      "key1": "val1",
      "key2": "val2"
    }
  },
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

