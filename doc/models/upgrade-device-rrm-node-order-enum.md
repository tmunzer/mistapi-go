
# Upgrade Device Rrm Node Order Enum

For APs only and if `strategy`==`rrm`. Used in rrm to determine whether to start upgrade from fringe or center AP’s. enum: `center_to_fringe`, `fringe_to_center`

## Enumeration

`UpgradeDeviceRrmNodeOrderEnum`

## Fields

| Name |
|  --- |
| `center_to_fringe` |
| `fringe_to_center` |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    upgradeDeviceRrmNodeOrder := models.UpgradeDeviceRrmNodeOrderEnum_CENTERTOFRINGE

}
```

