
# Site Setting Auto Placement

Automatically determined AP placement coordinates and orientation

## Structure

`SiteSettingAutoPlacement`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Orientation` | `*int` | Optional | AP orientation angle in degrees on the map |
| `X` | `*float64` | Optional | Map x-coordinate determined by auto placement |
| `Y` | `*float64` | Optional | Map y-coordinate determined by auto placement |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    siteSettingAutoPlacement := models.SiteSettingAutoPlacement{
        Orientation:          models.ToPointer(45),
        X:                    models.ToPointer(float64(30)),
        Y:                    models.ToPointer(float64(60)),
    }

}
```

