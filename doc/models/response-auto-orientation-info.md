
# Response Auto Orientation Info

*This model accepts additional fields of type interface{}.*

## Structure

`ResponseAutoOrientationInfo`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `EstTimeLeft` | `*float64` | Optional | Only when `status`==`inprogress`, estimate of the time to completion |
| `StartTime` | `*float64` | Optional | time when auto orient process was last queued for this map |
| `Status` | [`*models.ResponseAutoOrientationInfoStatusEnum`](../../doc/models/response-auto-orientation-info-status-enum.md) | Optional | The status of auto orient for a given map. enum:<br><br>* `pending`: Auto orient has not been requested for this map<br>* `inprogress`: Auto orient is currently processing<br>* `done`: The auto orient process has completed<br>* `error`: There was an error in the auto orient process |
| `StopTime` | `*float64` | Optional | time when auto orient completed or was manually stopped |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "est_time_left": 122.9,
  "start_time": 151.7,
  "status": "done",
  "stop_time": 157.08,
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

