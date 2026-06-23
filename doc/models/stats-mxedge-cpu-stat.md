
# Stats Mxedge Cpu Stat

Aggregate and per-core CPU utilization statistics for a Mist Edge

## Structure

`StatsMxedgeCpuStat`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Cpus` | [`map[string]models.CpuStat`](../../doc/models/cpu-stat.md) | Optional | Per-core CPU utilization statistics keyed by CPU name |
| `Idle` | `*int` | Optional | Percentage of Idle, Idle/(Idle + Busy) since last sampling |
| `Interrupt` | `*int` | Optional | Percentage of Interrupt, (Irq + SoftIrq)/(Idle + Busy) since last sampling |
| `System` | `*int` | Optional | Percentage of System, System/(Idle + Busy) since last sampling |
| `Usage` | `*int` | Optional | Percentage of load, Busy/(Idle + Busy) since last sampling |
| `User` | `*int` | Optional | Percentage of User, User/(Idle + Busy) since last sampling |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    statsMxedgeCpuStat := models.StatsMxedgeCpuStat{
        Cpus:                 map[string]models.CpuStat{
            "cpu0": models.CpuStat{
                Idle:                 models.NewOptional(models.ToPointer(float64(89))),
                Interrupt:            models.NewOptional(models.ToPointer(float64(0))),
                LoadAvg:              []float64{
                    float64(8.63),
                },
                System:               models.NewOptional(models.ToPointer(float64(8))),
                Usage:                models.NewOptional(models.ToPointer(float64(10))),
                User:                 models.NewOptional(models.ToPointer(float64(1))),
            },
            "cpu1": models.CpuStat{
                Idle:                 models.NewOptional(models.ToPointer(float64(81))),
                Interrupt:            models.NewOptional(models.ToPointer(float64(0))),
                LoadAvg:              []float64{
                    float64(8.63),
                },
                System:               models.NewOptional(models.ToPointer(float64(4))),
                Usage:                models.NewOptional(models.ToPointer(float64(18))),
                User:                 models.NewOptional(models.ToPointer(float64(13))),
            },
            "cpu2": models.CpuStat{
                Idle:                 models.NewOptional(models.ToPointer(float64(81))),
                Interrupt:            models.NewOptional(models.ToPointer(float64(0))),
                LoadAvg:              []float64{
                    float64(8.63),
                },
                System:               models.NewOptional(models.ToPointer(float64(4))),
                Usage:                models.NewOptional(models.ToPointer(float64(18))),
                User:                 models.NewOptional(models.ToPointer(float64(13))),
            },
            "cpu3": models.CpuStat{
                Idle:                 models.NewOptional(models.ToPointer(float64(2))),
                Interrupt:            models.NewOptional(models.ToPointer(float64(0))),
                LoadAvg:              []float64{
                    float64(8.63),
                },
                System:               models.NewOptional(models.ToPointer(float64(50))),
                Usage:                models.NewOptional(models.ToPointer(float64(97))),
                User:                 models.NewOptional(models.ToPointer(float64(46))),
            },
        },
        Idle:                 models.ToPointer(62),
        Interrupt:            models.ToPointer(0),
        System:               models.ToPointer(17),
        Usage:                models.ToPointer(37),
        User:                 models.ToPointer(19),
    }

}
```

