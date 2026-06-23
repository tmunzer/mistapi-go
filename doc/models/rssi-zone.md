
# Rssi Zone

RSSI-based zone configuration for a site

*This model accepts additional fields of type interface{}.*

## Structure

`RssiZone`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `CreatedTime` | `*float64` | Optional, Read-only | When the object has been created, in epoch |
| `Devices` | [`[]models.RssiZoneDevice`](../../doc/models/rssi-zone-device.md) | Required | List of devices and the respective RSSI values to be considered in the zone<br><br>**Constraints**: *Unique Items Required* |
| `ForSite` | `*bool` | Optional, Read-only | Whether the RSSI zone is scoped to a site |
| `Id` | `*uuid.UUID` | Optional, Read-only | Unique ID of the object instance in the Mist Organization |
| `ModifiedTime` | `*float64` | Optional, Read-only | When the object has been modified for the last time, in epoch |
| `Name` | `*string` | Optional | Display name of the RSSI zone |
| `OrgId` | `*uuid.UUID` | Optional, Read-only | Unique identifier of a Mist organization |
| `SiteId` | `*uuid.UUID` | Optional, Read-only | Unique identifier of a Mist site |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    rssiZone := models.RssiZone{
        CreatedTime:          models.ToPointer(float64(116.3)),
        Devices:              []models.RssiZoneDevice{
            models.RssiZoneDevice{
                DeviceId:             uuid.MustParse("00000000-0000-0000-1000-d8695a0f9e61"),
                Rssi:                 0,
            },
        },
        ForSite:              models.ToPointer(false),
        Id:                   models.ToPointer(uuid.MustParse("53f10664-3ce8-4c27-b382-0ef66432349f")),
        ModifiedTime:         models.ToPointer(float64(218.66)),
        Name:                 models.ToPointer("zone name"),
        OrgId:                models.ToPointer(uuid.MustParse("a97c1b22-a4e9-411e-9bfd-d8695a0f9e61")),
        SiteId:               models.ToPointer(uuid.MustParse("441a1214-6928-442a-8e92-e1d34b8ec6a6")),
        AdditionalProperties: map[string]interface{}{
            "exampleAdditionalProperty": interface{}("[key1, val1][key2, val2]"),
        },
    }

}
```

