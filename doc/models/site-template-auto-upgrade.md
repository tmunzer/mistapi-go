
# Site Template Auto Upgrade

## Structure

`SiteTemplateAutoUpgrade`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `DayOfWeek` | [`*models.DayOfWeekEnum`](../../doc/models/day-of-week-enum.md) | Optional | enum: `any`, `fri`, `mon`, `sat`, `sun`, `thu`, `tue`, `wed` |
| `Enabled` | `*bool` | Optional | - |
| `TimeOfDay` | `*string` | Optional | - |
| `Version` | `*string` | Optional | - |

## Example (as JSON)

```json
{
  "day_of_week": "tue",
  "enabled": false,
  "time_of_day": "time_of_day8",
  "version": "version4"
}
```

