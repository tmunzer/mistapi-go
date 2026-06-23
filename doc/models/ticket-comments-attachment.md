
# Ticket Comments Attachment

Metadata for an attachment on a ticket comment

## Structure

`TicketCommentsAttachment`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `ContentType` | `*string` | Optional | MIME type of the ticket comment attachment |
| `ContentUrl` | `*string` | Optional | Download URL for the ticket comment attachment |
| `CreatedAt` | `*int` | Optional | Time when this ticket comment attachment was created, in epoch seconds |
| `FileName` | `*string` | Optional | Original file name of the ticket comment attachment |
| `Id` | `*uuid.UUID` | Optional, Read-only | Unique ID of the object instance in the Mist Organization |
| `SizeInBytes` | `*int` | Optional | Attachment file size, in bytes |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    ticketCommentsAttachment := models.TicketCommentsAttachment{
        ContentType:          models.ToPointer("image/png"),
        ContentUrl:           models.ToPointer("https://api.mist.com/api/v1/forward/download?jwt=..."),
        CreatedAt:            models.ToPointer(1453908369),
        FileName:             models.ToPointer("crash.png"),
        Id:                   models.ToPointer(uuid.MustParse("53f10664-3ce8-4c27-b382-0ef66432349f")),
        SizeInBytes:          models.ToPointer(1943),
    }

}
```

