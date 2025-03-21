
# Webhook Ping

Ping webhook

*This model accepts additional fields of type interface{}.*

## Structure

`WebhookPing`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Events` | [`[]models.WebhookPingEvent`](../../doc/models/webhook-ping-event.md) | Required | **Constraints**: *Minimum Items*: `1`, *Unique Items Required* |
| `Topic` | `string` | Required | **Default**: `"ping"` |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "events": [
    {
      "id": "53f10664-3ce8-4c27-b382-0ef66432349f",
      "name": "name0",
      "site_id": "441a1214-6928-442a-8e92-e1d34b8ec6a6",
      "timestamp": 188.18,
      "exampleAdditionalProperty": {
        "key1": "val1",
        "key2": "val2"
      }
    }
  ],
  "topic": "ping",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

