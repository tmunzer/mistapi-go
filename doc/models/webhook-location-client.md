
# Webhook Location Client

Location Client sample

## Structure

`WebhookLocationClient`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Events` | [`[]models.WebhookLocationClientEvent`](../../doc/models/webhook-location-client-event.md) | Required | list of events |
| `Topic` | `string` | Required | topic subscribed to<br>**Default**: `"location_client"` |

## Example (as JSON)

```json
{
  "events": [
    {
      "mac": "5684dae9ac8b",
      "map_id": "845a23bf-bed9-e43c-4c86-6fa474be7ae5",
      "site_id": "4ac1dcf4-9d8b-7211-65c4-057819f0862b",
      "timestamp": 1461220784,
      "type": "wifi",
      "x": 13.5,
      "y": 3.2
    }
  ],
  "topic": "location_client"
}
```

