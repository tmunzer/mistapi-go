
# Gateway Port Config Wan Speedtest Mode Enum

Controls whether Marvis/scheduler can run speedtest on this port. enum: `auto`, `enabled`, `disabled`

## Enumeration

`GatewayPortConfigWanSpeedtestModeEnum`

## Fields

| Name |
|  --- |
| `auto` |
| `enabled` |
| `disabled` |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    gatewayPortConfigWanSpeedtestMode := models.GatewayPortConfigWanSpeedtestModeEnum_ENABLED

}
```

