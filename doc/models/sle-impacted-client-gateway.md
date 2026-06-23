
# Sle Impacted Client Gateway

Gateway association for an impacted client

## Structure

`SleImpactedClientGateway`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `ChassisMac` | `*string` | Optional | MAC address of the chassis associated with this client gateway row |
| `GatewayMac` | `*string` | Optional | MAC address of the gateway associated with this client row |
| `GatewayName` | `*string` | Optional | Display name of the gateway associated with this client row |
| `Interfaces` | `[]string` | Optional | Unique string values returned or accepted by this schema<br><br>**Constraints**: *Unique Items Required* |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    sleImpactedClientGateway := models.SleImpactedClientGateway{
        ChassisMac:           models.ToPointer("chassis_mac0"),
        GatewayMac:           models.ToPointer("gateway_mac6"),
        GatewayName:          models.ToPointer("gateway_name4"),
        Interfaces:           []string{
            "interfaces9",
            "interfaces0",
        },
    }

}
```

