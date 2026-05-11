
# Auto Map Assignment Request

## Structure

`AutoMapAssignmentRequest`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `MapIds` | `[]uuid.UUID` | Optional | Optional list of specific map IDs to apply/clear. If not provided or empty, all pending map assignments are accepted/rejected. |

## Example (as JSON)

```json
{
  "map_ids": [
    "00001c56-0000-0000-0000-000000000000",
    "00001c57-0000-0000-0000-000000000000",
    "00001c58-0000-0000-0000-000000000000"
  ]
}
```

