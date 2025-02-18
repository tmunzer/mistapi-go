
# Ticket Attachment

*This model accepts additional fields of type interface{}.*

## Structure

`TicketAttachment`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `ContentUrl` | `*string` | Optional | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "content_url": "https://api.mist.com/api/v1/forward/download?jwt=...",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

