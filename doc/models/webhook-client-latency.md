
# Webhook Client Latency

Sample of the `client-latency` webhook payload.

## Structure

`WebhookClientLatency`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Events` | [`[]models.WebhookClientLatencyEvent`](../../doc/models/webhook-client-latency-event.md) | Optional | - |
| `Topic` | [`*models.WebhookClientLatencyTopicEnum`](../../doc/models/webhook-client-latency-topic-enum.md) | Optional | enum: `client-latency` |

## Example (as JSON)

```json
{
  "events": [
    {
      "avg_auth": 161.94,
      "avg_dhcp": 179.42,
      "avg_dns": 84.7,
      "max_auth": 92.74,
      "max_dhcp": 117.02
    },
    {
      "avg_auth": 161.94,
      "avg_dhcp": 179.42,
      "avg_dns": 84.7,
      "max_auth": 92.74,
      "max_dhcp": 117.02
    },
    {
      "avg_auth": 161.94,
      "avg_dhcp": 179.42,
      "avg_dns": 84.7,
      "max_auth": 92.74,
      "max_dhcp": 117.02
    }
  ],
  "topic": "client-latency"
}
```

