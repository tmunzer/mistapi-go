
# Deviceprofile Switch

Switch Device Profiles can be applied to one or multiple switches. The settings from the Device Profile will override the settings from the Switch Template and the Site Settings.

*This model accepts additional fields of type interface{}.*

## Structure

`DeviceprofileSwitch`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AclPolicies` | [`[]models.AclPolicy`](../../doc/models/acl-policy.md) | Optional | - |
| `AclTags` | [`map[string]models.AclTag`](../../doc/models/acl-tag.md) | Optional | ACL Tags to identify traffic source or destination. Key name is the tag name |
| `AdditionalConfigCmds` | `[]string` | Optional | additional CLI commands to append to the generated Junos config. **Note**: no check is done |
| `AggregateRoutes` | [`map[string]models.AggregateRoute`](../../doc/models/aggregate-route.md) | Optional | Property key is the destination subnet (e.g. "172.16.3.0/24") |
| `AggregateRoutes6` | [`map[string]models.AggregateRoute`](../../doc/models/aggregate-route.md) | Optional | Property key is the destination subnet (e.g. "2a02:1234:420a:10c9::/64") |
| `CreatedTime` | `*float64` | Optional | When the object has been created, in epoch |
| `DhcpSnooping` | [`*models.DhcpSnooping`](../../doc/models/dhcp-snooping.md) | Optional | - |
| `DhcpdConfig` | [`*models.SwitchDhcpdConfig`](../../doc/models/switch-dhcpd-config.md) | Optional | - |
| `DnsServers` | `[]string` | Optional | Global dns settings. To keep compatibility, dns settings in `ip_config` and `oob_ip_config` will overwrite this setting |
| `DnsSuffix` | `[]string` | Optional | Global dns settings. To keep compatibility, dns settings in `ip_config` and `oob_ip_config` will overwrite this setting |
| `EvpnConfig` | [`*models.EvpnConfig`](../../doc/models/evpn-config.md) | Optional | EVPN Junos settings |
| `ExtraRoutes` | [`map[string]models.ExtraRoute`](../../doc/models/extra-route.md) | Optional | Property key is the destination CIDR (e.g. "10.0.0.0/8") |
| `ExtraRoutes6` | [`map[string]models.ExtraRoute6`](../../doc/models/extra-route-6.md) | Optional | Property key is the destination CIDR (e.g. "2a02:1234:420a:10c9::/64") |
| `Id` | `*uuid.UUID` | Optional | Unique ID of the object instance in the Mist Organization |
| `IotConfig` | [`map[string]models.SwitchIotPort`](../../doc/models/switch-iot-port.md) | Optional | Property Key is the IOT port name, e.g.:<br><br>* `IN0` or `IN1` for the FPC0 input port with 5V triggered inputs<br>* `OUT1` for the FPC0 output port (can only be triggered by either IN0 or IN1)<br>* "X/IN0`, `X/IN1`and`X/OUT` are used to define IOT ports on VC members |
| `IpConfig` | [`*models.JunosIpConfig`](../../doc/models/junos-ip-config.md) | Optional | Junos IP Config |
| `MistNac` | [`*models.SwitchMistNac`](../../doc/models/switch-mist-nac.md) | Optional | Enable mist_nac to use RadSec |
| `ModifiedTime` | `*float64` | Optional | When the object has been modified for the last time, in epoch |
| `Name` | `string` | Required | - |
| `Networks` | [`map[string]models.SwitchNetwork`](../../doc/models/switch-network.md) | Optional | Property key is network name |
| `NtpServers` | `[]string` | Optional | List of NTP servers specific to this device. By default, those in Site Settings will be used |
| `OobIpConfig` | [`*models.SwitchOobIpConfig`](../../doc/models/switch-oob-ip-config.md) | Optional | Switch OOB IP Config:<br><br>- If HA configuration: key parameter will be nodeX (eg: node1)<br>- If there are 2 routing engines, re1 mgmt IP has to be set separately (if desired): key parameter = `re1` |
| `OrgId` | `*uuid.UUID` | Optional | - |
| `OspfAreas` | [`map[string]models.OspfArea`](../../doc/models/ospf-area.md) | Optional | Junos OSPF areas. Property key is the OSPF Area (Area should be a number (0-255) / IP address) |
| `OtherIpConfigs` | [`map[string]models.JunosOtherIpConfig`](../../doc/models/junos-other-ip-config.md) | Optional | Property key is the network name. Defines the additional IP Addresses configured on the device. |
| `PortConfig` | [`map[string]models.JunosPortConfig`](../../doc/models/junos-port-config.md) | Optional | Property key is the port name or range (e.g. "ge-0/0/0-10") |
| `PortMirroring` | [`map[string]models.SwitchPortMirroringProperty`](../../doc/models/switch-port-mirroring-property.md) | Optional | Property key is the port mirroring instance name. `port_mirroring` can be added under device/site settings. It takes interface and ports as input for ingress, interface as input for egress and can take interface and port as output. A maximum 4 mirroring ports is allowed |
| `PortUsages` | [`map[string]models.SwitchPortUsage`](../../doc/models/switch-port-usage.md) | Optional | Property key is the port usage name. Defines the profiles of port configuration configured on the switch |
| `RadiusConfig` | [`*models.SwitchRadiusConfig`](../../doc/models/switch-radius-config.md) | Optional | Junos Radius config |
| `RemoteSyslog` | [`*models.RemoteSyslog`](../../doc/models/remote-syslog.md) | Optional | - |
| `SiteId` | `*uuid.UUID` | Optional | - |
| `SnmpConfig` | [`*models.SnmpConfig`](../../doc/models/snmp-config.md) | Optional | - |
| `StpConfig` | [`*models.SwitchStpConfig`](../../doc/models/switch-stp-config.md) | Optional | - |
| `SwitchMgmt` | [`*models.SwitchMgmt`](../../doc/models/switch-mgmt.md) | Optional | Switch settings |
| `Type` | `string` | Required, Constant | Device Type. enum: `switch`<br><br>**Value**: `"switch"` |
| `UseRouterIdAsSourceIp` | `*bool` | Optional | Whether to use it for snmp / syslog / tacplus / radius<br><br>**Default**: `false` |
| `VrfConfig` | [`*models.VrfConfig`](../../doc/models/vrf-config.md) | Optional | - |
| `VrfInstances` | [`map[string]models.SwitchVrfInstance`](../../doc/models/switch-vrf-instance.md) | Optional | Property key is the network name |
| `VrrpConfig` | [`*models.VrrpConfig`](../../doc/models/vrrp-config.md) | Optional | Junos VRRP config |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "aggregate_routes": {
    "172.16.3.0/24": {
      "discard": false,
      "metric": 28,
      "preference": 30
    }
  },
  "aggregate_routes6": {
    "2a02:1234:420a:10c9::/64": {
      "discard": false,
      "metric": 126,
      "preference": 30
    }
  },
  "extra_routes": {
    "0.0.0.0/0": {
      "via": "192.168.1.10"
    }
  },
  "extra_routes6": {
    "2a02:1234:420a:10c9::/64": {
      "via": "2a02:1234:200a::100"
    }
  },
  "id": "53f10664-3ce8-4c27-b382-0ef66432349f",
  "name": "name4",
  "org_id": "a97c1b22-a4e9-411e-9bfd-d8695a0f9e61",
  "site_id": "441a1214-6928-442a-8e92-e1d34b8ec6a6",
  "type": "switch",
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
    },
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
      "ether_types": [
        "ether_types8",
        "ether_types9"
      ],
      "gbp_tag": 14,
      "macs": [
        "macs1"
      ],
      "network": "network2",
      "port_usage": "port_usage0",
      "type": "network"
    }
  },
  "additional_config_cmds": [
    "additional_config_cmds4",
    "additional_config_cmds3",
    "additional_config_cmds2"
  ],
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

