
# Asset of Interest

BLE beacon that matched a named asset or asset filter

## Structure

`AssetOfInterest`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `ApMac` | `*string` | Optional | Access point MAC address that heard this BLE beacon<br><br>**Constraints**: *Minimum Length*: `1` |
| `Beam` | `*float64` | Optional | BLE beam number that detected this beacon |
| `By` | `*string` | Optional | Match source for this BLE beacon, such as asset or asset filter<br><br>**Constraints**: *Minimum Length*: `1` |
| `CurrSite` | `*string` | Optional | Site currently associated with this BLE beacon<br><br>**Constraints**: *Minimum Length*: `1` |
| `DeviceName` | `*string` | Optional | AP name reported for the device that heard this BLE beacon |
| `Id` | `*uuid.UUID` | Optional, Read-only | Unique ID of the object instance in the Mist Organization |
| `LastSeen` | `models.Optional[float64]` | Optional, Read-only | Timestamp indicating when the entity was last seen |
| `Mac` | `*string` | Optional | Bluetooth MAC address of this BLE beacon<br><br>**Constraints**: *Minimum Length*: `1` |
| `Manufacture` | `*string` | Optional | BLE manufacturer name resolved for this beacon<br><br>**Constraints**: *Minimum Length*: `1` |
| `MapId` | `*string` | Optional | Identifier of the map associated with this BLE beacon observation<br><br>**Constraints**: *Minimum Length*: `1` |
| `Name` | `*string` | Optional | Display name of the matched asset or filter<br><br>**Constraints**: *Minimum Length*: `1` |
| `Rssi` | `*float64` | Optional | Signal strength of this BLE beacon, in dBm |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    assetOfInterest := models.AssetOfInterest{
        ApMac:                models.ToPointer("ap_mac8"),
        Beam:                 models.ToPointer(float64(123.2)),
        By:                   models.ToPointer("by8"),
        CurrSite:             models.ToPointer("curr_site2"),
        DeviceName:           models.ToPointer("device_name4"),
        Id:                   models.ToPointer(uuid.MustParse("53f10664-3ce8-4c27-b382-0ef66432349f")),
        LastSeen:             models.NewOptional(models.ToPointer(float64(1470417522))),
    }

}
```

