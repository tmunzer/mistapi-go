
# Stats Switch Clients Stats Total

Total AP and wired-client counts for switch client statistics

## Structure

`StatsSwitchClientsStatsTotal`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `NumAps` | `[]int` | Optional | AP count values included in the aggregate switch client statistics |
| `NumWiredClients` | `*int` | Optional | Number of wired clients included in the aggregate switch client statistics |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    statsSwitchClientsStatsTotal := models.StatsSwitchClientsStatsTotal{
        NumAps:               []int{
            245,
        },
        NumWiredClients:      models.ToPointer(188),
    }

}
```

