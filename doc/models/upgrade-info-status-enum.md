
# Upgrade Info Status Enum

enum: `error`, `inprogress`, `scheduled`, `starting`, `success`

## Enumeration

`UpgradeInfoStatusEnum`

## Fields

| Name |
|  --- |
| `error` |
| `inprogress` |
| `scheduled` |
| `starting` |
| `success` |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    upgradeInfoStatus := models.UpgradeInfoStatusEnum_STARTING

}
```

