
# Stats Zone Details Client Waits

Client wait time right now

*This model accepts additional fields of type interface{}.*

## Structure

`StatsZoneDetailsClientWaits`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Avg` | `int` | Required | Average wait time in seconds |
| `Max` | `int` | Required | Longest wait time in seconds |
| `Min` | `int` | Required | Shortest wait time in seconds |
| `P95` | `int` | Required | 95th percentile of all the wait time(s) |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "avg": 1200,
  "max": 3610,
  "min": 600,
  "p95": 2800,
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

