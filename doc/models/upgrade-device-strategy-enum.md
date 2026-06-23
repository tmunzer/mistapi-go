
# Upgrade Device Strategy Enum

enum: `big_bang` (upgrade all at once), `canary`, `rrm` (APs only), `serial` (one at a time)

## Enumeration

`UpgradeDeviceStrategyEnum`

## Fields

| Name |
|  --- |
| `big_bang` |
| `canary` |
| `rrm` |
| `serial` |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    upgradeDeviceStrategy := models.UpgradeDeviceStrategyEnum_RRM

}
```

