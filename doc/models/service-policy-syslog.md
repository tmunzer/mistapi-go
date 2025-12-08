
# Service Policy Syslog

Required for syslog logging

## Structure

`ServicePolicySyslog`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Enabled` | `*bool` | Optional | **Default**: `false` |
| `ServerNames` | `[]string` | Optional | - |

## Example (as JSON)

```json
{
  "enabled": false,
  "server_names": [
    "dc_syslog_server"
  ]
}
```

