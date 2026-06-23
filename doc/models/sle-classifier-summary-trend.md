
# Sle Classifier Summary Trend

Time-series trend response for an SLE classifier

## Structure

`SleClassifierSummaryTrend`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Classifier` | [`models.SleTrendClassifier`](../../doc/models/sle-trend-classifier.md) | Required | Time-series classifier trend detail for an SLE metric |
| `End` | `float64` | Required | Last timestamp in the classifier trend window |
| `Metric` | `string` | Required | SLE metric name represented by the trend<br><br>**Constraints**: *Minimum Length*: `1` |
| `Start` | `float64` | Required | First timestamp in the classifier trend window |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    sleClassifierSummaryTrend := models.SleClassifierSummaryTrend{
        Classifier:           models.SleTrendClassifier{
            Interval:             float64(24.52),
            Name:                 "name4",
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
            XLabel:               "x_label0",
            YLabel:               "y_label2",
        },
        End:                  float64(200.6),
        Metric:               "metric6",
        Start:                float64(156.66),
    }

}
```

