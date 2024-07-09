
# Device Switch

Switch Configuration.
You can configure `port_usages` and `networks` settings at the device level, but most of the time it's better use the Site Setting to achieve better consistency and be able to re-use the same settings across switches entries defined here will "replace" those defined in Site Setting/Network Template

## Structure

`DeviceSwitch`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AclPolicies` | [`[]models.AclPolicy`](../../doc/models/acl-policy.md) | Optional | - |
| `AclTags` | [`map[string]models.AclTag`](../../doc/models/acl-tag.md) | Optional | ACL Tags to identify traffic source or destination. Key name is the tag name |
| `AdditionalConfigCmds` | `[]string` | Optional | additional CLI commands to append to the generated Junos config<br><br>**Note**: no check is done |
| `CreatedTime` | `*float64` | Optional | - |
| `DeviceprofileId` | `*uuid.UUID` | Optional | - |
| `DhcpSnooping` | [`*models.DhcpSnooping`](../../doc/models/dhcp-snooping.md) | Optional | - |
| `DhcpdConfig` | [`*models.DhcpdConfig`](../../doc/models/dhcpd-config.md) | Optional | - |
| `DisableAutoConfig` | `*bool` | Optional | for a claimed switch, we control the configs by default. This option (disables the behavior)<br>**Default**: `false` |
| `DnsServers` | `[]string` | Optional | Global dns settings. To keep compatibility, dns settings in `ip_config` and `oob_ip_config` will overwrite this setting |
| `DnsSuffix` | `[]string` | Optional | Global dns settings. To keep compatibility, dns settings in `ip_config` and `oob_ip_config` will overwrite this setting |
| `EvpnConfig` | [`*models.EvpnConfig`](../../doc/models/evpn-config.md) | Optional | EVPN Junos settings |
| `ExtraRoutes` | [`map[string]models.ExtraRouteProperties`](../../doc/models/extra-route-properties.md) | Optional | - |
| `ExtraRoutes6` | [`map[string]models.ExtraRoute6Properties`](../../doc/models/extra-route-6-properties.md) | Optional | Property key is the destination CIDR (e.g. "2a02:1234:420a:10c9::/64") |
| `Id` | `*uuid.UUID` | Optional | - |
| `Image1Url` | `models.Optional[string]` | Optional | - |
| `Image2Url` | `models.Optional[string]` | Optional | - |
| `Image3Url` | `models.Optional[string]` | Optional | - |
| `IpConfig` | [`*models.JunosIpConfig`](../../doc/models/junos-ip-config.md) | Optional | Junos IP Config |
| `Mac` | `*string` | Optional | device MAC address |
| `Managed` | `*bool` | Optional | for an adopted switch, we don’t overwrite their existing configs automatically<br>**Default**: `false` |
| `MapId` | `*uuid.UUID` | Optional | map where the device belongs to |
| `MistNac` | [`*models.SwitchMistNac`](../../doc/models/switch-mist-nac.md) | Optional | enable mist_nac to use radsec |
| `Model` | `*string` | Optional | device Model |
| `ModifiedTime` | `*float64` | Optional | - |
| `Name` | `*string` | Optional | - |
| `Networks` | [`map[string]models.SwitchNetwork`](../../doc/models/switch-network.md) | Optional | Property key is network name |
| `Notes` | `*string` | Optional | - |
| `NtpServers` | `[]string` | Optional | list of NTP servers specific to this device. By default, those in Site Settings will be used |
| `OobIpConfig` | [`*models.SwitchOobIpConfig`](../../doc/models/switch-oob-ip-config.md) | Optional | - If HA configuration: key parameter will be nodeX (eg: node1)<br>- If there are 2 routing engines, re1 mgmt IP has to be set separately (if desired): key parameter = `re1` |
| `OrgId` | `*uuid.UUID` | Optional | - |
| `OspfConfig` | [`*models.OspfConfig`](../../doc/models/ospf-config.md) | Optional | Junos OSPF config |
| `OtherIpConfigs` | [`map[string]models.JunosOtherIpConfig`](../../doc/models/junos-other-ip-config.md) | Optional | Property key is the network name |
| `PortConfig` | [`map[string]models.JunosPortConfig`](../../doc/models/junos-port-config.md) | Optional | Property key is the port name or range (e.g. "ge-0/0/0-10") |
| `PortMirroring` | [`map[string]models.SwitchPortMirroringProperty`](../../doc/models/switch-port-mirroring-property.md) | Optional | Property key is the port mirroring instance name<br>port_mirroring can be added under device/site settings. It takes interface and ports as input for ingress, interface as input for egress and can take interface and port as output. |
| `PortUsages` | [`map[string]models.SwitchPortUsage`](../../doc/models/switch-port-usage.md) | Optional | - |
| `RadiusConfig` | [`*models.RadiusConfig`](../../doc/models/radius-config.md) | Optional | Junos Radius config |
| `RemoteSyslog` | [`*models.RemoteSyslog`](../../doc/models/remote-syslog.md) | Optional | - |
| `Role` | `*string` | Optional | - |
| `RouterId` | `*string` | Optional | used for OSPF / BGP / EVPN |
| `Serial` | `*string` | Optional | device Serial |
| `SiteId` | `*uuid.UUID` | Optional | - |
| `SnmpConfig` | [`*models.SnmpConfig`](../../doc/models/snmp-config.md) | Optional | - |
| `StpConfig` | [`*models.SwitchStpConfig`](../../doc/models/switch-stp-config.md) | Optional | - |
| `SwitchMgmt` | [`*models.SwitchMgmt`](../../doc/models/switch-mgmt.md) | Optional | - |
| `Type` | [`*models.DeviceTypeSwitchEnum`](../../doc/models/device-type-switch-enum.md) | Optional | Device Type |
| `UseRouterIdAsSourceIp` | `*bool` | Optional | whether to use it for snmp / syslog / tacplus / radius<br>**Default**: `false` |
| `Vars` | `map[string]string` | Optional | a dictionary of name->value, the vars can then be used in Wlans. This can overwrite those from Site Vars |
| `VirtualChassis` | [`*models.SwitchVirtualChassis`](../../doc/models/switch-virtual-chassis.md) | Optional | required for preprovisioned Virtual Chassis |
| `VrfConfig` | [`*models.VrfConfig`](../../doc/models/vrf-config.md) | Optional | - |
| `VrfInstances` | [`map[string]models.VrfInstance`](../../doc/models/vrf-instance.md) | Optional | Property key is the network name |
| `VrrpConfig` | [`*models.VrrpConfig`](../../doc/models/vrrp-config.md) | Optional | Junos VRRP config |
| `X` | `*float64` | Optional | x in pixel |
| `Y` | `*float64` | Optional | y in pixel |

## Example (as JSON)

```json
{
  "disable_auto_config": false,
  "extra_routes6": {
    "2a02:1234:420a:10c9::/64": {
      "via": "2a02:1234:200a::100"
    }
  },
  "managed": false,
  "map_id": "63eda950-c6da-11e4-a628-60f81dd250cc",
  "org_id": "a97c1b22-a4e9-411e-9bfd-d8695a0f9e61",
  "router_id": "10.2.1.10",
  "site_id": "441a1214-6928-442a-8e92-e1d34b8ec6a6",
  "use_router_id_as_source_ip": false,
  "vars": {
    "RADIUS_IP1": "172.31.2.5",
    "RADIUS_SECRET": "11s64632d"
  },
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
  "x": 53.5,
  "y": 173.1,
  "acl_policies": [
    {
      "actions": [
        {
          "action": "allow",
          "dst_tag": "dst_tag0"
        }
      ],
      "name": "name2",
      "src_tags": [
        "src_tags1",
        "src_tags0"
      ]
    }
  ],
  "acl_tags": {
    "key0": {
      "gbp_tag": 14,
      "macs": [
        "macs1"
      ],
      "network": "network2",
      "radius_group": "radius_group8",
      "specs": [
        {
          "port_range": "port_range8",
          "protocol": "protocol6"
        }
      ],
      "type": "any"
    },
    "key1": {
      "gbp_tag": 14,
      "macs": [
        "macs1"
      ],
      "network": "network2",
      "radius_group": "radius_group8",
      "specs": [
        {
          "port_range": "port_range8",
          "protocol": "protocol6"
        }
      ],
      "type": "any"
    },
    "key2": {
      "gbp_tag": 14,
      "macs": [
        "macs1"
      ],
      "network": "network2",
      "radius_group": "radius_group8",
      "specs": [
        {
          "port_range": "port_range8",
          "protocol": "protocol6"
        }
      ],
      "type": "any"
    }
  },
  "additional_config_cmds": [
    "additional_config_cmds4",
    "additional_config_cmds3",
    "additional_config_cmds2"
  ],
  "created_time": 63.2,
  "deviceprofile_id": "0000012c-0000-0000-0000-000000000000"
}
```
