
# Webhook Site Sle

Sample of the `site-sle` webhook payload.

## Structure

`WebhookSiteSle`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Events` | [`[]models.WebhookSiteSleEvent`](../../doc/models/webhook-site-sle-event.md) | Required | Site SLE score events included in a webhook delivery |
| `Topic` | `string` | Required, Constant | Webhook topic name for site SLE score deliveries. enum: `site-sle`<br><br>**Value**: `"site-sle"` |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    webhookSiteSle := models.WebhookSiteSle{
        Events:               []models.WebhookSiteSleEvent{
            models.WebhookSiteSleEvent{
                OrgId:                models.ToPointer(uuid.MustParse("a97c1b22-a4e9-411e-9bfd-d8695a0f9e61")),
                SiteId:               models.ToPointer(uuid.MustParse("441a1214-6928-442a-8e92-e1d34b8ec6a6")),
                Sle:                  models.ToPointer(models.WebhookSiteSleEventSle{
                    ApAvailability:       models.ToPointer(float64(199.22)),
                    SuccessfulConnect:    models.ToPointer(float64(14.8)),
                    TimeToConnect:        models.ToPointer(float64(125.56)),
                }),
                Timestamp:            models.ToPointer(float64(188.18)),
            },
        },
        Topic:                "site-sle",
    }

}
```

