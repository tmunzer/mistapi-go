
# Ticket Comment

*This model accepts additional fields of type interface{}.*

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
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "attachment_ids": [
    "00000000-0000-0000-0000-15231a659c78"
  ],
  "author": "author6",
  "comment": "comment2",
  "created_at": 248,
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
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

