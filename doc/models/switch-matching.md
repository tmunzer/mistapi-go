
# Switch Matching

Defines custom switch configuration based on different criteria

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
      "default_port_usage": "default_port_usage4",
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
      "exampleAdditionalProperty": "switch_matching_rule_additionalProperties2"
    },
    {
      "additional_config_cmds": [
        "additional_config_cmds8"
      ],
      "default_port_usage": "default_port_usage4",
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
      "exampleAdditionalProperty": "switch_matching_rule_additionalProperties2"
    }
  ]
}
```

