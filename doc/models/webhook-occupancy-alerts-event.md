
# Webhook Occupancy Alerts Event

## Structure

`WebhookOccupancyAlertsEvent`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AlertEvents` | [`[]models.WebhookOccupancyAlertsEventAlertEventsItems`](../../doc/models/webhook-occupancy-alerts-event-alert-events-items.md) | Optional | list of occupancy alerts for non-compliance zones within the site detected around the same time<br>**Constraints**: *Minimum Items*: `1`, *Unique Items Required* |
| `ForSite` | `*bool` | Optional | - |
| `SiteId` | `uuid.UUID` | Required | - |
| `SiteName` | `string` | Required | - |

## Example (as JSON)

```json
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
      "type": "COMPLIANCE-VIOLATION",
      "zone_id": "000010f6-0000-0000-0000-000000000000",
      "zone_name": "zone_name6"
    },
    {
      "current_occupancy": 30,
      "map_id": "00000fb2-0000-0000-0000-000000000000",
      "occupancy_limit": 78,
      "org_id": "00000f52-0000-0000-0000-000000000000",
      "timestamp": 148.24,
      "type": "COMPLIANCE-VIOLATION",
      "zone_id": "000010f6-0000-0000-0000-000000000000",
      "zone_name": "zone_name6"
    }
  ],
  "for_site": false
}
```

