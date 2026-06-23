
# Marvis Client Insights

Time-series performance metrics for a Marvis Client device

## Structure

`MarvisClientInsights`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AvgBattery` | `[]float64` | Optional | Average battery level per interval bucket |
| `AvgCellularRssi` | `[]float64` | Optional | Average cellular RSSI per interval bucket, in dBm |
| `AvgCpu` | `[]float64` | Optional | Average CPU utilization per interval bucket (0–100) |
| `AvgMemory` | `[]float64` | Optional | Average memory utilization per interval bucket (0–100) |
| `AvgWifiRssi` | `[]float64` | Optional | Average Wi-Fi RSSI per interval bucket, in dBm |
| `End` | `*int` | Optional | End of the reporting window, in epoch seconds |
| `Interval` | `*int` | Optional | Duration of each interval bucket, in seconds |
| `Limit` | `*int` | Optional | Maximum number of results requested |
| `Page` | `*int` | Optional | Current page number |
| `Rt` | `[]string` | Optional | List of ISO 8601 timestamp strings for each interval bucket |
| `Start` | `*int` | Optional | Start of the reporting window, in epoch seconds |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    marvisClientInsights := models.MarvisClientInsights{
        AvgBattery:           []float64{
            float64(0.8),
        },
        AvgCellularRssi:      []float64{
            float64(48.5),
        },
        AvgCpu:               []float64{
            float64(87.84),
            float64(87.85),
        },
        AvgMemory:            []float64{
            float64(139.04),
        },
        AvgWifiRssi:          []float64{
            float64(206.13),
        },
    }

}
```

