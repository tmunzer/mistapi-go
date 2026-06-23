
# Webhook Client Sessions

Sample of the `client-sessions` webhook payload.

## Structure

`WebhookClientSessions`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Events` | [`[]models.WebhookClientSessionsEvent`](../../doc/models/webhook-client-sessions-event.md) | Required | Client session events included in a webhook delivery<br><br>**Constraints**: *Minimum Items*: `1`, *Unique Items Required* |
| `Topic` | `string` | Required, Constant | Webhook topic name for client session deliveries. enum: `client-sessions`<br><br>**Value**: `"client-sessions"` |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    webhookClientSessions := models.WebhookClientSessions{
        Events:               []models.WebhookClientSessionsEvent{
            models.WebhookClientSessionsEvent{
                Ap:                   "ap6",
                ApName:               "ap_name8",
                Band:                 "band2",
                Bssid:                "bssid4",
                ClientFamily:         "client_family4",
                ClientManufacture:    "client_manufacture6",
                ClientModel:          "client_model4",
                ClientOs:             "client_os8",
                Connect:              64,
                ConnectFloat:         float64(130.42),
                Disconnect:           14,
                DisconnectFloat:      float64(24.2),
                Duration:             68,
                Mac:                  "mac4",
                NextAp:               "next_ap8",
                OrgId:                uuid.MustParse("a97c1b22-a4e9-411e-9bfd-d8695a0f9e61"),
                Rssi:                 float64(58.22),
                SiteId:               uuid.MustParse("441a1214-6928-442a-8e92-e1d34b8ec6a6"),
                SiteName:             "site_name2",
                Ssid:                 "ssid8",
                TerminationReason:    198,
                Timestamp:            float64(188.18),
                Version:              float64(15.76),
                WlanId:               uuid.MustParse("0000177c-0000-0000-0000-000000000000"),
            },
        },
        Topic:                "client-sessions",
    }

}
```

