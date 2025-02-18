
# Sle Threshold

*This model accepts additional fields of type interface{}.*

## Structure

`SleThreshold`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Default` | `*float64` | Optional | - |
| `Direction` | `*string` | Optional | **Constraints**: *Minimum Length*: `1` |
| `Maximum` | `*float64` | Optional | - |
| `Metric` | `*string` | Optional | **Constraints**: *Minimum Length*: `1` |
| `Minimum` | `*float64` | Optional | - |
| `Threshold` | `*string` | Optional | **Constraints**: *Minimum Length*: `1` |
| `Units` | `*string` | Optional | **Constraints**: *Minimum Length*: `1` |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "default": 158.88,
  "direction": "direction4",
  "maximum": 155.3,
  "metric": "metric0",
  "minimum": 127.18,
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

