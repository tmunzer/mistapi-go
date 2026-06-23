
# Stats Call

Statistics record for a detected collaboration call, such as Zoom or Teams

## Structure

`StatsCall`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `App` | `*string` | Optional | Third-party collaboration application that reported the call statistics |
| `AudioQuality` | `*int` | Optional | Quality score reported for the call audio stream |
| `EndTime` | `*int` | Optional | Time when the call ended, in epoch seconds |
| `Mac` | `*string` | Optional | Client MAC address associated with the call statistics record |
| `MeetingId` | `*string` | Optional | Collaboration meeting identifier reported for the call |
| `OrgId` | `*uuid.UUID` | Optional, Read-only | Unique identifier of a Mist organization |
| `Rating` | `*int` | Optional | Overall call rating reported by the collaboration application, when available |
| `ScreenShareQuality` | `*int` | Optional | Quality score reported for screen sharing during the call |
| `SiteId` | `*uuid.UUID` | Optional, Read-only | Unique identifier of a Mist site |
| `StartTime` | `*int` | Optional | Time when the call started, in epoch seconds |
| `VideoQuality` | `*int` | Optional | Quality score reported for the call video stream |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    statsCall := models.StatsCall{
        App:                  models.ToPointer("app0"),
        AudioQuality:         models.ToPointer(146),
        EndTime:              models.ToPointer(10),
        Mac:                  models.ToPointer("mac6"),
        MeetingId:            models.ToPointer("meeting_id8"),
        OrgId:                models.ToPointer(uuid.MustParse("a97c1b22-a4e9-411e-9bfd-d8695a0f9e61")),
        SiteId:               models.ToPointer(uuid.MustParse("441a1214-6928-442a-8e92-e1d34b8ec6a6")),
    }

}
```

