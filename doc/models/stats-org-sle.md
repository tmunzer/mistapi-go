
# Stats Org Sle

*This model accepts additional fields of type interface{}.*

## Structure

`StatsOrgSle`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Path` | `string` | Required | - |
| `UserMinutes` | [`*models.StatsOrgSleUserMinutes`](../../doc/models/stats-org-sle-user-minutes.md) | Optional | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "path": "path0",
  "user_minutes": {
    "ok": 13.84,
    "total": 12.38,
    "exampleAdditionalProperty": {
      "key1": "val1",
      "key2": "val2"
    }
  },
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

