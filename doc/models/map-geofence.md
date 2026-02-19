
# Map Geofence

## Structure

`MapGeofence`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Name` | `*string` | Optional | Name of the geofence |
| `Vertices` | [`[]models.MapGeofenceVertice`](../../doc/models/map-geofence-vertice.md) | Optional | List of vertices defining the geofence |

## Example (as JSON)

```json
{
  "name": "example",
  "vertices": [
    {
      "X": 86.66,
      "Y": 252.2
    },
    {
      "X": 86.66,
      "Y": 252.2
    },
    {
      "X": 86.66,
      "Y": 252.2
    }
  ]
}
```

