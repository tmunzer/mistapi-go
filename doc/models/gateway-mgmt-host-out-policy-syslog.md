
# Gateway Mgmt Host Out Policy Syslog

Host-out path policy for gateway syslog traffic

## Structure

`GatewayMgmtHostOutPolicySyslog`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `PathPreference` | `*string` | Optional | Preferred path name used by default for gateway syslog traffic |
| `Servers` | [`[]models.GatewayMgmtHostOutPolicySyslogServer`](../../doc/models/gateway-mgmt-host-out-policy-syslog-server.md) | Optional | Per-server host-out path policies for gateway syslog traffic |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    gatewayMgmtHostOutPolicySyslog := models.GatewayMgmtHostOutPolicySyslog{
        PathPreference:       models.ToPointer("broadband_wans"),
        Servers:              []models.GatewayMgmtHostOutPolicySyslogServer{
            models.GatewayMgmtHostOutPolicySyslogServer{
                Host:                 models.ToPointer("host4"),
                PathPreference:       models.ToPointer("path_preference8"),
                ServerName:           models.ToPointer("server_name8"),
            },
        },
    }

}
```

