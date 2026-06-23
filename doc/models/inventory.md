
# Inventory

Organization inventory record for a claimed device

## Structure

`Inventory`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Adopted` | `*bool` | Optional | Only if `type`==`switch` or `type`==`gateway`, whether the switch/gateway is adopted |
| `ChassisMac` | `*string` | Optional | For Virtual Chassis only, MAC address of the FPC0 member |
| `ChassisSerial` | `*string` | Optional | For Virtual Chassis only, the Serial Number of the FPC0 |
| `Connected` | `*bool` | Optional | Whether the device is connected |
| `CreatedTime` | `*float64` | Optional, Read-only | When the object has been created, in epoch |
| `DeviceprofileId` | `models.Optional[string]` | Optional | Deviceprofile id if assigned, null if not assigned |
| `Hostname` | `*string` | Optional | Inventory hostname value reported by the device |
| `HwRev` | `*string` | Optional | Device hardware revision number |
| `Id` | `*uuid.UUID` | Optional, Read-only | Unique ID of the object instance in the Mist Organization |
| `Jsi` | `*bool` | Optional | Whether the inventory device is in JSI mode |
| `LastDisconnected` | `*int` | Optional | Timestamp when the device last disconnected, in epoch seconds |
| `Mac` | `*string` | Optional | Device MAC address for this inventory record |
| `Magic` | `*string` | Optional | Claim code used to add this device to inventory |
| `Model` | `*string` | Optional | Device model reported in inventory |
| `ModifiedTime` | `*float64` | Optional, Read-only | When the object has been modified for the last time, in epoch |
| `Name` | `*string` | Optional | Device name if configured |
| `OrgId` | `*uuid.UUID` | Optional, Read-only | Unique identifier of a Mist organization |
| `Serial` | `*string` | Optional | Device serial number reported in inventory |
| `SiteId` | `*uuid.UUID` | Optional, Read-only | Unique identifier of a Mist site |
| `Sku` | `*string` | Optional | Device stock keeping unit |
| `Type` | [`*models.DeviceTypeDefaultApEnum`](../../doc/models/device-type-default-ap-enum.md) | Optional | enum: `ap`, `gateway`, `switch`<br><br>**Default**: `"ap"` |
| `VcMac` | `*string` | Optional | If `type`==`switch` and the device is part of a Virtual Chassis, MAC address of the Virtual Chassis. If `type`==`gateway` and the device is part of a cluster, MAC address of the cluster |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    inventory := models.Inventory{
        Adopted:              models.ToPointer(false),
        ChassisMac:           models.ToPointer("chassis_mac6"),
        ChassisSerial:        models.ToPointer("chassis_serial6"),
        Connected:            models.ToPointer(false),
        CreatedTime:          models.ToPointer(float64(98.34)),
        Id:                   models.ToPointer(uuid.MustParse("53f10664-3ce8-4c27-b382-0ef66432349f")),
        OrgId:                models.ToPointer(uuid.MustParse("a97c1b22-a4e9-411e-9bfd-d8695a0f9e61")),
        SiteId:               models.ToPointer(uuid.MustParse("441a1214-6928-442a-8e92-e1d34b8ec6a6")),
        Type:                 models.ToPointer(models.DeviceTypeDefaultApEnum_AP),
    }

}
```

