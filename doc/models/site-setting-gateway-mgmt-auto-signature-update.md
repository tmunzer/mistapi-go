
# Site Setting Gateway Mgmt Auto Signature Update

## Structure

`SiteSettingGatewayMgmtAutoSignatureUpdate`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `DayOfWeek` | [`*models.DayOfWeekEnum`](../../doc/models/day-of-week-enum.md) | Optional | - |
| `Enable` | `*bool` | Optional | **Default**: `true` |
| `TimeOfDay` | `*string` | Optional | optional, Mist will decide the timing |

## Example (as JSON)

```json
{
  "enable": true,
  "day_of_week": "tue",
  "time_of_day": "time_of_day2"
}
```

