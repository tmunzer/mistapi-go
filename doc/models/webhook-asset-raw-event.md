
# Webhook Asset Raw Event

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

## Example (as JSON)

```json
{
  "asset_id": "000004dc-0000-0000-0000-000000000000",
  "beam": 2,
  "device_id": "00000bd0-0000-0000-0000-000000000000",
  "mac": "mac2",
  "map_id": "000003b2-0000-0000-0000-000000000000",
  "mfg_company_id": 1.12,
  "mfg_data": "mfg_data0",
  "rssi": 249.0,
  "site_id": "441a1214-6928-442a-8e92-e1d34b8ec6a6",
  "timestamp": 122.96,
  "ibeacon_major": 56,
  "ibeacon_minor": 174,
  "ibeacon_uuid": "00002230-0000-0000-0000-000000000000",
  "service_data_data": "service_data_data6",
  "service_data_last_rx_time": 76
}
```

