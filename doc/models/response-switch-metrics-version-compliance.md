
# Response Switch Metrics Version Compliance

Version compliance metric for switches in the requested scope

## Structure

`ResponseSwitchMetricsVersionCompliance`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Details` | [`*models.ResponseSwitchMetricsVersionComplianceDetails`](../../doc/models/response-switch-metrics-version-compliance-details.md) | Optional | Detail values for the switch software version compliance metric |
| `Score` | `*int` | Optional | Reported metric score for switch software version compliance |
| `TotalSwitchCount` | `*int` | Optional | Number of switches included in the version compliance metric |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    responseSwitchMetricsVersionCompliance := models.ResponseSwitchMetricsVersionCompliance{
        Details:              models.ToPointer(models.ResponseSwitchMetricsVersionComplianceDetails{
            MajorVersions:        []models.SwitchMetricsComplianceMajorVersion{
                models.SwitchMetricsComplianceMajorVersion{
                    MajorCount:           models.ToPointer(170),
                    MajorVersion:         models.ToPointer("major_version0"),
                    Model:                models.ToPointer("model6"),
                    SystemNames:          []string{
                        "system_names6",
                        "system_names7",
                    },
                },
            },
        }),
        Score:                models.ToPointer(96),
        TotalSwitchCount:     models.ToPointer(140),
    }

}
```

