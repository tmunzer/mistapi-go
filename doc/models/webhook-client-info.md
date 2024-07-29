
# Webhook Client Info

## Structure

`WebhookClientInfo`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Events` | [`[]models.WebhookClientInfoEvent`](../../doc/models/webhook-client-info-event.md) | Optional | - |
| `Topic` | [`*models.WebhookClientInfoTopicEnum`](../../doc/models/webhook-client-info-topic-enum.md) | Optional | enum: `client-info` |

## Example (as JSON)

```json
{
  "events": [
    {
      "hostname": "hostname6",
      "ip": "ip4",
      "mac": "mac4",
      "org_id": "00000dbc-0000-0000-0000-000000000000",
      "site_id": "0000245a-0000-0000-0000-000000000000"
    }
  ],
  "topic": "client-info"
}
```

