
# Service Policy Syslog

Required for syslog logging

*This model accepts additional fields of type interface{}.*

## Structure

`ServicePolicySyslog`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Enabled` | `*bool` | Optional | **Default**: `false` |
| `ServerNames` | `[]string` | Optional | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "enabled": false,
  "server_names": [
    "dc_syslog_server"
  ],
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

