
# Dswitches Metrics Switch Ap Affinity Details

*This model accepts additional fields of type interface{}.*

## Structure

`DswitchesMetricsSwitchApAffinityDetails`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `SystemName` | `[]string` | Required | **Constraints**: *Unique Items Required* |
| `Threshold` | `float64` | Required | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "system_name": [
    "system_name8",
    "system_name9"
  ],
  "threshold": 195.06,
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

