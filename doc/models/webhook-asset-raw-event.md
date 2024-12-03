
# Webhook Asset Raw Event

*This model accepts additional fields of type interface{}.*

## Structure

`WebhookAssetRawEvent`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AssetId` | `uuid.UUID` | Required | asset id |
| `Beam` | `int` | Required | antenna index, from 1-8, clock-wise starting from the LED |
| `DeviceId` | `uuid.UUID` | Required | device where the asset reading is from |
| `IbeaconMajor` | `*int` | Optional | iBeacon major |
| `IbeaconMinor` | `*int` | Optional | iBeacon minor |
| `IbeaconUuid` | `*uuid.UUID` | Optional | iBeacon UUID |
| `Mac` | `string` | Required | MAC of the beacon |
| `MapId` | `uuid.UUID` | Required | map id |
| `MfgCompanyId` | `float64` | Required | optional, BLE manufacturing company ID |
| `MfgData` | `string` | Required | optional, BLE manufacturing data in hex byte-string format (ie: “112233AABBCC”) |
| `Rssi` | `float64` | Required | signal strength |
| `ServiceDataData` | `*string` | Optional | optional, data from service data |
| `ServiceDataLastRxTime` | `*int` | Optional | optional, last data transmit time from service data |
| `ServiceDataRxCnt` | `*int` | Optional | optional, data transmit count from service data |
| `ServiceDataUuid` | `*uuid.UUID` | Optional | optional, UUID from service data |
| `ServicePackets` | [`[]models.ServicePacket`](../../doc/models/service-packet.md) | Optional | list of service data packets heard from the asset/ beacon |
| `SiteId` | `uuid.UUID` | Required | - |
| `Timestamp` | `float64` | Required | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "asset_id": "00001ea0-0000-0000-0000-000000000000",
  "beam": 166,
  "device_id": "00002594-0000-0000-0000-000000000000",
  "mac": "mac8",
  "map_id": "000010fe-0000-0000-0000-000000000000",
  "mfg_company_id": 123.08,
  "mfg_data": "mfg_data6",
  "rssi": 114.96,
  "site_id": "441a1214-6928-442a-8e92-e1d34b8ec6a6",
  "timestamp": 244.92,
  "ibeacon_major": 220,
  "ibeacon_minor": 82,
  "ibeacon_uuid": "0000122c-0000-0000-0000-000000000000",
  "service_data_data": "service_data_data2",
  "service_data_last_rx_time": 168,
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

