
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
  "major_count": 19.3,
  "model": "model8",
  "system_names": [
    "system_names8",
    "system_names9",
    "system_names0"
  ]
}
```

