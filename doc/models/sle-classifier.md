
# Sle Classifier

Time-series classifier detail for an SLE metric

## Structure

`SleClassifier`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Impact` | [`models.SleClassifierImpact`](../../doc/models/sle-classifier-impact.md) | Required | AP and user impact counts for an SLE classifier |
| `Interval` | `float64` | Required | Sample interval in seconds for classifier time-series data |
| `Name` | `string` | Required | Classifier name for the SLE metric<br><br>**Constraints**: *Minimum Length*: `1` |
| `Samples` | [`*models.SleClassifierSamples`](../../doc/models/sle-classifier-samples.md) | Optional | Per-interval sample arrays for an SLE classifier |
| `XLabel` | `string` | Required | Label for the classifier sample x-axis<br><br>**Constraints**: *Minimum Length*: `1` |
| `YLabel` | `string` | Required | Label for the classifier sample y-axis<br><br>**Constraints**: *Minimum Length*: `1` |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    sleClassifier := models.SleClassifier{
        Impact:               models.SleClassifierImpact{
            NumAps:               float64(250.52),
            NumUsers:             float64(4.56),
            TotalAps:             float64(243.84),
            TotalUsers:           float64(4.14),
        },
        Interval:             float64(210.96),
        Name:                 "name8",
        Samples:              models.ToPointer(models.SleClassifierSamples{
            Degraded:             []models.NumberOrNull{
                models.NumberOrNullContainer.FromPrecision(float64(170.03)),
            },
            Duration:             []float64{
                float64(249.08),
                float64(249.09),
                float64(249.1),
            },
            Total:                []models.NumberOrNull{
                models.NumberOrNullContainer.FromPrecision(float64(144.95)),
            },
        }),
        XLabel:               "x_label6",
        YLabel:               "y_label6",
    }

}
```

