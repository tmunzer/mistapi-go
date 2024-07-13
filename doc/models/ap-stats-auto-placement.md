
# Ap Stats Auto Placement

## Structure

`ApStatsAutoPlacement`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Info` | [`*models.ApStatsAutoPlacementInfo`](../../doc/models/ap-stats-auto-placement-info.md) | Optional | Additional information about auto placements AP data |
| `RecommendedAnchor` | `*bool` | Optional | Flag to represent if AP is recommended as an anchor by auto placement service |
| `Status` | `*string` | Optional | Basic Placement Status |
| `StatusDetail` | `*string` | Optional | Additional info about placement status |
| `UseAutoPlacement` | `*bool` | Optional | Flag to represent if auto_placement values are currently utilized |
| `X` | `*float64` | Optional | X Autoplaced Position in pixels |
| `XM` | `*float64` | Optional | X Autoplaced Position in meters |
| `Y` | `*float64` | Optional | Y Autoplaced Position in pixels |
| `YM` | `*float64` | Optional | X Autoplaced Position in meters |

## Example (as JSON)

```json
{
  "status": "localized",
  "status_detail": "localized",
  "x": 53.5,
  "x_m": 5.35,
  "y": 173.1,
  "y_m": 17.31,
  "info": {
    "cluster_number": 112,
    "orientation_stats": 90,
    "probability_surface": {
      "radius": 74.96,
      "radius_m": 19.46,
      "x": 93.54
    }
  },
  "recommended_anchor": false,
  "use_auto_placement": false
}
```

