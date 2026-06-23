
# Stats Beacon

BLE beacon runtime statistics and placement data

## Structure

`StatsBeacon`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `BatteryVoltage` | `*float64` | Optional | Battery voltage, in mV |
| `EddystoneInstance` | `*string` | Optional | Eddystone-UID instance value advertised by the beacon |
| `EddystoneNamespace` | `*string` | Optional | Eddystone-UID namespace value advertised by the beacon |
| `LastSeen` | `models.Optional[float64]` | Optional, Read-only | Timestamp indicating when the entity was last seen |
| `Mac` | `string` | Required | Beacon MAC address observed by Mist |
| `MapId` | `uuid.UUID` | Required | Map identifier where the beacon is placed |
| `Name` | `string` | Required | Display name of the beacon |
| `Power` | `int` | Required | Transmit power configured for the beacon |
| `Type` | `string` | Required | Advertisement protocol type used by the beacon, such as iBeacon or Eddystone |
| `X` | `float64` | Required | Map X coordinate of the beacon placement, in pixels |
| `Y` | `float64` | Required | Map Y coordinate of the beacon placement, in pixels |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    statsBeacon := models.StatsBeacon{
        BatteryVoltage:       models.ToPointer(float64(83.32)),
        EddystoneInstance:    models.ToPointer("eddystone_instance2"),
        EddystoneNamespace:   models.ToPointer("eddystone_namespace8"),
        LastSeen:             models.NewOptional(models.ToPointer(float64(1470417522))),
        Mac:                  "mac0",
        MapId:                uuid.MustParse("00002402-0000-0000-0000-000000000000"),
        Name:                 "name6",
        Power:                98,
        Type:                 "type4",
        X:                    float64(97.08),
        Y:                    float64(34.2),
    }

}
```

