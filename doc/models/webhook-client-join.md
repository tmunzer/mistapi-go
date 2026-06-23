
# Webhook Client Join

Sample of the `client-join` webhook payload.

## Structure

`WebhookClientJoin`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Events` | [`[]models.WebhookClientJoinEvent`](../../doc/models/webhook-client-join-event.md) | Required | Client join events included in a webhook delivery<br><br>**Constraints**: *Minimum Items*: `1`, *Unique Items Required* |
| `Topic` | `string` | Required, Constant | Webhook topic name for client join deliveries. enum: `client-join`<br><br>**Value**: `"client-join"` |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    webhookClientJoin := models.WebhookClientJoin{
        Events:               []models.WebhookClientJoinEvent{
            models.WebhookClientJoinEvent{
                Ap:                   "ap6",
                ApName:               "ap_name8",
                Band:                 "band2",
                Bssid:                "bssid4",
                Connect:              64,
                ConnectFloat:         float64(130.42),
                Mac:                  "mac4",
                OrgId:                uuid.MustParse("a97c1b22-a4e9-411e-9bfd-d8695a0f9e61"),
                Rssi:                 float64(58.22),
                SiteId:               uuid.MustParse("441a1214-6928-442a-8e92-e1d34b8ec6a6"),
                SiteName:             "site_name2",
                Ssid:                 "ssid8",
                Timestamp:            float64(188.18),
                Version:              float64(15.76),
                WlanId:               uuid.MustParse("0000177c-0000-0000-0000-000000000000"),
            },
        },
        Topic:                "client-join",
    }

}
```

