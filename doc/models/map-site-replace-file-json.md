
# Map Site Replace File Json

## Structure

`MapSiteReplaceFileJson`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Transform` | [`*models.MapSiteReplaceFileJsonTransform`](../../doc/models/map-site-replace-file-json-transform.md) | Optional | If `transform` is provided, all the locations of the objects on the map (AP, Zone, Vbeacon, Beacon) will be transformed as well (relative to the new Map) |

## Example (as JSON)

```json
{
  "transform": {
    "rotation": 130.62,
    "scale": 195.12,
    "x": 28.38,
    "y": 102.9
  }
}
```

