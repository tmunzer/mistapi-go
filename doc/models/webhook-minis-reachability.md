
# Webhook Minis Reachability

Sample of the `minis-reachability` webhook payload.

## Structure

`WebhookMinisReachability`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Events` | [`[]models.WebhookMinisReachabilityEvent`](../../doc/models/webhook-minis-reachability-event.md) | Required | Marvis Minis reachability test events included in a webhook delivery |
| `Topic` | `string` | Required, Constant | Webhook topic name for Minis reachability test deliveries. enum: `minis-reachability`<br><br>**Value**: `"minis-reachability"` |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    webhookMinisReachability := models.WebhookMinisReachability{
        Events:               []models.WebhookMinisReachabilityEvent{
            models.WebhookMinisReachabilityEvent{
                AvgLatency:           models.ToPointer(float64(11.76)),
                DeviceMac:            models.ToPointer("7cb68d8f0440"),
                LossPercentage:       models.ToPointer(float64(105.92)),
                MaxLatency:           models.ToPointer(float64(79.28)),
                MinLatency:           models.ToPointer(float64(246.38)),
                OrgId:                models.ToPointer(uuid.MustParse("a97c1b22-a4e9-411e-9bfd-d8695a0f9e61")),
                ProbeName:            models.ToPointer("google ping"),
                ProbeTarget:          models.ToPointer("google.com"),
                ProbeType:            models.ToPointer("reachability"),
                Protocol:             models.ToPointer("icmp"),
                SiteId:               models.ToPointer(uuid.MustParse("441a1214-6928-442a-8e92-e1d34b8ec6a6")),
                TestType:             models.ToPointer("ping"),
                Vlan:                 models.ToPointer(12),
            },
        },
        Topic:                "minis-reachability",
    }

}
```

