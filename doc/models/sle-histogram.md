
# Sle Histogram

Histogram response for an SLE metric

## Structure

`SleHistogram`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Data` | [`[]models.SleHistogramDataItem`](../../doc/models/sle-histogram-data-item.md) | Required | Histogram buckets for an SLE metric<br><br>**Constraints**: *Unique Items Required* |
| `End` | `float64` | Required | Last timestamp in the histogram window |
| `Metric` | `string` | Required | SLE metric name represented by the histogram<br><br>**Constraints**: *Minimum Length*: `1` |
| `Start` | `float64` | Required | First timestamp in the histogram window |
| `XLabel` | `string` | Required | Label for the histogram x-axis<br><br>**Constraints**: *Minimum Length*: `1` |
| `YLabel` | `string` | Required | Label for the histogram y-axis<br><br>**Constraints**: *Minimum Length*: `1` |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    sleHistogram := models.SleHistogram{
        Data:                 []models.SleHistogramDataItem{
            models.SleHistogramDataItem{
                Range:                []float64{
                    float64(7.03),
                    float64(7.04),
                },
                Value:                float64(95.62),
            },
        },
        End:                  float64(232.96),
        Metric:               "metric0",
        Start:                float64(189.02),
        XLabel:               "x_label4",
        YLabel:               "y_label6",
    }

}
```

