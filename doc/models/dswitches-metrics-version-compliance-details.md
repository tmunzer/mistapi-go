
# Dswitches Metrics Version Compliance Details

## Structure

`DswitchesMetricsVersionComplianceDetails`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `MajorVersions` | [`[]models.DswitchesComplianceMajorVersion`](../../doc/models/dswitches-compliance-major-version.md) | Required | **Constraints**: *Unique Items Required* |
| `TotalSwitchCount` | `int` | Required | - |

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
      ]
    }
  ],
  "total_switch_count": 226
}
```

