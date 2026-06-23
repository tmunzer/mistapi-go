
# Sle Impacted Chassis

Paginated list of chassis impacted by an SLE metric

## Structure

`SleImpactedChassis`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Chassis` | [`[]models.SleImpactedChassisChassisItem`](../../doc/models/sle-impacted-chassis-chassis-item.md) | Optional | Impacted chassis rows returned for an SLE query |
| `Classifier` | `*string` | Optional | Requested SLE classifier filter applied to the query |
| `End` | `*int` | Optional | Last timestamp in the impacted chassis window |
| `Failure` | `*string` | Optional | Requested SLE failure filter applied to the query |
| `Limit` | `*int` | Optional | Maximum number of impacted chassis rows returned per page |
| `Metric` | `*string` | Optional | SLE metric name used for the impacted chassis query |
| `Page` | `*int` | Optional | Current page number for impacted chassis results |
| `Start` | `*int` | Optional | First timestamp in the impacted chassis window |
| `TotalCount` | `*int` | Optional | Number of impacted chassis rows matching the query |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    sleImpactedChassis := models.SleImpactedChassis{
        Chassis:              []models.SleImpactedChassisChassisItem{
            models.SleImpactedChassisChassisItem{
                Chassis:              models.ToPointer("chassis0"),
                Degraded:             models.ToPointer(float64(171.46)),
                Duration:             models.ToPointer(float64(44.52)),
                Role:                 models.ToPointer("role0"),
                SwitchMac:            models.ToPointer("switch_mac4"),
            },
            models.SleImpactedChassisChassisItem{
                Chassis:              models.ToPointer("chassis0"),
                Degraded:             models.ToPointer(float64(171.46)),
                Duration:             models.ToPointer(float64(44.52)),
                Role:                 models.ToPointer("role0"),
                SwitchMac:            models.ToPointer("switch_mac4"),
            },
        },
        Classifier:           models.ToPointer("classifier0"),
        End:                  models.ToPointer(52),
        Failure:              models.ToPointer("failure8"),
        Limit:                models.ToPointer(118),
    }

}
```

