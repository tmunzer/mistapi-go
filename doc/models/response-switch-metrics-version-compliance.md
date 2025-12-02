
# Response Switch Metrics Version Compliance

## Structure

`ResponseSwitchMetricsVersionCompliance`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Details` | [`*models.ResponseSwitchMetricsVersionComplianceDetails`](../../doc/models/response-switch-metrics-version-compliance-details.md) | Optional | - |
| `Score` | `*int` | Optional | - |
| `TotalSwitchCount` | `*int` | Optional | - |

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
        ]
      }
    ]
  },
  "score": 228,
  "total_switch_count": 16
}
```

