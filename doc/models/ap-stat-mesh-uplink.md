
# Ap Stat Mesh Uplink

*This model accepts additional fields of type interface{}.*

## Structure

`ApStatMeshUplink`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Band` | `*string` | Optional | - |
| `Channel` | `*int` | Optional | - |
| `IdleTime` | `*int` | Optional | - |
| `LastSeen` | `*float64` | Optional | - |
| `Proto` | `*string` | Optional | - |
| `Rssi` | `*int` | Optional | - |
| `RxBps` | `*int` | Optional | - |
| `RxBytes` | `*int` | Optional | - |
| `RxPackets` | `*int` | Optional | - |
| `RxRate` | `*int` | Optional | - |
| `RxRetries` | `*int` | Optional | - |
| `SiteId` | `*uuid.UUID` | Optional | - |
| `Snr` | `*int` | Optional | - |
| `TxBps` | `*int` | Optional | - |
| `TxBytes` | `*int` | Optional | - |
| `TxPackets` | `*int` | Optional | - |
| `TxRate` | `*int` | Optional | - |
| `TxRetries` | `*int` | Optional | - |
| `UplinkApId` | `*uuid.UUID` | Optional | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "band": "5",
  "channel": 36,
  "idle_time": 3,
  "last_seen": 1470417522.0,
  "proto": "n",
  "rssi": -65,
  "rx_bps": 12,
  "rx_bytes": 217416,
  "rx_packets": 2337,
  "rx_rate": 65,
  "rx_retries": 5,
  "site_id": "441a1214-6928-442a-8e92-e1d34b8ec6a6",
  "snr": 31,
  "tx_bps": 6,
  "tx_bytes": 175132,
  "tx_packets": 1566,
  "tx_rate": 65,
  "tx_retries": 500,
  "uplink_ap_id": "00000000-0000-0000-1000-5c5b35000010",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

