
# Switch Metrics Compliance Major Version

*This model accepts additional fields of type interface{}.*

## Structure

`SwitchMetricsComplianceMajorVersion`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `MajorCount` | `*int` | Optional | - |
| `MajorVersion` | `*string` | Optional | - |
| `Model` | `*string` | Optional | - |
| `SystemNames` | `[]string` | Optional | **Constraints**: *Unique Items Required* |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "major_count": 136,
  "major_version": "major_version6",
  "model": "model0",
  "system_names": [
    "system_names0",
    "system_names1"
  ],
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

