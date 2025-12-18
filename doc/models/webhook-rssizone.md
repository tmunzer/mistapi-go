
# Webhook Rssizone

Sample of the `rssizone` webhook payload.

## Structure

`WebhookRssizone`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Events` | [`[]models.WebhookRssizoneEvent`](../../doc/models/webhook-rssizone-event.md) | Required | **Constraints**: *Minimum Items*: `1`, *Unique Items Required* |
| `Topic` | `string` | Required, Constant | enum: `rssizone`<br><br>**Value**: `"rssizone"` |

## Example (as JSON)

```json
{
  "events": [
    {
      "mac": "mac4",
      "map_id": "00001148-0000-0000-0000-000000000000",
      "rssizone_id": "00000056-0000-0000-0000-000000000000",
      "site_id": "441a1214-6928-442a-8e92-e1d34b8ec6a6",
      "timestamp": 188.18,
      "trigger": "enter",
      "type": "sdk"
    }
  ],
  "topic": "rssizone"
}
```

