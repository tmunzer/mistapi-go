
# Marvis

Marvis automation and client settings

## Structure

`Marvis`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AutoOperations` | [`*models.MarvisAutoOperations`](../../doc/models/marvis-auto-operations.md) | Optional | Marvis automatic remediation operation toggles |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    marvis := models.Marvis{
        AutoOperations:       models.ToPointer(models.MarvisAutoOperations{
            ApInsufficientCapacity:                 models.ToPointer(false),
            ApLoop:                                 models.ToPointer(false),
            ApNonCompliant:                         models.ToPointer(false),
            BouncePortForAbnormalPoeClient:         models.ToPointer(false),
            DisablePortWhenDdosProtocolViolation:   models.ToPointer(false),
        }),
    }

}
```

