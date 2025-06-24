
# Vpn Path Traffic Shaping

*This model accepts additional fields of type interface{}.*

## Structure

`VpnPathTrafficShaping`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `ClassPercentage` | `[]int` | Optional | percentages for different class of traffic: high / medium / low / best-effort adding up to 100<br><br>**Constraints**: *Minimum Items*: `4`, *Maximum Items*: `4` |
| `Enabled` | `*bool` | Optional | - |
| `MaxTxKbps` | `models.Optional[int]` | Optional | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "class_percentage": [
    5,
    6,
    7
  ],
  "enabled": false,
  "max_tx_kbps": 14,
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

