
# Site Setting Auto Placement

if we're able to determine its x/y/orientation, this will be populated

*This model accepts additional fields of type interface{}.*

## Structure

`SiteSettingAutoPlacement`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Orientation` | `*int` | Optional | - |
| `X` | `*float64` | Optional | - |
| `Y` | `*float64` | Optional | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "orientation": 45,
  "x": 30.0,
  "y": 60.0,
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

