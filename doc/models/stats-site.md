
# Stats Site

Site metadata and aggregate device/client counts returned by organization site stats endpoints

## Structure

`StatsSite`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Address` | `*string` | Optional | Configured physical address for the site |
| `AlarmtemplateId` | `models.Optional[uuid.UUID]` | Optional | Alarm template applied to the site, when one overrides the organization-level alarm template |
| `AnalyticEnabled` | `*bool` | Optional | Whether analytics features are enabled for the site |
| `AptemplateId` | `models.Optional[uuid.UUID]` | Optional | AP template applied to access points in the site, when configured |
| `CountryCode` | `string` | Required | Two-letter country code used for site configuration generation |
| `CreatedTime` | `float64` | Required, Read-only | When the object has been created, in epoch |
| `EngagementEnabled` | `*bool` | Optional | Whether engagement features are enabled for the site |
| `GatewaytemplateId` | `models.Optional[uuid.UUID]` | Optional | Gateway template applied to gateways in the site, when configured |
| `Id` | `uuid.UUID` | Required, Read-only | Unique ID of the object instance in the Mist Organization |
| `Lat` | `*float64` | Optional | Geographic latitude coordinate for the site location |
| `Latlng` | [`models.LatLng`](../../doc/models/lat-lng.md) | Required | Geographic latitude and longitude coordinate pair |
| `Lng` | `*float64` | Optional | Longitude coordinate for the site location |
| `ModifiedTime` | `float64` | Required, Read-only | When the object has been modified for the last time, in epoch |
| `MspId` | `*uuid.UUID` | Optional, Read-only | Managed service provider identifier |
| `Name` | `string` | Required | Display name of the site |
| `NetworktemplateId` | `models.Optional[uuid.UUID]` | Optional | Network template applied to the site, when configured |
| `Notes` | `*string` | Optional | Free-form notes configured for the site |
| `NumAp` | `int` | Required | Number of access points in the site |
| `NumApConnected` | `int` | Required | Number of access points currently connected in the site |
| `NumClients` | `int` | Required | Number of clients currently counted in the site |
| `NumDevices` | `int` | Required | Number of managed devices in the site |
| `NumDevicesConnected` | `int` | Required | Number of managed devices currently connected in the site |
| `NumGateway` | `int` | Required | Number of gateways in the site |
| `NumGatewayConnected` | `int` | Required | Number of gateways currently connected in the site |
| `NumSwitch` | `int` | Required | Number of switches in the site |
| `NumSwitchConnected` | `int` | Required | Number of switches currently connected in the site |
| `OrgId` | `uuid.UUID` | Required, Read-only | Unique identifier of a Mist organization |
| `RftemplateId` | `models.Optional[uuid.UUID]` | Optional | RF template applied to the site, when configured |
| `SecpolicyId` | `models.Optional[uuid.UUID]` | Optional | Security policy applied to the site, when configured |
| `SitegroupIds` | `[]uuid.UUID` | Optional | Site group identifiers associated with a site statistics record |
| `SitetemplateId` | `models.Optional[uuid.UUID]` | Optional | Site template applied to the site, when configured |
| `Timezone` | `string` | Required | IANA time zone name for the site |
| `Tzoffset` | `int` | Required | Time zone offset value derived from the site's timezone |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    statsSite := models.StatsSite{
        Address:              models.ToPointer("address8"),
        AlarmtemplateId:      models.NewOptional(models.ToPointer(uuid.MustParse("00000f8e-0000-0000-0000-000000000000"))),
        AnalyticEnabled:      models.ToPointer(false),
        AptemplateId:         models.NewOptional(models.ToPointer(uuid.MustParse("0000053e-0000-0000-0000-000000000000"))),
        CountryCode:          "country_code2",
        CreatedTime:          float64(3.42),
        EngagementEnabled:    models.ToPointer(false),
        Id:                   uuid.MustParse("53f10664-3ce8-4c27-b382-0ef66432349f"),
        Latlng:               models.LatLng{
            Lat:                  float64(37.295833),
            Lng:                  float64(-122.032946),
        },
        ModifiedTime:         float64(75.54),
        MspId:                models.ToPointer(uuid.MustParse("b9d42c2e-88ee-41f8-b798-f009ce7fe909")),
        Name:                 "name2",
        NumAp:                132,
        NumApConnected:       174,
        NumClients:           118,
        NumDevices:           94,
        NumDevicesConnected:  50,
        NumGateway:           198,
        NumGatewayConnected:  18,
        NumSwitch:            114,
        NumSwitchConnected:   238,
        OrgId:                uuid.MustParse("a97c1b22-a4e9-411e-9bfd-d8695a0f9e61"),
        Timezone:             "timezone2",
        Tzoffset:             148,
    }

}
```

