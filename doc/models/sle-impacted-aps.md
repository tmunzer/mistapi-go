
# Sle Impacted Aps

Paginated list of APs impacted by an SLE metric

## Structure

`SleImpactedAps`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Aps` | [`[]models.SleImpactedApsAp`](../../doc/models/sle-impacted-aps-ap.md) | Required | Impacted AP rows returned for an SLE query<br><br>**Constraints**: *Unique Items Required* |
| `Classifier` | `string` | Required | Requested SLE classifier filter applied to the query |
| `End` | `float64` | Required | Last timestamp in the impacted APs window |
| `Failure` | `string` | Required | Requested SLE failure filter applied to the query |
| `Limit` | `float64` | Required | Maximum number of impacted AP rows returned per page |
| `Metric` | `string` | Required | SLE metric name used for the impacted APs query<br><br>**Constraints**: *Minimum Length*: `1` |
| `Page` | `float64` | Required | Current page number for impacted AP results |
| `Start` | `float64` | Required | First timestamp in the impacted APs window |
| `TotalCount` | `float64` | Required | Number of impacted AP rows matching the query |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    sleImpactedAps := models.SleImpactedAps{
        Aps:                  []models.SleImpactedApsAp{
            models.SleImpactedApsAp{
                ApMac:                "ap_mac6",
                Degraded:             float64(30.24),
                Duration:             float64(159.3),
                Name:                 "name4",
                Total:                float64(197.76),
            },
        },
        Classifier:           "classifier2",
        End:                  float64(119.64),
        Failure:              "failure0",
        Limit:                float64(125.62),
        Metric:               "metric8",
        Page:                 float64(3.52),
        Start:                float64(75.7),
        TotalCount:           float64(189.82),
    }

}
```

