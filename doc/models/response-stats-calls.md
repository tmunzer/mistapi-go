
# Response Stats Calls

*This model accepts additional fields of type interface{}.*

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
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

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
      "meeting_id": "meeting_id2",
      "exampleAdditionalProperty": {
        "key1": "val1",
        "key2": "val2"
      }
    },
    {
      "app": "app6",
      "audio_quality": 66,
      "end_time": 186,
      "mac": "mac0",
      "meeting_id": "meeting_id2",
      "exampleAdditionalProperty": {
        "key1": "val1",
        "key2": "val2"
      }
    }
  ],
  "start": 107.88,
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

