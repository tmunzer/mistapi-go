
# Webhook Zone Event

## Structure

`WebhookZoneEvent`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AssetId` | `*uuid.UUID` | Optional | uuid of named asset |
| `Id` | `uuid.UUID` | Required | Unique ID of the object instance in the Mist Organnization |
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
  "id": "53f10664-3ce8-4c27-b382-0ef66432349f",
  "map_id": "00000996-0000-0000-0000-000000000000",
  "site_id": "441a1214-6928-442a-8e92-e1d34b8ec6a6",
  "timestamp": 52,
  "trigger": "enter",
  "type": "type0",
  "zone_id": "00001712-0000-0000-0000-000000000000",
  "asset_id": "00002608-0000-0000-0000-000000000000",
  "mac": "mac4",
  "name": "name0"
}
```

