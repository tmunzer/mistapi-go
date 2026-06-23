
# Stats Switch Ap Redundancy

AP switch redundancy coverage summary for a switch

## Structure

`StatsSwitchApRedundancy`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Modules` | [`map[string]models.StatsSwitchApRedundancyModule`](../../doc/models/stats-switch-ap-redundancy-module.md) | Optional | Per-module AP redundancy counts for VC or stacked switches |
| `NumAps` | `*int` | Optional | Total number of APs considered for switch redundancy |
| `NumApsWithSwitchRedundancy` | `*int` | Optional | Number of APs that have switch redundancy coverage |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    statsSwitchApRedundancy := models.StatsSwitchApRedundancy{
        Modules:                    map[string]models.StatsSwitchApRedundancyModule{
            "key0": models.StatsSwitchApRedundancyModule{
                NumAps:                     models.ToPointer(2),
                NumApsWithSwitchRedundancy: models.ToPointer(254),
            },
        },
        NumAps:                     models.ToPointer(15),
        NumApsWithSwitchRedundancy: models.ToPointer(8),
    }

}
```

