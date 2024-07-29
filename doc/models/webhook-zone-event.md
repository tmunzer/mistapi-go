
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
  "asset_id": "00001e84-0000-0000-0000-000000000000",
  "id": "00001e2a-0000-0000-0000-000000000000",
  "mac": "mac6",
  "map_id": "00000002-0000-0000-0000-000000000000",
  "name": "name2",
  "site_id": "00000e90-0000-0000-0000-000000000000",
  "timestamp": 136,
  "trigger": "enter",
  "type": "type8",
  "zone_id": "000020a6-0000-0000-0000-000000000000"
}
```

