
# Stats Unconnected Client

Location statistics for an unconnected Wi-Fi client observed by an AP

## Structure

`StatsUnconnectedClient`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `ApMac` | `string` | Required | MAC address of the AP that heard the client |
| `LastSeen` | `models.Optional[float64]` | Optional, Read-only | Timestamp indicating when the entity was last seen |
| `Mac` | `string` | Required | Unconnected client MAC address observed by an AP |
| `Manufacture` | `string` | Required | Device manufacture, through fingerprinting or OUI |
| `MapId` | `models.Optional[uuid.UUID]` | Optional | Map identifier for the unconnected client's location, if known |
| `Rssi` | `int` | Required | Client RSSI observed by the AP that heard the client (in dBm) |
| `X` | `*float64` | Optional | Horizontal map coordinate of the unconnected client location, in pixels, if known |
| `Y` | `float64` | Required | Vertical map coordinate of the unconnected client location, in pixels, if known |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    statsUnconnectedClient := models.StatsUnconnectedClient{
        ApMac:                "ap_mac6",
        LastSeen:             models.NewOptional(models.ToPointer(float64(1470417522))),
        Mac:                  "mac8",
        Manufacture:          "manufacture8",
        MapId:                models.NewOptional(models.ToPointer(uuid.MustParse("00001af4-0000-0000-0000-000000000000"))),
        Rssi:                 242,
        X:                    models.ToPointer(float64(126.1)),
        Y:                    float64(1.38),
    }

}
```

