
# Sle Impacted Gateways Gateway

SLE impact row for a gateway

## Structure

`SleImpactedGatewaysGateway`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Degraded` | `*float64` | Optional | Portion of the SLE total that was degraded for this gateway |
| `Duration` | `*int` | Optional | Observation time represented by this gateway impact row |
| `GatewayMac` | `*string` | Optional | MAC address of the gateway represented by this impacted row |
| `GatewayModel` | `*string` | Optional | Model of the gateway represented by this impacted row |
| `GatewayVersion` | `*string` | Optional | Firmware version of the gateway represented by this impacted row |
| `Name` | `*string` | Optional | Display name for the gateway impact row |
| `Total` | `*int` | Optional | Overall SLE total measured for this gateway impact row |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    sleImpactedGatewaysGateway := models.SleImpactedGatewaysGateway{
        Degraded:             models.ToPointer(float64(176.52)),
        Duration:             models.ToPointer(94),
        GatewayMac:           models.ToPointer("gateway_mac0"),
        GatewayModel:         models.ToPointer("gateway_model4"),
        GatewayVersion:       models.ToPointer("gateway_version6"),
    }

}
```

