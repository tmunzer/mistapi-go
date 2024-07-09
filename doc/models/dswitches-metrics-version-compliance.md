
# Dswitches Metrics Version Compliance

## Structure

`DswitchesMetricsVersionCompliance`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Details` | [`models.DswitchesMetricsVersionComplianceDetails`](../../doc/models/dswitches-metrics-version-compliance-details.md) | Required | - |
| `Score` | `float64` | Required | - |

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
        ]
      }
    ],
    "total_switch_count": 14
  },
  "score": 141.72
}
```

