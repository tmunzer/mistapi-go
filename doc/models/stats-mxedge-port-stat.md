
# Stats Mxedge Port Stat

*This model accepts additional fields of type interface{}.*

## Structure

`StatsMxedgePortStat`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `FullDuplex` | `*bool` | Optional | - |
| `Mac` | `*string` | Optional | - |
| `RxBytes` | `models.Optional[int64]` | Optional | Amount of traffic received since connection |
| `RxErrors` | `*int` | Optional | - |
| `RxPkts` | `models.Optional[int64]` | Optional | Amount of packets received since connection |
| `Speed` | `*int` | Optional | - |
| `State` | `*string` | Optional | - |
| `TxBytes` | `models.Optional[int64]` | Optional | Amount of traffic sent since connection |
| `TxErrors` | `*int` | Optional | - |
| `TxPkts` | `models.Optional[int64]` | Optional | Amount of packets sent since connection |
| `Up` | `*bool` | Optional | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "rx_bytes": 8515104416,
  "rx_pkts": 57770567,
  "tx_bytes": 211217389682,
  "tx_pkts": 812204062,
  "full_duplex": false,
  "mac": "mac4",
  "rx_errors": 252,
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

