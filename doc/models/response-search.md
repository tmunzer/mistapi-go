
# Response Search

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

## Example (as JSON)

```json
{
  "limit": 68,
  "page": 182,
  "results": [
    {
      "id": "000023ba-0000-0000-0000-000000000000",
      "text": "text4",
      "type": "type4"
    }
  ],
  "total": 162,
  "next": "next0"
}
```

