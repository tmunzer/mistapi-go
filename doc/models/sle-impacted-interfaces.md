
# Sle Impacted Interfaces

Paginated list of interfaces impacted by an SLE metric

## Structure

`SleImpactedInterfaces`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Classifier` | `*string` | Optional | Requested SLE classifier filter applied to the query |
| `End` | `*int` | Optional | Last timestamp in the impacted interfaces window |
| `Failure` | `*string` | Optional | Requested SLE failure filter applied to the query |
| `Interfaces` | [`[]models.SleImpactedInterfacesInterface`](../../doc/models/sle-impacted-interfaces-interface.md) | Optional | Impacted interface rows returned for an SLE query |
| `Limit` | `*int` | Optional | Maximum number of impacted interface rows returned per page |
| `Metric` | `*string` | Optional | SLE metric name used for the impacted interfaces query |
| `Page` | `*int` | Optional | Current page number for impacted interface results |
| `Start` | `*int` | Optional | First timestamp in the impacted interfaces window |
| `TotalCount` | `*int` | Optional | Number of impacted interface rows matching the query |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    sleImpactedInterfaces := models.SleImpactedInterfaces{
        Classifier:           models.ToPointer("classifier8"),
        End:                  models.ToPointer(92),
        Failure:              models.ToPointer("failure6"),
        Interfaces:           []models.SleImpactedInterfacesInterface{
            models.SleImpactedInterfacesInterface{
                Degraded:             models.ToPointer(float64(98.24)),
                Duration:             models.ToPointer(float64(227.3)),
                InterfaceName:        models.ToPointer("interface_name4"),
                SwitchMac:            models.ToPointer("switch_mac2"),
                SwitchName:           models.ToPointer("switch_name6"),
            },
        },
        Limit:                models.ToPointer(178),
    }

}
```

