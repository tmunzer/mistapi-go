
# Rrm Band Metric

*This model accepts additional fields of type interface{}.*

## Structure

`RrmBandMetric`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `CochannelNeighbors` | `float64` | Required | average number of co-channel neighbors |
| `Density` | `float64` | Required | defined by how APs can hear from one and another, 0 - 1 (everyone can hear everyone)<br>**Constraints**: `>= 0`, `<= 1` |
| `Interferences` | [`map[string]models.RrmBandMetricInterference`](../../doc/models/rrm-band-metric-interference.md) | Optional | Property key is the channel number |
| `Neighbors` | `float64` | Required | average number of neighbors |
| `Noise` | `float64` | Required | average noise in dBm |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "cochannel_neighbors": 110.94,
  "density": 56.56,
  "interferences": {
    "149": {
      "radar": 0.3,
      "exampleAdditionalProperty": {
        "key1": "val1",
        "key2": "val2"
      }
    },
    "153": {
      "radar": 0.2,
      "exampleAdditionalProperty": {
        "key1": "val1",
        "key2": "val2"
      }
    }
  },
  "neighbors": 217.56,
  "noise": 202.26,
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

