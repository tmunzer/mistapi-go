
# Webhook Site Sle Event Sle

Wi-Fi SLE scores reported by a site SLE webhook event

## Structure

`WebhookSiteSleEventSle`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `ApAvailability` | `*float64` | Optional | Wireless AP availability SLE score for the site |
| `SuccessfulConnect` | `*float64` | Optional | Connection success SLE score for the site |
| `TimeToConnect` | `*float64` | Optional | Client connection-time SLE score for the site |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    webhookSiteSleEventSle := models.WebhookSiteSleEventSle{
        ApAvailability:       models.ToPointer(float64(0.6)),
        SuccessfulConnect:    models.ToPointer(float64(0.7)),
        TimeToConnect:        models.ToPointer(float64(0.9)),
    }

}
```

