
# Ticket Comment

Comment on a support ticket

*This model accepts additional fields of type interface{}.*

## Structure

`TicketComment`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AttachmentIds` | `[]uuid.UUID` | Optional, Read-only | List of ticket comment attachment identifiers |
| `Attachments` | [`[]models.TicketCommentsAttachment`](../../doc/models/ticket-comments-attachment.md) | Optional, Read-only | List of attachment metadata objects for a ticket comment |
| `Author` | `string` | Required, Read-only | User who wrote this ticket comment |
| `Comment` | `string` | Required | Text body of this ticket comment |
| `CreatedAt` | `int` | Required, Read-only | Time when this ticket comment was created, in epoch seconds |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    ticketComment := models.TicketComment{
        AttachmentIds:        []uuid.UUID{
            uuid.MustParse("00000000-0000-0000-0000-15231a659c78"),
        },
        Attachments:          []models.TicketCommentsAttachment{
            models.TicketCommentsAttachment{
                ContentType:          models.ToPointer("content_type4"),
                ContentUrl:           models.ToPointer("content_url4"),
                CreatedAt:            models.ToPointer(174),
                FileName:             models.ToPointer("file_name8"),
                Id:                   models.ToPointer(uuid.MustParse("000001f4-0000-0000-0000-000000000000")),
            },
            models.TicketCommentsAttachment{
                ContentType:          models.ToPointer("content_type4"),
                ContentUrl:           models.ToPointer("content_url4"),
                CreatedAt:            models.ToPointer(174),
                FileName:             models.ToPointer("file_name8"),
                Id:                   models.ToPointer(uuid.MustParse("000001f4-0000-0000-0000-000000000000")),
            },
        },
        Author:               "author8",
        Comment:              "comment2",
        CreatedAt:            104,
        AdditionalProperties: map[string]interface{}{
            "exampleAdditionalProperty": interface{}("[key1, val1][key2, val2]"),
        },
    }

}
```

