
# Switch Matching Rule

Property key defines the type of matching, value is the string to match. e.g:

* `match_name[0:3]`: switch name must match the first 3 letters of the property value
* `match_name[2:6]`: switch name must match the property value from the 2nd to the 6th letter
* `match_model[0-8]`: switch model must match the first 8 letters of the property value
* `match_role`: switch role must match the property value

*This model accepts additional fields of type string.*

## Structure

`SwitchMatchingRule`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AdditionalConfigCmds` | `[]string` | Optional | additional CLI commands to append to the generated Junos config. **Note**: no check is done |
| `IpConfig` | [`*models.SwitchMatchingRuleIpConfig`](../../doc/models/switch-matching-rule-ip-config.md) | Optional | In-Band Management interface configuration |
| `Name` | `*string` | Optional | Rule name. WARNING: the name `default` is reserved and can only be used for the last rule in the list<br><br>**Constraints**: *Minimum Length*: `1`, *Maximum Length*: `32` |
| `OobIpConfig` | [`*models.SwitchMatchingRuleOobIpConfig`](../../doc/models/switch-matching-rule-oob-ip-config.md) | Optional | Out-of-Band Management interface configuration |
| `PortConfig` | [`map[string]models.JunosPortConfig`](../../doc/models/junos-port-config.md) | Optional | Property key is the port name or range (e.g. "ge-0/0/0-10") |
| `PortMirroring` | [`map[string]models.SwitchPortMirroringProperty`](../../doc/models/switch-port-mirroring-property.md) | Optional | Property key is the port mirroring instance name. `port_mirroring` can be added under device/site settings. It takes interface and ports as input for ingress, interface as input for egress and can take interface and port as output. A maximum 4 mirroring ports is allowed |
| `StpConfig` | [`*models.SwitchStpConfig`](../../doc/models/switch-stp-config.md) | Optional | - |
| `SwitchMgmt` | [`*models.SwitchMgmt`](../../doc/models/switch-mgmt.md) | Optional | Switch settings |
| `AdditionalProperties` | `map[string]string` | Optional | - |

## Example (as JSON)

```json
{
  "match_model": "EX4300",
  "match_name[0:3]": "abc",
  "additional_config_cmds": [
    "additional_config_cmds4"
  ],
  "ip_config": {
    "network": "network6",
    "type": "dhcp",
    "exampleAdditionalProperty": {
      "key1": "val1",
      "key2": "val2"
    }
  },
  "name": "name2",
  "oob_ip_config": {
    "type": "dhcp",
    "use_mgmt_vrf": false,
    "use_mgmt_vrf_for_host_out": false,
    "exampleAdditionalProperty": {
      "key1": "val1",
      "key2": "val2"
    }
  },
  "port_config": {
    "key0": {
      "ae_disable_lacp": false,
      "ae_idx": 230,
      "ae_lacp_slow": false,
      "aggregated": false,
      "critical": false,
      "usage": "usage6",
      "exampleAdditionalProperty": {
        "key1": "val1",
        "key2": "val2"
      }
    },
    "key1": {
      "ae_disable_lacp": false,
      "ae_idx": 230,
      "ae_lacp_slow": false,
      "aggregated": false,
      "critical": false,
      "usage": "usage6",
      "exampleAdditionalProperty": {
        "key1": "val1",
        "key2": "val2"
      }
    },
    "key2": {
      "ae_disable_lacp": false,
      "ae_idx": 230,
      "ae_lacp_slow": false,
      "aggregated": false,
      "critical": false,
      "usage": "usage6",
      "exampleAdditionalProperty": {
        "key1": "val1",
        "key2": "val2"
      }
    }
  }
}
```

