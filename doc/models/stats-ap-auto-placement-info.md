
# Stats Ap Auto Placement Info

Additional information about auto placements AP data

*This model accepts additional fields of type interface{}.*

## Structure

`StatsApAutoPlacementInfo`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `ClusterNumber` | `*int` | Optional | All APs sharing a given cluster number can be placed relative to each other |
| `OrientationStats` | `*int` | Optional | The orientation of an AP |
| `ProbabilitySurface` | [`*models.StatsApAutoPlacementInfoProbabilitySurface`](../../doc/models/stats-ap-auto-placement-info-probability-surface.md) | Optional | Coordinates representing a circle where the AP is most likely exists in the event of an inaccurate placement result |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "cluster_number": 0,
  "orientation_stats": 0,
  "probability_surface": {
    "radius": 74.96,
    "radius_m": 19.46,
    "x": 93.54,
    "exampleAdditionalProperty": {
      "key1": "val1",
      "key2": "val2"
    }
  },
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

