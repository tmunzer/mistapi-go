
# Site Template

## Structure

`SiteTemplate`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AutoUpgrade` | [`*models.SiteTemplateAutoUpgrade`](../../doc/models/site-template-auto-upgrade.md) | Optional | - |
| `Name` | `*string` | Optional | - |
| `Vars` | `map[string]string` | Optional | a dictionary of name->value, the vars can then be used in Wlans. This can overwrite those from Site Vars |

## Example (as JSON)

```json
{
  "vars": {
    "RADIUS_IP1": "172.31.2.5",
    "RADIUS_SECRET": "11s64632d"
  },
  "auto_upgrade": {
    "day_of_week": "sun",
    "enabled": false,
    "time_of_day": "time_of_day0",
    "version": "version2"
  },
  "name": "name8"
}
```

