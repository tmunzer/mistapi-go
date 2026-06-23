
# Webhook Location Event

Generic location update with map coordinates and optional BLE or Wi-Fi metadata

## Structure

`WebhookLocationEvent`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `BatteryVoltage` | `*int` | Optional | Battery voltage value reported by the located entity |
| `EddystoneUidInstance` | `*string` | Optional | Eddystone UID instance value advertised by the located entity |
| `EddystoneUidNamespace` | `*string` | Optional | Eddystone UID namespace value advertised by the located entity |
| `EddystoneUrlUrl` | `*string` | Optional | Eddystone URL advertised by the located entity |
| `IbeaconMajor` | `models.Optional[int]` | Optional | Major number for iBeacon<br><br>**Constraints**: `>= 1`, `<= 65535` |
| `IbeaconMinor` | `models.Optional[int]` | Optional | Minor number for iBeacon<br><br>**Constraints**: `>= 1`, `<= 65535` |
| `IbeaconUuid` | `models.Optional[uuid.UUID]` | Optional | iBeacon UUID value, or null when no iBeacon UUID is configured |
| `Id` | `uuid.UUID` | Required, Read-only | Unique ID of the object instance in the Mist Organization |
| `Mac` | `*string` | Optional | Located entity MAC address for this location event |
| `MapId` | `uuid.UUID` | Required | Identifier of the map where the location was calculated |
| `MfgCompanyId` | `*int` | Optional | Optional, BLE manufacturing company ID |
| `MfgData` | `*string` | Optional | Optional, BLE manufacturing data in hex byte-string format (ie "112233AABBCC") |
| `Name` | `*string` | Optional | Client or asset display name, when available |
| `SiteId` | `uuid.UUID` | Required, Read-only | Unique identifier of a Mist site |
| `Timestamp` | `float64` | Required, Read-only | Epoch timestamp, in seconds |
| `Type` | `string` | Required | Location object type for this event |
| `WifiBeaconExtendedInfo` | [`[]models.WifiBeaconExtendedInfoItems`](../../doc/models/wifi-beacon-extended-info-items.md) | Optional | Optional, list of extended beacon info packets heard from the client, frame and sequence control included with the payload |
| `X` | `float64` | Required | Horizontal map coordinate of the located entity, in meters |
| `Y` | `float64` | Required | Vertical map coordinate of the located entity, in meters |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    webhookLocationEvent := models.WebhookLocationEvent{
        BatteryVoltage:         models.ToPointer(212),
        EddystoneUidInstance:   models.ToPointer("eddystone_uid_instance2"),
        EddystoneUidNamespace:  models.ToPointer("eddystone_uid_namespace4"),
        EddystoneUrlUrl:        models.ToPointer("eddystone_url_url4"),
        IbeaconMajor:           models.NewOptional(models.ToPointer(1234)),
        IbeaconMinor:           models.NewOptional(models.ToPointer(1234)),
        IbeaconUuid:            models.NewOptional(models.ToPointer(uuid.MustParse("f3f17139-704a-f03a-2786-0400279e37c3"))),
        Id:                     uuid.MustParse("53f10664-3ce8-4c27-b382-0ef66432349f"),
        MapId:                  uuid.MustParse("00000806-0000-0000-0000-000000000000"),
        SiteId:                 uuid.MustParse("441a1214-6928-442a-8e92-e1d34b8ec6a6"),
        Timestamp:              float64(252.96),
        Type:                   "type8",
        X:                      float64(159.64),
        Y:                      float64(34.92),
    }

}
```

