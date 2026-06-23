
# Webhook Sdkclient Scan Data Event

SDK client connection state and background Wi-Fi scan observations

## Structure

`WebhookSdkclientScanDataEvent`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `ConnectionAp` | `string` | Required | MAC address of the AP the client is connected to |
| `ConnectionBand` | `string` | Required | 5GHz or 2.4GHz band, of the BSSID the client is connected to |
| `ConnectionBssid` | `string` | Required | Connected AP BSSID for the SDK client |
| `ConnectionChannel` | `int` | Required | Channel of the band the client is connected to |
| `ConnectionRssi` | `float64` | Required | RSSI of the client’s connection to the AP/BSSID |
| `LastSeen` | `models.Optional[float64]` | Optional, Read-only | Timestamp indicating when the entity was last seen |
| `Mac` | `string` | Required | SDK client MAC address |
| `ScanData` | [`[]models.WebhookSdkclientScanDataEventScanDataItem`](../../doc/models/webhook-sdkclient-scan-data-event-scan-data-item.md) | Optional | Background Wi-Fi scan observations reported by an SDK client<br><br>**Constraints**: *Minimum Items*: `1`, *Unique Items Required* |
| `SiteId` | `uuid.UUID` | Required, Read-only | Unique identifier of a Mist site |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    webhookSdkclientScanDataEvent := models.WebhookSdkclientScanDataEvent{
        ConnectionAp:         "connection_ap6",
        ConnectionBand:       "connection_band6",
        ConnectionBssid:      "connection_bssid2",
        ConnectionChannel:    192,
        ConnectionRssi:       float64(200.5),
        LastSeen:             models.NewOptional(models.ToPointer(float64(1470417522))),
        Mac:                  "mac8",
        ScanData:             []models.WebhookSdkclientScanDataEventScanDataItem{
            models.WebhookSdkclientScanDataEventScanDataItem{
                Ap:                   "ap6",
                Band:                 models.ScanDataItemBandEnum_ENUM24,
                Bssid:                "bssid2",
                Channel:              72,
                Rssi:                 float64(228.1),
                Ssid:                 "ssid4",
                Timestamp:            float64(102.06),
            },
            models.WebhookSdkclientScanDataEventScanDataItem{
                Ap:                   "ap6",
                Band:                 models.ScanDataItemBandEnum_ENUM24,
                Bssid:                "bssid2",
                Channel:              72,
                Rssi:                 float64(228.1),
                Ssid:                 "ssid4",
                Timestamp:            float64(102.06),
            },
        },
        SiteId:               uuid.MustParse("441a1214-6928-442a-8e92-e1d34b8ec6a6"),
    }

}
```

