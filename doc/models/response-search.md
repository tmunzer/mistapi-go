
# Response Search

*This model accepts additional fields of type interface{}.*

## Structure

`ResponseSearch`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Limit` | `int` | Required | - |
| `Next` | `*string` | Optional | - |
| `Page` | `int` | Required | - |
| `Results` | [`[]models.ResponseSearchItem`](../../doc/models/response-search-item.md) | Required | **Constraints**: *Unique Items Required* |
| `Total` | `int` | Required | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "limit": 68,
  "page": 182,
  "results": [
    {
      "id": "53f10664-3ce8-4c27-b382-0ef66432349f",
      "text": "text4",
      "type": "type4",
      "exampleAdditionalProperty": {
        "key1": "val1",
        "key2": "val2"
      }
    }
  ],
  "total": 162,
  "next": "next0",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

