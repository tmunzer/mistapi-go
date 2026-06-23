
# Vbeacon

Virtual beacon configuration for SDK proximity notifications

*This model accepts additional fields of type interface{}.*

## Structure

`Vbeacon`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `CreatedTime` | `*float64` | Optional, Read-only | When the object has been created, in epoch |
| `ForSite` | `*bool` | Optional, Read-only | Whether this virtual beacon is defined at site scope |
| `Id` | `*uuid.UUID` | Optional, Read-only | Unique ID of the object instance in the Mist Organization |
| `Major` | `*int` | Optional | Bluetooth beacon major value used for proximity matching |
| `MapId` | `*uuid.UUID` | Optional | Floorplan map identifier containing the virtual beacon |
| `Message` | `*string` | Optional | Notification message displayed by the SDK application when a client is near the virtual beacon |
| `Minor` | `*int` | Optional | Bluetooth beacon minor value used for proximity matching |
| `ModifiedTime` | `*float64` | Optional, Read-only | When the object has been modified for the last time, in epoch |
| `Name` | `*string` | Optional | Display name or label for the virtual beacon |
| `OrgId` | `*uuid.UUID` | Optional, Read-only | Unique identifier of a Mist organization |
| `Power` | `*int` | Optional | Required if `power_mode`==`custom`, -30 - 100, in dBm. For default power_mode, power = 4 dBm.<br><br>**Default**: `4`<br><br>**Constraints**: `>= -30`, `<= 100` |
| `PowerMode` | [`*models.BleConfigPowerModeEnum`](../../doc/models/ble-config-power-mode-enum.md) | Optional | Transmit power mode for BLE beacons; use `custom` to set explicit power. enum: `custom`, `default`<br><br>**Default**: `"default"` |
| `SiteId` | `*uuid.UUID` | Optional, Read-only | Unique identifier of a Mist site |
| `Url` | `*string` | Optional | Destination URL shown by the SDK application when a client is near the virtual beacon |
| `Uuid` | `*uuid.UUID` | Optional | Bluetooth beacon UUID used for proximity matching |
| `WayfindingNodename` | `*string` | Optional | Name to be used in wayfinding_path or wayfinding_grid blob |
| `X` | `*float64` | Optional | Horizontal pixel coordinate of the virtual beacon on the map |
| `Y` | `*float64` | Optional | Vertical pixel coordinate of the virtual beacon on the map |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    vbeacon := models.Vbeacon{
        CreatedTime:          models.ToPointer(float64(135.22)),
        ForSite:              models.ToPointer(false),
        Id:                   models.ToPointer(uuid.MustParse("53f10664-3ce8-4c27-b382-0ef66432349f")),
        Major:                models.ToPointer(1356),
        MapId:                models.ToPointer(uuid.MustParse("63eda950-c6da-11e4-a628-60f81dd250cc")),
        Message:              models.ToPointer("Welcome to Mist"),
        Minor:                models.ToPointer(21),
        Name:                 models.ToPointer("conference room"),
        OrgId:                models.ToPointer(uuid.MustParse("a97c1b22-a4e9-411e-9bfd-d8695a0f9e61")),
        Power:                models.ToPointer(4),
        PowerMode:            models.ToPointer(models.BleConfigPowerModeEnum_CUSTOM),
        SiteId:               models.ToPointer(uuid.MustParse("441a1214-6928-442a-8e92-e1d34b8ec6a6")),
        Url:                  models.ToPointer("https://www.mist.com/any"),
        Uuid:                 models.ToPointer(uuid.MustParse("31375aeb-b8d3-1ea6-83bf-a31eb04e1c38")),
        WayfindingNodename:   models.ToPointer("node1"),
        X:                    models.ToPointer(float64(53.5)),
        Y:                    models.ToPointer(float64(173.1)),
        AdditionalProperties: map[string]interface{}{
            "exampleAdditionalProperty": interface{}("[key1, val1][key2, val2]"),
        },
    }

}
```

