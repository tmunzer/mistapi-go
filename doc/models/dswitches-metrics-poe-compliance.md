
# Dswitches Metrics Poe Compliance

*This model accepts additional fields of type interface{}.*

## Structure

`DswitchesMetricsPoeCompliance`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Details` | [`models.DswitchesMetricsPoeComplianceDetails`](../../doc/models/dswitches-metrics-poe-compliance-details.md) | Required | - |
| `Score` | `float64` | Required | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "details": {
    "total_aps": 132,
    "total_power": 137.54,
    "exampleAdditionalProperty": {
      "key1": "val1",
      "key2": "val2"
    }
  },
  "score": 126.52,
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

