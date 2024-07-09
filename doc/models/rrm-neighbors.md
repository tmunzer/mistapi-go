
# Rrm Neighbors

## Structure

`RrmNeighbors`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Mac` | `*string` | Optional | - |
| `Neighbors` | [`[]models.RrmNeighborsNeighbor`](../../doc/models/rrm-neighbors-neighbor.md) | Optional | - |

## Example (as JSON)

```json
{
  "mac": "5c5b35000001",
  "neighbors": [
    {
      "mac": "mac4",
      "rssi": 56
    },
    {
      "mac": "mac4",
      "rssi": 56
    }
  ]
}
```

