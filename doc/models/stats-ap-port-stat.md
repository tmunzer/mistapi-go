
# Stats Ap Port Stat

*This model accepts additional fields of type interface{}.*

## Structure

`StatsApPortStat`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `FullDuplex` | `models.Optional[bool]` | Optional | - |
| `RxBytes` | `models.Optional[int]` | Optional | - |
| `RxErrors` | `models.Optional[int]` | Optional | - |
| `RxPeakBps` | `models.Optional[int]` | Optional | - |
| `RxPkts` | `models.Optional[int]` | Optional | - |
| `Speed` | `models.Optional[int]` | Optional | - |
| `TxBytes` | `models.Optional[int]` | Optional | - |
| `TxPeakBps` | `models.Optional[int]` | Optional | - |
| `TxPkts` | `models.Optional[int]` | Optional | - |
| `Up` | `models.Optional[bool]` | Optional | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "full_duplex": true,
  "rx_bytes": 2056,
  "rx_errors": 0,
  "rx_peak_bps": 22185,
  "rx_pkts": 670,
  "speed": 1000,
  "tx_bytes": 2056,
  "tx_peak_bps": 43922,
  "tx_pkts": 670,
  "up": true,
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

