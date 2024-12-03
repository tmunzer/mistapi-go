
# Map Node

Nodes on maps

*This model accepts additional fields of type interface{}.*

## Structure

`MapNode`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Edges` | `map[string]string` | Optional | - |
| `Name` | `string` | Required | - |
| `Position` | [`*models.MapNodePosition`](../../doc/models/map-node-position.md) | Optional | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "edges": {
    "N1": "1"
  },
  "name": "N1",
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
```

