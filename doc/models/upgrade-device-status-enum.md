
# Upgrade Device Status Enum

status upgrade is in. enum: `cancelled`, `completed`, `created`, `downloaded`, `downloading`, `failed`, `upgrading`, `queued`

## Enumeration

`UpgradeDeviceStatusEnum`

## Fields

| Name |
|  --- |
| `cancelled` |
| `completed` |
| `created` |
| `downloaded` |
| `downloading` |
| `failed` |
| `queued` |
| `upgrading` |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    upgradeDeviceStatus := models.UpgradeDeviceStatusEnum_QUEUED

}
```

