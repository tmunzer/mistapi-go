
# Lat Lng

*This model accepts additional fields of type interface{}.*

## Structure

`LatLng`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Lat` | `float64` | Required | - |
| `Lng` | `float64` | Required | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "lat": 37.295833,
  "lng": -122.032946,
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

