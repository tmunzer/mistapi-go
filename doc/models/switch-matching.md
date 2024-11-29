
# Switch Matching

Switch template

## Structure

`SwitchMatching`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Enable` | `*bool` | Optional | - |
| `Rules` | [`[]models.SwitchMatchingRule`](../../doc/models/switch-matching-rule.md) | Optional | **Constraints**: *Unique Items Required* |

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
        "type": "dhcp"
      },
      "name": "name8",
      "oob_ip_config": {
        "type": "dhcp",
        "use_mgmt_vrf": false,
        "use_mgmt_vrf_for_host_out": false
      },
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
      }
    },
    {
      "additional_config_cmds": [
        "additional_config_cmds8"
      ],
      "ip_config": {
        "network": "network6",
        "type": "dhcp"
      },
      "name": "name8",
      "oob_ip_config": {
        "type": "dhcp",
        "use_mgmt_vrf": false,
        "use_mgmt_vrf_for_host_out": false
      },
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
      }
    }
  ]
}
```

