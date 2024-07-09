
# Response Stats Calls

## Structure

`ResponseStatsCalls`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `End` | `*float64` | Optional | - |
| `Limit` | `*int` | Optional | - |
| `Next` | `*string` | Optional | - |
| `Results` | [`[]models.CallStats`](../../doc/models/call-stats.md) | Optional | - |
| `Start` | `*float64` | Optional | - |
| `Total` | `*int` | Optional | - |

## Example (as JSON)

```json
{
  "end": 2.24,
  "limit": 54,
  "next": "next4",
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
  "start": 214.3
}
```

