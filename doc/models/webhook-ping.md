
# Webhook Ping

ping webhook

## Structure

`WebhookPing`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Events` | [`[]models.WebhookPingEvent`](../../doc/models/webhook-ping-event.md) | Required | **Constraints**: *Minimum Items*: `1`, *Unique Items Required* |
| `Topic` | `string` | Required | **Default**: `"ping"` |

## Example (as JSON)

```json
{
  "events": [
    {
      "id": "00000ce4-0000-0000-0000-000000000000",
      "name": "name0",
      "site_id": "43e9c864-a7e4-4310-8031-d9817d2c5a43",
      "timestamp": 188.18
    }
  ],
  "topic": "ping"
}
```

