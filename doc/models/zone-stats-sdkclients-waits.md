
# Zone Stats Sdkclients Waits

sdkclient wait time right now

## Structure

`ZoneStatsSdkclientsWaits`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Avg` | `*float64` | Optional | average wait time in seconds |
| `Max` | `*float64` | Optional | longest wait time in seconds |
| `Min` | `*float64` | Optional | shortest wait time in seconds |
| `P95` | `*float64` | Optional | 95th percentile of all the wait time(s) |

## Example (as JSON)

```json
{
  "avg": 1200.0,
  "max": 3610.0,
  "min": 600.0,
  "p95": 2800.0
}
```

