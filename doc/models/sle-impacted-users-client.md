
# Sle Impacted Users Client

SLE impact row for a client in an impacted users response

## Structure

`SleImpactedUsersClient`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Degraded` | `*float64` | Optional | Portion of the SLE total that was degraded for this client |
| `Duration` | `*float64` | Optional | Observation time represented by this client impact row |
| `Gateways` | [`[]models.SleImpactedClientGateway`](../../doc/models/sle-impacted-client-gateway.md) | Optional | Gateway associations for impacted clients |
| `Mac` | `*string` | Optional | Client MAC address for the impacted client |
| `Name` | `*string` | Optional | Display name for the client impact row |
| `SrcIp` | `*string` | Optional | Client source IP address for the impacted client |
| `Total` | `*float64` | Optional | Overall SLE total measured for this client impact row |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    sleImpactedUsersClient := models.SleImpactedUsersClient{
        Degraded:             models.ToPointer(float64(75.48)),
        Duration:             models.ToPointer(float64(204.54)),
        Gateways:             []models.SleImpactedClientGateway{
            models.SleImpactedClientGateway{
                ChassisMac:           models.ToPointer("chassis_mac4"),
                GatewayMac:           models.ToPointer("gateway_mac2"),
                GatewayName:          models.ToPointer("gateway_name0"),
                Interfaces:           []string{
                    "interfaces5",
                    "interfaces6",
                },
            },
            models.SleImpactedClientGateway{
                ChassisMac:           models.ToPointer("chassis_mac4"),
                GatewayMac:           models.ToPointer("gateway_mac2"),
                GatewayName:          models.ToPointer("gateway_name0"),
                Interfaces:           []string{
                    "interfaces5",
                    "interfaces6",
                },
            },
            models.SleImpactedClientGateway{
                ChassisMac:           models.ToPointer("chassis_mac4"),
                GatewayMac:           models.ToPointer("gateway_mac2"),
                GatewayName:          models.ToPointer("gateway_name0"),
                Interfaces:           []string{
                    "interfaces5",
                    "interfaces6",
                },
            },
        },
        Mac:                  models.ToPointer("mac2"),
        Name:                 models.ToPointer("name8"),
    }

}
```

