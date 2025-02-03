
# Response Auto Placement Info

*This model accepts additional fields of type interface{}.*

## Structure

`ResponseAutoPlacementInfo`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `EndTime` | `*float64` | Optional | Time when autoplacement completed or was manually stopped |
| `EstTimeLeft` | `*float64` | Optional | (Only when inprogress) estimate of the time to completion |
| `StartTime` | `*int` | Optional | Time when autoplacement process was last queued for this map |
| `Status` | [`*models.AutoPlacementInfoStatusEnum`](../../doc/models/auto-placement-info-status-enum.md) | Optional | the status of autoplacement for a given map. enum: `done`, `error`, `inprogress`, `pending` |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "end_time": 184.76,
  "est_time_left": 52.68,
  "start_time": 84,
  "status": "done",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

