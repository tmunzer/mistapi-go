
# Response Switch Metrics Version Compliance

*This model accepts additional fields of type interface{}.*

## Structure

`ResponseSwitchMetricsVersionCompliance`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Details` | [`*models.ResponseSwitchMetricsVersionComplianceDetails`](../../doc/models/response-switch-metrics-version-compliance-details.md) | Optional | - |
| `Score` | `*int` | Optional | - |
| `TotalSwitchCount` | `*int` | Optional | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "details": {
    "major_versions": [
      {
        "major_count": 170,
        "major_version": "major_version0",
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
    "exampleAdditionalProperty": {
      "key1": "val1",
      "key2": "val2"
    }
  },
  "score": 228,
  "total_switch_count": 16,
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

