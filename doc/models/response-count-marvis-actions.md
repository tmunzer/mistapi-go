
# Response Count Marvis Actions

*This model accepts additional fields of type interface{}.*

## Structure

`ResponseCountMarvisActions`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Distinct` | `*string` | Optional | - |
| `Limit` | `*int` | Optional | - |
| `Results` | [`[]models.ResponseCountMarvisActionsResult`](../../doc/models/response-count-marvis-actions-result.md) | Optional | - |
| `Total` | `*int` | Optional | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "distinct": "status",
  "limit": 1000,
  "results": [
    {
      "count": 24,
      "status": "002e176a-0000-000-1111-002e208b20e1"
    },
    {
      "count": 12,
      "status": "2d3f176a-0000-000-2222-002e208f176a"
    },
    {
      "count": 15,
      "status": "08b2176a-0000-000-3333-002e208b2d3f"
    }
  ],
  "total": 3,
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

