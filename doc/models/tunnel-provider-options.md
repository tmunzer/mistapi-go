
# Tunnel Provider Options

Provider-specific options for gateway tunnel auto provisioning

## Structure

`TunnelProviderOptions`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Jse` | [`*models.TunnelProviderOptionsJse`](../../doc/models/tunnel-provider-options-jse.md) | Optional | For jse-ipsec, this allows provisioning of adequate resource on JSE. Make sure adequate licenses are added |
| `Prisma` | [`*models.TunnelProviderOptionsPrisma`](../../doc/models/tunnel-provider-options-prisma.md) | Optional | Prisma Access provider options for tunnel auto provisioning |
| `Zscaler` | [`*models.TunnelProviderOptionsZscaler`](../../doc/models/tunnel-provider-options-zscaler.md) | Optional | For zscaler-ipsec and zscaler-gre |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    tunnelProviderOptions := models.TunnelProviderOptions{
        Jse:                  models.ToPointer(models.TunnelProviderOptionsJse{
            NumUsers:             models.ToPointer(186),
            OrgName:              models.ToPointer("org_name6"),
        }),
        Prisma:               models.ToPointer(models.TunnelProviderOptionsPrisma{
            ServiceAccountName:   models.ToPointer("service_account_name6"),
        }),
        Zscaler:              models.ToPointer(models.TunnelProviderOptionsZscaler{
            AupBlockInternetUntilAccepted:       models.ToPointer(false),
            AupEnabled:                          models.ToPointer(false),
            AupForceSslInspection:               models.ToPointer(false),
            AupTimeoutInDays:                    models.ToPointer(104),
            AuthRequired:                        models.ToPointer(false),
        }),
    }

}
```

