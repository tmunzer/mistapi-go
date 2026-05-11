
# Response Auto Map Assignment Clear

## Structure

`ResponseAutoMapAssignmentClear`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Message` | `string` | Required | Human-readable description of the operation result |
| `RejectedMaps` | `[]uuid.UUID` | Required | List of map IDs that were successfully rejected |

## Example (as JSON)

```json
{
  "message": "message0",
  "rejected_maps": [
    "000015b6-0000-0000-0000-000000000000"
  ]
}
```

