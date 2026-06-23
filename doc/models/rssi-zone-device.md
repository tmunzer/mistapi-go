
# Rssi Zone Device

AP device and RSSI threshold used by an RSSI zone

## Structure

`RssiZoneDevice`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `DeviceId` | `uuid.UUID` | Required, Read-only | Read-only UUID identifier for a device |
| `Rssi` | `int` | Required | Minimum RSSI threshold for considering the device inside the zone |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    rssiZoneDevice := models.RssiZoneDevice{
        DeviceId:             uuid.MustParse("00000000-0000-0000-1000-d8695a0f9e61"),
        Rssi:                 0,
    }

}
```

