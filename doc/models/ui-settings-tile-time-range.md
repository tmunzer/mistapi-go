
# Ui Settings Tile Time Range

*This model accepts additional fields of type interface{}.*

## Structure

`UiSettingsTileTimeRange`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `End` | `*float64` | Optional | - |
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
  "end": 1508823743.0,
  "endDate": "10/23/2017",
  "interval": "1d",
  "name": "Past 7 Days",
  "shortName": "7d",
  "start": 1508223600,
  "usePreset": true,
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

