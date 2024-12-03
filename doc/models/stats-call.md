
# Stats Call

*This model accepts additional fields of type interface{}.*

## Structure

`StatsCall`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `App` | `*string` | Optional | - |
| `AudioQuality` | `*int` | Optional | - |
| `EndTime` | `*int` | Optional | - |
| `Mac` | `*string` | Optional | - |
| `MeetingId` | `*string` | Optional | - |
| `OrgId` | `*uuid.UUID` | Optional | - |
| `Rating` | `*int` | Optional | - |
| `ScreenShareQuality` | `*int` | Optional | - |
| `SiteId` | `*uuid.UUID` | Optional | - |
| `StartTime` | `*int` | Optional | - |
| `VideoQuality` | `*int` | Optional | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "org_id": "a97c1b22-a4e9-411e-9bfd-d8695a0f9e61",
  "site_id": "441a1214-6928-442a-8e92-e1d34b8ec6a6",
  "app": "app6",
  "audio_quality": 28,
  "end_time": 148,
  "mac": "mac0",
  "meeting_id": "meeting_id2",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

