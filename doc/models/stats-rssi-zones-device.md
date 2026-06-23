
# Stats Rssi Zones Device

AP device and RSSI threshold used by an RSSI zone

## Structure

`StatsRssiZonesDevice`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `DeviceId` | `*uuid.UUID` | Optional | Identifier of the AP device used by this RSSI zone |
| `Rssi` | `*int` | Optional | Minimum RSSI threshold for considering a device inside the zone |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    statsRssiZonesDevice := models.StatsRssiZonesDevice{
        DeviceId:             models.ToPointer(uuid.MustParse("000014d6-0000-0000-0000-000000000000")),
        Rssi:                 models.ToPointer(10),
    }

}
```

