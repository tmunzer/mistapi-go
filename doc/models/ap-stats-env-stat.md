
# Ap Stats Env Stat

device environment, including CPU temperature, Ambient temperature, Humidity, Attitude, Pressure, Accelerometers, Magnetometers and vCore Voltage

## Structure

`ApStatsEnvStat`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AccelX` | `*float64` | Optional | - |
| `AccelY` | `*float64` | Optional | - |
| `AccelZ` | `*float64` | Optional | - |
| `AmbientTemp` | `*int` | Optional | - |
| `Attitude` | `*int` | Optional | - |
| `CpuTemp` | `*int` | Optional | - |
| `Humidity` | `*int` | Optional | - |
| `MagneX` | `*float64` | Optional | - |
| `MagneY` | `*float64` | Optional | - |
| `MagneZ` | `*float64` | Optional | - |
| `Pressure` | `*int` | Optional | - |
| `VcoreVoltage` | `*float64` | Optional | - |

## Example (as JSON)

```json
{
  "accel_x": 184.92,
  "accel_y": 28.22,
  "accel_z": 174.96,
  "ambient_temp": 252,
  "attitude": 174
}
```

