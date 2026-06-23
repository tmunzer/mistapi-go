
# Webhook Sdkclient Scan Data

Sample of the `sdkclient-scan-data` webhook payload.

## Structure

`WebhookSdkclientScanData`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Events` | [`[]models.WebhookSdkclientScanDataEvent`](../../doc/models/webhook-sdkclient-scan-data-event.md) | Required | SDK client scan data events included in a webhook delivery<br><br>**Constraints**: *Minimum Items*: `1`, *Unique Items Required* |
| `Topic` | `string` | Required, Constant | Webhook topic name for SDK client scan data deliveries. enum: `sdkclient-scan-data`<br><br>**Value**: `"sdkclient-scan-data"` |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    webhookSdkclientScanData := models.WebhookSdkclientScanData{
        Events:               []models.WebhookSdkclientScanDataEvent{
            models.WebhookSdkclientScanDataEvent{
                ConnectionAp:         "connection_ap0",
                ConnectionBand:       "connection_band2",
                ConnectionBssid:      "connection_bssid8",
                ConnectionChannel:    46,
                ConnectionRssi:       float64(33.56),
                LastSeen:             models.NewOptional(models.ToPointer(float64(1470417522))),
                Mac:                  "mac4",
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
                },
                SiteId:               uuid.MustParse("441a1214-6928-442a-8e92-e1d34b8ec6a6"),
            },
        },
        Topic:                "sdkclient-scan-data",
    }

}
```

