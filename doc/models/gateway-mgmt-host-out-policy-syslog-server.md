
# Gateway Mgmt Host Out Policy Syslog Server

Allows to define the host_out_policy per Syslog Server. The Property key is the Syslog name

## Structure

`GatewayMgmtHostOutPolicySyslogServer`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Host` | `*string` | Optional | - |
| `PathPreference` | `*string` | Optional | - |
| `ServerName` | `*string` | Optional | - |

## Example (as JSON)

```json
{
  "host": "103.35.3.5",
  "path_preference": "dc_only",
  "server_name": "dc_syslog_server"
}
```

