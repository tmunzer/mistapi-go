
# Org Sle Stat

## Structure

`OrgSleStat`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Path` | `string` | Required | - |
| `UserMinutes` | [`*models.OrgSleStatUserMinutes`](../../doc/models/org-sle-stat-user-minutes.md) | Optional | - |

## Example (as JSON)

```json
{
  "path": "path6",
  "user_minutes": {
    "ok": 13.84,
    "total": 12.38
  }
}
```

