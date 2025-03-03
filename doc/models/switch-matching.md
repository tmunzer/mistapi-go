
# Switch Matching

Defines custom switch configuration based on different criteria

*This model accepts additional fields of type interface{}.*

## Structure

`SwitchMatching`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Enable` | `*bool` | Optional | - |
| `Rules` | [`[]models.SwitchMatchingRule`](../../doc/models/switch-matching-rule.md) | Optional | **Constraints**: *Unique Items Required* |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "enable": false,
  "rules": [
    {
      "additional_config_cmds": [
        "additional_config_cmds8"
      ],
      "ip_config": {
        "network": "network6",
        "type": "dhcp",
        "exampleAdditionalProperty": {
          "key1": "val1",
          "key2": "val2"
        }
      },
      "name": "name8",
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
        }
      },
      "exampleAdditionalProperty": "switch_matching_rule_additionalProperties2"
    },
    {
      "additional_config_cmds": [
        "additional_config_cmds8"
      ],
      "ip_config": {
        "network": "network6",
        "type": "dhcp",
        "exampleAdditionalProperty": {
          "key1": "val1",
          "key2": "val2"
        }
      },
      "name": "name8",
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
        }
      },
      "exampleAdditionalProperty": "switch_matching_rule_additionalProperties2"
    }
  ],
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

