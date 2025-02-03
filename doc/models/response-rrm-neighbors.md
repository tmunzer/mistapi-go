
# Response Rrm Neighbors

*This model accepts additional fields of type interface{}.*

## Structure

`ResponseRrmNeighbors`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `End` | `int` | Required | - |
| `Limit` | `int` | Required | - |
| `Next` | `*string` | Optional | Link to query next set of results. value is null if no next page exists. |
| `Results` | [`[]models.RrmNeighbors`](../../doc/models/rrm-neighbors.md) | Required | **Constraints**: *Unique Items Required* |
| `Start` | `int` | Required | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "end": 246,
  "limit": 180,
  "results": [
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
        }
      ],
      "exampleAdditionalProperty": {
        "key1": "val1",
        "key2": "val2"
      }
    }
  ],
  "start": 204,
  "next": "next6",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

