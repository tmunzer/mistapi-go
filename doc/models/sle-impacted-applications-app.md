
# Sle Impacted Applications App

SLE impact row for an application

## Structure

`SleImpactedApplicationsApp`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `App` | `*string` | Optional | Identifier of the application represented by this impacted row |
| `Degraded` | `*int` | Optional | Portion of the SLE total that was degraded for this application |
| `Duration` | `*int` | Optional | Observation time represented by this application impact row |
| `Name` | `*string` | Optional | Display name for the application impact row |
| `Threshold` | `*int` | Optional | SLE threshold value used for this application |
| `Total` | `*int` | Optional | Overall SLE total measured for this application impact row |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    sleImpactedApplicationsApp := models.SleImpactedApplicationsApp{
        App:                  models.ToPointer("app0"),
        Degraded:             models.ToPointer(10),
        Duration:             models.ToPointer(116),
        Name:                 models.ToPointer("name2"),
        Threshold:            models.ToPointer(228),
    }

}
```

