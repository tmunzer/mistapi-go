
# Sle Impacted Switches

*This model accepts additional fields of type interface{}.*

## Structure

`SleImpactedSwitches`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Classifier` | `*string` | Optional | - |
| `End` | `*int` | Optional | - |
| `Failure` | `*string` | Optional | - |
| `Limit` | `*int` | Optional | - |
| `Metric` | `*string` | Optional | - |
| `Page` | `*int` | Optional | - |
| `Start` | `*int` | Optional | - |
| `Switches` | [`[]models.SleImpactedSwitchesSwitch`](../../doc/models/sle-impacted-switches-switch.md) | Optional | - |
| `TotalCount` | `*int` | Optional | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "classifier": "classifier8",
  "end": 74,
  "failure": "failure6",
  "limit": 96,
  "metric": "metric4",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

