
# Map Wayfinding Path

JSON blob for wayfinding (using Dijkstra’s algorithm)

## Structure

`MapWayfindingPath`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Coordinate` | `*string` | Optional | - |
| `Nodes` | [`[]models.MapNode`](../../doc/models/map-node.md) | Optional | **Constraints**: *Minimum Items*: `0` |

## Example (as JSON)

```json
{
  "coordinate": "actual",
  "nodes": [
    {
      "edges": {
        "key0": "edges6"
      },
      "name": "name6",
      "position": {
        "x": 224.72,
        "y": 100.0
      }
    },
    {
      "edges": {
        "key0": "edges6"
      },
      "name": "name6",
      "position": {
        "x": 224.72,
        "y": 100.0
      }
    }
  ]
}
```

