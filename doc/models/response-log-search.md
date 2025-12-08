
# Response Log Search

## Structure

`ResponseLogSearch`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `End` | `int` | Required | - |
| `Limit` | `int` | Required | - |
| `Next` | `*string` | Optional | - |
| `Results` | [`[]models.LogEvent`](../../doc/models/log-event.md) | Required | **Constraints**: *Unique Items Required* |
| `Start` | `int` | Required | - |
| `Total` | `int` | Required | - |

## Example (as JSON)

```json
{
  "end": 48,
  "limit": 134,
  "results": [
    {
      "id": "53f10664-3ce8-4c27-b382-0ef66432349f",
      "message": "message6",
      "org_id": "a97c1b22-a4e9-411e-9bfd-d8695a0f9e61",
      "timestamp": 2.64,
      "admin_id": "00000014-0000-0000-0000-000000000000",
      "admin_name": "admin_name4",
      "after": {
        "key1": "val1",
        "key2": "val2"
      },
      "before": {
        "key1": "val1",
        "key2": "val2"
      },
      "for_site": false
    }
  ],
  "start": 6,
  "total": 40,
  "next": "next8"
}
```

