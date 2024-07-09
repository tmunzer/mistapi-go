
# Zone Stats Details Client Waits

client wait time right now

## Structure

`ZoneStatsDetailsClientWaits`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Avg` | `int` | Required | average wait time in seconds |
| `Max` | `int` | Required | longest wait time in seconds |
| `Min` | `int` | Required | shortest wait time in seconds |
| `P95` | `int` | Required | 95th percentile of all the wait time(s) |

## Example (as JSON)

```json
{
  "avg": 1200,
  "max": 3610,
  "min": 600,
  "p95": 2800
}
```
