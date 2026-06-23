
# Installer Device

Recently claimed device visible to installer workflows

## Structure

`InstallerDevice`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `BleStat` | [`*models.InstallerDeviceBleStat`](../../doc/models/installer-device-ble-stat.md) | Optional | BLE statistics for the device |
| `Connected` | `*bool` | Optional | Whether the device is currently connected to Mist |
| `DeviceprofileName` | `*string` | Optional | Device profile name associated with this installer device |
| `ExtIp` | `*string` | Optional | External IP address observed for device management traffic |
| `Height` | `*float64` | Optional | Mounting height recorded for map placement |
| `Ip` | `*string` | Optional | Management IP address currently reported for the device |
| `LastSeen` | `models.Optional[float64]` | Optional, Read-only | Timestamp indicating when the entity was last seen |
| `Mac` | `*string` | Optional | Device MAC address shown to installers |
| `MapId` | `*uuid.UUID` | Optional | Map where the installer placed this device |
| `Model` | `*string` | Optional | Device model reported for this installer device |
| `Name` | `*string` | Optional | Device name configured through the installer workflow |
| `Orientation` | `*int` | Optional | Device orientation in degrees from 0 to 359, where 0 is up and 90 is right |
| `Serial` | `*string` | Optional | Device serial number reported to installer workflows |
| `SiteName` | `*string` | Optional | Site name associated with this installer device |
| `Uptime` | `*int` | Optional | Device uptime, in seconds |
| `VcMac` | `models.Optional[string]` | Optional | Virtual Chassis MAC address when this device is part of a VC |
| `Version` | `*string` | Optional | Software version currently running on the device |
| `X` | `*float64` | Optional | Horizontal map position of the device, in pixels |
| `Y` | `*float64` | Optional | Vertical map position of the device, in pixels |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    installerDevice := models.InstallerDevice{
        BleStat:              models.ToPointer(models.InstallerDeviceBleStat{
            Major:                models.ToPointer(210),
            Minors:               []int{
                237,
                238,
            },
            Uuid:                 models.ToPointer(uuid.MustParse("00001d0e-0000-0000-0000-000000000000")),
        }),
        Connected:            models.ToPointer(true),
        DeviceprofileName:    models.ToPointer("SJ1"),
        ExtIp:                models.ToPointer("12.34.56.78"),
        Height:               models.ToPointer(float64(2.7)),
        Ip:                   models.ToPointer("192.168.1.111"),
        LastSeen:             models.NewOptional(models.ToPointer(float64(1470417522))),
        Mac:                  models.ToPointer("5c5b35000018"),
        MapId:                models.ToPointer(uuid.MustParse("845a23bf-bed9-e43c-4c86-6fa474be7ae5")),
        Model:                models.ToPointer("AP41"),
        Name:                 models.ToPointer("hallway"),
        Orientation:          models.ToPointer(90),
        Serial:               models.ToPointer("FXLH2015150025"),
        SiteName:             models.ToPointer("SJ1"),
        Uptime:               models.ToPointer(12345),
        Version:              models.ToPointer("0.10.24362"),
        X:                    models.ToPointer(float64(150)),
        Y:                    models.ToPointer(float64(300)),
    }

}
```

