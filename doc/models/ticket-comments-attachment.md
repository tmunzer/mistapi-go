
# Ticket Comments Attachment

## Structure

`TicketCommentsAttachment`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `ContentType` | `*string` | Optional | - |
| `ContentUrl` | `*string` | Optional | - |
| `CreatedAt` | `*int` | Optional | - |
| `FileName` | `*string` | Optional | - |
| `Id` | `*uuid.UUID` | Optional | Unique ID of the object instance in the Mist Organnization |
| `SizeInBytes` | `*int` | Optional | - |

## Example (as JSON)

```json
{
  "content_type": "image/png",
  "content_url": "https://api.mist.com/api/v1/forward/download?jwt=...",
  "created_at": 1453908369,
  "file_name": "crash.png",
  "id": "53f10664-3ce8-4c27-b382-0ef66432349f",
  "size_in_bytes": 1943
}
```

