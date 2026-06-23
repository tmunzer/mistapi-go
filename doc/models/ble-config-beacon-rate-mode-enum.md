
# Ble Config Beacon Rate Mode Enum

Beacon rate mode for Mist BLE beacons; use custom to set beacon_rate. enum: `custom`, `default`

## Enumeration

`BleConfigBeaconRateModeEnum`

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
    bleConfigBeaconRateMode := models.BleConfigBeaconRateModeEnum_CUSTOM

}
```

