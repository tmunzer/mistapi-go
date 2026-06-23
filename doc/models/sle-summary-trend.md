
# Sle Summary Trend

Time-series SLE summary trend response

## Structure

`SleSummaryTrend`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Classifiers` | [`[]models.SleTrendClassifier`](../../doc/models/sle-trend-classifier.md) | Required | SLE classifier trend details included in a summary trend response<br><br>**Constraints**: *Unique Items Required* |
| `End` | `float64` | Required | Last timestamp in the SLE summary trend window |
| `Sle` | [`models.SleSummarySle`](../../doc/models/sle-summary-sle.md) | Required | Time-series SLE metric samples for a summary response |
| `Start` | `float64` | Required | First timestamp in the SLE summary trend window |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    sleSummaryTrend := models.SleSummaryTrend{
        Classifiers:          []models.SleTrendClassifier{
            models.SleTrendClassifier{
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
        End:                  float64(217.08),
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
        Start:                float64(173.14),
    }

}
```

