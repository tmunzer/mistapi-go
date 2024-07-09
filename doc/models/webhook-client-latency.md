
# Webhook Client Latency

## Structure

`WebhookClientLatency`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Events` | [`[]models.WebhookClientLatencyEvent`](../../doc/models/webhook-client-latency-event.md) | Optional | - |
| `Topic` | `*string` | Optional | - |

## Example (as JSON)

```json
{
  "topic": "client-latency",
  "events": [
    {
      "avg_auth": 161.94,
      "avg_dhcp": 179.42,
      "avg_dns": 84.7,
      "max_auth": 92.74,
      "max_dhcp": 117.02
    }
  ]
}
```

