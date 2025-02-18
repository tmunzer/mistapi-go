
# Webhook Occupancy Alerts Event Alert Events Items

*This model accepts additional fields of type interface{}.*

## Structure

`WebhookOccupancyAlertsEventAlertEventsItems`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `CurrentOccupancy` | `int` | Required | - |
| `MapId` | `uuid.UUID` | Required | - |
| `OccupancyLimit` | `int` | Required | - |
| `OrgId` | `uuid.UUID` | Required | - |
| `Timestamp` | `float64` | Required | - |
| `Type` | [`models.WebhookOccupancyAlertTypeEnum`](../../doc/models/webhook-occupancy-alert-type-enum.md) | Required | enum: `COMPLIANCE-OK`, `COMPLIANCE-VIOLATION` |
| `ZoneId` | `uuid.UUID` | Required | - |
| `ZoneName` | `string` | Required | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "current_occupancy": 80,
  "map_id": "00002274-0000-0000-0000-000000000000",
  "occupancy_limit": 188,
  "org_id": "a97c1b22-a4e9-411e-9bfd-d8695a0f9e61",
  "timestamp": 144.22,
  "type": "COMPLIANCE-OK",
  "zone_id": "00002544-0000-0000-0000-000000000000",
  "zone_name": "zone_name4",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

