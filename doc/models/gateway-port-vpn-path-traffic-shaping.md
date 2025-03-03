
# Gateway Port Vpn Path Traffic Shaping

Only if the VPN `type`==`hub_spoke`

*This model accepts additional fields of type interface{}.*

## Structure

`GatewayPortVpnPathTrafficShaping`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `ClassPercentages` | `[]int` | Optional | percentages for different class of traffic: high / medium / low / best-effort. Sum must be equal to 100 |
| `Enabled` | `*bool` | Optional | **Default**: `false` |
| `MaxTxKbps` | `*int` | Optional | Interface Transmit Cap in kbps |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "enabled": false,
  "class_percentages": [
    163,
    164
  ],
  "max_tx_kbps": 142,
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

