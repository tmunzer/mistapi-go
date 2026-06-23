
# Stats Ap Auto Placement Info Probability Surface

Circular uncertainty area for an AP auto placement result

## Structure

`StatsApAutoPlacementInfoProbabilitySurface`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Radius` | `*float64` | Optional | The RADIUS representing placement uncertainty, measured in pixels |
| `RadiusM` | `*float64` | Optional | The RADIUS representing placement uncertainty, measured in meters |
| `X` | `*float64` | Optional | Potential placement center X coordinate, in pixels |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    statsApAutoPlacementInfoProbabilitySurface := models.StatsApAutoPlacementInfoProbabilitySurface{
        Radius:               models.ToPointer(float64(2.1)),
        RadiusM:              models.ToPointer(float64(75.58)),
        X:                    models.ToPointer(float64(17)),
    }

}
```

