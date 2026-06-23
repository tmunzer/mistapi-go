
# Site Template Auto Upgrade

Automatic upgrade settings applied by a site template

## Structure

`SiteTemplateAutoUpgrade`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `DayOfWeek` | [`*models.DayOfWeekEnum`](../../doc/models/day-of-week-enum.md) | Optional | enum: `any`, `fri`, `mon`, `sat`, `sun`, `thu`, `tue`, `wed` |
| `Enabled` | `*bool` | Optional | Whether automatic upgrades are enabled for sites using this site template |
| `TimeOfDay` | `*string` | Optional | Local time of day when the automatic upgrade window starts |
| `Version` | `*string` | Optional | Target firmware version installed during automatic upgrades |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    siteTemplateAutoUpgrade := models.SiteTemplateAutoUpgrade{
        DayOfWeek:            models.ToPointer(models.DayOfWeekEnum_ANY),
        Enabled:              models.ToPointer(false),
        TimeOfDay:            models.ToPointer("time_of_day8"),
        Version:              models.ToPointer("version4"),
    }

}
```

