
# Site Setting Gateway Mgmt Auto Signature Update

*This model accepts additional fields of type interface{}.*

## Structure

`SiteSettingGatewayMgmtAutoSignatureUpdate`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `DayOfWeek` | [`*models.DayOfWeekEnum`](../../doc/models/day-of-week-enum.md) | Optional | enum: `any`, `fri`, `mon`, `sat`, `sun`, `thu`, `tue`, `wed` |
| `Enable` | `*bool` | Optional | **Default**: `true` |
| `TimeOfDay` | `*string` | Optional | Optional, Mist will decide the timing |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "enable": true,
  "day_of_week": "any",
  "time_of_day": "time_of_day8",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

