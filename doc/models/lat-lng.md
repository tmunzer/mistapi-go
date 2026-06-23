
# Lat Lng

Geographic latitude and longitude coordinate pair

## Structure

`LatLng`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Lat` | `float64` | Required | Geographic latitude in decimal degrees |
| `Lng` | `float64` | Required | Geographic longitude in decimal degrees |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    latLng := models.LatLng{
        Lat:                  float64(37.295833),
        Lng:                  float64(-122.032946),
    }

}
```

