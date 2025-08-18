
# Gateway Mgmt Host Out Policy Syslog

*This model accepts additional fields of type interface{}.*

## Structure

`GatewayMgmtHostOutPolicySyslog`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `PathPreference` | `*string` | Optional | - |
| `Servers` | [`[]models.GatewayMgmtHostOutPolicySyslogServer`](../../doc/models/gateway-mgmt-host-out-policy-syslog-server.md) | Optional | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "path_preference": "broadband_wans",
  "servers": [
    {
      "host": "host4",
      "path_preference": "path_preference8",
      "server_name": "server_name8",
      "exampleAdditionalProperty": {
        "key1": "val1",
        "key2": "val2"
      }
    },
    {
      "host": "host4",
      "path_preference": "path_preference8",
      "server_name": "server_name8",
      "exampleAdditionalProperty": {
        "key1": "val1",
        "key2": "val2"
      }
    },
    {
      "host": "host4",
      "path_preference": "path_preference8",
      "server_name": "server_name8",
      "exampleAdditionalProperty": {
        "key1": "val1",
        "key2": "val2"
      }
    }
  ],
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

