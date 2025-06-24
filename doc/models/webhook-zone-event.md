
# Webhook Zone Event

*This model accepts additional fields of type interface{}.*

## Structure

`WebhookZoneEvent`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AssetId` | `*uuid.UUID` | Optional | Only if `type`==`asset`. UUID of named asset |
| `Id` | `*uuid.UUID` | Optional | Only if `type`==`sdk`. UUID of the SDK Client |
| `Mac` | `*string` | Optional | MAC address of Wi-Fi client, SDK Client or Asset |
| `MapId` | `uuid.UUID` | Required | Map id |
| `Name` | `*string` | Optional | Name of the client, may be empty |
| `SiteId` | `uuid.UUID` | Required | - |
| `Timestamp` | `float64` | Required | Epoch (seconds) |
| `Trigger` | [`models.WebhookZoneEventTriggerEnum`](../../doc/models/webhook-zone-event-trigger-enum.md) | Required | enum: `enter`, `exit` |
| `Type` | [`models.WebhookZoneEventTypeEnum`](../../doc/models/webhook-zone-event-type-enum.md) | Required | Type of client. enum: `asset` (BLE Tag), `sdk`, `wifi` |
| `ZoneId` | `uuid.UUID` | Required | Zone id |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "map_id": "00000996-0000-0000-0000-000000000000",
  "site_id": "441a1214-6928-442a-8e92-e1d34b8ec6a6",
  "timestamp": 207.88,
  "trigger": "enter",
  "type": "asset",
  "zone_id": "00001712-0000-0000-0000-000000000000",
  "asset_id": "00002608-0000-0000-0000-000000000000",
  "id": "00001496-0000-0000-0000-000000000000",
  "mac": "mac4",
  "name": "name0",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

