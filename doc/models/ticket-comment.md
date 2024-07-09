
# Ticket Comment

## Structure

`TicketComment`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AttachmentIds` | `[]uuid.UUID` | Optional | - |
| `Attachments` | [`[]models.TicketCommentsAttachment`](../../doc/models/ticket-comments-attachment.md) | Optional | - |
| `Author` | `string` | Required | - |
| `Comment` | `string` | Required | - |
| `CreatedAt` | `int` | Required | - |

## Example (as JSON)

```json
{
  "attachment_ids": [
    "00000000-0000-0000-0000-15231a659c78"
  ],
  "author": "author2",
  "comment": "comment8",
  "created_at": 182,
  "attachments": [
    {
      "content_type": "content_type4",
      "created_at": 174,
      "file_name": "file_name8",
      "id": "000001f4-0000-0000-0000-000000000000",
      "size_in_bytes": 42
    },
    {
      "content_type": "content_type4",
      "created_at": 174,
      "file_name": "file_name8",
      "id": "000001f4-0000-0000-0000-000000000000",
      "size_in_bytes": 42
    }
  ]
}
```

