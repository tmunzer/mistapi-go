
# Webhook Location

location webhook sample

## Structure

`WebhookLocation`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Events` | [`[]models.WebhookLocationEvent`](../../doc/models/webhook-location-event.md) | Required | list of events<br>**Constraints**: *Minimum Items*: `1`, *Unique Items Required* |
| `Topic` | `string` | Required | topic subscribed to<br>**Default**: `"location"` |

## Example (as JSON)

```json
{
  "events": [
    {
      "id": "00000ce4-0000-0000-0000-000000000000",
      "map_id": "00001148-0000-0000-0000-000000000000",
      "site_id": "441a1214-6928-442a-8e92-e1d34b8ec6a6",
      "timestamp": 130,
      "type": "type0",
      "x": 94.86,
      "y": 226.14,
      "battery_voltage": 134,
      "eddystone_uid_instance": "eddystone_uid_instance4",
      "eddystone_uid_namespace": "eddystone_uid_namespace6",
      "eddystone_url_url": "eddystone_url_url4",
      "ibeacon_major": 178
    }
  ],
  "topic": "location"
}
```

