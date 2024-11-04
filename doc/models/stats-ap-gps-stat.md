
# Stats Ap Gps Stat

## Structure

`StatsApGpsStat`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Accuracy` | `*float64` | Optional | The estimated accuracy or accuracy of the GPS coordinates, measured in meters. |
| `Altitude` | `*float64` | Optional | The elevation of the AP above sea level, measured in meters. |
| `Latitude` | `*float64` | Optional | The geographic latitude of the AP, measured in degrees. |
| `Longitude` | `*float64` | Optional | The geographic longitude of the AP, measured in degrees. |
| `Src` | [`*models.StatsApGpsStatSrcEnum`](../../doc/models/stats-ap-gps-stat-src-enum.md) | Optional | The origin of the GPS data. enum:<br><br>* `gps`: from this device’s GPS estimates<br>* `other_ap` from neighboring device GPS estimates |
| `Timestamp` | `*float64` | Optional | The unix timestamp when the GPS data was recorded. |

## Example (as JSON)

```json
{
  "accuracy": 12.5,
  "altitude": 99.939,
  "latitude": 37.29548,
  "longitude": -122.03304,
  "timestamp": 1428949501,
  "src": "gps"
}
```

