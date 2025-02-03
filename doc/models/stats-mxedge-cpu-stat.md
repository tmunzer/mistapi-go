
# Stats Mxedge Cpu Stat

CPU/core stats list

*This model accepts additional fields of type interface{}.*

## Structure

`StatsMxedgeCpuStat`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Cpus` | [`map[string]models.CpuStat`](../../doc/models/cpu-stat.md) | Optional | - |
| `Idle` | `*int` | Optional | Percentage of Idle, Idle/(Idle + Busy) since last sampling |
| `Interrupt` | `*int` | Optional | Percentage of Interrupt, (Irq + SoftIrq)/(Idle + Busy) since last sampling |
| `System` | `*int` | Optional | Percentage of System, System/(Idle + Busy) since last sampling |
| `Usage` | `*int` | Optional | Percentage of load, Busy/(Idle + Busy) since last sampling |
| `User` | `*int` | Optional | Percentage of User, User/(Idle + Busy) since last sampling |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "cpus": {
    "cpu0": {
      "idle": 89.0,
      "interrupt": 0.0,
      "system": 8.0,
      "usage": 10,
      "user": 1.0,
      "load_avg": [
        8.63
      ]
    },
    "cpu1": {
      "idle": 81.0,
      "interrupt": 0.0,
      "system": 4.0,
      "usage": 18,
      "user": 13.0,
      "load_avg": [
        8.63
      ]
    },
    "cpu2": {
      "idle": 81.0,
      "interrupt": 0.0,
      "system": 4.0,
      "usage": 18,
      "user": 13.0,
      "load_avg": [
        8.63
      ]
    },
    "cpu3": {
      "idle": 2.0,
      "interrupt": 0.0,
      "system": 50.0,
      "usage": 97,
      "user": 46.0,
      "load_avg": [
        8.63
      ]
    }
  },
  "idle": 62,
  "interrupt": 0,
  "system": 17,
  "usage": 37,
  "user": 19,
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

