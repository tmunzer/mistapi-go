
# Stats Org Sle

## Structure

`StatsOrgSle`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Path` | `string` | Required | - |
| `UserMinutes` | [`*models.StatsOrgSleUserMinutes`](../../doc/models/stats-org-sle-user-minutes.md) | Optional | - |

## Example (as JSON)

```json
{
  "path": "path0",
  "user_minutes": {
    "ok": 13.84,
    "total": 12.38
  }
}
```

