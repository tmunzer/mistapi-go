
# Ticket

Support Ticket

*This model accepts additional fields of type interface{}.*

## Structure

`Ticket`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `CaseNumber` | `*string` | Optional | - |
| `Comments` | [`[]models.TicketComment`](../../doc/models/ticket-comment.md) | Optional | - |
| `CreatedAt` | `*int` | Optional | - |
| `Id` | `*uuid.UUID` | Optional | Unique ID of the object instance in the Mist Organnization |
| `Requester` | `*string` | Optional | - |
| `RequesterEmail` | `*string` | Optional | Email of the requester |
| `Status` | [`*models.TicketStatusEnum`](../../doc/models/ticket-status-enum.md) | Optional | Ticket status. enum:<br><br>* open: ticket is open, Mist is working on it<br>* pending: ticket is open and Requester attention is needed (e.g. Mist is asking for some more information)<br>* solved: ticket is marked as solved / considered by Mist (requester can update it, causing it to re-open; or rate it)<br>* closed: ticket is archived and cannot be changed. |
| `Subject` | `string` | Required | - |
| `Type` | `string` | Required | Question (default) / bug / critical |
| `UpdatedAt` | `*int` | Optional | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "id": "53f10664-3ce8-4c27-b382-0ef66432349f",
  "subject": "subject2",
  "type": "type6",
  "case_number": "case_number6",
  "comments": [
    {
      "attachment_ids": [
        "00001052-0000-0000-0000-000000000000",
        "00001053-0000-0000-0000-000000000000",
        "00001054-0000-0000-0000-000000000000"
      ],
      "attachments": [
        {
          "content_type": "content_type4",
          "content_url": "content_url4",
          "created_at": 174,
          "file_name": "file_name8",
          "id": "000001f4-0000-0000-0000-000000000000",
          "exampleAdditionalProperty": {
            "key1": "val1",
            "key2": "val2"
          }
        }
      ],
      "author": "author2",
      "comment": "comment8",
      "created_at": 176,
      "exampleAdditionalProperty": {
        "key1": "val1",
        "key2": "val2"
      }
    },
    {
      "attachment_ids": [
        "00001052-0000-0000-0000-000000000000",
        "00001053-0000-0000-0000-000000000000",
        "00001054-0000-0000-0000-000000000000"
      ],
      "attachments": [
        {
          "content_type": "content_type4",
          "content_url": "content_url4",
          "created_at": 174,
          "file_name": "file_name8",
          "id": "000001f4-0000-0000-0000-000000000000",
          "exampleAdditionalProperty": {
            "key1": "val1",
            "key2": "val2"
          }
        }
      ],
      "author": "author2",
      "comment": "comment8",
      "created_at": 176,
      "exampleAdditionalProperty": {
        "key1": "val1",
        "key2": "val2"
      }
    },
    {
      "attachment_ids": [
        "00001052-0000-0000-0000-000000000000",
        "00001053-0000-0000-0000-000000000000",
        "00001054-0000-0000-0000-000000000000"
      ],
      "attachments": [
        {
          "content_type": "content_type4",
          "content_url": "content_url4",
          "created_at": 174,
          "file_name": "file_name8",
          "id": "000001f4-0000-0000-0000-000000000000",
          "exampleAdditionalProperty": {
            "key1": "val1",
            "key2": "val2"
          }
        }
      ],
      "author": "author2",
      "comment": "comment8",
      "created_at": 176,
      "exampleAdditionalProperty": {
        "key1": "val1",
        "key2": "val2"
      }
    }
  ],
  "created_at": 160,
  "requester": "requester8",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

