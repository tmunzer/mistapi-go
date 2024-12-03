
# Sle Impacted Applications App

*This model accepts additional fields of type interface{}.*

## Structure

`SleImpactedApplicationsApp`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `App` | `*string` | Optional | - |
| `Degraded` | `*int` | Optional | - |
| `Duration` | `*int` | Optional | - |
| `Name` | `*string` | Optional | - |
| `Threshold` | `*int` | Optional | - |
| `Total` | `*int` | Optional | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "app": "app6",
  "degraded": 192,
  "duration": 42,
  "name": "name6",
  "threshold": 154,
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

