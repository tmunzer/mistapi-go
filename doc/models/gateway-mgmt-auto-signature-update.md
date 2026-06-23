
# Gateway Mgmt Auto Signature Update

Automatic security signature update schedule

## Structure

`GatewayMgmtAutoSignatureUpdate`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `DayOfWeek` | [`*models.DayOfWeekEnum`](../../doc/models/day-of-week-enum.md) | Optional | enum: `any`, `fri`, `mon`, `sat`, `sun`, `thu`, `tue`, `wed` |
| `Enable` | `*bool` | Optional | Whether automatic security signature updates are enabled<br><br>**Default**: `true` |
| `TimeOfDay` | `*string` | Optional | Optional, Mist will decide the timing |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    gatewayMgmtAutoSignatureUpdate := models.GatewayMgmtAutoSignatureUpdate{
        DayOfWeek:            models.ToPointer(models.DayOfWeekEnum_ANY),
        Enable:               models.ToPointer(true),
        TimeOfDay:            models.ToPointer("time_of_day6"),
    }

}
```

