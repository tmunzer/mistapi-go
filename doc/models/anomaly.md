
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
    "events6",
    "events5",
    "events4"
  ],
  "since": 193.66,
  "sle_baseline": 169.14,
  "sle_deviation": 83.6,
  "timestamp": 63.24
}
```

