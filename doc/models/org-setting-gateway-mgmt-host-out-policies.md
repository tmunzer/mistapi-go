
# Org Setting Gateway Mgmt Host Out Policies

Optional path preferences for gateway-originated management traffic; ECMP is used across available paths when no preference is specified

## Structure

`OrgSettingGatewayMgmtHostOutPolicies`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Dns` | [`*models.GatewayMgmtHostOutPolicy`](../../doc/models/gateway-mgmt-host-out-policy.md) | Optional | Host-out path policy for gateway-originated management traffic |
| `Ntp` | [`*models.GatewayMgmtHostOutPolicy`](../../doc/models/gateway-mgmt-host-out-policy.md) | Optional | Host-out path policy for gateway-originated management traffic |
| `Syslog` | [`*models.GatewayMgmtHostOutPolicySyslog`](../../doc/models/gateway-mgmt-host-out-policy-syslog.md) | Optional | Host-out path policy for gateway syslog traffic |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    orgSettingGatewayMgmtHostOutPolicies := models.OrgSettingGatewayMgmtHostOutPolicies{
        Dns:                  models.ToPointer(models.GatewayMgmtHostOutPolicy{
            PathPreference:       models.ToPointer("path_preference8"),
        }),
        Ntp:                  models.ToPointer(models.GatewayMgmtHostOutPolicy{
            PathPreference:       models.ToPointer("path_preference4"),
        }),
        Syslog:               models.ToPointer(models.GatewayMgmtHostOutPolicySyslog{
            PathPreference:       models.ToPointer("path_preference2"),
            Servers:              []models.GatewayMgmtHostOutPolicySyslogServer{
                models.GatewayMgmtHostOutPolicySyslogServer{
                    Host:                 models.ToPointer("host4"),
                    PathPreference:       models.ToPointer("path_preference8"),
                    ServerName:           models.ToPointer("server_name8"),
                },
            },
        }),
    }

}
```

