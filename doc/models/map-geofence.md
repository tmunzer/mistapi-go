
# Map Geofence

Geofence drawn on a map

## Structure

`MapGeofence`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Name` | `*string` | Optional | Display name for the map geofence |
| `Vertices` | [`[]models.MapGeofenceVertice`](../../doc/models/map-geofence-vertice.md) | Optional | List of vertices defining the geofence |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    mapGeofence := models.MapGeofence{
        Name:                 models.ToPointer("example"),
        Vertices:             []models.MapGeofenceVertice{
            models.MapGeofenceVertice{
                X:                    models.ToPointer(float64(86.66)),
                Y:                    models.ToPointer(float64(252.2)),
            },
        },
    }

}
```

