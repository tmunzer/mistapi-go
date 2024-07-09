
# Gateway Matching Rule

## Structure

`GatewayMatchingRule`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AdditionalConfigCmds` | `[]string` | Optional | additional CLI commands to append to the generated Junos config<br><br>**Note**: no check is done |
| `Name` | `*string` | Optional | - |
| `PortConfig` | [`map[string]models.JunosPortConfig`](../../doc/models/junos-port-config.md) | Optional | - |

## Example (as JSON)

```json
{
  "additional_config_cmds": [
    "additional_config_cmds2"
  ],
  "name": "name8",
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
  }
}
```

