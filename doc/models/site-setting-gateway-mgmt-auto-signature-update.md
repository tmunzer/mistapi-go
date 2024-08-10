
# Site Setting Gateway Mgmt Auto Signature Update

## Structure

`SiteSettingGatewayMgmtAutoSignatureUpdate`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `DayOfWeek` | [`*models.DayOfWeekEnum`](../../doc/models/day-of-week-enum.md) | Optional | enum: `any`, `fri`, `mon`, `sat`, `sun`, `thu`, `tue`, `wed` |
| `Enable` | `*bool` | Optional | **Default**: `true` |
| `TimeOfDay` | `*string` | Optional | optional, Mist will decide the timing |

## Example (as JSON)

```json
{
  "enable": true,
  "day_of_week": "any",
  "time_of_day": "time_of_day8"
}
```

