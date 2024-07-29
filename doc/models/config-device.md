
# Config Device

## Structure

`ConfigDevice`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Aeroscout` | [`*models.ApAeroscout`](../../doc/models/ap-aeroscout.md) | Optional | Aeroscout AP settings |
| `BleConfig` | [`*models.BleConfig`](../../doc/models/ble-config.md) | Optional | BLE AP settings |
| `Centrak` | [`*models.ApCentrak`](../../doc/models/ap-centrak.md) | Optional | - |
| `ClientBridge` | [`*models.ApClientBridge`](../../doc/models/ap-client-bridge.md) | Optional | - |
| `CreatedTime` | `*float64` | Optional | - |
| `DeviceprofileId` | `models.Optional[uuid.UUID]` | Optional | - |
| `DisableEth1` | `*bool` | Optional | whether to disable eth1 port<br>**Default**: `false` |
| `DisableEth2` | `*bool` | Optional | whether to disable eth2 port<br>**Default**: `false` |
| `DisableEth3` | `*bool` | Optional | whether to disable eth3 port<br>**Default**: `false` |
| `DisableModule` | `*bool` | Optional | whether to disable module port<br>**Default**: `false` |
| `EslConfig` | [`*models.ApEslConfig`](../../doc/models/ap-esl-config.md) | Optional | - |
| `ForSite` | `*bool` | Optional | - |
| `Height` | `*float64` | Optional | height, in meters, optional |
| `Id` | `*uuid.UUID` | Optional | - |
| `Image1Url` | `models.Optional[string]` | Optional | - |
| `Image2Url` | `models.Optional[string]` | Optional | - |
| `Image3Url` | `models.Optional[string]` | Optional | - |
| `IotConfig` | [`*models.ApIot`](../../doc/models/ap-iot.md) | Optional | IoT AP settings |
| `IpConfig` | [`*models.ApIpConfig1`](../../doc/models/ap-ip-config-1.md) | Optional | IP AP settings |
| `Led` | [`*models.ApLed`](../../doc/models/ap-led.md) | Optional | LED AP settings |
| `Locked` | `*bool` | Optional | whether this map is considered locked down |
| `Mac` | `*string` | Optional | device MAC address |
| `MapId` | `*uuid.UUID` | Optional | map where the device belongs to |
| `Mesh` | [`*models.ApMesh`](../../doc/models/ap-mesh.md) | Optional | Mesh AP settings |
| `Model` | `*string` | Optional | device Model |
| `ModifiedTime` | `*float64` | Optional | - |
| `Name` | `*string` | Optional | - |
| `Notes` | `*string` | Optional | any notes about this AP |
| `NtpServers` | `[]string` | Optional | list of NTP servers specific to this device. By default, those in Site Settings will be used<br>**Constraints**: *Unique Items Required* |
| `OrgId` | `*uuid.UUID` | Optional | - |
| `Orientation` | `*int` | Optional | orientation, 0-359, in degrees, up is 0, right is 90.<br>**Constraints**: `>= 0`, `<= 359` |
| `PoePassthrough` | `*bool` | Optional | whether to enable power out through module port (for APH) or eth1 (for APL/BT11)<br>**Default**: `false` |
| `PortConfig` | [`*models.PortConfig1`](../../doc/models/port-config-1.md) | Optional | eth0 is not allowed here.<br>Property key is the interface(s) name (e.g. "eth1" or"eth1,eth2") |
| `PwrConfig` | [`*models.ApPwrConfig`](../../doc/models/ap-pwr-config.md) | Optional | power related configs |
| `RadioConfig` | [`*models.ApRadio`](../../doc/models/ap-radio.md) | Optional | Radio AP settings |
| `Serial` | `*string` | Optional | device Serial |
| `SiteId` | `*uuid.UUID` | Optional | - |
| `Type` | [`*models.DeviceTypeApEnum`](../../doc/models/device-type-ap-enum.md) | Optional | Device Type. enum: `ap` |
| `UplinkPortConfig` | [`*models.ApUplinkPortConfig`](../../doc/models/ap-uplink-port-config.md) | Optional | - |
| `UsbConfig` | [`*models.ApUsb`](../../doc/models/ap-usb.md) | Optional | USB AP settings<br>Note: if native imagotag is enabled, BLE will be disabled automatically<br>Note: legacy, new config moved to ESL Config. |
| `Vars` | `map[string]string` | Optional | a dictionary of name->value, the vars can then be used in Wlans. This can overwrite those from Site Vars |
| `X` | `*float64` | Optional | x in pixel |
| `Y` | `*float64` | Optional | y in pixel |
| `AclPolicies` | [`[]models.AclPolicy`](../../doc/models/acl-policy.md) | Optional | - |
| `AclTags` | [`map[string]models.AclTag`](../../doc/models/acl-tag.md) | Optional | ACL Tags to identify traffic source or destination. Key name is the tag name |
| `AdditionalConfigCmds` | `[]string` | Optional | additional CLI commands to append to the generated Junos config<br><br>**Note**: no check is done |
| `DhcpSnooping` | [`*models.DhcpSnooping`](../../doc/models/dhcp-snooping.md) | Optional | - |
| `DhcpdConfig` | [`*models.DhcpdConfig1`](../../doc/models/dhcpd-config-1.md) | Optional | - |
| `DisableAutoConfig` | `*bool` | Optional | for a claimed switch, we control the configs by default. This option (disables the behavior)<br>**Default**: `false` |
| `DnsServers` | `[]string` | Optional | Global dns settings. To keep compatibility, dns settings in `ip_config` and `oob_ip_config` will overwrite this setting |
| `DnsSuffix` | `[]string` | Optional | Global dns settings. To keep compatibility, dns settings in `ip_config` and `oob_ip_config` will overwrite this setting |
| `EvpnConfig` | [`*models.EvpnConfig`](../../doc/models/evpn-config.md) | Optional | EVPN Junos settings |
| `ExtraRoutes` | [`*models.ExtraRoutes`](../../doc/models/extra-routes.md) | Optional | Property key is the destination CIDR (e.g. "10.0.0.0/8") |
| `ExtraRoutes6` | [`*models.ExtraRoutes6`](../../doc/models/extra-routes-6.md) | Optional | Property key is the destination CIDR (e.g. "2a02:1234:420a:10c9::/64") |
| `Managed` | `*bool` | Optional | for an adopted switch, we don’t overwrite their existing configs automatically<br>**Default**: `false` |
| `MistNac` | [`*models.SwitchMistNac`](../../doc/models/switch-mist-nac.md) | Optional | enable mist_nac to use radsec |
| `Networks` | [`map[string]models.SwitchNetwork`](../../doc/models/switch-network.md) | Optional | Property key is network name |
| `OobIpConfig` | [`*models.SwitchOobIpConfig1`](../../doc/models/switch-oob-ip-config-1.md) | Optional | - If HA configuration: key parameter will be nodeX (eg: node1)<br>- If there are 2 routing engines, re1 mgmt IP has to be set separately (if desired): key parameter = `re1` |
| `OspfConfig` | [`*models.OspfConfig`](../../doc/models/ospf-config.md) | Optional | Junos OSPF config |
| `OtherIpConfigs` | [`map[string]models.JunosOtherIpConfig`](../../doc/models/junos-other-ip-config.md) | Optional | Property key is the network name |
| `PortMirroring` | [`*models.SwitchPortMirroring`](../../doc/models/switch-port-mirroring.md) | Optional | Property key is the port mirroring instance name<br>port_mirroring can be added under device/site settings. It takes interface and ports as input for ingress, interface as input for egress and can take interface and port as output. |
| `PortUsages` | [`map[string]models.SwitchPortUsage`](../../doc/models/switch-port-usage.md) | Optional | - |
| `RadiusConfig` | [`*models.RadiusConfig`](../../doc/models/radius-config.md) | Optional | Junos Radius config |
| `RemoteSyslog` | [`*models.RemoteSyslog`](../../doc/models/remote-syslog.md) | Optional | - |
| `Role` | `*string` | Optional | - |
| `RouterId` | `*string` | Optional | used for OSPF / BGP / EVPN |
| `SnmpConfig` | [`*models.SnmpConfig`](../../doc/models/snmp-config.md) | Optional | - |
| `StpConfig` | [`*models.SwitchStpConfig`](../../doc/models/switch-stp-config.md) | Optional | - |
| `SwitchMgmt` | [`*models.SwitchMgmt`](../../doc/models/switch-mgmt.md) | Optional | - |
| `UseRouterIdAsSourceIp` | `*bool` | Optional | whether to use it for snmp / syslog / tacplus / radius<br>**Default**: `false` |
| `VirtualChassis` | [`*models.SwitchVirtualChassis`](../../doc/models/switch-virtual-chassis.md) | Optional | required for preprovisioned Virtual Chassis |
| `VrfConfig` | [`*models.VrfConfig`](../../doc/models/vrf-config.md) | Optional | - |
| `VrfInstances` | [`*models.SwitchVrfInstances`](../../doc/models/switch-vrf-instances.md) | Optional | Property key is the network name |
| `VrrpConfig` | [`*models.VrrpConfig`](../../doc/models/vrrp-config.md) | Optional | Junos VRRP config |
| `BgpConfig` | [`map[string]models.BgpConfig`](../../doc/models/bgp-config.md) | Optional | - |
| `IdpProfiles` | [`map[string]models.IdpProfile`](../../doc/models/idp-profile.md) | Optional | Property key is the profile name |
| `IpConfigs` | [`map[string]models.GatewayIpConfigProperty`](../../doc/models/gateway-ip-config-property.md) | Optional | Property key is the network name |
| `MspId` | `*uuid.UUID` | Optional | - |
| `PathPreferences` | [`map[string]models.GatewayPathPreferences`](../../doc/models/gateway-path-preferences.md) | Optional | Property key is the path name |
| `RoutingPolicies` | [`map[string]models.RoutingPolicy`](../../doc/models/routing-policy.md) | Optional | Property key is the routing policy name |
| `ServicePolicies` | [`[]models.ServicePolicy`](../../doc/models/service-policy.md) | Optional | - |
| `TunnelConfigs` | [`map[string]models.TunnelConfigs`](../../doc/models/tunnel-configs.md) | Optional | Property key is the tunnel name |
| `TunnelProviderOptions` | [`*models.TunnelProviderOptions`](../../doc/models/tunnel-provider-options.md) | Optional | - |

## Example (as JSON)

```json
{
  "deviceprofile_id": "6f4bf402-45f9-2a56-6c8b-7f83d3bc98e9",
  "disable_eth1": false,
  "disable_eth2": false,
  "disable_eth3": false,
  "disable_module": false,
  "height": 2.75,
  "map_id": "63eda950-c6da-11e4-a628-60f81dd250cc",
  "name": "conference room",
  "notes": "slightly off center",
  "org_id": "a97c1b22-a4e9-411e-9bfd-d8695a0f9e61",
  "orientation": 45,
  "poe_passthrough": false,
  "site_id": "441a1214-6928-442a-8e92-e1d34b8ec6a6",
  "vars": {
    "RADIUS_IP1": "172.31.2.5",
    "RADIUS_SECRET": "11s64632d"
  },
  "x": 53.5,
  "y": 173.1,
  "disable_auto_config": false,
  "extra_routes6": {
    "2a02:1234:420a:10c9::/64": {
      "via": "2a02:1234:200a::100"
    }
  },
  "managed": false,
  "router_id": "10.2.1.10",
  "use_router_id_as_source_ip": false,
  "vrf_instances": {
    "guest": {
      "extra_routes": {
        "0.0.0.0/0": {
          "via": "192.168.31.1"
        }
      },
      "networks": [
        "guest"
      ]
    }
  },
  "aeroscout": {
    "enabled": false,
    "host": "host6",
    "locate_connected": false
  },
  "ble_config": {
    "beacon_enabled": false,
    "beacon_rate": 110,
    "beacon_rate_mode": "custom",
    "beam_disabled": [
      113,
      114,
      115
    ],
    "custom_ble_packet_enabled": false
  },
  "centrak": {
    "enabled": false
  },
  "client_bridge": {
    "auth": {
      "psk": "psk4",
      "type": "open"
    },
    "enabled": false,
    "ssid": "ssid0"
  },
  "created_time": 233.98
}
```

