
# Webhook Occupancy Alerts

occupancy alert webhook sample

*This model accepts additional fields of type interface{}.*

## Structure

`WebhookOccupancyAlerts`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Events` | [`[]models.WebhookOccupancyAlertsEvent`](../../doc/models/webhook-occupancy-alerts-event.md) | Required | **Constraints**: *Minimum Items*: `1`, *Unique Items Required* |
| `Topic` | `string` | Required | **Default**: `"occupancy-alerts"` |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

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
          "zone_name": "zone_name6",
          "exampleAdditionalProperty": {
            "key1": "val1",
            "key2": "val2"
          }
        },
        {
          "current_occupancy": 30,
          "map_id": "00000fb2-0000-0000-0000-000000000000",
          "occupancy_limit": 78,
          "org_id": "00000f52-0000-0000-0000-000000000000",
          "timestamp": 148.24,
          "type": "COMPLIANCE-OK",
          "zone_id": "000010f6-0000-0000-0000-000000000000",
          "zone_name": "zone_name6",
          "exampleAdditionalProperty": {
            "key1": "val1",
            "key2": "val2"
          }
        },
        {
          "current_occupancy": 30,
          "map_id": "00000fb2-0000-0000-0000-000000000000",
          "occupancy_limit": 78,
          "org_id": "00000f52-0000-0000-0000-000000000000",
          "timestamp": 148.24,
          "type": "COMPLIANCE-OK",
          "zone_id": "000010f6-0000-0000-0000-000000000000",
          "zone_name": "zone_name6",
          "exampleAdditionalProperty": {
            "key1": "val1",
            "key2": "val2"
          }
        }
      ],
      "for_site": false,
      "exampleAdditionalProperty": {
        "key1": "val1",
        "key2": "val2"
      }
    }
  ],
  "topic": "occupancy-alerts",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

