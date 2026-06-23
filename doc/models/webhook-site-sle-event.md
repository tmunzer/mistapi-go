
# Webhook Site Sle Event

Site-level Wi-Fi SLE score snapshot

## Structure

`WebhookSiteSleEvent`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `OrgId` | `*uuid.UUID` | Optional, Read-only | Unique identifier of a Mist organization |
| `SiteId` | `*uuid.UUID` | Optional, Read-only | Unique identifier of a Mist site |
| `Sle` | [`*models.WebhookSiteSleEventSle`](../../doc/models/webhook-site-sle-event-sle.md) | Optional | Wi-Fi SLE scores reported by a site SLE webhook event |
| `Timestamp` | `*float64` | Optional, Read-only | Epoch timestamp, in seconds |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    webhookSiteSleEvent := models.WebhookSiteSleEvent{
        OrgId:                models.ToPointer(uuid.MustParse("a97c1b22-a4e9-411e-9bfd-d8695a0f9e61")),
        SiteId:               models.ToPointer(uuid.MustParse("441a1214-6928-442a-8e92-e1d34b8ec6a6")),
        Sle:                  models.ToPointer(models.WebhookSiteSleEventSle{
            ApAvailability:       models.ToPointer(float64(199.22)),
            SuccessfulConnect:    models.ToPointer(float64(14.8)),
            TimeToConnect:        models.ToPointer(float64(125.56)),
        }),
        Timestamp:            models.ToPointer(float64(25.92)),
    }

}
```

