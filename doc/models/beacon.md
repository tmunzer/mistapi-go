
# Beacon

Beacon configuration and placement data

*This model accepts additional fields of type interface{}.*

## Structure

`Beacon`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `CreatedTime` | `*float64` | Optional, Read-only | When the object has been created, in epoch |
| `EddystoneInstance` | `*string` | Optional | Eddystone-UID instance (6 bytes) in hexstring format |
| `EddystoneNamespace` | `*string` | Optional | Eddystone-UID namespace (10 bytes) in hexstring format |
| `EddystoneUrl` | `*string` | Optional | Eddystone-URL value broadcast by the beacon |
| `ForSite` | `*bool` | Optional, Read-only | Whether this beacon is scoped directly to a site |
| `IbeaconMajor` | `models.Optional[int]` | Optional | Major number for iBeacon<br><br>**Constraints**: `>= 1`, `<= 65535` |
| `IbeaconMinor` | `models.Optional[int]` | Optional | Minor number for iBeacon<br><br>**Constraints**: `>= 1`, `<= 65535` |
| `IbeaconUuid` | `models.Optional[uuid.UUID]` | Optional | iBeacon UUID value, or null when no iBeacon UUID is configured |
| `Id` | `*uuid.UUID` | Optional, Read-only | Unique ID of the object instance in the Mist Organization |
| `Mac` | `*string` | Optional | Optional beacon MAC address, currently used only to identify battery voltage |
| `MapId` | `*uuid.UUID` | Optional | Map where the beacon is placed |
| `ModifiedTime` | `*float64` | Optional, Read-only | When the object has been modified for the last time, in epoch |
| `Name` | `*string` | Optional | Display name of the beacon |
| `OrgId` | `*uuid.UUID` | Optional, Read-only | Unique identifier of a Mist organization |
| `Power` | `*int` | Optional | Beacon transmit power, in dBm<br><br>**Default**: `-12`<br><br>**Constraints**: `>= -12`, `<= 100` |
| `SiteId` | `*uuid.UUID` | Optional, Read-only | Unique identifier of a Mist site |
| `Type` | [`*models.BeaconTypeEnum`](../../doc/models/beacon-type-enum.md) | Optional | enum: `eddystone-uid`, `eddystone-url`, `ibeacon`<br><br>**Default**: `"eddystone-uid"` |
| `X` | `*float64` | Optional | Horizontal map position of the beacon, in pixels |
| `Y` | `*float64` | Optional | Vertical map position of the beacon, in pixels |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    beacon := models.Beacon{
        CreatedTime:          models.ToPointer(float64(53.64)),
        EddystoneInstance:    models.ToPointer("eddystone_instance6"),
        EddystoneNamespace:   models.ToPointer("eddystone_namespace0"),
        EddystoneUrl:         models.ToPointer("eddystone_url8"),
        ForSite:              models.ToPointer(false),
        IbeaconMajor:         models.NewOptional(models.ToPointer(1234)),
        IbeaconMinor:         models.NewOptional(models.ToPointer(1234)),
        IbeaconUuid:          models.NewOptional(models.ToPointer(uuid.MustParse("f3f17139-704a-f03a-2786-0400279e37c3"))),
        Id:                   models.ToPointer(uuid.MustParse("53f10664-3ce8-4c27-b382-0ef66432349f")),
        OrgId:                models.ToPointer(uuid.MustParse("a97c1b22-a4e9-411e-9bfd-d8695a0f9e61")),
        Power:                models.ToPointer(-12),
        SiteId:               models.ToPointer(uuid.MustParse("441a1214-6928-442a-8e92-e1d34b8ec6a6")),
        Type:                 models.ToPointer(models.BeaconTypeEnum_EDDYSTONEUID),
        AdditionalProperties: map[string]interface{}{
            "exampleAdditionalProperty": interface{}("[key1, val1][key2, val2]"),
        },
    }

}
```

