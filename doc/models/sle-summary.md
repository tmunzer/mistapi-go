
# Sle Summary

Deprecated SLE summary response for a metric window

## Structure

`SleSummary`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Classifiers` | [`[]models.SleClassifier`](../../doc/models/sle-classifier.md) | Required | SLE classifier details included in a summary response<br><br>**Constraints**: *Unique Items Required* |
| `End` | `float64` | Required | Last timestamp in the SLE summary window |
| `Events` | `[]interface{}` | Required | Event records included in an SLE summary response |
| `Impact` | [`models.SleSummaryImpact`](../../doc/models/sle-summary-impact.md) | Required | AP and user impact counts for an SLE summary window |
| `Sle` | [`models.SleSummarySle`](../../doc/models/sle-summary-sle.md) | Required | Time-series SLE metric samples for a summary response |
| `Start` | `float64` | Required | First timestamp in the SLE summary window |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    sleSummary := models.SleSummary{
        Classifiers:          []models.SleClassifier{
            models.SleClassifier{
                Impact:               models.SleClassifierImpact{
                    NumAps:               float64(250.52),
                    NumUsers:             float64(4.56),
                    TotalAps:             float64(243.84),
                    TotalUsers:           float64(4.14),
                },
                Interval:             float64(217.3),
                Name:                 "name2",
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
                XLabel:               "x_label8",
                YLabel:               "y_label0",
            },
        },
        End:                  float64(20.6),
        Events:               []interface{}{
            interface{}("[key1, val1][key2, val2]"),
            interface{}("[key1, val1][key2, val2]"),
        },
        Impact:               models.SleSummaryImpact{
            NumAps:               float64(250.52),
            NumUsers:             float64(4.56),
            TotalAps:             float64(243.84),
            TotalUsers:           float64(4.14),
        },
        Sle:                  models.SleSummarySle{
            Interval:             float64(227.76),
            Name:                 "name8",
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
            XLabel:               "x_label4",
            YLabel:               "y_label6",
        },
        Start:                float64(232.66),
    }

}
```

