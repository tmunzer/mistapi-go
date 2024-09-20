
# Switch Matching Rule

property key define the type of matching, value is the string to match. e.g: `match_name[0:3]`, `match_name[2:6]`, `match_model`,  `match_model[0-6]`

## Structure

`SwitchMatchingRule`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AdditionalConfigCmds` | `[]string` | Optional | additional CLI commands to append to the generated Junos config<br><br>**Note**: no check is done |
| `IpConfig` | [`*models.SwitchMatchingRuleIpConfig`](../../doc/models/switch-matching-rule-ip-config.md) | Optional | In-Band Management interface configuration |
| `MatchRole` | `*string` | Optional | role to match |
| `Name` | `*string` | Optional | - |
| `OobIpConfig` | [`*models.SwitchMatchingRuleOobIpConfig`](../../doc/models/switch-matching-rule-oob-ip-config.md) | Optional | Out-of-Band Management interface configuration |
| `PortConfig` | [`map[string]models.JunosPortConfig`](../../doc/models/junos-port-config.md) | Optional | Propery key is the interface name or interface range |
| `PortMirroring` | [`map[string]models.SwitchPortMirroringProperty`](../../doc/models/switch-port-mirroring-property.md) | Optional | Property key is the port mirroring instance name<br>port_mirroring can be added under device/site settings. It takes interface and ports as input for ingress, interface as input for egress and can take interface and port as output. A maximum 4 port mirrorings is allowed |
| `SwitchMgmt` | [`*models.SwitchMgmt`](../../doc/models/switch-mgmt.md) | Optional | Switch settings |

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
    "type": "dhcp"
  },
  "match_role": "match_role2",
  "name": "name2",
  "oob_ip_config": {
    "type": "dhcp",
    "use_mgmt_vrf": false,
    "use_mgmt_vrf_for_host_out": false
  }
}
```

