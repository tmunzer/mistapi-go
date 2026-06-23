
# Stats Ap Radio Config Band

Per-band radio configuration reported by an AP

## Structure

`StatsApRadioConfigBand`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AllowRrmDisable` | `models.Optional[bool]` | Optional, Read-only | Whether RRM disablement is allowed for this radio band |
| `Bandwidth` | `models.Optional[float64]` | Optional, Read-only | Channel bandwidth configured for this radio band |
| `Channel` | `*int` | Optional | Operating channel configured for this radio band |
| `Disabled` | `models.Optional[bool]` | Optional, Read-only | Whether this radio band is disabled |
| `DynamicChainingEnabled` | `models.Optional[bool]` | Optional, Read-only | Whether dynamic chaining is enabled for this radio band |
| `Power` | `models.Optional[float64]` | Optional, Read-only | Transmit power configured for this radio band |
| `PowerMax` | `models.Optional[float64]` | Optional, Read-only | Maximum transmit power allowed for this radio band |
| `PowerMin` | `models.Optional[float64]` | Optional, Read-only | Minimum transmit power allowed for this radio band |
| `RxChain` | `models.Optional[int]` | Optional, Read-only | Number of receive chains enabled for this radio band |
| `TxChain` | `models.Optional[int]` | Optional, Read-only | Number of transmit chains enabled for this radio band |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    statsApRadioConfigBand := models.StatsApRadioConfigBand{
        AllowRrmDisable:        models.NewOptional(models.ToPointer(false)),
        Bandwidth:              models.NewOptional(models.ToPointer(float64(20))),
        Channel:                models.ToPointer(1),
        Disabled:               models.NewOptional(models.ToPointer(false)),
        DynamicChainingEnabled: models.NewOptional(models.ToPointer(false)),
        Power:                  models.NewOptional(models.ToPointer(float64(10))),
        PowerMax:               models.NewOptional(models.ToPointer(float64(10))),
        PowerMin:               models.NewOptional(models.ToPointer(float64(10))),
        RxChain:                models.NewOptional(models.ToPointer(4)),
        TxChain:                models.NewOptional(models.ToPointer(4)),
    }

}
```

