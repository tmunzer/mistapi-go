
# Dswitches Metrics Version Compliance

Version compliance metric for discovered switches

## Structure

`DswitchesMetricsVersionCompliance`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Details` | [`models.DswitchesMetricsVersionComplianceDetails`](../../doc/models/dswitches-metrics-version-compliance-details.md) | Required | Detail values used by the discovered-switch version compliance metric |
| `Score` | `float64` | Required | Compliance score for the discovered-switch version metric |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    dswitchesMetricsVersionCompliance := models.DswitchesMetricsVersionCompliance{
        Details:              models.DswitchesMetricsVersionComplianceDetails{
            MajorVersions:        []models.DswitchesComplianceMajorVersion{
                models.DswitchesComplianceMajorVersion{
                    MajorCount:           float64(47.78),
                    Model:                "model6",
                    SystemNames:          []string{
                        "system_names6",
                        "system_names7",
                    },
                },
            },
            TotalSwitchCount:     14,
        },
        Score:                float64(252.04),
    }

}
```

