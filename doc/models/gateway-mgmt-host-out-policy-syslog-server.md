
# Gateway Mgmt Host Out Policy Syslog Server

Allows to define the host_out_policy per Syslog Server. The Property key is the Syslog name

*This model accepts additional fields of type interface{}.*

## Structure

`GatewayMgmtHostOutPolicySyslogServer`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Host` | `*string` | Optional | - |
| `Name` | `*string` | Optional | - |
| `PathPreference` | `*string` | Optional | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "host": "103.35.3.5",
  "name": "server1",
  "path_preference": "dc_only",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

