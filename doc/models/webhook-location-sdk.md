
# Webhook Location Sdk

Location SDK sample

*This model accepts additional fields of type interface{}.*

## Structure

`WebhookLocationSdk`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Events` | [`[]models.WebhookLocationSdkEvent`](../../doc/models/webhook-location-sdk-event.md) | Required | list of events |
| `Topic` | `string` | Required | topic subscribed to<br>**Default**: `"location_sdk"` |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "events": [
    {
      "id": "53f10664-3ce8-4c27-b382-0ef66432349f",
      "map_id": "845a23bf-bed9-e43c-4c86-6fa474be7ae5",
      "name": "optional",
      "site_id": "441a1214-6928-442a-8e92-e1d34b8ec6a6",
      "timestamp": 1461220784,
      "type": "sdk",
      "x": 13.5,
      "y": 3.2,
      "exampleAdditionalProperty": {
        "key1": "val1",
        "key2": "val2"
      }
    }
  ],
  "topic": "location_sdk",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

