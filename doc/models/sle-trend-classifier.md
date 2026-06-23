
# Sle Trend Classifier

Time-series classifier trend detail for an SLE metric

## Structure

`SleTrendClassifier`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Interval` | `float64` | Required | Sample interval in seconds for classifier trend data |
| `Name` | `string` | Required | Classifier name for the SLE metric trend<br><br>**Constraints**: *Minimum Length*: `1` |
| `Samples` | [`*models.SleClassifierSamples`](../../doc/models/sle-classifier-samples.md) | Optional | Per-interval sample arrays for an SLE classifier |
| `XLabel` | `string` | Required | Label for the classifier trend x-axis<br><br>**Constraints**: *Minimum Length*: `1` |
| `YLabel` | `string` | Required | Label for the classifier trend y-axis<br><br>**Constraints**: *Minimum Length*: `1` |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    sleTrendClassifier := models.SleTrendClassifier{
        Interval:             float64(182.58),
        Name:                 "name0",
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
        XLabel:               "x_label4",
        YLabel:               "y_label8",
    }

}
```

