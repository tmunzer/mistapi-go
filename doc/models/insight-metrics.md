
# Insight Metrics

Insight metric response for a requested time range and aggregation interval

## Structure

`InsightMetrics`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `End` | `int` | Required | Window end timestamp for the returned insight metrics |
| `Interval` | `int` | Required | Aggregation interval used for the metric results |
| `Limit` | `*int` | Optional | Maximum number of insight metric result items returned |
| `Results` | [`[]models.InsightMetricsResultsItem`](../../doc/models/containers/insight-metrics-results-item.md) | Optional | Insight metric result item, returned either as a number or an object depending on the requested metric<br><br>**Constraints**: *Unique Items Required* |
| `Start` | `int` | Required | Window start timestamp for the returned insight metrics |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    insightMetrics := models.InsightMetrics{
        End:                  54,
        Interval:             220,
        Limit:                models.ToPointer(116),
        Results:              []models.InsightMetricsResultsItem{
            models.InsightMetricsResultsItemContainer.FromPrecision(float64(56.15)),
            models.InsightMetricsResultsItemContainer.FromPrecision(float64(56.16)),
        },
        Start:                12,
    }

}
```

