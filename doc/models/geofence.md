
# Geofence

*This model accepts additional fields of type interface{}.*

## Structure

`Geofence`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Name` | `*string` | Optional | Name of the geofence |
| `Vertices` | [`[]models.Vertex`](../../doc/models/vertex.md) | Optional | List of vertices defining the geofence |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "name": "example",
  "vertices": [
    {
      "X": 86.66,
      "Y": 252.2,
      "exampleAdditionalProperty": {
        "key1": "val1",
        "key2": "val2"
      }
    },
    {
      "X": 86.66,
      "Y": 252.2,
      "exampleAdditionalProperty": {
        "key1": "val1",
        "key2": "val2"
      }
    },
    {
      "X": 86.66,
      "Y": 252.2,
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

