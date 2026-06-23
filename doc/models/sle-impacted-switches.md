
# Sle Impacted Switches

Paginated list of switches impacted by an SLE metric

## Structure

`SleImpactedSwitches`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Classifier` | `*string` | Optional | Requested SLE classifier filter applied to the query |
| `End` | `*int` | Optional | Last timestamp in the impacted switches window |
| `Failure` | `*string` | Optional | Requested SLE failure filter applied to the query |
| `Limit` | `*int` | Optional | Maximum number of impacted switch rows returned per page |
| `Metric` | `*string` | Optional | SLE metric name used for the impacted switches query |
| `Page` | `*int` | Optional | Current page number for impacted switch results |
| `Start` | `*int` | Optional | First timestamp in the impacted switches window |
| `Switches` | [`[]models.SleImpactedSwitchesSwitch`](../../doc/models/sle-impacted-switches-switch.md) | Optional | Impacted switch rows returned for an SLE query |
| `TotalCount` | `*int` | Optional | Number of impacted switch rows matching the query |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    sleImpactedSwitches := models.SleImpactedSwitches{
        Classifier:           models.ToPointer("classifier4"),
        End:                  models.ToPointer(254),
        Failure:              models.ToPointer("failure2"),
        Limit:                models.ToPointer(172),
        Metric:               models.ToPointer("metric0"),
    }

}
```

