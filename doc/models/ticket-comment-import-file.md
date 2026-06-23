
# Ticket Comment Import File

Multipart payload for adding a ticket comment with an attachment

## Structure

`TicketCommentImportFile`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Comment` | `*string` | Optional | Text body for the ticket comment submitted with the uploaded file |
| `File` | `*[]byte` | Optional | Binary file payload to attach to the ticket comment |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    ticketCommentImportFile := models.TicketCommentImportFile{
        Comment:              models.ToPointer("this is urgent"),
        File:                 models.ToPointer(nil),
    }

}
```

