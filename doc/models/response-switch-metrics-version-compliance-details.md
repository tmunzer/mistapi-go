
# Response Switch Metrics Version Compliance Details

Detail values for the switch software version compliance metric

## Structure

`ResponseSwitchMetricsVersionComplianceDetails`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `MajorVersions` | [`[]models.SwitchMetricsComplianceMajorVersion`](../../doc/models/switch-metrics-compliance-major-version.md) | Optional | Software version compliance groupings by switch model and version |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    responseSwitchMetricsVersionComplianceDetails := models.ResponseSwitchMetricsVersionComplianceDetails{
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
    }

}
```

