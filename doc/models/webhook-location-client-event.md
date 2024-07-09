
# Webhook Location Client Event

## Structure

`WebhookLocationClientEvent`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Mac` | `*string` | Optional | - |
| `MapId` | `*uuid.UUID` | Optional | - |
| `SiteId` | `*uuid.UUID` | Optional | - |
| `Timestamp` | `*int` | Optional | - |
| `Type` | `*string` | Optional | **Default**: `"wifi"` |
| `WifiBeaconExtendedInfo` | [`[]models.WifiBeaconExtendedInfoItems`](../../doc/models/wifi-beacon-extended-info-items.md) | Optional | optional, list of extended beacon info packets heard from the client, frame and sequence control included with the payload |
| `X` | `*float64` | Optional | x, in meter |
| `Y` | `*float64` | Optional | y, in meter |

## Example (as JSON)

```json
{
  "mac": "5684dae9ac8b",
  "map_id": "845a23bf-bed9-e43c-4c86-6fa474be7ae5",
  "site_id": "4ac1dcf4-9d8b-7211-65c4-057819f0862b",
  "timestamp": 1461220784,
  "type": "wifi",
  "x": 13.5,
  "y": 3.2
}
```

