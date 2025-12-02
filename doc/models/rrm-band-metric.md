
# Rrm Band Metric

## Structure

`RrmBandMetric`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AvgApsPerChannel` | `*float64` | Optional | Average number of APs per channel |
| `ChannelDistributionUniformity` | `*float64` | Optional | Distribution of channel across the Access Points |
| `CochannelNeighbors` | `float64` | Required | Average number of co-channel neighbors |
| `Density` | `float64` | Required | defined by how APs can hear from one and another, 0 - 1 (everyone can hear everyone)<br><br>**Constraints**: `>= 0`, `<= 1` |
| `Interferences` | [`map[string]models.RrmBandMetricInterference`](../../doc/models/rrm-band-metric-interference.md) | Optional | Property key is the channel number |
| `NapsByChannel` | `map[string]float64` | Optional | Property key is the channel number, value is number of APs on that channel |
| `NapsByPower` | `map[string]float64` | Optional | Property key is the power level, value is number of APs on that power level |
| `Neighbors` | `float64` | Required | Average number of neighbors |
| `Noise` | `float64` | Required | Average noise in dBm |

## Example (as JSON)

```json
{
  "cochannel_neighbors": 110.94,
  "density": 1.0,
  "interferences": {
    "149": {
      "radar": 0.3
    },
    "153": {
      "radar": 0.2
    }
  },
  "neighbors": 217.56,
  "noise": 202.26,
  "avg_aps_per_channel": 125.46,
  "channel_distribution_uniformity": 141.82,
  "naps_by_channel": {
    "key0": 62.01,
    "key1": 62.0,
    "key2": 61.99
  },
  "naps_by_power": {
    "key0": 179.22
  }
}
```

