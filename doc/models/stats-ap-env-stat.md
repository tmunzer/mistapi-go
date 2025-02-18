
# Stats Ap Env Stat

Device environment, including CPU temperature, Ambient temperature, Humidity, Attitude, Pressure, Accelerometers, Magnetometers and vCore Voltage

*This model accepts additional fields of type interface{}.*

## Structure

`StatsApEnvStat`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AccelX` | `models.Optional[float64]` | Optional | - |
| `AccelY` | `models.Optional[float64]` | Optional | - |
| `AccelZ` | `models.Optional[float64]` | Optional | - |
| `AmbientTemp` | `models.Optional[int]` | Optional | - |
| `Attitude` | `models.Optional[int]` | Optional | - |
| `CpuTemp` | `models.Optional[int]` | Optional | - |
| `Humidity` | `models.Optional[int]` | Optional | - |
| `MagneX` | `models.Optional[float64]` | Optional | - |
| `MagneY` | `models.Optional[float64]` | Optional | - |
| `MagneZ` | `models.Optional[float64]` | Optional | - |
| `Pressure` | `models.Optional[float64]` | Optional | - |
| `VcoreVoltage` | `models.Optional[int]` | Optional | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "accel_x": 0.0,
  "accel_y": 0.032,
  "accel_z": -1.088,
  "ambient_temp": 43,
  "attitude": 0,
  "cpu_temp": 61,
  "humidity": 9,
  "magne_x": 0,
  "magne_y": 0,
  "magne_z": 0,
  "pressure": 968,
  "vcore_voltage": 0,
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

