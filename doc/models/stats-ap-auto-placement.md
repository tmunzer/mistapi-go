
# Stats Ap Auto Placement

Auto placement result and status for an AP

## Structure

`StatsApAutoPlacement`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Info` | [`*models.StatsApAutoPlacementInfo`](../../doc/models/stats-ap-auto-placement-info.md) | Optional | Additional diagnostics for AP auto placement |
| `RecommendedAnchor` | `*bool` | Optional | Flag to represent if AP is recommended as an anchor by auto placement service |
| `Status` | `*string` | Optional | Current auto placement status for the AP |
| `StatusDetail` | `*string` | Optional | Additional info about placement status |
| `X` | `*float64` | Optional | Auto-placed X coordinate, in pixels |
| `XM` | `*float64` | Optional | Auto-placed X coordinate, in meters |
| `Y` | `*float64` | Optional | Auto-placed Y coordinate, in pixels |
| `YM` | `*float64` | Optional | Auto-placed Y coordinate, in meters |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    statsApAutoPlacement := models.StatsApAutoPlacement{
        Info:                 models.ToPointer(models.StatsApAutoPlacementInfo{
            ClusterNumber:        models.ToPointer(112),
            OrientationStats:     models.ToPointer(90),
            ProbabilitySurface:   models.ToPointer(models.StatsApAutoPlacementInfoProbabilitySurface{
                Radius:               models.ToPointer(float64(74.96)),
                RadiusM:              models.ToPointer(float64(19.46)),
                X:                    models.ToPointer(float64(93.54)),
            }),
        }),
        RecommendedAnchor:    models.ToPointer(false),
        Status:               models.ToPointer("localized"),
        StatusDetail:         models.ToPointer("localized"),
        X:                    models.ToPointer(float64(53.5)),
        XM:                   models.ToPointer(float64(5.35)),
        Y:                    models.ToPointer(float64(173.1)),
        YM:                   models.ToPointer(float64(17.31)),
    }

}
```

