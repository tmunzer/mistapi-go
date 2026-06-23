
# Ticket Upload Attachments File

Multipart upload payload containing a ticket attachment file

## Structure

`TicketUploadAttachmentsFile`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `File` | `*[]byte` | Optional | Ekahau or ibwave file |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    ticketUploadAttachmentsFile := models.TicketUploadAttachmentsFile{
        File:                 models.ToPointer(nil),
    }

}
```

