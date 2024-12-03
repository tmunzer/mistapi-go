
# Dswitches Metrics Switch Ap Affinity

*This model accepts additional fields of type interface{}.*

## Structure

`DswitchesMetricsSwitchApAffinity`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Details` | [`models.DswitchesMetricsSwitchApAffinityDetails`](../../doc/models/dswitches-metrics-switch-ap-affinity-details.md) | Required | - |
| `Score` | `float64` | Required | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "details": {
    "system_name": [
      "system_name0",
      "system_name1"
    ],
    "threshold": 56.78,
    "exampleAdditionalProperty": {
      "key1": "val1",
      "key2": "val2"
    }
  },
  "score": 250.06,
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

