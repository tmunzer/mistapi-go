
# Webhook Location Asset Event

Asset location update with map coordinates and optional BLE beacon metadata

## Structure

`WebhookLocationAssetEvent`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `BatteryVoltage` | `*int` | Optional | Battery voltage value reported by the asset tag |
| `EddystoneUidInstance` | `*string` | Optional | Eddystone UID instance value advertised by the asset tag |
| `EddystoneUidNamespace` | `*string` | Optional | Eddystone UID namespace value advertised by the asset tag |
| `EddystoneUrlUrl` | `*string` | Optional | Eddystone URL advertised by the asset tag |
| `IbeaconMajor` | `models.Optional[int]` | Optional | Major number for iBeacon<br><br>**Constraints**: `>= 1`, `<= 65535` |
| `IbeaconMinor` | `models.Optional[int]` | Optional | Minor number for iBeacon<br><br>**Constraints**: `>= 1`, `<= 65535` |
| `IbeaconUuid` | `models.Optional[uuid.UUID]` | Optional | iBeacon UUID value, or null when no iBeacon UUID is configured |
| `Mac` | `*string` | Optional | Asset MAC address used as the location identifier |
| `MapId` | `*uuid.UUID` | Optional | Map where the asset location was calculated |
| `MfgCompanyId` | `*int` | Optional | Optional, BLE manufacturing company ID |
| `MfgData` | `*string` | Optional | Optional, BLE manufacturing data in hex byte-string format (ie: "112233AABBCC") |
| `SiteId` | `*uuid.UUID` | Optional, Read-only | Unique identifier of a Mist site |
| `Timestamp` | `*float64` | Optional, Read-only | Epoch timestamp, in seconds |
| `Type` | `*string` | Optional | Location object type for the asset event; defaults to `asset`<br><br>**Default**: `"asset"` |
| `X` | `*float64` | Optional | Horizontal map coordinate of the asset position, in meters |
| `Y` | `*float64` | Optional | Vertical map coordinate of the asset position, in meters |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    webhookLocationAssetEvent := models.WebhookLocationAssetEvent{
        BatteryVoltage:        models.ToPointer(3370),
        EddystoneUidInstance:  models.ToPointer("5c5b35000001"),
        EddystoneUidNamespace: models.ToPointer("2818e3868dec25629ede"),
        EddystoneUrlUrl:       models.ToPointer("https://www.abc.com"),
        IbeaconMajor:          models.NewOptional(models.ToPointer(1234)),
        IbeaconMinor:          models.NewOptional(models.ToPointer(1234)),
        IbeaconUuid:           models.NewOptional(models.ToPointer(uuid.MustParse("f3f17139-704a-f03a-2786-0400279e37c3"))),
        Mac:                   models.ToPointer("7fc2936fd243"),
        MapId:                 models.ToPointer(uuid.MustParse("845a23bf-bed9-e43c-4c86-6fa474be7ae5")),
        MfgCompanyId:          models.ToPointer(935),
        MfgData:               models.ToPointer("648520a1020000"),
        SiteId:                models.ToPointer(uuid.MustParse("441a1214-6928-442a-8e92-e1d34b8ec6a6")),
        Type:                  models.ToPointer("asset"),
        X:                     models.ToPointer(float64(13.5)),
        Y:                     models.ToPointer(float64(3.2)),
    }

}
```

