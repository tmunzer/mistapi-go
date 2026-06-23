
# Stats Ap Auto Placement Info

Additional diagnostics for AP auto placement

## Structure

`StatsApAutoPlacementInfo`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `ClusterNumber` | `*int` | Optional | All APs sharing a given cluster number can be placed relative to each other |
| `OrientationStats` | `*int` | Optional | The orientation of an AP |
| `ProbabilitySurface` | [`*models.StatsApAutoPlacementInfoProbabilitySurface`](../../doc/models/stats-ap-auto-placement-info-probability-surface.md) | Optional | Circular uncertainty area for an AP auto placement result |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    statsApAutoPlacementInfo := models.StatsApAutoPlacementInfo{
        ClusterNumber:        models.ToPointer(0),
        OrientationStats:     models.ToPointer(0),
        ProbabilitySurface:   models.ToPointer(models.StatsApAutoPlacementInfoProbabilitySurface{
            Radius:               models.ToPointer(float64(74.96)),
            RadiusM:              models.ToPointer(float64(19.46)),
            X:                    models.ToPointer(float64(93.54)),
        }),
    }

}
```

