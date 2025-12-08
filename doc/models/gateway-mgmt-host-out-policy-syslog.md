
# Gateway Mgmt Host Out Policy Syslog

## Structure

`GatewayMgmtHostOutPolicySyslog`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `PathPreference` | `*string` | Optional | - |
| `Servers` | [`[]models.GatewayMgmtHostOutPolicySyslogServer`](../../doc/models/gateway-mgmt-host-out-policy-syslog-server.md) | Optional | - |

## Example (as JSON)

```json
{
  "path_preference": "broadband_wans",
  "servers": [
    {
      "host": "host4",
      "path_preference": "path_preference8",
      "server_name": "server_name8"
    },
    {
      "host": "host4",
      "path_preference": "path_preference8",
      "server_name": "server_name8"
    },
    {
      "host": "host4",
      "path_preference": "path_preference8",
      "server_name": "server_name8"
    }
  ]
}
```

