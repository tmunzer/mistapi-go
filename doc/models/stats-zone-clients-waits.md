
# Stats Zone Clients Waits

Client wait time right now

## Structure

`StatsZoneClientsWaits`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Avg` | `*float64` | Optional | Average wait time in seconds |
| `Max` | `*float64` | Optional | Longest wait time in seconds |
| `Min` | `*float64` | Optional | Shortest wait time in seconds |
| `P95` | `*float64` | Optional | 95th percentile of all the wait time(s) |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    statsZoneClientsWaits := models.StatsZoneClientsWaits{
        Avg:                  models.ToPointer(float64(1200)),
        Max:                  models.ToPointer(float64(3610)),
        Min:                  models.ToPointer(float64(600)),
        P95:                  models.ToPointer(float64(2800)),
    }

}
```

