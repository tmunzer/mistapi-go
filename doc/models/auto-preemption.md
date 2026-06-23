
# Auto Preemption

Schedule to preempt AP tunnels that are not connected to their preferred peer

## Structure

`AutoPreemption`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `DayOfWeek` | [`*models.DayOfWeekEnum`](../../doc/models/day-of-week-enum.md) | Optional | enum: `any`, `fri`, `mon`, `sat`, `sun`, `thu`, `tue`, `wed` |
| `Enabled` | `*bool` | Optional | Whether auto preemption is enabled<br><br>**Default**: `false` |
| `TimeOfDay` | `*string` | Optional | `any` / HH:MM (24-hour format)<br><br>**Default**: `"any"` |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    autoPreemption := models.AutoPreemption{
        DayOfWeek:            models.ToPointer(models.DayOfWeekEnum_TUE),
        Enabled:              models.ToPointer(false),
        TimeOfDay:            models.ToPointer("12:00"),
    }

}
```

