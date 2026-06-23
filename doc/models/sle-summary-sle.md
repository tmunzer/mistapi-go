
# Sle Summary Sle

Time-series SLE metric samples for a summary response

## Structure

`SleSummarySle`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Interval` | `float64` | Required | Sample interval in seconds for SLE summary data |
| `Name` | `string` | Required | SLE metric name represented by the summary samples<br><br>**Constraints**: *Minimum Length*: `1` |
| `Samples` | [`models.SleSummarySleSamples`](../../doc/models/sle-summary-sle-samples.md) | Required | Per-interval sample arrays for an SLE summary |
| `XLabel` | `string` | Required | Label for the SLE summary x-axis<br><br>**Constraints**: *Minimum Length*: `1` |
| `YLabel` | `string` | Required | Label for the SLE summary y-axis<br><br>**Constraints**: *Minimum Length*: `1` |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    sleSummarySle := models.SleSummarySle{
        Interval:             float64(237.88),
        Name:                 "name0",
        Samples:              models.SleSummarySleSamples{
            Degraded:             []models.NumberOrNull{
                models.NumberOrNullContainer.FromPrecision(float64(170.03)),
            },
            Total:                []models.NumberOrNull{
                models.NumberOrNullContainer.FromPrecision(float64(144.95)),
            },
            Value:                []models.NumberOrNull{
                models.NumberOrNullContainer.FromPrecision(float64(13.46)),
                models.NumberOrNullContainer.FromPrecision(float64(13.47)),
                models.NumberOrNullContainer.FromPrecision(float64(13.48)),
            },
        },
        XLabel:               "x_label6",
        YLabel:               "y_label8",
    }

}
```

