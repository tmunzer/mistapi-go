
# Map Node

Nodes on maps

## Structure

`MapNode`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Edges` | `map[string]string` | Optional | - |
| `Name` | `string` | Required | - |
| `Position` | [`*models.MapNodePosition`](../../doc/models/map-node-position.md) | Optional | - |

## Example (as JSON)

```json
{
  "edges": {
    "N1": "1"
  },
  "name": "N1",
  "position": {
    "x": 224.72,
    "y": 100.0
  }
}
```

