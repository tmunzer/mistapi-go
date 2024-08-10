
# Response Log Search

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

## Example (as JSON)

```json
{
  "end": 48,
  "limit": 134,
  "results": [
    {
      "admin_id": "00000014-0000-0000-0000-000000000000",
      "admin_name": "admin_name4",
      "message": "message6",
      "org_id": "00002492-0000-0000-0000-000000000000",
      "site_id": "441a1214-6928-442a-8e92-e1d34b8ec6a6",
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
      "id": "000023ba-0000-0000-0000-000000000000"
    }
  ],
  "start": 6,
  "total": 40,
  "next": "next8"
}
```

