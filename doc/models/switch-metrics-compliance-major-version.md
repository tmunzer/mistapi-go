
# Switch Metrics Compliance Major Version

Version compliance grouping for one switch model

## Structure

`SwitchMetricsComplianceMajorVersion`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `MajorCount` | `*int` | Optional | Number of switches in this software version grouping |
| `MajorVersion` | `*string` | Optional | Software version represented by this compliance grouping |
| `Model` | `*string` | Optional | Switch model represented by this compliance grouping |
| `SystemNames` | `[]string` | Optional | Unique string values returned or accepted by this schema<br><br>**Constraints**: *Unique Items Required* |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    switchMetricsComplianceMajorVersion := models.SwitchMetricsComplianceMajorVersion{
        MajorCount:           models.ToPointer(48),
        MajorVersion:         models.ToPointer("major_version4"),
        Model:                models.ToPointer("model0"),
        SystemNames:          []string{
            "system_names0",
        },
    }

}
```

