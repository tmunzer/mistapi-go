
# Log Event

Audit log event recorded for an organization or site

## Structure

`LogEvent`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AdminId` | `models.Optional[uuid.UUID]` | Optional, Read-only | Admin user identifier associated with the log event |
| `AdminName` | `models.Optional[string]` | Optional, Read-only | Name of the admin that performs the action |
| `After` | `*interface{}` | Optional, Read-only | field values after the change |
| `Before` | `*interface{}` | Optional, Read-only | field values prior to the change |
| `DeviceId` | `models.Optional[uuid.UUID]` | Optional, Read-only | Device identifier associated with the log event |
| `ForSite` | `*bool` | Optional, Read-only | Whether this log event is scoped to a site |
| `Id` | `*uuid.UUID` | Optional, Read-only | Unique ID of the object instance in the Mist Organization |
| `Message` | `string` | Required, Read-only | Human-readable log message describing the event |
| `OrgId` | `uuid.UUID` | Required, Read-only | Unique identifier of a Mist organization |
| `SiteId` | `models.Optional[uuid.UUID]` | Optional, Read-only | Site associated with the log event, if any |
| `SrcIp` | `*string` | Optional | sender source IP address |
| `Timestamp` | `float64` | Required, Read-only | Epoch timestamp, in seconds |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    logEvent := models.LogEvent{
        AdminId:              models.NewOptional(models.ToPointer(uuid.MustParse("00000244-0000-0000-0000-000000000000"))),
        AdminName:            models.NewOptional(models.ToPointer("admin_name4")),
        After:                models.ToPointer(interface{}("[key1, val1][key2, val2]")),
        Before:               models.ToPointer(interface{}("[key1, val1][key2, val2]")),
        DeviceId:             models.NewOptional(models.ToPointer(uuid.MustParse("00001740-0000-0000-0000-000000000000"))),
        Id:                   models.ToPointer(uuid.MustParse("53f10664-3ce8-4c27-b382-0ef66432349f")),
        Message:              "message6",
        OrgId:                uuid.MustParse("a97c1b22-a4e9-411e-9bfd-d8695a0f9e61"),
        Timestamp:            float64(8.24),
    }

}
```

