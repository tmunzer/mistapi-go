
# Map Site Replace File Json Transform

If `transform` is provided, all the locations of the objects on the map (AP, Zone, Vbeacon, Beacon) will be transformed as well (relative to the new Map)

## Structure

`MapSiteReplaceFileJsonTransform`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Rotation` | `*float64` | Optional | whether to rotate the replacing image, in degrees<br>**Default**: `0` |
| `Scale` | `*float64` | Optional | whether to scale the replacing image<br>**Default**: `1` |
| `X` | `*float64` | Optional | where the (0, 0) of the new image is relative to the original map<br>**Default**: `0` |
| `Y` | `*float64` | Optional | where the (0, 0) of the new image is relative to the original map<br>**Default**: `0` |

## Example (as JSON)

```json
{
  "rotation": 0.0,
  "scale": 0.98,
  "x": 3.16,
  "y": 12.0
}
```

