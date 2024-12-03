
# Webhook Client Info

*This model accepts additional fields of type interface{}.*

## Structure

`WebhookClientInfo`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Events` | [`[]models.WebhookClientInfoEvent`](../../doc/models/webhook-client-info-event.md) | Optional | - |
| `Topic` | [`*models.WebhookClientInfoTopicEnum`](../../doc/models/webhook-client-info-topic-enum.md) | Optional | enum: `client-info` |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "events": [
    {
      "hostname": "hostname6",
      "ip": "ip4",
      "mac": "mac4",
      "org_id": "00000dbc-0000-0000-0000-000000000000",
      "site_id": "0000245a-0000-0000-0000-000000000000",
      "exampleAdditionalProperty": {
        "key1": "val1",
        "key2": "val2"
      }
    }
  ],
  "topic": "client-info",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

