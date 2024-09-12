
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
      "match_role": "match_role8",
      "name": "name8",
      "oob_ip_config": {
        "type": "dhcp",
        "use_mgmt_vrf": {
          "key1": "val1",
          "key2": "val2"
        },
        "use_mgmt_vrf_for_host_out": {
          "key1": "val1",
          "key2": "val2"
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
      "match_role": "match_role8",
      "name": "name8",
      "oob_ip_config": {
        "type": "dhcp",
        "use_mgmt_vrf": {
          "key1": "val1",
          "key2": "val2"
        },
        "use_mgmt_vrf_for_host_out": {
          "key1": "val1",
          "key2": "val2"
        }
      }
    }
  ]
}
```

