
# Dswitches Compliance Major Version

## Structure

`DswitchesComplianceMajorVersion`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `MajorCount` | `float64` | Required | - |
| `Model` | `string` | Required | - |
| `SystemNames` | `[]string` | Optional | **Constraints**: *Unique Items Required* |

## Example (as JSON)

```json
{
  "major_count": 95.34,
  "model": "model4",
  "system_names": [
    "system_names4",
    "system_names5",
    "system_names6"
  ]
}
```

