
# Map Wall Path

JSON blob for wall definition (same format as wayfinding_path)

*This model accepts additional fields of type interface{}.*

## Structure

`MapWallPath`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Coordinate` | `*string` | Optional | - |
| `Nodes` | [`[]models.MapNode`](../../doc/models/map-node.md) | Optional | **Constraints**: *Minimum Items*: `0` |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

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
        "y": 100.0,
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
  ],
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

