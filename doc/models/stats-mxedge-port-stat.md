
# Stats Mxedge Port Stat

*This model accepts additional fields of type interface{}.*

## Structure

`StatsMxedgePortStat`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `FullDuplex` | `*bool` | Optional | - |
| `Mac` | `*string` | Optional | - |
| `RxBytes` | `*float64` | Optional | - |
| `RxErrors` | `*int` | Optional | - |
| `RxPkts` | `*int` | Optional | - |
| `Speed` | `*int` | Optional | - |
| `State` | `*string` | Optional | - |
| `TxBytes` | `*int` | Optional | - |
| `TxErrors` | `*int` | Optional | - |
| `TxPkts` | `*int` | Optional | - |
| `Up` | `*bool` | Optional | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "full_duplex": false,
  "mac": "mac4",
  "rx_bytes": 3.38,
  "rx_errors": 252,
  "rx_pkts": 140,
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

