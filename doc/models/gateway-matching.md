
# Gateway Matching

Gateway matching

## Structure

`GatewayMatching`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Enable` | `*bool` | Optional | - |
| `Rules` | [`[]models.GatewayMatchingRule`](../../doc/models/gateway-matching-rule.md) | Optional | **Constraints**: *Unique Items Required* |

## Example (as JSON)

```json
{
  "enable": false,
  "rules": [
    {
      "additional_config_cmds": [
        "additional_config_cmds8"
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
        }
      },
      "exampleAdditionalProperty": "gateway_matching_rule_additionalProperties8"
    },
    {
      "additional_config_cmds": [
        "additional_config_cmds8"
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
        }
      },
      "exampleAdditionalProperty": "gateway_matching_rule_additionalProperties8"
    },
    {
      "additional_config_cmds": [
        "additional_config_cmds8"
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
        }
      },
      "exampleAdditionalProperty": "gateway_matching_rule_additionalProperties8"
    }
  ]
}
```

