
# Response Stats Calls

## Structure

`ResponseStatsCalls`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `End` | `*float64` | Optional | - |
| `Limit` | `*int` | Optional | - |
| `Next` | `*string` | Optional | - |
| `Results` | [`[]models.StatsCall`](../../doc/models/stats-call.md) | Optional | - |
| `Start` | `*float64` | Optional | - |
| `Total` | `*int` | Optional | - |

## Example (as JSON)

```json
{
  "end": 151.82,
  "limit": 164,
  "next": "next2",
  "results": [
    {
      "app": "app6",
      "audio_quality": 66,
      "end_time": 186,
      "mac": "mac0",
      "meeting_id": "meeting_id2"
    },
    {
      "app": "app6",
      "audio_quality": 66,
      "end_time": 186,
      "mac": "mac0",
      "meeting_id": "meeting_id2"
    }
  ],
  "start": 107.88
}
```

