
# Webhook Asset Raw Rssi Event

## Structure

`WebhookAssetRawRssiEvent`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `ApLoc` | `[]float64` | Optional | optional, coordinates (if any) of reporting AP (updated once in 60s per client) |
| `Beam` | `*int` | Optional | antenna index, clock-wise starting from the LED<br><br>**Constraints**: `>= 1`, `<= 9` |
| `DeviceId` | `*uuid.UUID` | Optional | Unique ID of the object instance in the Mist Organization |
| `IbeaconMajor` | `models.Optional[int]` | Optional | Major number for iBeacon<br><br>**Constraints**: `>= 1`, `<= 65535` |
| `IbeaconMinor` | `models.Optional[int]` | Optional | Minor number for iBeacon<br><br>**Constraints**: `>= 1`, `<= 65535` |
| `IbeaconUuid` | `models.Optional[uuid.UUID]` | Optional | - |
| `IsAsset` | `*bool` | Optional | - |
| `Mac` | `*string` | Optional | MAC of the asset/ beacon |
| `MapId` | `*uuid.UUID` | Optional | Unique ID of the object instance in the Mist Organization |
| `MfgCompanyId` | `models.Optional[int]` | Optional | optional, BLE manufacturing company ID |
| `MfgData` | `models.Optional[string]` | Optional | optional, BLE manufacturing data in hex byte-string format (ie “112233AABBCC”) |
| `OrgId` | `*uuid.UUID` | Optional | - |
| `Rssi` | `*int` | Optional | - |
| `ServicePackets` | [`[]models.WebhookAssetRawRssiEventServicePacket`](../../doc/models/webhook-asset-raw-rssi-event-service-packet.md) | Optional | - |
| `SiteId` | `*uuid.UUID` | Optional | - |
| `Timestamp` | `*float64` | Optional | Epoch (seconds) |

## Example (as JSON)

```json
{
  "device_id": "53f10664-3ce8-4c27-b382-0ef66432349f",
  "ibeacon_major": 1234,
  "ibeacon_minor": 1234,
  "ibeacon_uuid": "f3f17139-704a-f03a-2786-0400279e37c3",
  "map_id": "53f10664-3ce8-4c27-b382-0ef66432349f",
  "org_id": "a97c1b22-a4e9-411e-9bfd-d8695a0f9e61",
  "site_id": "441a1214-6928-442a-8e92-e1d34b8ec6a6",
  "ap_loc": [
    134.98
  ],
  "beam": 9
}
```

