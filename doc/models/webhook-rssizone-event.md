
# Webhook Rssizone Event

## Structure

`WebhookRssizoneEvent`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Mac` | `string` | Required | Client MAC address |
| `MapId` | `uuid.UUID` | Required | - |
| `RssizoneId` | `uuid.UUID` | Required | RSSI zone name |
| `SiteId` | `uuid.UUID` | Required | - |
| `Timestamp` | `float64` | Required | Epoch (seconds) |
| `Trigger` | [`models.WebhookZoneEventTriggerEnum`](../../doc/models/webhook-zone-event-trigger-enum.md) | Required | enum: `enter`, `exit` |
| `Type` | [`models.WebhookZoneEventTypeEnum`](../../doc/models/webhook-zone-event-type-enum.md) | Required | Type of client. enum: `asset` (BLE Tag), `sdk`, `wifi` |

## Example (as JSON)

```json
{
  "mac": "mac4",
  "map_id": "0000212e-0000-0000-0000-000000000000",
  "rssizone_id": "0000103c-0000-0000-0000-000000000000",
  "site_id": "441a1214-6928-442a-8e92-e1d34b8ec6a6",
  "timestamp": 47.48,
  "trigger": "enter",
  "type": "wifi"
}
```

