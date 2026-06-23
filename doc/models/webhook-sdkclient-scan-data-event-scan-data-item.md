
# Webhook Sdkclient Scan Data Event Scan Data Item

Single BSSID observation from an SDK client's background Wi-Fi scan

## Structure

`WebhookSdkclientScanDataEventScanDataItem`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Ap` | `string` | Required | MAC address of the AP associated with the BSSID scanned |
| `Band` | [`models.ScanDataItemBandEnum`](../../doc/models/scan-data-item-band-enum.md) | Required | 5GHz or 2.4GHz band, associated with the BSSID scanned. enum: `2.4`, `5`<br><br>**Constraints**: *Minimum Length*: `1` |
| `Bssid` | `string` | Required | Scanned BSSID found during the SDK client's background Wi-Fi scan |
| `Channel` | `int` | Required | Radio channel found in the SDK client's background scan |
| `Rssi` | `float64` | Required | Signal strength measured by the SDK client for the scanned BSSID |
| `Ssid` | `string` | Required | Network SSID containing the BSSID scanned by the SDK client |
| `Timestamp` | `float64` | Required, Read-only | Epoch timestamp, in seconds |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    webhookSdkclientScanDataEventScanDataItem := models.WebhookSdkclientScanDataEventScanDataItem{
        Ap:                   "ap2",
        Band:                 models.ScanDataItemBandEnum_ENUM24,
        Bssid:                "bssid0",
        Channel:              162,
        Rssi:                 float64(4.48),
        Ssid:                 "ssid4",
        Timestamp:            float64(134.44),
    }

}
```

