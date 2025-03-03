
# Gateway Traffic Shaping

*This model accepts additional fields of type interface{}.*

## Structure

`GatewayTrafficShaping`

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
    201
  ],
  "max_tx_kbps": 250,
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

