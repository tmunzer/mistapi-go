
# Stats Asset Service Packet

Service data advertisement from a BLE asset

## Structure

`StatsAssetServicePacket`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Data` | `*string` | Optional | Service data payload (hex encoded) |
| `LastRxTime` | `*int` | Optional | Unix timestamp when this service data was last received |
| `RxCnt` | `*int` | Optional | Total number of times this service data was received |
| `Uuid` | `*string` | Optional | Service UUID |

## Example (as JSON)

```json
{
  "data": "640",
  "last_rx_time": 1645855923,
  "rx_cnt": 213065,
  "uuid": "00003e10-0000-1000-8000-00805f9b34fb"
}
```

