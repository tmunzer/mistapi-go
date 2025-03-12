
# Ap Stat Mesh Downlink

*This model accepts additional fields of type interface{}.*

## Structure

`ApStatMeshDownlink`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Band` | `*string` | Optional | - |
| `Channel` | `*int` | Optional | - |
| `IdleTime` | `*int` | Optional | - |
| `LastSeen` | `models.Optional[float64]` | Optional | Last seen timestamp |
| `Proto` | `*string` | Optional | - |
| `Rssi` | `*int` | Optional | - |
| `RxBps` | `models.Optional[int64]` | Optional | Rate of receiving traffic, bits/seconds, last known |
| `RxBytes` | `models.Optional[int64]` | Optional | Amount of traffic received since connection |
| `RxPackets` | `models.Optional[int64]` | Optional | Amount of packets received since connection |
| `RxRate` | `models.Optional[float64]` | Optional | RX Rate, Mbps |
| `RxRetries` | `models.Optional[int]` | Optional | Amount of rx retries |
| `SiteId` | `*uuid.UUID` | Optional | - |
| `Snr` | `*int` | Optional | - |
| `TxBps` | `models.Optional[int64]` | Optional | Rate of transmitting traffic, bits/seconds, last known |
| `TxBytes` | `models.Optional[int64]` | Optional | Amount of traffic sent since connection |
| `TxPackets` | `models.Optional[int64]` | Optional | Amount of packets sent since connection |
| `TxRate` | `models.Optional[float64]` | Optional | TX Rate, Mbps |
| `TxRetries` | `models.Optional[int]` | Optional | Amount of tx retries |
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
  "rx_bps": 60003,
  "rx_bytes": 8515104416,
  "rx_packets": 57770567,
  "site_id": "441a1214-6928-442a-8e92-e1d34b8ec6a6",
  "snr": 31,
  "tx_bps": 634301,
  "tx_bytes": 211217389682,
  "tx_packets": 812204062,
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

