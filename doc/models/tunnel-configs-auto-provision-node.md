
# Tunnel Configs Auto Provision Node

## Structure

`TunnelConfigsAutoProvisionNode`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `NumHosts` | `*string` | Optional | - |
| `WanNames` | `[]string` | Optional | optional, only needed if `vars_only`==`false` |

## Example (as JSON)

```json
{
  "num_hosts": "num_hosts4",
  "wan_names": [
    "wan_names6",
    "wan_names7"
  ]
}
```

