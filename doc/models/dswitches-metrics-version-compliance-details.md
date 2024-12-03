
# Dswitches Metrics Version Compliance Details

*This model accepts additional fields of type interface{}.*

## Structure

`DswitchesMetricsVersionComplianceDetails`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `MajorVersions` | [`[]models.DswitchesComplianceMajorVersion`](../../doc/models/dswitches-compliance-major-version.md) | Required | **Constraints**: *Unique Items Required* |
| `TotalSwitchCount` | `int` | Required | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
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
  "total_switch_count": 226,
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

