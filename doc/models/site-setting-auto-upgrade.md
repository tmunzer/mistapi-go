
# Site Setting Auto Upgrade

Auto Upgrade Settings

*This model accepts additional fields of type interface{}.*

## Structure

`SiteSettingAutoUpgrade`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `CustomVersions` | `map[string]string` | Optional | Custom versions for different models. Property key is the model name (e.g. "AP41") |
| `DayOfWeek` | [`*models.DayOfWeekEnum`](../../doc/models/day-of-week-enum.md) | Optional | enum: `any`, `fri`, `mon`, `sat`, `sun`, `thu`, `tue`, `wed` |
| `Enabled` | `*bool` | Optional | Whether auto upgrade should happen (Note that Mist may auto-upgrade if the version is not supported)<br><br>**Default**: `false` |
| `TimeOfDay` | `*string` | Optional | `any` / HH:MM (24-hour format), upgrade will happen within up to 1-hour from this time |
| `Version` | [`*models.SiteAutoUpgradeVersionEnum`](../../doc/models/site-auto-upgrade-version-enum.md) | Optional | desired version. enum: `beta`, `custom`, `stable`<br><br>**Default**: `"stable"` |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "custom_versions": {
    "AP21": "stable",
    "AP41": "0.1.5135",
    "AP61": "0.1.7215"
  },
  "enabled": false,
  "time_of_day": "12:00",
  "version": "beta",
  "day_of_week": "any",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

