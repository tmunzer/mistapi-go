
# Webhook Location Centrak

location webhook sample

*This model accepts additional fields of type interface{}.*

## Structure

`WebhookLocationCentrak`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Events` | [`[]models.WebhookLocationCentrakEvent`](../../doc/models/webhook-location-centrak-event.md) | Required | List of events<br>**Constraints**: *Minimum Items*: `1`, *Unique Items Required* |
| `Topic` | `string` | Required | Topic subscribed to<br>**Default**: `"location"` |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "events": [
    {
      "map_id": "map_id4",
      "mfg_company_id": 234,
      "mfg_data": "mfg_data2",
      "timestamp": 130,
      "wifi_beacon_extended_info": [
        {
          "frame_ctrl": 244,
          "payload": "payload0",
          "seq_ctrl": 156,
          "exampleAdditionalProperty": {
            "key1": "val1",
            "key2": "val2"
          }
        }
      ],
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

