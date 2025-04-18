
# Stats Zone Clients Waits

Client wait time right now

*This model accepts additional fields of type interface{}.*

## Structure

`StatsZoneClientsWaits`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Avg` | `*float64` | Optional | Average wait time in seconds |
| `Max` | `*float64` | Optional | Longest wait time in seconds |
| `Min` | `*float64` | Optional | Shortest wait time in seconds |
| `P95` | `*float64` | Optional | 95th percentile of all the wait time(s) |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "avg": 1200.0,
  "max": 3610.0,
  "min": 600.0,
  "p95": 2800.0,
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

