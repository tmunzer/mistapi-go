
# Ble Config Power Mode Enum

Transmit power mode for BLE beacons; use `custom` to set explicit power. enum: `custom`, `default`

## Enumeration

`BleConfigPowerModeEnum`

## Fields

| Name |
|  --- |
| `custom` |
| `default` |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    bleConfigPowerMode := models.BleConfigPowerModeEnum_CUSTOM

}
```

