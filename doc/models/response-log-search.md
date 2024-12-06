
# Response Log Search

*This model accepts additional fields of type interface{}.*

## Structure

`ResponseLogSearch`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `End` | `int` | Required | - |
| `Limit` | `int` | Required | - |
| `Next` | `*string` | Optional | - |
| `Results` | [`[]models.ResponseLogSearchItem`](../../doc/models/response-log-search-item.md) | Required | **Constraints**: *Unique Items Required* |
| `Start` | `int` | Required | - |
| `Total` | `int` | Required | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "end": 48,
  "limit": 134,
  "results": [
    {
      "admin_id": "00000014-0000-0000-0000-000000000000",
      "admin_name": "admin_name4",
      "id": "53f10664-3ce8-4c27-b382-0ef66432349f",
      "message": "message6",
      "org_id": "a97c1b22-a4e9-411e-9bfd-d8695a0f9e61",
      "site_id": "00001420-0000-0000-0000-000000000000",
      "timestamp": 2.64,
      "after": {
        "key1": "val1",
        "key2": "val2"
      },
      "before": {
        "key1": "val1",
        "key2": "val2"
      },
      "for_site": false,
      "src_ip": "src_ip0",
      "exampleAdditionalProperty": {
        "key1": "val1",
        "key2": "val2"
      }
    }
  ],
  "start": 6,
  "total": 40,
  "next": "next8",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

