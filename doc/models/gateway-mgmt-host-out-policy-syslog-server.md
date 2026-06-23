
# Gateway Mgmt Host Out Policy Syslog Server

Per-syslog-server host-out path policy override

## Structure

`GatewayMgmtHostOutPolicySyslogServer`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Host` | `*string` | Optional | Syslog server hostname or IP address |
| `PathPreference` | `*string` | Optional | Preferred path name used for this syslog server |
| `ServerName` | `*string` | Optional | Remote syslog server name referenced by the policy |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    gatewayMgmtHostOutPolicySyslogServer := models.GatewayMgmtHostOutPolicySyslogServer{
        Host:                 models.ToPointer("103.35.3.5"),
        PathPreference:       models.ToPointer("dc_only"),
        ServerName:           models.ToPointer("dc_syslog_server"),
    }

}
```

