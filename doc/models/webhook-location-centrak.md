
# Webhook Location Centrak

Sample of the `location-centrak` webhook payload.

## Structure

`WebhookLocationCentrak`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Events` | [`[]models.WebhookLocationCentrakEvent`](../../doc/models/webhook-location-centrak-event.md) | Required | List of events<br><br>**Constraints**: *Minimum Items*: `1`, *Unique Items Required* |
| `Topic` | `string` | Required, Constant | enum: `location-centrak`<br><br>**Value**: `"location-centrak"` |

## Example (as JSON)

```json
{
  "events": [
    {
      "mac": "mac4",
      "map_id": "map_id4",
      "mfg_company_id": 234,
      "mfg_data": "mfg_data2",
      "site_id": "0000245a-0000-0000-0000-000000000000"
    }
  ],
  "topic": "location-centrak"
}
```

