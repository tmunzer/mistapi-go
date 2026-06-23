
# Sle Impacted Applications

Paginated list of applications impacted by an SLE metric

## Structure

`SleImpactedApplications`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Apps` | [`[]models.SleImpactedApplicationsApp`](../../doc/models/sle-impacted-applications-app.md) | Optional | Impacted application rows returned for an SLE query |
| `Classifier` | `*string` | Optional | Requested SLE classifier filter applied to the query |
| `End` | `*int` | Optional | Last timestamp in the impacted applications window |
| `Failure` | `*string` | Optional | Requested SLE failure filter applied to the query |
| `Limit` | `*string` | Optional | Maximum number of impacted application rows returned per page |
| `Metric` | `*string` | Optional | SLE metric name used for the impacted applications query |
| `Page` | `*int` | Optional | Current page number for impacted application results |
| `Start` | `*int` | Optional | First timestamp in the impacted applications window |
| `TotalCount` | `*int` | Optional | Number of impacted application rows matching the query |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    sleImpactedApplications := models.SleImpactedApplications{
        Apps:                 []models.SleImpactedApplicationsApp{
            models.SleImpactedApplicationsApp{
                App:                  models.ToPointer("app4"),
                Degraded:             models.ToPointer(58),
                Duration:             models.ToPointer(164),
                Name:                 models.ToPointer("name6"),
                Threshold:            models.ToPointer(20),
            },
            models.SleImpactedApplicationsApp{
                App:                  models.ToPointer("app4"),
                Degraded:             models.ToPointer(58),
                Duration:             models.ToPointer(164),
                Name:                 models.ToPointer("name6"),
                Threshold:            models.ToPointer(20),
            },
        },
        Classifier:           models.ToPointer("classifier8"),
        End:                  models.ToPointer(114),
        Failure:              models.ToPointer("failure6"),
        Limit:                models.ToPointer("limit2"),
    }

}
```

