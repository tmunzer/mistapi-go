
# Dswitches Metrics Version Compliance

*This model accepts additional fields of type interface{}.*

## Structure

`DswitchesMetricsVersionCompliance`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Details` | [`models.DswitchesMetricsVersionComplianceDetails`](../../doc/models/dswitches-metrics-version-compliance-details.md) | Required | - |
| `Score` | `float64` | Required | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "details": {
    "major_versions": [
      {
        "major_count": 47.78,
        "model": "model6",
        "system_names": [
          "system_names6",
          "system_names7"
        ],
        "exampleAdditionalProperty": {
          "key1": "val1",
          "key2": "val2"
        }
      }
    ],
    "total_switch_count": 14,
    "exampleAdditionalProperty": {
      "key1": "val1",
      "key2": "val2"
    }
  },
  "score": 9.44,
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

