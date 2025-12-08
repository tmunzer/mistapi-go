
# Webhook Location Unclient

Sample of the `location-unclient` webhook payload.

## Structure

`WebhookLocationUnclient`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Events` | [`[]models.WebhookLocationUnclientEvent`](../../doc/models/webhook-location-unclient-event.md) | Required | List of events |
| `Topic` | `string` | Required, Constant | **Value**: `"location-unclient"` |

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
      "timestamp": 188.18
    }
  ],
  "topic": "location-unclient"
}
```

