
# Dswitches Metrics Version Compliance Details

Detail values used by the discovered-switch version compliance metric

## Structure

`DswitchesMetricsVersionComplianceDetails`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `MajorVersions` | [`[]models.DswitchesComplianceMajorVersion`](../../doc/models/dswitches-compliance-major-version.md) | Required | Version-compliance groupings by discovered switch model<br><br>**Constraints**: *Unique Items Required* |
| `TotalSwitchCount` | `int` | Required | Number of discovered switches evaluated for version compliance |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    dswitchesMetricsVersionComplianceDetails := models.DswitchesMetricsVersionComplianceDetails{
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
        TotalSwitchCount:     12,
    }

}
```

