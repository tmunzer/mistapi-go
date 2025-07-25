
# Webhook Location Client

Location Client sample

*This model accepts additional fields of type interface{}.*

## Structure

`WebhookLocationClient`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Events` | [`[]models.WebhookLocationClientEvent`](../../doc/models/webhook-location-client-event.md) | Required | List of events |
| `Topic` | `string` | Required | Topic subscribed to<br><br>**Default**: `"location_client"` |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "events": [
    {
      "mac": "5684dae9ac8b",
      "map_id": "845a23bf-bed9-e43c-4c86-6fa474be7ae5",
      "site_id": "441a1214-6928-442a-8e92-e1d34b8ec6a6",
      "type": "wifi",
      "x": 13.5,
      "y": 3.2,
      "timestamp": 188.18,
      "exampleAdditionalProperty": {
        "key1": "val1",
        "key2": "val2"
      }
    }
  ],
  "topic": "location_client",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

