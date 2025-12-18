
# Webhook Location

Sample of the `location` webhook payload.

## Structure

`WebhookLocation`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Events` | [`[]models.WebhookLocationEvent`](../../doc/models/webhook-location-event.md) | Required | List of events<br><br>**Constraints**: *Minimum Items*: `1`, *Unique Items Required* |
| `Topic` | `string` | Required, Constant | enum: `location`<br><br>**Value**: `"location"` |

## Example (as JSON)

```json
{
  "events": [
    {
      "ibeacon_major": 1234,
      "ibeacon_minor": 1234,
      "ibeacon_uuid": "f3f17139-704a-f03a-2786-0400279e37c3",
      "id": "53f10664-3ce8-4c27-b382-0ef66432349f",
      "map_id": "00001148-0000-0000-0000-000000000000",
      "site_id": "441a1214-6928-442a-8e92-e1d34b8ec6a6",
      "timestamp": 188.18,
      "type": "type0",
      "x": 94.86,
      "y": 226.14,
      "battery_voltage": 134,
      "eddystone_uid_instance": "eddystone_uid_instance4",
      "eddystone_uid_namespace": "eddystone_uid_namespace6",
      "eddystone_url_url": "eddystone_url_url4"
    }
  ],
  "topic": "location"
}
```

