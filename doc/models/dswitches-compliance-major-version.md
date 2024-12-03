
# Dswitches Compliance Major Version

*This model accepts additional fields of type interface{}.*

## Structure

`DswitchesComplianceMajorVersion`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `MajorCount` | `float64` | Required | - |
| `Model` | `string` | Required | - |
| `SystemNames` | `[]string` | Optional | **Constraints**: *Unique Items Required* |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "major_count": 95.34,
  "model": "model4",
  "system_names": [
    "system_names4",
    "system_names5",
    "system_names6"
  ],
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

