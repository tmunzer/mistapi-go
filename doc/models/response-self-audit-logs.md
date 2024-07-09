
# Response Self Audit Logs

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

## Example (as JSON)

```json
{
  "end": 48,
  "limit": 122,
  "page": 236,
  "results": [
    {
      "admin_id": "00000014-0000-0000-0000-000000000000",
      "admin_name": "admin_name4",
      "id": "000023ba-0000-0000-0000-000000000000",
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
      "for_site": false
    }
  ],
  "start": 6,
  "total": 216
}
```

