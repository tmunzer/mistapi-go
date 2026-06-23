
# Ap Redundancy Module

AP switch redundancy counts for one VC member

## Structure

`ApRedundancyModule`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `NumAps` | `*int` | Optional | Total number of APs connected through this VC member |
| `NumApsWithSwitchRedundancy` | `*int` | Optional | Number of APs on this VC member with switch redundancy coverage |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    apRedundancyModule := models.ApRedundancyModule{
        NumAps:                     models.ToPointer(15),
        NumApsWithSwitchRedundancy: models.ToPointer(8),
    }

}
```

