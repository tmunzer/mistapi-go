
# Vpn Path Selection

Only if `type`==`hub_spoke`; path selection behavior for VPN paths

## Structure

`VpnPathSelection`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Strategy` | [`*models.VpnPathSelectionStrategyEnum`](../../doc/models/vpn-path-selection-strategy-enum.md) | Optional | enum: `disabled`, `simple`, `manual`<br><br>**Default**: `"disabled"` |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    vpnPathSelection := models.VpnPathSelection{
        Strategy:             models.ToPointer(models.VpnPathSelectionStrategyEnum_DISABLED),
    }

}
```

