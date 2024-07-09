
# Switch Matching Rule

property key define the type of matching, value is the string to match. e.g: `match_name[0:3]`, `match_name[2:6]`, `match_model`,  `match_model[0-6]`

## Structure

`SwitchMatchingRule`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AdditionalConfigCmds` | `[]string` | Optional | additional CLI commands to append to the generated Junos config<br><br>**Note**: no check is done |
| `MatchRole` | `*string` | Optional | role to match |
| `Name` | `*string` | Optional | - |
| `PortConfig` | [`map[string]models.JunosPortConfig`](../../doc/models/junos-port-config.md) | Optional | Propery key is the interface name or interface range |
| `PortMirroring` | [`map[string]models.SwitchPortMirroringProperty`](../../doc/models/switch-port-mirroring-property.md) | Optional | Property key is the port mirroring instance name<br>port_mirroring can be added under device/site settings. It takes interface and ports as input for ingress, interface as input for egress and can take interface and port as output. |
| `SwitchMgmt` | [`*models.ConfigSwitch`](../../doc/models/config-switch.md) | Optional | Switch settings |

## Example (as JSON)

```json
{
  "match_model": "EX4300",
  "match_name[0:3]": "abc",
  "additional_config_cmds": [
    "additional_config_cmds8"
  ],
  "match_role": "match_role0",
  "name": "name0",
  "port_config": {
    "key0": {
      "ae_disable_lacp": false,
      "ae_idx": 230,
      "ae_lacp_slow": false,
      "aggregated": false,
      "critical": false,
      "usage": "usage6"
    },
    "key1": {
      "ae_disable_lacp": false,
      "ae_idx": 230,
      "ae_lacp_slow": false,
      "aggregated": false,
      "critical": false,
      "usage": "usage6"
    },
    "key2": {
      "ae_disable_lacp": false,
      "ae_idx": 230,
      "ae_lacp_slow": false,
      "aggregated": false,
      "critical": false,
      "usage": "usage6"
    }
  },
  "port_mirroring": {
    "key0": {
      "input_networks_ingress": [
        "input_networks_ingress8"
      ],
      "input_port_ids_egress": [
        "input_port_ids_egress4",
        "input_port_ids_egress5"
      ],
      "input_port_ids_ingress": [
        "input_port_ids_ingress2"
      ],
      "output_network": "output_network4",
      "output_port_id": "output_port_id2"
    },
    "key1": {
      "input_networks_ingress": [
        "input_networks_ingress8"
      ],
      "input_port_ids_egress": [
        "input_port_ids_egress4",
        "input_port_ids_egress5"
      ],
      "input_port_ids_ingress": [
        "input_port_ids_ingress2"
      ],
      "output_network": "output_network4",
      "output_port_id": "output_port_id2"
    },
    "key2": {
      "input_networks_ingress": [
        "input_networks_ingress8"
      ],
      "input_port_ids_egress": [
        "input_port_ids_egress4",
        "input_port_ids_egress5"
      ],
      "input_port_ids_ingress": [
        "input_port_ids_ingress2"
      ],
      "output_network": "output_network4",
      "output_port_id": "output_port_id2"
    }
  }
}
```

