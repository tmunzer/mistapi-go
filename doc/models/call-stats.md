
# Call Stats

## Structure

`CallStats`

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

## Example (as JSON)

```json
{
  "org_id": "a97c1b22-a4e9-411e-9bfd-d8695a0f9e61",
  "site_id": "441a1214-6928-442a-8e92-e1d34b8ec6a6",
  "app": "app2",
  "audio_quality": 232,
  "end_time": 96,
  "mac": "mac4",
  "meeting_id": "meeting_id6"
}
```

