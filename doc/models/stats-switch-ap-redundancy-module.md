
# Stats Switch Ap Redundancy Module

AP switch redundancy counts for one VC or stack member

## Structure

`StatsSwitchApRedundancyModule`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `NumAps` | `*int` | Optional | Total number of APs connected through this switch member |
| `NumApsWithSwitchRedundancy` | `*int` | Optional | Number of APs on this switch member with switch redundancy coverage |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    statsSwitchApRedundancyModule := models.StatsSwitchApRedundancyModule{
        NumAps:                     models.ToPointer(15),
        NumApsWithSwitchRedundancy: models.ToPointer(8),
    }

}
```

