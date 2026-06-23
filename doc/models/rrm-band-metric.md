
# Rrm Band Metric

Aggregate RRM metrics for a radio band

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

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    rrmBandMetric := models.RrmBandMetric{
        AvgApsPerChannel:              models.ToPointer(float64(59.66)),
        ChannelDistributionUniformity: models.ToPointer(float64(43.3)),
        CochannelNeighbors:            float64(40.06),
        Density:                       float64(1),
        Interferences:                 map[string]models.RrmBandMetricInterference{
            "149": models.RrmBandMetricInterference{
                Radar:                models.ToPointer(float64(0.3)),
            },
            "153": models.RrmBandMetricInterference{
                Radar:                models.ToPointer(float64(0.2)),
            },
        },
        NapsByChannel:                 map[string]float64{
            "key0": float64(123.11),
            "key1": float64(123.12),
        },
        NapsByPower:                   map[string]float64{
            "key0": float64(5.9),
            "key1": float64(5.91),
            "key2": float64(5.92),
        },
        Neighbors:                     float64(146.68),
        Noise:                         float64(131.38),
    }

}
```

