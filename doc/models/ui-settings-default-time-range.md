
# Ui Settings Default Time Range

Default time range applied to a site UI databoard

## Structure

`UiSettingsDefaultTimeRange`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `End` | `*int` | Optional | Upper bound of the default databoard time range, in epoch seconds |
| `EndDate` | `*string` | Optional | Display date for the end of the default databoard time range |
| `Interval` | `*string` | Optional | Bucket interval used for the default databoard time range |
| `Name` | `*string` | Optional | Display name for the default databoard time range |
| `ShortName` | `*string` | Optional | Compact display label for the default databoard time range |
| `Start` | `*int` | Optional | Lower bound of the default databoard time range, in epoch seconds |
| `UsePreset` | `*bool` | Optional | Whether the default databoard time range uses a named preset |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    uiSettingsDefaultTimeRange := models.UiSettingsDefaultTimeRange{
        End:                  models.ToPointer(1508828400),
        EndDate:              models.ToPointer("10/23/2017"),
        Interval:             models.ToPointer("1d"),
        Name:                 models.ToPointer("This Week"),
        ShortName:            models.ToPointer("thisWeek"),
        Start:                models.ToPointer(1508655600),
        UsePreset:            models.ToPointer(true),
    }

}
```

