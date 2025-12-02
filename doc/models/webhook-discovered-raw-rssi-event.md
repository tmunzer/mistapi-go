
# Webhook Discovered Raw Rssi Event

## Structure

`WebhookDiscoveredRawRssiEvent`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `ApLoc` | `[]float64` | Optional | coordinates (if any) of reporting AP (updated once in 60s per client) |
| `Beam` | `int` | Required | Antenna index, from 1-8, clock-wise starting from the LED |
| `DeviceId` | `uuid.UUID` | Required | Device id of the reporting AP |
| `IbeaconMajor` | `models.Optional[int]` | Optional | **Constraints**: `>= 1`, `<= 65535` |
| `IbeaconMinor` | `models.Optional[int]` | Optional | **Constraints**: `>= 1`, `<= 65535` |
| `IbeaconUuid` | `models.Optional[uuid.UUID]` | Optional | - |
| `IsAsset` | `*bool` | Optional | - |
| `Mac` | `string` | Required | MAC of the asset/ beacon |
| `MapId` | `uuid.UUID` | Required | - |
| `MfgCompanyId` | `*string` | Optional | BLE manufacturing company ID |
| `MfgData` | `*string` | Optional | BLE manufacturing data in hex byte-string format (ie: "112233AABBCC") |
| `OrgId` | `uuid.UUID` | Required | - |
| `Rssi` | `float64` | Required | Signal strength |
| `ServicePackets` | [`[]models.ServicePacket`](../../doc/models/service-packet.md) | Optional | List of service data packets heard from the asset/ beacon |
| `SiteId` | `uuid.UUID` | Required | - |
| `Timestamp` | `*float64` | Optional | Epoch (seconds) |

## Example (as JSON)

```json
{
  "beam": 12,
  "device_id": "000000da-0000-0000-0000-000000000000",
  "ibeacon_major": 1234,
  "ibeacon_minor": 1234,
  "ibeacon_uuid": "f3f17139-704a-f03a-2786-0400279e37c3",
  "mac": "mac6",
  "map_id": "00000ea8-0000-0000-0000-000000000000",
  "org_id": "a97c1b22-a4e9-411e-9bfd-d8695a0f9e61",
  "rssi": 220.94,
  "site_id": "441a1214-6928-442a-8e92-e1d34b8ec6a6",
  "ap_loc": [
    189.7
  ],
  "is_asset": false
}
```

