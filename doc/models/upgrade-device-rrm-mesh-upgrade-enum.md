
# Upgrade Device Rrm Mesh Upgrade Enum

For APs only and if `strategy`==`rrm`. Whether to upgrade mesh AP’s parallelly or sequentially at the end of the upgrade. enum: `parallel`, `sequential`

## Enumeration

`UpgradeDeviceRrmMeshUpgradeEnum`

## Fields

| Name |
|  --- |
| `parallel` |
| `sequential` |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    upgradeDeviceRrmMeshUpgrade := models.UpgradeDeviceRrmMeshUpgradeEnum_PARALLEL

}
```

