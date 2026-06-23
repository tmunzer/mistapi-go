
# Const Device Ap Vble

Virtual BLE (vBLE) capability settings for an AP model

## Structure

`ConstDeviceApVble`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `BeaconRate` | `*int` | Optional | Advertisement rate for the virtual BLE beacon |
| `Beams` | `*int` | Optional | Number of virtual BLE beams supported by the AP model |
| `Power` | `*int` | Optional | Transmit power for virtual BLE beacons, in dBm |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    constDeviceApVble := models.ConstDeviceApVble{
        BeaconRate:           models.ToPointer(4),
        Beams:                models.ToPointer(9),
        Power:                models.ToPointer(8),
    }

}
```

