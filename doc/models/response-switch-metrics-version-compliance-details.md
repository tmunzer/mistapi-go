
# Response Switch Metrics Version Compliance Details

*This model accepts additional fields of type interface{}.*

## Structure

`ResponseSwitchMetricsVersionComplianceDetails`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `MajorVersions` | [`[]models.SwitchMetricsComplianceMajorVersion`](../../doc/models/switch-metrics-compliance-major-version.md) | Optional | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
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
}
```

