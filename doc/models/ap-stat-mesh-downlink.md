
# Ap Stat Mesh Downlink

## Structure

`ApStatMeshDownlink`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Band` | `*string` | Optional | - |
| `Channel` | `*int` | Optional | - |
| `IdleTime` | `*int` | Optional | - |
| `LastSeen` | `*int` | Optional | - |
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

## Example (as JSON)

```json
{
  "band": "5",
  "channel": 36,
  "idle_time": 3,
  "last_seen": 1470417522,
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
  "tx_retries": 500
}
```

