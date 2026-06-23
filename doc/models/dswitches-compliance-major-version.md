
# Dswitches Compliance Major Version

Version-compliance grouping for one discovered switch model

## Structure

`DswitchesComplianceMajorVersion`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `MajorCount` | `float64` | Required | Number of major software versions observed for this switch model |
| `Model` | `string` | Required | Switch model represented by this version-compliance grouping |
| `SystemNames` | `[]string` | Optional | Unique string values returned or accepted by this schema<br><br>**Constraints**: *Unique Items Required* |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    dswitchesComplianceMajorVersion := models.DswitchesComplianceMajorVersion{
        MajorCount:           float64(169.32),
        Model:                "model6",
        SystemNames:          []string{
            "system_names6",
            "system_names7",
        },
    }

}
```

