
# Webhook Discovered Raw Rssi Event

Raw RSSI sample reported for a passive BLE device packet

## Structure

`WebhookDiscoveredRawRssiEvent`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `ApLoc` | `[]float64` | Optional | coordinates (if any) of reporting AP (updated once in 60s per client) |
| `Beam` | `int` | Required | Antenna index, from 1-8, clock-wise starting from the LED |
| `DeviceId` | `uuid.UUID` | Required | Identifier of the AP that reported the raw RSSI sample |
| `IbeaconMajor` | `models.Optional[int]` | Optional | Major number for iBeacon<br><br>**Constraints**: `>= 1`, `<= 65535` |
| `IbeaconMinor` | `models.Optional[int]` | Optional | Minor number for iBeacon<br><br>**Constraints**: `>= 1`, `<= 65535` |
| `IbeaconUuid` | `models.Optional[uuid.UUID]` | Optional | iBeacon UUID value, or null when no iBeacon UUID is configured |
| `IsAsset` | `*bool` | Optional | Whether the advertising MAC address is recognized as an asset |
| `Mac` | `string` | Required | Client MAC address for the passive BLE device or beacon that emitted the packet |
| `MapId` | `uuid.UUID` | Required | Map associated with the reporting AP when location context is available |
| `MfgCompanyId` | `*string` | Optional | BLE manufacturing company ID |
| `MfgData` | `*string` | Optional | BLE manufacturing data in hex byte-string format (ie: "112233AABBCC") |
| `OrgId` | `uuid.UUID` | Required, Read-only | Unique identifier of a Mist organization |
| `Rssi` | `float64` | Required | Received signal strength of the packet at the reporting AP, in dBm |
| `ServicePackets` | [`[]models.ServicePacket`](../../doc/models/service-packet.md) | Optional | List of service data packets heard from the asset/ beacon |
| `SiteId` | `uuid.UUID` | Required, Read-only | Unique identifier of a Mist site |
| `Timestamp` | `*float64` | Optional, Read-only | Epoch timestamp, in seconds |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    webhookDiscoveredRawRssiEvent := models.WebhookDiscoveredRawRssiEvent{
        ApLoc:                []float64{
            float64(144.18),
            float64(144.17),
        },
        Beam:                 68,
        DeviceId:             uuid.MustParse("00001622-0000-0000-0000-000000000000"),
        IbeaconMajor:         models.NewOptional(models.ToPointer(1234)),
        IbeaconMinor:         models.NewOptional(models.ToPointer(1234)),
        IbeaconUuid:          models.NewOptional(models.ToPointer(uuid.MustParse("f3f17139-704a-f03a-2786-0400279e37c3"))),
        IsAsset:              models.ToPointer(false),
        Mac:                  "mac4",
        MapId:                uuid.MustParse("00002070-0000-0000-0000-000000000000"),
        OrgId:                uuid.MustParse("a97c1b22-a4e9-411e-9bfd-d8695a0f9e61"),
        Rssi:                 float64(175.42),
        SiteId:               uuid.MustParse("441a1214-6928-442a-8e92-e1d34b8ec6a6"),
    }

}
```

