
# Site Setting Auto Upgrade Esl

auto upgrade AP ESL. When both firmware and ESL auto-upgrade are enabled, ESL upgrade will be done only after firmware upgrade

## Structure

`SiteSettingAutoUpgradeEsl`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AllowDowngrade` | `*bool` | Optional | If true, it will allow downgrade to a lower version<br><br>**Default**: `false` |
| `CustomVersions` | `map[string]string` | Optional | Custom versions for different models. Property key is the model name (e.g. "AP41") |
| `DayOfWeek` | [`*models.DayOfWeekEnum`](../../doc/models/day-of-week-enum.md) | Optional | enum: `any`, `fri`, `mon`, `sat`, `sun`, `thu`, `tue`, `wed` |
| `Enabled` | `*bool` | Optional | Whether auto upgrade should happen (Note that Mist may auto-upgrade if the version is not supported)<br><br>**Default**: `false` |
| `TimeOfDay` | `*string` | Optional | `any` / HH:MM (24-hour format), upgrade will happen within up to 1-hour from this time |
| `Version` | `*string` | Optional | - |

## Example (as JSON)

```json
{
  "allow_downgrade": false,
  "custom_versions": {
    "AP41": "2.4.6",
    "AP61": "2.5.0"
  },
  "enabled": false,
  "time_of_day": "12:00",
  "version": "2.5.0",
  "day_of_week": "sun"
}
```

