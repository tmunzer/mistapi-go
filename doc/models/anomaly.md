
# Anomaly

Anomaly

*This model accepts additional fields of type interface{}.*

## Structure

`Anomaly`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Events` | `[]string` | Required | - |
| `Since` | `*float64` | Optional | - |
| `SleBaseline` | `float64` | Required | - |
| `SleDeviation` | `float64` | Required | - |
| `Timestamp` | `float64` | Required | Epoch (seconds) |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "events": [
    "events4"
  ],
  "since": 194.04,
  "sle_baseline": 169.52,
  "sle_deviation": 83.98,
  "timestamp": 63.62,
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

