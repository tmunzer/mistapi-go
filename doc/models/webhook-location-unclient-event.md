
# Webhook Location Unclient Event

Unconnected Wi-Fi client location update with map coordinates

## Structure

`WebhookLocationUnclientEvent`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Mac` | `*string` | Optional | Unconnected client MAC address used as the location identifier |
| `MapId` | `*uuid.UUID` | Optional | Map where the unconnected client location was calculated |
| `SiteId` | `*uuid.UUID` | Optional, Read-only | Unique identifier of a Mist site |
| `Timestamp` | `*float64` | Optional, Read-only | Epoch timestamp, in seconds |
| `Type` | `*string` | Optional | Location object type for the unconnected client event; defaults to `wifi`<br><br>**Default**: `"wifi"` |
| `WifiBeaconExtendedInfo` | [`[]models.WifiBeaconExtendedInfoItems`](../../doc/models/wifi-beacon-extended-info-items.md) | Optional | Optional, list of extended beacon info packets heard from the client, frame and sequence control included with the payload |
| `X` | `*float64` | Optional | Horizontal map coordinate of the unconnected client, in meters |
| `Y` | `*float64` | Optional | Vertical map coordinate of the unconnected client, in meters |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    webhookLocationUnclientEvent := models.WebhookLocationUnclientEvent{
        Mac:                    models.ToPointer("5684dae9ac8b"),
        MapId:                  models.ToPointer(uuid.MustParse("845a23bf-bed9-e43c-4c86-6fa474be7ae5")),
        SiteId:                 models.ToPointer(uuid.MustParse("441a1214-6928-442a-8e92-e1d34b8ec6a6")),
        Timestamp:              models.ToPointer(float64(158.5)),
        Type:                   models.ToPointer("wifi"),
        X:                      models.ToPointer(float64(13.5)),
        Y:                      models.ToPointer(float64(3.2)),
    }

}
```

