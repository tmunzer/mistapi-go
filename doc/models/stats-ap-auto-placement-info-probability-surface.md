
# Stats Ap Auto Placement Info Probability Surface

Coordinates representing a circle where the AP is most likely exists in the event of an inaccurate placement result

*This model accepts additional fields of type interface{}.*

## Structure

`StatsApAutoPlacementInfoProbabilitySurface`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Radius` | `*float64` | Optional | The radius representing placement uncertainty, measured in pixels |
| `RadiusM` | `*float64` | Optional | The radius representing placement uncertainty, measured in meters |
| `X` | `*float64` | Optional | Y-coordinate of the potential placement’s center, measured in pixels |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "radius": 2.1,
  "x": 17.0,
  "radius_m": 200.0,
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

