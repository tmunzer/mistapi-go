
# Ticket

Support ticket record with status, comments, and metadata

*This model accepts additional fields of type interface{}.*

## Structure

`Ticket`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `CaseNumber` | `*string` | Optional, Read-only | Support case number associated with this ticket |
| `Comments` | [`[]models.TicketComment`](../../doc/models/ticket-comment.md) | Optional | List of comments on a support ticket |
| `CreatedAt` | `*int` | Optional, Read-only | Time when this support ticket was created, in epoch seconds |
| `Id` | `*uuid.UUID` | Optional, Read-only | Unique ID of the object instance in the Mist Organization |
| `Requester` | `*string` | Optional, Read-only | User who opened or requested this support ticket |
| `RequesterEmail` | `*string` | Optional | Email of the requester |
| `Status` | [`*models.TicketStatusEnum`](../../doc/models/ticket-status-enum.md) | Optional | Ticket status. enum: `closed`, `open`, `pending`, `solved`. `open` means Mist is working on it, `pending` means requester attention is needed, `solved` means Mist considers it resolved but it can still be updated or rated, and `closed` means it is archived |
| `Subject` | `string` | Required | Short summary of the support request |
| `Type` | `string` | Required | Question (default) / bug / critical |
| `UpdatedAt` | `*int` | Optional, Read-only | Time when this support ticket was last updated, in epoch seconds |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    ticket := models.Ticket{
        CaseNumber:           models.ToPointer("case_number6"),
        Comments:             []models.TicketComment{
            models.TicketComment{
                AttachmentIds:        []uuid.UUID{
                    uuid.MustParse("00001052-0000-0000-0000-000000000000"),
                    uuid.MustParse("00001053-0000-0000-0000-000000000000"),
                    uuid.MustParse("00001054-0000-0000-0000-000000000000"),
                },
                Attachments:          []models.TicketCommentsAttachment{
                    models.TicketCommentsAttachment{
                        ContentType:          models.ToPointer("content_type4"),
                        ContentUrl:           models.ToPointer("content_url4"),
                        CreatedAt:            models.ToPointer(174),
                        FileName:             models.ToPointer("file_name8"),
                        Id:                   models.ToPointer(uuid.MustParse("000001f4-0000-0000-0000-000000000000")),
                    },
                },
                Author:               "author2",
                Comment:              "comment8",
                CreatedAt:            176,
                AdditionalProperties: map[string]interface{}{
                    "exampleAdditionalProperty": interface{}("[key1, val1][key2, val2]"),
                },
            },
            models.TicketComment{
                AttachmentIds:        []uuid.UUID{
                    uuid.MustParse("00001052-0000-0000-0000-000000000000"),
                    uuid.MustParse("00001053-0000-0000-0000-000000000000"),
                    uuid.MustParse("00001054-0000-0000-0000-000000000000"),
                },
                Attachments:          []models.TicketCommentsAttachment{
                    models.TicketCommentsAttachment{
                        ContentType:          models.ToPointer("content_type4"),
                        ContentUrl:           models.ToPointer("content_url4"),
                        CreatedAt:            models.ToPointer(174),
                        FileName:             models.ToPointer("file_name8"),
                        Id:                   models.ToPointer(uuid.MustParse("000001f4-0000-0000-0000-000000000000")),
                    },
                },
                Author:               "author2",
                Comment:              "comment8",
                CreatedAt:            176,
                AdditionalProperties: map[string]interface{}{
                    "exampleAdditionalProperty": interface{}("[key1, val1][key2, val2]"),
                },
            },
            models.TicketComment{
                AttachmentIds:        []uuid.UUID{
                    uuid.MustParse("00001052-0000-0000-0000-000000000000"),
                    uuid.MustParse("00001053-0000-0000-0000-000000000000"),
                    uuid.MustParse("00001054-0000-0000-0000-000000000000"),
                },
                Attachments:          []models.TicketCommentsAttachment{
                    models.TicketCommentsAttachment{
                        ContentType:          models.ToPointer("content_type4"),
                        ContentUrl:           models.ToPointer("content_url4"),
                        CreatedAt:            models.ToPointer(174),
                        FileName:             models.ToPointer("file_name8"),
                        Id:                   models.ToPointer(uuid.MustParse("000001f4-0000-0000-0000-000000000000")),
                    },
                },
                Author:               "author2",
                Comment:              "comment8",
                CreatedAt:            176,
                AdditionalProperties: map[string]interface{}{
                    "exampleAdditionalProperty": interface{}("[key1, val1][key2, val2]"),
                },
            },
        },
        CreatedAt:            models.ToPointer(160),
        Id:                   models.ToPointer(uuid.MustParse("53f10664-3ce8-4c27-b382-0ef66432349f")),
        Requester:            models.ToPointer("requester8"),
        Subject:              "subject2",
        Type:                 "type6",
        AdditionalProperties: map[string]interface{}{
            "exampleAdditionalProperty": interface{}("[key1, val1][key2, val2]"),
        },
    }

}
```

