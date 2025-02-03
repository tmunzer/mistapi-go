
# Tunnel Config Auto Provision Node

*This model accepts additional fields of type interface{}.*

## Structure

`TunnelConfigAutoProvisionNode`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `ProbeIps` | `[]string` | Optional | **Constraints**: *Unique Items Required* |
| `WanNames` | `[]string` | Optional | Optional, only needed if `vars_only`==`false` |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "probe_ips": [
    "probe_ips7"
  ],
  "wan_names": [
    "wan_names8",
    "wan_names9"
  ],
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

