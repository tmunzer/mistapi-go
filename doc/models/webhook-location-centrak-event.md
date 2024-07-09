
# Webhook Location Centrak Event

## Structure

`WebhookLocationCentrakEvent`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `MapId` | `*string` | Optional | map id |
| `MfgCompanyId` | `*int` | Optional | optional, BLE manufacturing company ID |
| `MfgData` | `*string` | Optional | optional, BLE manufacturing data in hex byte-string format (i.e. “112233AABBCC”) |
| `Timestamp` | `*int` | Optional | timestamp of the event, epoch |
| `WifiBeaconExtendedInfo` | [`[]models.WifiBeaconExtendedInfoItems`](../../doc/models/wifi-beacon-extended-info-items.md) | Optional | optional, list of extended beacon info packets heard from the client, frame and sequence control included with the payload |
| `X` | `*float64` | Optional | x, in meter |
| `Y` | `*float64` | Optional | y, in meter |

## Example (as JSON)

```json
{
  "map_id": "map_id4",
  "mfg_company_id": 206,
  "mfg_data": "mfg_data2",
  "timestamp": 102,
  "wifi_beacon_extended_info": [
    {
      "frame_ctrl": 244,
      "payload": "payload0",
      "seq_ctrl": 156
    },
    {
      "frame_ctrl": 244,
      "payload": "payload0",
      "seq_ctrl": 156
    }
  ]
}
```

