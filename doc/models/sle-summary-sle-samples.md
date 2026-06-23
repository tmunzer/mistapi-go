
# Sle Summary Sle Samples

Per-interval sample arrays for an SLE summary

## Structure

`SleSummarySleSamples`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Degraded` | [`[]models.NumberOrNull`](../../doc/models/containers/number-or-null.md) | Required | Numeric value or null |
| `Total` | [`[]models.NumberOrNull`](../../doc/models/containers/number-or-null.md) | Required | Numeric value or null |
| `Value` | [`[]models.NumberOrNull`](../../doc/models/containers/number-or-null.md) | Required | Numeric value or null |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    sleSummarySleSamples := models.SleSummarySleSamples{
        Degraded:             []models.NumberOrNull{
            models.NumberOrNullContainer.FromPrecision(float64(137.33)),
            models.NumberOrNullContainer.FromPrecision(float64(137.34)),
        },
        Total:                []models.NumberOrNull{
            models.NumberOrNullContainer.FromPrecision(float64(51.43)),
            models.NumberOrNullContainer.FromPrecision(float64(51.42)),
            models.NumberOrNullContainer.FromPrecision(float64(51.41)),
        },
        Value:                []models.NumberOrNull{
            models.NumberOrNullContainer.FromPrecision(float64(236.76)),
        },
    }

}
```

