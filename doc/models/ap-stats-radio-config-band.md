
# Ap Stats Radio Config Band

## Structure

`ApStatsRadioConfigBand`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AllowRrmDisable` | `models.Optional[bool]` | Optional | - |
| `Bandwidth` | `models.Optional[float64]` | Optional | - |
| `Channel` | `*int` | Optional | - |
| `Disabled` | `models.Optional[bool]` | Optional | - |
| `DynamicChainingEnabled` | `models.Optional[bool]` | Optional | - |
| `Power` | `models.Optional[float64]` | Optional | - |
| `PowerMax` | `models.Optional[float64]` | Optional | - |
| `PowerMin` | `models.Optional[float64]` | Optional | - |
| `RxChain` | `models.Optional[int]` | Optional | - |
| `TxChain` | `models.Optional[int]` | Optional | - |

## Example (as JSON)

```json
{
  "bandwidth": 20.0,
  "channel": 1,
  "dynamic_chaining_enabled": false,
  "power": 10,
  "power_max": 10,
  "power_min": 10,
  "rx_chain": 4,
  "tx_chain": 4,
  "allow_rrm_disable": false,
  "disabled": false
}
```

