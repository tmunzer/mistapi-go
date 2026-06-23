
# Response Switch Metrics

Switch metrics returned for the requested site or switch scope

## Structure

`ResponseSwitchMetrics`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `ActivePortsSummary` | [`*models.ResponseSwitchMetricsActivePortsSummary`](../../doc/models/response-switch-metrics-active-ports-summary.md) | Optional | Active-port summary metric for switches in the requested scope |
| `ConfigSuccess` | [`*models.ResponseSwitchMetricsConfigSuccess`](../../doc/models/response-switch-metrics-config-success.md) | Optional | Configuration success metric for switches in the requested scope |
| `VersionCompliance` | [`*models.ResponseSwitchMetricsVersionCompliance`](../../doc/models/response-switch-metrics-version-compliance.md) | Optional | Version compliance metric for switches in the requested scope |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    responseSwitchMetrics := models.ResponseSwitchMetrics{
        ActivePortsSummary:   models.ToPointer(models.ResponseSwitchMetricsActivePortsSummary{
            Details:              models.ToPointer(models.SwitchMetricsActivePortsSummaryDetails{
                ActivePortCount:      models.ToPointer(136),
                TotalPortCount:       models.ToPointer(42),
            }),
            Score:                models.ToPointer(238),
            TotalSwitchCount:     models.ToPointer(26),
        }),
        ConfigSuccess:        models.ToPointer(models.ResponseSwitchMetricsConfigSuccess{
            Details:              models.ToPointer(models.ResponseSwitchMetricsConfigSuccessDetails{
                ConfigSuccessCount:   models.ToPointer(166),
            }),
            Score:                models.ToPointer(52),
            TotalSwitchCount:     models.ToPointer(160),
        }),
        VersionCompliance:    models.ToPointer(models.ResponseSwitchMetricsVersionCompliance{
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
            Score:                models.ToPointer(94),
            TotalSwitchCount:     models.ToPointer(138),
        }),
    }

}
```

