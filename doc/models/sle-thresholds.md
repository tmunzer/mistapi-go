
# Sle Thresholds

## Structure

`SleThresholds`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Capacity` | `*int` | Optional | Capacity, in %<br><br>**Default**: `20`<br><br>**Constraints**: `>= 5`, `<= 50` |
| `Coverage` | `*int` | Optional | Coverage, in dBm<br><br>**Default**: `-72`<br><br>**Constraints**: `>= -90`, `<= -60` |
| `Throughput` | `*int` | Optional | Throughput, in Mbps<br><br>**Default**: `10`<br><br>**Constraints**: `>= 1`, `<= 100` |
| `TimeToConnect` | `*int` | Optional | Time to connect, in seconds<br><br>**Default**: `4`<br><br>**Constraints**: `>= 2`, `<= 10` |

## Example (as JSON)

```json
{
  "capacity": 20,
  "coverage": -72,
  "throughput": 10,
  "time-to-connect": 4
}
```

