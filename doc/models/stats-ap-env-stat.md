
# Stats Ap Env Stat

Device environment, including CPU temperature, Ambient temperature, Humidity, Attitude, Pressure, Accelerometers, Magnetometers and vCore Voltage

## Structure

`StatsApEnvStat`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AccelX` | `models.Optional[float64]` | Optional, Read-only | X-axis accelerometer reading reported by the AP |
| `AccelY` | `models.Optional[float64]` | Optional, Read-only | Y-axis accelerometer reading reported by the AP |
| `AccelZ` | `models.Optional[float64]` | Optional, Read-only | Z-axis accelerometer reading reported by the AP |
| `AmbientTemp` | `models.Optional[int]` | Optional, Read-only | Temperature reading from the AP ambient sensor |
| `Attitude` | `models.Optional[int]` | Optional, Read-only | Device attitude or orientation reading reported by the AP |
| `CpuTemp` | `models.Optional[int]` | Optional, Read-only | Temperature reading from the AP CPU sensor |
| `Humidity` | `models.Optional[int]` | Optional, Read-only | Relative humidity sensor reading reported by the AP |
| `MagneX` | `models.Optional[float64]` | Optional, Read-only | X-axis magnetometer reading reported by the AP |
| `MagneY` | `models.Optional[float64]` | Optional, Read-only | Y-axis magnetometer reading reported by the AP |
| `MagneZ` | `models.Optional[float64]` | Optional, Read-only | Z-axis magnetometer reading reported by the AP |
| `Pressure` | `models.Optional[float64]` | Optional, Read-only | Barometric pressure sensor reading reported by the AP |
| `VcoreVoltage` | `models.Optional[int]` | Optional, Read-only | Core voltage sensor reading reported by the AP |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    statsApEnvStat := models.StatsApEnvStat{
        AccelX:               models.NewOptional(models.ToPointer(float64(0))),
        AccelY:               models.NewOptional(models.ToPointer(float64(0.032))),
        AccelZ:               models.NewOptional(models.ToPointer(float64(-1.088))),
        AmbientTemp:          models.NewOptional(models.ToPointer(43)),
        Attitude:             models.NewOptional(models.ToPointer(0)),
        CpuTemp:              models.NewOptional(models.ToPointer(61)),
        Humidity:             models.NewOptional(models.ToPointer(9)),
        MagneX:               models.NewOptional(models.ToPointer(float64(0))),
        MagneY:               models.NewOptional(models.ToPointer(float64(0))),
        MagneZ:               models.NewOptional(models.ToPointer(float64(0))),
        Pressure:             models.NewOptional(models.ToPointer(float64(968))),
        VcoreVoltage:         models.NewOptional(models.ToPointer(0)),
    }

}
```

