
# Upgrade Org Devices Download Strategy Enum

enum:

* `big_bang`: download all at once, no orchestration
* `serial`: one at a time'
* `canary`: upgrade in phases

## Enumeration

`UpgradeOrgDevicesDownloadStrategyEnum`

## Fields

| Name |
|  --- |
| `canary` |
| `big_bang` |
| `serial` |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    upgradeOrgDevicesDownloadStrategy := models.UpgradeOrgDevicesDownloadStrategyEnum_BIGBANG

}
```

