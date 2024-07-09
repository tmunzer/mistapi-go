
# Webhook Location Sdk

Location SDK sample

## Structure

`WebhookLocationSdk`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Events` | [`[]models.WebhookLocationSdkEvent`](../../doc/models/webhook-location-sdk-event.md) | Required | list of events |
| `Topic` | `string` | Required | topic subscribed to<br>**Default**: `"location_sdk"` |

## Example (as JSON)

```json
{
  "events": [
    {
      "id": "de87bf9d-183f-e383-cc68-6ba43947d403",
      "map_id": "845a23bf-bed9-e43c-4c86-6fa474be7ae5",
      "name": "optional",
      "site_id": "4ac1dcf4-9d8b-7211-65c4-057819f0862b",
      "timestamp": 1461220784,
      "type": "sdk",
      "x": 13.5,
      "y": 3.2
    }
  ],
  "topic": "location_sdk"
}
```

