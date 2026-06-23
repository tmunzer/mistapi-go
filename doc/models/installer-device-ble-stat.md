
# Installer Device Ble Stat

BLE statistics for the device

## Structure

`InstallerDeviceBleStat`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Major` | `*int` | Optional | iBeacon major value reported by the device |
| `Minors` | `[]int` | Optional | iBeacon minor values reported by the device |
| `Uuid` | `*uuid.UUID` | Optional | iBeacon UUID advertised by the device |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    installerDeviceBleStat := models.InstallerDeviceBleStat{
        Major:                models.ToPointer(12345),
        Minors:               []int{
            17,
        },
        Uuid:                 models.ToPointer(uuid.MustParse("ada72f8f-1643-e5c6-94db-f2a5636f1a64")),
    }

}
```

