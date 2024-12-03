
# Network Template

Network Template

*This model accepts additional fields of type interface{}.*

## Structure

`NetworkTemplate`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AclPolicies` | [`[]models.AclPolicy`](../../doc/models/acl-policy.md) | Optional | - |
| `AclTags` | [`map[string]models.AclTag`](../../doc/models/acl-tag.md) | Optional | ACL Tags to identify traffic source or destination. Key name is the tag name |
| `AdditionalConfigCmds` | `[]string` | Optional | additional CLI commands to append to the generated Junos config. **Note**: no check is done |
| `CreatedTime` | `*float64` | Optional | when the object has been created, in epoch |
| `DhcpSnooping` | [`*models.DhcpSnooping`](../../doc/models/dhcp-snooping.md) | Optional | - |
| `DnsServers` | `[]string` | Optional | Global dns settings. To keep compatibility, dns settings in `ip_config` and `oob_ip_config` will overwrite this setting |
| `DnsSuffix` | `[]string` | Optional | Global dns settings. To keep compatibility, dns settings in `ip_config` and `oob_ip_config` will overwrite this setting |
| `ExtraRoutes` | [`map[string]models.ExtraRoute`](../../doc/models/extra-route.md) | Optional | - |
| `ExtraRoutes6` | [`map[string]models.ExtraRoute6`](../../doc/models/extra-route-6.md) | Optional | Property key is the destination CIDR (e.g. "2a02:1234:420a:10c9::/64") |
| `Id` | `*uuid.UUID` | Optional | Unique ID of the object instance in the Mist Organnization |
| `ImportOrgNetworks` | `[]string` | Optional | Org Networks that we'd like to import |
| `MistNac` | [`*models.SwitchMistNac`](../../doc/models/switch-mist-nac.md) | Optional | enable mist_nac to use radsec |
| `ModifiedTime` | `*float64` | Optional | when the object has been modified for the last time, in epoch |
| `Name` | `*string` | Optional | - |
| `Networks` | [`map[string]models.SwitchNetwork`](../../doc/models/switch-network.md) | Optional | Property key is network name |
| `NtpServers` | `[]string` | Optional | list of NTP servers specific to this device. By default, those in Site Settings will be used |
| `OrgId` | `*uuid.UUID` | Optional | - |
| `OspfAreas` | [`map[string]models.OspfArea`](../../doc/models/ospf-area.md) | Optional | Junos OSPF areas |
| `PortMirroring` | [`map[string]models.SwitchPortMirroringProperty`](../../doc/models/switch-port-mirroring-property.md) | Optional | Property key is the port mirroring instance name<br>port_mirroring can be added under device/site settings. It takes interface and ports as input for ingress, interface as input for egress and can take interface and port as output. A maximum 4 port mirrorings is allowed |
| `PortUsages` | [`map[string]models.SwitchPortUsage`](../../doc/models/switch-port-usage.md) | Optional | Property key is the port usage name. Defines the profiles of port configuration configured on the switch |
| `RadiusConfig` | [`*models.SwitchRadiusConfig`](../../doc/models/switch-radius-config.md) | Optional | Junos Radius config |
| `RemoteSyslog` | [`*models.RemoteSyslog`](../../doc/models/remote-syslog.md) | Optional | - |
| `RemoveExistingConfigs` | `*bool` | Optional | by default, when we configure a device, we only clean up config we generates. Remove existing configs if enabled<br>**Default**: `false` |
| `SnmpConfig` | [`*models.SnmpConfig`](../../doc/models/snmp-config.md) | Optional | - |
| `SwitchMatching` | [`*models.SwitchMatching`](../../doc/models/switch-matching.md) | Optional | defines custom switch configuration based on different criterias |
| `SwitchMgmt` | [`*models.SwitchMgmt`](../../doc/models/switch-mgmt.md) | Optional | Switch settings |
| `VrfConfig` | [`*models.VrfConfig`](../../doc/models/vrf-config.md) | Optional | - |
| `VrfInstances` | [`map[string]models.SwitchVrfInstance`](../../doc/models/switch-vrf-instance.md) | Optional | Property key is the network name |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "extra_routes6": {
    "2a02:1234:420a:10c9::/64": {
      "via": "2a02:1234:200a::100"
    }
  },
  "id": "53f10664-3ce8-4c27-b382-0ef66432349f",
  "org_id": "a97c1b22-a4e9-411e-9bfd-d8695a0f9e61",
  "remove_existing_configs": false,
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
          "dst_tag": "dst_tag0",
          "exampleAdditionalProperty": {
            "key1": "val1",
            "key2": "val2"
          }
        }
      ],
      "name": "name2",
      "src_tags": [
        "src_tags1",
        "src_tags0"
      ],
      "exampleAdditionalProperty": {
        "key1": "val1",
        "key2": "val2"
      }
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
          "protocol": "protocol6",
          "exampleAdditionalProperty": {
            "key1": "val1",
            "key2": "val2"
          }
        }
      ],
      "type": "dynamic_gbp",
      "exampleAdditionalProperty": {
        "key1": "val1",
        "key2": "val2"
      }
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
          "protocol": "protocol6",
          "exampleAdditionalProperty": {
            "key1": "val1",
            "key2": "val2"
          }
        }
      ],
      "type": "dynamic_gbp",
      "exampleAdditionalProperty": {
        "key1": "val1",
        "key2": "val2"
      }
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
          "protocol": "protocol6",
          "exampleAdditionalProperty": {
            "key1": "val1",
            "key2": "val2"
          }
        }
      ],
      "type": "dynamic_gbp",
      "exampleAdditionalProperty": {
        "key1": "val1",
        "key2": "val2"
      }
    }
  },
  "additional_config_cmds": [
    "additional_config_cmds0",
    "additional_config_cmds9",
    "additional_config_cmds8"
  ],
  "created_time": 167.36,
  "dhcp_snooping": {
    "all_networks": false,
    "enable_arp_spoof_check": false,
    "enable_ip_source_guard": false,
    "enabled": false,
    "networks": [
      "networks8",
      "networks9"
    ],
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

