
# Response Auto Map Assignment Apply

## Structure

`ResponseAutoMapAssignmentApply`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AcceptedMaps` | `[]uuid.UUID` | Required | List of map IDs that were successfully accepted |
| `Message` | `string` | Required | Human-readable description of the operation result |

## Example (as JSON)

```json
{
  "accepted_maps": [
    "00002229-0000-0000-0000-000000000000",
    "00002228-0000-0000-0000-000000000000",
    "00002227-0000-0000-0000-000000000000"
  ],
  "message": "message0"
}
```

