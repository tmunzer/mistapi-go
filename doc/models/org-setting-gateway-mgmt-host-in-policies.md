
# Org Setting Gateway Mgmt Host in Policies

Host-in access policies for gateway management services

## Structure

`OrgSettingGatewayMgmtHostInPolicies`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Icmp` | [`*models.OrgSettingGatewayMgmtHostInPolicy`](../../doc/models/org-setting-gateway-mgmt-host-in-policy.md) | Optional | Host-in access policy for a gateway management service |
| `Snmp` | [`*models.OrgSettingGatewayMgmtHostInPolicy`](../../doc/models/org-setting-gateway-mgmt-host-in-policy.md) | Optional | Host-in access policy for a gateway management service |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    orgSettingGatewayMgmtHostInPolicies := models.OrgSettingGatewayMgmtHostInPolicies{
        Icmp:                 models.ToPointer(models.OrgSettingGatewayMgmtHostInPolicy{
            Tenants:              []string{
                "tenants3",
            },
        }),
        Snmp:                 models.ToPointer(models.OrgSettingGatewayMgmtHostInPolicy{
            Tenants:              []string{
                "tenants5",
            },
        }),
    }

}
```

