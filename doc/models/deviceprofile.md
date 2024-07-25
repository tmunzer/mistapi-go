
# Deviceprofile

## Structure

`Deviceprofile`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Aeroscout` | [`*models.ApAeroscout`](../../doc/models/ap-aeroscout.md) | Optional | Aeroscout AP settings |
| `BleConfig` | [`*models.BleConfig`](../../doc/models/ble-config.md) | Optional | BLE AP settings |
| `CreatedTime` | `*float64` | Optional | - |
| `DisableEth1` | `*bool` | Optional | whether to disable eth1 port<br>**Default**: `false` |
| `DisableEth2` | `*bool` | Optional | whether to disable eth2 port<br>**Default**: `false` |
| `DisableEth3` | `*bool` | Optional | whether to disable eth3 port<br>**Default**: `false` |
| `DisableModule` | `*bool` | Optional | whether to disable module port<br>**Default**: `false` |
| `EslConfig` | [`*models.ApEslConfig`](../../doc/models/ap-esl-config.md) | Optional | - |
| `ForSite` | `*bool` | Optional | - |
| `Id` | `*uuid.UUID` | Optional | - |
| `IotConfig` | [`*models.ApIot`](../../doc/models/ap-iot.md) | Optional | IoT AP settings |
| `IpConfig` | [`*models.ApIpConfig`](../../doc/models/ap-ip-config.md) | Optional | IP AP settings |
| `Led` | [`*models.ApLed`](../../doc/models/ap-led.md) | Optional | LED AP settings |
| `Mesh` | [`*models.ApMesh`](../../doc/models/ap-mesh.md) | Optional | Mesh AP settings |
| `ModifiedTime` | `*float64` | Optional | - |
| `Name` | `models.Optional[string]` | Optional | - |
| `NtpServers` | `[]string` | Optional | list of NTP servers specific to this device. By default, those in Site Settings will be used<br>**Constraints**: *Unique Items Required* |
| `OrgId` | `*uuid.UUID` | Optional | - |
| `PoePassthrough` | `*bool` | Optional | whether to enable power out through module port (for APH) or eth1 (for APL/BT11)<br>**Default**: `false` |
| `PortConfig` | [`*models.PortConfig`](../../doc/models/port-config.md) | Optional | Property key is the interface(s) name (e.g. "eth1,eth2") |
| `PwrConfig` | [`*models.ApPwrConfig`](../../doc/models/ap-pwr-config.md) | Optional | power related configs |
| `RadioConfig` | [`*models.ApRadio`](../../doc/models/ap-radio.md) | Optional | Radio AP settings |
| `SiteId` | `*uuid.UUID` | Optional | - |
| `SwitchConfig` | [`*models.ApSwitch`](../../doc/models/ap-switch.md) | Optional | for people who want to fully control the vlans (advanced) |
| `Type` | [`*models.DeviceTypeApEnum`](../../doc/models/device-type-ap-enum.md) | Optional | Device Type |
| `UplinkPortConfig` | [`*models.ApUplinkPortConfig`](../../doc/models/ap-uplink-port-config.md) | Optional | - |
| `UsbConfig` | [`*models.ApUsb`](../../doc/models/ap-usb.md) | Optional | USB AP settings<br>Note: if native imagotag is enabled, BLE will be disabled automatically<br>Note: legacy, new config moved to ESL Config. |
| `Vars` | `map[string]string` | Optional | a dictionary of name->value, the vars can then be used in Wlans. This can overwrite those from Site Vars |
| `AdditionalConfigCmds` | `[]string` | Optional | additional CLI commands to append to the generated Junos config<br><br>**Note**: no check is done |
| `BgpConfig` | [`map[string]models.BgpConfig`](../../doc/models/bgp-config.md) | Optional | - |
| `DhcpdConfig` | [`*models.DhcpdConfig`](../../doc/models/dhcpd-config.md) | Optional | - |
| `DnsOverride` | `*bool` | Optional | **Default**: `false` |
| `DnsServers` | `[]string` | Optional | Global dns settings. To keep compatibility, dns settings in `ip_config` and `oob_ip_config` will overwrite this setting |
| `DnsSuffix` | `[]string` | Optional | Global dns settings. To keep compatibility, dns settings in `ip_config` and `oob_ip_config` will overwrite this setting |
| `ExtraRoutes` | [`map[string]models.GatewayExtraRoute`](../../doc/models/gateway-extra-route.md) | Optional | Property key is the destination CIDR (e.g. "10.0.0.0/8") |
| `ExtraRoutes6` | [`map[string]models.GatewayExtraRoute`](../../doc/models/gateway-extra-route.md) | Optional | Property key is the destination CIDR (e.g. "2a02:1234:420a:10c9::/64") |
| `GatewayMatching` | [`*models.GatewayMatching`](../../doc/models/gateway-matching.md) | Optional | Gateway matching |
| `IdpProfiles` | [`map[string]models.IdpProfile`](../../doc/models/idp-profile.md) | Optional | Property key is the profile name |
| `IpConfigs` | [`map[string]models.GatewayIpConfigProperty`](../../doc/models/gateway-ip-config-property.md) | Optional | Property key is the network name |
| `Networks` | [`[]models.Network`](../../doc/models/network.md) | Optional | - |
| `NtpOverride` | `*bool` | Optional | **Default**: `false` |
| `OobIpConfig` | [`*models.GatewayOobIpConfig`](../../doc/models/gateway-oob-ip-config.md) | Optional | out-of-band (vme/em0/fxp0) IP config |
| `PathPreferences` | [`map[string]models.GatewayPathPreferences`](../../doc/models/gateway-path-preferences.md) | Optional | Property key is the path name |
| `RouterId` | `*string` | Optional | auto assigned if not set |
| `RoutingPolicies` | [`map[string]models.RoutingPolicy`](../../doc/models/routing-policy.md) | Optional | Property key is the routing policy name |
| `ServicePolicies` | [`[]models.ServicePolicy`](../../doc/models/service-policy.md) | Optional | - |
| `TunnelConfigs` | [`map[string]models.TunnelConfigs`](../../doc/models/tunnel-configs.md) | Optional | Property key is the tunnel name |
| `TunnelProviderOptions` | [`*models.TunnelProviderOptions`](../../doc/models/tunnel-provider-options.md) | Optional | - |

## Example (as JSON)

```json
{
  "disable_eth1": false,
  "disable_eth2": false,
  "disable_eth3": false,
  "disable_module": false,
  "name": "gw_template",
  "org_id": "a97c1b22-a4e9-411e-9bfd-d8695a0f9e61",
  "poe_passthrough": false,
  "site_id": "441a1214-6928-442a-8e92-e1d34b8ec6a6",
  "vars": {
    "RADIUS_IP1": "172.31.2.5",
    "RADIUS_SECRET": "11s64632d"
  },
  "dnsOverride": false,
  "extra_routes6": {
    "2a02:1234:420a:10c9::/64": {
      "via": "2a02:1234:200a::100"
    }
  },
  "ntpOverride": false,
  "router_id": "10.2.1.10",
  "aeroscout": {
    "enabled": false,
    "host": "host6",
    "locate_connected": false
  },
  "ble_config": {
    "beacon_enabled": false,
    "beacon_rate": 110,
    "beacon_rate_mode": "default",
    "beam_disabled": [
      113,
      114,
      115
    ],
    "custom_ble_packet_enabled": false
  },
  "created_time": 79.24
}
```

