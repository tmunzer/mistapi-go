
# Response Auto Map Assignment Info

## Structure

`ResponseAutoMapAssignmentInfo`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `EstTimeLeft` | `*float64` | Optional | Only when `status`==`in_progress`, estimated seconds remaining |
| `StartTime` | `*float64` | Optional | Unix timestamp when auto map assignment was started |
| `Status` | [`models.ResponseAutoMapAssignmentInfoStatusEnum`](../../doc/models/response-auto-map-assignment-info-status-enum.md) | Required | The status of auto map assignment for a given site. enum:<br><br>* `not_started`: Auto map assignment has not been requested<br>* `in_progress`: Auto map assignment is currently processing<br>* `completed`: The auto map assignment process has completed<br>* `error`: There was an error in the auto map assignment process |
| `StopTime` | `*float64` | Optional | Only when `status`==`completed`, Unix timestamp when auto map assignment stopped |
| `TimeUpdated` | `*float64` | Optional | Unix timestamp when status was last updated |

## Example (as JSON)

```json
{
  "est_time_left": 164.66,
  "start_time": 135.86,
  "status": "completed",
  "stop_time": 188.64,
  "time_updated": 207.0
}
```

