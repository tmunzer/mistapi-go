
# Stats Ap Port Stat

*This model accepts additional fields of type interface{}.*

## Structure

`StatsApPortStat`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `FullDuplex` | `models.Optional[bool]` | Optional | - |
| `RxBytes` | `models.Optional[int64]` | Optional | Amount of traffic received since connection |
| `RxErrors` | `models.Optional[int]` | Optional | - |
| `RxPeakBps` | `models.Optional[int]` | Optional | - |
| `RxPkts` | `models.Optional[int64]` | Optional | Amount of packets received since connection |
| `Speed` | `models.Optional[int]` | Optional | - |
| `TxBytes` | `models.Optional[int64]` | Optional | Amount of traffic sent since connection |
| `TxPeakBps` | `models.Optional[int]` | Optional | - |
| `TxPkts` | `models.Optional[int64]` | Optional | Amount of packets sent since connection |
| `Up` | `models.Optional[bool]` | Optional | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "full_duplex": true,
  "rx_bytes": 8515104416,
  "rx_errors": 0,
  "rx_peak_bps": 22185,
  "rx_pkts": 57770567,
  "speed": 1000,
  "tx_bytes": 211217389682,
  "tx_peak_bps": 43922,
  "tx_pkts": 812204062,
  "up": true,
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

