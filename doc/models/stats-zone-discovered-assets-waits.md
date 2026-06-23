
# Stats Zone Discovered Assets Waits

Discovered asset wait time right now

## Structure

`StatsZoneDiscoveredAssetsWaits`

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
    statsZoneDiscoveredAssetsWaits := models.StatsZoneDiscoveredAssetsWaits{
        Avg:                  models.ToPointer(float64(0)),
        Max:                  models.ToPointer(float64(0)),
        Min:                  models.ToPointer(float64(0)),
        P95:                  models.ToPointer(float64(0)),
    }

}
```

