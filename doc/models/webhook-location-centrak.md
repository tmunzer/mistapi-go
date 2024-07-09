
# Webhook Location Centrak

location webhook sample

## Structure

`WebhookLocationCentrak`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Events` | [`[]models.WebhookLocationCentrakEvent`](../../doc/models/webhook-location-centrak-event.md) | Required | list of events<br>**Constraints**: *Minimum Items*: `1`, *Unique Items Required* |
| `Topic` | `string` | Required | topic subscribed to<br>**Default**: `"location"` |

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
          "seq_ctrl": 156
        }
      ]
    }
  ],
  "topic": "location"
}
```

