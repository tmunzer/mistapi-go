
# Response Device Metrics

Time-series insight metric response for a device

## Structure

`ResponseDeviceMetrics`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `End` | `int` | Required | Epoch timestamp for the end of the metric query window |
| `Interval` | `int` | Required | Aggregation interval in seconds for each metric sample |
| `Limit` | `*int` | Optional | Maximum number of metric samples returned in this page |
| `Page` | `*int` | Optional | Returned page number for paginated metric samples |
| `Results` | [`[]models.ResponseDeviceMetricsResultsItems`](../../doc/models/containers/response-device-metrics-results-items.md) | Required | Device metric result value, returned as a string or integer |
| `Rt` | `[]string` | Optional | Unique string values returned or accepted by this schema<br><br>**Constraints**: *Unique Items Required* |
| `Start` | `int` | Required | Epoch timestamp for the start of the metric query window |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    responseDeviceMetrics := models.ResponseDeviceMetrics{
        End:                  172,
        Interval:             82,
        Limit:                models.ToPointer(254),
        Page:                 models.ToPointer(112),
        Results:              []models.ResponseDeviceMetricsResultsItems{
            models.ResponseDeviceMetricsResultsItemsContainer.FromString("String6"),
            models.ResponseDeviceMetricsResultsItemsContainer.FromString("String7"),
        },
        Rt:                   []string{
            "rt0",
            "rt1",
        },
        Start:                130,
    }

}
```

