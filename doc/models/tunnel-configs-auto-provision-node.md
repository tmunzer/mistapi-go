
# Tunnel Configs Auto Provision Node

*This model accepts additional fields of type interface{}.*

## Structure

`TunnelConfigsAutoProvisionNode`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `NumHosts` | `*string` | Optional | - |
| `WanNames` | `[]string` | Optional | optional, only needed if `vars_only`==`false` |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "num_hosts": "num_hosts4",
  "wan_names": [
    "wan_names6",
    "wan_names7"
  ],
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

