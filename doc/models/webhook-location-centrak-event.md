
# Webhook Location Centrak Event

*This model accepts additional fields of type interface{}.*

## Structure

`WebhookLocationCentrakEvent`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `MapId` | `*string` | Optional | Map id |
| `MfgCompanyId` | `*int` | Optional | Optional, BLE manufacturing company ID |
| `MfgData` | `*string` | Optional | Optional, BLE manufacturing data in hex byte-string format (i.e. “112233AABBCC”) |
| `Timestamp` | `*int` | Optional | Timestamp of the event, epoch |
| `WifiBeaconExtendedInfo` | [`[]models.WifiBeaconExtendedInfoItems`](../../doc/models/wifi-beacon-extended-info-items.md) | Optional | Optional, list of extended beacon info packets heard from the client, frame and sequence control included with the payload |
| `X` | `*float64` | Optional | x, in meter |
| `Y` | `*float64` | Optional | y, in meter |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "map_id": "map_id6",
  "mfg_company_id": 254,
  "mfg_data": "mfg_data0",
  "timestamp": 102,
  "wifi_beacon_extended_info": [
    {
      "frame_ctrl": 244,
      "payload": "payload0",
      "seq_ctrl": 156,
      "exampleAdditionalProperty": {
        "key1": "val1",
        "key2": "val2"
      }
    },
    {
      "frame_ctrl": 244,
      "payload": "payload0",
      "seq_ctrl": 156,
      "exampleAdditionalProperty": {
        "key1": "val1",
        "key2": "val2"
      }
    }
  ],
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

