
# Cpu Stat

CPU utilization breakdown for a device

## Structure

`CpuStat`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Idle` | `models.Optional[float64]` | Optional, Read-only | Percentage of CPU time that is idle |
| `Interrupt` | `models.Optional[float64]` | Optional, Read-only | Percentage of CPU time being used by interrupts |
| `LoadAvg` | `[]float64` | Optional | Load averages for the last 1, 5, and 15 minutes |
| `System` | `models.Optional[float64]` | Optional, Read-only | Percentage of CPU time being used by system processes |
| `Usage` | `models.Optional[float64]` | Optional, Read-only | Overall CPU usage percentage |
| `User` | `models.Optional[float64]` | Optional, Read-only | Percentage of CPU time being used by user processes |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    cpuStat := models.CpuStat{
        LoadAvg:              []float64{
            float64(76.35),
            float64(76.36),
            float64(76.37),
        },
    }

}
```

