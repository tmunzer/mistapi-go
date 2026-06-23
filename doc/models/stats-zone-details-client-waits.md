
# Stats Zone Details Client Waits

Client wait time right now

## Structure

`StatsZoneDetailsClientWaits`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Avg` | `int` | Required | Average wait time in seconds |
| `Max` | `int` | Required | Longest wait time in seconds |
| `Min` | `int` | Required | Shortest wait time in seconds |
| `P95` | `int` | Required | 95th percentile of all the wait time(s) |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    statsZoneDetailsClientWaits := models.StatsZoneDetailsClientWaits{
        Avg:                  1200,
        Max:                  3610,
        Min:                  600,
        P95:                  2800,
    }

}
```

