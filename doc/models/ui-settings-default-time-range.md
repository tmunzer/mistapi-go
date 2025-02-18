
# Ui Settings Default Time Range

*This model accepts additional fields of type interface{}.*

## Structure

`UiSettingsDefaultTimeRange`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `End` | `*int` | Optional | - |
| `EndDate` | `*string` | Optional | - |
| `Interval` | `*string` | Optional | - |
| `Name` | `*string` | Optional | - |
| `ShortName` | `*string` | Optional | - |
| `Start` | `*int` | Optional | - |
| `UsePreset` | `*bool` | Optional | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "end": 1508828400,
  "endDate": "10/23/2017",
  "interval": "1d",
  "name": "This Week",
  "shortName": "thisWeek",
  "start": 1508655600,
  "usePreset": true,
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

