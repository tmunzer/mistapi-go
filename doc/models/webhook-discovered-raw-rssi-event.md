
# Webhook Discovered Raw Rssi Event

## Structure

`WebhookDiscoveredRawRssiEvent`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `ApLoc` | `[]float64` | Optional | coordinates (if any) of reporting AP (updated once in 60s per client) |
| `Beam` | `int` | Required | antenna index, from 1-8, clock-wise starting from the LED |
| `DeviceId` | `uuid.UUID` | Required | device id of the reporting AP |
| `IbeaconMajor` | `*int` | Optional | - |
| `IbeaconMinor` | `*int` | Optional | - |
| `IbeaconUuid` | `*uuid.UUID` | Optional | - |
| `IsAsset` | `*bool` | Optional | - |
| `Mac` | `string` | Required | MAC of the asset/ beacon |
| `MapId` | `uuid.UUID` | Required | - |
| `MfgCompanyId` | `*string` | Optional | BLE manufacturing company ID |
| `MfgData` | `*string` | Optional | BLE manufacturing data in hex byte-string format (ie: “112233AABBCC”) |
| `OrgId` | `uuid.UUID` | Required | - |
| `Rssi` | `float64` | Required | signal strength |
| `ServicePackets` | [`[]models.ServicePacket`](../../doc/models/service-packet.md) | Optional | list of service data packets heard from the asset/ beacon |
| `SiteId` | `uuid.UUID` | Required | - |
| `Timestamp` | `*int` | Optional | - |

## Example (as JSON)

```json
{
  "beam": 204,
  "device_id": "000007ca-0000-0000-0000-000000000000",
  "mac": "mac2",
  "map_id": "000007b8-0000-0000-0000-000000000000",
  "org_id": "a97c1b22-a4e9-411e-9bfd-d8695a0f9e61",
  "rssi": 194.7,
  "site_id": "441a1214-6928-442a-8e92-e1d34b8ec6a6",
  "ap_loc": [
    92.54
  ],
  "ibeacon_major": 2,
  "ibeacon_minor": 120,
  "ibeacon_uuid": "000008e6-0000-0000-0000-000000000000",
  "is_asset": false
}
```

