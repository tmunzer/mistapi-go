
# Webhook Asset Raw Rssi Event

Raw RSSI sample reported for an asset or beacon packet

## Structure

`WebhookAssetRawRssiEvent`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `ApLoc` | `[]float64` | Optional | optional, coordinates (if any) of reporting AP (updated once in 60s per client) |
| `Beam` | `*int` | Optional | antenna index, clock-wise starting from the LED<br><br>**Constraints**: `>= 1`, `<= 9` |
| `DeviceId` | `*uuid.UUID` | Optional, Read-only | Unique ID of the object instance in the Mist Organization |
| `IbeaconMajor` | `models.Optional[int]` | Optional | Major number for iBeacon<br><br>**Constraints**: `>= 1`, `<= 65535` |
| `IbeaconMinor` | `models.Optional[int]` | Optional | Minor number for iBeacon<br><br>**Constraints**: `>= 1`, `<= 65535` |
| `IbeaconUuid` | `models.Optional[uuid.UUID]` | Optional | iBeacon UUID value, or null when no iBeacon UUID is configured |
| `IsAsset` | `*bool` | Optional | Whether the advertising MAC address is recognized as an asset |
| `Mac` | `*string` | Optional | Client MAC address for the asset or beacon that emitted the packet |
| `MapId` | `*uuid.UUID` | Optional, Read-only | Unique ID of the object instance in the Mist Organization |
| `MfgCompanyId` | `models.Optional[int]` | Optional | optional, BLE manufacturing company ID |
| `MfgData` | `models.Optional[string]` | Optional | optional, BLE manufacturing data in hex byte-string format (ie “112233AABBCC”) |
| `OrgId` | `*uuid.UUID` | Optional, Read-only | Unique identifier of a Mist organization |
| `Rssi` | `*int` | Optional | Received signal strength of the packet at the reporting AP, in dBm |
| `ServicePackets` | [`[]models.WebhookAssetRawRssiEventServicePacket`](../../doc/models/webhook-asset-raw-rssi-event-service-packet.md) | Optional | BLE service data packets decoded from an asset raw RSSI event |
| `SiteId` | `*uuid.UUID` | Optional, Read-only | Unique identifier of a Mist site |
| `Timestamp` | `*float64` | Optional, Read-only | Epoch timestamp, in seconds |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    webhookAssetRawRssiEvent := models.WebhookAssetRawRssiEvent{
        ApLoc:                []float64{
            float64(3.22),
            float64(3.21),
            float64(3.2),
        },
        Beam:                 models.ToPointer(9),
        DeviceId:             models.ToPointer(uuid.MustParse("53f10664-3ce8-4c27-b382-0ef66432349f")),
        IbeaconMajor:         models.NewOptional(models.ToPointer(1234)),
        IbeaconMinor:         models.NewOptional(models.ToPointer(1234)),
        IbeaconUuid:          models.NewOptional(models.ToPointer(uuid.MustParse("f3f17139-704a-f03a-2786-0400279e37c3"))),
        MapId:                models.ToPointer(uuid.MustParse("53f10664-3ce8-4c27-b382-0ef66432349f")),
        OrgId:                models.ToPointer(uuid.MustParse("a97c1b22-a4e9-411e-9bfd-d8695a0f9e61")),
        SiteId:               models.ToPointer(uuid.MustParse("441a1214-6928-442a-8e92-e1d34b8ec6a6")),
    }

}
```

