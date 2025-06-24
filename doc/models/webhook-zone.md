
# Webhook Zone

Zone webhook sample

*This model accepts additional fields of type interface{}.*

## Structure

`WebhookZone`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Events` | [`[]models.WebhookZoneEvent`](../../doc/models/webhook-zone-event.md) | Required | List of events<br><br>**Constraints**: *Minimum Items*: `1`, *Unique Items Required* |
| `Topic` | `string` | Required | Topic subscribed to<br><br>**Default**: `"zone"` |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "events": [
    {
      "map_id": "00001148-0000-0000-0000-000000000000",
      "site_id": "441a1214-6928-442a-8e92-e1d34b8ec6a6",
      "timestamp": 188.18,
      "trigger": "enter",
      "type": "sdk",
      "zone_id": "00000f60-0000-0000-0000-000000000000",
      "asset_id": "00001e56-0000-0000-0000-000000000000",
      "id": "00000ce4-0000-0000-0000-000000000000",
      "mac": "mac4",
      "name": "name0",
      "exampleAdditionalProperty": {
        "key1": "val1",
        "key2": "val2"
      }
    }
  ],
  "topic": "zone",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

