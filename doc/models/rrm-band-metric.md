
# Rrm Band Metric

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

## Example (as JSON)

```json
{
  "cochannel_neighbors": 248.02,
  "density": 62.36,
  "interferences": {
    "149": {
      "radar": 0.3
    },
    "153": {
      "radar": 0.2
    }
  },
  "neighbors": 98.64,
  "noise": 172.66
}
```

