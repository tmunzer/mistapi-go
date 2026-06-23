
# Webhook Client Latency

Sample of the `client-latency` webhook payload.

## Structure

`WebhookClientLatency`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Events` | [`[]models.WebhookClientLatencyEvent`](../../doc/models/webhook-client-latency-event.md) | Optional | Client latency summary events included in a webhook delivery |
| `Topic` | [`*models.WebhookClientLatencyTopicEnum`](../../doc/models/webhook-client-latency-topic-enum.md) | Optional | Webhook topic name for client latency deliveries. enum: `client-latency` |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    webhookClientLatency := models.WebhookClientLatency{
        Events:               []models.WebhookClientLatencyEvent{
            models.WebhookClientLatencyEvent{
                AvgAuth:              models.ToPointer(float64(161.94)),
                AvgDhcp:              models.ToPointer(float64(179.42)),
                AvgDns:               models.ToPointer(float64(84.7)),
                MaxAuth:              models.ToPointer(float64(92.74)),
                MaxDhcp:              models.ToPointer(float64(117.02)),
            },
        },
        Topic:                models.ToPointer(models.WebhookClientLatencyTopicEnum_CLIENTLATENCY),
    }

}
```

