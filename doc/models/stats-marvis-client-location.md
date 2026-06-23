
# Stats Marvis Client Location

Last known location fix for a Marvis Client device

## Structure

`StatsMarvisClientLocation`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `MapId` | `*uuid.UUID` | Optional | UUID of the floor-plan map |
| `SiteId` | `*uuid.UUID` | Optional | UUID of the site the device was located in |
| `Timestamp` | `*int` | Optional | Timestamp of the location fix, in epoch seconds |
| `X` | `*float64` | Optional | X coordinate on the floor-plan map, in pixels |
| `Y` | `*float64` | Optional | Y coordinate on the floor-plan map, in pixels |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    statsMarvisClientLocation := models.StatsMarvisClientLocation{
        MapId:                models.ToPointer(uuid.MustParse("00000252-0000-0000-0000-000000000000")),
        SiteId:               models.ToPointer(uuid.MustParse("00000c40-0000-0000-0000-000000000000")),
        Timestamp:            models.ToPointer(56),
        X:                    models.ToPointer(float64(10.84)),
        Y:                    models.ToPointer(float64(120.44)),
    }

}
```

