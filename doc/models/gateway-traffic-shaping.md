
# Gateway Traffic Shaping

Traffic shaping settings for a gateway interface or VPN path

## Structure

`GatewayTrafficShaping`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `ClassPercentages` | `[]int` | Optional | percentages for different class of traffic: high / medium / low / best-effort. Sum must be equal to 100 |
| `Enabled` | `*bool` | Optional | Whether traffic shaping is enabled<br><br>**Default**: `false` |
| `MaxTxKbps` | `*int` | Optional | Maximum transmit bandwidth for the interface, in Kbps |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    gatewayTrafficShaping := models.GatewayTrafficShaping{
        ClassPercentages:     []int{
            237,
            236,
            235,
        },
        Enabled:              models.ToPointer(false),
        MaxTxKbps:            models.ToPointer(30),
    }

}
```

