
# Anomaly

Anomaly

## Structure

`Anomaly`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Events` | `[]string` | Required | - |
| `Since` | `*float64` | Optional | - |
| `SleBaseline` | `float64` | Required | - |
| `SleDeviation` | `float64` | Required | - |
| `Timestamp` | `float64` | Required | - |

## Example (as JSON)

```json
{
  "events": [
    "events4"
  ],
  "since": 194.04,
  "sle_baseline": 169.52,
  "sle_deviation": 83.98,
  "timestamp": 63.62
}
```

