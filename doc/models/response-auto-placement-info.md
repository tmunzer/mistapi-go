
# Response Auto Placement Info

## Structure

`ResponseAutoPlacementInfo`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `EndTime` | `*float64` | Optional | time when autoplacement completed or was manually stopped |
| `EstTimeLeft` | `*float64` | Optional | (Only when inprogress) estimate of the time to completion |
| `StartTime` | `*float64` | Optional | time when autoplacement process was last queued for this map |
| `Status` | [`*models.AutoPlacementInfoStatusEnum`](../../doc/models/auto-placement-info-status-enum.md) | Optional | the status of autoplacement for a given map |

## Example (as JSON)

```json
{
  "end_time": 38.78,
  "est_time_left": 93.3,
  "start_time": 122.1,
  "status": "done"
}
```

