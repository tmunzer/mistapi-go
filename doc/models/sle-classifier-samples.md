
# Sle Classifier Samples

Per-interval sample arrays for an SLE classifier

## Structure

`SleClassifierSamples`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Degraded` | [`[]models.NumberOrNull`](../../doc/models/containers/number-or-null.md) | Required | Numeric value or null |
| `Duration` | `[]float64` | Required | Per-interval observation durations for a classifier |
| `Total` | [`[]models.NumberOrNull`](../../doc/models/containers/number-or-null.md) | Required | Numeric value or null |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    sleClassifierSamples := models.SleClassifierSamples{
        Degraded:             []models.NumberOrNull{
            models.NumberOrNullContainer.FromPrecision(float64(57.69)),
        },
        Duration:             []float64{
            float64(136.74),
            float64(136.75),
            float64(136.76),
        },
        Total:                []models.NumberOrNull{
            models.NumberOrNullContainer.FromPrecision(float64(129.39)),
        },
    }

}
```

