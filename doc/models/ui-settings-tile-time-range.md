
# Ui Settings Tile Time Range

Time range override for a site UI databoard tile

## Structure

`UiSettingsTileTimeRange`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `End` | `*float64` | Optional | Upper bound of the tile time range, in epoch seconds |
| `EndDate` | `*string` | Optional | Display date for the end of the tile time range |
| `Interval` | `*string` | Optional | Bucket interval used for the tile time range |
| `Name` | `*string` | Optional | Display name for the tile time range |
| `ShortName` | `*string` | Optional | Compact display label for the tile time range |
| `Start` | `*int` | Optional | Lower bound of the tile time range, in epoch seconds |
| `UsePreset` | `*bool` | Optional | Whether the tile time range uses a named preset |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    uiSettingsTileTimeRange := models.UiSettingsTileTimeRange{
        End:                  models.ToPointer(float64(1508823743)),
        EndDate:              models.ToPointer("10/23/2017"),
        Interval:             models.ToPointer("1d"),
        Name:                 models.ToPointer("Past 7 Days"),
        ShortName:            models.ToPointer("7d"),
        Start:                models.ToPointer(1508223600),
        UsePreset:            models.ToPointer(true),
    }

}
```

