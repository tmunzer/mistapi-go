
# Response Self Audit Logs

*This model accepts additional fields of type interface{}.*

## Structure

`ResponseSelfAuditLogs`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `End` | `int` | Required | - |
| `Limit` | `int` | Required | - |
| `Page` | `int` | Required | - |
| `Results` | [`[]models.AuditLog`](../../doc/models/audit-log.md) | Required | **Constraints**: *Unique Items Required* |
| `Start` | `int` | Required | - |
| `Total` | `int` | Required | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "end": 138,
  "limit": 32,
  "page": 146,
  "results": [
    {
      "admin_id": "456b7016-a916-a4b1-78dd-72b947c152b7",
      "admin_name": "admin_name4",
      "id": "53f10664-3ce8-4c27-b382-0ef66432349f",
      "message": "message6",
      "org_id": "a97c1b22-a4e9-411e-9bfd-d8695a0f9e61",
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
      "exampleAdditionalProperty": {
        "key1": "val1",
        "key2": "val2"
      }
    }
  ],
  "start": 96,
  "total": 130,
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

