
# Installer Provision Device

Installer payload for provisioning or replacing a device

*This model accepts additional fields of type interface{}.*

## Structure

`InstallerProvisionDevice`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `DeviceprofileName` | `*string` | Optional | Profile name applied to the device during installer provisioning |
| `ForSite` | `*bool` | Optional, Read-only | Whether this installer provisioning request is scoped to a site |
| `Height` | `*float64` | Optional | Mounting height recorded for map placement |
| `MapId` | `*uuid.UUID` | Optional | Identifier of the map where the device should be placed |
| `Name` | `string` | Required | Device name to set during installer provisioning |
| `Orientation` | `*int` | Optional | Device orientation in degrees from 0 to 359, where 0 is up and 90 is right |
| `ReplacingMac` | `*string` | Optional | Only if replacing an existing device; MAC address of the device being replaced |
| `Role` | `*string` | Optional | Optional role for switch / gateway |
| `SiteId` | `*uuid.UUID` | Optional | Identifier of the destination site for installer provisioning |
| `SiteName` | `*string` | Optional | Destination site name for installer provisioning |
| `X` | `*float64` | Optional | Horizontal map position for the device, in pixels |
| `Y` | `*float64` | Optional | Vertical map position for the device, in pixels |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    installerProvisionDevice := models.InstallerProvisionDevice{
        DeviceprofileName:    models.ToPointer("SJ1"),
        ForSite:              models.ToPointer(false),
        Height:               models.ToPointer(float64(2.7)),
        MapId:                models.ToPointer(uuid.MustParse("845a23bf-bed9-e43c-4c86-6fa474be7ae5")),
        Name:                 "SJ1-AP1",
        Orientation:          models.ToPointer(90),
        ReplacingMac:         models.ToPointer("5c5b3500003"),
        SiteId:               models.ToPointer(uuid.MustParse("72771e6a-6f5e-4de4-a5b9-1266c4197811")),
        SiteName:             models.ToPointer("SJ1"),
        X:                    models.ToPointer(float64(150)),
        Y:                    models.ToPointer(float64(300)),
        AdditionalProperties: map[string]interface{}{
            "exampleAdditionalProperty": interface{}("[key1, val1][key2, val2]"),
        },
    }

}
```

