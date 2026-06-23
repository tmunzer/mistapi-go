
# Device Other

Third-party device discovered or managed through an external vendor integration

## Structure

`DeviceOther`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `CreatedTime` | `*float64` | Optional, Read-only | When the object has been created, in epoch |
| `DeviceMac` | `*string` | Optional | MAC address of the Mist device this third-party device is attached to |
| `Id` | `*uuid.UUID` | Optional, Read-only | Unique ID of the object instance in the Mist Organization |
| `Mac` | `*string` | Optional | Third-party device MAC address |
| `Model` | `*string` | Optional | Third-party device model reported by the vendor integration |
| `ModifiedTime` | `*float64` | Optional, Read-only | When the object has been modified for the last time, in epoch |
| `Name` | `*string` | Optional | Display name reported or configured for the third-party device |
| `OrgId` | `*uuid.UUID` | Optional, Read-only | Unique identifier of a Mist organization |
| `Serial` | `*string` | Optional | Manufacturer serial number reported for the third-party device |
| `SiteId` | `*uuid.UUID` | Optional, Read-only | Unique identifier of a Mist site |
| `State` | `*string` | Optional | Operational state reported for the third-party device |
| `Vendor` | `*string` | Optional | External vendor integration that reported the third-party device |
| `VendorApiId` | `*string` | Optional | Identifier assigned by the vendor API for this third-party device |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    deviceOther := models.DeviceOther{
        CreatedTime:          models.ToPointer(float64(150.44)),
        DeviceMac:            models.ToPointer("device_mac8"),
        Id:                   models.ToPointer(uuid.MustParse("53f10664-3ce8-4c27-b382-0ef66432349f")),
        Mac:                  models.ToPointer("mac8"),
        Model:                models.ToPointer("model2"),
        OrgId:                models.ToPointer(uuid.MustParse("a97c1b22-a4e9-411e-9bfd-d8695a0f9e61")),
        SiteId:               models.ToPointer(uuid.MustParse("441a1214-6928-442a-8e92-e1d34b8ec6a6")),
    }

}
```

