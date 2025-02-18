
# Sle Impact Summary Band Item

*This model accepts additional fields of type interface{}.*

## Structure

`SleImpactSummaryBandItem`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Band` | `string` | Required | **Constraints**: *Minimum Length*: `1` |
| `Degraded` | `float64` | Required | - |
| `Duration` | `float64` | Required | - |
| `Name` | `string` | Required | **Constraints**: *Minimum Length*: `1` |
| `Total` | `float64` | Required | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "band": "band0",
  "degraded": 86.18,
  "duration": 215.24,
  "name": "name8",
  "total": 114.18,
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

