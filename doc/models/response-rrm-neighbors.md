
# Response Rrm Neighbors

## Structure

`ResponseRrmNeighbors`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `End` | `int` | Required | - |
| `Limit` | `int` | Required | - |
| `Next` | `*string` | Optional | the link to query next set of results. value is null if no next page exists. |
| `Results` | [`[]models.RrmNeighbors`](../../doc/models/rrm-neighbors.md) | Required | **Constraints**: *Unique Items Required* |
| `Start` | `int` | Required | - |

## Example (as JSON)

```json
{
  "end": 110,
  "limit": 196,
  "results": [
    {
      "mac": "5c5b35000001",
      "neighbors": [
        {
          "mac": "mac4",
          "rssi": 56
        }
      ]
    }
  ],
  "start": 68,
  "next": "next4"
}
```

