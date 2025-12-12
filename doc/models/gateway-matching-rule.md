
# Gateway Matching Rule

*This model accepts additional fields of type string.*

## Structure

`GatewayMatchingRule`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AdditionalConfigCmds` | `[]string` | Optional | additional CLI commands to append to the generated Junos config. **Note**: no check is done |
| `Name` | `*string` | Optional | - |
| `PortConfig` | [`*models.GatewayPortConfig`](../../doc/models/gateway-port-config.md) | Optional | Gateway port config |
| `AdditionalProperties` | `map[string]string` | Optional | property key defines the type of matching. e.g: `match_name[0:3]`, `match_model[0-6]` or `match_role` |

## Example (as JSON)

```json
{
  "additional_config_cmds": [
    "additional_config_cmds4"
  ],
  "name": "name0",
  "port_config": {
    "ae_disable_lacp": false,
    "ae_idx": "ae_idx8",
    "ae_lacp_force_up": false,
    "aggregated": false,
    "critical": false,
    "usage": "lan"
  },
  "exampleAdditionalProperty": "gateway_matching_rule_additionalProperties8"
}
```

