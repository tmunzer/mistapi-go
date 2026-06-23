
# Wlan Schedule

WLAN operating schedule, default is disabled

## Structure

`WlanSchedule`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Enabled` | `*bool` | Optional | Whether the WLAN operating schedule is enabled<br><br>**Default**: `false` |
| `Hours` | [`*models.Hours`](../../doc/models/hours.md) | Optional | Day-of-week operating hour filters using hour ranges such as 09:00-17:00 |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    wlanSchedule := models.WlanSchedule{
        Enabled:              models.ToPointer(false),
        Hours:                models.ToPointer(models.Hours{
            Fri:                  models.ToPointer("fri2"),
            Mon:                  models.ToPointer("mon8"),
            Sat:                  models.ToPointer("sat0"),
            Sun:                  models.ToPointer("sun6"),
            Thu:                  models.ToPointer("thu6"),
        }),
    }

}
```

