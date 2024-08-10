
# Webhook Zone Event

## Structure

`WebhookZoneEvent`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AssetId` | `*uuid.UUID` | Optional | uuid of named asset |
| `Id` | `uuid.UUID` | Required | uuid of SDK-client |
| `Mac` | `*string` | Optional | mac address of wifi client or asset |
| `MapId` | `uuid.UUID` | Required | map id |
| `Name` | `*string` | Optional | name of the client, may be empty |
| `SiteId` | `uuid.UUID` | Required | - |
| `Timestamp` | `int` | Required | timestamp of the event, epoch |
| `Trigger` | [`models.WebhookZoneEventTriggerEnum`](../../doc/models/webhook-zone-event-trigger-enum.md) | Required | enum: `enter`, `exit` |
| `Type` | `string` | Required | - |
| `ZoneId` | `uuid.UUID` | Required | zone id |

## Example (as JSON)

```json
{
  "asset_id": "00002608-0000-0000-0000-000000000000",
  "id": "00001496-0000-0000-0000-000000000000",
  "mac": "mac4",
  "map_id": "00000996-0000-0000-0000-000000000000",
  "name": "name0",
  "site_id": "000004fc-0000-0000-0000-000000000000",
  "timestamp": 52,
  "trigger": "enter",
  "type": "type0",
  "zone_id": "00001712-0000-0000-0000-000000000000"
}
```

