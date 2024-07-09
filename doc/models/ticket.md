
# Ticket

Support Ticket

## Structure

`Ticket`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `CaseNumber` | `*string` | Optional | - |
| `Comments` | [`[]models.TicketComment`](../../doc/models/ticket-comment.md) | Optional | - |
| `CreatedAt` | `*int` | Optional | - |
| `Id` | `*uuid.UUID` | Optional | - |
| `Requester` | `*string` | Optional | - |
| `RequesterEmail` | `*string` | Optional | email of the requester |
| `Status` | [`*models.TicketStatusEnum`](../../doc/models/ticket-status-enum.md) | Optional | Status open: ticket is open, Mist is working on it<br><br>* pending: ticket is open and Requester attention is needed (e.g. Mist is asking for some more information)<br>* solved: ticket is marked as solved / considered by Mist (requester can update it, causing it to re-open; or rate it)<br>* closed: ticket is archived and cannot be changed |
| `Subject` | `string` | Required | - |
| `Type` | `string` | Required | question (default) / bug / critical |
| `UpdatedAt` | `*int` | Optional | - |

## Example (as JSON)

```json
{
  "case_number": "case_number0",
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
          "created_at": 174,
          "file_name": "file_name8",
          "id": "000001f4-0000-0000-0000-000000000000",
          "size_in_bytes": 42
        }
      ],
      "author": "author2",
      "comment": "comment8",
      "created_at": 176
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
          "created_at": 174,
          "file_name": "file_name8",
          "id": "000001f4-0000-0000-0000-000000000000",
          "size_in_bytes": 42
        }
      ],
      "author": "author2",
      "comment": "comment8",
      "created_at": 176
    }
  ],
  "created_at": 136,
  "id": "00001ebe-0000-0000-0000-000000000000",
  "requester": "requester8",
  "subject": "subject4",
  "type": "type0"
}
```

