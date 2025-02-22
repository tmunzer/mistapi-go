
# Stats Ap Port Stat

*This model accepts additional fields of type interface{}.*

## Structure

`StatsApPortStat`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `FullDuplex` | `models.Optional[bool]` | Optional | - |
| `RxBytes` | `models.Optional[float64]` | Optional | - |
| `RxErrors` | `models.Optional[float64]` | Optional | - |
| `RxPkts` | `models.Optional[float64]` | Optional | - |
| `Speed` | `models.Optional[int]` | Optional | - |
| `TxBytes` | `models.Optional[float64]` | Optional | - |
| `TxPkts` | `models.Optional[float64]` | Optional | - |
| `Up` | `models.Optional[bool]` | Optional | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "full_duplex": true,
  "rx_bytes": 2056.0,
  "rx_errors": 0.0,
  "rx_pkts": 670.0,
  "speed": 1000,
  "tx_bytes": 2056,
  "tx_pkts": 670,
  "up": true,
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

