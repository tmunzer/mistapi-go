
# Webhook Location Centrak Event

## Structure

`WebhookLocationCentrakEvent`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Mac` | `*string` | Optional | MAC address of the device |
| `MapId` | `*string` | Optional | Map id |
| `MfgCompanyId` | `*int` | Optional | Optional, BLE manufacturing company ID |
| `MfgData` | `*string` | Optional | Optional, BLE manufacturing data in hex byte-string format (i.e. "112233AABBCC") |
| `SiteId` | `*uuid.UUID` | Optional | - |
| `Timestamp` | `*float64` | Optional | Epoch (seconds) |
| `Type` | [`*models.WebhookLocationCentrakEventTypeEnum`](../../doc/models/webhook-location-centrak-event-type-enum.md) | Optional | - |
| `WifiBeaconExtendedInfo` | [`[]models.WifiBeaconExtendedInfoItems`](../../doc/models/wifi-beacon-extended-info-items.md) | Optional | Optional, list of extended beacon info packets heard from the client, frame and sequence control included with the payload |
| `X` | `*float64` | Optional | x, in meter |
| `Y` | `*float64` | Optional | y, in meter |

## Example (as JSON)

```json
{
  "mac": "mac2",
  "map_id": "map_id6",
  "mfg_company_id": 254,
  "mfg_data": "mfg_data0",
  "site_id": "000019c2-0000-0000-0000-000000000000"
}
```

