
# Webhook Client Join Event

Wireless client association event delivered when a client joins a WLAN

## Structure

`WebhookClientJoinEvent`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Ap` | `string` | Required | MAC address of the AP the client connected to |
| `ApName` | `string` | Required | user-friendly name of the AP the client connected to. |
| `Band` | `string` | Required | 5GHz or 2.4GHz band |
| `Bssid` | `string` | Required | WLAN radio BSSID that the client associated with |
| `Connect` | `int` | Required | Time when the client connected, in epoch seconds |
| `ConnectFloat` | `float64` | Required | Client connection timestamp with millisecond precision |
| `Mac` | `string` | Required | Client MAC address that joined the WLAN |
| `OrgId` | `uuid.UUID` | Required, Read-only | Unique identifier of a Mist organization |
| `Rssi` | `float64` | Required | Signal strength when the client associated, in dBm |
| `SiteId` | `uuid.UUID` | Required, Read-only | Unique identifier of a Mist site |
| `SiteName` | `string` | Required | Site name associated with the client join event |
| `Ssid` | `string` | Required | WLAN SSID that the client joined |
| `Timestamp` | `float64` | Required, Read-only | Epoch timestamp, in seconds |
| `Version` | `float64` | Required | schema version of this message |
| `WlanId` | `uuid.UUID` | Required | Unique identifier of the WLAN the client joined |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    webhookClientJoinEvent := models.WebhookClientJoinEvent{
        Ap:                   "ap2",
        ApName:               "ap_name4",
        Band:                 "band8",
        Bssid:                "bssid0",
        Connect:              218,
        ConnectFloat:         float64(111.48),
        Mac:                  "mac0",
        OrgId:                uuid.MustParse("a97c1b22-a4e9-411e-9bfd-d8695a0f9e61"),
        Rssi:                 float64(39.28),
        SiteId:               uuid.MustParse("441a1214-6928-442a-8e92-e1d34b8ec6a6"),
        SiteName:             "site_name8",
        Ssid:                 "ssid4",
        Timestamp:            float64(169.24),
        Version:              float64(252.82),
        WlanId:               uuid.MustParse("00001016-0000-0000-0000-000000000000"),
    }

}
```

