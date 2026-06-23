
# Sle Thresholds

Site SLE threshold overrides for capacity, coverage, throughput, and time to connect

## Structure

`SleThresholds`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Capacity` | `*int` | Optional | Threshold percentage for capacity SLE scoring<br><br>**Default**: `20`<br><br>**Constraints**: `>= 5`, `<= 50` |
| `Coverage` | `*int` | Optional | RSSI threshold for coverage SLE scoring, in dBm<br><br>**Default**: `-72`<br><br>**Constraints**: `>= -90`, `<= -60` |
| `Throughput` | `*int` | Optional | Minimum throughput threshold for SLE scoring, in Mbps<br><br>**Default**: `10`<br><br>**Constraints**: `>= 1`, `<= 100` |
| `TimeToConnect` | `*int` | Optional | Time to connect, in seconds<br><br>**Default**: `4`<br><br>**Constraints**: `>= 2`, `<= 10` |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    sleThresholds := models.SleThresholds{
        Capacity:             models.ToPointer(20),
        Coverage:             models.ToPointer(-72),
        Throughput:           models.ToPointer(10),
        TimeToConnect:        models.ToPointer(4),
    }

}
```

