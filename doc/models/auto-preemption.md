
# Auto Preemption

schedule to preempt ap’s which are not connected to preferred peer

## Structure

`AutoPreemption`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `DayOfWeek` | [`*models.DayOfWeekEnum`](../../doc/models/day-of-week-enum.md) | Optional | enum: `any`, `fri`, `mon`, `sat`, `sun`, `thu`, `tue`, `wed` |
| `Enabled` | `*bool` | Optional | whether auto preemption should happen<br>**Default**: `false` |
| `TimeOfDay` | `*string` | Optional | any / HH:MM (24-hour format)<br>**Default**: `"any"` |

## Example (as JSON)

```json
{
  "enabled": false,
  "time_of_day": "12:00",
  "day_of_week": "mon"
}
```

