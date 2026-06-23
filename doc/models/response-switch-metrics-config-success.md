
# Response Switch Metrics Config Success

Configuration success metric for switches in the requested scope

## Structure

`ResponseSwitchMetricsConfigSuccess`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Details` | [`*models.ResponseSwitchMetricsConfigSuccessDetails`](../../doc/models/response-switch-metrics-config-success-details.md) | Optional | Detail values for the switch configuration success metric |
| `Score` | `*int` | Optional | Reported metric score for switch configuration success |
| `TotalSwitchCount` | `*int` | Optional | Number of switches included in the configuration success metric |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    responseSwitchMetricsConfigSuccess := models.ResponseSwitchMetricsConfigSuccess{
        Details:              models.ToPointer(models.ResponseSwitchMetricsConfigSuccessDetails{
            ConfigSuccessCount:   models.ToPointer(166),
        }),
        Score:                models.ToPointer(56),
        TotalSwitchCount:     models.ToPointer(156),
    }

}
```

