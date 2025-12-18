
# Webhook Occupancy Alerts

Sample of the `occupancy-alerts` webhook payload.

## Structure

`WebhookOccupancyAlerts`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Events` | [`[]models.WebhookOccupancyAlertsEvent`](../../doc/models/webhook-occupancy-alerts-event.md) | Required | **Constraints**: *Minimum Items*: `1`, *Unique Items Required* |
| `Topic` | `string` | Required, Constant | enum: `occupeancy-alerts`<br><br>**Value**: `"occupancy-alerts"` |

## Example (as JSON)

```json
{
  "events": [
    {
      "site_id": "441a1214-6928-442a-8e92-e1d34b8ec6a6",
      "site_name": "site_name2",
      "alert_events": [
        {
          "current_occupancy": 30,
          "map_id": "00000fb2-0000-0000-0000-000000000000",
          "occupancy_limit": 78,
          "org_id": "00000f52-0000-0000-0000-000000000000",
          "timestamp": 148.24,
          "type": "COMPLIANCE-OK",
          "zone_id": "000010f6-0000-0000-0000-000000000000",
          "zone_name": "zone_name6"
        },
        {
          "current_occupancy": 30,
          "map_id": "00000fb2-0000-0000-0000-000000000000",
          "occupancy_limit": 78,
          "org_id": "00000f52-0000-0000-0000-000000000000",
          "timestamp": 148.24,
          "type": "COMPLIANCE-OK",
          "zone_id": "000010f6-0000-0000-0000-000000000000",
          "zone_name": "zone_name6"
        },
        {
          "current_occupancy": 30,
          "map_id": "00000fb2-0000-0000-0000-000000000000",
          "occupancy_limit": 78,
          "org_id": "00000f52-0000-0000-0000-000000000000",
          "timestamp": 148.24,
          "type": "COMPLIANCE-OK",
          "zone_id": "000010f6-0000-0000-0000-000000000000",
          "zone_name": "zone_name6"
        }
      ],
      "for_site": false
    }
  ],
  "topic": "occupancy-alerts"
}
```

