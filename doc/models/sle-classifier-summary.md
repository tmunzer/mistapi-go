
# Sle Classifier Summary

Deprecated SLE classifier summary response

## Structure

`SleClassifierSummary`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Classifier` | [`models.SleClassifier`](../../doc/models/sle-classifier.md) | Required | Time-series classifier detail for an SLE metric |
| `End` | `float64` | Required | Last timestamp in the classifier summary window |
| `Failures` | `[]interface{}` | Required | Failure records included in a deprecated classifier summary |
| `Impact` | [`models.SleClassifierSummaryImpact`](../../doc/models/sle-classifier-summary-impact.md) | Required | AP and user impact counts for a classifier summary window |
| `Metric` | `string` | Required | SLE metric name summarized by this response<br><br>**Constraints**: *Minimum Length*: `1` |
| `Start` | `float64` | Required | First timestamp in the classifier summary window |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    sleClassifierSummary := models.SleClassifierSummary{
        Classifier:           models.SleClassifier{
            Impact:               models.SleClassifierImpact{
                NumAps:               float64(250.52),
                NumUsers:             float64(4.56),
                TotalAps:             float64(243.84),
                TotalUsers:           float64(4.14),
            },
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
        End:                  float64(216.22),
        Failures:             []interface{}{
            interface{}("[key1, val1][key2, val2]"),
        },
        Impact:               models.SleClassifierSummaryImpact{
            NumAps:               float64(250.52),
            NumUsers:             float64(4.56),
            TotalAps:             float64(243.84),
            TotalUsers:           float64(4.14),
        },
        Metric:               "metric6",
        Start:                float64(172.28),
    }

}
```

