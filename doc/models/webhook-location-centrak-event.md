
# Webhook Location Centrak Event

CenTrak location update with map coordinates and Wi-Fi beacon metadata

## Structure

`WebhookLocationCentrakEvent`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Mac` | `*string` | Optional | Device MAC address used as the CenTrak location identifier |
| `MapId` | `*string` | Optional | Identifier of the map where the CenTrak location was calculated |
| `MfgCompanyId` | `*int` | Optional | Optional, BLE manufacturing company ID |
| `MfgData` | `*string` | Optional | Optional, BLE manufacturing data in hex byte-string format (i.e. "112233AABBCC") |
| `SiteId` | `*uuid.UUID` | Optional | Site associated with the CenTrak location event |
| `Timestamp` | `*float64` | Optional, Read-only | Epoch timestamp, in seconds |
| `Type` | [`*models.WebhookLocationCentrakEventTypeEnum`](../../doc/models/webhook-location-centrak-event-type-enum.md) | Optional | Location object type for CenTrak Wi-Fi beacon events. enum: `wifi` |
| `WifiBeaconExtendedInfo` | [`[]models.WifiBeaconExtendedInfoItems`](../../doc/models/wifi-beacon-extended-info-items.md) | Optional | Optional, list of extended beacon info packets heard from the client, frame and sequence control included with the payload |
| `X` | `*float64` | Optional | Horizontal map coordinate of the CenTrak device, in meters |
| `Y` | `*float64` | Optional | Vertical map coordinate of the CenTrak device, in meters |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    webhookLocationCentrakEvent := models.WebhookLocationCentrakEvent{
        Mac:                    models.ToPointer("mac8"),
        MapId:                  models.ToPointer("map_id0"),
        MfgCompanyId:           models.ToPointer(154),
        MfgData:                models.ToPointer("mfg_data6"),
        SiteId:                 models.ToPointer(uuid.MustParse("00000406-0000-0000-0000-000000000000")),
    }

}
```

