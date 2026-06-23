
# Map Geofence Vertice

Vertex coordinate for a map geofence polygon

## Structure

`MapGeofenceVertice`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `X` | `*float64` | Optional | Geofence vertex X coordinate in map units |
| `Y` | `*float64` | Optional | Geofence vertex Y coordinate in map units |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    mapGeofenceVertice := models.MapGeofenceVertice{
        X:                    models.ToPointer(float64(700)),
        Y:                    models.ToPointer(float64(100)),
    }

}
```

