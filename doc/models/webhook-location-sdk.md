
# Webhook Location Sdk

Sample of the `location-sdk` webhook payload.

## Structure

`WebhookLocationSdk`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Events` | [`[]models.WebhookLocationSdkEvent`](../../doc/models/webhook-location-sdk-event.md) | Required | List of events |
| `Topic` | `string` | Required, Constant | **Value**: `"location-sdk"` |

## Example (as JSON)

```json
{
  "events": [
    {
      "id": "53f10664-3ce8-4c27-b382-0ef66432349f",
      "map_id": "845a23bf-bed9-e43c-4c86-6fa474be7ae5",
      "name": "optional",
      "site_id": "441a1214-6928-442a-8e92-e1d34b8ec6a6",
      "type": "sdk",
      "x": 13.5,
      "y": 3.2,
      "timestamp": 188.18
    }
  ],
  "topic": "location-sdk"
}
```

