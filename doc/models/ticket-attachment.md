
# Ticket Attachment

Download information for a support ticket attachment

## Structure

`TicketAttachment`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `ContentUrl` | `*string` | Optional | Download URL for the support ticket attachment |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    ticketAttachment := models.TicketAttachment{
        ContentUrl:           models.ToPointer("https://api.mist.com/api/v1/forward/download?jwt=..."),
    }

}
```

