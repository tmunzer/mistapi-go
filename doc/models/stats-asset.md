
# Stats Asset

BLE asset location and advertisement statistics

## Structure

`StatsAsset`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Ttl` | `*int` | Optional | Time-to-live in seconds; how long this asset data is valid in cache |
| `BatteryPercent` | `*int` | Optional | Estimated battery level (1–100%); currently supported for Aruba/HPE asset tags |
| `BatteryVoltage` | `*float64` | Optional | Battery voltage, in mV |
| `Beam` | `*int` | Optional | BLE beam number where the asset was observed |
| `By` | `*string` | Optional | Observation source type for the asset statistic |
| `DeviceId` | `*uuid.UUID` | Optional, Read-only | Device ID of the loudest AP |
| `DeviceName` | `*string` | Optional | Display name of the loudest AP observing the asset |
| `Duration` | `*int` | Optional | Length of the current asset observation, in seconds |
| `EddystoneUidInstance` | `*string` | Optional | Eddystone-UID instance value advertised by the asset |
| `EddystoneUidNamespace` | `*string` | Optional | Eddystone-UID namespace value advertised by the asset |
| `EddystoneUrlUrl` | `*string` | Optional | URL value advertised by the asset through Eddystone-URL |
| `IbeaconMajor` | `models.Optional[int]` | Optional | Major number for iBeacon<br><br>**Constraints**: `>= 1`, `<= 65535` |
| `IbeaconMinor` | `models.Optional[int]` | Optional | Minor number for iBeacon<br><br>**Constraints**: `>= 1`, `<= 65535` |
| `IbeaconUuid` | `models.Optional[uuid.UUID]` | Optional | iBeacon UUID value, or null when no iBeacon UUID is configured |
| `Id` | `*uuid.UUID` | Optional, Read-only | Unique ID of the object instance in the Mist Organization |
| `LastSeen` | `models.Optional[float64]` | Optional, Read-only | Timestamp indicating when the entity was last seen |
| `Mac` | `string` | Required | Bluetooth MAC address for the asset |
| `Manufacture` | `*string` | Optional | Vendor name resolved from the BLE manufacturer company ID |
| `MapId` | `*uuid.UUID` | Optional | Map where the device belongs to |
| `MfgCompanyId` | `*int` | Optional | BLE manufacturer company ID from advertisement |
| `MfgData` | `*string` | Optional | Manufacturer-specific data (hex encoded) |
| `Name` | `*string` | Optional | Display label for the BLE asset |
| `Rssi` | `*int` | Optional | Signal strength (RSSI) of the loudest AP in dBm |
| `Rssizones` | [`[]models.AssetRssiZone`](../../doc/models/asset-rssi-zone.md) | Optional | Only send this for individual asset stat |
| `ServicePackets` | [`[]models.StatsAssetServicePacket`](../../doc/models/stats-asset-service-packet.md) | Optional | List of all service data advertisements (maximum length of 10)<br><br>**Constraints**: *Maximum Items*: `10` |
| `Temperature` | `*float64` | Optional | Reported temperature value from the BLE asset |
| `X` | `*float64` | Optional | Map X coordinate of the asset location, in pixels |
| `Y` | `*float64` | Optional | Map Y coordinate of the asset location, in pixels |
| `Zones` | [`[]models.AssetZone`](../../doc/models/asset-zone.md) | Optional | Only send this for individual asset stat |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    statsAsset := models.StatsAsset{
        Ttl:                   models.ToPointer(196),
        BatteryPercent:        models.ToPointer(50),
        BatteryVoltage:        models.ToPointer(float64(2970)),
        Beam:                  models.ToPointer(6),
        By:                    models.ToPointer("asset"),
        DeviceId:              models.ToPointer(uuid.MustParse("00000000-0000-0000-1000-5c5b35000001")),
        DeviceName:            models.ToPointer("a"),
        Duration:              models.ToPointer(120),
        EddystoneUidInstance:  models.ToPointer("5c5b35000001"),
        EddystoneUidNamespace: models.ToPointer("2818e3868dec25629ede"),
        EddystoneUrlUrl:       models.ToPointer("https://www.abc.com"),
        IbeaconMajor:          models.NewOptional(models.ToPointer(1234)),
        IbeaconMinor:          models.NewOptional(models.ToPointer(1234)),
        IbeaconUuid:           models.NewOptional(models.ToPointer(uuid.MustParse("f3f17139-704a-f03a-2786-0400279e37c3"))),
        Id:                    models.ToPointer(uuid.MustParse("53f10664-3ce8-4c27-b382-0ef66432349f")),
        LastSeen:              models.NewOptional(models.ToPointer(float64(1470417522))),
        Mac:                   "6fa474be7ae5",
        MapId:                 models.ToPointer(uuid.MustParse("c45be59f-854d-4ef7-b782-dcd6309c84a9")),
        MfgCompanyId:          models.ToPointer(935),
        MfgData:               models.ToPointer("648520a1020000"),
        Name:                  models.ToPointer("6fa474be7ae5"),
        Rssi:                  models.ToPointer(-60),
        Temperature:           models.ToPointer(float64(23)),
        X:                     models.ToPointer(float64(280.19918140310193)),
        Y:                     models.ToPointer(float64(420.2987721046529)),
    }

}
```

