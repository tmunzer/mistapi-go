
# Webhook Client Sessions Event

Wireless client session event delivered when a client roams or disconnects

## Structure

`WebhookClientSessionsEvent`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Ap` | `string` | Required | MAC address of the AP the client roamed or disconnected from |
| `ApName` | `string` | Required | user-friendly name of the AP the client roamed or disconnected from. |
| `Band` | `string` | Required | 5GHz or 2.4GHz band |
| `Bssid` | `string` | Required | WLAN radio BSSID that the client was associated with |
| `ClientFamily` | `string` | Required | Device family E.g. "Mac", "iPhone", "Apple watch" |
| `ClientManufacture` | `string` | Required | Device manufacturer E.g. "Apple" |
| `ClientModel` | `string` | Required | Device model E.g. "8+", "XS" |
| `ClientOs` | `string` | Required | Device operating system E.g. "Mojave", "Windows 10", "Linux" |
| `Connect` | `int` | Required | Time when the client connected, in epoch seconds |
| `ConnectFloat` | `float64` | Required | Client connection timestamp with millisecond precision |
| `Disconnect` | `int` | Required | Time when the client disconnected, in epoch seconds |
| `DisconnectFloat` | `float64` | Required | Client disconnect timestamp with millisecond precision |
| `Duration` | `int` | Required | Length of the roamed or completed client session indicated by the `termination_reason` value |
| `Mac` | `string` | Required | Client MAC address for the roaming or disconnected session |
| `NextAp` | `string` | Required | the AP the client has roamed to. |
| `OrgId` | `uuid.UUID` | Required, Read-only | Unique identifier of a Mist organization |
| `Rssi` | `float64` | Required | Latest average RSSI before the user disconnects |
| `SiteId` | `uuid.UUID` | Required, Read-only | Unique identifier of a Mist site |
| `SiteName` | `string` | Required | Site name associated with the client session event |
| `Ssid` | `string` | Required | WLAN SSID for the client session |
| `TerminationReason` | `int` | Required | 1 disassociate - when the client disassociates. 2 inactive - when the client is timeout. 3 roamed - when the client is roamed between APs |
| `Timestamp` | `float64` | Required, Read-only | Epoch timestamp, in seconds |
| `Version` | `float64` | Required | schema version of this message |
| `WlanId` | `uuid.UUID` | Required | Unique identifier of the WLAN for the client session |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    webhookClientSessionsEvent := models.WebhookClientSessionsEvent{
        Ap:                   "ap8",
        ApName:               "ap_name6",
        Band:                 "band8",
        Bssid:                "bssid0",
        ClientFamily:         "client_family0",
        ClientManufacture:    "client_manufacture0",
        ClientModel:          "client_model8",
        ClientOs:             "client_os4",
        Connect:              64,
        ConnectFloat:         float64(120.18),
        Disconnect:           14,
        DisconnectFloat:      float64(13.96),
        Duration:             68,
        Mac:                  "mac0",
        NextAp:               "next_ap4",
        OrgId:                uuid.MustParse("a97c1b22-a4e9-411e-9bfd-d8695a0f9e61"),
        Rssi:                 float64(47.98),
        SiteId:               uuid.MustParse("441a1214-6928-442a-8e92-e1d34b8ec6a6"),
        SiteName:             "site_name2",
        Ssid:                 "ssid6",
        TerminationReason:    58,
        Timestamp:            float64(78.06),
        Version:              float64(5.52),
        WlanId:               uuid.MustParse("00000ecc-0000-0000-0000-000000000000"),
    }

}
```

