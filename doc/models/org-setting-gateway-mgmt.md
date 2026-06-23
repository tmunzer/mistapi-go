
# Org Setting Gateway Mgmt

Organization-level gateway management settings

## Structure

`OrgSettingGatewayMgmt`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AppProbing` | [`*models.OrgSettingGatewayMgmtAppProbing`](../../doc/models/org-setting-gateway-mgmt-app-probing.md) | Optional | Application probing settings for organization gateway management |
| `AppUsage` | `*bool` | Optional | For SRX only, whether gateway application usage collection is enabled; requires App Track license |
| `FipsEnabled` | `*bool` | Optional | Whether FIPS mode is enabled for managed gateways<br><br>**Default**: `false` |
| `HostInPolicies` | [`*models.OrgSettingGatewayMgmtHostInPolicies`](../../doc/models/org-setting-gateway-mgmt-host-in-policies.md) | Optional | Host-in access policies for gateway management services |
| `HostOutPolicies` | [`*models.OrgSettingGatewayMgmtHostOutPolicies`](../../doc/models/org-setting-gateway-mgmt-host-out-policies.md) | Optional | Optional path preferences for gateway-originated management traffic; ECMP is used across available paths when no preference is specified |
| `OverlayIp` | [`*models.OrgSettingGatewayMgmtOverlayIp`](../../doc/models/org-setting-gateway-mgmt-overlay-ip.md) | Optional | Overlay IP configuration used for gateway management traffic |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    orgSettingGatewayMgmt := models.OrgSettingGatewayMgmt{
        AppProbing:           models.ToPointer(models.OrgSettingGatewayMgmtAppProbing{
            Apps:                 []string{
                "apps8",
                "apps9",
                "apps0",
            },
        }),
        AppUsage:             models.ToPointer(false),
        FipsEnabled:          models.ToPointer(false),
        HostInPolicies:       models.ToPointer(models.OrgSettingGatewayMgmtHostInPolicies{
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
        }),
        HostOutPolicies:      models.ToPointer(models.OrgSettingGatewayMgmtHostOutPolicies{
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
        }),
    }

}
```

