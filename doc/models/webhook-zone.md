
# Webhook Zone

zone webhook sample

## Structure

`WebhookZone`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Events` | [`[]models.WebhookZoneEvent`](../../doc/models/webhook-zone-event.md) | Required | list of events<br>**Constraints**: *Minimum Items*: `1`, *Unique Items Required* |
| `Topic` | `string` | Required | topic subscribed to<br>**Default**: `"zone"` |

## Example (as JSON)

```json
{
  "events": [
    {
      "asset_id": "00001e56-0000-0000-0000-000000000000",
      "id": "00000ce4-0000-0000-0000-000000000000",
      "mac": "mac4",
      "map_id": "00001148-0000-0000-0000-000000000000",
      "name": "name0",
      "site_id": "0000245a-0000-0000-0000-000000000000",
      "timestamp": 130,
      "trigger": "enter",
      "type": "type0",
      "zone_id": "00000f60-0000-0000-0000-000000000000"
    }
  ],
  "topic": "zone"
}
```

