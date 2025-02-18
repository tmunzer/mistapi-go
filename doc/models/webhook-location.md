
# Webhook Location

location webhook sample

*This model accepts additional fields of type interface{}.*

## Structure

`WebhookLocation`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Events` | [`[]models.WebhookLocationEvent`](../../doc/models/webhook-location-event.md) | Required | List of events<br>**Constraints**: *Minimum Items*: `1`, *Unique Items Required* |
| `Topic` | `string` | Required | Topic subscribed to<br>**Default**: `"location"` |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "events": [
    {
      "id": "53f10664-3ce8-4c27-b382-0ef66432349f",
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
      "ibeacon_major": 178,
      "exampleAdditionalProperty": {
        "key1": "val1",
        "key2": "val2"
      }
    }
  ],
  "topic": "location",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

