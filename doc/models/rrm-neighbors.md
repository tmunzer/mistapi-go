
# Rrm Neighbors

*This model accepts additional fields of type interface{}.*

## Structure

`RrmNeighbors`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Mac` | `*string` | Optional | - |
| `Neighbors` | [`[]models.RrmNeighborsNeighbor`](../../doc/models/rrm-neighbors-neighbor.md) | Optional | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "mac": "5c5b35000001",
  "neighbors": [
    {
      "mac": "mac4",
      "rssi": 56,
      "exampleAdditionalProperty": {
        "key1": "val1",
        "key2": "val2"
      }
    },
    {
      "mac": "mac4",
      "rssi": 56,
      "exampleAdditionalProperty": {
        "key1": "val1",
        "key2": "val2"
      }
    },
    {
      "mac": "mac4",
      "rssi": 56,
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

