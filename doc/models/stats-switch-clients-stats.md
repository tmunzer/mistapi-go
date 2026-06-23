
# Stats Switch Clients Stats

Aggregate switch client counts

## Structure

`StatsSwitchClientsStats`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Total` | [`*models.StatsSwitchClientsStatsTotal`](../../doc/models/stats-switch-clients-stats-total.md) | Optional | Total AP and wired-client counts for switch client statistics |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    statsSwitchClientsStats := models.StatsSwitchClientsStats{
        Total:                models.ToPointer(models.StatsSwitchClientsStatsTotal{
            NumAps:               []int{
                23,
                22,
                21,
            },
            NumWiredClients:      models.ToPointer(222),
        }),
    }

}
```

