
# Ap Redundancy

AP switch redundancy coverage summary

## Structure

`ApRedundancy`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Modules` | [`map[string]models.ApRedundancyModule`](../../doc/models/ap-redundancy-module.md) | Optional | Property key is the node id |
| `NumAps` | `*int` | Optional | Total number of APs considered for switch redundancy |
| `NumApsWithSwitchRedundancy` | `*int` | Optional | Number of APs that have switch redundancy coverage |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    apRedundancy := models.ApRedundancy{
        Modules:                    map[string]models.ApRedundancyModule{
            "key0": models.ApRedundancyModule{
                NumAps:                     models.ToPointer(2),
                NumApsWithSwitchRedundancy: models.ToPointer(254),
            },
            "key1": models.ApRedundancyModule{
                NumAps:                     models.ToPointer(2),
                NumApsWithSwitchRedundancy: models.ToPointer(254),
            },
        },
        NumAps:                     models.ToPointer(15),
        NumApsWithSwitchRedundancy: models.ToPointer(8),
    }

}
```

