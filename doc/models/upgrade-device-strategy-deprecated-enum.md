
# Upgrade Device Strategy Deprecated Enum

Deprecated; use `download_strategy` and `reboot_strategy` instead. `big_bang` (upgrade all at once, no orchestration), `serial` (one at a time), `canary`, or `rrm` (AP only); default is big_bang

## Enumeration

`UpgradeDeviceStrategyDeprecatedEnum`

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
    upgradeDeviceStrategyDeprecated := models.UpgradeDeviceStrategyDeprecatedEnum_BIGBANG

}
```

