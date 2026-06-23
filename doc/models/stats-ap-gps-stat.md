
# Stats Ap Gps Stat

GPS location data reported by an AP

## Structure

`StatsApGpsStat`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Accuracy` | `*float64` | Optional | The estimated accuracy or accuracy of the GPS coordinates, measured in meters. |
| `Altitude` | `*float64` | Optional | The elevation of the AP above sea level, measured in meters. |
| `Latitude` | `*float64` | Optional | The geographic latitude of the AP, measured in degrees. |
| `Longitude` | `*float64` | Optional | The geographic longitude of the AP, measured in degrees. |
| `Src` | [`*models.StatsApGpsStatSrcEnum`](../../doc/models/stats-ap-gps-stat-src-enum.md) | Optional | The origin of the GPS data. enum: `gps`: from this device GPS estimates, `other_ap` from neighboring device GPS estimates. Note: API responses may return `other_aps` which should be treated as `other_ap` |
| `Timestamp` | `*float64` | Optional, Read-only | Epoch timestamp, in seconds |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    statsApGpsStat := models.StatsApGpsStat{
        Accuracy:             models.ToPointer(float64(12.5)),
        Altitude:             models.ToPointer(float64(99.939)),
        Latitude:             models.ToPointer(float64(37.29548)),
        Longitude:            models.ToPointer(float64(-122.03304)),
        Src:                  models.ToPointer(models.StatsApGpsStatSrcEnum_OTHERAPS),
    }

}
```

