
# Stats Ap Gps Stat

*This model accepts additional fields of type interface{}.*

## Structure

`StatsApGpsStat`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Accuracy` | `*float64` | Optional | The estimated accuracy or accuracy of the GPS coordinates, measured in meters. |
| `Altitude` | `*float64` | Optional | The elevation of the AP above sea level, measured in meters. |
| `Latitude` | `*float64` | Optional | The geographic latitude of the AP, measured in degrees. |
| `Longitude` | `*float64` | Optional | The geographic longitude of the AP, measured in degrees. |
| `Src` | [`*models.StatsApGpsStatSrcEnum`](../../doc/models/stats-ap-gps-stat-src-enum.md) | Optional | The origin of the GPS data. enum: `gps`: from this device GPS estimates, `other_ap` from neighboring device GPS estimates. Note: API responses may return `other_aps` which should be treated as `other_ap` |
| `Timestamp` | `*float64` | Optional | Epoch (seconds) |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "accuracy": 12.5,
  "altitude": 99.939,
  "latitude": 37.29548,
  "longitude": -122.03304,
  "src": "gps",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

