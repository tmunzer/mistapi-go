
# Cpu Stat

## Structure

`CpuStat`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Idle` | `models.Optional[float64]` | Optional | Percentage of CPU time that is idle |
| `Interrupt` | `models.Optional[float64]` | Optional | Percentage of CPU time being used by interrupts |
| `LoadAvg` | `[]float64` | Optional | Load averages for the last 1, 5, and 15 minutes |
| `System` | `models.Optional[float64]` | Optional | Percentage of CPU time being used by system processes |
| `User` | `models.Optional[float64]` | Optional | Percentage of CPU time being used by user processe |

## Example (as JSON)

```json
{
  "idle": 71.24,
  "interrupt": 185.0,
  "load_avg": [
    75.07,
    75.08
  ],
  "system": 211.56,
  "user": 173.68
}
```

